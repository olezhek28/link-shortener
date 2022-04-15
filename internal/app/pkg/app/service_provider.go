package app

import (
	"context"

	"github.com/olezhek28/link-shortener/internal/app/repository"
	linkShortener "github.com/olezhek28/link-shortener/internal/app/service/link_shortener"
)

type serviceProvider struct {
	linkShortenerService *linkShortener.Service

	linksRepository repository.ILinks
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetLinkShortenerService(ctx context.Context) *linkShortener.Service {
	if s.linkShortenerService == nil {
		s.linkShortenerService = linkShortener.NewLinkShortenerService(s.GetLinksRepository(ctx))
	}

	return s.linkShortenerService
}

func (s *serviceProvider) GetLinksRepository(ctx context.Context) repository.ILinks {
	if s.linksRepository == nil {
		s.linksRepository = repository.NewLinksRepository(ctx)
	}

	return s.linksRepository
}
