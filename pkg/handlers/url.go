package handlers

import (
	"context"
	"errors"
	tiny "github.com/LittleMikle/gRPC_tinyurl"
	"github.com/LittleMikle/gRPC_tinyurl/proto"
)

func (i *Implementation) CreateURL(ctx context.Context, request *proto.FullURLRequest) (*proto.TinyURLResponse, error) {
	if request.FullURL == "" {
		return nil, errors.New("URL CANT BE EMPTY")
	}
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

func (i *Implementation) GetURL(ctx context.Context, request *proto.TinyURLRequest) (*proto.FullURLResponse, error) {
	if request.TinyURL == "" {
		return nil, errors.New("URL CANT BE EMPTY")
	}
	fullURL, err := i.urlService.GetURL(ctx, request.TinyURL)
	if err != nil {
		return nil, err
	}
	return &proto.FullURLResponse{
		FullURL: fullURL,
	}, nil
}
