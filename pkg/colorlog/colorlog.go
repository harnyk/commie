package colorlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/mgutz/ansi"
)

type ColorConsoleHandler struct {
	writer *os.File
}

func NewColorConsoleHandler(writer *os.File) slog.Handler {
	return &ColorConsoleHandler{writer: writer}
}

func (h *ColorConsoleHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

// var clDimGreen = ansi.ColorFunc("green+d")
var clDimBlue = ansi.ColorFunc("blue+d")
var clDim = ansi.ColorFunc("reset+d")

var clDebug = ansi.ColorFunc("cyan")
var clInfo = ansi.ColorFunc("green")
var clWarn = ansi.ColorFunc("yellow")
var clError = ansi.ColorFunc("red")
var clUnknown = ansi.ColorFunc("magenta")

func (h *ColorConsoleHandler) Handle(_ context.Context, record slog.Record) error {

	timeStr := record.Time.Format(time.RFC3339)
	levelStr := levelToColor(record.Level)
	msg := record.Message

	attrs := ""
	record.Attrs(func(attr slog.Attr) bool {
		attrs += fmt.Sprintf(" %s=%v", clDimBlue(attr.Key), clDim(attr.Value.String()))
		return true
	})

	fmt.Fprintf(h.writer, "%s [%s] %s%s\n", clDim(timeStr), levelStr, msg, attrs)
	return nil
}

func (h *ColorConsoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *ColorConsoleHandler) WithGroup(name string) slog.Handler {
	return h
}

func levelToColor(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return clDebug("DEBUG")
	case slog.LevelInfo:
		return clInfo("INFO")
	case slog.LevelWarn:
		return clWarn("WARN")
	case slog.LevelError:
		return clError("ERROR")
	default:
		return clUnknown("UNKNOWN")
	}
}
