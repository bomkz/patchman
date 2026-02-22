//go:build windows

package ipc

import (
	"bytes"
	"os/exec"
	"syscall"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc/procman"
)

func patchAssets() {

	cmd := exec.Command(pwdDir+global.PathSeparator()+"patchman-unity.exe", "batchimportasset", pwdDir+global.PathSeparator()+"operations.json", pwdDir+global.PathSeparator()+"classdata.tpk")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Dir = pwdDir

	err := cmd.Run()
	if err != nil {
		procman.Write("error", []byte(err.Error()))
	}

	if out.String() != "Done!" {
		procman.Write("error", out.Bytes())
	}
	procman.Write("success", nil)

}
func patchBundles(CompressionType string) {

	cmd := exec.Command(pwdDir+global.PathSeparator()+"patchman-unity.exe", "batchimportbundle", pwdDir+global.PathSeparator()+"operations.json", CompressionType)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
	var out bytes.Buffer
	cmd.Dir = pwdDir

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
