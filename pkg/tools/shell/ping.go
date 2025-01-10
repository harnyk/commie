package shell

import (
	"bytes"
	"errors"
	"os/exec"

	"github.com/harnyk/gena"
)

// PingParams holds the parameters for the Ping tool.
type PingParams struct {
	Host string `mapstructure:"host"`
}

// Ping is the Ping tool.
type Ping struct{}

// NewPingHandler creates a new handler for the Ping tool.
func NewPingHandler() gena.ToolHandler {
	return &Ping{}
}

// Execute runs the Ping tool with the given parameters.
func (p *Ping) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(p.execute, params)
}

// execute pings a host.
func (p *Ping) execute(params PingParams) (string, error) {
	if params.Host == "" {
		return "", errors.New("no host specified")
	}

	cmd := exec.Command("ping", "-c", "4", params.Host)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// NewPing creates a new Ping tool.
func NewPing() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("shell_ping").
		WithDescription("Pings a host").
		WithHandler(NewPingHandler()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"host": H{
						"type":        "string",
						"description": "The host to ping",
					},
				},
				"required": []string{"host"},
			},
		)

	return tool
}
