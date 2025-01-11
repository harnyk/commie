package shell

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"github.com/harnyk/gena"
)

const (
	LIMIT_LINES  = 200
	LIMIT_BYTES = 4096
)

// ShellParams holds the parameters for the shell command.
type ShellParams struct {
	Command               string `mapstructure:"command"`
	AskedUserConfirmation bool   `mapstructure:"checklistIHaveExplicitelyAskedUserConfirmation"`
}

// ShellHandler handles the execution of shell commands.
type ShellHandler struct {
	envContext EnvironmentContext
}

// NewShellHandler creates a new ShellHandler with the specified shell.
func NewShellHandler(envContext EnvironmentContext) gena.ToolHandler {
	return &ShellHandler{envContext: envContext}
}

// Execute executes the shell command with the given parameters.
func (h *ShellHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

// execute runs the shell command and returns the output.
func (h *ShellHandler) execute(params ShellParams) (string, error) {
	if !params.AskedUserConfirmation {
		return "", errors.New("you have'nt asked user's confirmation. Do it now!")
	}

	shell := h.envContext.Shell

	if shell == "" {
		shell = "/bin/sh"
	}

	args := []string{
		"-c",
		params.Command,
	}

	if h.envContext.IsWindowsStyleFlags {
		args = []string{
			"/c",
			params.Command,
		}
	}

	cmd := exec.Command(shell, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, error: %v", params.Command, err)
	}

	limitedOutput, wasLimitedByLines, wasLimitedByBytes := limitOutput(string(output), LIMIT_LINES, LIMIT_BYTES)

	if wasLimitedByLines {
		limitedOutput = fmt.Sprintf("%s\n(only last %d lines are shown)", limitedOutput, LIMIT_LINES)
	} else if wasLimitedByBytes {
		limitedOutput = fmt.Sprintf("%s\n(only last %d bytes are shown)", limitedOutput, LIMIT_BYTES)
	}

	return string(limitedOutput), nil
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
func New() *gena.Tool {
	envContext, err := NewEnvironmentContext()
	if err != nil {
		return nil
	}

	return gena.NewTool().
		WithName("shell").
		WithDescription("Executes an arbitrary shell command. It is very dangerous, you MUST always ask the user's confirmation before executing a shell command. For example: Assistant: I am going to run the following command in your shell:\n```shell\nifconfig\n```. Do you agree? Answer 'yes(y)' or 'no(n)'.\n").
		WithHandler(NewShellHandler(envContext)).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"checklistIHaveExplicitelyAskedUserConfirmation": gena.H{
						"type":        "boolean",
						"description": "I have explicitly asked the user's confirmation before executing this shell command",
					},
					"command": gena.H{
						"type":        "string",
						"description": "The shell command to execute.",
					},
				},
				"required": []string{"command"},
			},
		)
}