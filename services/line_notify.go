package services

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func LineNotify(token, message string) {

	url := "https://notify-api.line.me/api/notify"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("message=%s", message))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if _, err := io.ReadAll(res.Body); err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
}
