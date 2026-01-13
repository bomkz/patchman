package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/bomkz/patchman/global"
	"github.com/rivo/tview"
	"golang.org/x/sys/windows"
)

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

// Initializes Tview
func initTview() {
	global.App = tview.NewApplication()
	global.App.EnableMouse(true)

	global.Root = tview.NewPages()

	global.Root.SetBorder(false).SetTitle("VTOL VR Patch Manager")
}
