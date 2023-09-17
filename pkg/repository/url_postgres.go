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
	query := fmt.Sprintf("INSERT INTO urls (fullurl, tinyurl) values ($1, $2) RETURNING tinyurl")
	row := r.db.QueryRow(query, url.FullURL, url.TinyURL)
	err := row.Scan(&url.TinyURL)
	if err != nil {
		return "", err
	}
	return url.TinyURL, nil
}

func (r *URLPostgres) GetURL(ctx context.Context, tinyURL string) (string, error) {
	fmt.Println(tinyURL)
	var fullURL string
	query := fmt.Sprintf("SELECT fullurl FROM urls WHERE tinyurl=$1")
	if err := r.db.Get(&fullURL, query, tinyURL); err != nil {
		return "", err
	}
	return fullURL, nil
}
