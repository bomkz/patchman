package ione

import (
	"log"

	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func HandleForm(indexbyte []byte, motd string) error {
	buildForm()
	return nil
}

func buildForm() {

	form := tview.NewForm()
	form.AddTextView("VTOL VR Version", global.VtolVersion, 0, 0, false, false).
		AddTextView("MOTD", "", 0, 0, false, false).
		AddButton("Patch", nil).
		AddButton("Unpatch", nil).
		AddButton("Custom", buildDeveloperForm).
		AddButton("Quit", cancel)

	form.SetBorder(false)
	global.Root.AddAndSwitchToPage("installform", form, true)

}

func cancel() {
	log.Fatal("User Quit Installer")

}
