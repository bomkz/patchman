package guiHandler

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
)

func buildCustomPatchHandler() {
	var openedFile fyne.URIReadCloser

	fileDir := ""
	info := widget.NewLabel("Select your custom patch file:")
	dir := widget.NewLabel("None Selected.")
	blank := widget.NewLabel("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                                                                                                                                                      ")
	filePicker := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			global.FatalError(err)
		}
		if reader == nil {
			return
		}
		openedFile = reader
		dir.SetText(openedFile.URI().Path())
		fileDir = openedFile.URI().Path()
		openedFile.Close()
	}, global.MainWindow)

	filePicker.SetFilter(storage.NewExtensionFileFilter([]string{".zip"}))

	global.MainWindow.SetContent(container.NewVBox(
		info,
		dir,
		widget.NewButton("Open Picker", func() { filePicker.Show() }),
		blank,
		container.NewHBox(
			widget.NewButton("Next", func() {
				writeToPwd(fileDir)
				global.UnzipIntoProgramWorkingDirectory("./patchman.zip")
				buildContentHandler()
			}),
			widget.NewButton("Cancel", func() { global.ExitSuccess() }),
		),
	))
}

func writeToPwd(openedFile string) {
	// Open file in read-write mode
	file := global.Assure(os.OpenFile(global.Directory+"\\patchman.zip", os.O_RDWR|os.O_CREATE, 0644))

	defer file.Close()
	custom := global.Assure(os.ReadFile(openedFile))

	// Write data to the file
	global.Assure(file.Write(custom))

}
