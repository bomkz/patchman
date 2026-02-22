package guiHandler

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
)

func handleFinish() {
	finishText := widget.NewLabel("Patchman has finished installing\nyour patch successfully.")
	global.MainWindow.SetContent(container.NewVBox(
		finishText,
		container.NewHBox(
			widget.NewButton("Finish", func() { global.ExitSuccess() }),
		),
	))
	global.MainWindow.Resize(fyne.NewSize(100, 50))

}
