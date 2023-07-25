package model

import (
	"context"
	"fmt"
	"net/http"

	"github.com/haski007/img-fetcher/api/translate"
	"github.com/sirupsen/logrus"
)

const (
	XKomSpecsClass       = "iFAJlN"
	XKomSpecsClass_Key   = "sc-13p5mv-1"
	XKomSpecsClass_Value = "UfEQd"
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
		//if strings.Count(value, "\n") > 1 {
		text += fmt.Sprintf("%s: %s\n", key, value)
		//} else {
		//	text += fmt.Sprintf("%s%s", key, value)
		//}
	}

	if language == Polish {
		return text, nil
	}

	logrus.Printf("Translating specifications on [%s] language", language.GetLocale())
	translation, err := translate.Concurrently(
		http.DefaultClient,
		text,
		Polish.GetLocale().String(),
		language.GetLocale().String(),
	)
	if err != nil {
		return "", fmt.Errorf("translate err: %w", err)
	}

	return translation, nil
}
