package guihandler

import (
	"os"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bomkz/patchman/global"
)

func buildPatchHandler() {

	for _, x := range index[currentGame].Patches {
		patches = append(patches, x.PatchName)
	}

	patchDescPreset := "Patch Description: "
	patchDescWidget := widget.NewLabel(patchDescPreset + "None Selected")

	patchAuthPreset := "Patch Author: "
	patchAuthWidget := widget.NewLabel(patchAuthPreset + "None Selected")

	patchLinkPreset := "Patch Link: "
	patchLinkWidget := widget.NewLabel(patchLinkPreset + "None Selected")

	variants = append(variants, "No Patch Selected")

	variantSelect := widget.NewSelect(variants, func(s string) {
		for x, y := range index[currentGame].Patches[currentPatch].PatchVariants {
			if y.Variant == s {
				currentVariant = x
			}
		}
	})

	patchSelect := widget.NewSelect(patches, func(s string) {
		for x, y := range index[currentGame].Patches {
			if y.PatchName == s {
				currentPatch = x
				patchDescWidget.SetText(patchDescPreset + y.PatchDesc)
				patchAuthWidget.SetText(patchAuthPreset + y.PatchAuthor)
				patchLinkWidget.SetText(patchLinkPreset + y.PatchLink)
				variants = []string{}
				for _, z := range y.PatchVariants {
					variants = append(variants, z.Variant)
				}
				variantSelect.SetOptions(variants)
			}
		}
	})

	global.MainWindow.SetContent(container.NewVBox(
		patchSelect,
		variantSelect,
		patchDescWidget,
		patchAuthWidget,
		patchLinkWidget,
		container.NewHBox(
			widget.NewButton("Next", func() { handlePatch() }),
			widget.NewButton("Custom", func() { os.Exit(0) }),
			widget.NewButton("Cancel", func() { global.ExitSuccess() }),
		),
	))
}

func handlePatch() {

	global.DownloadFileToProgramWorkingDirectory("./patchman.zip", index[currentGame].Patches[currentPatch].PatchVariants[currentVariant].DownloadLink)
	global.UnzipIntoProgramWorkingDirectory("./patchman.zip")
	buildContentHandler()

}
