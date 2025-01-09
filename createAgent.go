package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/harnyk/commie/pkg/colorlog"
	"github.com/harnyk/commie/pkg/tools/dump"
	"github.com/harnyk/commie/pkg/tools/git"
	"github.com/harnyk/commie/pkg/tools/list"
	"github.com/harnyk/commie/pkg/tools/ls"
	"github.com/harnyk/commie/pkg/tools/memory"
	"github.com/harnyk/commie/pkg/tools/rm"
	"github.com/harnyk/gena"
)

func createAgent() *gena.Agent {
	log := slog.New(colorlog.NewColorConsoleHandler(os.Stderr))
	memResolver := memory.NewMemoryFileResolver(log)
	memFile, err := memResolver.File()
	if err != nil {
		log.Error("failed to get memory file", "error", err)
		os.Exit(1)
	}
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
		WithLogger(slog.New(colorlog.NewColorConsoleHandler(os.Stderr))).
		WithTemperature(0.7).
		WithTool(ls.New()).
		WithTool(list.New()).
		WithTool(rm.New()).
		WithTool(dump.New()).
		// WithTool(httpget.New()).
		WithTool(git.NewStatus()).
		WithTool(git.NewDiff()).
		WithTool(git.NewCommit()).
		WithTool(git.NewPush()).
		WithTool(git.NewAdd()).
		WithTool(git.NewLog()).
		WithTool(git.NewPRDiff()).
		WithTool(memory.NewSet(memoryRepo)).
		WithTool(memory.NewGet(memoryRepo))

	if cfg.OpenAIAPIURL != "" {
		agent.WithAPIURL(cfg.OpenAIAPIURL)
	}

	return agent.Build()
}
