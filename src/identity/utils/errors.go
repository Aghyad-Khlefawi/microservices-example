package utils

import (
	"fmt"
	"log"
)

func LogFatalError(message string, err error) {
	wrapped:= LogError(message,err)
	panic(wrapped)
}

func LogFatal(message string) {
	err := fmt.Errorf(message)
	log.Print(err.Error())
	panic(err)
}

func LogError(message string, err error) error {
	wrapped := fmt.Errorf(message + " -- %w",err)
	log.Print(err)
	return wrapped
}
