package shell

import (
	"os/exec"

	"github.com/harnyk/gena"
)

type ShellParams struct {
	Command string `mapstructure:"command"`
}

type ShellHandler struct {
}

func NewShellHandler() gena.ToolHandler {
	return &ShellHandler{}
}

func (h *ShellHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *ShellHandler) execute(params ShellParams) (string, error) {
	output, err := exec.Command("bash", "-c", params.Command).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func New() *gena.Tool {
	return gena.NewTool().
		WithName("shell").
		WithDescription("Executes an arbitrary shell command.").
		WithHandler(NewShellHandler()).
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
