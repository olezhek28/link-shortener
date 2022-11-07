package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEscSeq    = "{password}"
	postgresEnvName = "POSTGRES_DSN"
)

type Config struct {
	DbDsn           string
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	MaxIdleConns    int
}

// GetDbConfig ...
func GetDbConfig() (*pgxpool.Config, error) {
	password := "jw8s0F4"

	dbDsn := "postgresql://links:{password}@localhost:6432/links?sslmode=disable&pool_max_conns=20" // os.Getenv(postgresEnvName)
	if len(dbDsn) == 0 {
		return nil, fmt.Errorf("DB params not found")
	}

	dbDsn = strings.ReplaceAll(dbDsn, dbPassEscSeq, password)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true

	return poolConfig, nil
}
