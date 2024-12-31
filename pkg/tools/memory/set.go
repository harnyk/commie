package memory

import (
	"github.com/harnyk/commie/pkg/agent"
)

type SetParams struct {
	ID      string   `json:"id"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func NewSet(repo MemoryRepo) *agent.Tool {
	var set = agent.NewTypedHandler(func(params SetParams) (any, error) {
		err := repo.Save(&MemoryItem{
			ID:      params.ID,
			Content: params.Content,
			Tags:    params.Tags,
		})
		return nil, err
	})

	return agent.NewTool().
		WithName("memory_set").
		WithDescription("Sets the content of a memory item. Use it when you need to save some knowledge between sessions. Never set an item without reading it first!").
		WithHandler(set.AcceptingMapOfAny()).
		WithSchema(
			agent.H{
				"type": "object",
				"properties": agent.H{
					"id": agent.H{
						"type":        "string",
						"description": "The ID of the memory item to set. Also serves as the title. For example: \"Conversation Language\"",
					},
					"content": agent.H{
						"type":        "string",
						"description": "The information to set. If omitted, the item will be deleted. Short advice on the topic described in the title.",
						"minLength":   10,
						"maxLength":   600,
					},
					"tags": agent.H{
						"type":        "array",
						"description": "The tags to set. 1 to 4. Prefer consistent short tags. For example: [personalization, git, user-info, guidline",
						"items":       agent.H{"type": "string"},
					},
				},
				"required": []string{"id", "tags"},
			},
		)
}
