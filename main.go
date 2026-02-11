//go:generate goversioninfo -icon=aircraft.ico -manifest=patchman.exe.manifest
package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"fyne.io/fyne/v2/app"
	"github.com/bomkz/patchman/global"
	guihandler "github.com/bomkz/patchman/guiHandler"
	"github.com/bomkz/patchman/indexHandler"
	"golang.org/x/sys/windows"
)

func main() {

	// Check for admin rights
	if isAdmin := checkAdmin(); !isAdmin {
		promptElevate()
		os.Exit(0)
	}

	// Handle Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Cleanup on exit
	go func() {
		<-c
		os.Exit(1)
	}()

	global.OsName = runtime.GOOS

	if global.OsName == "windows" {
		global.InitSteamReader()
	}

	initFyne()

	indexHandler.BuildIndex()

	guihandler.InitGui()
}

func initFyne() {
	global.App = app.New()

	global.MainWindow = global.App.NewWindow("Patchman")

}

// Elevates self as admin
func promptElevate() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

// Returns whether running as admin or not.
func checkAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	isadmin := false
	if err != nil {
		isadmin = false
	} else {
		isadmin = true
	}
	return isadmin
}
