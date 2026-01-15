package installer

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/bomkz/patchman/global"
)

// patcher.exe exportfrombundle --bundle "C:\BundlePath\unity.assets" --assetName "exampleAsset" --exportPath "C:\ExportPath\ExportName"

func runPatchmanUnityBundles() {

	switch global.OsName {
	case "windows":
		cmd := exec.Command(global.Directory+"\\patchman-unity.exe", "batchimportbundle", global.Directory+".\\operations.json", CompressionType)
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			panic(err)
		}

		if out.String() != "Done!" {
			panic(out.String())
		}

	case "linux":
		cmd := exec.Command(global.Directory+"\\patchman-unity", "batchimportbundle", global.Directory+".\\operations.json", CompressionType)
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			panic(err)
		}

		if out.String() != "Done!" {
			panic(out.String())
		}

	}

}

func runPatchmanUnityAssets() {
	switch global.OsName {
	case "windows":
		cmd := exec.Command(global.Directory+"\\patchman-unity.exe", "batchimportasset", global.Directory+".\\operations.json")
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			panic(err)
		}

		if out.String() != "Done!" {
			panic(out.String())
		}
	case "linux":
		cmd := exec.Command(global.Directory+"\\patchman-unity", "batchimportasset", global.Directory+".\\operations.json")
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			panic(err)
		}

		if out.String() != "Done!" {
			panic(out.String())
		}
	}
}

func createOperationsFile(opData PatchmanUnityStruct) {

	file := global.Assure(os.Create(global.Directory + "\\operations.json"))
	defer file.Close()

	jsonData := global.Assure(json.Marshal(opData))

	global.Assure(file.Write(jsonData))
}
