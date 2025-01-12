package main

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/harnyk/commie/pkg/toolmw"
	"github.com/harnyk/commie/pkg/tools/filesystem"
	"github.com/harnyk/commie/pkg/tools/git"
	"github.com/harnyk/commie/pkg/tools/memory"
	"github.com/harnyk/commie/pkg/tools/shell"
	"github.com/harnyk/gena"
)

func createAgent(profileDir string, log *slog.Logger) *gena.Agent {
	memFile := filepath.Join(profileDir, "memory.yaml")
	log.Debug("memory file", "path", memFile)
	memoryRepo := memory.NewMemoryRepoYAMLFile(memFile)

	promptTextWithMemory := strings.Builder{}
	promptTextWithMemory.WriteString(promptText)

	toc, _ := memoryRepo.GetTOC()
	if len(toc) > 0 {
		promptTextWithMemory.WriteString("\nCurrent memory items:\n")
		for _, item := range toc {
			tagsString := strings.Join(item.Tags, ",")
			promptTextWithMemory.WriteString(fmt.Sprintf("- id:'%s', tags:%s\n", item.ID, tagsString))
		}
	}

	agent := gena.NewAgent().
		WithOpenAIKey(cfg.OpenAIKey).
		WithOpenAIModel(cfg.OpenAIModel).
		WithSystemPrompt(promptTextWithMemory.String()).
		WithLogger(log).
		WithTemperature(0.7).
		// fs tools
		WithTool(filesystem.NewLs()).
		WithTool(filesystem.NewRealpath()).
		WithTool(filesystem.NewList()).
		WithTool(filesystem.NewRm()).
		WithTool(filesystem.NewDump()).
		WithTool(filesystem.NewMkdir()).
		WithTool(
			shell.New().
				WithMiddleware(toolmw.NewConsentMmiddleware("The agent is about to execute the following command:\n```shell\n{{.command}}\n```\n"))).
		WithTool(shell.NewPing()).
		// git tools
		WithTool(git.NewStatus()).
		WithTool(git.NewDiff()).
		WithTool(git.NewCommit()).
		WithTool(git.NewPush()).
		WithTool(git.NewAdd()).
		WithTool(git.NewLog()).
		WithTool(git.NewPRDiff()).
		// memory tools
		WithTool(memory.NewSet(memoryRepo)).
		WithTool(memory.NewGet(memoryRepo))

	if cfg.OpenAIAPIURL != "" {
		agent.WithAPIURL(cfg.OpenAIAPIURL)
	}

	return agent.Build()
}
