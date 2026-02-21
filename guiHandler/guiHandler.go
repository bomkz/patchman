package guiHandler

import (
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc"
)

func InitGui() {
	global.SteamPath = ipc.SteamPath()

	buildGameListSteam()

	global.MainWindow.ShowAndRun()
}
