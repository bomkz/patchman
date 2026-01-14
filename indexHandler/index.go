package indexHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/bomkz/patchman/formHandler"
	"github.com/bomkz/patchman/global"
)

// Builds Index by downloading and parsing, then sends to patchScriptHandler
func BuildIndex() {

	// Download and parse Index
	handleIndex()

	// Go through Preindex content to find correct version
	var indexData []byte
	var motd string
	for _, x := range preindex.Content {
		version := global.Assure(strconv.Atoi(x.Version))
		if version == useIndexVersion {
			indexData = x.Content
			motd = x.Motd
		}
	}

	// Check if indexData is nil
	if indexData == nil {
		panic(errors.New("form content is nil"))
	}
	// Send indexData to patchScriptHandler
	formHandler.HandleForm(indexData, motd)
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

	parseIndex()

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
	} else {
		//TODO
	}

}

func loadPreIndex() (data []byte, err error) {
	if !global.Internet {
		indexmem, err = os.ReadFile("C:\\patchman\\index.json")
		if err != nil {
			return
		}
	}

	data = indexmem
	return
}

// Downloads Patchman Index from given URL and stores in indexmem
func downloadIndex(url string) (err error) {

	resp := global.Assure(http.Get(url))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	indexmem = global.Assure(io.ReadAll(resp.Body))
	return

}

func parseIndex() {
	preindexbyte := global.Assure(loadPreIndex())

	global.AssureNoReturn(json.Unmarshal(preindexbyte, &preindex))

	for _, x := range preindex.Content {
		preindexversion := global.Assure(strconv.Atoi(x.Version))

		if preindexversion == PreIndexVersion {
			useIndexVersion = preindexversion
		} else {
			useIndexVersion = 99
		}
	}

}

func checkLocalDbNoInternet() bool {
	_, error := os.Stat("C:\\patchman\\index.json")
	return !errors.Is(error, os.ErrNotExist)
}
