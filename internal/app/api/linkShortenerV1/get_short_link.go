package linkShortenerV1

import (
	"context"

	desc "github.com/olezhek28/link-shortener/pkg/link_shortener/v1"
)

func (i *Implementation) GetShortLink(ctx context.Context, req *desc.GetShortLinkRequest) (*desc.GetShortLinkResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	shortLink, err := i.linkShortenerService.GetShortLink(ctx, req.GetLongLink())
	if err != nil {
		return nil, err
	}

	return &desc.GetShortLinkResponse{
		Result: &desc.GetShortLinkResponse_Result{
			ShortLink: shortLink,
		},
	}, nil
}
