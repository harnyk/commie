package git

import (
	"errors"

	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/gena"
)

type CommitParams struct {
	Message string `mapstructure:"message"`
}

type CommitHandler struct {
	commandRunner *shell.CommandRunner
}

func NewCommitHandler(commandRunner *shell.CommandRunner) gena.ToolHandler {
	return &CommitHandler{
		commandRunner: commandRunner,
	}
}

func (h *CommitHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *CommitHandler) execute(params CommitParams) (string, error) {
	if params.Message == "" {
		return "", errors.New("no commit message specified")
	}

	args := []string{"commit", "-m", params.Message}
	return h.commandRunner.Run("git", args...)
}

func NewCommit(commandRunner *shell.CommandRunner) *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("git_commit").
		WithDescription("Commits staged changes to the repository with a message").
		WithHandler(NewCommitHandler(commandRunner)).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"message": H{
						"type":        "string",
						"description": "The commit message",
					},
				},
				"required": []string{"message"},
			},
		)

	return tool
}
