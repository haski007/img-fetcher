package fetcher

import (
	"net/http"
)

type Fetcher struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) *Fetcher {
	return &Fetcher{
		httpClient: httpClient,
	}
}
