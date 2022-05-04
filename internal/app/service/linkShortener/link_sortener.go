package linkShortener

import (
	"github.com/olezhek28/link-shortener/internal/app/pkg/api/redis"
	"github.com/olezhek28/link-shortener/internal/app/repository"
)

type Service struct {
	redisClient redis.IRedisClient

	linksRepository repository.ILinks
}

func NewLinkShortenerService(
	redisClient redis.IRedisClient,

	linksRepository repository.ILinks,
) *Service {
	return &Service{
		redisClient: redisClient,

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
