package main

import (
	"errors"
	"fmt"
	"os"
)

func HandleError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg+"%s", err)
		os.Exit(-1)
	}
}

func IsFileExist(fileName string) bool {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
