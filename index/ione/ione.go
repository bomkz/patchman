package ione

import (
	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func HandleForm(indexbyte []byte, motd string) (*tview.Form, error) {
	return buildForm(), nil
}

func buildForm() *tview.Form {

	form := tview.NewForm()
	form.AddTextView("VTOL VR Version", global.VtolVersion, 0, 0, false, false).
		AddTextView("MOTD", "", 0, 0, false, false).
		AddButton("Patch", nil).
		AddButton("Unpatch", nil).
		AddButton("DevTools", nil).
		AddButton("Quit", nil)

	form.SetBorder(false)
	return form
}
