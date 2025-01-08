package git

import (
	"os/exec"

	"github.com/harnyk/gena"
)

type GitStatusParams struct {
}

type GitStatusHandler struct{}

func NewGitStatusHandler() gena.ToolHandler {
	return &GitStatusHandler{}
}

func (h *GitStatusHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *GitStatusHandler) execute(params GitStatusParams) (string, error) {
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
		WithHandler(NewGitStatusHandler()).
		WithSchema(
			gena.H{
				"type":       "object",
				"properties": gena.H{},
				"required":   []string{},
			},
		)
}
