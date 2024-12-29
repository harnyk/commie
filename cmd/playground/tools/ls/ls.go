package ls

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/harnyk/commie/pkg/agent"
)

type LsParams struct {
	Directory string
}

var Ls agent.TypedHandler[LsParams, []string] = func(params LsParams) ([]string, error) {
	if params.Directory == "" {
		return nil, errors.New("no directory specified")
	}

	files, err := os.ReadDir(params.Directory)
	if err != nil {
		return nil, err
	}

	var result []string

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, err
		}

		var line string
		if file.IsDir() {
			line = fmt.Sprintf("d: %s", file.Name())
		} else {
			line = fmt.Sprintf("f: %s %8d %s", file.Name(), info.Size(), info.Mode().String())
		}
		result = append(result, line)
	}

	sort.Strings(result)

	return result, nil
}

func New() *agent.Tool {
	return agent.NewTool().
		WithName("ls").
		WithDescription("Returns the list of files with permissions and length in bytes").
		WithHandler(Ls.AcceptingMapOfAny()).
		WithSchema(
			agent.H{
				"type": "object",
				"properties": agent.H{
					"directory": agent.H{
						"type":        "string",
						"description": "The directory to list. Use '.' for the current directory",
					},
				},
				"required": []string{"directory"},
			},
		)
}
