package agent

import "errors"

type H map[string]interface{}

type Tool struct {
	Name        string
	Description string
	Schema      H
	Handler     func(any) (any, error)
}

func NewTool() *Tool {
	return &Tool{}
}

func (t *Tool) WithName(name string) *Tool {
	t.Name = name
	return t
}

func (t *Tool) WithDescription(description string) *Tool {
	t.Description = description
	return t
}

func (t *Tool) WithSchema(schema H) *Tool {
	t.Schema = schema
	return t
}

func (t *Tool) WithHandler(handler func(any) (any, error)) *Tool {
	t.Handler = handler
	return t
}

func (t *Tool) Run(params any) (any, error) {
	// TODO: validate params with some JSON-schema validator
	if t.Handler == nil {
		return nil, errors.New("no handler defined for tool")
	}
	return t.Handler(params)
}

func (t *Tool) String() string {
	return t.Name
}
