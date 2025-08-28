package cdm_framework

import (
	"os"

	"github.com/joho/godotenv"
)

type Logger struct {
	Level string
}

func NewLogger() *Logger {
	godotenv.Load()

	var logLevel string

	if os.Getenv("LOG_LEVEL") == "" {
		logLevel = "info"
	} else {
		logLevel = os.Getenv("LOG_LEVEL")
	}

	return &Logger{
		Level: logLevel,
	}
}

func (l *Logger) Debug(msg string) {
	if l.Level == "debug" {
		println(msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.Level == "debug" || l.Level == "info" {
		println(msg)
	}
}

func (l *Logger) Warn(msg string) {
	if l.Level == "debug" || l.Level == "info" || l.Level == "warn" {
		println(msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.Level == "debug" || l.Level == "info" || l.Level == "warn" || l.Level == "error" {
		println(msg)
	}
}
