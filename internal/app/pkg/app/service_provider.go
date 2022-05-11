package app

import (
	"context"
	"log"

	"github.com/olezhek28/link-shortener/internal/app/pkg/api/redis"
	"github.com/olezhek28/link-shortener/internal/app/pkg/db"
	"github.com/olezhek28/link-shortener/internal/app/repository"
	"github.com/olezhek28/link-shortener/internal/app/service/linkShortener"
)

type serviceProvider struct {
	DB db.IClient

	linkShortenerService *linkShortener.Service

	redisClient redis.IRedisClient

	linksRepository repository.ILinks
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GetLinkShortenerService ...
func (s *serviceProvider) GetLinkShortenerService(ctx context.Context) *linkShortener.Service {
	if s.linkShortenerService == nil {
		s.linkShortenerService = linkShortener.NewLinkShortenerService(
			s.GetRedisClient(),

			s.GetLinksRepository(ctx),
		)
	}

	return s.linkShortenerService
}

func (s *serviceProvider) GetRedisClient() redis.IRedisClient {
	if s.redisClient == nil {
		s.redisClient = redis.NewRedisClient()
	}

	return s.redisClient
}

// GetLinksRepository ...
func (s *serviceProvider) GetLinksRepository(ctx context.Context) repository.ILinks {
	if s.linksRepository == nil {
		s.linksRepository = repository.NewLinksRepository(ctx, s.GetDB())
	}

	return s.linksRepository
}

// GetDB ...
func (s *serviceProvider) GetDB() db.IClient {
	if s.DB == nil {
		config, err := db.GetDbConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		s.DB = db.NewClient(config)
	}

	return s.DB
}
