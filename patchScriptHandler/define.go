package patchScriptHandler

import (
	_ "embed"
)

type IndexModifiableContentStruct struct {
	AssetName string `json:"assetName"`
	AssetPath string `json:"assetPath"`
}

type IndexStruct struct {
	AppName           string                         `json:"appName"`
	AppID             string                         `json:"appID"`
	AppPath           string                         `json:"appPath"`
	LinuxPathCheck    string                         `json:"linuxPathCheck"`
	Motd              string                         `json:"motd"`
	ModifiableAssets  []string                       `json:"modifiableAssets"`
	ModifiableContent []IndexModifiableContentStruct `json:"modifiableContent"`
	Content           []IndexContentStruct           `json:"content"`
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

//go:embed patchman-unity
var PatchmanUnityLinux []byte

//go:embed classData.tpk
var ClassDataTpk []byte

var Motd string
var patchData []IndexContentStruct
var patches = []string{}
var variants = []string{}
var currentVariant int = 0
var currentGame int = 0
var currentSelection int = 0
