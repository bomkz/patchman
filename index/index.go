package index

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/index/ione"
	"github.com/bomkz/patchman/index/izero"
	"github.com/rivo/tview"
)

func BuildIndex() (*tview.Form, error) {

	handleIndex()

	switch useIndexVersion {
	case 0:
		var indexData []byte
		for _, x := range preindex.Content {
			version, err := strconv.Atoi(x.Version)
			if err != nil {
				log.Panic(err)
			}
			if version == useIndexVersion {
				indexData = x.Content
			}
		}

		if indexData == nil {
			return nil, errors.New("form content is nil")
		}

		form, err := izero.HandleForm(indexData)

		if err != nil {
			return nil, err
		}

		return form, nil
	case 1:
		var indexData []byte
		for _, x := range preindex.Content {
			version, err := strconv.Atoi(x.Version)
			if err != nil {
				log.Panic(err)
			}
			if version == useIndexVersion {
				indexData = x.Content
			}
		}

		if indexData == nil {
			return nil, errors.New("form content is nil")
		}
		form, err := ione.HandleForm(indexData)

		if err != nil {
			return nil, err
		}

		return form, nil
	}

	return nil, errors.New("could not handle index idk why. good luk")
}

func handleIndex() {

	err := downloadIndex(IndexURL)
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

	err = parseIndex()

	if err != nil {
		os.RemoveAll(global.Directory)
		log.Fatal(err)
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
		if preindexversion <= MaxPreIndexVersion {
			useIndexVersion = preindexversion
		}
	}
	return nil
}

func checkLocalDbNoInternet() bool {
	_, error := os.Stat("C:\\patchman\\index.json")
	return !errors.Is(error, os.ErrNotExist)
}
func ReadTaint() {
	vtolvrpath := global.FindVtolPath()

	taint, err := os.ReadFile(vtolvrpath + "\\patchman.json")
	if err != nil {
		global.StopApp <- true
		log.Fatal("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman.")
	}
	err = json.Unmarshal(taint, &global.Status)
	if err != nil {
		log.Fatal("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman.")
	}
	global.InstalledVersion = global.Status.InstalledVersion
}

func TaintInfo() string {
	if global.InstalledVersion == 0 {
		return izero.BuildTaintInfo()
	}
	return ""
}
