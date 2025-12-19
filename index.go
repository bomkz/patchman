package main

import (
	"log"

	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/index"
)

func buildForm() {

	form, err := index.BuildIndex()
	if err != nil {
		log.Panic()
	}
	global.Root.AddAndSwitchToPage("installform", form, true)
}
