package db

import (
	"fmt"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	//_ "github.com/lib/pq"
)

type IClient interface {
	Open() error
	Close() error
	DB() *sqlx.DB
}

type client struct {
	db     *sqlx.DB
	config *Config
}

func NewClient(config *Config) IClient {
	return &client{
		config: config,
	}
}

func (c *client) Open() error {
	var err error
	c.db, err = sqlx.Open("pgx", c.config.DbDsn)
	if err != nil {
		return fmt.Errorf("failed to opening connection to db: %s", err.Error())
	}

	c.db.SetMaxOpenConns(c.config.MaxOpenConns)
	c.db.SetConnMaxIdleTime(c.config.ConnMaxIdleTime)
	c.db.SetMaxIdleConns(c.config.MaxIdleConns)

	return nil
}

func (c *client) Close() error {
	if c != nil && c.db != nil {
		err := c.db.Close()
		if err != nil {
			return fmt.Errorf("error occurred while closing db connections: %s", err.Error())
		}
	}

	return nil
}

func (c *client) DB() *sqlx.DB {
	return c.db
}
