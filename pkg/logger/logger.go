package logger

import (
	"famcache/domain/logger"
	"log/slog"
)

type Logger struct{}

func (l *Logger) Info(format string, args ...interface{}) {
	slog.Info(format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	slog.Debug(format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	slog.Error(format, args...)
}

func NewLogger() logger.Logger {
	return &Logger{}
}
