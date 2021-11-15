package fetcher

import (
	"context"

	"github.com/haski007/img-fetcher/internal/fetcher/model"
)

type Resource interface {
	GetTitle() string
	GetImages() []string
	GetSpecifications() map[string]string
	GetSpecificationsTxt(ctx context.Context, language model.Language) (string, error)
}
