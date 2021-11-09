package fetcher

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/haski007/img-fetcher/internal/fetcher/model"
	"strings"
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

	var images []string
	document.Find("body img").Each(func(i int, selection *goquery.Selection) {
		link, exists := selection.Attr("src")
		if exists && strings.Contains(link, "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big") {
			images = append(images, link)
		}
	})
	xkom.Images = images

	if err := CreateFolder(xkom); err != nil {
		return fmt.Errorf("create folder err: %w", err)
	}

	return nil
}
