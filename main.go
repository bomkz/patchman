//go:generate goversioninfo -icon=aircraft.ico -manifest=patchman.exe.manifest
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/index"
)

func main() {

	isAdmin := checkAdmin()

	if !isAdmin {
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

	var err error
	global.SteamPath, err = global.GetSteamPath()
	if err != nil {
		log.Fatal(err)
	}

	global.TargetVersion, err = global.GetAppIDBuildIDVersion("667970")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(global.TargetVersion)

	err = createDir()
	if err != nil {
		log.Fatal(err)
	}

	global.TargetPath, err = global.FindAppIDPath("667970")
	global.TargetPath += "\\steamapps\\common\\VTOL VR\\"

	if global.Exists(global.TargetPath + "\\patchman.json") {
		index.ReadTaint()
	}
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "/?", "-?", "?", "/help", "/h", "-h", "--help", "h", "help":
			fmt.Println(helpArgument)
			os.Exit(0)

		case "/version", "/v", "--version", "-v", "v", "version":
			fmt.Println(versionArgument)
			os.Exit(0)
		default:
			global.TargetVersion = os.Args[1]
		}
	} else if len(os.Args) > 2 {
		log.Fatal("Unrecognized argument: " + os.Args[1] + "\nValid examples:\npatchman.exe [game buildid override] \npatchman.exe 18407725\npatchman.exe version\n patchman.exe help\npatchman.exe patchstatus")
	}

	initTview()

	err = index.BuildIndex()
	if err != nil {
		global.FatalError(err)
	}

	defer os.RemoveAll(global.Directory)
	if err := global.App.SetRoot(global.Root, true).Run(); err != nil {
		global.FatalError(err)
	}

	global.ExitAppWithMessage("Done!")
}
