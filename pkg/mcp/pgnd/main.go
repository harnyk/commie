package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/harnyk/gena"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	c, err := client.NewStdioMCPClient(
		"mcp-server-git",
		[]string{},
	)
	if err != nil {
		panic(err)
	}

	c.Initialize(context.Background(), mcp.InitializeRequest{})
	defer c.Close()

	tools, err := c.ListTools(context.Background(), mcp.ListToolsRequest{})
	if err != nil {
		panic(err)
	}

	agent := gena.NewAgent().
		WithSystemPrompt("You can work with git. Always use your tools to do the work. The local repository is always located in the current directory. Use dot.").
		WithLogger(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))).
		WithOpenAIModel("gpt-4o").
		WithOpenAIKey(os.Getenv("OPENAI_API_KEY"))

	for _, tool := range tools.Tools {
		gtool := gena.NewTool()
		gtool.WithName(tool.Name)
		gtool.WithDescription(tool.Description)
		schema := tool.InputSchema
		gschema := gena.H{
			"type":       schema.Type,
			"required":   schema.Required,
			"properties": schema.Properties,
		}
		gtool.WithSchema(gschema)
		gtool.WithHandler(NewMcpToolHandler(c, &tool))

		agent.WithTool(gtool)
	}

	agent.Build()

	answer, err := agent.Ask(context.Background(), "check the diffs for the current branch")
	if err != nil {
		panic(err)
	}
	fmt.Println("answer:", answer)
}
