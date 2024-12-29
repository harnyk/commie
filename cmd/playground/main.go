package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/harnyk/commie/cmd/playground/tools/cat"
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
	_, err := toml.DecodeFile("config.toml", &cfg)
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

			If you compose commit messages, you should analyze the changes and provide a commit message
			following conventional commits format.
		`).
		WithTool(ls.New()).
		WithTool(cat.New()).
		WithTool(git.NewStatus()).
		WithTool(git.NewDiff()).
		Build()

	//-------------------------------------------------------------------------

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your question: ")
		question, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		answer, err := inforg.Ask(context.Background(), question)
		if err != nil {
			fmt.Println("Error processing question:", err)
			continue
		}

		fmt.Println("Answer:", answer)
	}
}
