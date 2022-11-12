package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/olezhek28/link-shortener/internal/logger"
)

func Init(_ context.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Info("No .env file found")
	}

	return nil
}
