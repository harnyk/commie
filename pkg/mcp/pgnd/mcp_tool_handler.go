package main

import (
	"context"

	"github.com/harnyk/gena"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

type McpToolHandler struct {
	tool   *mcp.Tool
	client client.MCPClient
}

func NewMcpToolHandler(mcpClient client.MCPClient, tool *mcp.Tool) gena.ToolHandler {
	return &McpToolHandler{
		tool:   tool,
		client: mcpClient,
	}
}

func (h *McpToolHandler) Execute(params gena.H) (any, error) {
	req := mcp.CallToolRequest{}
	req.Params.Name = h.tool.Name
	req.Params.Arguments = params

	res, err := h.client.CallTool(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
