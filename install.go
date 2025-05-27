package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/bomkz/patchman/steamutils"
)

func install() {
	defer CleanDir()
	app.Stop()
	assertTargets()
	patched := true
	if Status.InstalledName == "Unpatched" {
		patched = false
	}
	if Status.InstalledName == "" {
		patched = false
	}
	if patched {
		SilentUninstall()
	}
	if !internet {

		if !exists("C:\\patchman\\" + StatusTarget.TargetUUID + "\\patch.zip") {
			CleanDir()
			log.Fatal("Error: C:\\patchman\\" + StatusTarget.TargetUUID + "\\patch.zip does not exist.\nPlease download the zip from " + StatusTarget.TargetVariantPatchURL + " on another device, and save it to C:\\patchman\\" + StatusTarget.TargetUUID + "\\patch.zip\nCreate any missing folders if necessary.")
		}
		if !exists("C:\\patchman\\zstd.exe") {
			CleanDir()
			log.Fatal("Error: C:\\patchman\\zstd.exe does not exist.\nPlease download the file from " + zstdURL + " on another device, and save it to C:\\patchman\\zstd.exe\nCreate any missing folders if necessary.")
		}
		err := copyFile(directory+"\\patch.zip", "C:\\patchman\\"+StatusTarget.TargetUUID+"\\patch.zip")
		if err != nil {
			CleanDir()
			log.Fatal(err)
		}
		err = copyFile(directory+"\\zstd.exe", "C:\\patchman\\zstd.exe")
		if err != nil {
			CleanDir()
			log.Fatal(err)
		}
		unzip(directory+"\\patch.zip", directory+"\\")
		PatchFiles()
		enablePatch()
		taintDirectory()
		InstallDone()
		return
	}

	downloadFile(directory+"\\patch.zip", StatusTarget.TargetVariantPatchURL)
	downloadFile(directory+"\\zstd.exe", zstdURL)
	unzip(directory+"\\patch.zip", directory+"\\")
	PatchFiles()
	enablePatch()
	taintDirectory()
	InstallDone()

}

