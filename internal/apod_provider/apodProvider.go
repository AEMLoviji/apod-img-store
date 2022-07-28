package apodprovider

import (
	"apod-img-store/internal/http_client"
	"encoding/json"
	"errors"
	"io"
)

type ApodProvider interface {
	GetImage() (*ApodResponse, error)
}

type apodProvider struct {
	HttpClient http_client.HttpClient
}

func NewApodProvider(httpClient http_client.HttpClient) ApodProvider {
	return &apodProvider{HttpClient: httpClient}
}

const (
	urlGetImage = "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY"
)

func (p *apodProvider) GetImage() (*ApodResponse, error) {
	response, err := p.HttpClient.Get(urlGetImage)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("invalid response body")
	}

	defer response.Body.Close()

	var result ApodResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, errors.New("error when trying to unmarshal get apod image response")
	}

	return &result, nil
}
