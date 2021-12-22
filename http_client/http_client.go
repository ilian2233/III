package http_client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func VisitURL(fileURLAddress string) error {

	// Build fileName from fullPath
	fileURL, err := url.Parse(fileURLAddress)
	if err != nil {
		return err
	}
	segments := strings.Split(fileURL.Path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	// Gets file from internet
	resp, err := http.Get(fileURLAddress)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// Saves body of response for console print and file save
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	// Fills local file with response
	size, err := file.Write(body)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

	fmt.Printf("Content of web page: %s", string(body))

	return nil
}
