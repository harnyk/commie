package main

import (
	"fmt"
	"strings"

	"github.com/harnyk/commie/pkg/tools/dump"
	"github.com/harnyk/commie/pkg/tools/git"
	"github.com/harnyk/commie/pkg/tools/list"
	"github.com/harnyk/commie/pkg/tools/ls"
	"github.com/harnyk/commie/pkg/tools/memory"
	"github.com/harnyk/commie/pkg/tools/patch"
	"github.com/harnyk/commie/pkg/tools/rm"
	"github.com/harnyk/gena"
)

func createAgent() *gena.Agent {
	memoryRepo := memory.NewMemoryRepoYAMLFile("./.commie/memory.yaml")

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

	fmt.Println(promptTextWithMemory.String())

	return gena.NewAgent().
		WithOpenAIKey(cfg.OpenAIKey).
		WithOpenAIModel(cfg.OpenAIModel).
		WithSystemPrompt(promptTextWithMemory.String()).
		WithTool(ls.New()).
		WithTool(list.New()).
		WithTool(rm.New()).
		WithTool(dump.New()).
		WithTool(patch.New()).
		WithTool(git.NewStatus()).
		WithTool(git.NewDiff()).
		WithTool(git.NewCommit()).
		WithTool(git.NewPush()).
		WithTool(git.NewAdd()).
		WithTool(git.NewLog()).
		WithTool(memory.NewSet(memoryRepo)).
		WithTool(memory.NewGet(memoryRepo)).
		Build()
}
