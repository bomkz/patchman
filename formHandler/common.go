package formHandler

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/formHandler/installHandler"
	"github.com/bomkz/patchman/formHandler/installHandler/installer"
	"github.com/bomkz/patchman/global"
)

func savePreset() {
	preset.PatchAssetSelection = installer.Assets
	preset.PatchContentSelection = installer.Content
	presetByte := global.Assure(json.Marshal(preset))
	presetFile := global.Assure(os.Create(savePath))
	preset.Compression = installer.CompressionType
	defer presetFile.Close()

	global.Assure(presetFile.Write(presetByte))
	global.ExitAppWithMessage("Preset saved!")
}

func install(filePath string) {
	global.UnpackDependencies()
	global.UnzipIntoProgramWorkingDirectory(filePath)
	patchscript := global.Assure(os.ReadFile(global.Directory + "\\patchscript.json"))
	installHandler.HandleActionScript(patchscript)

	global.ExitTview()

	global.ExitAppWithMessage("Done!")

}

func HandleForm(indexbyte []byte, motd string) {

	// Unmarshals index to global variable
	global.AssureNoReturn(json.Unmarshal(indexbyte, &index))

	buildGameForm(motd)
}

func buildAssetList() {
	preset.AssetString = ""
	for _, x := range installer.Assets {
		if x.Modify {
			if preset.AssetString == "" {
				preset.AssetString += x.AssetName
			} else {
				preset.AssetString += "\n" + x.AssetName
			}

		}
	}
}
func buildContentList() {
	preset.ContentString = ""
	for _, x := range installer.Content {
		if x.Modify {

			if preset.ContentString == "" {
				preset.ContentString += x.ContentName
			} else {
				preset.ContentString += "\n" + x.ContentName

			}
		}
	}
}

func selectedAsset(option string, optionIndex int) {
	preset.CurrentAsset = optionIndex
}

func selectedContent(option string, optionIndex int) {
	preset.CurrentContent = optionIndex
}

func setContentList() {
	var tmpslice []string
	preset.Content = []string{}
	for _, x := range installer.Content {
		tmpslice = append(tmpslice, x.ContentName)
	}

	for _, x := range tmpslice {
		preset.Content = append(preset.Content, x)
	}
}

func checkGamePath(path string, filecheck string) bool {
	_, err := os.Stat(path + filecheck)
	return !os.IsNotExist(err)
}

func setAssetList() {
	var tmpslice []string
	preset.Assets = []string{}
	for _, x := range installer.Assets {
		tmpslice = append(tmpslice, x.AssetName)
	}

	for _, x := range tmpslice {
		preset.Assets = append(preset.Assets, x)
	}
}
