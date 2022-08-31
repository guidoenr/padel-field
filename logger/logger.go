package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var Loginfo *log.Logger
var Logwarning *log.Logger
var Logerror *log.Logger

// init initialize the three logs
func init() {
	logFile, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	Loginfo = log.New(logFile, "[INFO]:  ", log.Ldate|log.Ltime|log.Lshortfile)
	Logerror = log.New(logFile, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	Logwarning = log.New(logFile, "[WARN]:  ", log.Ldate|log.Ltime|log.Lshortfile)

	// setting output to write both terminal and logs.log
	mw := io.MultiWriter(os.Stdout, logFile)
	Loginfo.SetOutput(mw)
	Logerror.SetOutput(mw)
	Logwarning.SetOutput(mw)
}
