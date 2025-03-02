package main

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/harnyk/commie/pkg/chat"
	"github.com/harnyk/commie/pkg/koop"
	"github.com/harnyk/commie/pkg/toolfactories"
	"github.com/harnyk/commie/pkg/tools/memory"
	"github.com/harnyk/gena"
)

func createChat(profileDir string, log *slog.Logger) *chat.Chat {
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

	// cmdRunner := shellService.NewCommandRunner()

	// gitFactory := toolfactories.NewGitToolFactory(cmdRunner)

	// fsFactory := toolfactories.NewFsToolFactory()

	memoryFactory := toolfactories.NewMemoryToolFactory(memoryRepo)

	agent := gena.NewAgent().
		WithOpenAIKey(cfg.OpenAIKey).
		WithOpenAIModel(cfg.OpenAIModel).
		WithSystemPrompt(promptTextWithMemory.String()).
		WithLogger(log).
		WithTemperature(0.7).
		// // fs tools
		// WithTool(fsFactory.NewLs()).
		// WithTool(fsFactory.NewRealpath()).
		// WithTool(fsFactory.NewList()).
		// WithTool(fsFactory.NewRm()).
		// WithTool(fsFactory.NewRename()).
		// WithTool(fsFactory.NewDump()).
		// WithTool(fsFactory.NewMkdir()).
		// WithTool(
		// 	shell.New(cmdRunner).
		// 		WithMiddleware(toolmw.NewConsentMiddleware("Commie is about to execute the following command:\n```shell\n{{.command}}\n```\n"))).
		// // git tools
		// WithTool(gitFactory.NewStatus()).
		// WithTool(gitFactory.NewListTags()).
		// WithTool(gitFactory.NewAdd()).
		// WithTool(gitFactory.NewDiff()).
		// WithTool(gitFactory.NewCommit()).
		// WithTool(gitFactory.NewPush()).
		// WithTool(gitFactory.NewLog()).
		// WithTool(gitFactory.NewPRDiff()).
		// // memory tools
		WithTool(memoryFactory.NewGet()).
		WithTool(memoryFactory.NewSet()).
		WithTool(memoryFactory.NewDel())

	if cfg.OpenAIAPIURL != "" {
		agent.WithAPIURL(cfg.OpenAIAPIURL)
	}

	agent.Build()

	chat := chat.New(agent)

	exampleKoop := koop.NewKoop()
	if err := exampleKoop.LoadFromFile("./examples/koops/dev/koop.yaml"); err != nil {
		log.Error("failed to load example koop", "error", err)
	}
	koop.UseKoop(agent, exampleKoop, "default")

	prompts := exampleKoop.ListPrompts()
	for _, promptName := range prompts {
		p, _ := exampleKoop.GetPrompt(promptName)
		chat.AddSystemPrompt(promptName, p)
	}

	return chat
}
