package link_shortener

import (
	"context"
)

func (s *Service) GetShortLink(ctx context.Context, longLink string) (string, error) {
	return s.linksRepository.GetShortLink(ctx, longLink)
}
