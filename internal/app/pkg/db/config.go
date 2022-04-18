package db

import (
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
	password := ""
	DbDsn := ""
	dbDsn := strings.ReplaceAll(DbDsn, dbPassEscSeq, password)

	return &Config{
		DbDsn:                dbDsn,
		DbMaxOpenConnects:    20,
		DbTimeout:            3 * time.Second,
		DbMaxIdleConnections: 2,
	}
}
