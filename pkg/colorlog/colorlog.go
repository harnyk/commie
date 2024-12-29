package colorlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

type GrayConsoleHandler struct {
	writer *os.File
}

func NewGrayConsoleHandler(writer *os.File) slog.Handler {
	return &GrayConsoleHandler{writer: writer}
}

func (h *GrayConsoleHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *GrayConsoleHandler) Handle(_ context.Context, record slog.Record) error {

	timeStr := record.Time.Format(time.RFC3339)
	levelStr := levelToColor(record.Level)
	msg := record.Message

	attrs := ""
	record.Attrs(func(attr slog.Attr) bool {
		attrs += fmt.Sprintf(" %s=%v", attr.Key, attr.Value)
		return true
	})

	fmt.Fprintf(h.writer, "%s [%s] %s%s\n", timeStr, levelStr, msg, attrs)
	return nil
}

func (h *GrayConsoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *GrayConsoleHandler) WithGroup(name string) slog.Handler {
	return h
}

func levelToColor(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return "\033[34mDEBUG\033[0m"
	case slog.LevelInfo:
		return "\033[32mINFO\033[0m"
	case slog.LevelWarn:
		return "\033[33mWARN\033[0m"
	case slog.LevelError:
		return "\033[31mERROR\033[0m"
	default:
		return "\033[90mUNKNOWN\033[0m"
	}
}
