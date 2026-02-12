package indexHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bomkz/patchman/global"
	"tawesoft.co.uk/go/dialog"
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
		global.Internet = false
		dialog.Alert("%s", global.NoInternet)

	}
	// Send indexData to patchScriptHandler

	global.IndexData = indexData
	global.IndexMotd = motd
}

// Downloads Index and parses it into index struct
func handleIndex() {

	err := downloadIndex(IndexURL)

	// If error exists, assume no internet connection, go offline mode.
	if err != nil {
		global.Internet = false
		dialog.Alert("%s", global.NoInternet)
		return
	}

	parseIndex()
	if PreIndexVersion == 99 {
		global.Internet = false
		dialog.Alert("%s", global.NoInternet)
		return

	}
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

	if err := json.Unmarshal(indexmem, &preindex); err != nil {
		global.Internet = false
		dialog.Alert("%s", global.NoInternet)
		return
	}

	for _, x := range preindex.Content {
		preindexversion := global.Assure(strconv.Atoi(x.Version))

		if preindexversion == PreIndexVersion {
			useIndexVersion = preindexversion
		} else {
			useIndexVersion = 99
		}
	}

}
