package global

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed classdata.tpk
var ClassDataTpk []byte

var Internet bool = true

var NoInternet string = `
Trouble downloading or reading https://github.com/bomkz/patchman-index, possible internet-related issue or unsupported patchman version, switching to offline mode.
`
var Directory string

var OsName string

var Installed = make(chan bool)

var SteamFound bool

var PatchmanUnityDir string

var pwdDir string

var gwdDir string

var IndexData []byte

var IndexMotd string

var SteamPath string

var Helper bool

var App fyne.App

var MainWindow fyne.Window

type Message struct {
	Type    string `json:"type"`
	Payload []byte `json:"payload"`
}
