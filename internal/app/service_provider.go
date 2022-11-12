package app

import (
	"context"
	"log"

	"github.com/olezhek28/link-shortener/internal/config"
	"github.com/olezhek28/link-shortener/internal/pkg/db"
	"github.com/olezhek28/link-shortener/internal/pkg/redis"
	"github.com/olezhek28/link-shortener/internal/repository/links"
	"github.com/olezhek28/link-shortener/internal/service/linkShortener"
)

type serviceProvider struct {
	db db.Client

	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	linkShortenerService *linkShortener.Service
	redisClient          redis.Client
	linksRepository      links.Repository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GetDB ...
func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		config, err := config.GetDbConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		db, err := db.NewClient(ctx, config)
		if err != nil {
			log.Fatalf("failed to connect to db: %s", err.Error())
		}

		s.db = db
	}

	return s.db
}

// GetGRPCConfig ...
func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.GetGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// GetHTTPConfig ...
func (s *serviceProvider) GetHTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.GetHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
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
		s.linksRepository = links.NewLinksRepository(ctx, s.GetDB(ctx))
	}

	return s.linksRepository
}


