package list

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/harnyk/commie/pkg/agent"
)

type ListParams struct {
	File  string `mapstructure:"file"`
	Start int    `mapstructure:"start"`
	End   int    `mapstructure:"end"`
}

var List agent.TypedHandler[ListParams, map[string]string] = func(params ListParams) (map[string]string, error) {
	if params.File == "" {
		return nil, errors.New("no file specified")
	}

	if params.Start < 1 {
		return nil, errors.New("start line must be greater than or equal to 1")
	}

	if params.End < params.Start {
		return nil, errors.New("end line must be greater than or equal to start line")
	}

	file, err := os.Open(params.File)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output strings.Builder
	lineNumber := 1
	totalLines := 0

	for scanner.Scan() {
		totalLines++
		if lineNumber >= params.Start && lineNumber <= params.End {
			output.WriteString(strconv.Itoa(lineNumber) + "|" + scanner.Text() + "\n")
		}

		if lineNumber > params.End {
			break
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return map[string]string{
		"content": output.String(),
		"stats":   "lines " + strconv.Itoa(params.Start) + "..." + strconv.Itoa(params.End) + " of " + strconv.Itoa(totalLines),
	}, nil
}

func New() *agent.Tool {
	type H = agent.H

	tool := agent.NewTool().
		WithName("list").
		WithDescription("Prints specific lines of a text file with line numbers and statistics. Line numbers are separated by '|' and are not a part of content").
		WithHandler(List.AcceptingMapOfAny()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"file": H{
						"type":        "string",
						"description": "The file path to read",
					},
					"start": H{
						"type":        "integer",
						"description": "The starting line number (1-based)",
						"minimum":     1,
					},
					"end": H{
						"type":        "integer",
						"description": "The ending line number (inclusive)",
						"minimum":     1,
					},
				},
				"required": []string{"file", "start", "end"},
			},
		)

	return tool
}
