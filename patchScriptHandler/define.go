package patchScriptHandler

import _ "embed"

type IndexStruct struct {
	AppName string               `json:"appName"`
	AppID   string               `json:"appID"`
	AppPath string               `json:"appPath"`
	Motd    string               `json:"motd"`
	Content []IndexContentStruct `json:"content"`
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

var games []string

var modPath string

var index []IndexStruct

//go:embed patchman-unity.exe
var PatchmanUnityExe []byte

//go:embed classData.tpk
var ClassDataTpk []byte

var Motd string
var patchData []IndexContentStruct
var patches = []string{}
var variants = []string{}
var currentVariant int = 0
var currentGame int = 0
var currentSelection int = 0
