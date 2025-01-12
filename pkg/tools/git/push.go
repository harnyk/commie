package git

import (
	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/gena"
)

type PushParams struct {
	Remote string `mapstructure:"remote"`
	Branch string `mapstructure:"branch"`
}

type PushHandler struct {
	commandRunner *shell.CommandRunner
}

func NewPushHandler(commandRunner *shell.CommandRunner) gena.ToolHandler {
	return &PushHandler{
		commandRunner: commandRunner,
	}
}

func (h *PushHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *PushHandler) execute(params PushParams) (string, error) {
	remote := "origin"
	branch := "main"

	if params.Remote != "" {
		remote = params.Remote
	}

	if params.Branch != "" {
		branch = params.Branch
	}

	args := []string{"push", remote, branch}
	output, err := h.commandRunner.Run("git", args...)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func NewPush(commandRunner *shell.CommandRunner) *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("push").
		WithDescription("Pushes commits to the remote repository").
		WithHandler(NewPushHandler(commandRunner)).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"remote": H{
						"type":        "string",
						"description": "The remote to push to (default: origin)",
					},
					"branch": H{
						"type":        "string",
						"description": "The branch to push (default: main)",
					},
				},
			},
		)

	return tool
}
