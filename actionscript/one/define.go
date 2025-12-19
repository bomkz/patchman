package actionScriptOne

type ActionScriptStruct struct {
	Action     string `json:"action"`
	ActionData []byte `json:"actionData"`
}

type PatchmanUnityStruct struct {
	Version          int                             `json:"version"`
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
