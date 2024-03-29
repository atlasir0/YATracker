package main

import (
	"fmt"
	"net/http"
	"os"
)

type Request struct {
}

type Response struct {
	Message string `json:"message"`
}

func Handler() (*Response, error) {
	apiUrl := "https://login.yandex.ru/info"
	httpReq, err := http.NewRequest("GET", apiUrl, nil)

	if err != nil {
		return nil, err
	}

	token := os.Getenv("YANDEX_TOKEN")
	httpReq.Header.Set("Authorization", "OAuth "+token)
	resp, err := http.DefaultClient.Do(httpReq)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("HTTP ответ:", resp.Status)

	if resp.StatusCode == http.StatusOK {
		return &Response{Message: "Доступ получен"}, nil
	} else {
		return &Response{Message: resp.Status}, nil
	}
}

func main() {
	response, err := Handler()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Сообщение:", response.Message)
}
