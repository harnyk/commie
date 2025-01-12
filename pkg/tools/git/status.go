package git

import (
	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/gena"
)

type GitStatusParams struct {
}

type StatusHandler struct {
	commandRunner *shell.CommandRunner
}

func NewGitStatusHandler(commandRunner *shell.CommandRunner) gena.ToolHandler {
	return &StatusHandler{
		commandRunner: commandRunner,
	}
}

func (h *StatusHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *StatusHandler) execute(params GitStatusParams) (string, error) {
	args := []string{"status"}
	return h.commandRunner.Run("git", args...)
}

func NewStatus(commandRunner *shell.CommandRunner) *gena.Tool {
	return gena.NewTool().
		WithName("git_status").
		WithDescription("Returns the git status").
		WithHandler(NewGitStatusHandler(commandRunner)).
		WithSchema(
			gena.H{
				"type":       "object",
				"properties": gena.H{},
				"required":   []string{},
			},
		)
}
