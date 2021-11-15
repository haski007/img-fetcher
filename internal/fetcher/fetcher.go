package fetcher

import (
	"net/http"

	"github.com/haski007/img-fetcher/internal/fetcher/model"
)

type Fetcher struct {
	httpClient *http.Client
	Language   model.Language
}

func New(httpClient *http.Client, lang model.Language) *Fetcher {
	return &Fetcher{
		httpClient: httpClient,
		Language:   lang,
	}
}
