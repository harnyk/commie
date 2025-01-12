package shell

import (
	"fmt"
	"strings"

	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/gena"
)

const (
	LIMIT_LINES = 200
	LIMIT_BYTES = 4096
)

// ShellParams holds the parameters for the shell command.
type ShellParams struct {
	Command string `mapstructure:"command"`
}

// ShellHandler handles the execution of shell commands.
type ShellHandler struct {
	runner shell.CommandRunner
}

// NewShellHandler creates a new ShellHandler with the specified shell.
func NewShellHandler(
	runner shell.CommandRunner,
) gena.ToolHandler {
	return &ShellHandler{
		runner: runner,
	}
}

// Execute executes the shell command with the given parameters.
func (h *ShellHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

// execute runs the shell command and returns the output.
func (h *ShellHandler) execute(params ShellParams) (string, error) {
	combinedOutput, err := h.runner.RunString(params.Command)
	if err != nil {
		return "", err
	}

	limitedOutput, wasLimitedByLines, wasLimitedByBytes := limitOutput(combinedOutput, LIMIT_LINES, LIMIT_BYTES)

	if wasLimitedByLines {
		limitedOutput = fmt.Sprintf("%s\n(only last %d lines are shown)", limitedOutput, LIMIT_LINES)
	} else if wasLimitedByBytes {
		limitedOutput = fmt.Sprintf("%s\n(only last %d bytes are shown)", limitedOutput, LIMIT_BYTES)
	}

	return limitedOutput, nil
}

func limitOutput(output string, lineLimit int, byteLimit int) (result string, wasLimitedByLines bool, wasLimitedByBytes bool) {
	lines := strings.Split(output, "\n")
	if len(lines) > lineLimit {
		lines = lines[len(lines)-lineLimit:]
		wasLimitedByLines = true
		limitedOutput := strings.Join(lines, "\n")
		if len(limitedOutput) > byteLimit {
			limitedOutput = limitedOutput[len(limitedOutput)-byteLimit:]
			wasLimitedByLines = false
			wasLimitedByBytes = true
		}
		return limitedOutput, wasLimitedByLines, wasLimitedByBytes
	}
	limitedOutput := strings.Join(lines, "\n")
	if len(limitedOutput) > byteLimit {
		limitedOutput = limitedOutput[len(limitedOutput)-byteLimit:]
		wasLimitedByBytes = true
	}
	return limitedOutput, wasLimitedByLines, wasLimitedByBytes
}

// New creates a new shell tool.
func New(runner shell.CommandRunner) *gena.Tool {
	return gena.NewTool().
		WithName("shell").
		WithDescription("Executes an arbitrary shell command.").
		WithHandler(NewShellHandler(runner)).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"command": gena.H{
						"type":        "string",
						"description": "The shell command to execute.",
					},
				},
				"required": []string{"command"},
			},
		)
}
