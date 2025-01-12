package shell

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

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

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("failed to get stdout pipe: %v", err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("failed to get stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start command: %v", err)
	}

	outputBuilder := &strings.Builder{}
	stdoutScanner := bufio.NewScanner(stdoutPipe)
	stderrScanner := bufio.NewScanner(stderrPipe)

	stdoutDone := make(chan bool)
	stderrDone := make(chan bool)

	go func() {
		for stdoutScanner.Scan() {
			outputBuilder.WriteString(stdoutScanner.Text() + "\n")
		}
		stdoutDone <- true
	}()

	go func() {
		for stderrScanner.Scan() {
			outputBuilder.WriteString(stderrScanner.Text() + "\n")
		}
		stderrDone <- true
	}()

	<-stdoutDone
	<-stderrDone

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, outputBuilder.String())
	}

	combinedOutput := outputBuilder.String()
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
func New() *gena.Tool {
	envContext, err := NewEnvironmentContext()
	if err != nil {
		return nil
	}

	return gena.NewTool().
		WithName("shell").
		WithDescription("Executes an arbitrary shell command.").
		WithHandler(NewShellHandler(envContext)).
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
