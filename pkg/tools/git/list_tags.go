package git

import (
	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/gena"
)

type ListTagsParams struct {
}

type ListTagsHandler struct {
	commandRunner *shell.CommandRunner
}

func NewListTagsHandler(commandRunner *shell.CommandRunner) gena.ToolHandler {
	return &ListTagsHandler{
		commandRunner: commandRunner,
	}
}

func (h *ListTagsHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *ListTagsHandler) execute(params ListTagsParams) (string, error) {
	args := []string{"tag"}
	return h.commandRunner.Run("git", args...)
}

func NewListTags(commandRunner *shell.CommandRunner) *gena.Tool {
	return gena.NewTool().
		WithName("git_list_tags").
		WithDescription("Returns the list of git tags").
		WithHandler(NewListTagsHandler(commandRunner)).
		WithSchema(
			gena.H{
				"type":       "object",
				"properties": gena.H{},
				"required":   []string{},
			},
		)
}
