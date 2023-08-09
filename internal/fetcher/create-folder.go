package fetcher

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/haski007/img-fetcher/internal/fetcher/model"

	"github.com/sirupsen/logrus"
)

func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func CreateFolder(ctx context.Context, r Resource, language model.Language) error {
	// ---> create directory
	dirName := strings.ReplaceAll(strings.ReplaceAll(r.GetTitle(), string(os.PathSeparator), "-"), " ", "_")
	dirName = strings.ReplaceAll(dirName, "/", "-")

	folderNameExists := folderExists(dirName)
	var i = 0
	for folderNameExists {
		i++
		folderNameExists = folderExists(fmt.Sprintf("%s_%d", dirName, i))
	}

	if i > 0 {
		dirName = fmt.Sprintf("%s_%d", dirName, i)
	}

	if err := os.Mkdir(dirName, 0755); err != nil {
		return fmt.Errorf("os mkdir err: %w", err)
	}

	// ---> download images
	for i, image := range r.GetImages() {
		extention := strings.Split(image, ".")[len(strings.Split(image, "."))-1]
		if err := downloadFile(image, fmt.Sprintf("./%s/%d.%s", dirName, i, extention)); err != nil {
			err = fmt.Errorf("download file err: %w", err)
			if errf := os.Remove(dirName); errf != nil {
				err = fmt.Errorf("%s failed to remove created folder err: %w", err, errf)
			}
			return err
		}
	}

	// ---> create specs file
	translatedText, err := r.GetSpecificationsTxt(ctx, language)
	if err != nil {
		return fmt.Errorf("[GetSpecificationsTxt] err: %w", err)
	}
	specsFileName := dirName + string(os.PathSeparator) + "specs.txt"
	if err := ioutil.WriteFile(specsFileName, []byte(translatedText), 0644); err != nil {
		logrus.Errorf("create and write file [%s] err: %s", specsFileName, err)
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

	//Write the bytes to the field
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
