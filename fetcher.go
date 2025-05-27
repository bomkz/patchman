package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func checkLocalDbNoInternet() bool {
	_, error := os.Stat("C:\\patchman\\index.json")
	return !errors.Is(error, os.ErrNotExist)
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
func downloadIndex(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making HTTP GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	indexmem, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error copying response body to file: %w", err)
	}

	return nil
}

func createDir() error {
	var err error
	directory, err = os.MkdirTemp(".\\", "patchman-")
	if err != nil {
		return err
	}
	return nil
}
