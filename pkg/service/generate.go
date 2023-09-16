package service

import gonanoid "github.com/matoous/go-nanoid/v2"

func Generate() (string, error) {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	token, err := gonanoid.Generate(alphabet, 10)
	return token, err
}
