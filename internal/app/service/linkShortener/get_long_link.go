package linkShortener

import (
	"context"
)

func (s *Service) GetLongLink(ctx context.Context, shortLink string) (string, error) {
	longLink, err := s.redisClient.Get(ctx, shortLink)
	if nil == err {
		return longLink, err
	}

	return s.linksRepository.GetLongLink(ctx, shortLink)
}
