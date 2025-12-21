package global

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bomkz/patchman/steamutils"
	"github.com/inancgumus/screen"
)

func ExitTview() {
	App.Stop()
	screen.Clear()
}

func ExitApp() {
	ExitTview()
	screen.Clear()
	os.Exit(0)
}

func ExitAppWithMessage(message string) {
	ExitTview()
	screen.Clear()
	fmt.Println(message)
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
	os.Exit(0)
}

func FatalError(err error) {
	ExitTview()
	fmt.Println(err)
	fmt.Println("Press Enter to exit")
	fmt.Scanln()

	ExitApp()

}

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("Couldn't open dest file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("Couldn't copy to dest from source: %v", err)
	}

	inputFile.Close() // for Windows, close before trying to remove: https://stackoverflow.com/a/64943554/246801

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't remove source file: %v", err)
	}
	return nil
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
