package git

import (
	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/gena"
)

const maxDiffLength = 4096


type GitDiffParams struct {
	AgainstRevision string   `mapstructure:"against_revision"`
	Files           []string `mapstructure:"files"`
	Offset          int      `mapstructure:"offset"`
	Length          int      `mapstructure:"length"`
}


type DiffHandler struct {
	commandRunner *shell.CommandRunner
}


func NewDiffHandler(commandRunner *shell.CommandRunner) gena.ToolHandler {
	return &DiffHandler{
		commandRunner: commandRunner,
	}
}


func (h *DiffHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}


func (h *DiffHandler) execute(params GitDiffParams) (string, error) {
	args := []string{"diff"}
	if params.AgainstRevision != "" {
		args = append(args, params.AgainstRevision)
	}
	args = append(args, params.Files...) // Добавляем файлы, если указаны

	output, err := h.commandRunner.Run("git", args...)
	if err != nil {
		return "", err
	}

	length := params.Length
	if length <= 0 || length > maxDiffLength {
		length = maxDiffLength
	}

	diff := string(output)
	if params.Offset < 0 {
		params.Offset = 0
	}
	if params.Offset >= len(diff) {
		return "", nil
	}

	end := params.Offset + length
	if end > len(diff) {
		end = len(diff)
	}

	return diff[params.Offset:end], nil
}


func NewDiff(commandRunner *shell.CommandRunner) *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("gitDiff").
		WithDescription("Returns a chunk of the diff between current state and specified revision, starting from offset with specified length").
		WithHandler(NewDiffHandler(commandRunner)).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"against_revision": H{
						"type":        "string",
						"description": "The revision to compare to. Optional",
					},
					"files": H{
						"type":        "array",
						"description": "List of files to include in the diff. Optional",
						"items":       H{"type": "string"},
					},
					"offset": H{
						"type":        "integer",
						"description": "The offset in bytes to start the chunk. Default is 0.",
					},
					"length": H{
						"type":        "integer",
						"description": "The maximum length of the chunk in bytes. Max is " + string(maxDiffLength) + ". Default is " + string(maxDiffLength),
					},
				},
				"required": []string{},
			},
		)

	return tool
}