package translate

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/haski007/img-fetcher/pkg/throttler"
	"golang.org/x/sync/errgroup"
)

const (
	rateLimit = 5
)

func Concurrently(
	client *http.Client,
	text, sourceLang, targetLang string,
) (translated string, err error) {
	var (
		errG errgroup.Group

		t = throttler.New(time.Second, rateLimit)
	)

	arr := strings.Split(text, "\n")

	for i, row := range arr {
		errG.Go(worker(row, i, arr, t, client, sourceLang, targetLang))
	}

	if err := errG.Wait(); err != nil {
		return "", fmt.Errorf("error group wait err: %w", err)
	}

	return strings.Join(arr, "\n"), nil
}

func worker(row string, index int, arr []string, t *throttler.Throttler, client *http.Client, sourceLang, targetLang string) func() error {
	return func() error {
		t.Throttle()

		var err error
		t.Lock()
		arr[index], err = Row(client, row, sourceLang, targetLang)
		t.Unlock()
		if err != nil {
			return fmt.Errorf("translate row err: %w", err)
		}
		return nil
	}
}
