package actionScriptOne

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"
)

// patcher.exe exportfrombundle --bundle "C:\BundlePath\unity.assets" --assetName "exampleAsset" --exportPath "C:\ExportPath\ExportName"


func runPatchmanUnity() {
	cmd := exec.Command(".\\patchman-unity.exe", "batchimportbundle", ".\\operations.json")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if out.String() == "Done!" {
		return
	} else {
		log.Panic("Uh oh...")
	}
}

func createOperationsFile() {
	file, err := os.Create("operations.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	jsonData, err := json.Marshal(testData)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

var testData PatchmanUnityStruct = PatchmanUnityStruct{
	OriginalFilePath: "C:\\Program Files (x86)\\Steam\\steamapps\\common\\VTOL VR\\DLC\\1770480\\1770480",
	Operations: []PatchmanUnityOperationsStruct{
		{
			AssetName: "ttsw_pullUp",
			AssetType: "AudioClip",
			AssetPath: "modded.resources",
			Type:      "import",
		},
		{
			AssetName: "ttsw_autopilotOff",
			AssetType: "AudioClip",
			AssetPath: "modded.resources",
			Type:      "import",
		},
	},
	ModifiedFilePath: ".\\1770480.mod",
}
