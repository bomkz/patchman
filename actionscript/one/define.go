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

type PatchmanUnityOperationsStruct struct {
	Type      string `json:"type"`
	AssetType string `json:"assetType"`
	AssetName string `json:"assetName"`
	AssetPath string `json:"assetPath"`
}
