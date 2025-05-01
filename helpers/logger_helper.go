package helpers

import (
	"fmt"
	"log"
	"runtime"
)

// CODE LINE INFO
func getCallerInfo() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", file, line)
}

func LogWithSeverity(severity string, message any) {
	// GET FILE NAME AND LINE NUMBER
	callerInfo := getCallerInfo()

	// CHECK MESSAGE DATA TYPE
	var logMessage string
	switch msg := message.(type) {
	case string:
		logMessage = msg
	case error:
		logMessage = fmt.Sprintf("Error: %s", msg.Error())
	default:
		logMessage = "Invalid log message type"
	}

	// SET LOG FLAGS
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// LOG MESSAGE
	log.Printf("[%s] %s: %s", severity, callerInfo, logMessage)
}
