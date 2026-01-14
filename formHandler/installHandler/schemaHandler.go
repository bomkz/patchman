package installHandler

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/formHandler/installHandler/installer"
	"github.com/bomkz/patchman/global"
)

func HandleActionScript(actionscript []byte) {
	var actionScript ActionScriptStruct

	global.AssureNoReturn(json.Unmarshal(actionscript, &actionScript))

	if actionScript.Patchscriptversion == 1 {
		installer.HandleActions(actionScript.Data)
	}
}

func BeginTestJson() {
	jsonFile := global.Assure(os.ReadFile(global.Directory + ".\\patchscript.json"))
	HandleActionScript(jsonFile)
}