func copyFile(dst string, src string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func SilentUninstall() {
	if !internet {
		if !exists("C:\\patchman\\" + Status.InstalledUUID + "\\unpatch.zip") {
			CleanDir()
			log.Fatal("Error: C:\\patchman\\" + Status.InstalledUUID + "\\unpatch.zip does not exist.\nPlease download the zip from " + Status.InstalledVariantUnpatchURL + " on another device, and save it to C:\\patchman\\" + StatusTarget.TargetUUID + "\\unpatch.zip\nCreate any missing folders if necessary.")
		}
		if !exists("C:\\patchman\\zstd.exe") {
			CleanDir()
			log.Fatal("Error: C:\\patchman\\zstd.exe does not exist.\nPlease download the file from " + zstdURL + " on another device, and save it to C:\\patchman\\zstd.exe\nCreate any missing folders if necessary.")
		}
		err := copyFile(directory+"\\unpatch.zip", "C:\\patchman\\"+Status.InstalledUUID+"\\unpatch.zip")
		if err != nil {
			CleanDir()
			log.Fatal(err)
		}
		err = copyFile(directory+"\\zstd.exe", "C:\\patchman\\zstd.exe")
		if err != nil {
			CleanDir()
			log.Fatal(err)
		}
		unzip(directory+"\\unpatch.zip", directory+"\\")
		UnpatchFiles()
		enablePatch()
		untaintDirectory()
		os.RemoveAll(directory)
		directory, err = os.MkdirTemp(".\\", "patchman-")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	downloadFile(directory+"\\unpatch.zip", Status.InstalledVariantUnpatchURL)
	downloadFile(directory+"\\zstd.exe", zstdURL)
	unzip(directory+"\\unpatch.zip", directory+"\\")
	UnpatchFiles()
	enablePatch()
	untaintDirectory()
	os.RemoveAll(directory)
	var err error
	directory, err = os.MkdirTemp(".\\", "patchman-")
	if err != nil {
		log.Fatal(err)
	}
}

func Uninstall() {
	app.Stop()
	time.Sleep(500 * time.Millisecond)
	patched := true
	if Status.InstalledName == "Unpatched" {
		patched = false
	}
	if Status.InstalledName == "" {
		patched = false
	}
	if !patched {
		fmt.Println("VTOL VR is not marked as patched. However, if it still is, please verify game files on Steam to fully uninstall.")
		os.Exit(1)
	}
	if !internet {
		if !exists("C:\\patchman\\" + Status.InstalledUUID + "\\unpatch.zip") {
			CleanDir()
			log.Fatal("Error: C:\\patchman\\" + Status.InstalledUUID + "\\unpatch.zip does not exist.\nPlease download the zip from " + Status.InstalledVariantUnpatchURL + " on another device, and save it to C:\\patchman\\" + StatusTarget.TargetUUID + "\\unpatch.zip\nCreate any missing folders if necessary.")
		}
		if !exists("C:\\patchman\\zstd.exe") {
			CleanDir()
			log.Fatal("Error: C:\\patchman\\zstd.exe does not exist.\nPlease download the file from " + zstdURL + " on another device, and save it to C:\\patchman\\zstd.exe\nCreate any missing folders if necessary.")
		}
		err := copyFile(directory+"\\unpatch.zip", "C:\\patchman\\"+Status.InstalledUUID+"\\unpatch.zip")
		if err != nil {
			CleanDir()
			log.Fatal(err)
		}
		err = copyFile(directory+"\\zstd.exe", "C:\\patchman\\zstd.exe")
		if err != nil {
			CleanDir()
			log.Fatal(err)
		}
		unzip(directory+"\\unpatch.zip", directory+"\\")
		UnpatchFiles()
		enablePatch()
		untaintDirectory()
		InstallDone()
		return
	}
	downloadFile(directory+"\\unpatch.zip", Status.InstalledVariantUnpatchURL)
	downloadFile(directory+"\\zstd.exe", zstdURL)
	unzip(directory+"\\unpatch.zip", directory+"\\")
	UnpatchFiles()
	enablePatch()
	untaintDirectory()
	InstallDone()
}

func InstallDone() {
	fmt.Println("VTOL VR patch attempted, please read the log for any uncaught error, and verify functionality in game.\n\n\nPress enter to continue...")
	fmt.Scanln()
}

func UninstallDone() {
	fmt.Println("VTOL VR unpatch attempted, please read the log for any uncaught error, and verify functionality in game.\nIf game is broken, verify game files on Steam to fix VTOL VR.\n\n\nPress enter to continue...")
	fmt.Scanln()
}

func readTaint() {
	vtolvrpath := findVtolPath()

	taint, err := os.ReadFile(vtolvrpath + "\\patchman.json")
	if err != nil {
		app.Stop()
		log.Fatal("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman.")
	}
	err = json.Unmarshal(taint, &Status)
	if err != nil {
		log.Fatal("Error Reading patchman.json. Please verify game files and delete the following file: " + vtolvrpath + "\\patchman.json then rerun patchman.")
	}
}

func enablePatch() {
	vtolvrpath := findVtolPath()
	if exists(vtolvrpath + "\\VTOLVR_Data\\resources.resource.mod") {
		os.Remove(vtolvrpath + "\\VTOLVR_Data\\resources.resource")
		os.Rename(vtolvrpath+"\\VTOLVR_Data\\resources.resource.mod", vtolvrpath+"\\VTOLVR_Data\\resources.resource")
	}
	if exists(vtolvrpath + "\\VTOLVR_Data\\resources.assets.mod") {
		os.Remove(vtolvrpath + "\\VTOLVR_Data\\resources.assets")
		os.Rename(vtolvrpath+"\\VTOLVR_Data\\resources.assets.mod", vtolvrpath+"\\VTOLVR_Data\\resources.assets")
	}
	if exists(vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS.mod") {
		os.Remove(vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS")
		os.Rename(vtolvrpath+"\\VTOLVR_Data\\resources.assets.resS.mod", vtolvrpath+"\\VTOLVR_Data\\resources.assets.resS")
	}
}

func untaintDirectory() {
	Status.InstalledName = "Unpatched"
	Status.InstalledObjectId = "N/A"
	Status.InstalledVersionId = vtolversion
	Status.InstalledVariantId = "N/A"
	Status.InstalledVariantPatchURL = "N/A"
	Status.InstalledVariantUnpatchURL = "N/A"
	Status.InstalledUUID = ""

	patchmanstatus, err := json.Marshal(Status)
	vtolvrpath := findVtolPath()
	if err != nil {
		log.Fatal(err)
	}
	if exists(vtolvrpath + "\\patchman.json") {
		os.Remove(vtolvrpath + "\\patchman.json")
	}
	os.WriteFile(vtolvrpath+"\\patchman.json", patchmanstatus, os.ModeAppend)
}

func taintDirectory() {
	Status.InstalledName = selection.Name
	Status.InstalledObjectId = StatusTarget.TargetObjectId
	Status.InstalledVersionId = vtolversion
	Status.InstalledVariantId = StatusTarget.TargetVariantId
	Status.InstalledVariantPatchURL = StatusTarget.TargetVariantPatchURL
	Status.InstalledVariantUnpatchURL = StatusTarget.TargetVariantUnpatchURL
	Status.InstalledUUID = StatusTarget.TargetUUID

	patchmanstatus, err := json.Marshal(Status)
	vtolvrpath := findVtolPath()
	if err != nil {
		log.Fatal(err)
	}
	if exists(vtolvrpath + "\\patchman.json") {
		os.Remove(vtolvrpath + "\\patchman.json")
	}
	os.WriteFile(vtolvrpath+"\\patchman.json", patchmanstatus, os.ModeAppend)
}
func CleanDir() {
	os.RemoveAll(directory)
}

func findVtolPath() string {

	steamPath, err := steamutils.GetSteamPath()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.ReadFile(steamPath + "\\steamapps\\libraryfolders.vdf")
	if err != nil {
		log.Fatal(err)
	}
	steamMap, err := steamutils.Unmarshal(f)
	if err != nil {
		log.Fatal(err)
	}
	dir, err := steamutils.FindGameLibraryPath(steamMap, "667970")
	if err != nil {
		log.Fatal(err)
	}
	return dir + "\\VTOL VR\\"

}

func PatchFiles() {
	vtolvrpath := findVtolPath()

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.resource' '" + directory + "\\resources.resource.patch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.resource.mod'"); err != nil {
		log.Fatal(err)
	}

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS' '" + directory + "\\resources.assets.resS.patch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS.mod'"); err != nil {
		log.Fatal(err)
	}

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.assets' '" + directory + "\\resources.assets.patch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.mod'"); err != nil {
		log.Fatal(err)
	}

}

