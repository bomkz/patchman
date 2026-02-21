package installer

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/global"
)

func createOperationsFile(opData PatchmanUnityStruct) {

	file := global.Assure(os.Create(global.Directory + "\\operations.json"))
	defer file.Close()

	jsonData := global.Assure(json.Marshal(opData))

	global.Assure(file.Write(jsonData))
}
