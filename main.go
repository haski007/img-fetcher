package main

import (
	"context"
	"flag"
	"github.com/haski007/img-fetcher/internal/fetcher"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	url := flag.String("url", "", "Link on page with your x-kom product!")
	flag.Parse()

	fetch := fetcher.New(&http.Client{})

	if *url == "" {
		logrus.Fatalf("flag url should not be empty!\n%s", flag.ErrHelp)
	}

	if err := fetch.XKom(context.Background(), *url); err != nil {
		logrus.Fatalf("fetch x-kom err: %s", err)
	}
}
