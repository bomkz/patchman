package ione

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bomkz/patchman/global"
	actionScript "github.com/bomkz/patchman/index/ione/actionscript"
	"github.com/bomkz/patchman/index/ione/actionscript/actionScriptOne"
	"github.com/bomkz/patchman/index/izero"
)

func install(filePath string) {
	if global.Status.InstalledVersion != 99 {
		switch global.Status.InstalledVersion {
		case 0:
			izero.ReadTaint()
			izero.SilentUninstall()
		case 1:
			actionScriptOne.Uninstall()
		}
	}

	err := unzip(filePath, global.Directory)
	if err != nil {
		global.FatalError(err)
	}

	defer cleanup()

	patchscript, err := os.ReadFile(global.Directory + "\\patchscript.json")
	if err != nil {
		global.FatalError(err)

	}

	actionScript.HandleActionScript(patchscript)

	os.RemoveAll(global.Directory)

	global.StopApp <- true
}

func cleanup() {
	for _, x := range cleanupQueue {
		os.Remove(x)
	}
	cleanupQueue = []string{}
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		os.MkdirAll(filepath.Dir(fpath), os.ModePerm)

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
