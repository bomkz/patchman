package patchScript

import (
	"encoding/json"
	"os"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller/patchScriptOne"
)

func HandleActionScript(actionscript []byte) {
	var actionScript ActionScriptStruct

	err := json.Unmarshal(actionscript, &actionScript)
	if err != nil {
		global.FatalError(err)
	}

	if actionScript.Patchscriptversion == 1 {
		patchScriptOne.HandleActions(actionScript.Data)
	}
}

func BeginTestJson() {
	jsonFile, err := os.ReadFile(".\\patchscript.json")
	if err != nil {
		global.FatalError(err)
	}
	HandleActionScript(jsonFile)
}
