package guiHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
)

func buildCustomGame() {
	gameNameWidget := widget.NewEntry()
	gameNameWidget.SetPlaceHolder("Game Name")

	gameLocationEntryWidget := widget.NewEntry()
	gameLocationFolderPickerDialog := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil {
			global.FatalError(err)
		}

		if lu != nil {
			filepath, found := strings.CutPrefix(lu.String(), "file://")
			if found {
				filepath = strings.ReplaceAll(filepath, "/", PathSeparator())
				gameLocationEntryWidget.SetText(filepath)
			} else {
				filepath = strings.ReplaceAll(filepath, "/", PathSeparator())
				gameLocationEntryWidget.SetText(lu.String())
			}
		}
	}, global.MainWindow)
	gameLocationEntryWidget.SetPlaceHolder("Game folder location")

	var contents []Content

	contentListWidget := widget.NewSelect([]string{}, func(s string) {})

	contentNameEntryWidget := widget.NewEntry()
	contentNameEntryWidget.SetPlaceHolder("Content Name")
	contentNameEntryWidget.Resize(fyne.NewSize(100, 20))

	contentLocationEntryWidget := widget.NewEntry()
	contentLocationEntryWidget.SetPlaceHolder("Content File Path")
	contentLocationEntryWidget.Resize(fyne.NewSize(100, 20))
	contentLocationFilePickerDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			global.FatalError(err)
		}
		if reader != nil {
			filepath, found := strings.CutPrefix(reader.URI().String(), "file://")
			if found {
				filepath = strings.ReplaceAll(filepath, "/", PathSeparator())

				contentLocationEntryWidget.SetText(filepath)
			} else {
				filepath = strings.ReplaceAll(filepath, "/", PathSeparator())

				contentLocationEntryWidget.SetText(reader.URI().String())
			}
		}
	}, global.MainWindow)

	global.MainWindow.SetContent(container.NewVBox(
		gameNameWidget,
		gameLocationEntryWidget,
		widget.NewButton("Folder Picker", func() {
			gameLocationFolderPickerDialog.Show()
		}),
		contentNameEntryWidget,
		contentLocationEntryWidget,
		widget.NewButton("File Picker", func() { contentLocationFilePickerDialog.Show() }),
		container.NewHBox(
			widget.NewButton("Add", func() {

				if contentLocationEntryWidget.Text == "" {
					return
				}
				if contentNameEntryWidget.Text == "" {
					return
				}
				newContent := Content{
					Name: contentNameEntryWidget.Text,
					Path: contentLocationEntryWidget.Text,
				}
				contents = append(contents, newContent)
				contentsString := []string{}

				for _, x := range contents {
					contentsString = append(contentsString, x.Name)
				}
				contentListWidget.SetOptions(contentsString)
			}),
			widget.NewButton("Remove", func() {
				tmpContents := []Content{}
				for _, x := range contents {

					if x.Name == contentListWidget.Selected {
						continue
					}
					tmpContents = append(tmpContents, x)
					contents = tmpContents

					contentsString := []string{}
					for _, x := range contents {
						contentsString = append(contentsString, x.Name)
					}
					contentListWidget.SetOptions(contentsString)
					contentListWidget.ClearSelected()
				}

			}),
			contentListWidget,
		),
		container.NewHBox(
			widget.NewButton("Save Custom Game", func() {
				if len(contents) == 0 {
					return
				}
				if gameNameWidget.Text == "" {
					return
				}
				if gameLocationEntryWidget.Text == "" {
					return
				}

				type customGame struct {
					Name     string    `json:"name"`
					Location string    `json:"location"`
					Content  []Content `json:"content"`
				}

				newGame := CustomGame{
					Name:     gameNameWidget.Text,
					Location: gameLocationEntryWidget.Text,
					Content:  contents,
				}
				saveGame(newGame)

			}),
			widget.NewButton("Import Custom Game", func() {
				importPicker := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
					if err != nil {
						global.FatalError(err)
					}

					if reader != nil {
						tmpGame := CustomGame{}
						gameByte := global.Assure(io.ReadAll(reader))
						global.AssureNoReturn(json.Unmarshal(gameByte, &tmpGame))
						tmpGame.Location = strings.ReplaceAll(tmpGame.Location, "\\", PathSeparator())
						tmpGame.Location = strings.ReplaceAll(tmpGame.Location, "/", PathSeparator())
						for x := range tmpGame.Content {
							tmpGame.Content[x].Path = strings.ReplaceAll(tmpGame.Content[x].Path, "\\", PathSeparator())
							tmpGame.Content[x].Path = strings.ReplaceAll(tmpGame.Content[x].Path, "/", PathSeparator())

						}

						os.WriteFile(global.Assure(os.UserHomeDir())+tmpGame.Name, gameByte, 0755)
					}
				}, global.MainWindow)
				importPicker.Show()
			}),
			widget.NewButton("Cancel", func() { global.App.Quit() }),
		),
	))
}

type Content struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type CustomGame struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Content  []Content `json:"content"`
}

func saveGame(game CustomGame) {
	checkDirExists()
	game.Name = checkFileExists(game.Name)
	if !strings.HasSuffix(game.Location, PathSeparator()) {
		game.Location += PathSeparator()
	}

	for x := range game.Content {
		game.Content[x].Path, _ = strings.CutPrefix(game.Content[x].Path, game.Location)
	}
	gameFile := global.Assure(os.Create(global.Assure(os.UserHomeDir()) + PathSeparator() + "patchman" + PathSeparator() + game.Name))
	gameByte := global.Assure(json.Marshal(game))

	global.Assure(gameFile.Write(gameByte))
	gameFile.Close()

	savedWidget := widget.NewLabel("Custom Game Saved")

	global.MainWindow.SetContent(container.NewVBox(savedWidget,
		widget.NewButton("Finish", func() {
			os.Exit(0)
		})))
}

func checkFileExists(fileName string) string {
	_, err := os.Stat(global.Assure(os.UserHomeDir()) + PathSeparator() + "patchman" + PathSeparator() + fileName)
	if errors.Is(err, os.ErrNotExist) {
		return fileName
	}
	return fmt.Sprint(time.Now().Unix()) + "-" + fileName

}

func checkDirExists() {

	_, err := os.Stat(global.Assure(os.UserHomeDir()) + PathSeparator() + "patchman")
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(global.Assure(os.UserHomeDir())+PathSeparator()+"patchman", 0755)
	}

}

func PathSeparator() string {
	if global.OsName == "windows" {
		return "\\"
	} else {
		return "/"
	}
}
