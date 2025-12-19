package actionScript

import (
	"encoding/json"
	"log"
	"os"

	"github.com/bomkz/patchman/index/ione/actionscript/actionScriptOne"
)

func HandleActionScript(actionscript []byte) {
	var actionScript ActionScriptStruct

	err := json.Unmarshal(actionscript, &actionScript)
	if err != nil {
		log.Fatal(err)
	}

	if actionScript.Patchscriptversion == 1 {
		actionScriptOne.HandleActions(actionScript.Data)
	}
}

func BeginTestJson() {
	jsonFile, err := os.ReadFile(".\\patchscript.json")
	if err != nil {
		log.Fatal(err)
	}
	HandleActionScript(jsonFile)
}
