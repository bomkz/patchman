//go:build windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bomkz/patchman/global"
)

func promptElevate() {
	exe := global.Assure(os.Executable())
	exe = global.Assure(filepath.EvalSymlinks(exe))

	psCmd := fmt.Sprintf("Start-Process -FilePath %q -ArgumentList %q -Verb RunAs -Wait", exe)
	global.Cmd = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", psCmd)

	global.AssureNoReturn(global.Cmd.Start())
	global.AssureNoReturn(global.Cmd.Start())
}

func checkAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	isadmin := false
	if err != nil {
		isadmin = false
	} else {
		isadmin = true
	}
	return isadmin
}
