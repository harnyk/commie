package git

import (
	"errors"
	"os/exec"

	"github.com/harnyk/gena"
)

type AddParams struct {
	Files []string `mapstructure:"files"`
}

var Add gena.TypedHandler[AddParams, string] = func(params AddParams) (string, error) {
	if len(params.Files) == 0 {
		return "", errors.New("no files specified")
	}

	args := append([]string{"add"}, params.Files...)
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func NewAdd() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("add").
		WithDescription("Adds files to the git staging area").
		WithHandler(Add.AcceptingMapOfAny()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"files": H{
						"type": "array",
						"items": H{
							"type": "string",
						},
						"description": "List of files to add to staging",
					},
				},
				"required": []string{"files"},
			},
		)

	return tool
}
