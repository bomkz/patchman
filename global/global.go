package global

import (
	"log"
	"os"

	"github.com/bomkz/patchman/steamutils"
)

func FatalError(err error) {
	StopApp <- true
	log.Fatal(err)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CleanDir() {
	os.RemoveAll(Directory)
}

func FindVtolPath() string {

	steamPath, err := steamutils.GetSteamPath()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.ReadFile(steamPath + "\\steamapps\\libraryfolders.vdf")
	if err != nil {
		log.Fatal(err)
	}
	steamMap, err := steamutils.Unmarshal(f)
	if err != nil {
		log.Fatal(err)
	}
	dir, err := steamutils.FindGameLibraryPath(steamMap, "667970")
	if err != nil {
		log.Fatal(err)
	}
	return dir + "\\steamapps\\common\\VTOL VR\\"

}
