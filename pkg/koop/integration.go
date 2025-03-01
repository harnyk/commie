package koop

import (
	"fmt"

	"github.com/harnyk/gena"
)

func UseKoop(
	agent *gena.Agent,
	k *Koop,
	kPromptName string,
) error {
	prompt, ok := k.GetPrompt(kPromptName)
	if !ok {
		return fmt.Errorf("koop prompt %q not found", kPromptName)
	}
	agent.WithSystemPrompt(prompt)

	tools := k.ListTools()
	for _, toolName := range tools {
		kTool, ok := k.GetTool(toolName)
		if !ok {
			return fmt.Errorf("koop tool %q not found", toolName)
		}

		gTool := gena.NewTool().
			WithName(toolName).
			WithDescription(kTool.Description).
			WithSchema(gena.H{
				"type":       kTool.Parameters.Type,
				"required":   kTool.Parameters.Required,
				"properties": kTool.Parameters.Properties,
			}).
			WithHandler(NewKoopHandlerShim(k, toolName))

		agent.WithTool(gTool)
	}

	return nil
}
