package agent

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

// example usage:
/*

agent = agent.NewAgent().
	WithOpenAIKey("your-openai-key").
	WithOpenAIModel("gpt-3.5-turbo").
	WithSystemPrompt("Hello, how are you?").
	WithMaxTokens(5000).
	WithTemperature(0.1).
	WithTool(currentTimeTool).
	Build()


answer, err := agent.Ask("What is the current date?")

*/

type Agent struct {
	OpenAIKey    string
	OpenAIModel  string
	SystemPrompt string
	MaxTokens    int
	Temperature  float64
	Tools        []Tool[any, any]
	client       *openai.Client
}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) Build() *Agent {
	a.client = openai.NewClient(a.OpenAIKey)
	return a
}

func (a *Agent) WithOpenAIKey(key string) *Agent {
	a.OpenAIKey = key
	return a
}

func (a *Agent) WithOpenAIModel(model string) *Agent {
	a.OpenAIModel = model
	return a
}

func (a *Agent) WithSystemPrompt(prompt string) *Agent {
	a.SystemPrompt = prompt
	return a
}

func (a *Agent) WithMaxTokens(tokens int) *Agent {
	a.MaxTokens = tokens
	return a
}

func (a *Agent) WithTemperature(temperature float64) *Agent {
	a.Temperature = temperature
	return a
}

func (a *Agent) WithTool(tool Tool[any, any]) *Agent {
	a.Tools = append(a.Tools, tool)
	return a
}

func (a *Agent) Ask(ctx context.Context, question string) (string, error) {
	a.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: a.OpenAIModel,
			Tools: a.getOpenAITools(),
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: a.SystemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)
}

func (a *Agent) getOpenAITools() []openai.Tool {
	var tools []openai.Tool
	for _, tool := range a.Tools {
		tools = append(tools, openai.Tool{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        tool.Name,
				Description: tool.Description,
				Parameters:  tool.Schema,
				Strict:      true,
			},
		})
	}
	return tools
}
