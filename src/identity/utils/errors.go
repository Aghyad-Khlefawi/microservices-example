package utils

import (
	"fmt"
	"log"
)

func LogFatalError(message string, err error) {
	wrapped := fmt.Errorf(message + " -- %w",err)
	log.Print(wrapped)
	panic(wrapped)
}

func LogFatal(message string) {
	err := fmt.Errorf(message)
	log.Print(err)
	panic(err)
}
