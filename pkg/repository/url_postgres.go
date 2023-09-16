package repository

import (
	"context"
	"fmt"
	tiny "github.com/LittleMikle/gRPC_tinyurl"
	"github.com/jmoiron/sqlx"
)

type URLPostgres struct {
	db *sqlx.DB
}

func NewURLPostgres(db *sqlx.DB) *URLPostgres {
	return &URLPostgres{
		db: db,
	}
}

func (r *URLPostgres) CreateURL(ctx context.Context, url tiny.URL) (string, error) {
	query := fmt.Sprintf("INSERT INTO urls (fullURL, tinyURL) values ($1, $2) RETURNING tinyURL")
	row := r.db.QueryRow(query, url.FullURL, url.TinyURL)
	err := row.Scan(&url.TinyURL)
	if err != nil {
		return "", err
	}
	return url.TinyURL, nil
}

func (r *URLPostgres) GetURL(ctx context.Context, tinyURL string) (string, error) {
	return "", nil
}
