package patchScriptHandler

import (
	"encoding/json"

	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func HandleForm(indexbyte []byte, motd string) {

	// Unmarshals index to global variable
	global.AssureNoReturn(json.Unmarshal(indexbyte, &index))

	buildgameForm(motd)
}

// Builds game choice form
func buildgameForm(motd string) {

	// Builds game list string
	buildGameList()

	gameDropBox := tview.NewDropDown().
		SetOptions(games, selectedGame).SetLabel("Select your game.")

	form := tview.NewForm()
	form.
		AddTextView("Select your game", "", 40, 1, false, false).
		AddTextView("MOTD", motd, 40, 1, false, false).
		AddFormItem(gameDropBox).AddButton("Next", gameNext)

	form.SetBorder(false)
	global.Root.AddAndSwitchToPage("gameform", form, true)
}

func beginInstall() {
	for _, x := range patchData {
		if x.PatchName == patches[currentPatch] {
			for _, y := range x.PatchVariants {
				if y.Variant == variants[currentVariant] {
					global.DownloadFileToProgramWorkingDirectory("patchman.zip", y.DownloadLink)
				}
			}
		}
	}

	install("patchman.zip")
	global.Root.RemovePage("installform")
}

func setPatchList() {
	patches = []string{}
	for _, patch := range patchData {
		patches = append(patches, patch.PatchName)
	}
}

func setVariantList() {
	variants = []string{}
	for _, variant := range patchData[currentPatch].PatchVariants {
		variants = append(variants, variant.Variant)
	}

}

func selectedVariant(option string, optionIndex int) {
	currentVariant = optionIndex
}
