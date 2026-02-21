package installer

import (
	"encoding/json"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc"
)

func HandleActions(actionData []byte) {
	var actionScript []ActionScriptStruct

	if CompressionType == "" {
		CompressionType = "none"
	}

	global.AssureNoReturn(json.Unmarshal(actionData, &actionScript))

	for _, x := range actionScript {
		switch x.Action {
		case "importbundle":
			batchBundleImport(x.ActionData)
		case "importasset":
			batchAssetImport(x.ActionData)
		case "copy":
			ipc.Copy(x.ActionData)
		}
	}

}

func batchBundleImport(patchmanJson []byte) {
	var patchmanData PatchmanUnityStruct

	global.AssureNoReturn(json.Unmarshal(patchmanJson, &patchmanData))
	if !ipc.ExistsAtGwd(patchmanData.OriginalFilePath) {
		return
	}
	if len(Content) >= 1 && Content[0].ContentName != "none" && Content[0].ContentPath != "none" {
		for _, x := range Content {
			if patchmanData.OriginalFilePath == x.ContentPath && !x.Modify {
				return
			}
		}
	}
	gwd := global.GetGwd()

	renameFile := patchmanData.OriginalFilePath
	patchmanData.OriginalFilePath = gwd + "\\" + patchmanData.OriginalFilePath

	patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"

	if len(Assets) >= 1 && Assets[0].AssetName != "none" {
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

	ipc.PatchBundles(CompressionType)

	ipc.RenameGameWorkingDirectoryFile(renameFile)

}

func batchAssetImport(patchmanJson []byte) {
	var patchmanData PatchmanUnityStruct

	global.AssureNoReturn(json.Unmarshal(patchmanJson, &patchmanData))

	if len(Content) >= 1 && Content[0].ContentName != "none" && Content[0].ContentPath != "none" {
		for _, x := range Content {
			if patchmanData.OriginalFilePath == x.ContentPath && !x.Modify {
				return
			}
		}
	}

	gwd := global.GetGwd()

	renameFile := patchmanData.OriginalFilePath

	patchmanData.OriginalFilePath = gwd + "\\" + patchmanData.OriginalFilePath

	patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"

	if len(Assets) >= 1 && Assets[0].AssetName != "none" {
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

	ipc.PatchAssets()

	ipc.RenameGameWorkingDirectoryFile(renameFile)

}
