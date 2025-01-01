package rm

import (
	"errors"
	"os"

	"github.com/harnyk/gena"
)

type RmParams struct {
	File string `mapstructure:"file"`
}

var Rm gena.TypedHandler[RmParams, string] = func(params RmParams) (string, error) {
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
		WithHandler(Rm.AcceptingMapOfAny()).
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
