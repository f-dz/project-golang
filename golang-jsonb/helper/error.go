package helper

import "log"

func CheckErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
