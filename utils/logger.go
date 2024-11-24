package utils

import (
	"log/slog"
	"os"
)

func LoggerSetting() *slog.Logger {
	opt := slog.HandlerOptions{
		AddSource: true,
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opt))
	return logger
}