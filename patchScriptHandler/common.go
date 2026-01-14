package patchScriptHandler

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller/patchScriptOne"
)

func savePreset() {
	preset.PatchAssetSelection = patchScriptOne.Assets
	preset.PatchContentSelection = patchScriptOne.Content
	presetByte := global.Assure(json.Marshal(preset))
	presetFile := global.Assure(os.Create(savePath))
	preset.Compression = patchScriptOne.CompressionType
	defer presetFile.Close()

	global.Assure(presetFile.Write(presetByte))
	global.ExitAppWithMessage("Preset saved!")
}

func buildAssetList() {
	preset.AssetString = ""
	for _, x := range patchScriptOne.Assets {
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
	for _, x := range patchScriptOne.Content {
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
	for _, x := range patchScriptOne.Content {
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
	for _, x := range patchScriptOne.Assets {
		tmpslice = append(tmpslice, x.AssetName)
	}

	for _, x := range tmpslice {
		preset.Assets = append(preset.Assets, x)
	}
}
