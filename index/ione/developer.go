package ione

import (
	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func buildDeveloperForm() {

	form := tview.NewForm()

	form.AddTextView("VTOL VR Version", global.VtolVersion, 0, 0, false, false).
		AddInputField("Path", "", 10, nil, pathField).
		AddButton("Build Patch", nil).
		AddButton("Validate Patch", nil).
		AddButton("Quit", cancel)

	form.SetBorder(false)

}

func pathField(path string) {
}
