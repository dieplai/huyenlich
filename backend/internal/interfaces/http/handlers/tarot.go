package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarot/backend/internal/application/tarot"
)

type TarotHandler struct {
	svc *tarot.Service
}

func NewTarotHandler(svc *tarot.Service) *TarotHandler {
	return &TarotHandler{svc: svc}
}

// GET /api/v1/tarot/deck
func (h *TarotHandler) Deck(c *gin.Context) {
	resp, err := h.svc.Deck(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "deck failed"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// POST /api/v1/readings/tarot/prepare
func (h *TarotHandler) Prepare(c *gin.Context) {
	var req tarot.PrepareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.svc.Prepare(c.Request.Context(), req)
	if err != nil {
		if tarot.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "prepare failed"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// POST /api/v1/readings/tarot
func (h *TarotHandler) Read(c *gin.Context) {
	var req tarot.ReadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.svc.Read(c.Request.Context(), req)
	if err != nil {
		if tarot.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "reading failed"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// POST /api/v1/readings/tarot/stream
func (h *TarotHandler) ReadStream(c *gin.Context) {
	var req tarot.ReadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")
	c.Status(http.StatusOK)

	flusher, _ := c.Writer.(http.Flusher)
	send := func(event string, payload any) error {
		c.SSEvent(event, payload)
		if flusher != nil {
			flusher.Flush()
		}
		return nil
	}

	if err := h.svc.StreamRead(c.Request.Context(), req, tarot.ReadStreamHandlers{
		OnStart: func(resp *tarot.ReadResponse) error {
			return send("start", resp)
		},
		OnDelta: func(delta string) error {
			return send("delta", gin.H{"text": delta})
		},
		OnDone: func(resp *tarot.ReadResponse) error {
			return send("done", resp)
		},
	}); err != nil {
		if tarot.IsValidationError(err) {
			_ = send("error", gin.H{"error": err.Error()})
			return
		}
		_ = send("error", gin.H{"error": "reading failed"})
	}
}
