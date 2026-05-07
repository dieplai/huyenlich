package tarot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tarot/backend/internal/application/ai"
	"github.com/tarot/backend/internal/domain/reading"
)

type rawCard struct {
	ID         string   `json:"id"`
	DeckIndex  int      `json:"deck_index"`
	NameVI     string   `json:"name_vi"`
	Arcana     string   `json:"arcana"`
	Number     int      `json:"number"`
	Suit       *string  `json:"suit"`
	Rank       *string  `json:"rank"`
	Element    *string  `json:"element"`
	KeywordsVI []string `json:"keywords_vi"`
	UprightVI  string   `json:"upright_vi"`
	ReversedVI string   `json:"reversed_vi"`
	Image      struct {
		LocalPath string `json:"local_path"`
	} `json:"image"`
}

type rawDeck struct {
	Cards []rawCard `json:"cards"`
}

const (
	aiStatusReady    = "ready"
	aiStatusFallback = "fallback"
)

type Service struct {
	readingRepo reading.Repository
	ai          *ai.Service
	cards       map[string]rawCard
	deckOrder   []string
	spreads     map[string]SpreadInfo
	analyzer    questionAnalyzer
}

type ReadStreamHandlers struct {
	OnStart func(*ReadResponse) error
	OnDelta func(string) error
	OnDone  func(*ReadResponse) error
}

func NewService(repo reading.Repository, aiSvc *ai.Service, dataDir string) (*Service, error) {
	svc := &Service{
		readingRepo: repo,
		ai:          aiSvc,
		cards:       make(map[string]rawCard),
		deckOrder:   make([]string, 0, 78),
		spreads:     defaultSpreadRegistry(),
		analyzer:    questionAnalyzer{},
	}
	if err := svc.loadDeck(dataDir); err != nil {
		return nil, err
	}
	return svc, nil
}

func (s *Service) loadDeck(dataDir string) error {
	path, data, err := readDeckFile(dataDir)
	if err != nil {
		return err
	}

	var deck rawDeck
	if err := json.Unmarshal(data, &deck); err != nil {
		return fmt.Errorf("parse tarot deck %s: %w", path, err)
	}
	if len(deck.Cards) == 0 {
		return fmt.Errorf("tarot deck %s has no cards", path)
	}

	for _, card := range deck.Cards {
		if card.ID == "" {
			return fmt.Errorf("tarot deck contains a card without id")
		}
		s.cards[card.ID] = card
		s.deckOrder = append(s.deckOrder, card.ID)
	}
	return nil
}

func readDeckFile(dataDir string) (string, []byte, error) {
	candidates := []string{dataDir}
	if dataDir != "data" {
		candidates = append(candidates, "data")
	}
	if dataDir != filepath.Join("..", "data") {
		candidates = append(candidates, filepath.Join("..", "data"))
	}

	var lastErr error
	for _, dir := range candidates {
		path := filepath.Join(dir, "tarot", "rider-waite-smith.json")
		data, err := os.ReadFile(path)
		if err == nil {
			return path, data, nil
		}
		lastErr = err
	}
	return "", nil, fmt.Errorf("load tarot deck from %v: %w", candidates, lastErr)
}

func (s *Service) Prepare(ctx context.Context, req PrepareRequest) (*PrepareResponse, error) {
	_ = ctx
	if strings.TrimSpace(req.Question) == "" {
		return nil, validationError("question is required")
	}
	analysis, aiPreferredCount := s.analyzeQuestion(ctx, req.Question)
	preferredCount := req.PreferredCardCount
	if preferredCount == 0 {
		preferredCount = aiPreferredCount
	}
	spread := s.recommendSpread(analysis, preferredCount)
	return &PrepareResponse{Analysis: analysis, Spread: spread}, nil
}

func (s *Service) Read(ctx context.Context, req ReadRequest) (*ReadResponse, error) {
	spread, free, fallbackContent, inputJSON, freeJSON, err := s.buildReadArtifacts(ctx, req)
	if err != nil {
		return nil, err
	}

	status := aiStatusFallback
	content, fromAI := s.writeReading(ctx, free, spread, fallbackContent)
	if fromAI {
		status = aiStatusReady
	}

	r := reading.New(reading.ServiceTarot, inputJSON, freeJSON)
	r.AIContent = &content
	if err := s.readingRepo.Save(ctx, r); err != nil {
		return nil, err
	}

	return &ReadResponse{
		ReadingID:  r.ID.String(),
		Spread:     spread,
		Free:       free,
		AIContent:  &content,
		AIStatus:   status,
		IsUnlocked: false,
	}, nil
}

