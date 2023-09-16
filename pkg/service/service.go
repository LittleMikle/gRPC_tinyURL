package service

import (
	"context"
	tiny "github.com/LittleMikle/gRPC_tinyurl"
	"github.com/LittleMikle/gRPC_tinyurl/pkg/repository"
)

type URL interface {
	CreateURL(ctx context.Context, url tiny.URL) (string, error)
	GetURL(ctx context.Context, tinyURL string) (string, error)
}

type Service struct {
	URL
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		URL: NewURLService(repos.URL),
	}
}
