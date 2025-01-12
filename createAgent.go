package main

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"

	shellService "github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/commie/pkg/toolfactories"
	"github.com/harnyk/commie/pkg/toolmw"
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

	cmdRunner := shellService.NewCommandRunner()

	gitFactory := toolfactories.NewGitToolFactory(cmdRunner)

	fsFactory := toolfactories.NewFsToolFactory()

	agent := gena.NewAgent().
		WithOpenAIKey(cfg.OpenAIKey).
		WithOpenAIModel(cfg.OpenAIModel).
		WithSystemPrompt(promptTextWithMemory.String()).
		WithLogger(log).
		WithTemperature(0.7).
		// fs tools
		WithTool(fsFactory.NewLs()).
		WithTool(fsFactory.NewRealpath()).
		WithTool(fsFactory.NewList()).
		WithTool(fsFactory.NewRm()).
		WithTool(fsFactory.NewRename()).
		WithTool(fsFactory.NewDump()).
		WithTool(fsFactory.NewMkdir()).
		WithTool(
			shell.New(cmdRunner).
				WithMiddleware(toolmw.NewConsentMiddleware("Commie is about to execute the following command:\n```shell\n{{.command}}\n```\n"))).
		// git tools
		WithTool(gitFactory.NewStatus()).
		WithTool(gitFactory.NewAdd()).
		WithTool(gitFactory.NewDiff()).
		WithTool(gitFactory.NewCommit()).
		WithTool(gitFactory.NewPush()).
		WithTool(gitFactory.NewLog()).
		WithTool(gitFactory.NewPRDiff()).
		// memory tools
		WithTool(memory.NewSet(memoryRepo)).
		WithTool(memory.NewGet(memoryRepo))

	if cfg.OpenAIAPIURL != "" {
		agent.WithAPIURL(cfg.OpenAIAPIURL)
	}

	return agent.Build()
}
