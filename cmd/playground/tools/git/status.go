package git

import (
	"os/exec"

	"github.com/harnyk/commie/pkg/agent"
)

type GitStatusParams struct {
}

var GitStatusHandler agent.TypedHandler[GitStatusParams, string] = func(params GitStatusParams) (string, error) {
	output, err := exec.Command("git", "status").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func NewStatus() *agent.Tool {
	return agent.NewTool().
		WithName("gitStatus").
		WithDescription("Returns the git status").
		WithHandler(GitStatusHandler.AcceptingMapOfAny()).
		WithSchema(
			agent.H{
				"type":       "object",
				"properties": agent.H{},
				"required":   []string{},
			},
		)
}
