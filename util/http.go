package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const REQ_TIMEOUT = 10 * time.Second
const DOMAIN = "https://thunderstore.io/"

var client = resty.NewWithClient(&http.Client{Timeout: REQ_TIMEOUT})

func post(url string, contentType string, body any) ([]byte, error) {
	data, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	// req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	// req.Header.Set("Content-Type", contentType)
	// req.AddCookie(http.Cookie{})

	response, err := client.R().
		SetBody(bytes.NewReader(data)).
		SetHeader("Content-Type", contentType).
		Post(url)

	if err != nil {
		return nil, err
	}

	fmt.Println(response.StatusCode())
	return response.Body(), nil
}

func get(url string) ([]byte, error) {
	response, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	return response.Body(), nil
}

func asJSON[T interface{}](res []byte, err error) (T, error) {
	var data T

	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(res), &data)
	return data, nil
}

func JsonGetRequest[T interface{}](endpoint string) (T, error) {
	return asJSON[T](get(DOMAIN + endpoint))
}

func JsonPostRequest[T interface{}](endpoint string, body any) (T, error) {
	return asJSON[T](post(DOMAIN+endpoint, "application/json", body))
}