func (s *Service) StreamRead(ctx context.Context, req ReadRequest, handlers ReadStreamHandlers) error {
	spread, free, fallbackContent, inputJSON, freeJSON, err := s.buildReadArtifacts(ctx, req)
	if err != nil {
		return err
	}

	emptyContent := ""
	r := reading.New(reading.ServiceTarot, inputJSON, freeJSON)
	r.AIContent = &emptyContent
	if err := s.readingRepo.Save(ctx, r); err != nil {
		return err
	}

	resp := &ReadResponse{
		ReadingID:  r.ID.String(),
		Spread:     spread,
		Free:       free,
		AIContent:  &emptyContent,
		AIStatus:   aiStatusFallback,
		IsUnlocked: false,
	}
	if handlers.OnStart != nil {
		if err := handlers.OnStart(resp); err != nil {
			return err
		}
	}

	status := aiStatusFallback
	content := ""
	if s.ai != nil && s.ai.Enabled() {
		streamed, err := s.ai.LongFormStream(ctx, buildReadingWriterPrompt(free, spread), func(delta string) error {
			if handlers.OnDelta == nil {
				return nil
			}
			return handlers.OnDelta(delta)
		})
		switch {
		case err == nil && strings.TrimSpace(streamed) != "":
			status = aiStatusReady
			content = normalizeAIReadingText(streamed)
		case strings.TrimSpace(streamed) != "":
			log.Printf("tarot ai stream partial fallback: %v", err)
			content = normalizeAIReadingText(streamed)
		default:
			if err != nil {
				log.Printf("tarot ai stream fallback: %v", err)
			}
			content, status = s.writeAndEmitFallback(ctx, free, spread, fallbackContent, handlers.OnDelta)
		}
	} else {
		content = fallbackContent
		if err := emitTextChunks(ctx, content, handlers.OnDelta); err != nil {
			return err
		}
	}

	r.AIContent = &content
	if err := s.readingRepo.Update(ctx, r); err != nil {
		return err
	}

	resp.AIContent = &content
	resp.AIStatus = status
	if handlers.OnDone != nil {
		return handlers.OnDone(resp)
	}
	return nil
}

func (s *Service) buildReadArtifacts(ctx context.Context, req ReadRequest) (SpreadInfo, FreeResult, string, []byte, []byte, error) {
	if err := s.validateReadRequest(req); err != nil {
		return SpreadInfo{}, FreeResult{}, "", nil, nil, err
	}

	analysis := s.resolveRequestAnalysis(ctx, req)
	spread, err := s.resolveSpread(req.SpreadID, analysis, len(req.CardIDs))
	if err != nil {
		return SpreadInfo{}, FreeResult{}, "", nil, nil, err
	}

	free, fallbackContent, err := s.buildReading(req.Question, analysis, spread, req.CardIDs, req.Orientations)
	if err != nil {
		return SpreadInfo{}, FreeResult{}, "", nil, nil, err
	}

	freeJSON, err := json.Marshal(free)
	if err != nil {
		return SpreadInfo{}, FreeResult{}, "", nil, nil, err
	}
	inputJSON, err := json.Marshal(req)
	if err != nil {
		return SpreadInfo{}, FreeResult{}, "", nil, nil, err
	}

	return spread, free, fallbackContent, inputJSON, freeJSON, nil
}

func (s *Service) writeAndEmitFallback(ctx context.Context, free FreeResult, spread SpreadInfo, fallbackContent string, onDelta func(string) error) (string, string) {
	content, fromAI := s.writeReading(ctx, free, spread, fallbackContent)
	status := aiStatusFallback
	if fromAI {
		status = aiStatusReady
	}
	if err := emitTextChunks(ctx, content, onDelta); err != nil {
		log.Printf("tarot fallback stream interrupted: %v", err)
	}
	return content, status
}

func emitTextChunks(ctx context.Context, text string, onDelta func(string) error) error {
	if onDelta == nil || text == "" {
		return nil
	}
	runes := []rune(text)
	const chunkSize = 22
	for start := 0; start < len(runes); start += chunkSize {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		end := start + chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		if err := onDelta(string(runes[start:end])); err != nil {
			return err
		}
		time.Sleep(12 * time.Millisecond)
	}
	return nil
}

func (s *Service) resolveRequestAnalysis(ctx context.Context, req ReadRequest) QuestionAnalysis {
	fallback := s.analyzer.Analyze(req.Question)
	if req.Analysis != nil {
		return mergeAnalysis(fallback, *req.Analysis)
	}
	if req.SpreadID != "" {
		return fallback
	}
	analysis, _ := s.analyzeQuestion(ctx, req.Question)
	return analysis
}

func (s *Service) analyzeQuestion(ctx context.Context, question string) (QuestionAnalysis, int) {
	fallback := s.analyzer.Analyze(question)
	if s.ai == nil || !s.ai.Enabled() {
		return fallback, 0
	}

	raw, err := s.ai.FastJSON(ctx, buildQuestionAnalysisPrompt(question))
	if err != nil {
		log.Printf("tarot ai prepare fallback: %v", err)
		return fallback, 0
	}
	analysis, preferredCount, err := parseAIQuestionAnalysis(raw, fallback)
	if err != nil {
		log.Printf("tarot ai prepare parse fallback: %v", err)
		return fallback, 0
	}
	return analysis, preferredCount
}

