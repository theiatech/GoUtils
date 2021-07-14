package HttpUtil

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpRequestParams struct {
	Method string
	Url string `json:"url"`
	Body []byte
	Headers     map[string]string
	BasicAuth  struct {
		Username string
		Password string
	}
	Timeout time.Duration
}

type HttpResponse struct {
	Header     map[string][]string
	StatusCode int
	Body       []byte
	Error      error
}

func Requests(httpRequestParams HttpRequestParams) HttpResponse {
	if 0 == httpRequestParams.Timeout {
		httpRequestParams.Timeout = 10 * time.Second
	}
	client := &http.Client{Timeout: httpRequestParams.Timeout}
	if "" == httpRequestParams.Method {
		httpRequestParams.Method = http.MethodGet
	}
	req, err := http.NewRequest(httpRequestParams.Method, httpRequestParams.Url, bytes.NewReader(httpRequestParams.Body))
	if err != nil {
		return HttpResponse{nil, 0, nil, err}
	}
	if 0 < len(httpRequestParams.BasicAuth.Username) {
		req.SetBasicAuth(httpRequestParams.BasicAuth.Username,httpRequestParams.BasicAuth.Password)
	}
	req.Header.Set("User-Agent", "curl/7.64.1")
	for k, v := range httpRequestParams.Headers {
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
