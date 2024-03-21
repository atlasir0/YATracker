package auth

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

func GetToken(key string) string {
	return store.tokens[key]
}
