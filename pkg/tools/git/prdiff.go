package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/harnyk/gena"
)

type GitPRDiffParams struct {
	AgainstRevision string   `mapstructure:"against_revision"`
	Files           []string `mapstructure:"files"`
}

type PRDiffHandler struct {
}

func NewPRDiffHandler() gena.ToolHandler {
	return &PRDiffHandler{}
}

func (h *PRDiffHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *PRDiffHandler) execute(params GitPRDiffParams) (string, error) {
	args := []string{"-c", "git diff $(git merge-base " + params.AgainstRevision + " HEAD)..HEAD"}
	args[1] += " " + strings.Join(params.Files, " ")

	fmt.Println(args)

	output, err := exec.Command("bash", args...).CombinedOutput()
	if err != nil {
		return "", err
	}

	diff := string(output)

	return diff, nil
}

func NewPRDiff() *gena.Tool {
	return gena.NewTool().
		WithName("gitPullRequestDiff").
		WithDescription("Returns a chunk of the diff between the merge base of the specified revision and HEAD. Use it when you want to get changes of the current pull request").
		WithHandler(NewPRDiffHandler()).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"against_revision": gena.H{
						"type":        "string",
						"description": "The revision to compare to. Required",
					},
					"files": gena.H{
						"type":        "array",
						"description": "List of files to include in the diff. Optional",
						"items":       gena.H{"type": "string"},
					},
				},
				"required": []string{"against_revision"},
			},
		)
}