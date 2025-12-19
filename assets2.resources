package main

import (
	"os"

	"github.com/bomkz/patchman/global"
)

func createDir() error {
	var err error
	global.Directory, err = os.MkdirTemp(".\\", "patchman-")
	if err != nil {
		return err
	}
	return nil
}
