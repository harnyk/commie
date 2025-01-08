package git

import (
	"errors"
	"os/exec"

	"github.com/harnyk/gena"
)

type CommitParams struct {
	Message string `mapstructure:"message"`
}

type CommitHandler struct{}

func NewCommitHandler() gena.ToolHandler {
	return &CommitHandler{}
}

func (h *CommitHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *CommitHandler) execute(params CommitParams) (string, error) {
	if params.Message == "" {
		return "", errors.New("no commit message specified")
	}

	cmd := exec.Command("git", "commit", "-m", params.Message)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}

func NewCommit() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("commit").
		WithDescription("Commits staged changes to the repository with a message").
		WithHandler(NewCommitHandler()).
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
