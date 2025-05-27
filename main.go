//go:generate goversioninfo -icon=aircraft.ico -manifest=patchman.exe.manifest
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rivo/tview"
)

var help string = `Valid arguments: patchman.exe
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
var patchmanversion string = "Patchman " + timestamp
var timestamp string = "1748349190"

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
		CleanDir()
		os.Exit(1)
	}()

	var err error
	vtolversion, err = getVtolVersion()
	if err != nil {
		log.Fatal(err)
	}

	err = createDir()
	if err != nil {
		log.Fatal(err)
	}

	err = downloadIndex(indexURL)
	if err != nil {
		internet = false
		exists := checkLocalDbNoInternet()
		if !exists {
			fmt.Println(err, nointernetinstruct)
			os.RemoveAll(directory)
			fmt.Scanln()
			os.Exit(1)
		} else {

			fmt.Println(nointernet)
			fmt.Scanln()
		}
	}

	err = parseIndex()
	if err != nil {
		os.RemoveAll(directory)
		log.Fatal(err)
	}
	vtolvrpath := findVtolPath()

	if exists(vtolvrpath + "\\patchman.json") {
		readTaint()
	}
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "help":
			fmt.Println(help)
			os.Exit(0)
		case "h":
			fmt.Println(help)
			os.Exit(0)
		case "--help":
			fmt.Println(help)
			os.Exit(0)
		case "-h":
			fmt.Println(help)
			os.Exit(0)
		case "/h":
			fmt.Println(help)
			os.Exit(0)
		case "/help":
			fmt.Println(help)
			os.Exit(0)
		case "?":
			fmt.Println(help)
			os.Exit(0)
		case "-?":
			fmt.Println(help)
			os.Exit(0)
		case "/?":
			fmt.Println(help)
			os.Exit(0)
		case "version":
			fmt.Println(patchmanversion)
			os.Exit(0)
		case "v":
			fmt.Println(patchmanversion)
			os.Exit(0)
		case "-v":
			fmt.Println(patchmanversion)
			os.Exit(0)
		case "--version":
			fmt.Println(patchmanversion)
			os.Exit(0)
		case "/v":
			fmt.Println(patchmanversion)
			os.Exit(0)
		case "/version":
			fmt.Println(patchmanversion)
			os.Exit(0)
		case "status":
			if exists(vtolvrpath + "\\patchman.json") {
				fmt.Println("Current Variant " + Status.InstalledName + " Object ID " + Status.InstalledObjectId + " Variant ID " + Status.InstalledVariantId + " Version ID " + Status.InstalledVersionId)
			} else {
				fmt.Println("patchman.json does not exist, game is likely unpatched, or user removed patchman.json")
			}
			os.Exit(0)
		case "s":
			if exists(vtolvrpath + "\\patchman.json") {
				fmt.Println("Current Variant " + Status.InstalledName + " Object ID " + Status.InstalledObjectId + " Variant ID " + Status.InstalledVariantId + " Version ID " + Status.InstalledVersionId)
			} else {
				fmt.Println("patchman.json does not exist, game is likely unpatched, or user removed patchman.json")
			}
			os.Exit(0)
		case "--status":
			if exists(vtolvrpath + "\\patchman.json") {
				fmt.Println("Current Variant " + Status.InstalledName + " Object ID " + Status.InstalledObjectId + " Variant ID " + Status.InstalledVariantId + " Version ID " + Status.InstalledVersionId)
			} else {
				fmt.Println("patchman.json does not exist, game is likely unpatched, or user removed patchman.json")
			}
			os.Exit(0)
		case "-s":
			if exists(vtolvrpath + "\\patchman.json") {
				fmt.Println("Current Variant " + Status.InstalledName + " Object ID " + Status.InstalledObjectId + " Variant ID " + Status.InstalledVariantId + " Version ID " + Status.InstalledVersionId)
			} else {
				fmt.Println("patchman.json does not exist, game is likely unpatched, or user removed patchman.json")
			}
			os.Exit(0)
		case "/s":
			if exists(vtolvrpath + "\\patchman.json") {
				fmt.Println("Current Variant " + Status.InstalledName + " Object ID " + Status.InstalledObjectId + " Variant ID " + Status.InstalledVariantId + " Version ID " + Status.InstalledVersionId)
			} else {
				fmt.Println("patchman.json does not exist, game is likely unpatched, or user removed patchman.json")
			}
			os.Exit(0)
		case "/status":
			if exists(vtolvrpath + "\\patchman.json") {
				fmt.Println("Current Variant " + Status.InstalledName + " Object ID " + Status.InstalledObjectId + " Variant ID " + Status.InstalledVariantId + " Version ID " + Status.InstalledVersionId)
			} else {
				fmt.Println("patchman.json does not exist, game is likely unpatched, or user removed patchman.json")
			}
			os.Exit(0)
		default:
			vtolversion = os.Args[1]
		}
	} else if len(os.Args) > 2 {
		log.Fatal("Unrecognized argument: " + os.Args[1] + "\nValid examples:\npatchman.exe [game buildid override] \npatchman.exe 18407725\npatchman.exe version\n patchman.exe help\npatchman.exe patchstatus")
	}

	app = tview.NewApplication()
	app.EnableMouse(true)

	root = tview.NewPages()

	root.SetBorder(false).SetTitle("VTOL VR Patch Manager")

	buildInitialSelection()

	buildForm()

	runApp()
}

func runApp() {
	go keepAlive()

	if err := app.SetRoot(root, true).Run(); err != nil {
		log.Panic(err)
	}

}

func keepAlive() {
	for {
		select {
		case <-stop:
			return // Exits the goroutine
		default:
			time.Sleep(50 * time.Millisecond)
			app.Draw()
		}
	}
}
