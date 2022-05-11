package linkShortener

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
)

const (
	prefix     = "https://olezhek28.com/"
	expiration = time.Hour
)

func (s *Service) AddLink(ctx context.Context, longLink string) (string, error) {
	hash := sha256.Sum256([]byte(longLink))
	var builder strings.Builder
	builder.WriteString(prefix)
	builder.WriteString(hex.EncodeToString(hash[:]))
	shortLink := builder.String()

	// check exist
	_, err := s.linksRepository.GetLongLink(ctx, shortLink)
	if err == nil {
		return shortLink, nil
	}

	err = s.linksRepository.AddLink(ctx, longLink, shortLink)
	if err != nil {
		return "", err
	}

	err = s.redisClient.Set(ctx, shortLink, longLink, expiration)
	if err != nil {
		return "", err
	}

	return shortLink, nil
}
