package main

import (
	"log"
	"os"

	"github.com/bomkz/patchman/steamutils"
	"github.com/iancoleman/orderedmap"
)

func getVtolVersion() (string, error) {

	steamPath, err := steamutils.GetSteamPath()
	if err != nil {
		log.Fatal(err)
	}

	libraryVDF, err := os.ReadFile(steamPath + "\\steamapps\\libraryfolders.vdf")
	if err != nil {
		log.Fatal(err)
	}

	steamMap, err := steamutils.Unmarshal(libraryVDF)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := steamutils.FindGameLibraryPath(steamMap, "667970")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.ReadFile(dir + "\\steamapps\\appmanifest_667970.acf")
	if err != nil {
		return "", err
	}

	acf, err := steamutils.Unmarshal(f)
	if err != nil {
		return "", err

	}

	var buildId string
	if appStateRaw, exists := acf.Get("AppState"); exists {
		if appState, ok := appStateRaw.(*orderedmap.OrderedMap); ok {
			buildIdInt, found := appState.Get("buildid")
			if found {
				buildId = buildIdInt.(string)
			}

		}
	}
	return buildId, nil

}
