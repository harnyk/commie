package git

import (
	"errors"
	"os/exec"

	"github.com/harnyk/commie/pkg/agent"
)

type CommitParams struct {
	Message string `mapstructure:"message"`
}

var Commit agent.TypedHandler[CommitParams, string] = func(params CommitParams) (string, error) {
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

func NewCommit() *agent.Tool {
	type H = agent.H

	tool := agent.NewTool().
		WithName("commit").
		WithDescription("Commits staged changes to the repository with a message").
		WithHandler(Commit.AcceptingMapOfAny()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"message": H{
						"type": "string",
						"description": "The commit message",
					},
				},
				"required": []string{"message"},
			},
		)

	return tool
}
