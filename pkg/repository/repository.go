package repository

import (
	"context"
	tiny "github.com/LittleMikle/gRPC_tinyurl"
	"github.com/jmoiron/sqlx"
)

type URL interface {
	CreateURL(ctx context.Context, url tiny.URL) (string, error)
	GetURL(ctx context.Context, tinyURL string) (string, error)
}

type Repository struct {
	URL
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		URL: NewURLPostgres(db),
	}
}
