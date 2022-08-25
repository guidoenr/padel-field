package logger

import (
	"fmt"
	"log"
	"os"
)

var Loginfo *log.Logger
var Logwarning *log.Logger
var Logerror *log.Logger

func init() {
	logFile, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	Loginfo = log.New(logFile, "[INFO]:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Logerror = log.New(logFile, "[ERROR]:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Logwarning = log.New(logFile, "[WARNING]:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
