package main

import (
	"github.com/rivo/tview"
)

var (
	root *tview.Pages
	app  *tview.Application
)

type SteamLibraryFolder struct {
	Path string `json:"path,omitempty"`
}

var stop = make(chan bool)

var helpArgument string = `Valid arguments: patchman.exe
	help    		- Displays this help message
	alias:
		h
		-h
		--help
		/h
		/help
	version 		- Displays patchman's version
	alias:
		v
		-v
		--version
		/v
		/version
	status  		- Displays the current status
	alias:
		s
		-s
		--status
		/s
		/status

If VTOL VR receives a new update and patches are yet to be marked as compatible, you could override the Build ID version by looking up a Build ID from https://github.com/bomkz/patchman-index and using it as follows:
	patchman.exe <buildid>
	patchman.exe 18407725`
var versionArgument string = "Patchman " + timestamp
var statusArgument string = "patchman.json does not exist, game is likely unpatched, or user removed patchman.json"

var timestamp string = "1748349190"
