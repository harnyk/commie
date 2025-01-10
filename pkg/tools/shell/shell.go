package shell

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/harnyk/gena"
)

// ShellParams holds the parameters for the shell command.
type ShellParams struct {
	Command               string `mapstructure:"command"`
	AskedUserConfirmation bool   `mapstructure:"checklistIHaveExplicitelyAskedUserConfirmation"`
}

// ShellHandler handles the execution of shell commands.
type ShellHandler struct {
	Shell string
}

// NewShellHandler creates a new ShellHandler with the specified shell.
func NewShellHandler(shell string) gena.ToolHandler {
	return &ShellHandler{Shell: shell}
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

	var cmd *exec.Cmd

	switch h.Shell {
	case "powershell":
		cmd = exec.Command("powershell", "-Command", params.Command)
	case "bash":
		cmd = exec.Command("bash", "-c", params.Command)
	default:
		return "", fmt.Errorf("unsupported shell: %s", h.Shell)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, error: %v", params.Command, err)
	}
	return string(output), nil
}

// New creates a new shell tool.
func New() *gena.Tool {
	shell := "bash"
	if runtime.GOOS == "windows" {
		shell = "powershell"
	}

	return gena.NewTool().
		WithName("shell").
		WithDescription("Executes an arbitrary shell command. It is very dangerous, you MUST always ask the user's confirmation before executing a shell command. For example: Assistant: I am going to run the following command in your shell:\n```shell\nifconfig\n```. Do you agree? Answer 'yes(y)' or 'no(n)'.").
		WithHandler(NewShellHandler(shell)).
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
