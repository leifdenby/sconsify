package sconsify

import (
	"log"
	"os"
)

var logger *log.Logger
var file *os.File

func InitialiseLogger() {
	filename := GetLogFileLocation()
	if filename != "" {
		DeleteLogFile(filename)
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			logger = log.New(file, "", log.Lshortfile)
		}
	}
}

func CloseLogger() {
	file.Close()
}

func Debug(message string) {
	if logger != nil {
		logger.Println(message)
	}
}

func Debugf(format string, v ...interface{}) {
	if logger != nil {
		logger.Printf(format, v...)
	}
}
