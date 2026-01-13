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

	global.InitSteamReader()

	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "/?", "-?", "?", "/help", "/h", "-h", "--help", "h", "help":
			fmt.Println(helpArgument)
			os.Exit(0)

		case "/version", "/v", "--version", "-v", "v", "version":
			fmt.Println(versionArgument)
			os.Exit(0)
		}
	} else if len(os.Args) > 2 {
		log.Fatal("Unrecognized argument: " + os.Args[1] + "\nValid examples:\npatchman.exe [game buildid override] \npatchman.exe 18407725\npatchman.exe version\n patchman.exe help\npatchman.exe patchstatus")
	}

	initTview()

	index.BuildIndex()

	defer os.RemoveAll(global.Directory)
	if err := global.App.SetRoot(global.Root, true).Run(); err != nil {
		global.FatalError(err)
	}

	global.CleanProgramWorkingDirectory()

	global.ExitAppWithMessage("Done!")
}
