package guiHandler

import (
	"encoding/json"
	"os"
	"strings"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc"
	alert "tawesoft.co.uk/go/dialog"
)

func buildGameListSteam() {

	if global.Internet {
		global.AssureNoReturn(json.Unmarshal(global.IndexData, &index))
	}

	customGameDir := global.Assure(os.ReadDir(global.Assure(os.UserHomeDir()) + PathSeparator() + "patchman"))
	for _, x := range customGameDir {

		gameByte := global.Assure(os.ReadFile(global.Assure(os.UserHomeDir()) + PathSeparator() + "patchman" + PathSeparator() + x.Name()))
		var tmpGame CustomGame
		if err := json.Unmarshal(gameByte, &tmpGame); err != nil {
			continue
		}
		tmpIndex := IndexStruct{
			AppName: "Custom - " + tmpGame.Name,
			AppPath: tmpGame.Location,
		}
		for _, y := range tmpGame.Content {
			tmpModifiableContent := IndexModifiableContentStruct{
				ContentName: y.Name,
				ContentPath: y.Path,
			}
			tmpIndex.ModifiableContent = append(tmpIndex.ModifiableContent, tmpModifiableContent)
		}
		index = append(index, tmpIndex)
	}

	gameOptions = append(gameOptions, "Custom")

	for _, x := range index {
		gameOptions = append(gameOptions, x.AppName)
	}

	steamPathText := "Steam Path: " + global.SteamPath
	steamPathTextWidget := widget.NewLabel(steamPathText)
	motdText := widget.NewLabel(global.IndexMotd)

	gamePathTextPreset := "Game Path: "
	gamePathText := gamePathTextPreset + "None Selected"
	gamePathTextWidget := widget.NewLabel(gamePathText)

	buildIdTextPreset := "Build ID: "
	buildIdText := buildIdTextPreset + "None Selected"
	buildIdTextWidget := widget.NewLabel(buildIdText)
	var gamePath string
	var gameSelect *widget.Select

	var customGame = true

	if !global.SteamFound {
		alert.Alert("Patchman encountered an error trying to find Steam,\nYou can still patch any custom Unity game.")
	}

	gameSelect = widget.NewSelect(gameOptions, func(s string) {

		if s == "Custom" {
			customGame = true
			gamePath = "Custom Unity Game"
			gamePathText = gamePathTextPreset + gamePath
			gamePathTextWidget.SetText(gamePathText)

			buildId := "Custom Unity Game"
			buildIdText = buildIdTextPreset + buildId
			buildIdTextWidget.SetText(buildIdText)
			return
		}
		customGame = false
		if strings.HasPrefix(s, "Custom - ") {
			for x, y := range index {
				if y.AppName == s {
					gamePath = y.AppPath
					gamePathText = gamePathTextPreset + gamePath
					gamePathTextWidget.SetText(gamePathText)

					buildId := "Custom Game"
					buildIdText = buildIdTextPreset + buildId
					buildIdTextWidget.SetText(buildIdText)
					currentGame = x
					return
				}
			}
		}

		for x, y := range index {
			if y.AppName == s {
				game, found := ipc.GameAppById(y.AppID)
				if !found {
					gamePath = "Game not found.\n(Check if game is installed or is on an unplugged secondary drive)"
					gamePathText = gamePathTextPreset + gamePath
					gamePathTextWidget.SetText(gamePathText)

					buildId := "N/A"
					buildIdText = buildIdTextPreset + buildId
					buildIdTextWidget.SetText(buildIdText)

					continue
				}
				gamePath = game.FullPath
				gamePathText = gamePathTextPreset + gamePath
				gamePathTextWidget.SetText(gamePathText)

				buildId := game.BuildID
				buildIdText = buildIdTextPreset + buildId
				buildIdTextWidget.SetText(buildIdText)
				currentGame = x

			}
		}

	})

	gameSelect.PlaceHolder = "Select game to modify"

	global.MainWindow.SetContent(container.NewVBox(
		steamPathTextWidget,
		gamePathTextWidget,
		buildIdTextWidget,
		gameSelect,
		motdText,
		container.NewHBox(
			widget.NewButton("Next", func() {

				if buildIdTextWidget.Text == buildIdTextPreset+"N/A" || buildIdTextWidget.Text == buildIdTextPreset+"None Selected" {
					return
				}
				if customGame {
					buildCustomGame()
					return
				}
				global.CreateWorkingDirectories(gamePath)

				ipc.Pwd(global.Directory)
				ipc.Gwd(gamePath)

				buildPatchHandler()
			}),
			widget.NewButton("Cancel", func() { global.App.Quit() }),
		),
	))
}
