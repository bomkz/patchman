package patchScriptHandler

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller/patchScriptOne"
	"github.com/rivo/tview"
)

func HandleForm(indexbyte []byte, motd string) {

	// Unmarshals index to global variable
	global.AssureNoReturn(json.Unmarshal(indexbyte, &index))

	buildgameForm(motd)
}

// Builds game choice form
func buildgameForm(motd string) {

	// Builds game list string
	buildGameList()

	gameDropBox := tview.NewDropDown().
		SetOptions(games, selectedGame).SetLabel("Select your game.")

	form := tview.NewForm()
	form.
		AddTextView("Select your game", "", 40, 1, false, false).
		AddTextView("MOTD", motd, 40, 1, false, false).
		AddFormItem(gameDropBox).AddButton("Next", gameNext)

	form.SetBorder(false)
	global.Root.AddAndSwitchToPage("gameform", form, true)
}

func gameNext() {
	global.TargetName = games[currentGame]

	for _, x := range index {
		if x.AppName == global.TargetName {
			patchData = x.Patches
			for _, y := range x.ModifiableAssets {
				newAsset := patchScriptOne.AssetSelection{
					AssetName: y,
					Modify:    true,
				}
				patchScriptOne.Assets = append(patchScriptOne.Assets, newAsset)
			}
			for _, y := range x.ModifiableContent {
				newContent := patchScriptOne.ContentSelection{
					ContentName: y.ContentName,
					ContentPath: y.ContentPath,
					Modify:      true,
				}
				patchScriptOne.Content = append(patchScriptOne.Content, newContent)
			}
			global.TargetAppID = x.AppID
			global.TargetName = x.AppName
			Motd = x.Motd
			global.TargetPathCheck = x.LinuxPathCheck

			switch global.OsName {
			case "windows":
				global.Root.RemovePage("gameform")
				basepath := global.Assure(global.SteamReader.FindAppIDPath(x.AppID))
				global.TargetBuildID = global.Assure(global.SteamReader.FindAppIDBuildID(x.AppID))
				global.CreateWorkingDirectories(basepath + x.AppPath)
			case "linux":

				checkTextView := tview.NewTextView().SetLabel("").SetText("").SetSize(1, 20)
				path := ""
				global.Root.RemovePage("gameform")
				form2 := tview.NewForm()
				form2.AddTextView("Linux OS Detected...", "Automatic game path finding is not yet supported for linux, please insert the full path to where the game is installed below.", 20, 4, false, false).
					AddInputField("Path to game", "", 20, nil, func(tmppath string) {
						path = tmppath
					}).
					AddButton("Next", func() {
						if checkGamePath(path, global.TargetPathCheck) {
							global.Root.RemovePage("linuxpathcheck")
							global.CreateWorkingDirectories(path)
						} else {
							checkTextView.SetText("Could not find game in path provided. Ensure path is the root of game folder.")
						}
					}).
					AddFormItem(checkTextView)
				global.Root.AddAndSwitchToPage("linuxpathcheck", form2, false)
			}

			break
		}
	}

	buildGameForm()
}

func buildGameForm() {

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

				patchScriptOne.Content = preset.PatchContentSelection
				patchScriptOne.Assets = preset.PatchAssetSelection
				patchScriptOne.CompressionType = preset.Compression

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
		if patchScriptOne.Content[preset.CurrentContent].Modify {
			patchScriptOne.Content[preset.CurrentContent].Modify = false
		} else {
			patchScriptOne.Content[preset.CurrentContent].Modify = true
		}
		buildContentList()
		contentTextView.SetText(preset.ContentString)
	}

	assetFunc := func() {
		if patchScriptOne.Assets[preset.CurrentAsset].Modify {
			patchScriptOne.Assets[preset.CurrentAsset].Modify = false
		} else {
			patchScriptOne.Assets[preset.CurrentAsset].Modify = true
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

func buildGameList() {
	games = []string{}
	for _, g := range index {
		games = append(games, g.AppName)
	}

}

func selectedGame(option string, optionIndex int) {
	currentGame = optionIndex
}

func selectedVariant(option string, optionIndex int) {
	currentVariant = optionIndex
}
func selectedAsset(option string, optionIndex int) {
	preset.CurrentAsset = optionIndex
}
func selectedContent(option string, optionIndex int) {
	preset.CurrentContent = optionIndex
}
