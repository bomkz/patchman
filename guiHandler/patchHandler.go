package guihandler

import (
	"os"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
)

func buildPatchHandler() {

	testWidget := widget.NewLabel("meow")
	global.MainWindow.SetContent(container.NewVBox(
		testWidget,
		container.NewHBox(
			widget.NewButton("Next", func() { os.Exit(0) }),
			widget.NewButton("Cancel", func() { os.Exit(0) }),
		),
	))
}
