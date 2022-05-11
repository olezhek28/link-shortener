package linkShortener

import (
	"context"
)

func (s *Service) GetLongLink(ctx context.Context, shortLink string) (string, error) {
	longLink, err := s.redisClient.Get(ctx, shortLink)
	if nil == err {
		return longLink, nil
	}

	longLink, err = s.linksRepository.GetLongLink(ctx, shortLink)
	if err != nil {
		return "", err
	}

	err = s.redisClient.Set(ctx, shortLink, longLink, expiration)
	if err != nil {
		return "", err
	}

	return longLink, nil
}
