package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tarot/backend/internal/interfaces/http/handlers"
	"github.com/tarot/backend/internal/interfaces/http/middleware"
)

type Router struct {
	tarot *handlers.TarotHandler
}

func NewRouter(tarot *handlers.TarotHandler) *Router {
	return &Router{tarot: tarot}
}

func (r *Router) Setup(engine *gin.Engine) {
	engine.Use(middleware.CORS())

	api := engine.Group("/api/v1")

	api.GET("/tarot/deck", r.tarot.Deck)

	readings := api.Group("/readings")
	readings.Use(middleware.OptionalAuth())
	{
		readings.POST("/tarot/prepare", r.tarot.Prepare)
		readings.POST("/tarot", r.tarot.Read)
		readings.POST("/tarot/stream", r.tarot.ReadStream)
	}

	// Health check
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
