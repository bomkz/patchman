package patchScriptHandler

import (
	_ "embed"

	"github.com/bomkz/patchman/patchScriptHandler/patchScriptInstaller/patchScriptOne"
)

// Main index content struct, stores game info, modifiable asset and content information, and patches.
type IndexStruct struct {

	// Name of the game, used for user friendliness. (i.e. VTOL VR)
	AppName string `json:"appName"`

	// AppID, used to find the path of the game using steam's library vdf.
	AppID string `json:"appID"`

	// The actual path of the game once the proper drive and folder is determined (i.e. <C:\\Program Files (x86)> --> \\Steam>\\steamapps\\common\VTOL VR <--)
	AppPath string `json:"appPath"`

	// Used to check if a user's give path is correct, patchman checks the existence of a file known to be at a given location, i.e. the game's executable ( i.e. VTOLVR.exe )
	LinuxPathCheck string `json:"linuxPathCheck"`

	// Motd, for informing users of potential bugs and updates.
	Motd string `json:"motd"`

	// List of modifiable game assets (i.e. ttsw_pullUp)
	ModifiableAssets []string `json:"modifiableAssets"`

	// List of modifiable game content like the base game or individual DLC.
	ModifiableContent []IndexModifiableContentStruct `json:"modifiableContent"`

	// List of actual patches for a given game.
	Patches []IndexPatchStruct `json:"content"`
}

// Part of IndexStruct, used for array of modifiable content (i.e. DLC and base games)
type IndexModifiableContentStruct struct {

	// Name of content, for user friendliness. (i.e. AH-94)
	ContentName string `json:"assetName"`

	// Path to content. (i.e. DLC\\1770480\\1770480)
	ContentPath string `json:"assetPath"`
}

// Part of IndexStruct, used for array of individual patches
type IndexPatchStruct struct {

	// Patch name, for user friendliness. (i.e. F16 RWR)
	PatchName string `json:"patchName"`

	// Patch description, for user friendliness. (i.e. This patch replaces the RWR sounds with the F16 RWR)
	PatchDesc string `json:"patchDesc"`

	// Patch author, for author credit. (i.e. bomkz)
	PatchAuthor string `json:"patchAuthor"`

	// Link to patch source (i.e. https://github.com/bomkz/f16rwr)
	PatchLink string `json:"patchLink"`

	// Individual patch variants list
	PatchVariants []IndexContentPatchVariantsStruct `json:"patchVariants"`
}

// Part of IndexContentStruct, used for array of individual patch variants
type IndexContentPatchVariantsStruct struct {

	// Variant name, for user friendliness (i.e. RWR Textures only)
	Variant string `json:"variant"`

	// Download link to Variant patchman zip file (i.e. https://examplelink.com/patchman.zip)
	DownloadLink string `json:"downloadLink"`
}

// Struct used to easily load and save presets
type PresetStruct struct {

	// List of assets for asset dropdown box
	Assets []string `json:"assets"`

	// Asset String for selected asset textview.
	AssetString string `json:"assetString"`

	// Asset currently selected by user as index.
	CurrentAsset int `json:"currentAsset"`

	// List of Content for content dropdown box
	Content []string `json:"content"`

	// Content String for selected content textview.
	ContentString string `json:"contentString"`

	// Content currently selected by user as index.
	CurrentContent int `json:"currentContent"`

	// Array storing assets selected for install
	PatchAssetSelection []patchScriptOne.AssetSelection `json:"assetSelection"`

	// Array storing content selected for install
	PatchContentSelection []patchScriptOne.ContentSelection `json:"contentSelection"`

	// Selected compression type
	Compression string `json:"compression"`
}

// Stores where presets should be saved
var savePath string

// Stores Message of the Day
var Motd string

// Array containing individual patch data.
var patchData []IndexPatchStruct

// Array containing patch list for dropdown box
var patches = []string{}

// Array containing variant list for dropdown box
var variants = []string{}

// Stores variant index currently selected by user
var currentVariant int = 0

// Stores game index currently selected by user
var currentGame int = 0

// Stores patch index currently selected by user
var currentPatch int = 0

// Preset contains information used to easily save and load content/asset presets
var preset PresetStruct

// Stores the game list for the dropdown box
var games []string

// Stores where the custom modpath is in custom form
var modPath string

// Variable contains main index json content
var index []IndexStruct

//go:embed patchman-unity.exe
var PatchmanUnityExe []byte

//go:embed patchman-unity
var PatchmanUnityLinux []byte

//go:embed classData.tpk
var ClassDataTpk []byte
