package guiHandler

import (
	"encoding/json"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/installHandler"
	"github.com/bomkz/patchman/installHandler/installer"
)

func buildContentHandler() {
	detectModifiedAssetsContent()
	buildAssetContentList()

	assetCheckGroup := widget.NewCheckGroup(index[currentGame].ModifiableAssets, func(s []string) {
		preset.PatchAssetSelection = []installer.AssetSelection{}
		for _, x := range preset.Assets {
			tmpAssetSelection := installer.AssetSelection{
				AssetName: x,
				Modify:    false,
			}
			preset.PatchAssetSelection = append(preset.PatchAssetSelection, tmpAssetSelection)
		}
		for _, x := range s {
			for z, y := range preset.PatchAssetSelection {
				if y.AssetName == x {
					preset.PatchAssetSelection[z].Modify = true
				}
			}
		}
	})
	contentCheckGroup := widget.NewCheckGroup(preset.Content, func(s []string) {
		preset.PatchContentSelection = []installer.ContentSelection{}

		for _, x := range index[currentGame].ModifiableContent {
			tmpContentSelection := installer.ContentSelection{
				ContentName: x.ContentName,
				ContentPath: x.ContentPath,
				Modify:      false,
			}
			preset.PatchContentSelection = append(preset.PatchContentSelection, tmpContentSelection)
		}
		for _, x := range s {
			for y, z := range preset.PatchContentSelection {
				if z.ContentName == x {
					preset.PatchContentSelection[y].Modify = true
				}
			}
		}
	})
	contentCheckGroup.SetSelected(preset.Content)
	assetCheckGroup.SetSelected(preset.Assets)

	compressionSelect := widget.NewSelect([]string{"None", "LZMA", "LZ4"}, func(s string) { preset.Compression = s })
	fyne.DoAndWait(func() {
		global.MainWindow.SetContent(container.NewVBox(
			container.NewHBox(
				assetCheckGroup,
				contentCheckGroup,
			),
			compressionSelect,

			container.NewHBox(
				widget.NewButton("Next", func() { beginInstall() }),
				widget.NewButton("Cancel", func() { global.ExitSuccess() }),
			),
		))
	})
}

func buildAssetContentList() {
	preset.Assets = index[currentGame].ModifiableAssets

	for _, x := range index[currentGame].ModifiableContent {
		preset.Content = append(preset.Content, x.ContentName)
	}

}

func detectModifiedAssetsContent() {
	var actionScript installHandler.ActionScriptStruct

	actionscript := global.Assure(os.ReadFile(global.Directory + global.PathSeparator() + "patchscript.json"))

	global.AssureNoReturn(json.Unmarshal(actionscript, &actionScript))

	var actionScriptData []installer.ActionScriptStruct

	global.AssureNoReturn(json.Unmarshal(actionScript.Data, &actionScriptData))

	index[currentGame].ModifiableAssets = []string{}
	var baseContent IndexModifiableContentStruct
	for _, x := range index[currentGame].ModifiableContent {
		if x.ContentName == "Base Game" {
			baseContent = x
		}
	}
	index[currentGame].ModifiableContent = []IndexModifiableContentStruct{}

	var tmpContent []IndexModifiableContentStruct
	for _, x := range actionScriptData {
		if x.Action == "importasset" {
			var patchmanData installer.PatchmanUnityStruct
			global.AssureNoReturn(json.Unmarshal(x.ActionData, &patchmanData))
			for _, y := range patchmanData.Operations {
				var found bool
				for _, z := range index[currentGame].ModifiableAssets {
					if z == y.AssetName {
						found = true
					}
				}
				if !found {
					index[currentGame].ModifiableAssets = append(index[currentGame].ModifiableAssets, y.AssetName)
				}
			}
		}
		if x.Action == "importbundle" {

			var patchmanData installer.PatchmanUnityStruct
			global.AssureNoReturn(json.Unmarshal(x.ActionData, &patchmanData))

			for _, y := range index[currentGame].ModifiableContent {

				if y.ContentPath == patchmanData.OriginalFilePath {
					tmpContent = append(tmpContent, y)
				}

			}

		}
	}

	tmpContent = append(tmpContent, baseContent)

	index[currentGame].ModifiableContent = tmpContent
}

func beginInstall() {
	global.UnpackDependencies()
	patchscript := global.Assure(os.ReadFile(global.Directory + global.PathSeparator() + "patchscript.json"))

	installer.Content = preset.PatchContentSelection
	installer.Assets = preset.PatchAssetSelection
	go installHandler.HandleActionScript(patchscript)
	buildInstaller()
}
