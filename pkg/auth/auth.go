package auth

import (
	"os"
)

type tokenStore struct {
	tokens map[string]string
}

var store tokenStore

func init() {
	store = tokenStore{
		tokens: make(map[string]string),
	}
}

func SetToken(key, value string) {
	store.tokens[key] = value
}

func GetToken() string {
	token := os.Getenv("YANDEX_TOKEN")
	return token
}
