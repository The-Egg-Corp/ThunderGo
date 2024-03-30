package util

import (
	"io"
	"net/http"
	"time"
)

const REQ_TIMEOUT = 10 * time.Second

func Request(url string) ([]byte, error) {
	client := http.Client{Timeout: REQ_TIMEOUT}

	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return io.ReadAll(response.Body)
}
