package actionScriptOne

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/bomkz/patchman/global"
)

// patcher.exe exportfrombundle --bundle "C:\BundlePath\unity.assets" --assetName "exampleAsset" --exportPath "C:\ExportPath\ExportName"

func runPatchmanUnityBundles() error {
	cmd := exec.Command(".\\patchman-unity.exe", "batchimportbundle", ".\\operations.json", CompressionType)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		global.FatalError(err)

	}

	if out.String() != "Done!" {
		return errors.New(out.String())
	}
	return nil
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

func runPatchmanUnityAssets() error {
	cmd := exec.Command(".\\patchman-unity.exe", "batchimportasset", ".\\operations.json")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return err
	}

	if out.String() != "Done!" {
		return errors.New(out.String())
	}

	return nil
}

func createOperationsFile(opData PatchmanUnityStruct) error {

	file, err := os.Create("operations.json")
	if err != nil {
		return err
	}

	defer file.Close()

	cleanupQueue = append(cleanupQueue, ".\\operations.json")

	jsonData, err := json.Marshal(opData)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}
