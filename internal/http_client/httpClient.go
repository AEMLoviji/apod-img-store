package http_client

import (
	"net/http"
)

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type httpClient struct{}

func NewHttpClient() HttpClient {
	return httpClient{}
}

func (c httpClient) Get(url string) (*http.Response, error) {
	return getInternal(url)
}

func getInternal(url string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", url, nil)

	client := http.Client{}

	return client.Do(request)
}
