package utils

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func InitLoggerWithName(name string, level string) *logrus.Entry {
	// Create a new logger instance
	myLogger := logrus.New()
	myLogger.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: true})

	// Set level
	l := strings.ToUpper(level)
	switch l {
	case "DEBUG":
		myLogger.SetLevel(logrus.DebugLevel)
	case "INFO":
		myLogger.SetLevel(logrus.InfoLevel)
	case "ERROR":
		myLogger.SetLevel(logrus.ErrorLevel)
	default:
		myLogger.SetLevel(logrus.InfoLevel)
	}

	// Return an entry with the name field attached
	return myLogger.WithField("name", name)
}

// Need to implement logging methods so that client is decoupled completely with the logging library used
// TODO: Logging methods
