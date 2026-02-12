//go:build windows
// +build windows

package main

import (
	"runtime"

	"fyne.io/fyne/v2/app"
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/guiHandler"
	"github.com/bomkz/patchman/indexHandler"
)

func main() {

	// Check for admin rights
	if isAdmin := checkAdmin(); !isAdmin {
		go promptElevate()
		global.OsName = runtime.GOOS

		initFyne()

		indexHandler.BuildIndex()
		guiHandler.InitGui()

	} else {
	}

}

func initFyne() {
	global.App = app.New()

	global.MainWindow = global.App.NewWindow("Patchman")

}
