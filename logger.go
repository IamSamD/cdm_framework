package cdm_framework

import (
	"log/slog"
	"os"
)

// Logger provides a standard text structured logger for use in cdm-checks
// The logger is based on slog and uses the slog.TextHandler
// This ensures a standard logging behaviour across all checks
// Log level is INFO by default
// When using Logger, log level can be set via the env var LOG_LEVEL
// Available options: DEBUG | WARN | ERROR
// Usage:
//
//	log := cdm_framework.Logger
//	log.Info("Some log message")
//
// If you want to be able to set the log level in code, for example from a cli flag use InitLogger
var Logger *slog.Logger

func init() {
	logLevel := os.Getenv("LOG_LEVEL")

	slogLogLevel := slog.LevelInfo

	switch logLevel {
	case "DEBUG":
		slogLogLevel = slog.LevelDebug
	case "WARN":
		slogLogLevel = slog.LevelWarn
	case "ERROR":
		slogLogLevel = slog.LevelError
	}

	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLogLevel,
	}))
}

// InitLogger allows for initiallisation of a cdm_framework.Logger where the log level is passed and an arguement
// Useful when you want to set the log level in code for example when setting log level from a cli flag
// The default log level is INFO and the default can be used by passing and empty string:
//
//	log := cdm_framework.InitLogger("")
//
// Otherwise pass your desired log level:
//
//	log := cdm_framework.InitLogger("DEBUG")
//
// Available options: DEBUG | WARN | ERROR
func InitLogger(logLevel string) *slog.Logger {
	var slogLogLevel slog.Level

	switch logLevel {
	case "DEBUG":
		slogLogLevel = slog.LevelDebug
	case "WARN":
		slogLogLevel = slog.LevelWarn
	case "ERROR":
		slogLogLevel = slog.LevelError
	default:
		slogLogLevel = slog.LevelInfo
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLogLevel,
	}))

	return logger
}
