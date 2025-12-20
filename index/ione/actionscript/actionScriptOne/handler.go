package actionScriptOne

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/bomkz/patchman/global"
)

// patcher.exe exportfrombundle --bundle "C:\BundlePath\unity.assets" --assetName "exampleAsset" --exportPath "C:\ExportPath\ExportName"

func runPatchmanUnityBundles() {
	cmd := exec.Command(".\\patchman-unity.exe", "batchimportbundle", ".\\operations.json", CompressionType)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		global.FatalError(err)

	}

	if out.String() == "Done!" {
		return
	} else {
		global.FatalError(errors.New(out.String()))
	}
}

func StatusUpdater() {

	for {
		<-refreshStatus

		time.Sleep(500 * time.Millisecond)
		ClearScreen()
		fmt.Println("Current Action: " + installStatus.Current.ActionName + " " + installStatus.Current.Filename)
		currentStepString := "(" + strconv.Itoa(installStatus.Current.StepsCompleted) + "/" + strconv.Itoa(installStatus.Current.TotalSteps) + ") "
		fmt.Println("Step: " + currentStepString + installStatus.Current.CurrentAction)
		overallProgress := "(" + strconv.Itoa(installStatus.completed) + "/" + strconv.Itoa(installStatus.total) + ")"
		fmt.Println("Total Progress: " + overallProgress)
	}

}

func runPatchmanUnityAssets() {
	cmd := exec.Command(".\\patchman-unity.exe", "batchimportasset", ".\\operations.json")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		global.FatalError(err)

	}

	if out.String() == "Done!" {
		return
	} else {
		log.Panic("Uh oh...")
	}

}

func createOperationsFile(opData PatchmanUnityStruct) {

	file, err := os.Create("operations.json")
	if err != nil {
		global.FatalError(err)

	}

	defer file.Close()

	cleanupQueue = append(cleanupQueue, ".\\operations.json")

	jsonData, err := json.Marshal(opData)
	if err != nil {
		global.FatalError(err)

	}

	_, err = file.Write(jsonData)
	if err != nil {
		global.FatalError(err)

	}
}
