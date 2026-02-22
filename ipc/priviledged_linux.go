//go:build linux

package ipc

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc/procman"
)

func patchBundles(CompressionType string) {
	os.Chmod(pwdDir+"/patchman-unity.exe", 0755)

	cmd := exec.Command(pwdDir+global.PathSeparator()+"patchman-unity.exe", "batchimportbundle", pwdDir+global.PathSeparator()+"operations.json", CompressionType)
	cmd.Dir = pwdDir
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

func patchAssets() {
	os.Chmod(pwdDir+global.PathSeparator()+"patchman-unity.exe", 0755)

	cmd := exec.Command(pwdDir+global.PathSeparator()+"patchman-unity.exe", "batchimportasset", pwdDir+global.PathSeparator()+"operations.json", pwdDir+global.PathSeparator()+"classdata.tpk")
	cmd.Dir = pwdDir
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
