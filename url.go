package tiny

type URL struct {
	Id      int    `json:"id" db:"id"`
	FullURL string `json:"fullURL" db:"fullURL"`
	TinyURL string `json:"tinyURL" db:"tinyURL"`
}
