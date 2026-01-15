package installer

import "encoding/json"

type ActionScriptStruct struct {
	Action     string          `json:"action"`
	ActionData json.RawMessage `json:"actionData"`
}

var CompressionType string

type PatchmanUnityStruct struct {
	OriginalFilePath string                          `json:"originalFilePath"`
	ModifiedFilePath string                          `json:"modifiedFilePath"`
	Operations       []PatchmanUnityOperationsStruct `json:"operations"`
}
type CopyStruct struct {
	FileName    string `json:"fileName"`
	Destination string `json:"destination"`
}

var Assets []AssetSelection

type AssetSelection struct {
	AssetName string
	Modify    bool
}

var Content []ContentSelection

type ContentSelection struct {
	ContentName string
	ContentPath string
	Modify      bool
}
type PatchmanUnityOperationsStruct struct {
	Type      string  `json:"type"`
	AssetType string  `json:"assetType"`
	AssetName string  `json:"assetName"`
	AssetPath string  `json:"assetPath"`
	Length    float32 `json:"length"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Offset    int64   `json:"offset"`
	Size      int64   `json:"size"`
}

type TaintInfoStruct struct {
	ModifiedFiles    []string `json:"modifiedFiles"`
	InstalledVersion int      `json:"installedVersion"`
}
