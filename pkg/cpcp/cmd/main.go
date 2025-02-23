package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/harnyk/commie/pkg/colorlog"
	"github.com/harnyk/commie/pkg/cpcp"
)

func main() {
	logger := slog.New(colorlog.NewColorConsoleHandler(os.Stdout, slog.LevelDebug))

	transport := cpcp.NewProcessClient(logger, "node", "./pkg/cpcp/cmd/example_plugin.js")

	reqres := cpcp.NewReqResClient(transport)

	err := reqres.Start()
	if err != nil {
		panic(err)
	}

	type AddRequestPayload = struct {
		Type string `json:"type"`
		A    int    `json:"a"`
		B    int    `json:"b"`
	}

	type AddResponsePayload = struct {
		C int `json:"c"`
	}

	req := &AddRequestPayload{
		Type: "add",
		A:    10,
		B:    20,
	}
	res := &AddResponsePayload{}

	if err = reqres.Send(req, res); err != nil {
		panic(err)
	}

	fmt.Println(res.C)
}
