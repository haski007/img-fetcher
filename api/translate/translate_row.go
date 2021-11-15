package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	apiUrl = "https://translate.googleapis.com/translate_a/single"
)

func Row(client *http.Client, row, sourceLang, targetLang string) (translated string, err error) {
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return "", fmt.Errorf("create new request err: %w", err)
	}

	q := req.URL.Query()
	q.Add("client", "gtx")
	q.Add("sl", sourceLang)
	q.Add("tl", targetLang)
	q.Add("dt", "t")
	q.Add("q", row)
	req.URL.RawQuery = q.Encode()

	rsp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http client do err: %s", err)
	}

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("read response body err: %s", err)
	}

	if rsp.StatusCode != 200 {
		logrus.Errorf(string(data))
		return "", fmt.Errorf("status code is not 200 status: %s", rsp.Status)
	}
	defer rsp.Body.Close()

	var face []interface{}

	if err := json.Unmarshal(data, &face); err != nil {
		return "", fmt.Errorf("unmarshall response data err: %s", err)
	}

	if face[0] != nil {
		translations := face[0].([]interface{})

		for _, t := range translations {
			if t == nil {
				continue
			}
			translated += t.([]interface{})[0].(string)
		}
	}
	return
}
