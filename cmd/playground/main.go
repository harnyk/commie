package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/harnyk/commie/cmd/playground/tools/cat"
	"github.com/harnyk/commie/cmd/playground/tools/dump"
	"github.com/harnyk/commie/cmd/playground/tools/git"
	"github.com/harnyk/commie/cmd/playground/tools/ls"
	"github.com/harnyk/commie/pkg/agent"
)

type Config struct {
	OpenAIKey   string `toml:"OPENAI_KEY"`
	OpenAIModel string `toml:"OPENAI_MODEL"`
}

func main() {
	cfg := Config{}
	_, err := toml.DecodeFile("config.toml", 	&cfg)
	if err != nil {
		panic(err)
	}

	inforg := agent.NewAgent().
		WithOpenAIKey(cfg.OpenAIKey).
		WithOpenAIModel(cfg.OpenAIModel).
		WithSystemPrompt(`
			You are a helpful assistant which helps a user to work with the file system, terminal and git.
			Your responses will be rendered directly to the modern Linux terminal,
			so you should use ASCII art, emojis for icons, ASCII terminal codes for colors.
			Markdown is not allowed, if you use it, the whole response will be broken.
			Reply with just a plain text with no markdown.

			If the user asks to do something, you should do your best and provide deep analysis using the
			available tools.

			If you compose commit messages, you should
			 - analyze the changes
			 - read the git diffs
			 - if necessary, read through the sources
			 - reason about the changes
			 - compose a concise commit message as a summary of the changes in "conventional commits" format.
		`).
		WithCommands(git.NewCommit(), git.NewPush())

	// Use the configured agent for the application logic
}
