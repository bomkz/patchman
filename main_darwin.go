//go:build darwin

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/bomkz/patchman/global"
)

func checkAdmin() bool {
	return os.Geteuid() == 0
}

func promptElevate() {
	exe, err := os.Executable()
	if err != nil {
		global.FatalError(errors.New("cannot determine executable: " + err.Error()))
	}

	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		global.FatalError(errors.New("cannot resolve symlink: " + err.Error()))
	}

	script := fmt.Sprintf("do shell script %s with administrator privileges", strconv.Quote(exe))
	cmd := exec.Command("osascript", "-e", script)

	out, err := cmd.CombinedOutput()
	if err != nil {
		global.FatalError(errors.New("elevation failed: " + err.Error() + ": " + string(out)))
	}
}
