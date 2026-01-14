package installer

import (
	"encoding/json"

	"github.com/bomkz/patchman/global"
)

func HandleActions(actionData []byte) {
	var actionScript []ActionScriptStruct

	global.AssureNoReturn(json.Unmarshal(actionData, &actionScript))

	global.ExitTview()

	for _, x := range actionScript {
		switch x.Action {
		case "importbundle":
			batchBundleImport(x.ActionData)
		case "importasset":
			batchAssetImport(x.ActionData)
		case "copy":
			handleCopy(x.ActionData)
		}
	}

	global.ExitAppWithMessage("Done!")
}

func handleCopy(actionData []byte) {
	var copyData CopyStruct

	global.AssureNoReturn(json.Unmarshal(actionData, &copyData))
	copyData.Destination = global.TargetPath + copyData.Destination
	global.CopyFromProgramWorkingDirectory(copyData.FileName, copyData.Destination)

}

func batchBundleImport(patchmanJson []byte) {
	var patchmanData PatchmanUnityStruct

	global.AssureNoReturn(json.Unmarshal(patchmanJson, &patchmanData))
	if !global.ExistsAtGwd(patchmanData.OriginalFilePath) {
		return
	}
	if len(Content) >= 1 && Content[1].ContentName != "none" && Content[1].ContentPath != "none" {
		for _, x := range Content {
			if patchmanData.OriginalFilePath == x.ContentPath && !x.Modify {
				return
			}
		}
	}
	gwd := global.GetGwd()

	renameFile := patchmanData.OriginalFilePath
	patchmanData.OriginalFilePath = gwd + patchmanData.OriginalFilePath

	patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"

	if len(Assets) >= 1 && Assets[1].AssetName != "none" {
		var tmpOperations []PatchmanUnityOperationsStruct

		for _, x := range Assets {
			for _, y := range patchmanData.Operations {
				if y.AssetName == x.AssetName && x.Modify {
					tmpOperations = append(tmpOperations, y)
				}
			}
		}

		if len(tmpOperations) == 0 {
			return
		}
		patchmanData.Operations = tmpOperations
	}

	createOperationsFile(patchmanData)

	runPatchmanUnityBundles()

	global.RenameGameWorkingDirectoryFile(renameFile)

}

func batchAssetImport(patchmanJson []byte) {
	var patchmanData PatchmanUnityStruct

	global.AssureNoReturn(json.Unmarshal(patchmanJson, &patchmanData))

	if len(Content) >= 1 && Content[1].ContentName != "none" && Content[1].ContentPath != "none" {
		for _, x := range Content {
			if patchmanData.OriginalFilePath == x.ContentPath && !x.Modify {
				return
			}
		}
	}

	gwd := global.GetGwd()

	renameFile := patchmanData.OriginalFilePath

	patchmanData.OriginalFilePath = gwd + patchmanData.OriginalFilePath

	patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"

	var tmpOperations []PatchmanUnityOperationsStruct

	if len(Assets) >= 1 && Assets[1].AssetName != "none" {
		var tmpOperations []PatchmanUnityOperationsStruct

		for _, x := range Assets {
			for _, y := range patchmanData.Operations {
				if y.AssetName == x.AssetName && x.Modify {
					tmpOperations = append(tmpOperations, y)
				}
			}
		}

		if len(tmpOperations) == 0 {
			return
		}
		patchmanData.Operations = tmpOperations
	}

	patchmanData.Operations = tmpOperations

	createOperationsFile(patchmanData)

	runPatchmanUnityAssets()

	global.RenameGameWorkingDirectoryFile(renameFile)

}
