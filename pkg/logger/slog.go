package logger

import (
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

var _ Interface = (*SlogLogger)(nil)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(level, logDisplayType string, output io.Writer) *SlogLogger {
	var logLevel slog.Level

	switch strings.ToLower(level) {
	case LogLevelDebug:
		logLevel = slog.LevelDebug
	case LogLevelInfo:
		logLevel = slog.LevelInfo
	case LogLevelWarn:
		logLevel = slog.LevelWarn
	case LogLevelError:
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelDebug
	}

	handler := slog.HandlerOptions{AddSource: false, Level: logLevel}

	jsonHandler := handler.NewJSONHandler(output)
	textHandler := handler.NewTextHandler(output)

	jsonLog := slog.New(jsonHandler)
	textLog := slog.New(textHandler)

	if strings.ToLower(logDisplayType) == "json" {
		return &SlogLogger{logger: jsonLog}
	}

	return &SlogLogger{logger: textLog}
}

func (l *SlogLogger) With(args ...any) {
	l.logger = l.logger.With(args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *SlogLogger) Fatal(msg string, args ...any) {
	l.logger.Warn(msg, args...)
	os.Exit(1)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}
