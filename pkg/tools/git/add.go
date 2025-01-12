package git

import (
	"errors"

	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/gena"
)

type AddParams struct {
	Files []string `mapstructure:"files"`
}

type Add struct {
	commandRunner *shell.CommandRunner
}

func NewAddHandler(commandRunner *shell.CommandRunner) gena.ToolHandler {
	return &Add{
		commandRunner: commandRunner,
	}
}

func (h *Add) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *Add) execute(params AddParams) (string, error) {
	if len(params.Files) == 0 {
		return "", errors.New("no files specified")
	}

	args := append([]string{"add"}, params.Files...)
	return h.commandRunner.Run("git", args...)
}

func NewAdd(commandRunner *shell.CommandRunner) *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("git_add").
		WithDescription("Adds files to the git staging area").
		WithHandler(NewAddHandler(commandRunner)).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"files": H{
						"type": "array",
						"items": H{
							"type": "string",
						},
						"description": "List of files to add to staging",
					},
				},
				"required": []string{"files"},
			},
		)

	return tool
}
