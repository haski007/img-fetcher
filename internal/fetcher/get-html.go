package fetcher

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (rcv *Fetcher) getHtmlDoc(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create get req err: %w", err)
	}

	req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("path", "/p/665466-karta-graficzna-nvidia-gigabyte-geforce-rtx-3060-ti-eagle-oc-lhr-8gb-gddr6.html")
	req.Header.Set("scheme", "https")
	req.Header.Set("accept-language", "ru,ru-RU;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("sec-ch-ua-platform", "macOS")
	req.Header.Set("User-Agent", "Golang/pizda/v1.2.0")

	rsp, err := rcv.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get page html err: %w", err)
	}
	defer func() {
		if err := rsp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("response status: %s", rsp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		logrus.Fatal(err)
	}

	return doc, nil
}
