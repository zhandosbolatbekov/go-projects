package pkg

import (
	"log/slog"
	"os"
)

type SlogLogger struct {
	logger *slog.Logger
}

type LoggerConfig struct {
	Level  slog.Level
	Format string // "json" or "text"
}

func NewLogger(cfg LoggerConfig) Logger {
	var handler slog.Handler
	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.Level})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.Level})
	}
	return &SlogLogger{
		logger: slog.New(handler),
	}
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SlogLogger) With(args ...any) Logger {
	return &SlogLogger{
		logger: l.logger.With(args...),
	}
}
