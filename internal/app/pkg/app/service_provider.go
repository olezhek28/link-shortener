package app

import (
	"context"

	"github.com/olezhek28/link-shortener/internal/app/pkg/db"
	"github.com/olezhek28/link-shortener/internal/app/repository"
	"github.com/olezhek28/link-shortener/internal/app/service/linkShortener"
)

type serviceProvider struct {
	DB db.IClient

	linkShortenerService *linkShortener.Service

	linksRepository repository.ILinks
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GetLinkShortenerService ...
func (s *serviceProvider) GetLinkShortenerService(ctx context.Context) *linkShortener.Service {
	if s.linkShortenerService == nil {
		s.linkShortenerService = linkShortener.NewLinkShortenerService(s.GetLinksRepository(ctx))
	}

	return s.linkShortenerService
}

// GetLinksRepository ...
func (s *serviceProvider) GetLinksRepository(ctx context.Context) repository.ILinks {
	if s.linksRepository == nil {
		s.linksRepository = repository.NewLinksRepository(ctx, s.GetDB(ctx))
	}

	return s.linksRepository
}

// GetDB ...
func (s *serviceProvider) GetDB(ctx context.Context) db.IClient {
	if s.DB == nil {
		s.DB = db.NewClient(db.GetDbConfig())
	}

	return s.DB
}
