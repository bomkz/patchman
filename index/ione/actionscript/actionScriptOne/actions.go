package actionScriptOne

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/global"
)

func Uninstall() {
	taintInfo := readTaint()
	for _, x := range taintInfo.ModifiedFiles {
		revertPatch(x)
	}
	vtolvrpath := global.FindVtolPath()

	err := os.Remove(vtolvrpath + "\\patchman.json")
	if err != nil {
		global.FatalError(err)

	}
}

func readTaint() TaintInfoStruct {
	var taintInfo TaintInfoStruct

	vtolvrpath := global.FindVtolPath()
	taintfile, err := os.ReadFile(vtolvrpath + ".\\patchman.json")
	if err != nil {
		global.FatalError(err)

	}

	err = json.Unmarshal(taintfile, &taintInfo)
	if err != nil {
		global.FatalError(err)

	}
	return taintInfo

}

func revertPatch(filePath string) {

	err := os.Remove(filePath)
	if err != nil {
		global.FatalError(err)

	}

	err = os.Rename(filePath+".orig", filePath)
	if err != nil {
		global.FatalError(err)

	}

}

func HandleActions(actionData []byte) {
	var actionScript []ActionScriptStruct

	err := json.Unmarshal(actionData, &actionScript)
	if err != nil {
		global.FatalError(err)

	}

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

	buildTaintInfo()

}

func handleCopy(actionData []byte) {
	var copyData CopyStruct

	err := json.Unmarshal(actionData, &copyData)
	if err != nil {
		global.FatalError(err)

	}

	vtolvrpath := global.FindVtolPath()
	copyData.Destination = vtolvrpath + copyData.Destination

	err = os.Rename(global.Directory+"\\"+copyData.FileName, copyData.Destination)
	if err != nil {
		global.FatalError(err)

	}
}

func batchBundleImport(patchmanJson []byte) {
	var patchmanData PatchmanUnityStruct

	err := json.Unmarshal(patchmanJson, &patchmanData)
	if err != nil {
		global.FatalError(err)

	}

	vtolvrpath := global.FindVtolPath()

	patchmanData.OriginalFilePath = vtolvrpath + "\\" + patchmanData.OriginalFilePath
	patchmanData.ModifiedFilePath = vtolvrpath + patchmanData.OriginalFilePath + ".mod"
	renameQueue = append(renameQueue, patchmanData.OriginalFilePath)
	taintQueue = append(taintQueue, patchmanData.OriginalFilePath)

	defer cleanup()
	createOperationsFile(patchmanData)
	runPatchmanUnityBundles()
	renameModifiedFiles()
	renameQueue = []string{}
}

func batchAssetImport(patchmanJson []byte) {
	var patchmanData PatchmanUnityStruct

	err := json.Unmarshal(patchmanJson, &patchmanData)
	if err != nil {
		global.FatalError(err)

	}

	vtolvrpath := global.FindVtolPath()

	patchmanData.OriginalFilePath = vtolvrpath + "\\" + patchmanData.OriginalFilePath
	patchmanData.ModifiedFilePath = vtolvrpath + patchmanData.OriginalFilePath + ".mod"
	renameQueue = append(renameQueue, patchmanData.OriginalFilePath)
	taintQueue = append(taintQueue, patchmanData.OriginalFilePath)

	defer cleanup()
	createOperationsFile(patchmanData)
	runPatchmanUnityAssets()
	renameModifiedFiles()
	renameQueue = []string{}

}

func buildTaintInfo() {
	var taintInfo TaintInfoStruct
	taintInfo.ModifiedFiles = taintQueue
	taintInfo.InstalledVersion = 1

	taintInfoJson, err := json.Marshal(taintInfo)
	if err != nil {
		global.FatalError(err)

	}

	vtolvrpath := global.FindVtolPath()

	if global.Exists(vtolvrpath + "\\patchman.json") {
		os.Remove(vtolvrpath + "\\patchman.json")
	}

	file, err := os.Create(vtolvrpath + "\\patchman.json")
	if err != nil {
		global.FatalError(err)

	}

	defer file.Close()

	_, err = file.Write(taintInfoJson)
	if err != nil {
		global.FatalError(err)

	}
}

func renameModifiedFiles() {
	for _, x := range renameQueue {
		err := os.Rename(x, x+".orig")
		if err != nil {
			global.FatalError(err)

		}
		err = os.Rename(x+".mod", x)
		if err != nil {
			global.FatalError(err)

		}
	}
}

func cleanup() {
	for _, x := range cleanupQueue {
		os.Remove(x)
	}
	cleanupQueue = []string{}
}
