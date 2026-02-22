package global

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"tawesoft.co.uk/go/dialog"
)

func UnpackDependencies() {
	CreateAndWriteProgramWorkingDirectory(PatchmanUnity, "patchman-unity.exe")
	CreateAndWriteProgramWorkingDirectory(ClassDataTpk, "classdata.tpk")
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
		FatalError(fmt.Errorf("bad status: %s", resp.Status))
	}

	Assure(io.Copy(outputFile, resp.Body))
}

func FatalError(err error) {

	logfile, err1 := os.OpenFile("./patchman.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		dialog.Alert("%s", "Could not create or write to log file: "+err1.Error()+"\n\n"+"Main Error: "+err.Error())
		os.Exit(1)

	}
	timestamp := time.Now().Truncate(time.Second).String()

	if _, err1 := logfile.WriteString(timestamp + ": " + err.Error() + "\n"); err1 != nil {
		dialog.Alert("%s", "Could not create or write to log file: "+err1.Error()+"\n\n"+"Main Error: "+err.Error())
	}

	dialog.Alert("%s", "Fatal Error: "+err.Error()+"\n\nError saved to log file patchman.log in current directory.")

	logfile.Close()

	os.Exit(1)

}

func WriteLog(logstring string) {

	logfile, err1 := os.OpenFile("./patchman.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		dialog.Alert("%s", "Could not create or write to log file: "+err1.Error()+"\n\n"+": "+logstring)

	}
	timestamp := time.Now().Truncate(time.Second).String()

	if _, err1 := logfile.WriteString(timestamp + ": " + logstring); err1 != nil {
		dialog.Alert("%s", "Could not create or write to log file: "+err1.Error()+"\n\n"+": "+logstring)
	}

	dialog.Alert(logstring)

	logfile.Close()

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

// Assure is a helper function to avoid boilerplate error handling.
func Assure[T any](v T, err error) T {
	if err != nil {
		FatalError(err)
	}
	return v
}

// AssureNoReturn is a helper function to avoid boilerplate error handling when a given functioning does not return a value.
func AssureNoReturn(err error) {
	if err != nil {
		FatalError(err)
	}
}

// Sanitizes file paths to prevent absolute paths or ../ usage.
func sanitizeFilePath(path string) (sanitizedPath string) {
	sanitizedPath = filepath.Clean(path)
	if filepath.IsAbs(path) || strings.HasPrefix(path, "..") {
		FatalError(errors.New("Cannot use absolute filepaths in copy argument."))
	}

	return
}

func ExitSuccess() {
	os.Exit(0)
}
