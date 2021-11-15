package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestTranslation(t *testing.T) {
	var url = "https://translate.googleapis.com/translate_a/single"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("create new request err: %s", err)
	}

	q := req.URL.Query()
	q.Add("client", "gtx")
	q.Add("sl", "pl")
	q.Add("tl", "en")
	q.Add("dt", "t")
	q.Add("q", "Powyższa odsłona gogli od HTC to synonim słowa premium. Gogle wykonane z najwyższej jakości materiał&oacute;w. Lata doświadczeń oraz bycie pionierem na rynku w zakresie VR dają gwarancję jakości i niezawodności. Gogle charakteryzują się najlepszą jakością obrazu na rynku, trwałością oraz szerokim wsparciem producenta w ramach platformy VIVE Port. Dodatkowo, gogle wsp&oacute;łpracują z platfo")
	req.URL.RawQuery = q.Encode()

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("http client do err: %s", err)
	}
	if rsp.StatusCode != 200 {
		t.Fatalf("status code is not 200 status: %s", rsp.Status)
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("read response body err: %s", err)
	}

	var face []interface{}

	if err := json.Unmarshal(data, &face); err != nil {
		t.Fatalf("unmarshall response data err: %s", err)
	}

	translations := face[0].([]interface{})

	var translatedText string
	for _, t := range translations {
		translatedText += t.([]interface{})[0].(string)
	}

	fmt.Println(translatedText)
}
