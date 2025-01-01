package git

import (
	"os/exec"

	"github.com/harnyk/gena"
)

type GitStatusParams struct {
}

var GitStatusHandler gena.TypedHandler[GitStatusParams, string] = func(params GitStatusParams) (string, error) {
	output, err := exec.Command("git", "status").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func NewStatus() *gena.Tool {
	return gena.NewTool().
		WithName("gitStatus").
		WithDescription("Returns the git status").
		WithHandler(GitStatusHandler.AcceptingMapOfAny()).
		WithSchema(
			gena.H{
				"type":       "object",
				"properties": gena.H{},
				"required":   []string{},
			},
		)
}
