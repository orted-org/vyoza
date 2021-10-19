package util

import (
	"io"
	"net/http"
	"time"
)

func Fetch(method string, url string, header map[string]string, body io.Reader, timeout int) (*http.Response, error) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},

		// since all the operations are in milliseconds, taking time.Millisecond
		Timeout: time.Duration(time.Duration(timeout) * time.Millisecond),
	}

	// creating the request
	req, _ := http.NewRequest(method, url, body)

	// adding the headers
	for key, value := range header {
		req.Header.Add(key, value)
	}

	// performing the call and returning the response and error
	return client.Do(req)
}