func mergeAnalysis(fallback QuestionAnalysis, incoming QuestionAnalysis) QuestionAnalysis {
	result := fallback
	if isAllowed(incoming.Topic, []string{topicGeneral, topicLove, topicCareer, topicMoney, topicInnerWork}) {
		result.Topic = incoming.Topic
	}
	if isAllowed(incoming.Intent, []string{intentClarity, intentDecision, intentTimeline, intentAdvice}) {
		result.Intent = incoming.Intent
	}
	if len(incoming.EmotionalTone) > 0 {
		result.EmotionalTone = incoming.EmotionalTone
	}
	if len(incoming.Signals) > 0 {
		result.Signals = incoming.Signals
	}
	if len(incoming.SafetyFlags) > 0 {
		result.SafetyFlags = incoming.SafetyFlags
	}
	if strings.TrimSpace(incoming.RewrittenQuestion) != "" {
		result.RewrittenQuestion = incoming.RewrittenQuestion
	}
	if incoming.Confidence > 0 {
		result.Confidence = incoming.Confidence
	}
	if incoming.Source != "" {
		result.Source = incoming.Source
	}
	return result
}

func (s *Service) writeReading(ctx context.Context, free FreeResult, spread SpreadInfo, fallback string) (string, bool) {
	if s.ai == nil || !s.ai.Enabled() {
		return fallback, false
	}
	content, err := s.ai.LongForm(ctx, buildReadingWriterPrompt(free, spread))
	if err != nil || strings.TrimSpace(content) == "" {
		if err != nil {
			log.Printf("tarot ai writer fallback: %v", err)
		}
		return fallback, false
	}
	return normalizeAIReadingText(content), true
}

func normalizeAIReadingText(content string) string {
	text := strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		cleaned := strings.TrimSpace(line)
		if cleaned == "---" || cleaned == "----" {
			lines[i] = ""
			continue
		}
		cleaned = strings.TrimPrefix(cleaned, "### ")
		cleaned = strings.TrimPrefix(cleaned, "## ")
		cleaned = strings.TrimPrefix(cleaned, "# ")
		cleaned = strings.ReplaceAll(cleaned, "**", "")
		cleaned = strings.ReplaceAll(cleaned, "__", "")
		lines[i] = cleaned
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func (s *Service) Deck(ctx context.Context) (*DeckResponse, error) {
	_ = ctx
	cards := make([]DeckCard, 0, len(s.deckOrder))
	for _, id := range s.deckOrder {
		card := s.cards[id]
		cards = append(cards, DeckCard{
			ID:         card.ID,
			NameVI:     card.NameVI,
			Arcana:     card.Arcana,
			Suit:       stringPtrValue(card.Suit),
			KeywordsVI: card.KeywordsVI,
			UprightVI:  card.UprightVI,
			ReversedVI: card.ReversedVI,
			ImagePath:  "/" + filepath.ToSlash(trimPublicPrefix(card.Image.LocalPath)),
		})
	}
	spreads := []SpreadInfo{s.spreads[spreadFiveCardCross]}

	return &DeckResponse{Cards: cards, Spreads: spreads}, nil
}

func (s *Service) validateReadRequest(req ReadRequest) error {
	if strings.TrimSpace(req.Question) == "" {
		return validationError("question is required")
	}
	if len(req.CardIDs) != 5 {
		return validationError("tarot reading currently supports exactly 5 cards")
	}
	if len(req.Orientations) > 0 && len(req.Orientations) != len(req.CardIDs) {
		return validationError("orientations must match card_ids length")
	}

	seen := make(map[string]bool, len(req.CardIDs))
	for _, id := range req.CardIDs {
		if id == "" {
			return validationError("card_ids cannot contain empty values")
		}
		if seen[id] {
			return validationError("card_ids cannot contain duplicate cards")
		}
		seen[id] = true
		if _, ok := s.cards[id]; !ok {
			return validationError(fmt.Sprintf("unknown tarot card: %s", id))
		}
	}
	return nil
}

func (s *Service) resolveSpread(spreadID string, analysis QuestionAnalysis, cardCount int) (SpreadInfo, error) {
	_ = analysis
	if spreadID != "" {
		spread, ok := s.spreads[spreadID]
		if !ok {
			return SpreadInfo{}, validationError(fmt.Sprintf("unknown spread_id: %s", spreadID))
		}
		if spread.ID != spreadFiveCardCross {
			return SpreadInfo{}, validationError("tarot reading currently supports only the 5-card cross spread")
		}
		if spread.CardCount != cardCount {
			return SpreadInfo{}, validationError(fmt.Sprintf("spread %s requires %d cards", spreadID, spread.CardCount))
		}
		return spread, nil
	}

	return s.recommendSpread(analysis, cardCount), nil
}

func (s *Service) recommendSpread(analysis QuestionAnalysis, preferredCardCount int) SpreadInfo {
	_, _ = analysis, preferredCardCount
	return s.spreads[spreadFiveCardCross]
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func trimPublicPrefix(path string) string {
	const prefix = "public/"
	if len(path) >= len(prefix) && path[:len(prefix)] == prefix {
		return path[len(prefix):]
	}
	return path
}
