package memory

import (
	"github.com/harnyk/gena"
)

type SetParams struct {
	ID      string   `json:"id"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func NewSet(repo MemoryRepo) *gena.Tool {
	var set = gena.NewTypedHandler(func(params SetParams) (any, error) {
		err := repo.Save(&MemoryItem{
			ID:      params.ID,
			Content: params.Content,
			Tags:    params.Tags,
		})
		return nil, err
	})

	return gena.NewTool().
		WithName("knowledge_write").
		WithDescription("Sets the content of a memory item. Use it when you need to save some knowledge between sessions. Never set an item without reading it first!").
		WithHandler(set.AcceptingMapOfAny()).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"id": gena.H{
						"type":        "string",
						"description": "The ID of the memory item to set. Also serves as the title. For example: \"Conversation Language\"",
					},
					"content": gena.H{
						"type":        "string",
						"description": "The information to set. If omitted, the item will be deleted. Short advice on the topic described in the title.",
						"minLength":   10,
						"maxLength":   600,
					},
					"tags": gena.H{
						"type":        "array",
						"description": "The tags to set. 1 to 4. Prefer consistent short tags. For example: [personalization, git, user-info, guidline",
						"items":       gena.H{"type": "string"},
					},
				},
				"required": []string{"id", "tags"},
			},
		)
}
