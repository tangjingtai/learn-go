package app

import (
	"github.com/astaxie/beego/validation"
	"log"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Printf(err.Key, err.Message)
	}

	return
}