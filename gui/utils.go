package gui

import "log"

func RaiseFatal(msg ...interface{}) {
	log.Fatal(msg...)
}