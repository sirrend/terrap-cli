package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// PerformRequestWithBody performs a request with a specified method (POST or GET) and a body.
func PerformRequest(method, url string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

// PerformGetRequestWithParams performs a GET request with query parameters.
func PerformRequestWithParams(url string, params map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	return client.Do(req)
}
