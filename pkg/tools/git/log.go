package git

import (
	"os/exec"
	"strconv"

	"github.com/harnyk/gena"
)

type GitLogParams struct {
	Revision string `mapstructure:"revision"`
	Skip     int    `mapstructure:"skip"`
	MaxCount int    `mapstructure:"max_count"`
}

var GitLogHandler gena.TypedHandler[GitLogParams, string] = func(params GitLogParams) (string, error) {
	args := []string{"log"}
	if params.Revision != "" {
		args = append(args, params.Revision)
	}

	skip := params.Skip
	if skip < 0 {
		skip = 0
	}

	maxCount := params.MaxCount
	if maxCount <= 0 {
		maxCount = 10
	}
	if maxCount > 100 {
		maxCount = 100
	}

	args = append(args, "--skip="+strconv.Itoa(skip))
	args = append(args, "--max-count="+strconv.Itoa(maxCount))

	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func NewLog() *gena.Tool {
	return gena.NewTool().
		WithName("gitLog").
		WithDescription("Returns the git log with pagination support").
		WithHandler(GitLogHandler.AcceptingMapOfAny()).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"revision": gena.H{"type": "string"},
					"skip": gena.H{
						"type":    "integer",
						"minimum": 0,
					},
					"max_count": gena.H{
						"type":    "integer",
						"minimum": 1,
						"maximum": 100,
					},
				},
				"required": []string{"skip", "max_count"},
			},
		)
}
