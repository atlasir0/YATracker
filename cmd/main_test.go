package main_test

import (
	"YAtracker/pkg/auth"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	os.Setenv("YANDEX_TOKEN", "test_yandex_token")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "OAuth test_yandex_token"
		assert.Equal(t, expectedAuthHeader, authHeader, "Неправильный заголовок Authorization")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success response"))
	}))
	defer server.Close()

	err := godotenv.Load("token.env")
	assert.NoError(t, err, "Не удалось загрузить файл .env")

	apiUrl := server.URL
	req, err := http.NewRequest("GET", apiUrl, nil)
	assert.NoError(t, err, "Ошибка создания HTTP-запроса")
	req.Header.Set("Authorization", "OAuth "+auth.GetToken())
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Ошибка отправки HTTP-запроса")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err, "Ошибка чтения ответа")

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Неправильный код состояния ответа")
	assert.Equal(t, "success response", string(body), "Неправильное тело ответа")
}
