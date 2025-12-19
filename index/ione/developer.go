package ione

import (
	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func buildDeveloperForm() {

	form := tview.NewForm()

	form.AddTextView("VTOL VR Version", global.VtolVersion, 0, 0, false, false).
		AddInputField("Custom mod path", "", 10, nil, pathField).
		AddButton("Install Patch", nil).
		AddButton("Quit", cancel)

	form.SetBorder(false)

	global.Root.AddAndSwitchToPage("DevForm", form, true)

}

func pathField(path string) {
	scriptPath = path
}
func scriptField(path string) {
	scriptPath = path
}
