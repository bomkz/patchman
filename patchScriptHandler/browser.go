package patchScriptHandler

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller/patchScriptOne"
	"github.com/rivo/tview"
)

func HandleForm(indexbyte []byte, motd string) {
	patchData := handleIndexJson(indexbyte)
	for _, patch := range patchData {
		patches = append(patches, patch.PatchName)
	}
	buildForm(motd)
}

func handleIndexJson(indexbyte []byte) (indexContent []IndexContentStruct) {
	indexContent = []IndexContentStruct{}
	global.AssureNoReturn(json.Unmarshal(indexbyte, &index))

	return
}

func buildForm(motd string) {

	buildGameList()
	varDropBox := tview.NewDropDown().
		SetOptions(games, selectedGame).SetLabel("Select your game.")

	form := tview.NewForm()
	form.
		AddTextView("Select your game", "", 40, 1, false, false).
		AddTextView("MOTD", motd, 40, 1, false, false).
		AddFormItem(varDropBox).AddButton("Next", func() {
		global.TargetName = games[currentGame]

		for _, x := range index {
			if x.AppName == global.TargetName {
				patchData = x.Content
				for _, y := range x.ModifiableAssets {
					newAsset := patchScriptOne.AssetSelection{
						AssetName: y,
						Modify:    true,
					}
					patchScriptOne.Assets = append(patchScriptOne.Assets, newAsset)
				}
				for _, y := range x.ModifiableContent {
					newContent := patchScriptOne.ContentSelection{
						ContentName: y.AssetName,
						ContentPath: y.AssetPath,
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
	})

	form.SetBorder(false)
	global.Root.AddAndSwitchToPage("gameform", form, true)
}

func checkGamePath(path string, filecheck string) bool {
	_, err := os.Stat(path + filecheck)
	return !os.IsNotExist(err)
}

func buildGameForm() {

	setPatchList()
	setContentList()
	setVariantList()
	setAssetList()
	buildAssetList()
	buildContentList()

	descTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchDesc).
		SetLabel("Description: ").
		SetSize(1, 40).
		SetScrollable(true)

	authTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchAuthor).
		SetLabel("Author:      ").
		SetSize(1, 40).
		SetScrollable(true)

	linkTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchLink).
		SetLabel("Source:      ").
		SetSize(1, 40).
		SetScrollable(true)

	motdTextView := tview.NewTextView().
		SetText(Motd).
		SetLabel("MOTD:        ").
		SetSize(1, 40).
		SetScrollable(true)

	varDropBox := tview.NewDropDown().
		SetOptions(variants, selectedVariant).SetLabel("Variants:    ").SetFieldWidth(40)

	assetDropBox := tview.NewDropDown().
		SetOptions(preset.Assets, selectedAsset).SetLabel("Assets:      ").SetFieldWidth(40)

	contentDropBox := tview.NewDropDown().
		SetOptions(preset.Content, selectedContent).SetLabel("Content:     ").SetFieldWidth(40)

	patchDropDown := tview.NewDropDown().SetLabel("Patch:       ").SetOptions(patches, func(option string, optionIndex int) {
		currentSelection = optionIndex
		authTextView.SetText(patchData[currentSelection].PatchAuthor)
		descTextView.SetText(patchData[currentSelection].PatchDesc)
		linkTextView.SetText(patchData[currentSelection].PatchLink)
		setVariantList()
		varDropBox.SetOptions(variants, selectedVariant)
	}).SetFieldWidth(40)

	compressionDropDown := tview.NewDropDown().SetLabel("Compression: ").SetOptions(Compression, setCompression).SetFieldWidth(40)

	assetTextView := tview.NewTextView().
		SetText(preset.AssetString).SetSize(40, 40).
		SetScrollable(true)

	contentTextView := tview.NewTextView().
		SetText(preset.ContentString).SetSize(40, 40).
		SetScrollable(true)

	textView1 := tview.NewTextView().SetLabel(global.TargetName+" BuildID "+global.TargetBuildID).SetSize(1, 40)

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

	presetButton := tview.NewButton("Presets").SetSelectedFunc(func() {
		prevpage = "installform"
		form := tview.NewForm().AddTextView("Presets", "", 0, 0, false, false).AddInputField("Path to save/load preset", "", 40, nil, func(newpath string) {
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
				global.Root.SwitchToPage(prevpage)
				buildAssetList()
				assetTextView.SetText(preset.AssetString)
				buildContentList()
				contentTextView.SetText(preset.ContentString)
			}).
			AddButton("Cancel", func() {
				global.Root.RemovePage("presetForm")
				global.Root.SwitchToPage(prevpage)
			})
		global.Root.AddAndSwitchToPage("presetForm", form, true)
	})
	presetButton.SetBorder(true)

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

	patchButton := tview.NewButton("Patch").SetSelectedFunc(beginInstall)
	patchButton.SetBorder(true)

	customButton := tview.NewButton("Custom").SetSelectedFunc(buildDeveloperForm)
	customButton.SetBorder(true)
	quitButton := tview.NewButton("Quit").SetSelectedFunc(global.ExitApp)
	quitButton.SetBorder(true)

	buttonFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(contentButton, 0, 3, false).
		AddItem(assetButton, 0, 3, false).
		AddItem(patchButton, 0, 2, false).
		AddItem(presetButton, 0, 2, false).
		AddItem(customButton, 0, 2, false).
		AddItem(quitButton, 0, 2, false)

	mainLeftFlex := tview.NewFlex().
		AddItem(textView1, 0, 1, false).
		AddItem(patchDropDown, 0, 1, false).
		AddItem(varDropBox, 0, 1, false).
		AddItem(contentDropBox, 0, 1, false).
		AddItem(assetDropBox, 0, 1, false).
		AddItem(compressionDropDown, 0, 1, false).
		AddItem(authTextView, 0, 1, false).
		AddItem(descTextView, 0, 1, false).
		AddItem(linkTextView, 0, 1, false).
		AddItem(motdTextView, 0, 1, false).
		SetDirection(tview.FlexRow).
		AddItem(buttonFlex, 0, 1, false)

	assetText := tview.NewTextView().SetText("Assets to modify (scrollable):").SetSize(2, 30).SetScrollable(true)
	contentText := tview.NewTextView().SetText("Content to modify:").SetSize(2, 18).SetScrollable(true)

	assetFlex := tview.NewFlex().AddItem(assetText, 0, 1, false).SetDirection(tview.FlexRow).AddItem(assetTextView, 0, 11, false)
	contentFlex := tview.NewFlex().AddItem(contentText, 0, 1, false).SetDirection(tview.FlexRow).AddItem(contentTextView, 0, 11, false)

	form1 := tview.NewFlex().AddItem(mainLeftFlex, 0, 10, true).SetDirection(tview.FlexColumn).AddItem(contentFlex, 0, 3, false).AddItem(assetFlex, 0, 5, false)

	global.Root.AddAndSwitchToPage("installform", form1, true)

}

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

func beginInstall() {
	for _, x := range patchData {
		if x.PatchName == patches[currentSelection] {
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

func setAssetList() {
	var tmpslice []string
	for _, x := range patchScriptOne.Assets {
		tmpslice = append(tmpslice, x.AssetName)
	}

	for _, x := range tmpslice {
		preset.Assets = append(preset.Assets, x)
	}
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

func setPatchList() {
	patches = []string{}
	for _, patch := range patchData {
		patches = append(patches, patch.PatchName)
	}
}

func setVariantList() {
	variants = []string{}
	for _, variant := range patchData[currentSelection].PatchVariants {
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
