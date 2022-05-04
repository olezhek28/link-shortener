package linkShortener

import (
	"context"
)

func (s *Service) GetLongLink(ctx context.Context, shortLink string) (string, error) {
	return s.linksRepository.GetLongLink(ctx, shortLink)
}
