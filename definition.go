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
