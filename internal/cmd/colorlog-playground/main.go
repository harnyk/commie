package main

import (
	"log/slog"
	"os"

	"github.com/harnyk/commie/pkg/colorlog"
)

func main() {
	// Initialize the ColorConsoleHandler
	handler := colorlog.NewColorConsoleHandler(os.Stdout)

	// Create a logger with the handler
	logger := slog.New(handler)

	// Log messages at different levels
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// Log with attributes
	logger.Info("Logging with attributes", slog.String("key1", "value1"), slog.Int("key2", 42))
}
