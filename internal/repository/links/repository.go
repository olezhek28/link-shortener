package links

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_links_repository.go -package=mocks . Repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/olezhek28/link-shortener/internal/pkg/db"
	"github.com/olezhek28/link-shortener/internal/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository interface {
	AddLink(ctx context.Context, longLink string, shortLink string) error
	GetLongLink(ctx context.Context, shortLink string) (string, error)
}

type repository struct {
	db db.Client
}

// NewLinksRepository ...
func NewLinksRepository(ctx context.Context, db db.Client) Repository {
	return &repository{
		db: db,
	}
}

// AddLink ...
func (r *repository) AddLink(ctx context.Context, longLink string, shortLink string) error {
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
func (r *repository) GetLongLink(ctx context.Context, shortLink string) (string, error) {
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
