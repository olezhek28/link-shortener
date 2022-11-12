package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

const (
	dsnEnvName                   = "PG_DSN"
	maxConnectionsEnvName        = "PG_MAX_CONNECTIONS"
	maxConnectionIdleTimeEnvName = "PG_CONNECTION_IDLE_TIME_SEC"

	defaultMaxConnections        = 20
	defaultMaxConnectionIdleTime = 5 * time.Second
)

type PGConfig interface {
	DSN() string
	MaxConnections() int32
	MaxConnectionIdleTime() time.Duration
}

type pgConfig struct {
	dsn                   string
	maxConnections        int32
	maxConnectionIdleTime time.Duration
}

// GetDbConfig ...
func GetDbConfig() (PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("db dsb not found")
	}

	maxConnectionsStr := os.Getenv(maxConnectionsEnvName)
	maxConnections, err := strconv.Atoi(maxConnectionsStr)
	if err != nil || maxConnections == 0 {
		maxConnections = defaultMaxConnections
	}

	maxConnectionIdleTimeStr := os.Getenv(maxConnectionIdleTimeEnvName)
	maxConnectionIdleTimeNum, err := strconv.Atoi(maxConnectionIdleTimeStr)
	maxConnectionIdleTime := time.Duration(maxConnectionIdleTimeNum) * time.Second
	if err != nil || maxConnectionIdleTime == 0 {
		maxConnectionIdleTime = defaultMaxConnectionIdleTime
	}

	return &pgConfig{
		dsn:                   dsn,
		maxConnections:        int32(maxConnections),
		maxConnectionIdleTime: maxConnectionIdleTime,
	}, nil
}

// DSN ...
func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}

// MaxConnections ...
func (cfg *pgConfig) MaxConnections() int32 {
	return cfg.maxConnections
}

// MaxConnectionIdleTime ...
func (cfg *pgConfig) MaxConnectionIdleTime() time.Duration {
	return cfg.maxConnectionIdleTime
}
