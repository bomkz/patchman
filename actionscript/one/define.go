package actionScriptOne

type ActionScriptStruct struct {
	Action     string `json:"action"`
	ActionData []byte `json:"actionData"`
}

type ExportFromBundleStruct struct {
	BundlePath      string                            `json:"bundlePath"`
	ExportSelection []ExportFromBundleSelectionStruct `json:"exportSelection"`
}

type ExportFromBundleSelectionStruct struct {
	SelectionName string `json:"selectionName"`
	ExportName    string `json:"exportName"`
	ExportPath    string `json:"exportPath"`
}

type ExportAllFromBundleStruct struct {
	BundlePath string `json:"bundlePath"`
	ExportPath string `json:"exportPath"`
}

type ImportIntoBundleStruct struct {
	BundlePath      string                            `json:"bundlePath"`
	ImportSelection []ImportIntoBundleSelectionStruct `json:"importSelection"`
}

type ImportIntoBundleSelectionStruct struct {
	SelectionName string `json:"selectionName"`
	ImportName    string `json:"importName"`
	ImportPath    string `json:"importPath"`
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

type GenerateHashStruct struct {
	HashPath string   `json:"hashPath"`
	FilePath []string `json:"filePath"`
}

type GenerateHashBatch struct {
	HashPath   string `json:"hashPath"`
	FolderPath string `json:"folderPath"`
}
