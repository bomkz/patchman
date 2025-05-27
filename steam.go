package main

import (
	"os"

	"github.com/bomkz/patchman/steamutils"
	"github.com/iancoleman/orderedmap"
)

func getVtolVersion() (string, error) {
	steamPath, err := steamutils.GetSteamPath()
	if err != nil {
		return "", err

	}

	f, err := os.ReadFile(steamPath + "\\steamapps\\appmanifest_667970.acf")
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
