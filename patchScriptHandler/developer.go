package patchScriptHandler

import (
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller/patchScriptOne"
	"github.com/rivo/tview"
)

var Compression = []string{
	"None",
	"LZMA",
	"LZ4",
	"LZ4Fast",
}

func buildDeveloperForm() {

	global.Root.RemovePage("installform")

	form := tview.NewForm()

	form.AddTextView(global.TargetName+" BuildID", global.TargetBuildID, 0, 0, false, false).
		AddInputField("Custom mod path", "", 40, nil, pathField).
		AddDropDown("Compression", Compression, 0, setCompression).
		AddButton("Install Patch", installFunc).
		AddButton("Quit", global.ExitApp)

	form.SetBorder(false)
	global.Root.AddAndSwitchToPage("DevForm", form, true)
}

func setCompression(option string, optionIndex int) {
	patchScriptOne.CompressionType = option
}
func installFunc() {
	install(modPath)
	global.Root.RemovePage("DevForm")
}

func pathField(path string) {
	modPath = path
}
