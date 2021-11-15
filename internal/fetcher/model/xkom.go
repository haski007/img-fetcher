package model

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/haski007/img-fetcher/api/translate"
)

const (
	XKomSpecsClass       = "juMOhe"
	XKomSpecsClass_Key   = "hnuwXU"
	XKomSpecsClass_Value = "iifYdH"
)

type XKom struct {
	ProdTitle      string            `json:"prod_title"`
	Images         []string          `json:"images"`
	Specifications map[string]string `json:"specifications"`
}

func (rcv XKom) GetTitle() string {
	return rcv.ProdTitle
}

func (rcv XKom) GetImages() []string {
	return rcv.Images
}

func (rcv XKom) GetSpecifications() map[string]string {
	return rcv.Specifications
}

func (rcv XKom) GetSpecificationsTxt(ctx context.Context, language Language) (string, error) {
	var text string
	for key, value := range rcv.GetSpecifications() {
		if strings.Count(value, "\n") > 1 {
			text += fmt.Sprintf("%s:\n%s", key, value)
		} else {
			text += fmt.Sprintf("%s%s", key, value)
		}
	}

	if language == Polish {
		return text, nil
	}

	translation, err := translate.Concurrently(
		http.DefaultClient,
		text,
		Polish.String(),
		language.String(),
	)
	if err != nil {
		return "", fmt.Errorf("translate err: %w", err)
	}

	return translation, nil
}
