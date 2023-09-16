package service

import (
	"context"
	tiny "github.com/LittleMikle/gRPC_tinyurl"
	"github.com/LittleMikle/gRPC_tinyurl/pkg/repository"
)

type URLService struct {
	repo repository.URL
}

func NewURLService(repo repository.URL) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (s *URLService) CreateURL(ctx context.Context, url tiny.URL) (string, error) {
	token, err := Generate()
	if err != nil {
		return "", err
	}

	url.TinyURL = token
	_, err = s.repo.CreateURL(ctx, url)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *URLService) GetURL(ctx context.Context, tinyURL string) (string, error) {
	return s.repo.GetURL(ctx, tinyURL)
}
