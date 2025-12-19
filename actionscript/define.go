package actionScript

type ActionScriptStruct struct {
	Patchscriptversion int    `json:"patchScriptVersion"`
	Motd               string `json:"motd"`
	Data               []byte `json:"data"`
}
