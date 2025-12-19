package actionScriptOne

import "encoding/json"

type ActionScriptStruct struct {
	Action     string          `json:"action"`
	ActionData json.RawMessage `json:"actionData"`
}

type PatchmanUnityStruct struct {
	OriginalFilePath string                          `json:"originalFilePath"`
	ModifiedFilePath string                          `json:"modifiedFilePath"`
	Operations       []PatchmanUnityOperationsStruct `json:"operations"`
}
type CopyStruct struct {
	FileName    string `json:"fileName"`
	Destination string `json:"destination"`
}

var renameQueue []string
var cleanupQueue []string

var taintQueue []string

type PatchmanUnityOperationsStruct struct {
	Type      string `json:"type"`
	AssetType string `json:"assetType"`
	AssetName string `json:"assetName"`
	AssetPath string `json:"assetPath"`
}

type TaintInfoStruct struct {
	ModifiedFiles    []string `json:"modifiedFiles"`
	InstalledVersion int      `json:"installedVersion"`
}
