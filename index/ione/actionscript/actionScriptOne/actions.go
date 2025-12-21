package actionScriptOne

import (
	"encoding/json"
	"os"
	"os/exec"
	"runtime"

	"github.com/bomkz/patchman/global"
)

func Uninstall() {
	taintInfo := readTaint()
	if len(taintInfo.ModifiedFiles) == 0 {
		return
	}
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
	taintfile, err := os.ReadFile(vtolvrpath + "patchman.json")
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
	vtolvrpath := global.FindVtolPath()

	err := os.Remove(filePath)
	if err != nil {
		os.Remove(vtolvrpath + "patchman.json")
		global.FatalError(err)
	}

	err = os.Rename(filePath+".orig", filePath)
	if err != nil {
		os.Remove(vtolvrpath + "patchman.json")
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
		var newInstallStatus installStatusActionsQueueStruct
		newInstallStatus.CurrentAction = ""
		installStatus.total += 1
		switch x.Action {
		case "importbundle":
			var tmpData PatchmanUnityStruct
			err := json.Unmarshal(x.ActionData, &tmpData)
			if err != nil {
				global.FatalError(err)
			}
			newInstallStatus.Filename = tmpData.OriginalFilePath
			newInstallStatus.ActionName = " patching bundle: "
			installStatus.Pending = append(installStatus.Pending, newInstallStatus)
		case "importasset":
			var tmpData PatchmanUnityStruct
			err := json.Unmarshal(x.ActionData, &tmpData)
			if err != nil {
				global.FatalError(err)
			}
			newInstallStatus.Filename = tmpData.OriginalFilePath
			newInstallStatus.ActionName = " patching asset: "
			installStatus.Pending = append(installStatus.Pending, newInstallStatus)
		case "copy":
			var tmpCopy CopyStruct
			err := json.Unmarshal(x.ActionData, &tmpCopy)
			if err != nil {
				global.FatalError(err)
			}
			newInstallStatus.Filename = tmpCopy.FileName

			newInstallStatus.ActionName = " copying file: "
			installStatus.Pending = append(installStatus.Pending, newInstallStatus)

		}
	}

	global.ExitTview()
	go StatusUpdater()

	for _, x := range actionScript {
		switch x.Action {
		case "importbundle":
			err = batchBundleImport(x.ActionData)
			if err != nil {
				global.FatalError(err)
				return
			}
		case "importasset":
			err = batchAssetImport(x.ActionData)
			if err != nil {
				global.FatalError(err)
				return
			}
		case "copy":
			err = handleCopy(x.ActionData)
			if err != nil {
				global.FatalError(err)
				return
			}
		}
	}

	buildTaintInfo()

	global.ExitAppWithMessage("Done!")
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func handleCopy(actionData []byte) error {
	var copyData CopyStruct

	err := json.Unmarshal(actionData, &copyData)
	if err != nil {
		return err
	}

	var tmpInstallStatus []installStatusActionsQueueStruct
	for x, y := range installStatus.Pending {
		if y.Filename == copyData.FileName {
			installStatus.Pending[x].Id = copyData.FileName
			installStatus.Pending[x].TotalSteps = 1
			installStatus.Pending[x].CurrentAction = "Copying file"
			installStatus.Current = installStatus.Pending[x]
		} else {
			tmpInstallStatus = append(tmpInstallStatus, y)
		}
	}

	installStatus.Pending = tmpInstallStatus

	vtolvrpath := global.FindVtolPath()
	copyData.Destination = vtolvrpath + copyData.Destination

	refreshStatus <- true

	err = global.MoveFile(global.Directory+"\\"+copyData.FileName, copyData.Destination)
	if err != nil {
		return err
	}
	installStatus.Current.CurrentAction = "Done!"
	installStatus.Current.StepsCompleted = 1
	installStatus.completed += 1
	refreshStatus <- true
	return nil
}

func batchBundleImport(patchmanJson []byte) error {
	var patchmanData PatchmanUnityStruct

	err := json.Unmarshal(patchmanJson, &patchmanData)
	if err != nil {
		global.FatalError(err)

	}
	var tmpPending []installStatusActionsQueueStruct
	for x, y := range installStatus.Pending {
		if y.Filename == patchmanData.OriginalFilePath {
			vtolvrpath := global.FindVtolPath()
			patchmanData.OriginalFilePath = vtolvrpath + "\\" + patchmanData.OriginalFilePath
			patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"
			installStatus.Pending[x].Id = patchmanData.OriginalFilePath

			installStatus.Current = installStatus.Pending[x]
		} else {
			tmpPending = append(tmpPending, y)
		}
	}

	if !global.Exists(patchmanData.OriginalFilePath) {
		return nil
	}

	renameQueue = append(renameQueue, patchmanData.OriginalFilePath)
	taintQueue = append(taintQueue, patchmanData.OriginalFilePath)

	installStatus.Pending = tmpPending
	installStatus.Current.TotalSteps = 4
	installStatus.Current.CurrentAction = "Creating Bundle Patch..."
	refreshStatus <- true

	err = createOperationsFile(patchmanData)
	if err != nil {
		return err
	}
	installStatus.Current.StepsCompleted = 1
	installStatus.Current.CurrentAction = "Installing Bundle Patch..."
	refreshStatus <- true

	err = runPatchmanUnityBundles()
	if err != nil {
		return err
	}
	installStatus.Current.StepsCompleted = 2
	installStatus.Current.CurrentAction = "Finalizing Bundle Patch..."
	refreshStatus <- true

	err = renameModifiedFiles()
	if err != nil {
		return err
	}
	installStatus.Current.StepsCompleted = 3
	installStatus.Current.CurrentAction = "Cleaning up temporary files..."
	refreshStatus <- true

	cleanup()
	installStatus.Current.StepsCompleted = 4
	installStatus.Current.CurrentAction = "Done!"
	refreshStatus <- true
	installStatus.Current = installStatusActionsQueueStruct{}
	installStatus.completed += 1

	renameQueue = []string{}
	return nil
}

func batchAssetImport(patchmanJson []byte) error {
	var patchmanData PatchmanUnityStruct

	err := json.Unmarshal(patchmanJson, &patchmanData)
	if err != nil {
		return err
	}

	var tmpPending []installStatusActionsQueueStruct
	for x, y := range installStatus.Pending {
		if y.Filename == patchmanData.OriginalFilePath {
			vtolvrpath := global.FindVtolPath()
			patchmanData.OriginalFilePath = vtolvrpath + "\\" + patchmanData.OriginalFilePath
			patchmanData.ModifiedFilePath = patchmanData.OriginalFilePath + ".mod"
			installStatus.Pending[x].Id = patchmanData.OriginalFilePath

			installStatus.Current = installStatus.Pending[x]
		} else {
			tmpPending = append(tmpPending, y)
		}
	}

	renameQueue = append(renameQueue, patchmanData.OriginalFilePath)
	taintQueue = append(taintQueue, patchmanData.OriginalFilePath)

	installStatus.Pending = tmpPending
	installStatus.Current.TotalSteps = 4
	installStatus.Current.CurrentAction = "Creating Asset Patch..."
	refreshStatus <- true

	err = createOperationsFile(patchmanData)
	if err != nil {
		return err
	}
	installStatus.Current.StepsCompleted = 1
	installStatus.Current.CurrentAction = "Installing Asset Patch..."
	refreshStatus <- true

	err = runPatchmanUnityAssets()
	if err != nil {
		return err
	}
	installStatus.Current.StepsCompleted = 2
	installStatus.Current.CurrentAction = "Finalizing Asset Patch..."
	refreshStatus <- true

	err = renameModifiedFiles()
	if err != nil {
		return err
	}
	installStatus.Current.StepsCompleted = 3
	installStatus.Current.CurrentAction = "Cleaning up temporary files..."
	refreshStatus <- true
	cleanup()
	installStatus.Current.StepsCompleted = 4
	installStatus.Current.CurrentAction = "Done!"
	refreshStatus <- true
	installStatus.Current = installStatusActionsQueueStruct{}

	installStatus.completed += 1

	renameQueue = []string{}
	return nil

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

func renameModifiedFiles() error {
	for _, x := range renameQueue {
		err := os.Rename(x, x+".orig")
		if err != nil {
			return err
		}
		err = os.Rename(x+".mod", x)
		if err != nil {
			return err
		}
	}
	return nil
}

func cleanup() {
	for _, x := range cleanupQueue {
		os.Remove(x)
	}
	cleanupQueue = []string{}
}
