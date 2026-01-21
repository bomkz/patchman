package global

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bomkz/patchman/steamutils"
	"github.com/inancgumus/screen"
)

func UnpackDependencies() {
	switch OsName {
	case "windows":
		CreateAndWriteProgramWorkingDirectory(PatchmanUnityExe, "patchman-unity.exe")
	case "linux":
		CreateAndWriteProgramWorkingDirectory(PatchmanUnityLinux, "patchman-unity")

	}
	CreateAndWriteProgramWorkingDirectory(ClassDataTpk, "classdata.tpk")
}

func ExitTview() {
	App.Stop()
	screen.Clear()
}

func ExitApp() {
	ExitTview()
	screen.Clear()
	os.Exit(0)
}

func ExitAppWithMessage(message string) {
	ExitTview()
	screen.Clear()
	fmt.Println(message)
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
	os.Exit(0)
}

func FatalError(err error) {
	ExitTview()
	fmt.Println(err)
	fmt.Println("Press Enter to exit")
	fmt.Scanln()

	ExitApp()

}

// Creates file at target path relative to patch root and writes byte array to it.
func CreateAndWriteProgramWorkingDirectory(fileByte []byte, target string) {
	dst := sanitizeFilePath(target)

	// Create file, defer for close.
	outputFile := Assure(os.Create(pwdDir + "\\.\\" + dst))
	defer outputFile.Close()

	// Write file contents
	Assure(outputFile.Write(fileByte))
}

func CreateWorkingDirectories(gameDirectory string) {

	pwdDir = Assure(os.MkdirTemp(".\\", "patchman-"))

	Directory = pwdDir

	gwdDir = gameDirectory
}

// Copies file from patchRoot to gameRoot
func CopyFromProgramWorkingDirectory(fileName string, target string) {
	src := sanitizeFilePath(fileName)
	dst := sanitizeFilePath(target)

	// Open src file
	inputFile := Assure(os.Open(pwdDir + "\\.\\" + src))
	defer inputFile.Close()

	// Create dst file and defer for closing
	outputFile := Assure(os.Create(gwdDir + "\\.\\" + dst))
	defer outputFile.Close()

	// Copy contents from src to dst
	Assure(io.Copy(outputFile, inputFile))
}

// Copies file from patchRoot to gameRoot
func CopyToProgramWorkingDirectory(fileName string, target string) {
	dst := sanitizeFilePath(target)

	// Open src file
	inputFile := Assure(os.Open(fileName))
	defer inputFile.Close()

	// Create dst file and defer for closing
	outputFile := Assure(os.Create(pwdDir + "\\.\\" + dst))
	defer outputFile.Close()

	// Copy contents from src to dst
	Assure(io.Copy(outputFile, inputFile))
}

// Deletes file from gwd
func DeleteFromGameWorkingDirectory(target string) {
	tgt := sanitizeFilePath(target)
	AssureNoReturn(os.Remove(gwdDir + "\\.\\" + tgt))
}

// Cleans up temporary pwd
func CleanProgramWorkingDirectory() {
	AssureNoReturn(os.RemoveAll(Directory))
}

// Checks if file exists at given path
func ExistsAtPwd(fileName string) bool {
	src := sanitizeFilePath(fileName)
	_, err := os.Stat(pwdDir + "\\.\\" + src)
	return !os.IsNotExist(err)
}

// Checks if file exists at given path
func ExistsAtGwd(fileName string) bool {
	src := sanitizeFilePath(fileName)
	_, err := os.Stat(gwdDir + "\\.\\" + src)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return !os.IsNotExist(err)
}

func GetGwd() string {
	return gwdDir
}

// Downloads file from URL to given path in pwd
func DownloadFileToProgramWorkingDirectory(filePath, url string) {
	dst := sanitizeFilePath(filePath)
	outputFile := Assure(os.Create(pwdDir + "\\.\\" + dst))
	defer outputFile.Close()

	resp := Assure(http.Get(url))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("bad status: %s", resp.Status))
	}

	Assure(io.Copy(outputFile, resp.Body))
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func RenameGameWorkingDirectoryFile(fileName string) {
	tgt := sanitizeFilePath(fileName)

	AssureNoReturn(os.Rename(gwdDir+"\\.\\"+tgt, gwdDir+"\\.\\"+tgt+".orig"))
	AssureNoReturn(os.Rename(gwdDir+"\\.\\"+tgt+".mod", gwdDir+"\\.\\"+tgt))

}

// Unzips given zipfile into pwd root
func UnzipIntoProgramWorkingDirectory(zipfile string) {
	r := Assure(zip.OpenReader(pwdDir + "\\.\\" + zipfile))
	defer r.Close()

	for _, f := range r.File {
		outFile := Assure(os.OpenFile(pwdDir+"\\.\\"+f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode()))

		rc := Assure(f.Open())

		Assure(io.Copy(outFile, rc))
		outFile.Close()
		rc.Close()
	}

}

func InitSteamReader() (err error) {
	SteamReader, err = steamutils.NewSteamReader(steamutils.SteamReaderConfig{})
	if err != nil {
		return
	}
	return
}

// Assure is a helper function to avoid boilerplate error handling.
func Assure[T any](v T, err error) T {
	if err != nil {
		CleanProgramWorkingDirectory()

		panic(err) // fail fast on critical fault
	}
	return v
}

// AssureNoReturn is a helper function to avoid boilerplate error handling when a given functioning does not return a value.
func AssureNoReturn(err error) {
	if err != nil {
		CleanProgramWorkingDirectory()

		panic(err) // fail fast on critical fault
	}
}

// Sanitizes file paths to prevent absolute paths or ../ usage.
func sanitizeFilePath(path string) (sanitizedPath string) {
	sanitizedPath = filepath.Clean(path)
	if filepath.IsAbs(path) || strings.HasPrefix(path, "..") {
		panic(errors.New("Cannot use absolute filepaths in copy argument."))
	}

	return
}
