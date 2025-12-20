package ione

import (
	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func buildDeveloperForm() {

	form := tview.NewForm()

	form.AddTextView("VTOL VR Version", global.VtolVersion, 0, 0, false, false).
		AddInputField("Custom mod path", "", 40, nil, pathField).
		AddButton("Install Patch", installFunc).
		AddButton("Quit", global.ExitApp)

	form.SetBorder(false)

	global.Root.AddAndSwitchToPage("DevForm", form, true)

}

func installFunc() {
	install(modPath)
}

func pathField(path string) {
	modPath = path
}
