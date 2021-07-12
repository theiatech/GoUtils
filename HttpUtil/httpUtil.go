package HttpUtil

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpResponse struct {
	Header     map[string][]string
	StatusCode int
	Body       []byte
	Error      error
}

func Get(u string, headers map[string]string) HttpResponse {
	client := &http.Client{Timeout: 60 * time.Second}
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return HttpResponse{nil, 0, nil, err}
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return HttpResponse{nil, 0, nil, err}
	}
	if http.StatusOK != resp.StatusCode {
		return HttpResponse{resp.Header, resp.StatusCode, nil, err}
	}
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return HttpResponse{resp.Header, resp.StatusCode, nil, err}
		}
	} else {
		reader = resp.Body
	}
	body, err := ioutil.ReadAll(reader)
	resp.Body.Close()
	reader.Close()
	return HttpResponse{resp.Header, resp.StatusCode, body, err}
}

func Post(u string, postBody string, headers map[string]string) HttpResponse {
	client := &http.Client{Timeout: 60 * time.Second}
	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(postBody))
	if err != nil {
		return HttpResponse{nil, 0, nil, err}
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return HttpResponse{nil, 0, nil, err}
	}
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return HttpResponse{resp.Header, resp.StatusCode, nil, err}
		}
	} else {
		reader = resp.Body
	}
	body, err := ioutil.ReadAll(reader)
	resp.Body.Close()
	reader.Close()
	return HttpResponse{resp.Header, resp.StatusCode, body, err}
}
