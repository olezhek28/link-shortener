package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IClient interface {
	Open(ctx context.Context) error
	Close() error
	DB() *pgxpool.Pool
}

type client struct {
	db     *pgxpool.Pool
	config *pgxpool.Config
}

func NewClient(config *pgxpool.Config) IClient {
	return &client{
		config: config,
	}
}

func (c *client) Open(ctx context.Context) error {
	var err error
	c.db, err = pgxpool.ConnectConfig(ctx, c.config)
	if err != nil {
		return err
	}

	return nil
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
