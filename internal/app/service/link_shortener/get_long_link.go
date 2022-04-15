package link_shortener

import (
	"context"
)

func (s *Service) GetLongLink(ctx context.Context, shortLink string) (string, error) {
	return s.linksRepository.GetLongLink(ctx, shortLink)
}
