package util

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const REQ_TIMEOUT = 10 * time.Second
const DOMAIN = "https://thunderstore.io/api/"

func get(url string) ([]byte, error) {
	client := http.Client{Timeout: REQ_TIMEOUT}

	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func asJSON[T interface{}](res []byte, err error) (T, error) {
	var data T

	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(res), &data)
	return data, nil
}

func JsonRequest[T interface{}](endpoint string) (T, error) {
	return asJSON[T](get(DOMAIN + endpoint))
}
