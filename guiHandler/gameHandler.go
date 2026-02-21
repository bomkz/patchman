package guiHandler

import (
	"encoding/json"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc"
)

func buildGameListSteam() {

	global.AssureNoReturn(json.Unmarshal(global.IndexData, &index))

	for _, x := range index {
		gameOptions = append(gameOptions, x.AppName)
	}

	steamPathText := "Steam Path: " + global.SteamPath
	steamPathTextWidget := widget.NewLabel(steamPathText)

	gamePathTextPreset := "Game Path: "
	gamePathText := gamePathTextPreset + "None Selected"
	gamePathTextWidget := widget.NewLabel(gamePathText)

	buildIdTextPreset := "Build ID: "
	buildIdText := buildIdTextPreset + "None Selected"
	buildIdTextWidget := widget.NewLabel(buildIdText)
	var gamePath string
	gameSelect := widget.NewSelect(gameOptions, func(s string) {
		for x, y := range index {
			if y.AppName == s {
				game := ipc.GameAppById(y.AppID)
				gamePath = game.FullPath
				gamePathText = gamePathTextPreset + gamePath
				gamePathTextWidget.SetText(gamePathText)

				buildId := game.BuildID
				buildIdText = buildIdTextPreset + buildId
				buildIdTextWidget.SetText(buildIdText)
			}
			currentGame = x
		}

	})
	gameSelect.PlaceHolder = "Select game to modify"

	global.MainWindow.SetContent(container.NewVBox(
		steamPathTextWidget,
		gamePathTextWidget,
		buildIdTextWidget,
		gameSelect,
		container.NewHBox(
			widget.NewButton("Next", func() {

				global.CreateWorkingDirectories(gamePath)

				ipc.Pwd(global.Directory)
				ipc.Gwd(gamePath)

				buildPatchHandler()
			}),
			widget.NewButton("Cancel", func() { global.App.Quit() }),
		),
	))
}
