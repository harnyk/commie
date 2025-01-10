package realpath

import (
	"errors"
	"path/filepath"

	"github.com/harnyk/gena"
)

// RealpathParams defines the parameters for the realpath tool.

type RealpathParams struct {
	Path string `json:"path"`
}

// RealpathHandler handles the realpath tool.

type RealpathHandler struct{}

// NewRealpathHandler creates a new realpath handler.

func NewRealpathHandler() gena.ToolHandler {
	return &RealpathHandler{}
}

// Execute executes the realpath tool.

func (h *RealpathHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

// execute performs the realpath operation.

func (h *RealpathHandler) execute(params RealpathParams) (any, error) {
	// Ensure the path is not empty
	if params.Path == "" {
		return nil, errors.New("no path specified")
	}

	// Resolve the absolute path
	resolvedPath, err := filepath.Abs(params.Path)
	if err != nil {
		return nil, err
	}

	return resolvedPath, nil
}

// New creates a new realpath tool.

func New() *gena.Tool {
	return gena.NewTool().
		WithName("realpath").
		WithDescription("Returns the real path of the specified path").
		WithHandler(NewRealpathHandler()).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"path": gena.H{
						"type":        "string",
						"description": "The path to resolve",
					},
				},
				"required": []string{"path"},
			},
		)
}
