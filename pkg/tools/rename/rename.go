package rename

import (
	"errors"
	"os"

	"github.com/harnyk/gena"
)

type RenameParams struct {
	OldPath string `mapstructure:"old_path"`
	NewPath string `mapstructure:"new_path"`
}

type RenameHandler struct {
}

func NewRenameHandler() gena.ToolHandler {
	return &RenameHandler{}
}

func (h *RenameHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *RenameHandler) execute(params RenameParams) (any, error) {
	if params.OldPath == "" {
		return nil, errors.New("no old path specified")
	}

	if params.NewPath == "" {
		return nil, errors.New("no new path specified")
	}

	// Rename (or move) the file
	if err := os.Rename(params.OldPath, params.NewPath); err != nil {
		return nil, err
	}

	return "File renamed/moved successfully", nil
}

func New() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("rename").
		WithDescription("Renames or moves a file").
		WithHandler(NewRenameHandler()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"old_path": H{
						"type":        "string",
						"description": "The current file path",
					},
					"new_path": H{
						"type":        "string",
						"description": "The new file path",
					},
				},
				"required": []string{"old_path", "new_path"},
			},
		)

	return tool
}
