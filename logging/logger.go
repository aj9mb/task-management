package logging

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "log: ", log.Lshortfile)
)

func GetLogger() *log.Logger {
	return logger
}