package memory

import (
	"errors"

	"github.com/harnyk/commie/pkg/agent"
)

type GetParamWhat string

const (
	GetParamWhatByID  GetParamWhat = "by_id"
	GetParamWhatByTag GetParamWhat = "by_tag"
	GetParamWhatTags  GetParamWhat = "tags"
	GetParamWhatToc   GetParamWhat = "toc"
)

type GetParams struct {
	What GetParamWhat `json:"what"`
	ID   string       `json:"id"`
	Tag  string       `json:"tag"`
}

func NewGet(repo MemoryRepo) *agent.Tool {
	var get = agent.NewTypedHandler(func(params GetParams) (any, error) {
		what := params.What
		id := params.ID
		tag := params.Tag
		switch what {
		case GetParamWhatByID:
			if id == "" {
				return nil, errors.New("no id specified")
			}
			return repo.GetById(id)
		case GetParamWhatByTag:
			if tag == "" {
				return nil, errors.New("no tag specified")
			}
			return repo.GetByTag(tag)
		case GetParamWhatTags:
			tags, err := repo.GetTags()
			if err != nil {
				return nil, err
			}
			return agent.H{
				"message": "tags loaded. use `memory_get what=by_tag tag=<tag>` to get a list of items with that tag",
				"tags":    tags}, nil
		case GetParamWhatToc:
			toc, err := repo.GetTOC()
			if err != nil {
				return nil, err
			}
			return agent.H{
				"message":       "TOC loaded. use `memory_get what=by_id id=<id>` to get an item",
				"ItemsWithTags": toc}, nil
		default:
			return nil, nil
		}
	})

	return agent.NewTool().
		WithName("memory_get").
		WithDescription("Gets the content of your memory notes, or a list of them, or list of tags.").
		WithHandler(get.AcceptingMapOfAny()).
		WithSchema(
			agent.H{
				"type": "object",
				"properties": agent.H{
					"what": agent.H{
						"type":        "string",
						"description": "Controls what to get",
						"enum":        []string{string(GetParamWhatByID), string(GetParamWhatByTag), string(GetParamWhatTags), string(GetParamWhatToc)},
					},
					"id": agent.H{
						"type":        "string",
						"description": "The ID of the memory item to get. Required if `what` is `by_id`",
					},
					"tag": agent.H{
						"type":        "string",
						"description": "The tag of the memory item to get. Required if `what` is `by_tag`",
					},
				},
				"required": []string{"what"},
			},
		)
}
