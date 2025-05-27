package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rivo/tview"
)

func parseIndex() error {
	preindexbyte, err := loadPreIndex()
	if err != nil {
		return err
	}
	err = json.Unmarshal(preindexbyte, &preindex)
	if err != nil {
		return err
	}
	for _, x := range preindex.Content {
		preindexversion, err := strconv.Atoi(x.Version)
		if err != nil {
			return err
		}
		if MaxPreIndexVersion == preindexversion {
			err = json.Unmarshal(x.Content, &index)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func loadPreIndex() ([]byte, error) {
	if !internet {
		var err error
		indexmem, err = os.ReadFile("C:\\patchman\\index.json")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return nil, err
		}
	}

	data := indexmem

	return data, nil
}

func buildInitialSelection() {
	VariantName := []string{}
	for _, x := range index {
		for _, y := range x.Versions {
			if y.Version == vtolversion {
				VariantName = append(VariantName, x.Name)
			}
		}
	}

	if len(VariantName) == 0 {
		log.Fatal(`Couldn't find any patches matching VTOL VR version: ` + vtolversion + `

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
			if y.Version == vtolversion {
				VariantName = append(VariantName, x.Name)
			}
		}
	}
	if len(VariantName) == 0 {
		fmt.Println("Couldn't find any patches matching VTOL VR version: " + vtolversion + `
		
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
				if y.Version == vtolversion {
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
	version := vtolversion
	form := tview.NewForm()
	form.AddTextView("VTOL VR Version", version, 0, 0, false, false).
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
						if y.Version == vtolversion {
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
		AddButton("Unpatch", Uninstall).
		AddButton("Quit", cancel)

	form.SetBorder(false)
	root.AddAndSwitchToPage("installform", form, true)
}
