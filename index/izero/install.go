package izero

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

	"github.com/bomkz/patchman/global"
)

func install() {
	if !global.Exists("C:\\patchman\\patchman-unity.exe") && !global.Internet {

	}
	defer global.CleanDir()
	global.ExitTview()
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
	if !global.Internet {

		if !global.Exists("C:\\patchman\\" + StatusTarget.TargetUUID + "\\patch.zip") {
			global.CleanDir()
			log.Fatal("Error: C:\\patchman\\" + StatusTarget.TargetUUID + "\\patch.zip does not exist.\nPlease download the zip from " + StatusTarget.TargetVariantPatchURL + " on another device, and save it to C:\\patchman\\" + StatusTarget.TargetUUID + "\\patch.zip\nCreate any missing folders if necessary.")
		}
		if !global.Exists("C:\\patchman\\zstd.exe") {
			global.CleanDir()
			log.Fatal("Error: C:\\patchman\\zstd.exe does not exist.\nPlease download the file from " + zstdURL + " on another device, and save it to C:\\patchman\\zstd.exe\nCreate any missing folders if necessary.")
		}
		err := copyFile(global.Directory+"\\patch.zip", "C:\\patchman\\"+StatusTarget.TargetUUID+"\\patch.zip")
		if err != nil {
			global.CleanDir()
			global.FatalError(err)

		}
		err = copyFile(global.Directory+"\\zstd.exe", "C:\\patchman\\zstd.exe")
		if err != nil {
			global.CleanDir()
			global.FatalError(err)

		}
		unzip(global.Directory+"\\patch.zip", global.Directory+"\\")
		patchFiles()
		enablePatch()
		taintDirectory()
		installDone()
		return
	}

	global.DownloadFile(global.Directory+"\\patch.zip", StatusTarget.TargetVariantPatchURL)
	global.DownloadFile(global.Directory+"\\zstd.exe", zstdURL)
	unzip(global.Directory+"\\patch.zip", global.Directory+"\\")
	patchFiles()
	enablePatch()
	taintDirectory()
	installDone()

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
	if !global.Internet {
		if !global.Exists("C:\\patchman\\" + Status.InstalledUUID + "\\unpatch.zip") {
			global.CleanDir()
			log.Fatal("Error: C:\\patchman\\" + Status.InstalledUUID + "\\unpatch.zip does not exist.\nPlease download the zip from " + Status.InstalledVariantUnpatchURL + " on another device, and save it to C:\\patchman\\" + StatusTarget.TargetUUID + "\\unpatch.zip\nCreate any missing folders if necessary.")
		}
		if !global.Exists("C:\\patchman\\zstd.exe") {
			global.CleanDir()
			log.Fatal("Error: C:\\patchman\\zstd.exe does not exist.\nPlease download the file from " + zstdURL + " on another device, and save it to C:\\patchman\\zstd.exe\nCreate any missing folders if necessary.")
		}
		err := copyFile(global.Directory+"\\unpatch.zip", "C:\\patchman\\"+Status.InstalledUUID+"\\unpatch.zip")
		if err != nil {
			global.CleanDir()
			global.FatalError(err)

		}
		err = copyFile(global.Directory+"\\zstd.exe", "C:\\patchman\\zstd.exe")
		if err != nil {
			global.CleanDir()
			global.FatalError(err)

		}
		unzip(global.Directory+"\\unpatch.zip", global.Directory+"\\")
		unpatchFiles()
		enablePatch()
		untaintDirectory()
		os.RemoveAll(global.Directory)
		global.Directory, err = os.MkdirTemp(".\\", "patchman-")
		if err != nil {
			global.FatalError(err)

		}
		return
	}
	global.DownloadFile(global.Directory+"\\unpatch.zip", Status.InstalledVariantUnpatchURL)
	global.DownloadFile(global.Directory+"\\zstd.exe", zstdURL)
	unzip(global.Directory+"\\unpatch.zip", global.Directory+"\\")
	unpatchFiles()
	enablePatch()
	untaintDirectory()
	os.RemoveAll(global.Directory)
	var err error
	global.Directory, err = os.MkdirTemp(".\\", "patchman-")
	if err != nil {
		global.FatalError(err)

	}
}

func uninstall() {
	global.ExitTview()
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
	if !global.Internet {
		if !global.Exists("C:\\patchman\\" + Status.InstalledUUID + "\\unpatch.zip") {
			global.CleanDir()
			log.Fatal("Error: C:\\patchman\\" + Status.InstalledUUID + "\\unpatch.zip does not exist.\nPlease download the zip from " + Status.InstalledVariantUnpatchURL + " on another device, and save it to C:\\patchman\\" + StatusTarget.TargetUUID + "\\unpatch.zip\nCreate any missing folders if necessary.")
		}
		if !global.Exists("C:\\patchman\\zstd.exe") {
			global.CleanDir()
			log.Fatal("Error: C:\\patchman\\zstd.exe does not exist.\nPlease download the file from " + zstdURL + " on another device, and save it to C:\\patchman\\zstd.exe\nCreate any missing folders if necessary.")
		}
		err := copyFile(global.Directory+"\\unpatch.zip", "C:\\patchman\\"+Status.InstalledUUID+"\\unpatch.zip")
		if err != nil {
			global.CleanDir()
			global.FatalError(err)

		}
		err = copyFile(global.Directory+"\\zstd.exe", "C:\\patchman\\zstd.exe")
		if err != nil {
			global.CleanDir()
			global.FatalError(err)

		}
		unzip(global.Directory+"\\unpatch.zip", global.Directory+"\\")
		unpatchFiles()
		enablePatch()
		untaintDirectory()
		uninstallDone()
		return
	}
	global.DownloadFile(global.Directory+"\\unpatch.zip", Status.InstalledVariantUnpatchURL)
	global.DownloadFile(global.Directory+"\\zstd.exe", zstdURL)
	unzip(global.Directory+"\\unpatch.zip", global.Directory+"\\")
	unpatchFiles()
	enablePatch()
	untaintDirectory()
	uninstallDone()
}

