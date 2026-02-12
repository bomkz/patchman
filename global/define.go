package global

import (
	_ "embed"
	"os/exec"

	"fyne.io/fyne/v2"
)

var TargetName string
var TargetBuildID string
var TargetAppID string
var TargetPath string
var TargetPathCheck string

//go:embed classData.tpk
var ClassDataTpk []byte

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
var Cmd *exec.Cmd
