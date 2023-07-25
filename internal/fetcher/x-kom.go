package fetcher

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
	"github.com/haski007/img-fetcher/internal/fetcher/model"
)

func (rcv *Fetcher) XKom(ctx context.Context, pageUrl string) error {
	document, err := rcv.getHtmlDoc(pageUrl)
	if err != nil {
		return fmt.Errorf("[getHtml] err: %w", err)
	}

	var xkom model.XKom
	splitted := strings.Split(document.Find("title").Contents().Text(), "-")
	if len(splitted) > 0 {
		xkom.ProdTitle = strings.TrimSpace(splitted[0])
	}

	// ---> Fetch images
	var images []string
	document.Find("body img").Each(func(i int, selection *goquery.Selection) {
		link, exists := selection.Attr("src")
		if exists && strings.Contains(link, "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big") {
			images = append(images, link)
		}
	})
	xkom.Images = images

	// ---> fetch specifications
	var specs = make(map[string]string)
	document.Find(fmt.Sprintf("body .%s", model.XKomSpecsClass)).
		Each(func(i int, row *goquery.Selection) {
			key := row.Children().Find(fmt.Sprintf(".%s", model.XKomSpecsClass_Key)).Text()
			value := row.Children().Find(fmt.Sprintf(".%s", model.XKomSpecsClass_Value)).Text()

			if key == "" || value == "" {
				if value == "" && key == "" {
					return
				} else if value == "" && key != "" {
					logrus.Errorf("there are no value for key: %s", key)
				} else {
					logrus.Errorf("there are no key for value: %s", value)
				}
			}
			specs[key] = value
		})

	xkom.Specifications = specs

	if err := CreateFolder(ctx, xkom, rcv.Language); err != nil {
		return fmt.Errorf("create folder err: %w", err)
	}

	return nil
}