func installDone() {
	fmt.Println("VTOL VR patch attempted, please read the log for any uncaught error, and verify functionality in game.\n\n\nPress enter to continue...")
	fmt.Scanln()
}

func uninstallDone() {
	fmt.Println("VTOL VR unpatch attempted, please read the log for any uncaught error, and verify functionality in game.\nIf game is broken, verify game files on Steam to fix VTOL VR.\n\n\nPress enter to continue...")
	fmt.Scanln()
}

func enablePatch() {
	vtolvrpath := global.FindVtolPath()
	if global.Exists(vtolvrpath + "\\VTOLVR_Data\\resources.resource.mod") {
		os.Remove(vtolvrpath + "\\VTOLVR_Data\\resources.resource")
		os.Rename(vtolvrpath+"\\VTOLVR_Data\\resources.resource.mod", vtolvrpath+"\\VTOLVR_Data\\resources.resource")
	}
	if global.Exists(vtolvrpath + "\\VTOLVR_Data\\resources.assets.mod") {
		os.Remove(vtolvrpath + "\\VTOLVR_Data\\resources.assets")
		os.Rename(vtolvrpath+"\\VTOLVR_Data\\resources.assets.mod", vtolvrpath+"\\VTOLVR_Data\\resources.assets")
	}
	if global.Exists(vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS.mod") {
		os.Remove(vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS")
		os.Rename(vtolvrpath+"\\VTOLVR_Data\\resources.assets.resS.mod", vtolvrpath+"\\VTOLVR_Data\\resources.assets.resS")
	}
}

func untaintDirectory() {
	Status.InstalledName = "Unpatched"
	Status.InstalledObjectId = "N/A"
	Status.InstalledVersionId = "N/A"
	Status.InstalledVariantId = "N/A"
	Status.InstalledVariantPatchURL = "N/A"
	Status.InstalledVariantUnpatchURL = "N/A"
	Status.InstalledUUID = ""
	Status.InstalledVersion = 0

	patchmanstatus, err := json.Marshal(Status)
	vtolvrpath := global.FindVtolPath()
	if err != nil {
		global.FatalError(err)

	}
	if global.Exists(vtolvrpath + "\\patchman.json") {
		os.Remove(vtolvrpath + "\\patchman.json")
	}
	os.WriteFile(vtolvrpath+"\\patchman.json", patchmanstatus, os.ModeAppend)
}

func taintDirectory() {
	Status.InstalledName = selection.Name
	Status.InstalledObjectId = StatusTarget.TargetObjectId
	Status.InstalledVersionId = global.VtolVersion
	Status.InstalledVariantId = StatusTarget.TargetVariantId
	Status.InstalledVariantPatchURL = StatusTarget.TargetVariantPatchURL
	Status.InstalledVariantUnpatchURL = StatusTarget.TargetVariantUnpatchURL
	Status.InstalledUUID = StatusTarget.TargetUUID
	Status.InstalledVersion = 0

	patchmanstatus, err := json.Marshal(Status)
	vtolvrpath := global.FindVtolPath()
	if err != nil {
		global.FatalError(err)

	}
	if global.Exists(vtolvrpath + "\\patchman.json") {
		os.Remove(vtolvrpath + "\\patchman.json")
	}
	os.WriteFile(vtolvrpath+"\\patchman.json", patchmanstatus, os.ModeAppend)
}

func patchFiles() {
	vtolvrpath := global.FindVtolPath()

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.resource' '" + global.Directory + "\\resources.resource.patch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.resource.mod'"); err != nil {
		global.FatalError(err)

	}

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS' '" + global.Directory + "\\resources.assets.resS.patch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS.mod'"); err != nil {
		global.FatalError(err)

	}

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.assets' '" + global.Directory + "\\resources.assets.patch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.mod'"); err != nil {
		global.FatalError(err)

	}

}

func unpatchFiles() {
	vtolvrpath := global.FindVtolPath()

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.resource' '" + global.Directory + "\\resources.resource.unpatch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.resource.mod'"); err != nil {
		global.FatalError(err)

	}

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS' '" + global.Directory + "\\resources.assets.resS.unpatch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.resS.mod'"); err != nil {
		global.FatalError(err)

	}

	if err := zstd("-d -f --long=31 --patch-from='" + vtolvrpath + "\\VTOLVR_Data\\resources.assets' '" + global.Directory + "\\resources.assets.unpatch' -o '" + vtolvrpath + "\\VTOLVR_Data\\resources.assets.mod'"); err != nil {
		global.FatalError(err)

	}

}

func zstd(arguments string) error {
	// Define the PowerShell command to decompress the file using zstd.exe
	cmd := exec.Command("powershell", "-Command", fmt.Sprint(`& {`+global.Directory+`\zstd.exe `+arguments+`}`))

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
				if y.Version == global.VtolVersion {
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
							StatusTarget.TargetVersionId = global.VtolVersion
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
							StatusTarget.TargetVersionId = global.VtolVersion
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
