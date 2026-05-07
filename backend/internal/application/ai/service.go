package ai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/tarot/backend/configs"
)

type Service struct {
	client      *http.Client
	baseURL     string
	apiKey      string
	fastModel   string
	writerModel string
	enabled     bool
}

type chatRequest struct {
	Model     string        `json:"model"`
	Messages  []chatMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
	Stream    bool          `json:"stream,omitempty"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Choices []struct {
		Message chatMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

type chatStreamChunk struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

type requestError struct {
	model      string
	statusCode int
	message    string
	err        error
}

func (e *requestError) Error() string {
	base := "ai request failed"
	if e.model != "" {
		base += ": model " + e.model
	}
	if e.statusCode > 0 {
		base += fmt.Sprintf(": status %d", e.statusCode)
	}
	if e.message != "" {
		base += ": " + e.message
	}
	if e.err != nil {
		base += ": " + e.err.Error()
	}
	return base
}

func (e *requestError) Unwrap() error {
	return e.err
}

func NewService() *Service {
	apiKey := strings.TrimSpace(configs.C.AIAPIKey)
	return &Service{
		client:      &http.Client{},
		baseURL:     strings.TrimRight(configs.C.AIBaseURL, "/"),
		apiKey:      apiKey,
		fastModel:   configs.C.AIFastModel,
		writerModel: configs.C.AIWriterModel,
		enabled:     apiKey != "",
	}
}

func (s *Service) Enabled() bool {
	return s != nil && s.enabled
}

func (s *Service) FastJSON(ctx context.Context, prompt string) (string, error) {
	return s.chatWithRetry(ctx, s.fastModel, prompt, 360, 14*time.Second, 1)
}

func (s *Service) LongForm(ctx context.Context, prompt string) (string, error) {
	content, err := s.chatWithRetry(ctx, s.writerModel, prompt, 1200, 80*time.Second, 1)
	if err == nil {
		return content, nil
	}
	if strings.TrimSpace(s.fastModel) == "" || s.fastModel == s.writerModel || !isRetryableAIError(err) {
		return "", err
	}

	fallbackContent, fallbackErr := s.chatWithRetry(ctx, s.fastModel, prompt, 900, 45*time.Second, 0)
	if fallbackErr == nil {
		return fallbackContent, nil
	}
	return "", fmt.Errorf("%w; fallback model %s failed: %v", err, s.fastModel, fallbackErr)
}

func (s *Service) LongFormStream(ctx context.Context, prompt string, onDelta func(string) error) (string, error) {
	content, err := s.chatStreamWithRetry(ctx, s.writerModel, prompt, 1200, 80*time.Second, 1, onDelta)
	if err == nil {
		return content, nil
	}
	if strings.TrimSpace(content) != "" {
		return content, err
	}
	if strings.TrimSpace(s.fastModel) == "" || s.fastModel == s.writerModel || !isRetryableAIError(err) {
		return "", err
	}

	fallbackContent, fallbackErr := s.chatStreamWithRetry(ctx, s.fastModel, prompt, 900, 45*time.Second, 0, onDelta)
	if fallbackErr == nil {
		return fallbackContent, nil
	}
	return "", fmt.Errorf("%w; fallback model %s failed: %v", err, s.fastModel, fallbackErr)
}

func (s *Service) chatWithRetry(ctx context.Context, model string, prompt string, maxTokens int, timeout time.Duration, retries int) (string, error) {
	var lastErr error
	for attempt := 0; attempt <= retries; attempt++ {
		attemptCtx, cancel := context.WithTimeout(ctx, timeout)
		content, err := s.chat(attemptCtx, model, prompt, maxTokens)
		cancel()
		if err == nil {
			return content, nil
		}
		lastErr = err
		if attempt == retries || !isRetryableAIError(err) || ctx.Err() != nil {
			break
		}
		time.Sleep(time.Duration(attempt+1) * 600 * time.Millisecond)
	}
	return "", lastErr
}

func (s *Service) chatStreamWithRetry(ctx context.Context, model string, prompt string, maxTokens int, timeout time.Duration, retries int, onDelta func(string) error) (string, error) {
	var lastErr error
	for attempt := 0; attempt <= retries; attempt++ {
		attemptCtx, cancel := context.WithTimeout(ctx, timeout)
		content, err := s.chatStream(attemptCtx, model, prompt, maxTokens, onDelta)
		cancel()
		if err == nil {
			return content, nil
		}
		if strings.TrimSpace(content) != "" {
			return content, err
		}
		lastErr = err
		if attempt == retries || !isRetryableAIError(err) || ctx.Err() != nil {
			break
		}
		time.Sleep(time.Duration(attempt+1) * 600 * time.Millisecond)
	}
	return "", lastErr
}

func (s *Service) chat(ctx context.Context, model string, prompt string, maxTokens int) (string, error) {
	if !s.Enabled() {
		return "", fmt.Errorf("ai service disabled: missing API key")
	}

	payload := chatRequest{
		Model: model,
		Messages: []chatMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: maxTokens,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", &requestError{model: model, err: err}
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var decoded chatResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return "", &requestError{model: model, statusCode: resp.StatusCode, message: strings.TrimSpace(string(data))}
		}
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if decoded.Error != nil && decoded.Error.Message != "" {
			return "", &requestError{model: model, statusCode: resp.StatusCode, message: decoded.Error.Message}
		}
		return "", &requestError{model: model, statusCode: resp.StatusCode}
	}
	if len(decoded.Choices) == 0 {
		return "", fmt.Errorf("no response from AI")
	}
	return strings.TrimSpace(decoded.Choices[0].Message.Content), nil
}

func (s *Service) chatStream(ctx context.Context, model string, prompt string, maxTokens int, onDelta func(string) error) (string, error) {
	if !s.Enabled() {
		return "", fmt.Errorf("ai service disabled: missing API key")
	}

	payload := chatRequest{
		Model: model,
		Messages: []chatMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: maxTokens,
		Stream:    true,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", &requestError{model: model, err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return "", readErr
		}
		var decoded chatResponse
		if err := json.Unmarshal(data, &decoded); err == nil && decoded.Error != nil && decoded.Error.Message != "" {
			return "", &requestError{model: model, statusCode: resp.StatusCode, message: decoded.Error.Message}
		}
		return "", &requestError{model: model, statusCode: resp.StatusCode, message: strings.TrimSpace(string(data))}
	}

	var builder strings.Builder
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 1024), 1024*1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ":") || !strings.HasPrefix(line, "data:") {
			continue
		}

		data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if data == "[DONE]" {
			break
		}

		var chunk chatStreamChunk
		if err := json.Unmarshal([]byte(data), &chunk); err != nil {
			return builder.String(), err
		}
		if chunk.Error != nil && chunk.Error.Message != "" {
			return builder.String(), &requestError{model: model, message: chunk.Error.Message}
		}

		for _, choice := range chunk.Choices {
			if choice.Delta.Content == "" {
				continue
			}
			builder.WriteString(choice.Delta.Content)
			if onDelta != nil {
				if err := onDelta(choice.Delta.Content); err != nil {
					return builder.String(), err
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return builder.String(), err
	}
	if strings.TrimSpace(builder.String()) == "" {
		return "", fmt.Errorf("no streamed response from AI")
	}
	return strings.TrimSpace(builder.String()), nil
}

func isRetryableAIError(err error) bool {
	var reqErr *requestError
	if errors.As(err, &reqErr) {
		if reqErr.statusCode == http.StatusTooManyRequests || reqErr.statusCode == http.StatusRequestTimeout {
			return true
		}
		if reqErr.statusCode >= 500 && reqErr.statusCode <= 599 {
			return true
		}
	}
	return errors.Is(err, context.DeadlineExceeded)
}
