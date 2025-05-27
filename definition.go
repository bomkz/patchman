package main

import (
	"encoding/json"

	"github.com/rivo/tview"
)

var (
	root *tview.Pages
	app  *tview.Application
)

type SteamLibraryFolder struct {
	Path string `json:"path,omitempty"`
}

var indexURL = "https://github.com/bomkz/patchman-index/releases/latest/download/index.json"
var indexmem []byte
var zstdURL = "https://github.com/bomkz/patchman-resources/releases/download/76f53ddd-7484-465e-a349-a63e35f84dc7/zstd.exe"
var directory string

var internet bool = true
var nointernetinstruct string = `


Trouble fetching index.json, possible internet-related issue.
If running on a system without internet connection, please download https://github.com/bomkz/patchman-index/releases/latest/download/index.json in a separate device, and place it in a new folder in C:\patchman\ as C:\patchman\index.json

Press enter to continue...
`

var index []IndexStruct
var preindex PreIndexStruct

var nointernet string = `


Trouble fetching index.json, possible internet-related issue, but C:\patchman\index.json exists, now running in offline mode.

Press enter to continue...
`
var MaxPreIndexVersion int = 0

type PreIndexStruct struct {
	Content []PreIndexContentStruct `json:"content"`
}

type PreIndexContentStruct struct {
	Version string          `json:"version"`
	Content json.RawMessage `json:"content"`
}

type IndexStruct struct {
	Version     string                              `json:"version"`
	ObjectID    string                              `json:"id"`
	Name        string                              `json:"name"`
	Descritpion string                              `json:"description"`
	Versions    []IndexContentContentVersionsStruct `json:"versions"`
}

type IndexContentContentVersionsStruct struct {
	Version string                                     `json:"version"`
	Forms   []IndexContentContentVersionsFormsStruct   `json:"forms"`
	Content []IndexContentContentVersionsContentStruct `json:"content"`
}

type IndexContentContentVersionsContentStruct struct {
	ObjectID   string                                          `json:"id"`
	UUID       string                                          `json:"uuid"`
	Items      []IndexContentContentVersionsContentItemsStruct `json:"items"`
	PatchURL   string                                          `json:"patchURL"`
	UnpatchURL string                                          `json:"unpatchURL"`
}

type IndexContentContentVersionsContentItemsStruct struct {
	ObjectID  string `json:"id"`
	Available bool   `json:"available"`
}

type IndexContentContentVersionsFormsStruct struct {
	ObjectID string `json:"id"`
	FormType string `json:"type"`
	Name     string `json:"name"`
}

type PossibleCombinationStruct struct {
	Name      string
	ObjectIDs []string
}

var vtolversion string

type selectionstruct struct {
	Name         string
	IndexID      int
	ObjectID     string
	Description  string
	Combinations []PossibleCombinationStruct
	Combination  []string
	VariantID    string
}

var selection selectionstruct

type statusstruct struct {
	InstalledObjectId          string `json:"objectId"`
	InstalledUUID              string `json:"objectUUID"`
	InstalledVersionId         string `json:"versionId"`
	InstalledVariantId         string `json:"variantId"`
	InstalledVariantPatchURL   string `json:"variantPatchURL"`
	InstalledVariantUnpatchURL string `json:"variantUnpatchURL"`
	InstalledName              string `json:"variantName"`
}
type statustargetstruct struct {
	TargetObjectId          string
	TargetUUID              string
	TargetVersionId         string
	TargetVariantId         string
	TargetVariantPatchURL   string
	TargetVariantUnpatchURL string
}

var StatusTarget statustargetstruct

var Status statusstruct

var stop = make(chan bool)
