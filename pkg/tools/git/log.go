package git

import (
	"os/exec"

	"github.com/harnyk/commie/pkg/agent"
)

type GitLogParams struct {
	Revision string `mapstructure:"revision"`
	Offset   int    `mapstructure:"offset"`
	Length   int    `mapstructure:"length"`
}

var GitLogHandler agent.TypedHandler[GitLogParams, string] = func(params GitLogParams) (string, error) {
	args := []string{"log"}
	if params.Revision != "" {
		args = append(args, params.Revision)
	}

	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return "", err
	}

	length := params.Length
	if length <= 0 || length > 1024 {
		length = 1024
	}

	log := string(output)
	if params.Offset < 0 {
		params.Offset = 0
	}
	if params.Offset >= len(log) {
		return "", nil
	}

	end := params.Offset + length
	if end > len(log) {
		end = len(log)
	}

	return log[params.Offset:end], nil
}

func NewLog() *agent.Tool {
	return agent.NewTool().
		WithName("gitLog").
		WithDescription("Returns the git log with pagination support").
		WithHandler(GitLogHandler.AcceptingMapOfAny()).
		WithSchema(
			agent.H{
				"type": "object",
				"properties": agent.H{
					"revision": agent.H{"type": "string"},
					"offset":   agent.H{"type": "integer"},
					"length":   agent.H{"type": "integer"},
				},
				"required": []string{"offset", "length"},
			},
		)
}