func UnpatchFiles() {
	vtolvrpath := findVtolPath()

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.resource' '" + directory + "\\resources.resource.unpatch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.resource.mod'"); err != nil {
		log.Fatal(err)
	}

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS' '" + directory + "\\resources.assets.resS.unpatch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS.mod'"); err != nil {
		log.Fatal(err)
	}

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.assets' '" + directory + "\\resources.assets.unpatch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.mod'"); err != nil {
		log.Fatal(err)
	}

}

func zstd(arguments string) error {
	// Define the PowerShell command to decompress the file using zstd.exe
	cmd := exec.Command("powershell", "-Command", fmt.Sprint(`& {`+directory+`\zstd.exe `+arguments+`}`))

	// Run the command and capture any errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to decompress file: %v\n%s", err, output)
	} else {
		fmt.Println(string(output))
	}

	return nil
}

func unzip(source, dest string) error {
	read, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer read.Close()
	for _, file := range read.File {
		if file.Mode().IsDir() {
			continue
		}
		open, err := file.Open()
		if err != nil {
			return err
		}
		name := path.Join(dest, file.Name)
		os.MkdirAll(path.Dir(name), os.ModeDir)
		create, err := os.Create(name)
		if err != nil {
			return err
		}
		defer create.Close()
		create.ReadFrom(open)
	}
	return nil
}

func assertTargets() {
	for _, x := range index {
		if x.ObjectID == selection.ObjectID {
			for _, y := range x.Versions {
				if y.Version == vtolversion {
					if selection.Combination != nil {
						for _, z := range y.Content {
							shouldcontinue := false

							for _, u := range z.Items {
								for _, h := range selection.Combination {
									if u.ObjectID == h && !u.Available {
										shouldcontinue = true
									}
								}
							}
							if shouldcontinue {
								continue
							}
							StatusTarget.TargetVariantId = z.ObjectID
							StatusTarget.TargetObjectId = selection.ObjectID
							StatusTarget.TargetVersionId = vtolversion
							StatusTarget.TargetVariantPatchURL = z.PatchURL
							StatusTarget.TargetVariantUnpatchURL = z.UnpatchURL
							selection.VariantID = z.ObjectID
							StatusTarget.TargetUUID = z.UUID
						}
					} else {
						for _, z := range y.Content {
							shouldcontinue := false
							for _, u := range z.Items {
								if u.Available {
									shouldcontinue = true
								}
							}
							if shouldcontinue {
								continue
							}
							StatusTarget.TargetVariantId = z.ObjectID
							StatusTarget.TargetObjectId = selection.ObjectID
							StatusTarget.TargetVersionId = vtolversion
							StatusTarget.TargetVariantPatchURL = z.PatchURL
							StatusTarget.TargetVariantUnpatchURL = z.UnpatchURL
							selection.VariantID = z.ObjectID
							StatusTarget.TargetUUID = z.UUID
						}
					}
				}
			}
		}
	}
}
