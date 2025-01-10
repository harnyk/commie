package filesystem

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/harnyk/gena"
)

type ListParams struct {
	File  string `mapstructure:"file"`
	Start int    `mapstructure:"start"`
	End   int    `mapstructure:"end"`
}

type List struct {
}

func NewListHandler() gena.ToolHandler {
	return &List{}
}

func (h *List) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *List) execute(params ListParams) (any, error) {
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
			output.WriteString(scanner.Text() + "\n")
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fixedEndLine := params.End
	if fixedEndLine > totalLines {
		fixedEndLine = totalLines
	}
	eofReached := fixedEndLine == totalLines

	stats := fmt.Sprintf("Lines %d...%d of %d", params.Start, fixedEndLine, totalLines)

	return map[string]string{
		"content": output.String(),
		"stats":   stats,
		"eof":     strconv.FormatBool(eofReached),
	}, nil
}

func NewList() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("filesystem_list").
		WithDescription("Prints specific lines of a text file with statistics. Prefer 1000 lines at a time").
		WithHandler(NewListHandler()).
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
						"description": "The ending line number (1-based, inclusive)",
						"minimum":     1,
					},
				},
				"required": []string{"file", "start", "end"},
			},
		)

	return tool
}
