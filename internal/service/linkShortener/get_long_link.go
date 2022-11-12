package linkShortener

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/olezhek28/link-shortener/internal/metrics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	prefix     = "https://olezhek28.com/"
	expiration = time.Hour
)

func (s *Service) GetLongLink(ctx context.Context, shortLink string) (string, error) {
	longLink, err := s.redisClient.Get(ctx, shortLink)
	if nil == err {
		metrics.IncCacheMiss("GetLongLink")
		return longLink, nil
	}

	metrics.IncCacheHit("GetLongLink")

	longLink, err = s.linksRepository.GetLongLink(ctx, shortLink)
	if err == pgx.ErrNoRows {
		return "", status.Error(codes.NotFound, "long link not found")
	}
	if err != nil {
		return "", err
	}

	err = s.redisClient.Set(ctx, shortLink, longLink, expiration)
	if err != nil {
		return "", err
	}

	return longLink, nil
}
