package dump

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/harnyk/commie/pkg/agent"
)

type DumpParams struct {
	File    string `mapstructure:"file"`
	Content string `mapstructure:"content"`
}

var Dump agent.TypedHandler[DumpParams, string] = func(params DumpParams) (string, error) {
	if params.File == "" {
		return "", errors.New("no file specified")
	}

	if params.Content == "" {
		return "", errors.New("no content to write")
	}

	// Create all parent directories if they don't exist
	dir := filepath.Dir(params.File)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	f, err := os.Create(params.File)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.WriteString(params.Content)
	if err != nil {
		return "", err
	}

	return "File written successfully", nil
}

func New() *agent.Tool {
	type H = agent.H

	tool := agent.NewTool().
		WithName("dump").
		WithDescription("Writes content to a file").
		WithHandler(Dump.AcceptingMapOfAny()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"file": H{
						"type":        "string",
						"description": "The file path to write to",
					},
					"content": H{
						"type":    "string",
						"description": "The content to write to the file",
					},
				},
				"required": []string{"file", "content"},
			},
		)

	return tool
}
