package cdm_framework

import "log/slog"

type Config map[string]string

type RunFunc func(Config) error

var Logger *slog.Logger
