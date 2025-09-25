package main

import (
	"log"

	"github.com/bomkz/patchman/index"
)

func buildForm() {

	form, err := index.BuildIndex()
	if err != nil {
		log.Panic()
	}
	root.AddAndSwitchToPage("installform", form, true)
}
