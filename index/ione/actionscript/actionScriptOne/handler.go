package actionScriptOne

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/bomkz/patchman/global"
)

// patcher.exe exportfrombundle --bundle "C:\BundlePath\unity.assets" --assetName "exampleAsset" --exportPath "C:\ExportPath\ExportName"

func runPatchmanUnityBundles() {
	cmd := exec.Command(".\\patchman-unity.exe", "batchimportbundle", ".\\operations.json")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		global.FatalError(err)

	}

	if out.String() == "Done!" {
		return
	} else {
		log.Panic("Uh oh...")
	}
}
func runPatchmanUnityAssets() {
	cmd := exec.Command(".\\patchman-unity.exe", "batchimportassets", ".\\operations.json")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		global.FatalError(err)

	}

	if out.String() == "Done!" {
		return
	} else {
		log.Panic("Uh oh...")
	}

}

func createOperationsFile(opData PatchmanUnityStruct) {
	file, err := os.Create("operations.json")
	if err != nil {
		global.FatalError(err)

	}

	defer file.Close()

	cleanupQueue = append(cleanupQueue, ".\\operations.json")

	jsonData, err := json.Marshal(opData)
	if err != nil {
		global.FatalError(err)

	}

	_, err = file.Write(jsonData)
	if err != nil {
		global.FatalError(err)

	}
}
