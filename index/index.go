package index

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/patchScriptHandler"
)

// Builds Index by downloading and parsing, then sends to patchScriptHandler
func BuildIndex() error {

	// Download and parse Index
	handleIndex()

	// Go through Preindex content to find correct version
	var indexData []byte
	for _, x := range preindex.Content {
		version, err := strconv.Atoi(x.Version)
		if err != nil {
			global.FatalError(err)
		}
		if version == useIndexVersion {
			indexData = x.Content
		}
	}

	// Check if indexData is nil
	if indexData == nil {
		return errors.New("form content is nil")
	}
	// Send indexData to patchScriptHandler
	err := patchScriptHandler.HandleForm(indexData, preindex.Content[1].Motd)
	if err != nil {
		return err
	}

	return nil
}

// Downloads Index and parses it into index struct
func handleIndex() {

	err := downloadIndex(IndexURL)

	// If error exists, assume no internet connection, go offline mode.
	if err != nil {
		global.Internet = false
		exists := checkLocalDbNoInternet()
		if !exists {
			fmt.Println(err, global.NoInternetInstruct)
			os.RemoveAll(global.Directory)
			fmt.Scanln()
			os.Exit(1)
		} else {

			fmt.Println(global.NoInternet)
			fmt.Scanln()
		}
	}

	if useIndexVersion == 99 {
		global.Internet = false
		exists := checkLocalDbNoInternet()
		if !exists {
			fmt.Println("Could not find a compatible index version, reverting to offline mode.", global.NoInternetInstruct)
			os.RemoveAll(global.Directory)
			fmt.Scanln()
			os.Exit(1)
		} else {

			fmt.Println("Could not find a compatible index version, reverting to offline mode.", global.NoInternet)
			fmt.Scanln()
		}
	}

	err = parseIndex()

	if err != nil {
		os.RemoveAll(global.Directory)
		global.FatalError(err)

	}
}

func loadPreIndex() ([]byte, error) {
	if !global.Internet {
		var err error
		indexmem, err = os.ReadFile("C:\\patchman\\index.json")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return nil, err
		}
	}

	data := indexmem

	return data, nil
}

// Downloads Patchman Index from given URL and stores in indexmem
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

func parseIndex() error {
	preindexbyte, err := loadPreIndex()
	if err != nil {
		return err
	}
	err = json.Unmarshal(preindexbyte, &preindex)
	if err != nil {
		return err
	}
	for _, x := range preindex.Content {
		preindexversion, err := strconv.Atoi(x.Version)
		if err != nil {
			return err
		}
		if preindexversion == MaxPreIndexVersion {
			useIndexVersion = preindexversion
		} else {
			useIndexVersion = 99
		}
	}
	return nil
}

func checkLocalDbNoInternet() bool {
	_, error := os.Stat("C:\\patchman\\index.json")
	return !errors.Is(error, os.ErrNotExist)
}
func ReadTaint() {
	vtolvrpath, err := global.FindAppIDPath("667970")
	vtolvrpath += "\\steamapps\\common\\VTOL VR\\"
	taint, err := os.ReadFile(vtolvrpath + "\\patchman.json")
	if err != nil {
		global.ExitTview()
		global.FatalError(errors.New("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman."))
	}
	err = json.Unmarshal(taint, &global.Status)
	if err != nil {
		global.FatalError(errors.New("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman."))
	}
	global.InstalledVersion = global.Status.InstalledVersion
}
