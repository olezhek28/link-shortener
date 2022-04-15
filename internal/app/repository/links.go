package repository

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_links_repository.go -package=mocks . ILinks

import "context"

type ILinks interface {
	GetLongLink(ctx context.Context, shortLink string) (string, error)
	GetShortLink(ctx context.Context, longLink string) (string, error)
}

type linksRepository struct {
}

func NewLinksRepository() ILinks {
	return &linksRepository{}
}

func (r *linksRepository) GetLongLink(ctx context.Context, shortLink string) (string, error) {
	return "", nil
}

func (r *linksRepository) GetShortLink(ctx context.Context, longLink string) (string, error) {
	return "", nil
}
