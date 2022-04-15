package link_shortener_v1

import (
	"context"

	desc "github.com/olezhek28/link-shortener/pkg/link_shortener/v1"
)

func (i *Implementation) GetLongLink(ctx context.Context, req *desc.GetLongLinkRequest) (*desc.GetLongLinkResponse, error) {
	longLink, err := i.linkShortenerService.GetLongLink(ctx, req.GetShortLink())
	if err != nil {
		return nil, err
	}

	return &desc.GetLongLinkResponse{
		Result: &desc.GetLongLinkResponse_Result{
			LongLink: longLink,
		},
	}, nil
}
