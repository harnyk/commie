package patch

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/harnyk/commie/pkg/agent"
)

type PatchParams struct {
	File  string `mapstructure:"file"`
	Patch string `mapstructure:"patch"`
	Fuzz  int    `mapstructure:"fuzz,omitempty"`
}

var Patch agent.TypedHandler[PatchParams, string] = func(params PatchParams) (string, error) {
	if params.File == "" {
		return "", errors.New("no file specified")
	}

	if params.Patch == "" {
		return "", errors.New("no patch content provided")
	}

	// Create all parent directories if they don't exist
	dir := filepath.Dir(params.File)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	// Prepare the patch command
	cmd := exec.Command("patch", params.File, "-i", "-", "--batch")

	// Add fuzz factor if specified
	if params.Fuzz > 0 {
		cmd.Args = append(cmd.Args, "--fuzz="+strconv.Itoa(params.Fuzz))
	}

	cmd.Stdin = strings.NewReader(params.Patch)

	// Run the patch command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(output))
	}

	return "Patch applied successfully", nil
}

func New() *agent.Tool {
	type H = agent.H

	tool := agent.NewTool().
		WithName("patch").
		WithDescription("Applies a universal diff patch to a file").
		WithHandler(Patch.AcceptingMapOfAny()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"file": H{
						"title":       "File path",
						"type":        "string",
						"description": "Path to the file to apply the patch to",
					},
					"patch": H{
						"title":       "Patch content",
						"type":        "string",
						"description": "The patch content to apply",
					},
					"fuzz": H{
						"title":       "Fuzz factor",
						"type":        "integer",
						"description": "The fuzz factor to use when applying the patch",
					},
				},
				"required": []string{"file", "patch", "fuzz"},
			},
		)

	return tool
}
