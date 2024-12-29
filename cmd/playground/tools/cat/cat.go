package cat

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/harnyk/commie/pkg/agent"
)

type CatParams struct {
	File   string `mapstructure:"file"`
	Length int    `mapstructure:"length"`
	Offset int    `mapstructure:"offset"`
}

var Cat agent.TypedHandler[CatParams, string] = func(params CatParams) (string, error) {
	if params.File == "" {
		return "", errors.New("no file specified")
	}

	if params.Length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	file, err := os.Open(params.File)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Seek(int64(params.Offset), io.SeekStart)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(file)
	data := make([]byte, params.Length)

	count, err := reader.Read(data)
	if err != nil && err != io.EOF {
		return "", err
	}

	return string(data[:count]), nil
}

func New() *agent.Tool {
	type H = agent.H

	tool := agent.NewTool().
		WithName("cat").
		WithDescription("Prints the contents of a file").
		WithHandler(Cat.AcceptingMapOfAny()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"file": H{
						"type":        "string",
						"description": "The file path to read",
					},
					"length": H{
						"type":        "integer",
						"description": "The number of bytes to read",
						"minimum":     1,
						"maximum":     1024,
					},
					"offset": H{
						"type":        "integer",
						"description": "The offset to start reading from",
						"minimum":     0,
					},
				},
				"required": []string{"file", "length"},
			},
		)

	return tool
}
