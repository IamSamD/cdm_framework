package cdm_framework

import (
	"log/slog"
	"os"
)

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
