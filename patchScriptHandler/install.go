package patchScriptHandler

import (
	"os"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller"
)

func install(filePath string) {
	unpackDependencies()
	global.UnzipIntoProgramWorkingDirectory(filePath)
	patchscript := global.Assure(os.ReadFile(global.Directory + "\\patchscript.json"))
	patchScriptInstaller.HandleActionScript(patchscript)

	global.ExitTview()

	global.ExitAppWithMessage("Done!")

}

func unpackDependencies() {
	global.CreateAndWriteProgramWorkingDirectory(PatchmanUnityExe, "patchman-unity.exe")
	global.CreateAndWriteProgramWorkingDirectory(ClassDataTpk, "classdata.tpk")
}
