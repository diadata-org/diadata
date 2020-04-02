package utils

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// DownloadResource is a simple utility that downloads a resource
// from @url and stores it into @filepath.
func DownloadResource(filepath, url string) error {

	log.Printf("Downloading data")

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
