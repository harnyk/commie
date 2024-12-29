package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
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

	type H = agent.H

	currentTimeTool := agent.NewTool().
		WithName("time").
		WithDescription("Returns the current time").
		WithHandler(func(_ any) (any, error) {
			message := "The current time is: " + time.Now().String()
			return H{"message": message}, nil
		}).
		WithSchema(
			H{
				"type":       "object",
				"properties": H{},
			},
		)

	inforg := agent.NewAgent().
		WithOpenAIKey(cfg.OpenAIKey).
		WithOpenAIModel(cfg.OpenAIModel).
		WithTool(currentTimeTool).
		Build()

	answer, err := inforg.Ask(context.Background(), "Сколько дней до конца года?")
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)

	answer, err = inforg.Ask(context.Background(), "Почему ты так считаешь?")
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}
