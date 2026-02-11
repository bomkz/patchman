package guihandler

import (
	"strings"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/steamutils"
)

func InitGui() {
	sr, err := steamutils.NewSteamReader(steamutils.SteamReaderConfig{FormatSteamPath: true})
	if err != nil {
		global.SteamPath = "Not Installed or Detected."
	} else {
		if strings.ToLower(global.OsName) == "windows" {

			global.SteamPath = sr.GetSteamPath()
			buildGameListWindowsSteam()
		}
	}

	global.MainWindow.ShowAndRun()
}
