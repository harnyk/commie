package cat

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/harnyk/gena"
)

type CatParams struct {
	File   string `mapstructure:"file"`
	Length int    `mapstructure:"length"`
	Offset int    `mapstructure:"offset"`
}

var Cat gena.TypedHandler[CatParams, string] = func(params CatParams) (string, error) {
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

func New() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
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
						"description": "The number of bytes to read. Recommended is 1024",
						"minimum":     1,
						"maximum":     2048,
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
