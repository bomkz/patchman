package main

import (
	"archive/zip"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/andygrunwald/vdf"
	"golang.org/x/sys/windows/registry"
)

func main() {

	checkAH94Installed()
	checkEF24GInstalled()

	removeResources()
	removeAH94Files()
	removeEF24GFiles()

	unpackFiles()
	install()
	cleanup()

}

func cleanup() {
	os.Remove(".\\resources.assets")
	os.Remove(".\\resources.assets.resS")
	os.Remove(".\\resources.resource")
	os.Remove(".\\1770480")
	os.Remove(".\\2531290")
	os.Remove(".\\CAB-609a7bd01976702a18d81971aebebeea.resource")
	os.Remove(".\\CAB-609a7bd01976702a18d81971aebebeea.resS")
	os.Remove(".\\CAB-db515831ae078197daa2fd6af388d061.resource")
	os.Remove(".\\CAB-db515831ae078197daa2fd6af388d061.resS")

}

func unpackFiles() {
	if err := os.WriteFile(".\\installer.zip", installerfiles, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	unzip(".\\installer.zip", ".\\")
	var err error

	ah94file, err := os.Open(".\\1770480")
	if err != nil {
		log.Fatal(err)
	}
	defer ah94file.Close()
	ah94, err = io.ReadAll(ah94file)
	if err != nil {
		log.Fatal(err)
	}

	ah94resSfile, err := os.Open(".\\CAB-609a7bd01976702a18d81971aebebeea.resS")
	if err != nil {
		log.Fatal(err)
	}
	defer ah94resSfile.Close()
	ah94resS, err = io.ReadAll(ah94resSfile)
	if err != nil {
		log.Fatal(err)
	}

	ah94resourcefile, err := os.Open(".\\CAB-609a7bd01976702a18d81971aebebeea.resource")
	if err != nil {
		log.Fatal(err)
	}
	defer ah94resourcefile.Close()
	ah94resource, err = io.ReadAll(ah94resourcefile)
	if err != nil {
		log.Fatal(err)
	}

	ef24gfile, err := os.Open(".\\2531290")
	if err != nil {
		log.Fatal(err)
	}
	defer ef24gfile.Close()
	ef24g, err = io.ReadAll(ef24gfile)
	if err != nil {
		log.Fatal(err)
	}

	ef24gresourcefile, err := os.Open(".\\CAB-db515831ae078197daa2fd6af388d061.resource")
	if err != nil {
		log.Fatal(err)
	}
	defer ef24gresourcefile.Close()
	ef24gresource, err = io.ReadAll(ef24gresourcefile)
	if err != nil {
		log.Fatal(err)
	}

	ef24gresSfile, err := os.Open(".\\CAB-db515831ae078197daa2fd6af388d061.resS")
	if err != nil {
		log.Fatal(err)
	}
	defer ef24gresSfile.Close()
	ef24gresS, err = io.ReadAll(ef24gresSfile)
	if err != nil {
		log.Fatal(err)
	}

	resourcesassetsfile, err := os.Open(".\\resources.assets")
	if err != nil {
		log.Fatal(err)
	}
	defer resourcesassetsfile.Close()
	resourcesassets, err = io.ReadAll(resourcesassetsfile)
	if err != nil {
		log.Fatal(err)
	}

	resourcesassetsressfile, err := os.Open(".\\resources.assets.resS")
	if err != nil {
		log.Fatal(err)
	}
	defer resourcesassetsressfile.Close()
	resourcesassetsress, err = io.ReadAll(resourcesassetsressfile)
	if err != nil {
		log.Fatal(err)
	}

	resourcesresourcefile, err := os.Open(".\\resources.resource")
	if err != nil {
		log.Fatal(err)
	}
	defer resourcesresourcefile.Close()
	resourcesresource, err = io.ReadAll(resourcesresourcefile)
	if err != nil {
		log.Fatal(err)
	}

}

func unzip(source, dest string) error {
	read, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer read.Close()
	for _, file := range read.File {
		if file.Mode().IsDir() {
			continue
		}
		open, err := file.Open()
		if err != nil {
			return err
		}
		name := path.Join(dest, file.Name)
		os.MkdirAll(path.Dir(name), os.ModeDir)
		create, err := os.Create(name)
		if err != nil {
			return err
		}
		defer create.Close()
		create.ReadFrom(open)
	}
	return nil
}

func install() {
	vtolvrpath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	vtolvrpath += "\\steamapps\\common\\VTOL VR"

	if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\resources.resource", resourcesresource, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\resources.assets", resourcesassets, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\resources.assets.resS", resourcesassetsress, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	if AH94Installed {
		if err := os.WriteFile(vtolvrpath+"\\DLC\\1770480\\1770480", ah94, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\CAB-609a7bd01976702a18d81971aebebeea.resource", ah94resource, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\CAB-609a7bd01976702a18d81971aebebeea.resS", ah94resS, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	}
	if EF24GInstalled {
		if err := os.WriteFile(vtolvrpath+"\\DLC\\2531290\\2531290", ef24g, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\CAB-db515831ae078197daa2fd6af388d061.resource", ef24gresource, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		if err := os.WriteFile(vtolvrpath+"\\VTOLVR_Data\\CAB-db515831ae078197daa2fd6af388d061.resS", ef24gresS, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	}
}

func removeResources() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\resources.assets")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\resources.assets.resS")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\resources.resource")
	if err != nil {
		fmt.Println(err)
	}
}

func removeAH94Files() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\1770480\\1770480")
	if err != nil {
		fmt.Println(err)
	}
	if AH94CabResSInstalled {
		err := os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\CAB-609a7bd01976702a18d81971aebebeea.resS")
		if err != nil {
			fmt.Println(err)
		}
	}
	if AH94CabResourceInstalled {
		err := os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\CAB-609a7bd01976702a18d81971aebebeea.resource")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func removeEF24GFiles() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\2531290\\2531290")
	if err != nil {
		fmt.Println(err)
	}
	if EF24GCabResSInstalled {
		err := os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\CAB-db515831ae078197daa2fd6af388d061.resS")
		if err != nil {
			fmt.Println(err)
		}
	}
	if EF24GCabResourceInstalled {
		err := os.Remove(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\CAB-db515831ae078197daa2fd6af388d061.resource")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func checkAH94Installed() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	exist := exists(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\1770480")
	if !exist {
		AH94Installed = false
	} else {
		AH94Installed = true
	}
	exist = exists(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\1770480\\1770480.manifest")
	if !exist {
		AH94Installed = false
	} else {
		AH94Installed = true
	}
	AH94CabResSInstalled = exists(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\CAB-609a7bd01976702a18d81971aebebeea.resS")
	AH94CabResourceInstalled = exists(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\CAB-609a7bd01976702a18d81971aebebeea.resource")

}

var AH94CabResourceInstalled = false
var AH94CabResSInstalled = false
var EF24GCabResourceInstalled = false
var EF24GCabResSInstalled = false

func checkEF24GInstalled() {
	libraryPath, err := readLibraryPaths()
	if err != nil {
		log.Fatal(err)
	}
	exist := exists(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\2531290")
	if !exist {
		EF24GInstalled = false
	} else {
		EF24GInstalled = true
	}
	exist = exists(libraryPath + "\\steamapps\\common\\VTOL VR\\DLC\\2531290\\2531290.manifest")
	if !exist {
		EF24GInstalled = false
	} else {
		EF24GInstalled = true
	}

	EF24GCabResSInstalled = exists(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\CAB-db515831ae078197daa2fd6af388d061.resS")
	EF24GCabResourceInstalled = exists(libraryPath + "\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\CAB-db515831ae078197daa2fd6af388d061.resource")

}

var AH94Installed = false
var EF24GInstalled = false

func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
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

type SteamLibraryFolder struct {
	Path string `json:"path,omitempty"`
}

//go:embed diffs.zip
var installerfiles []byte

var resourcesresource []byte
var resourcesassets []byte
var resourcesassetsress []byte
var ah94 []byte
var ah94resource []byte
var ah94resS []byte
var ef24g []byte
var ef24gresource []byte
var ef24gresS []byte
