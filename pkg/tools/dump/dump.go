package dump

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/harnyk/gena"
)

type DumpParams struct {
	File    string `mapstructure:"file"`
	Content string `mapstructure:"content"`
}

type Dump struct{}

func NewDump() gena.ToolHandler {
	return &Dump{}
}

func (d *Dump) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(d.execute, params)
}

func (d *Dump) execute(params DumpParams) (string, error) {
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

func New() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("dump").
		WithDescription("Writes content to a file. This function is very dangerous! Never call it if you haven't read the ENTIRE content of the file (from the first to the last line), otherwise you will lose part of the file").
		WithHandler(NewDump()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"file": H{
						"type":        "string",
						"description": "The file path to write to",
					},
					"content": H{
						"type":        "string",
						"description": "The content to write to the file",
					},
				},
				"required": []string{"file", "content"},
			},
		)

	return tool
}
