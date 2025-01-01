package memory

import (
	"errors"

	"github.com/harnyk/gena"
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

func NewGet(repo MemoryRepo) *gena.Tool {
	var get = gena.NewTypedHandler(func(params GetParams) (any, error) {
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
			return gena.H{
				"message": "tags loaded. use `knowledge_read what=by_tag tag=<tag>` to get a list of items with that tag",
				"tags":    tags}, nil
		case GetParamWhatToc:
			toc, err := repo.GetTOC()
			if err != nil {
				return nil, err
			}
			return gena.H{
				"message":       "TOC loaded. use `knowledge_read what=by_id id=<id>` to get an item",
				"ItemsWithTags": toc}, nil
		default:
			return nil, nil
		}
	})

	return gena.NewTool().
		WithName("knowledge_read").
		WithDescription("Gets the content of your memory notes, or a list of them, or list of tags.").
		WithHandler(get.AcceptingMapOfAny()).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"what": gena.H{
						"type":        "string",
						"description": "Controls what to get",
						"enum":        []string{string(GetParamWhatByID), string(GetParamWhatByTag), string(GetParamWhatTags), string(GetParamWhatToc)},
					},
					"id": gena.H{
						"type":        "string",
						"description": "The ID of the memory item to get. Required if `what` is `by_id`",
					},
					"tag": gena.H{
						"type":        "string",
						"description": "The tag of the memory item to get. Required if `what` is `by_tag`",
					},
				},
				"required": []string{"what"},
			},
		)
}
