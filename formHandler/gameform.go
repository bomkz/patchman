package formHandler

import (
	"github.com/bomkz/patchman/formHandler/installHandler/installer"
	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

// Builds game choice form
func buildGameForm(motd string) {

	buildGameList()

	gameDropBox := tview.NewDropDown().
		SetOptions(games, selectedGame).SetLabel("Select your game.")

	form := tview.NewForm()
	form.
		AddTextView("Select your game", "", 40, 1, false, false).
		AddTextView("MOTD", motd, 120, 4, false, true).
		AddFormItem(gameDropBox).
		AddButton("Custom", buildCustomGame).
		AddButton("Next", gameNext)

	form.SetBorder(false)
	global.Root.AddAndSwitchToPage("gameform", form, true)
}

// Handles gameNext button
func gameNext() {
	global.TargetName = games[currentGame]

	for _, x := range index {
		if x.AppName == global.TargetName {
			patchData = x.Patches
			for _, y := range x.ModifiableAssets {
				newAsset := installer.AssetSelection{
					AssetName: y,
					Modify:    true,
				}
				installer.Assets = append(installer.Assets, newAsset)
			}
			for _, y := range x.ModifiableContent {
				newContent := installer.ContentSelection{
					ContentName: y.ContentName,
					ContentPath: y.ContentPath,
					Modify:      true,
				}
				installer.Content = append(installer.Content, newContent)
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

	buildInstallForm()
}

// Sets selection index
func selectedGame(option string, optionIndex int) {
	currentGame = optionIndex
}

// Builds game list string
func buildGameList() {
	games = []string{}
	for _, g := range index {
		games = append(games, g.AppName)
	}

}
