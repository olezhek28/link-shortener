package app

import (
	"context"
	"log"

	"github.com/olezhek28/link-shortener/internal/pkg/db"
	"github.com/olezhek28/link-shortener/internal/pkg/redis"
	"github.com/olezhek28/link-shortener/internal/repository/links"
	"github.com/olezhek28/link-shortener/internal/service/linkShortener"
)

type serviceProvider struct {
	DB db.Client

	linkShortenerService *linkShortener.Service
	redisClient          redis.Client
	linksRepository      links.Repository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GetLinkShortenerService ...
func (s *serviceProvider) GetLinkShortenerService(ctx context.Context) *linkShortener.Service {
	if s.linkShortenerService == nil {
		s.linkShortenerService = linkShortener.NewLinkShortenerService(s.GetRedisClient(), s.GetLinksRepository(ctx))
	}

	return s.linkShortenerService
}

func (s *serviceProvider) GetRedisClient() redis.Client {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient()
	}

	return s.redisClient
}

// GetLinksRepository ...
func (s *serviceProvider) GetLinksRepository(ctx context.Context) links.Repository {
	if s.linksRepository == nil {
		s.linksRepository = links.NewLinksRepository(ctx, s.GetDB())
	}

	return s.linksRepository
}

// GetDB ...
func (s *serviceProvider) GetDB() db.Client {
	if s.DB == nil {
		config, err := db.GetDbConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		s.DB = db.NewClient(config)
	}

	return s.DB
}
