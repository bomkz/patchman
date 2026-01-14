package formHandler

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/formHandler/installHandler/installer"
	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func buildInstallForm() {

	setPatchList()
	setContentList()
	setVariantList()
	setAssetList()
	buildAssetList()
	buildContentList()

	//-------------------------------------------\\
	descTextView := tview.NewTextView().
		SetText(patchData[currentPatch].PatchDesc).
		SetLabel("Description: ").
		SetSize(1, 40).
		SetScrollable(true)

	authTextView := tview.NewTextView().
		SetText(patchData[currentPatch].PatchAuthor).
		SetLabel("Author:      ").
		SetSize(1, 40).
		SetScrollable(true)

	linkTextView := tview.NewTextView().
		SetText(patchData[currentPatch].PatchLink).
		SetLabel("Source:      ").
		SetSize(1, 40).
		SetScrollable(true)

	motdTextView := tview.NewTextView().
		SetText(Motd).
		SetLabel("MOTD:        ").
		SetSize(1, 40).
		SetScrollable(true)

	assetTextView := tview.NewTextView().
		SetText(preset.AssetString).SetSize(40, 40).
		SetScrollable(true)

	contentTextView := tview.NewTextView().
		SetText(preset.ContentString).SetSize(40, 40).
		SetScrollable(true)

	titleTextView := tview.NewTextView().
		SetLabel(global.TargetName+" BuildID "+global.TargetBuildID).
		SetSize(1, 40).
		SetScrollable(false)

	assetText := tview.NewTextView().
		SetText("Assets to modify (scrollable):").
		SetSize(2, 30).
		SetScrollable(true)

	contentText := tview.NewTextView().
		SetText("Content to modify:").
		SetSize(2, 18).
		SetScrollable(true)
	//-----------------------------------\\
	variantDropDown := tview.NewDropDown().
		SetOptions(variants, selectedVariant).
		SetLabel("Variants:    ").
		SetFieldWidth(40)

	assetDropDown := tview.NewDropDown().
		SetOptions(preset.Assets, selectedAsset).
		SetLabel("Assets:      ").
		SetFieldWidth(40)

	contentDropDown := tview.NewDropDown().
		SetOptions(preset.Content, selectedContent).
		SetLabel("Content:     ").
		SetFieldWidth(40)

	patchDropDown := tview.NewDropDown().
		SetLabel("Patch:       ").
		SetOptions(patches, func(option string, optionIndex int) {
			currentPatch = optionIndex
			authTextView.SetText(patchData[currentPatch].PatchAuthor)
			descTextView.SetText(patchData[currentPatch].PatchDesc)
			linkTextView.SetText(patchData[currentPatch].PatchLink)
			setVariantList()
			variantDropDown.SetOptions(variants, selectedVariant)
		}).
		SetFieldWidth(40)

	compressionDropDown := tview.NewDropDown().
		SetLabel("Compression: ").
		SetOptions(Compression, setCompression).
		SetFieldWidth(40)
	//------------------\\
	presetFunc := func() {
		form := tview.NewForm().
			AddTextView("Presets", "", 0, 0, false, false).
			AddInputField("Path to save/load preset", "", 40, nil, func(newpath string) {
				savePath = newpath
			}).
			AddButton("Save", savePreset).
			AddButton("Load", func() {
				jsonByte := global.Assure(os.ReadFile(savePath))
				global.AssureNoReturn(json.Unmarshal(jsonByte, &preset))

				installer.Content = preset.PatchContentSelection
				installer.Assets = preset.PatchAssetSelection
				installer.CompressionType = preset.Compression

				global.Root.RemovePage("presetForm")
				global.Root.SwitchToPage("installform")
				buildAssetList()
				assetTextView.SetText(preset.AssetString)
				buildContentList()
				contentTextView.SetText(preset.ContentString)
			}).
			AddButton("Cancel", func() {
				global.Root.RemovePage("presetForm")
				global.Root.SwitchToPage("installform")
			})
		global.Root.AddAndSwitchToPage("presetForm", form, true)
	}

	contentFunc := func() {
		if installer.Content[preset.CurrentContent].Modify {
			installer.Content[preset.CurrentContent].Modify = false
		} else {
			installer.Content[preset.CurrentContent].Modify = true
		}
		buildContentList()
		contentTextView.SetText(preset.ContentString)
	}

	assetFunc := func() {
		if installer.Assets[preset.CurrentAsset].Modify {
			installer.Assets[preset.CurrentAsset].Modify = false
		} else {
			installer.Assets[preset.CurrentAsset].Modify = true
		}
		buildAssetList()
		assetTextView.SetText(preset.AssetString)
	}
	//-------------------------------------------------------------------\\
	assetButton := tview.NewButton("Toggle Asset").SetSelectedFunc(assetFunc)
	assetButton.SetBorder(true)

	presetButton := tview.NewButton("Presets").
		SetSelectedFunc(presetFunc)
	presetButton.SetBorder(true)
	contentButton := tview.NewButton("Toggle Content").
		SetSelectedFunc(contentFunc)
	contentButton.SetBorder(true)

	patchButton := tview.NewButton("Patch").
		SetSelectedFunc(beginInstall)
	patchButton.SetBorder(true)

	customButton := tview.NewButton("Custom").
		SetSelectedFunc(buildDeveloperForm)
	customButton.SetBorder(true)

	quitButton := tview.NewButton("Quit").
		SetSelectedFunc(global.ExitApp)
	quitButton.SetBorder(true)
	//---------------------------------------------------------\\
	buttonFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(contentButton, 0, 3, false).
		AddItem(assetButton, 0, 3, false).
		AddItem(patchButton, 0, 2, false).
		AddItem(presetButton, 0, 2, false).
		AddItem(customButton, 0, 2, false).
		AddItem(quitButton, 0, 2, false)
	//------------------------\\
	mainFlex := tview.NewFlex().
		AddItem(titleTextView, 0, 1, false).
		AddItem(patchDropDown, 0, 1, false).
		AddItem(variantDropDown, 0, 1, false).
		AddItem(contentDropDown, 0, 1, false).
		AddItem(assetDropDown, 0, 1, false).
		AddItem(compressionDropDown, 0, 1, false).
		AddItem(authTextView, 0, 1, false).
		AddItem(descTextView, 0, 1, false).
		AddItem(linkTextView, 0, 1, false).
		AddItem(motdTextView, 0, 1, false).
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

	global.Root.AddAndSwitchToPage("installform", containerFlex, true)

}

func beginInstall() {
	for _, x := range patchData {
		if x.PatchName == patches[currentPatch] {
			for _, y := range x.PatchVariants {
				if y.Variant == variants[currentVariant] {
					global.DownloadFileToProgramWorkingDirectory("patchman.zip", y.DownloadLink)
				}
			}
		}
	}

	install("patchman.zip")
	global.Root.RemovePage("installform")
}

func setPatchList() {
	patches = []string{}
	for _, patch := range patchData {
		patches = append(patches, patch.PatchName)
	}
}

func setVariantList() {
	variants = []string{}
	for _, variant := range patchData[currentPatch].PatchVariants {
		variants = append(variants, variant.Variant)
	}

}

func selectedVariant(option string, optionIndex int) {
	currentVariant = optionIndex
}
