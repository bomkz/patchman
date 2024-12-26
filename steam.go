package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/andygrunwald/vdf"
	"golang.org/x/sys/windows/registry"
)

func getSteamPath() string {
	root := registry.CURRENT_USER
	keyPath := `Software\Valve\Steam`

	SteamPath, err := readStringValueWithDefault(root, keyPath, "SteamPath", "")
	if err != nil {
		log.Fatal(err)
	}

	SteamPath = strings.ReplaceAll(SteamPath, "/", "\\")
	return SteamPath
}

// ReadStringValueWithDefault reads a string value from the Windows Registry with a default value.
func readStringValueWithDefault(root registry.Key, keyPath, valueName, defaultValue string) (string, error) {
	k, err := registry.OpenKey(root, keyPath, registry.QUERY_VALUE)
	if err != nil {
		return defaultValue, nil // Return the default value if the key or value doesn't exist
	}
	defer k.Close()

	value, _, err := k.GetStringValue(valueName)
	if err != nil {
		return defaultValue, nil // Return the default value if the value doesn't exist
	}

	return value, nil
}

func readLibraryPaths() (string, error) {
	steamPath := getSteamPath()
	f, err := os.Open(steamPath + "\\steamapps\\libraryfolders.vdf")
	if err != nil {
		panic(err)
	}

	p := vdf.NewParser(f)
	m, err := p.Parse()
	if err != nil {
		panic(err)
	}

	folder := m["libraryfolders"]
	libraries := []SteamLibraryFolder{}
	for _, y := range folder.(map[string]interface{}) {
		library := SteamLibraryFolder{}
		w, _ := json.Marshal(y)
		json.Unmarshal(w, &library)
		libraries = append(libraries, library)
	}

	for _, x := range libraries {
		vtolexists := exists(x.Path + "\\steamapps\\common\\VTOL VR")
		if vtolexists {
			return x.Path, nil
		}
	}

	return "", errors.New("VTOL VR path not found")
}
