// +build !android,!windows,!js

package raylib

import (
	"fmt"
	"os"
)

// TraceLog - Show trace log messages (INFO, WARNING, ERROR, DEBUG)
func TraceLog(msgType int, text string, v ...interface{}) {
	switch msgType {
	case LogInfo:
		fmt.Printf("INFO: "+text+"\n", v...)
	case LogWarning:
		fmt.Printf("WARNING: "+text+"\n", v...)
	case LogError:
		fmt.Printf("ERROR: "+text+"\n", v...)
		os.Exit(1)
	case LogDebug:
		if traceDebugMsgs {
			fmt.Printf("DEBUG: "+text+"\n", v...)
		}
	}
}

// HomeDir - Returns user home directory
func HomeDir() string {
	return os.Getenv("HOME")
}
