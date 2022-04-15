package link_shortener

import "github.com/olezhek28/link-shortener/internal/app/repository"

type Service struct {
	linksRepository repository.ILinks
}

func NewLinkShortenerService(linksRepository repository.ILinks) *Service {
	return &Service{
		linksRepository: linksRepository,
	}
}

// NewMockInternalService ...
func NewMockInternalService(deps ...interface{}) *Service {
	lss := Service{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.ILinks:
			lss.linksRepository = s
		}
	}

	return &lss
}
