package ipc

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/bomkz/patchman/ipc/procman"
	"github.com/bomkz/steamutils"
)

func steamPath() {
	sr, err := steamutils.NewSteamReader(steamutils.SteamReaderConfig{})
	if err != nil {
		procman.Write("error", []byte(err.Error()))
		return
	}
	steamPath := sr.GetSteamPath()

	procman.Write("steampath", []byte(steamPath))
}
func gameByAppId(appID string) {
	//Section 1
	sr, err := steamutils.NewSteamReader(steamutils.SteamReaderConfig{})
	if err != nil {
		fatalErrorPriviledged(err)
	}

	//Section 2
	game, err := sr.GetInstalledAppByID(appID)
	if err != nil {
		procman.Write("notfound", []byte(err.Error()))
		return
	}

	//Section 3
	gameByte, err := json.Marshal(game)
	if err != nil {
		fatalErrorPriviledged(err)
	}

	procman.Write("gameappbyid", gameByte)
}

func fatalErrorPriviledged(err error) {
	procman.Write("error", []byte(err.Error()))
	logfile, _ := os.OpenFile("./patchman.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	timestamp := time.Now().Truncate(time.Second).String()

	logfile.WriteString(timestamp + ": " + err.Error() + "\n")
	logfile.Close()
	os.Exit(1)
}

func StartHelper() {
	messageHandler()
}

func patchAssets() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command(pwdDir+"\\patchman-unity.exe", "batchimportasset", pwdDir+".\\operations.json", pwdDir+".\\classdata.tpk")
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: 0x08000000,
		}
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			procman.Write("error", []byte(err.Error()))
		}

		if out.String() != "Done!" {
			procman.Write("error", out.Bytes())
		}
		procman.Write("success", nil)

	case "linux":
		cmd := exec.Command(pwdDir+"\\patchman-unity", "batchimportasset", pwdDir+".\\operations.json")
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			procman.Write("error", []byte(err.Error()))
		}

		if out.String() != "Done!" {
			procman.Write("error", out.Bytes())
		}

		procman.Write("success", nil)
	}

}
func renameGameWorkingDirectoryFile(fileName string) {
	tgt := fileName

	err := os.Rename(gwdDir+"\\.\\"+tgt, gwdDir+"\\.\\"+tgt+".orig")
	if err != nil {
		procman.Write("error", []byte(err.Error()))
	}
	err = os.Rename(gwdDir+"\\.\\"+tgt+".mod", gwdDir+"\\.\\"+tgt)
	if err != nil {
		procman.Write("error", []byte(err.Error()))
	}

	procman.Write("success", nil)
}

func patchBundles(CompressionType string) {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command(pwdDir+"\\patchman-unity.exe", "batchimportbundle", pwdDir+".\\operations.json", CompressionType)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: 0x08000000,
		}
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			procman.Write("error", []byte(err.Error()))
		}

		if out.String() != "Done!" {
			procman.Write("error", out.Bytes())
		}

		procman.Write("success", nil)

	default:
		cmd := exec.Command(pwdDir+"\\patchman-unity", "batchimportbundle", pwdDir+".\\operations.json", CompressionType)
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			procman.Write("error", []byte(err.Error()))
		}

		if out.String() != "Done!" {
			procman.Write("error", out.Bytes())
		}
		procman.Write("success", nil)

	}
}

// Copies file from patchRoot to gameRoot
func CopyFromProgramWorkingDirectory(fileName string, target string) {
	src := fileName
	dst := target

	// Open src file
	inputFile, err := os.Open(pwdDir + "\\.\\" + src)
	if err != nil {
		fatalErrorPriviledged(err)
	}
	defer inputFile.Close()

	// Create dst file and defer for closing
	outputFile, err := os.Create(gwdDir + "\\.\\" + dst)
	if err != nil {
		fatalErrorPriviledged(err)
	}
	defer outputFile.Close()

	// Copy contents from src to dst
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		fatalErrorPriviledged(err)
	}
}

type CopyStruct struct {
	FileName    string `json:"fileName"`
	Destination string `json:"destination"`
}

func copy(actionData []byte) {
	var copyData CopyStruct

	err := json.Unmarshal(actionData, &copyData)
	if err != nil {
		fatalErrorPriviledged(err)
	}

	CopyFromProgramWorkingDirectory(copyData.FileName, copyData.Destination)

	procman.Write("success", nil)

}

var pwdDir string

var gwdDir string

func checkExistsAtGwd(fileName string) bool {
	src := fileName
	_, err := os.Stat(gwdDir + "\\.\\" + src)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		procman.Write("error", []byte(err.Error()))
	}
	return !os.IsNotExist(err)
}

func existsAtGwd(fileName string) {
	exists := checkExistsAtGwd(fileName)
	if exists {
		procman.Write("true", nil)
	} else {
		procman.Write("false", nil)
	}

}

func messageHandler() {
	for {
		message := procman.Read()

		switch string(message.Type) {
		case "steampath":
			steamPath()
		case "gameappbyid":
			appId := string(message.Payload)
			gameByAppId(appId)
		case "patchbundles":
			compressionType := string(message.Payload)
			patchBundles(compressionType)
		case "patchassets":
			patchAssets()
		case "copy":
			copy(message.Payload)
		case "pwd":
			pwdDir = string(message.Payload)
		case "gwd":
			gwdDir = string(message.Payload)
		case "renamegameworkingdirectoryfile":
			renameGameWorkingDirectoryFile(string(message.Payload))
		case "existsatgwd":
			existsAtGwd(string(message.Payload))
		default:
			procman.Write("error", []byte("Error: Unrecognized command"+message.Type))
		}
	}
}
