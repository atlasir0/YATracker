package main

import (
	"YAtracker/pkg/auth"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("token.env"); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	apiUrl := "https://login.yandex.ru/info"
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", "OAuth "+auth.GetToken())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("HTTP ответ:", resp.Status)
	fmt.Println("Тело ответа:", string(body))
}
