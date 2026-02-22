package ipc

import (
	"encoding/json"
	"errors"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc/procman"
	"github.com/bomkz/steamutils"
)

func SteamPath() (steamPath string, found bool) {

	procman.Write("steampath", nil)
	msg := procman.Read()

	switch msg.Type {
	case "error":
		global.WriteLog(string(msg.Payload))
		return
	case "steampath":
		return string(msg.Payload), true
	default:
		global.FatalError(errors.New("Unrecognized response:" + msg.Type))
		return
	}
}

func GameAppById(appId string) (game steamutils.InstalledApp, found bool) {
	procman.Write("gameappbyid", []byte(appId))
	msg := procman.Read()

	switch msg.Type {
	case "error":
		global.FatalError(errors.New(string(msg.Payload)))
		return
	case "notfound":
		global.WriteLog(string(msg.Payload))
		return
	case "gameappbyid":
		json.Unmarshal(msg.Payload, &game)
		found = true
		return
	default:
		global.FatalError(errors.New("Unrecognized response: " + msg.Type))
		return

	}
}

func PatchBundles(compressionType string) {
	procman.Write("patchbundles", []byte(compressionType))
	msg := procman.Read()
	switch msg.Type {
	case "error":
		global.FatalError(errors.New(string(msg.Payload)))
	case "success":
		return
	default:
		global.FatalError(errors.New("Unrecognized Type: " + msg.Type))
	}
}

func Pwd(pwd string) {
	procman.Write("pwd", []byte(pwd))
}

func Gwd(gwd string) {
	procman.Write("gwd", []byte(gwd))
}

func Copy(actionData json.RawMessage) {
	procman.Write("copy", actionData)
	msg := procman.Read()

	switch msg.Type {
	case "error":
		global.FatalError(errors.New(string(msg.Payload)))
	case "success":
		return
	default:
		global.FatalError(errors.New("Unrecognized Type: " + msg.Type))
	}
}

func RenameGameWorkingDirectoryFile(fileName string) {
	procman.Write("renamegameworkingdirectoryfile", []byte(fileName))
	msg := procman.Read()
	switch msg.Type {
	case "error":
		global.FatalError(errors.New(string(msg.Payload)))
	case "success":
		return
	default:
		global.FatalError(errors.New("Unrecognized Type: " + msg.Type))
	}
}
func ExistsAtGwd(fileName string) bool {
	procman.Write("existsatgwd", []byte(fileName))
	msg := procman.Read()
	switch msg.Type {
	case "error":
		global.FatalError(errors.New(string(msg.Payload)))
	case "true":
		return true
	case "false":
		return false
	default:
		global.FatalError(errors.New("Unrecognized Type: " + msg.Type))
	}
	return false
}

func PatchAssets() {
	procman.Write("patchassets", nil)
	msg := procman.Read()
	switch msg.Type {
	case "error":
		global.FatalError(errors.New(string(msg.Payload)))
	case "success":
		return
	default:
		global.FatalError(errors.New("Unrecognized Type: " + msg.Type))
	}
}
