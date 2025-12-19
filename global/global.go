package global

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bomkz/patchman/steamutils"
)

func Exit() {
	StopApp <- true
}
func FatalError(err error) {
	StopApp <- true
	log.Fatal(err)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
func DownloadFile(filePath, url string) error {
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

func CleanDir() {
	os.RemoveAll(Directory)
}

func FindVtolPath() string {

	steamPath, err := steamutils.GetSteamPath()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.ReadFile(steamPath + "\\steamapps\\libraryfolders.vdf")
	if err != nil {
		log.Fatal(err)
	}
	steamMap, err := steamutils.Unmarshal(f)
	if err != nil {
		log.Fatal(err)
	}
	dir, err := steamutils.FindGameLibraryPath(steamMap, "667970")
	if err != nil {
		log.Fatal(err)
	}
	return dir + "\\steamapps\\common\\VTOL VR\\"

}
