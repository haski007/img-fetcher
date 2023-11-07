package fetcher

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func (rcv *Fetcher) getHtmlDoc(url string) (*goquery.Document, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var htmlContent string
	tasks := chromedp.Tasks{
		network.Enable(),       // Вмикаємо мережеві події
		chromedp.Navigate(url), // Переходимо на google.com
		chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(err)
	}

	res, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("new document from reader err: %w", err)
	}

	os.WriteFile("test.html", []byte(htmlContent), 0755)
	return res, nil
}
