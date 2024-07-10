package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pborman/getopt"
)

func retrieveURLAndFilename() (string, string) {
	urlPath := getopt.StringLong("url", 'u', "", "url")
	getopt.Parse()
	_, err := url.Parse(*urlPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*urlPath)
	splittedURL := strings.Split(*urlPath, "/")
	return *urlPath, splittedURL[len(splittedURL)-1]
}

func createFileWithGivenName(filename string) *os.File {
	fmt.Printf("Creating file with name: %s\n", filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func downloadDataFromURL(urlPath string, client *http.Client, file *os.File) int64 {
	resp, err := client.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return size
}

func main() {
	urlPath, filename := retrieveURLAndFilename()
	file := createFileWithGivenName(filename)
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	size := downloadDataFromURL(urlPath, &client, file)

	fmt.Printf("Successfully downloaded file from %s with size %d bytes\n", urlPath, size)
}
