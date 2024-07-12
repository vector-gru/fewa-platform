package logging

import (
	"log"
)

// Logger is a simple logging interface
type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

// SimpleLogger is a basic implementation of the Logger interface
type SimpleLogger struct{}

// NewLogger creates a new SimpleLogger
func NewLogger() Logger {
	return &SimpleLogger{}
}

// Error logs an error message
func (l *SimpleLogger) Error(args ...interface{}) {
	log.Println("ERROR:", args)
}

// Info logs an info message
func (l *SimpleLogger) Info(args ...interface{}) {
	log.Println("INFO:", args)
}
