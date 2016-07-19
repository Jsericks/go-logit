package logit

import (
	"log"
	"os"
	"strings"
)

var LogLevel int

func init() {
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))

	switch logLevel {
	case "debug":
		LogLevel = 0
	case "info":
		LogLevel = 1
	case "warn":
		LogLevel = 2
	case "error":
		LogLevel = 3
	case "fatal":
		LogLevel = 4
	default:
		LogLevel = 3
	}
}

func Debug(format string, v ...interface{}) {
	if LogLevel <= 0 {
		log.Printf("DEBUG:"+format, v...)
	}
}

func Info(format string, v ...interface{}) {
	if LogLevel <= 1 {
		log.Printf("INFO:"+format, v...)
	}
}

func Warn(format string, v ...interface{}) {
	if LogLevel <= 2 {
		log.Printf("WARN:"+format, v...)
	}
}

func Error(format string, v ...interface{}) {
	if LogLevel <= 3 {
		log.Printf("ERROR:"+format, v...)
	}
}

func Fatal(format string, v ...interface{}) {
	if LogLevel <= 4 {
		log.Printf("FATAL:"+format, v...)
	}
}
