package actionScriptOne

import (
	"encoding/json"
	"log"
	"os"
)

func HandleActions(actionData []byte) {
	var actionScript []ActionScriptStruct

	err := json.Unmarshal(actionData, &actionScript)
	if err != nil {
		log.Fatal(err)
	}

	for _, x := range actionScript {
		switch x.Action {
		case "importbundle":
			batchBundleImport(x.ActionData)
		case "importasset":
			batchAssetImport(x.ActionData)
		}
	}

}

func batchBundleImport(patchmanJson []byte) {
	var patchmanData PatchmanUnityStruct

	err := json.Unmarshal(patchmanJson, &patchmanData)
	if err != nil {
		log.Fatal(err)
	}

	patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"
	renameQueue = append(renameQueue, patchmanData.OriginalFilePath)

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
		log.Fatal(err)
	}

	patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"
	renameQueue = append(renameQueue, patchmanData.OriginalFilePath)

	defer cleanup()
	createOperationsFile(patchmanData)
	runPatchmanUnityAssets()
	renameModifiedFiles()
	renameQueue = []string{}

}

func renameModifiedFiles() {
	for _, x := range renameQueue {
		err := os.Rename(x, x+".orig")
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(x+".mod", x)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var renameQueue []string
var cleanupQueue []string

func cleanup() {
	for _, x := range cleanupQueue {
		os.Remove(x)
	}
}
