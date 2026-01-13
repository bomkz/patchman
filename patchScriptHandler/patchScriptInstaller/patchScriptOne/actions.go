package patchScriptOne

import (
	"encoding/json"

	"github.com/bomkz/patchman/global"
)

func HandleActions(actionData []byte) {
	var actionScript []ActionScriptStruct

	global.AssureNoReturn(json.Unmarshal(actionData, &actionScript))

	for _, x := range actionScript {
		switch x.Action {
		case "importbundle":
			var tmpData PatchmanUnityStruct
			global.AssureNoReturn(json.Unmarshal(x.ActionData, &tmpData))
		case "importasset":
			var tmpData PatchmanUnityStruct
			global.AssureNoReturn(json.Unmarshal(x.ActionData, &tmpData))

		case "copy":
			var tmpCopy CopyStruct
			global.AssureNoReturn(json.Unmarshal(x.ActionData, &tmpCopy))
		}
	}

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

	gwd := global.GetGwd()

	renameFile := patchmanData.OriginalFilePath
	patchmanData.OriginalFilePath = gwd + patchmanData.OriginalFilePath

	patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"

	createOperationsFile(patchmanData)

	runPatchmanUnityBundles()

	global.RenameGameWorkingDirectoryFile(renameFile)

}

func batchAssetImport(patchmanJson []byte) {
	var patchmanData PatchmanUnityStruct

	global.AssureNoReturn(json.Unmarshal(patchmanJson, &patchmanData))

	gwd := global.GetGwd()

	renameFile := patchmanData.OriginalFilePath

	patchmanData.OriginalFilePath = gwd + patchmanData.OriginalFilePath

	patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"

	createOperationsFile(patchmanData)

	runPatchmanUnityAssets()

	global.RenameGameWorkingDirectoryFile(renameFile)

}
