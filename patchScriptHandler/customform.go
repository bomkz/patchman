package patchScriptHandler

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller/patchScriptOne"
	"github.com/rivo/tview"
)

var Compression = []string{
	"None",
	"LZMA",
	"LZ4",
	"LZ4Fast",
}

func buildDeveloperForm() {
	global.Root.RemovePage("installform")

	setContentList()
	setAssetList()
	buildAssetList()
	buildContentList()

	//---------------------------------\\
	assetDropDown := tview.NewDropDown().
		SetOptions(preset.Assets, selectedAsset).SetLabel("Assets:      ").SetFieldWidth(40)
	contentDropDown := tview.NewDropDown().
		SetOptions(preset.Content, selectedContent).SetLabel("Content:     ").SetFieldWidth(40)
	compressionDropDown := tview.NewDropDown().SetLabel("Compression: ").
		SetOptions(Compression, setCompression).
		SetFieldWidth(40)
	//---------------------------------\\
	assetTextView := tview.NewTextView().
		SetText(preset.AssetString).SetSize(40, 40).
		SetScrollable(true)
	contentTextView := tview.NewTextView().
		SetText(preset.ContentString).SetSize(40, 40).
		SetScrollable(true)
	titleTextView := tview.NewTextView().
		SetLabel(global.TargetName+" BuildID "+global.TargetBuildID).
		SetSize(1, 40)
	//-------------------------------------------------------------------\\
	assetButton := tview.NewButton("Toggle Asset").SetSelectedFunc(func() {
		if patchScriptOne.Assets[preset.CurrentAsset].Modify {
			patchScriptOne.Assets[preset.CurrentAsset].Modify = false
		} else {
			patchScriptOne.Assets[preset.CurrentAsset].Modify = true
		}
		buildAssetList()
		assetTextView.SetText(preset.AssetString)
	})
	assetButton.SetBorder(true)

	contentButton := tview.NewButton("Toggle Content").SetSelectedFunc(func() {
		if patchScriptOne.Content[preset.CurrentContent].Modify {
			patchScriptOne.Content[preset.CurrentContent].Modify = false
		} else {
			patchScriptOne.Content[preset.CurrentContent].Modify = true
		}
		buildContentList()
		contentTextView.SetText(preset.ContentString)
	})
	contentButton.SetBorder(true)

	patchButton := tview.NewButton("Patch").
		SetSelectedFunc(installFunc)
	patchButton.SetBorder(true)

	quitButton := tview.NewButton("Quit").
		SetSelectedFunc(global.ExitApp)
	quitButton.SetBorder(true)

	presetButton := tview.NewButton("Presets").SetSelectedFunc(func() {
		form := tview.NewForm().
			AddTextView("Presets", "", 0, 0, false, false).
			AddInputField("Path to save/load preset", "", 40, nil, func(newpath string) {
				savePath = newpath
			}).
			AddButton("Save", savePreset).
			AddButton("Load", func() {
				jsonByte := global.Assure(os.ReadFile(savePath))
				global.AssureNoReturn(json.Unmarshal(jsonByte, &preset))

				patchScriptOne.Content = preset.PatchContentSelection
				patchScriptOne.Assets = preset.PatchAssetSelection
				patchScriptOne.CompressionType = preset.Compression

				global.Root.RemovePage("presetForm")
				global.Root.SwitchToPage("custom")
				buildAssetList()
				assetTextView.SetText(preset.AssetString)
				buildContentList()
				contentTextView.SetText(preset.ContentString)
			}).
			AddButton("Cancel", func() {
				global.Root.RemovePage("presetForm")
				global.Root.SwitchToPage("custom")
			})
		global.Root.AddAndSwitchToPage("presetForm", form, true)
	})
	presetButton.SetBorder(true)
	//---------------------------------\\
	customField := tview.NewInputField().
		SetChangedFunc(pathField).
		SetLabel("Custom patch path: ")

	assetText := tview.NewTextView().
		SetText("Assets to modify (scrollable):").
		SetSize(2, 30).
		SetScrollable(true)

	contentText := tview.NewTextView().
		SetText("Content to modify:").
		SetSize(2, 18).
		SetScrollable(true)
	//---------------------------------------------------------\\
	buttonFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(contentButton, 0, 3, false).
		AddItem(assetButton, 0, 3, false).
		AddItem(presetButton, 0, 2, false).
		AddItem(patchButton, 0, 2, false).
		AddItem(quitButton, 0, 2, false)

	mainFlex := tview.NewFlex().
		AddItem(titleTextView, 0, 1, false).
		AddItem(customField, 0, 1, false).
		AddItem(contentDropDown, 0, 1, false).
		AddItem(assetDropDown, 0, 1, false).
		AddItem(compressionDropDown, 0, 1, false).
		SetDirection(tview.FlexRow).
		AddItem(buttonFlex, 0, 1, false)
	//-------------------------\\
	assetFlex := tview.NewFlex().
		AddItem(assetText, 0, 1, false).
		SetDirection(tview.FlexRow).
		AddItem(assetTextView, 0, 11, false)

	contentFlex := tview.NewFlex().
		AddItem(contentText, 0, 1, false).
		SetDirection(tview.FlexRow).
		AddItem(contentTextView, 0, 11, false)
	//-----------------------------\\
	containerFlex := tview.NewFlex().
		AddItem(mainFlex, 0, 10, true).
		SetDirection(tview.FlexColumn).
		AddItem(contentFlex, 0, 3, false).
		AddItem(assetFlex, 0, 5, false)

	global.Root.AddAndSwitchToPage("custom", containerFlex, true)

}

func setCompression(option string, optionIndex int) {
	patchScriptOne.CompressionType = option
}
func installFunc() {
	install(modPath)
	global.Root.RemovePage("DevForm")
}

func pathField(path string) {
	modPath = path
}
