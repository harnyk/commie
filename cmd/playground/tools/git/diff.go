package git

import (
	"os/exec"

	"github.com/harnyk/commie/pkg/agent"
)

type GitDiffParams struct {
	AgainstRevision string   `mapstructure:"against_revision"`
	Files           []string `mapstructure:"files"`
	Offset          int      `mapstructure:"offset"`
	Length          int      `mapstructure:"length"`
}

var GitDiffHandler agent.TypedHandler[GitDiffParams, string] = func(params GitDiffParams) (string, error) {
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
	if length <= 0 || length > 1024 {
		length = 1024
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

func NewDiff() *agent.Tool {
	return agent.NewTool().
		WithName("gitDiff").
		WithDescription("Returns a chunk of the diff between current state and specified revision, starting from offset with specified length").
		WithHandler(GitDiffHandler.AcceptingMapOfAny()).
		WithSchema(
			agent.H{
				"type": "object",
				"properties": agent.H{
					"against_revision": agent.H{
						"type":        "string",
						"description": "The revision to compare to. Optional",
					},
					"files": agent.H{
						"type":        "array",
						"description": "List of files to include in the diff. Optional",
						"items":       agent.H{"type": "string"},
					},
					"offset": agent.H{
						"type":        "integer",
						"description": "The offset in bytes to start the chunk. Default is 0.",
					},
					"length": agent.H{
						"type":        "integer",
						"description": "The maximum length of the chunk in bytes. Max is 1024.",
					},
				},
				"required": []string{},
			},
		)
}
