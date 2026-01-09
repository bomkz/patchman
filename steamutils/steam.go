package steamutils

import (
	"fmt"
	"strings"

	"github.com/iancoleman/orderedmap"
	"golang.org/x/sys/windows/registry"
)

// Goes through all libraries in libraryfolders.vdf to find the path of the library containing the target appid
func FindGameLibraryPath(m *orderedmap.OrderedMap, targetAppID string) (string, error) {
	libFoldersVal, exists := m.Get("libraryfolders")
	if !exists {
		return "", fmt.Errorf("libraryfolders key not found in the VDF data")
	}

	libFolders, ok := libFoldersVal.(*orderedmap.OrderedMap)
	if !ok {
		return "", fmt.Errorf("libraryfolders is not of the expected type")
	}

	for _, libKey := range libFolders.Keys() {
		libraryVal, exists := libFolders.Get(libKey)
		if !exists {
			continue
		}

		library, ok := libraryVal.(*orderedmap.OrderedMap)
		if !ok {
			continue
		}

		appsVal, exists := library.Get("apps")
		if !exists {
			continue
		}

		apps, ok := appsVal.(*orderedmap.OrderedMap)
		if !ok {
			continue
		}

		for _, appKey := range apps.Keys() {
			if appKey == targetAppID {
				pathVal, exists := library.Get("path")
				if !exists {
					return "", fmt.Errorf("path not found in library %s for app %s", libKey, targetAppID)
				}
				pathStr, ok := pathVal.(string)
				if !ok {
					return "", fmt.Errorf("library path for library %s is not a string", libKey)
				}

				return pathStr, nil
			}
		}
	}

	return "", fmt.Errorf("app with appid %s not found in any library", targetAppID)
}

// Finds Steam's Path from Windows Registry
func GetSteamPath() (string, error) {
	root := registry.CURRENT_USER
	keyPath := `Software\Valve\Steam`

	SteamPath, err := readStringValueWithDefault(root, keyPath, "SteamPath", "")
	if err != nil {
		return "", err
	}

	SteamPath = strings.ReplaceAll(SteamPath, "/", "\\")
	return SteamPath, nil
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

// The code in this file was made by ChatGPT, use in production is highly discouraged as unexpected results may occur. The code in this file is not vetted for stability or edge cases.
