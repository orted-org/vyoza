package util

import (
	"errors"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	ErrHTTPRequestTimeoutExceeded = errors.New("http request timeout exceeded")
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
	res, err := client.Do(req)
	if err != nil {
		if isTimeoutError(err) {
			return nil, ErrHTTPRequestTimeoutExceeded
		}
		return nil, err
	}
	return res, nil
}

// if the error is due to timeout
func isTimeoutError(err error) bool {
	e, ok := err.(net.Error)
	return ok && e.Timeout()
}
