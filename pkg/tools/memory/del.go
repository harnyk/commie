package memory

import (
	"errors"

	"github.com/harnyk/gena"
)

type DelParams struct {
	ID string `json:"id"`
}

type DelHandler struct {
	repo MemoryRepo
}

func NewDelHandler(repo MemoryRepo) gena.ToolHandler {
	return &DelHandler{
		repo: repo,
	}
}

func (h *DelHandler) execute(params DelParams) (any, error) {
	existing, err := h.repo.GetById(params.ID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("item does not exist")
	}

	err = h.repo.Delete(params.ID)
	if err != nil {
		return nil, err
	}
	return "Item deleted", nil
}

func (h *DelHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped[DelParams, any](h.execute, params)
}

func NewDel(repo MemoryRepo) *gena.Tool {
	return gena.NewTool().
		WithName("knowledge_delete").
		WithDescription("Deletes a memory item by its ID").
		WithHandler(NewDelHandler(repo)).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"id": gena.H{
						"type":        "string",
						"description": "The ID of the memory item to delete",
					},
				},
				"required": []string{"id"},
			},
		)
}
