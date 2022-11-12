package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/olezhek28/link-shortener/internal/config"
)

type Client interface {
	Close() error
	DB() *pgxpool.Pool
}

type client struct {
	db *pgxpool.Pool
}

func NewClient(ctx context.Context, config config.PGConfig) (Client, error) {
	poolConfig, err := pgxpool.ParseConfig(config.DSN())
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = config.MaxConnections()
	poolConfig.MaxConnIdleTime = config.MaxConnectionIdleTime()
	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true

	db, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return &client{
		db: db,
	}, nil
}

func (c *client) Close() error {
	if c.db != nil {
		c.db.Close()
	}

	return nil
}

func (c *client) DB() *pgxpool.Pool {
	return c.db
}
