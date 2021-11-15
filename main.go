package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/haski007/img-fetcher/internal/fetcher"
	"github.com/haski007/img-fetcher/internal/fetcher/model"
	"github.com/sirupsen/logrus"
)

func main() {
	url := flag.String("url", "", "Link on page with your x-kom product!")
	langFlag := flag.String("lang", "en", "en/ru/ua/pl")
	flag.Parse()

	language := model.EncodeLanguage(*langFlag)
	if language == model.UnknownLanguage {
		logrus.Fatalf("Unknown language err: \n%s", flag.ErrHelp)
	}

	fetch := fetcher.New(&http.Client{}, language)

	if *url == "" {
		logrus.Fatalf("flag url should not be empty!\n%s", flag.ErrHelp)
	}

	if err := fetch.XKom(context.Background(), *url); err != nil {
		logrus.Fatalf("fetch x-kom err: %s", err)
	}
}
