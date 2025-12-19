package actionScriptOne

import (
	"encoding/json"
	"log"
	"os"
)

// patcher.exe exportfrombundle --bundle "C:\BundlePath\unity.assets" --assetName "exampleAsset" --exportPath "C:\ExportPath\ExportName"
func batchAssetImportHandler() {
	createOperationsFile()

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

var testData ImportIntoBundleStruct = ImportIntoBundleStruct{
	BundlePath: "C:\\Program Files (x86)\\Steam\\steamapps\\common\\VTOL VR\\DLC\\1770480\\1770480",
	ImportSelection: []ImportIntoBundleSelectionStruct{
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
	SavePath: ".\\1770480.mod",
}
