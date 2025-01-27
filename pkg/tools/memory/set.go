package memory

import (
	"errors"

	"github.com/harnyk/gena"
)

type SetParams struct {
	ID      string   `json:"id"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type SetHandler struct {
	repo MemoryRepo
}

func NewSetHandler(repo MemoryRepo) gena.ToolHandler {
	return &SetHandler{
		repo: repo,
	}
}

func (h *SetHandler) execute(params SetParams) (any, error) {
	existing, err := h.repo.GetById(params.ID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("item already exists")
	}

	err = h.repo.Save(&MemoryItem{
		ID:      params.ID,
		Content: params.Content,
		Tags:    params.Tags,
	})
	return "Item saved", err
}

func (h *SetHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped[SetParams, any](h.execute, params)
}

func NewSet(repo MemoryRepo) *gena.Tool {
	return gena.NewTool().
		WithName("knowledge_write").
		WithDescription("Updates or creates the content of a memory item. Use it when you need to save some knowledge between sessions. Prefer creating new items over updating the existing").
		WithHandler(NewSetHandler(repo)).
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
