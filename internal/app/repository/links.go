package repository

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_links_repository.go -package=mocks . ILinks

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/olezhek28/link-shortener/internal/app/pkg/db"
	"github.com/olezhek28/link-shortener/internal/app/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ILinks interface {
	AddLink(ctx context.Context, longLink string, shortLink string) error
	GetLongLink(ctx context.Context, shortLink string) (string, error)
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

// AddLink ...
func (r *linksRepository) AddLink(ctx context.Context, longLink string, shortLink string) error {
	builder := sq.Insert(table.Links).
		PlaceholderFormat(sq.Dollar).
		Columns("long_link", "short_link").
		Values(longLink, shortLink)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.DB().Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
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
