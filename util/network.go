package util

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	ErrHTTPRequestTimeoutExceeded = errors.New("http request timeout exceeded")
	ErrTLSRequestTimeoutExceeded  = errors.New("tls request timed out")
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

// does the same thing as Fetch but instead returns response time and status code
// Status -> -1 for error, anything else if request successful
// ResponseTime -> -1 for error and timeout, anything else if request successful
func StatusAndResponseTimeFor(method string, url string, header map[string]string, body io.Reader, timeout int) (int, int) {
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

	start := time.Now()
	// performing the call and returning the response and error
	res, err := client.Do(req)
	if err != nil {
		return -1, -1
	}
	return res.StatusCode, int(time.Since(start).Milliseconds())
}

type SSLCertificateDetails struct {
	IsValid bool
	Expiry  time.Time
	Remark  string
}

func GetSSLCertificateDetails(url string, timeout int) SSLCertificateDetails {
	url = convertURLToTLSURI(url)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))
	defer cancelFunc()
	result := make(chan SSLCertificateDetails)
	go func(result chan<- SSLCertificateDetails, url *string) {
		genRes := SSLCertificateDetails{
			IsValid: false,
		}
		conn, err := tls.Dial("tcp", *url, nil)
		if err != nil {
			genRes.Remark = err.Error()
		}
		defer conn.Close()
		genRes.IsValid = true
		genRes.Remark = conn.ConnectionState().PeerCertificates[0].Issuer.String()
		genRes.Expiry = conn.ConnectionState().PeerCertificates[0].NotAfter.UTC()
		result <- genRes
	}(result, &url)
	select {
	case <-ctx.Done():
		return SSLCertificateDetails{IsValid: false, Remark: ErrTLSRequestTimeoutExceeded.Error()}
	case <-result:
		return <-result
	}
}

func convertURLToTLSURI(url string) string {
	parts := strings.Split(url, "//")
	if len(parts) > 1 {
		url = parts[1]
	}
	return strings.Split(url, ":")[0] + ":443"
}
