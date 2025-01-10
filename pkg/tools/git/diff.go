package git

import (
	"os/exec"

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
}

func NewDiffHandler() gena.ToolHandler {
	return &DiffHandler{}
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

	output, err := exec.Command("git", args...).CombinedOutput()
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

func NewDiff() *gena.Tool {
	return gena.NewTool().
		WithName("gitDiff").
		WithDescription("Returns a chunk of the diff between current state and specified revision, starting from offset with specified length").
		WithHandler(NewDiffHandler()).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"against_revision": gena.H{
						"type":        "string",
						"description": "The revision to compare to. Optional",
					},
					"files": gena.H{
						"type":        "array",
						"description": "List of files to include in the diff. Optional",
						"items":       gena.H{"type": "string"},
					},
					"offset": gena.H{
						"type":        "integer",
						"description": "The offset in bytes to start the chunk. Default is 0.",
					},
					"length": gena.H{
						"type":        "integer",
						"description": "The maximum length of the chunk in bytes. Max is " + string(maxDiffLength) + ". Default is " + string(maxDiffLength),
					},
				},
				"required": []string{},
			},
		)
}
