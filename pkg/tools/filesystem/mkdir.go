package filesystem

import (
	"errors"
	"os"

	"github.com/harnyk/gena"
)

// MkdirParams holds the parameters for the Mkdir tool.
type MkdirParams struct {
	Dir string `mapstructure:"dir"`
}

// Mkdir is the Mkdir tool.
type Mkdir struct{}

// NewMkdirHandler creates a new handler for the Mkdir tool.
func NewMkdirHandler() gena.ToolHandler {
	return &Mkdir{}
}

// Execute runs the Mkdir tool with the given parameters.
func (m *Mkdir) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(m.execute, params)
}

// execute creates a directory.
func (m *Mkdir) execute(params MkdirParams) (string, error) {
	if params.Dir == "" {
		return "", errors.New("no directory specified")
	}

	err := os.MkdirAll(params.Dir, 0755)
	if err != nil {
		return "", err
	}

	return "Directory successfully created", nil
}

// NewMkdir creates a new Mkdir tool.
func NewMkdir() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("filesystem_mkdir").
		WithDescription("Creates a directory").
		WithHandler(NewMkdirHandler()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"dir": H{
						"type":        "string",
						"description": "The directory path to create",
					},
				},
				"required": []string{"dir"},
			},
		)

	return tool
}