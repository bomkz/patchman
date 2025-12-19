package actionScriptOne

type ActionScriptStruct struct {
	Action     string `json:"action"`
	ActionData []byte `json:"actionData"`
}

type ImportIntoBundleStruct struct {
	BundlePath      string                            `json:"originalFilePath"`
	ImportSelection []ImportIntoBundleSelectionStruct `json:"importSelection"`
	SavePath        string                            `json:"ModifiedFilePath"`
}

type ImportIntoBundleSelectionStruct struct {
	AssetName string `json:"assetName"`
	AssetType string `json:"assetType"`
	AssetPath string `json:"assetPath"`
	Type      string `json:"type"`
}

type ImportIntoAssetStruct struct {
	LoadAssetFiles []string `json:"loadAssetFiles"`
}

type ImportIntoAssetSelectionStruct struct {
	AssetName         string `json:"assetName"`
	AssetType         string `json:"assetType"`
	AssetHash         string `json:"assetHash"`
	ImportedAssetPath string `json:"importedAssetPath"`
	ImportedAssetHash string `json:"importedAssetHash"`
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
