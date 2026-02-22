package guiHandler

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc"
)

func InitGui() {
	global.SteamPath, global.SteamFound = ipc.SteamPath()
	global.MainWindow.SetContent(widget.NewLabel(""))
	go func() {
		validateCustomGames()
		buildGameListSteam()
	}()
	global.MainWindow.ShowAndRun()

}

func validateCustomGames() {
	_, err := os.Stat(customDir)

	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return

		}
		global.FatalError(err)

	}
	customGameDir := global.Assure(os.ReadDir(customDir))
	for _, x := range customGameDir {
		var fixed bool
		gameByte := global.Assure(os.ReadFile(customDir + x.Name()))
		tmpGame := CustomGame{}
		err := json.Unmarshal(gameByte, &tmpGame)
		if err != nil {
			continue
		}
		_, err = os.Stat(tmpGame.Location)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				newPath := fixPath
				tmpGame.Location = newPath(tmpGame.Location, "folder")
				if !strings.HasSuffix(tmpGame.Location, PathSeparator()) {
					tmpGame.Location += PathSeparator()
				}
				fixed = true
			}
		}

		for y, z := range tmpGame.Content {
			_, err = os.Stat(tmpGame.Location + PathSeparator() + z.Path)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					newPath := fixPath(z.Path, "file")
					tmpGame.Content[y].Path = newPath
					tmpGame.Content[y].Path, _ = strings.CutPrefix(tmpGame.Content[y].Path, tmpGame.Location)
					fixed = true
				}
			}
		}
		if fixed {
			newGameByte := global.Assure(json.Marshal(tmpGame))
			global.AssureNoReturn(os.Remove(customDir + x.Name()))
			global.AssureNoReturn(os.WriteFile(customDir+x.Name(), newGameByte, 0755))
		}

	}
}

func fixPath(resource string, resourceType string) (newPath string) {
	var newPathChan = make(chan string)
	go func() {
		global.MainWindow.Resize(fyne.NewSize(500, 500))

		errorLabel := widget.NewLabel("There was an error reading your custom game files...\n" + resource + "\nappears to not exist anymore or was moved.\nIf it was moved, input the new location below.\n")

		newLocation := widget.NewEntry()
		newLocation.SetPlaceHolder("Insert new location for missing resource: " + resource)
		if resourceType == "file" {
			newLocationFilePickerDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					global.FatalError(err)
				}

				if reader != nil {
					filePath, found := strings.CutPrefix(reader.URI().String(), "file://")
					if !found {
						filePath = reader.URI().String()
					}

					filePath = strings.ReplaceAll(filePath, "/", PathSeparator())
					filePath = strings.ReplaceAll(filePath, "\\", PathSeparator())

					newLocation.SetText(filePath)
				}
			}, global.MainWindow)

			global.MainWindow.SetContent(container.NewVBox(
				errorLabel,
				newLocation,
				widget.NewButton("Open File Picker", func() {
					newLocationFilePickerDialog.Show()
				}),
				widget.NewButton("Save", func() {
					_, err := os.Stat(newLocation.Text)
					if err != nil {
						if errors.Is(err, os.ErrNotExist) {
							return
						}
					}

					newPathChan <- newLocation.Text

				}),
			))
		}
		if resourceType == "folder" {
			newLocationFolderPickerDialog := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
				if err != nil {
					global.FatalError(err)
				}

				if lu != nil {
					filePath, found := strings.CutPrefix(lu.String(), "file://")
					if !found {
						filePath = lu.String()
					}

					filePath = strings.ReplaceAll(filePath, "/", PathSeparator())
					filePath = strings.ReplaceAll(filePath, "\\", PathSeparator())

					newLocation.SetText(filePath)
				}
			}, global.MainWindow)

			global.MainWindow.SetCloseIntercept(func() {
				os.Exit(0)
			})

			global.MainWindow.SetContent(container.NewVBox(
				errorLabel,
				newLocation,
				widget.NewButton("Open Folder Picker", func() {
					newLocationFolderPickerDialog.Show()
				}),
				widget.NewButton("Save", func() {

					_, err := os.Stat(newLocation.Text)
					if err != nil {
						if errors.Is(err, os.ErrNotExist) {
							return
						}
					}
					newPathChan <- newLocation.Text

				}),
			))
		}
	}()
	newPath = <-newPathChan
	return
}

var customDir = global.Assure(os.UserHomeDir()) + PathSeparator() + "patchman" + PathSeparator()
