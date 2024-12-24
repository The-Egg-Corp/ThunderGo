package util

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const REQ_TIMEOUT = 10 * time.Second
const DOMAIN = "https://thunderstore.io/"

//const AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 OPR/107.0.0.0 (Edition std-1)"

var client = resty.NewWithClient(&http.Client{Timeout: REQ_TIMEOUT})

func Get(url string, contentType string) (*[]byte, *int, error) {
	response, err := client.R().
		SetHeader("Content-Type", contentType).
		Get(url)

	if err != nil {
		return nil, nil, err
	}

	resBody := response.Body()
	statusCode := response.StatusCode()

	return &resBody, &statusCode, nil
}

func Post(url string, contentType string, body any) (*[]byte, *int, error) {
	data, err := json.Marshal(&body)
	if err != nil {
		return nil, nil, err
	}

	// req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	// req.Header.Set("Content-Type", contentType)
	// req.AddCookie(http.Cookie{})

	response, err := client.R().
		SetBody(bytes.NewReader(data)).
		SetHeader("Content-Type", contentType).
		Post(url)

	if err != nil {
		return nil, nil, err
	}

	resBody := response.Body()
	statusCode := response.StatusCode()

	return &resBody, &statusCode, nil
}

func JsonGetRequest[Expected interface{}](endpoint string) (*Expected, *int, error) {
	resBody, code, err := Get(DOMAIN+endpoint, "application/json")
	return fromJSON[Expected](resBody), code, err
}

func JsonPostRequest[Expected interface{}, Body interface{}](endpoint string, body Body) (*Expected, *int, error) {
	resBody, code, err := Post(DOMAIN+endpoint, "application/json", body)
	return fromJSON[Expected](resBody), code, err
}

func fromJSON[T interface{}](data *[]byte) *T {
	if data == nil {
		return nil
	}

	var val T

	err := json.Unmarshal([]byte(*data), &val)
	if err != nil {
		return nil
	}

	return &val
}
