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
	DbDsn                string
	DbMaxOpenConnects    int
	DbTimeout            time.Duration
	DbMaxIdleConnections int
}

// GetDbConfig ...
func GetDbConfig() *Config {
	password := "sample_pass"
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password={password} dbname=%s sslmode=%s",
		"localhost", "5445", "postgres", "sample_db", "disable")

	dbDsn := strings.ReplaceAll(DbDsn, dbPassEscSeq, password)

	return &Config{
		DbDsn:                dbDsn,
		DbMaxOpenConnects:    20,
		DbTimeout:            3 * time.Second,
		DbMaxIdleConnections: 2,
	}
}
