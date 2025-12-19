package ione

import (
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/index/ione/actionscript/actionScriptOne"
	"github.com/bomkz/patchman/index/izero"
)

func install(filePath string) {
	if global.Status.InstalledVersion != 99 {
		if global.Status.InstalledVersion == 0 {
			izero.ReadTaint()
			izero.SilentUninstall()
		} else if global.Status.InstalledVersion == 1 {
			actionScriptOne.Uninstall()
		}
	}
}
