package linkShortener

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"
)

const (
	prefix     = "https://olezhek28.com/"
	expiration = time.Hour
)

func (s *Service) AddLink(ctx context.Context, longLink string) (string, error) {
	hash := md5.Sum([]byte(longLink))
	var builder strings.Builder
	builder.WriteString(prefix)
	builder.WriteString(hex.EncodeToString(hash[:]))
	shortLink := builder.String()

	err := s.redisClient.Set(ctx, shortLink, longLink, expiration)
	if err != nil {
		return "", err
	}

	err = s.linksRepository.AddLink(ctx, longLink, shortLink)
	if err != nil {
		return "", err
	}

	return shortLink, nil
}
