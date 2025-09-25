package actionScript

type ActionScriptStruct struct {
	Patchscriptversion int    `json:"patchscriptVersion"`
	Motd               string `json:"motd"`
	Data               []byte `json:"data"`
}
