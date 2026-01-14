package formHandler

import (
	"os"

	"github.com/bomkz/patchman/formHandler/installHandler"
	"github.com/bomkz/patchman/global"
)

func install(filePath string) {
	global.UnpackDependencies()
	global.UnzipIntoProgramWorkingDirectory(filePath)
	patchscript := global.Assure(os.ReadFile(global.Directory + "\\patchscript.json"))
	installHandler.HandleActionScript(patchscript)

	global.ExitTview()

	global.ExitAppWithMessage("Done!")

}
