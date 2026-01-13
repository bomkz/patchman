package patchScriptHandler

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/global"
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
		AddTextView("Select your game", "", 0, 0, false, false).
		AddTextView("MOTD", motd, 20, 3, false, false).
		AddFormItem(varDropBox).AddButton("Next", func() {
		global.TargetName = games[currentGame]

		for _, x := range index {
			if x.AppName == global.TargetName {
				patchData = x.Content
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
	setVariantList()

	descTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchDesc).
		SetLabel("Description").
		SetSize(2, 20)

	authTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchAuthor).
		SetLabel("Author").
		SetSize(2, 20)

	linkTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchLink).
		SetLabel("Source").
		SetSize(2, 20)

	motdTextView := tview.NewTextView().
		SetText(Motd).
		SetLabel("MOTD").
		SetSize(2, 20)

	varDropBox := tview.NewDropDown().
		SetOptions(variants, selectedVariant)

	form := tview.NewForm()
	form.AddTextView(global.TargetName+" BuildID", global.TargetBuildID, 0, 0, false, false).
		AddDropDown("Select Patch...", patches, 0, func(option string, optionIndex int) {
			currentSelection = optionIndex
			authTextView.SetText(patchData[currentSelection].PatchAuthor)
			descTextView.SetText(patchData[currentSelection].PatchDesc)
			linkTextView.SetText(patchData[currentSelection].PatchLink)
			setVariantList()
			varDropBox.SetOptions(variants, selectedVariant)
		}).
		AddFormItem(varDropBox).
		AddDropDown("Select Compression Type...", Compression, 0, setCompression).
		AddFormItem(authTextView).
		AddFormItem(descTextView).
		AddFormItem(linkTextView).
		AddFormItem(motdTextView).
		AddButton("Patch", beginInstall).
		AddButton("Custom", buildDeveloperForm).
		AddButton("Quit", global.ExitApp)

	form.SetBorder(false)
	global.Root.AddAndSwitchToPage("installform", form, true)

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
