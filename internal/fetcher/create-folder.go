package fetcher

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

func CreateFolder(r Resource) error {
	if err := os.Mkdir(r.GetTitle(), 0755); err != nil {
		return fmt.Errorf("os mkdir err: %w", err)
	}

	for i, image := range r.GetImages() {
		extention := strings.Split(image, ".")[len(strings.Split(image, "."))-1]
		if err := downloadFile(image, fmt.Sprintf("./%s/%d.%s", r.GetTitle(), i, extention)); err != nil {
			err = fmt.Errorf("download file err: %w", err)
			if errf := os.Remove(r.GetTitle()); errf != nil {
				err = fmt.Errorf("%w failed to remove created folder err: %w", err, errf)
			}
			return err
		}
	}
	logrus.Printf("Created new folder `%s` with %d images", r.GetTitle(), len(r.GetImages()))
	return nil
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return fmt.Errorf("create get request err: %w", err)
	}

	req.Header.Set("User-Agent", "Golang/pizda/v1.2.0")

	cli := &http.Client{}
	response, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("http response status: %s", response.Status)
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
