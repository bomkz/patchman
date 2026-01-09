package patchScriptHandler

import (
	"encoding/json"

	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func HandleForm(indexbyte []byte, motds string) error {
	patchData, err := handleIndexJson(indexbyte)
	if err != nil {
		return err
	}
	for _, patch := range patchData {
		patches = append(patches, patch.PatchName)
	}
	motd = motds
	buildForm()
	return nil
}

var motd string
var patchData []IndexContentStruct

func handleIndexJson(indexbyte []byte) ([]IndexContentStruct, error) {
	tmp := []IndexContentStruct{}
	err := json.Unmarshal(indexbyte, &tmp)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}

var patches = []string{}
var variants = []string{}
var currentVariant int = 0
var currentSelection int = 0

func buildForm() {

	setVariantList()

	descTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchDesc).
		SetDynamicColors(false).SetScrollable(true).SetMaxLines(4)

	authTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchAuthor).
		SetDynamicColors(false).SetScrollable(false).SetMaxLines(2)

	linkTextView := tview.NewTextView().
		SetText(patchData[currentSelection].PatchLink).
		SetDynamicColors(false).SetScrollable(false).SetMaxLines(2)

	varDropBox := tview.NewDropDown().
		SetOptions(variants, selectedVariant)

	form := tview.NewForm()
	form.AddTextView("VTOL VR Version", global.TargetVersion, 0, 0, false, false).
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
		AddDropDown("Select Compression Type...", Compression, 0, setCompression).
		AddDropDown("Select Variant...", variants, currentVariant, selectedVariant).
		AddButton("Patch", nil).
		AddButton("Unpatch", uninstall).
		AddButton("Custom", buildDeveloperForm).
		AddButton("Quit", global.ExitApp)

	form.SetBorder(false)
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
