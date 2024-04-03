package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const REQ_TIMEOUT = 10 * time.Second
const DOMAIN = "https://thunderstore.io/"

var client = http.Client{Timeout: REQ_TIMEOUT}

func post(url string, contentType string, body any) ([]byte, error) {
	data, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, contentType, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	fmt.Println(response.StatusCode)
	return closeAndRead(response)
}

func get(url string) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return closeAndRead(response)
}

func closeAndRead(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	return io.ReadAll(res.Body)
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
