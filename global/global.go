package global

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bomkz/patchman/steamutils"
	"github.com/iancoleman/orderedmap"
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

func CreateRoots(gameDirectory string) error {
	var err error
	Directory, err = os.MkdirTemp(".\\", "patchman-")
	if err != nil {
		return err
	}

	patchRoot, err = os.OpenRoot(Directory)
	if err != nil {
		return err
	}

	gameRoot, err = os.OpenRoot(gameDirectory)
	if err != nil {
		return err
	}
	return nil
}

func CopyFromRoot(fileName string, target string) error {

	src := filepath.Clean(fileName)
	dst := filepath.Clean(target)

	if filepath.IsAbs(src) || filepath.IsAbs(dst) || strings.HasPrefix(src, "..") || strings.HasPrefix(dst, "..") {
		return errors.New("Cannot use absolute filepaths in copy argument.")
	}

	inputFile, err := patchRoot.Open(src)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := gameRoot.Create(dst)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("Couldn't copy to dest from source: %v", err)
	}

	return nil
}

func MoveFromRoot(fileName string, target string) error {
	src := filepath.Clean(fileName)
	dst := filepath.Clean(target)

	if filepath.IsAbs(src) || filepath.IsAbs(dst) || strings.HasPrefix(src, "..") || strings.HasPrefix(dst, "..") {
		return errors.New("Cannot use absolute filepaths in copy argument.")
	}

	inputFile, err := patchRoot.Open(src)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := gameRoot.Create(dst)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("Couldn't copy to dest from source: %v", err)
	}

	inputFile.Close()

	err = os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("Couldn't remove source file: %v", err)
	}

	return nil
}

func DeleteFromGameDirectory(target string) error {

	tgt := filepath.Clean(target)

	if filepath.IsAbs(tgt) || strings.HasPrefix(tgt, "..") {
		return errors.New("Cannot use absolute filepaths in copy argument.")
	}
	return gameRoot.Remove(tgt)
}

func CleanRoot() error {
	patchRoot.Close()
	err := os.RemoveAll(Directory)
	if err != nil {
		return err
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

func UnzipIntoRoot(zipfile string) error {
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		outFile, err := patchRoot.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func GetSteamPath() (steamPath string, err error) {
	steamPath, err = steamutils.GetSteamPath()
	return
}

// Gets current VTOL Version
func GetAppIDBuildIDVersion(AppID string) (buildId string, err error) {

	libraryVDF, err := os.ReadFile(SteamPath + "\\steamapps\\libraryfolders.vdf")
	if err != nil {
		return
	}

	steamMap, err := steamutils.Unmarshal(libraryVDF)
	if err != nil {
		return
	}

	dir, err := steamutils.FindGameLibraryPath(steamMap, AppID)
	if err != nil {
		return
	}

	f, err := os.ReadFile(dir + "\\steamapps\\appmanifest_" + AppID + ".acf")
	if err != nil {
		return
	}

	acf, err := steamutils.Unmarshal(f)
	if err != nil {
		return
	}

	if appStateRaw, exists := acf.Get("AppState"); exists {
		if appState, ok := appStateRaw.(*orderedmap.OrderedMap); ok {
			buildIdInt, found := appState.Get("buildid")
			if found {
				buildId = buildIdInt.(string)
			}

		}
	}
	return

}

// Finds Path for a given AppID.
func FindAppIDPath(AppID string) (string, error) {

	// Read library vdf
	f, err := os.ReadFile(SteamPath + "\\steamapps\\libraryfolders.vdf")
	if err != nil {
		log.Fatal(err)
	}

	steamMap, err := steamutils.Unmarshal(f)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := steamutils.FindGameLibraryPath(steamMap, AppID)
	if err != nil {
		log.Fatal(err)
	}
	return dir, err

}
