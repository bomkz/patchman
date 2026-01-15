package installHandler

import "encoding/json"

type ActionScriptStruct struct {
	Patchscriptversion int             `json:"patchScriptVersion"`
	Motd               string          `json:"motd"`
	Data               json.RawMessage `json:"data"`
}
