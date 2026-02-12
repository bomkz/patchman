package guihandler

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
)

func handleFinish() {
	finishText := widget.NewLabel("Patchman has finished installing your patch successfully.")
	global.MainWindow.SetContent(container.NewVBox(
		finishText,
		container.NewHBox(
			widget.NewButton("Finish", func() { global.ExitSuccess() }),
		),
	))

}
