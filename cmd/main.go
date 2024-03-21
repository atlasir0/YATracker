package main

import (
	"YAtracker/pkg/auth"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	auth.SetToken("yandex", "y0_AgAAAAAXBtATAAt7GwAAAAD-47QTAAAPIC1E5hNDO7tHK3Vgxuh3p7uWGw")

	apiUrl := "https://login.yandex.ru/info"
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", "OAuth "+auth.GetToken("yandex"))
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

	fmt.Println(string(body))
}
