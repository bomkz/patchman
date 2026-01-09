package patchScriptHandler

import _ "embed"

type IndexStruct struct {
	Version string               `json:"version"`
	Content []IndexContentStruct `json:"content"`
	MOTD    string               `json:"motd"`
}

type IndexContentStruct struct {
	PatchName     string                            `json:"patchName"`
	PatchDesc     string                            `json:"patchDesc"`
	PatchAuthor   string                            `json:"patchAuthor"`
	PatchLink     string                            `json:"patchLink"`
	PatchVariants []IndexContentPatchVariantsStruct `json:"patchVariants"`
}

type IndexContentPatchVariantsStruct struct {
	Variant      string `json:"variant"`
	DownloadLink string `json:"downloadLink"`
}

var modPath string

var Index IndexStruct
var cleanupQueue = []string{}

//go:embed patchman-unity.exe
var PatchmanUnityExe []byte

//go:embed classData.tpk
var ClassDataTpk []byte
