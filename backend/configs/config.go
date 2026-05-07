package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string
	DatabaseURL   string
	RedisURL      string
	JWTSecret     string
	AIBaseURL     string
	AIAPIKey      string
	AIModel       string
	AIFastModel   string
	AIWriterModel string
	DataDir       string
}

var C *Config

func Load() {
	_ = godotenv.Load()
	_ = godotenv.Load("backend/.env")

	C = &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/tarot?sslmode=disable"),
		RedisURL:      getEnv("REDIS_URL", "redis://localhost:6379/0"),
		JWTSecret:     getEnv("JWT_SECRET", "change-me-in-production"),
		AIBaseURL:     getEnv("AI_BASE_URL", "https://ckey.vn/v1"),
		AIAPIKey:      getEnv("AI_API_KEY", ""),
		AIModel:       getEnv("AI_MODEL", "gpt-5.5"),
		AIFastModel:   getEnv("AI_FAST_MODEL", getEnv("AI_MODEL", "gpt-5.4-mini")),
		AIWriterModel: getEnv("AI_WRITER_MODEL", getEnv("AI_MODEL", "gpt-5.5")),
		DataDir:       getEnv("DATA_DIR", "../data"),
	}

	if C.AIAPIKey == "" {
		log.Println("WARNING: AI_API_KEY not set")
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
