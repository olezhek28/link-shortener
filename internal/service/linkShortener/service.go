package linkShortener

import (
	"github.com/olezhek28/link-shortener/internal/pkg/redis"
	"github.com/olezhek28/link-shortener/internal/repository/links"
)

type Service struct {
	redisClient     redis.Client
	linksRepository links.Repository
}

func NewLinkShortenerService(
	redisClient redis.Client,
	linksRepository links.Repository) *Service {
	return &Service{
		redisClient:     redisClient,
		linksRepository: linksRepository,
	}
}

// NewMockInternalService ...
// func NewMockInternalService(deps ...interface{}) *Service {
//	lss := Service{}
//
//	for _, v := range deps {
//		switch s := v.(type) {
//		case repository.ILinks:
//			lss.linksRepository = s
//		}
//	}
//
//	return &lss
// }
