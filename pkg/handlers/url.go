package handlers

import (
	"context"
	tiny "github.com/LittleMikle/gRPC_tinyurl"
	"github.com/LittleMikle/gRPC_tinyurl/proto"
)

func (i *Implementation) CreateURL(ctx context.Context, request *proto.FullURLRequest) (*proto.TinyURLResponse, error) {
	var url tiny.URL
	url.FullURL = request.FullURL

	token, err := i.urlService.CreateURL(ctx, url)
	if err != nil {
		return nil, err
	}

	return &proto.TinyURLResponse{
		TinyURL: token,
	}, nil
}
