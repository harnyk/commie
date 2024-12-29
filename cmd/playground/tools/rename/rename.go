package rename

import (
	"errors"
	"os"

	"github.com/harnyk/commie/pkg/agent"
)

type RenameParams struct {
	OldPath string `mapstructure:"old_path"`
	NewPath string `mapstructure:"new_path"`
}

var Rename agent.TypedHandler[RenameParams, string] = func(params RenameParams) (string, error) {
	if params.OldPath == "" {
		return "", errors.New("no old path specified")
	}
	
	if params.NewPath == "" {
		return "", errors.New("no new path specified")
	}

	// Rename (or move) the file
	if err := os.Rename(params.OldPath, params.NewPath); err != nil {
		return "", err
	}

	return "File renamed/moved successfully", nil
}

func New() *agent.Tool {
	type H = agent.H

	tool := agent.NewTool().
		WithName("rename").
		WithDescription("Renames or moves a file").
		WithHandler(Rename.AcceptingMapOfAny()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"old_path": H{
						"type":        "string",
						"description": "The current file path",
					},
					"new_path": H{
						"type":   "string",
						"description": "The new file path",
					},
				},
				"required": []string{"old_path", "new_path"},
			},
		)

	return tool
}
