package fetcher

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetXKomImages(t *testing.T) {
	fetcher := New(&http.Client{})

	const (
		url = "https://www.x-kom.pl/p/665466-karta-graficzna-nvidia-gigabyte-geforce-rtx-3060-ti-eagle-oc-lhr-8gb-gddr6.html"
	)
	data, err := fetcher.getHtml(url)
	if err != nil {
		t.Fatalf("get html err: %s", err)
	}

	fmt.Println(string(data))
	if err := GetXKomImages(data); err != nil {
		t.Fatalf("get xkom images err: %s", err)
	}
}
