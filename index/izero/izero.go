package izero

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
)

func buildInitialSelection() {
	VariantName := []string{}
	for _, x := range index {
		for _, y := range x.Versions {
			if y.Version == global.VtolVersion {
				VariantName = append(VariantName, x.Name)
			}
		}
	}

	if len(VariantName) == 0 {
		log.Fatal(`Couldn't find any patches matching VTOL VR version: ` + global.VtolVersion + `

If VTOL VR receives a new update and patches are yet to be marked as compatible, you could override the Build ID version by looking up a Build ID from https://github.com/bomkz/patchman-index and using it as follows:
	patchman.exe <buildid>
	patchman.exe 18407725
	
Press enter to continue...`)
		fmt.Scanln()
		os.Exit(1)
	}

	for _, x := range index {
		if x.Name == VariantName[0] {
			selection.IndexID = 0
			selection.Name = x.Name
			selection.ObjectID = x.ObjectID
		}
	}

	selection.IndexID = 0

	for _, x := range index {
		if x.Name == selection.Name {
			selection.Description = x.Descritpion
		}
	}
}

func generatePossibleCombinations(forms []IndexContentContentVersionsFormsStruct) []PossibleCombinationStruct {
	var results []PossibleCombinationStruct
	n := len(forms)
	total := 1 << n

	for mask := 0; mask < total; mask++ {
		var names []string
		var objectIDs []string

		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				names = append(names, forms[i].Name)
				objectIDs = append(objectIDs, forms[i].ObjectID)
			}
		}

		combination := PossibleCombinationStruct{
			Name:      strings.Join(names, ", "),
			ObjectIDs: objectIDs,
		}
		results = append(results, combination)
	}

	return results
}

func buildForm() {

	VariantName := []string{}
	for _, x := range index {
		for _, y := range x.Versions {
			if y.Version == global.VtolVersion {
				VariantName = append(VariantName, x.Name)
			}
		}
	}
	if len(VariantName) == 0 {
		fmt.Println("Couldn't find any patches matching VTOL VR version: " + global.VtolVersion + `
		
If VTOL VR receives a new update and patches are yet to be marked as compatible, you could override the Build ID version by looking up a Build ID from https://github.com/bomkz/patchman-index and using it as follows:
	patchman.exe <buildid>
	patchman.exe 18407725
	
Press enter to continue...`)
		fmt.Scanln()
		os.Exit(1)
	}

	descTextView := tview.NewTextView().
		SetText(selection.Description).
		SetDynamicColors(false).SetScrollable(false)

	formOps := []IndexContentContentVersionsFormsStruct{}
	for _, x := range index {
		if x.ObjectID == selection.ObjectID {
			for _, y := range x.Versions {
				if y.Version == global.VtolVersion {
					formOps = y.Forms
				}
			}
		}
	}

	combined := generatePossibleCombinations(formOps)
	var combinedNames = []string{}
	for _, x := range combined {
		combinedNames = append(combinedNames, x.Name)
	}
	dynamicDropbox := tview.NewDropDown().SetOptions(combinedNames, func(text string, index int) {
		for _, x := range selection.Combinations {
			if x.Name == text {
				selection.Combination = x.ObjectIDs
			}
		}
	})
	selection.Combinations = combined
	form := tview.NewForm()
	form.AddTextView("VTOL VR Version", global.VtolVersion, 0, 0, false, false).
		AddTextView("Current Variant ", Status.InstalledName+" Object ID "+Status.InstalledObjectId+" Variant ID "+Status.InstalledVariantId+" Version ID "+Status.InstalledVersionId, 0, 0, false, false).
		AddDropDown("Select variant", VariantName, selection.IndexID, func(option string, optionIndex int) {
			for _, x := range index {
				if x.Name == option {
					selection.Name = x.Name
				}
			}

			for _, x := range index {
				if x.Name == option {
					selection.ObjectID = x.ObjectID
					selection.IndexID = optionIndex
				}
			}

			for _, x := range index {
				if x.Name == selection.Name {
					selection.Description = x.Descritpion
				}
			}

			formOps := []IndexContentContentVersionsFormsStruct{}
			for _, x := range index {
				if x.ObjectID == selection.ObjectID {
					for _, y := range x.Versions {
						if y.Version == global.VtolVersion {
							formOps = y.Forms
						}
					}
				}
			}

			combined := generatePossibleCombinations(formOps)
			selection.Combinations = combined
			selection.Combination = nil
			var combinedNames = []string{}
			for _, x := range combined {
				combinedNames = append(combinedNames, x.Name)
			}

			dynamicDropbox.SetOptions(combinedNames, func(text string, index int) {
				for _, x := range selection.Combinations {
					if x.Name == text {
						selection.Combination = x.ObjectIDs
					}
				}
			}).SetCurrentOption(0)

			descTextView.SetText(selection.Description)
		}).
		AddFormItem(descTextView).
		AddFormItem(dynamicDropbox).
		AddButton("Patch", install).
		AddButton("Unpatch", uninstall).
		AddButton("Quit", cancel)

	form.SetBorder(false)
	global.Root.AddAndSwitchToPage("installform", form, true)
}

func HandleForm(indexbyte []byte) error {

	readTaintNoFail()

	if err := json.Unmarshal(indexbyte, &index); err != nil {
		return err
	}

	buildInitialSelection()
	buildForm()
	return nil
}

func BuildTaintInfo() string {
	ReadTaint()
	return "Current Variant " + Status.InstalledName + " Object ID " + Status.InstalledObjectId + " Variant ID " + Status.InstalledVariantId + " Version ID " + Status.InstalledVersionId
}

func ReadTaint() {
	vtolvrpath := global.FindVtolPath()

	taint, err := os.ReadFile(vtolvrpath + "\\patchman.json")
	if err != nil {
		global.StopApp <- true
		log.Fatal("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman.")
	}
	err = json.Unmarshal(taint, &Status)
	if err != nil {
		log.Fatal("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman.")
	}
}

func readTaintNoFail() {
	vtolvrpath := global.FindVtolPath()

	taint, err := os.ReadFile(vtolvrpath + "\\patchman.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(taint, &Status)
	if err != nil {
		log.Fatal("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman.")
	}
}
