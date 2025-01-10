package filesystem

import (
	"errors"
	"os"

	"github.com/harnyk/gena"
)

type RmParams struct {
	File string `mapstructure:"file"`
}

type Rm struct{}

func NewRmHandler() gena.ToolHandler {
	return &Rm{}
}

func (r *Rm) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(r.execute, params)
}

func (r *Rm) execute(params RmParams) (string, error) {
	if params.File == "" {
		return "", errors.New("no file specified")
	}

	err := os.Remove(params.File)
	if err != nil {
		return "", err
	}

	return "File successfully deleted", nil
}

func NewRm() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("filesystem_rm").
		WithDescription("Deletes a specified file").
		WithHandler(NewRmHandler()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"file": H{
						"type":        "string",
						"description": "The file path to delete",
					},
				},
				"required": []string{"file"},
			},
		)

	return tool
}
