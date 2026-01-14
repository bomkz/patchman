package formHandler

import (
	"github.com/bomkz/patchman/global"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func buildCustomGame() {
	var appId string
	var appPath string

	appIdChanged := func(text string) {
		appId = text
	}

	appPathChanged := func(text string) {
		appPath = text
	}

	checkTextView := tview.NewTextView()

	nextButtonFunc := func() {
		var found bool
		for _, x := range index {
			if x.AppID == appId {
				if checkGamePath(appPath, global.TargetPathCheck) {
					global.Root.RemovePage("customgame")
					global.CreateWorkingDirectories(appPath)

					for z, y := range games {
						if x.AppName == y {
							currentGame = z
							found = true
							break
						}
					}
				} else {
					checkTextView.SetText("Could not find game in path provided. Ensure path is the root of game folder.")
				}
			}
		}
		if !found {
			var tmpgame = IndexStruct{
				AppName:           "custom",
				AppID:             appId,
				AppPath:           appPath,
				ModifiableAssets:  []string{"none"},
				ModifiableContent: []IndexModifiableContentStruct{IndexModifiableContentStruct{ContentName: "none", ContentPath: "none"}},
			}

			index = append(index, tmpgame)
			buildGameList()

			for z, y := range games {
				if y == "custom" {
					currentGame = z
					break
				}
			}
			buildCustomForm()
		}
	}

	appIdInput := tview.NewInputField().SetChangedFunc(appIdChanged).SetTitle("AppID (Optional): ")

	appPathInput := tview.NewInputField().SetChangedFunc(appPathChanged).SetTitle("AppPath (Required):")

	nextButton := tview.NewButton("Next").SetSelectedFunc(nextButtonFunc)

	mainFlex := tview.NewFlex().
		AddItem(checkTextView, 0, 1, false).
		AddItem(appIdInput, 0, 1, false).
		AddItem(appPathInput, 0, 1, false).
		AddItem(nextButton, 0, 1, false)

	mainFlex.
		SetTitle("Custom Game").
		SetTitleColor(tcell.ColorLightPink)

	global.Root.AddAndSwitchToPage("customgameform", mainFlex, true)
}
