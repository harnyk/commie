package git

import (
	"os/exec"

	"github.com/harnyk/gena"
)

type PushParams struct {
	Remote string `mapstructure:"remote"`
	Branch string `mapstructure:"branch"`
}

var Push gena.TypedHandler[PushParams, string] = func(params PushParams) (string, error) {
	remote := "origin"
	branch := "main"

	if params.Remote != "" {
		remote = params.Remote
	}

	if params.Branch != "" {
		branch = params.Branch
	}

	cmd := exec.Command("git", "push", remote, branch)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}

func NewPush() *gena.Tool {
	type H = gena.H

	tool := gena.NewTool().
		WithName("push").
		WithDescription("Pushes commits to the remote repository").
		WithHandler(Push.AcceptingMapOfAny()).
		WithSchema(
			H{
				"type": "object",
				"properties": H{
					"remote": H{
						"type":        "string",
						"description": "The remote to push to (default: origin)",
					},
					"branch": H{
						"type":        "string",
						"description": "The branch to push (default: main)",
					},
				},
			},
		)

	return tool
}
