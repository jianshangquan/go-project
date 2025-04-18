package logger

import (
	"fmt"
)

type Logger struct {}

func NewLogger() *Logger {
	println("New Logger created")
	return &Logger{}
}

func (l *Logger) Log(message string) {
	// Log the request details
	fmt.Println("Log: ", message)
}
