package linkShortenerV1

import (
	"context"

	desc "github.com/olezhek28/link-shortener/pkg/link_shortener/v1"
)

func (i *Implementation) AddLink(ctx context.Context, req *desc.AddLinkRequest) (*desc.AddLinkResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	shortLink, err := i.linkShortenerService.AddLink(ctx, req.GetLongLink())
	if err != nil {
		return nil, err
	}

	return &desc.AddLinkResponse{
		Result: &desc.AddLinkResponse_Result{
			ShortLink: shortLink,
		},
	}, nil
}
