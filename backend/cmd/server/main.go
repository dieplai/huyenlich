package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tarot/backend/configs"
	appAI "github.com/tarot/backend/internal/application/ai"
	"github.com/tarot/backend/internal/application/tarot"
	"github.com/tarot/backend/internal/infrastructure/persistence"
	httpRouter "github.com/tarot/backend/internal/interfaces/http"
	httpHandlers "github.com/tarot/backend/internal/interfaces/http/handlers"
)

func main() {
	configs.Load()

	// Infrastructure
	readingRepo := persistence.NewMemoryReadingRepo()
	aiSvc := appAI.NewService()

	// Application services
	tarotSvc, err := tarot.NewService(readingRepo, aiSvc, configs.C.DataDir)
	if err != nil {
		log.Fatal(err)
	}

	// HTTP handlers
	tarotH := httpHandlers.NewTarotHandler(tarotSvc)

	// Router
	engine := gin.Default()
	router := httpRouter.NewRouter(tarotH)
	router.Setup(engine)

	log.Printf("Server starting on :%s", configs.C.ServerPort)
	if err := engine.Run(":" + configs.C.ServerPort); err != nil {
		log.Fatal(err)
	}
}
