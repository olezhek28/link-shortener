package link_shortener_v1

import (
	linkShortener "github.com/olezhek28/link-shortener/internal/app/service/link_shortener"
	desc "github.com/olezhek28/link-shortener/pkg/link_shortener/v1"
)

type Implementation struct {
	desc.UnimplementedLinkShortenerV1Server

	linkShortenerService *linkShortener.Service
}

func NewLinkShortenerV1(linkShortenerService *linkShortener.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedLinkShortenerV1Server{},

		linkShortenerService,
	}
}

// func newMockLinkShortenerV1(i Implementation) *Implementation {
//	return &Implementation{
//		desc.UnimplementedLinkShortenerV1Server{},
//		i.linkShortenerService,
//	}
// }
