package repository

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_links_repository.go -package=mocks . ILinks

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/olezhek28/link-shortener/internal/pkg/db"
	"github.com/olezhek28/link-shortener/internal/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ILinks interface {
	GetLongLink(ctx context.Context, shortLink string) (string, error)
	GetShortLink(ctx context.Context, longLink string) (string, error)
}

type linksRepository struct {
	db db.IClient
}

// NewLinksRepository ...
func NewLinksRepository(ctx context.Context, db db.IClient) ILinks {
	return &linksRepository{
		db: db,
	}
}

// GetLongLink ...
func (r *linksRepository) GetLongLink(ctx context.Context, shortLink string) (string, error) {
	builder := sq.Select("long_link").
		PlaceholderFormat(sq.Dollar).
		From(table.Links).
		Where(sq.Eq{"short_link": shortLink}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	var longLink string
	err = r.db.DB().QueryRow(ctx, query, args...).Scan(&longLink)
	if err == pgx.ErrNoRows {
		return "", status.Error(codes.NotFound, "long link not found")
	}
	if err != nil {
		return "", err
	}

	return longLink, nil
}

// GetShortLink ...
func (r *linksRepository) GetShortLink(ctx context.Context, longLink string) (string, error) {
	builder := sq.Select("short_link").
		PlaceholderFormat(sq.Dollar).
		From(table.Links).
		Where(sq.Eq{"long_link": longLink}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	var shortLink string
	err = r.db.DB().QueryRow(ctx, query, args...).Scan(&shortLink)
	if err == pgx.ErrNoRows {
		return "", status.Error(codes.NotFound, "short link not found")
	}
	if err != nil {
		return "", err
	}

	return shortLink, nil
}
