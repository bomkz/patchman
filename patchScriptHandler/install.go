package patchScriptHandler

import (
	"os"
	"time"

	"github.com/bomkz/patchman/global"
	patchScript "github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller"
	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller/patchScriptOne"
)

func install(filePath string) {
	if global.Status.InstalledVersion != 99 {
		switch global.Status.InstalledVersion {
		case 0:
			global.ExitAppWithMessage("Verify Game FIles in Steam before retrying install...")
		case 1:
			patchScriptOne.Uninstall()
		}
	}

	if !global.Exists(".\\patchman-unity.exe") {
		os.WriteFile(".\\patchman-unity.exe", PatchmanUnityExe, os.ModeAppend)
	}
	if !global.Exists(".\\classdata.tpk") {
		os.WriteFile(".\\classdata.tpk", ClassDataTpk, os.ModeAppend)
	}
	err := global.UnzipIntoRoot(filePath)
	if err != nil {
		global.FatalError(err)
	}

	defer cleanup()

	patchscript, err := os.ReadFile(global.Directory + "\\patchscript.json")
	if err != nil {
		global.FatalError(err)

	}
	patchScript.HandleActionScript(patchscript)

	os.RemoveAll(global.Directory)
	global.ExitTview()

	time.Sleep(1000 * time.Second)
	global.ExitAppWithMessage("Done!")

}

func uninstall() {
	patchScriptOne.Uninstall()

}

func cleanup() {
	for _, x := range cleanupQueue {
		os.Remove(x)
	}
	cleanupQueue = []string{}
}
