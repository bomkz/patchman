package global

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"github.com/bomkz/steamutils"
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

var NoInternet string = `
Trouble downloading or reading https://github.com/bomkz/patchman-index, possible internet-related issue or unsupported patchman version, switching to offline mode.
`
var Directory string

var OsName string

var PatchmanUnityDir string

var pwdDir string

var gwdDir string

var IndexData []byte

var IndexMotd string

var SteamPath string

var App fyne.App

var MainWindow fyne.Window
