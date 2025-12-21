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
	"github.com/rivo/tview"
)

func main() {

	admin := checkAdmin()

	if !admin {
		promptElevate()
		os.Exit(0)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		global.CleanDir()
		os.Exit(1)
	}()

	var err error
	global.VtolVersion, err = getVtolVersion()
	fmt.Println(global.VtolVersion)
	if err != nil {
		log.Fatal(err)
	}

	err = createDir()
	if err != nil {
		log.Fatal(err)
	}

	vtolvrpath := global.FindVtolPath()

	if global.Exists(vtolvrpath + "\\patchman.json") {
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

		case "/status", "/s", "-s", "--status", "s", "status":
			if global.Exists(vtolvrpath + "\\patchman.json") {
				fmt.Println(index.TaintInfo())
			} else {
				fmt.Println(statusArgument)
			}
			os.Exit(0)
		default:
			global.VtolVersion = os.Args[1]
		}
	} else if len(os.Args) > 2 {
		log.Fatal("Unrecognized argument: " + os.Args[1] + "\nValid examples:\npatchman.exe [game buildid override] \npatchman.exe 18407725\npatchman.exe version\n patchman.exe help\npatchman.exe patchstatus")
	}

	global.App = tview.NewApplication()
	global.App.EnableMouse(true)

	global.Root = tview.NewPages()

	global.Root.SetBorder(false).SetTitle("VTOL VR Patch Manager")

	buildForm()

	defer os.RemoveAll(global.Directory)

	runApp()

	global.ExitAppWithMessage("Done!")
}

func runApp() {

	if err := global.App.SetRoot(global.Root, true).Run(); err != nil {
		log.Panic(err)
	}

}
