package indexOne

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func cancel() {
	log.Fatal("User Quit Installer")

}

func downloadFile(filePath, url string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making HTTP GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying response body to file: %w", err)
	}

	return nil
}
