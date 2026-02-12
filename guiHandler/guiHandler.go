package guiHandler

import (
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/steamutils"
)

func InitGui() {
	sr, err := steamutils.NewSteamReader(steamutils.SteamReaderConfig{FormatSteamPath: true})
	if err != nil {
		global.SteamPath = "Not Installed or Detected."
	} else {
		global.SteamPath = sr.GetSteamPath()
		buildGameListSteam()
	}

	global.MainWindow.ShowAndRun()
}
