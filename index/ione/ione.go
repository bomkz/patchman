package ione

import (
	"encoding/json"

	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func HandleForm(indexbyte []byte, motds string) error {
	patchData = handleIndexJson(indexbyte)
	for _, patch := range patchData {
		patches = append(patches, patch.PatchName)
	}
	motd = motds
	buildForm()
	return nil
}

var motd string
var patchData []IndexContentStruct

func handleIndexJson(indexbyte []byte) []IndexContentStruct {
	tmp := []IndexContentStruct{}
	err := json.Unmarshal(indexbyte, &tmp)
	if err != nil {
		global.FatalError(err)
	}

	return tmp
}

var patches = []string{}
var variants = []string{}
var currentVariant int = 0
var currentSelection int = 0

func buildForm() {

	setVariantList()

	descTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchDesc).
		SetDynamicColors(false).SetScrollable(false)

	authTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchAuthor).
		SetDynamicColors(false).SetScrollable(false)

	linkTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchLink).
		SetDynamicColors(false).SetScrollable(false)

	varDropBox := tview.NewDropDown().
		SetOptions(variants, selectedVariant)

	form := tview.NewForm()
	form.AddTextView("VTOL VR Version", global.VtolVersion, 0, 0, false, false).
		AddTextView("MOTD", motd, 0, 0, false, false).
		AddDropDown("Select Patch...", patches, 0, func(option string, optionIndex int) {
			currentSelection = optionIndex
			authTextView.SetText(patchData[currentSelection].PatchAuthor)
			descTextView.SetText(patchData[currentSelection].PatchDesc)
			linkTextView.SetText(patchData[currentSelection].PatchLink)
			setVariantList()
			varDropBox.SetOptions(variants, selectedVariant)
		}).
		AddFormItem(varDropBox).
		AddFormItem(authTextView).
		AddFormItem(descTextView).
		AddFormItem(linkTextView).
		AddDropDown("Select Variant...", variants, currentVariant, selectedVariant).
		AddButton("Patch", nil).
		AddButton("Unpatch", nil).
		AddButton("Custom", buildDeveloperForm).
		AddButton("Quit", global.ExitApp)

	form.SetBorder(false)
	global.Root.RemovePage("installform")
	global.Root.AddAndSwitchToPage("installform", form, true)

}

func setVariantList() {
	variants = []string{}
	for _, variant := range patchData[currentSelection].PatchVariants {
		variants = append(variants, variant.Variant)
	}

}

func selectedVariant(option string, optionIndex int) {
	currentVariant = optionIndex
}
