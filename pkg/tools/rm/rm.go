package rm

import (
	"errors"
	"os"

	"github.com/harnyk/gena"
)

type RmParams struct {
	File string `mapstructure:"file"`
}

type Rm struct{}

func NewRm() gena.ToolHandler {
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

func New() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("rm").
		WithDescription("Deletes a specified file").
		WithHandler(NewRm()).
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
