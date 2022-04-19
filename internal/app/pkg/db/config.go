package db

import (
	"fmt"
	"strings"
	"time"
)

const (
	dbPassEscSeq = "{password}"
)

type Config struct {
	DbDsn           string
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	MaxIdleConns    int
}

// GetDbConfig ...
func GetDbConfig() *Config {
	password := "sample_pass"
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password={password} dbname=%s sslmode=%s",
		"localhost", "5445", "postgres", "sample_db", "disable")

	dbDsn := strings.ReplaceAll(DbDsn, dbPassEscSeq, password)

	return &Config{
		DbDsn:           dbDsn,
		MaxOpenConns:    20,
		ConnMaxIdleTime: 3 * time.Second,
		MaxIdleConns:    2,
	}
}
