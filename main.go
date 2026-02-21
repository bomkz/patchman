//go:build windows
// +build windows

package main

import (
	"fmt"
	"runtime"

	"fyne.io/fyne/v2/app"
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/guiHandler"
	"github.com/bomkz/patchman/indexHandler"
	"github.com/bomkz/patchman/ipc"
	"github.com/bomkz/patchman/ipc/procman"
)

func main() {

	fmt.Println("meow")
	procman.InitProcess()

	if !global.Helper {
		global.OsName = runtime.GOOS

		initFyne()

		indexHandler.BuildIndex()
		guiHandler.InitGui()
	} else {
		ipc.StartHelper()
	}

}

func initFyne() {
	global.App = app.New()

	global.MainWindow = global.App.NewWindow("Patchman")

}
