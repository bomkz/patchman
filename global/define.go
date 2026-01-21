package global

import (
	_ "embed"

	"github.com/bomkz/patchman/steamutils"
	"github.com/rivo/tview"
)

var TargetName string
var TargetBuildID string
var TargetAppID string
var TargetPath string
var TargetPathCheck string

//go:embed patchman-unity.exe
var PatchmanUnityExe []byte

//go:embed patchman-unity
var PatchmanUnityLinux []byte

//go:embed classData.tpk
var ClassDataTpk []byte

var SteamReader steamutils.SteamReader
var Internet bool = true
var NoInternetInstruct string = `


Trouble fetching index.json, possible internet-related issue.
If running on a system without internet connection, please download https://github.com/bomkz/patchman-index/releases/latest/download/index.json in a separate device, and place it in a new folder in C:\patchman\ as C:\patchman\index.json

Press enter to continue...
`
var NoInternet string = `


Trouble fetching index.json, possible internet-related issue, but C:\patchman\index.json exists, now running in offline mode.

Press enter to continue...
`
var Directory string

var OsName string

var PatchmanUnityDir string

var App *tview.Application
var Root *tview.Pages

var pwdDir string

var gwdDir string
