package config

import (
	"github.com/joho/godotenv"
	"github.com/olezhek28/link-shortener/internal/logger"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Info("No .env file found")
	}
}
