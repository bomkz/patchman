//go:build linux

package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bomkz/patchman/global"
)

func checkAdmin() bool {

	return os.Geteuid() == 0
}

func promptElevate() {
	exe, _ := os.Executable()
	exe, _ = filepath.EvalSymlinks(exe)

	global.Cmd = exec.Command("pkexec", exe)
	global.AssureNoReturn(global.Cmd.Start())
	global.AssureNoReturn(global.Cmd.Wait())
}
