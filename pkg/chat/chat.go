package chat

import (
	"context"

	"github.com/harnyk/gena"
)

type Chat struct {
	Agent                   *gena.Agent
	systemPrompts           map[string]string
	currentSystemPromptName string
}

func New(agent *gena.Agent) *Chat {
	return &Chat{
		Agent:         agent,
		systemPrompts: map[string]string{},
	}
}

func (c *Chat) AddSystemPrompt(name string, prompt string) {
	c.systemPrompts[name] = prompt
}

func (c *Chat) SwitchSystemPrompt(name string) {
	c.currentSystemPromptName = name
}

func (c *Chat) SystemPrompt() string {
	return c.systemPrompts[c.currentSystemPromptName]
}

func (c *Chat) GetPromptNames() []string {
	names := make([]string, 0, len(c.systemPrompts))
	for name := range c.systemPrompts {
		names = append(names, name)
	}
	return names
}

func (c *Chat) Ask(ctx context.Context, message string) (string, error) {
	return c.Agent.AskWithOptions(ctx, message, gena.AskOptions{
		SystemPrompt: c.SystemPrompt(),
	})
}
