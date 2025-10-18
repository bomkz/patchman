package global

var VtolVersion string

var Internet bool = true
var InstalledVersion int

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

var StopApp = make(chan bool)

type StatusStruct struct {
	InstalledVersion int `json:"installedVersion"`
}

var Status StatusStruct
