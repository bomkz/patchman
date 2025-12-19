package main

import (
	"log"

	"github.com/bomkz/patchman/index"
)

func buildForm() {

	err := index.BuildIndex()
	if err != nil {
		log.Panic()
	}

}
