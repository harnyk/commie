package agent

// Tool is a struct that represents a tool that can be used by the agent.

/*
Example usage:

H = agent.H

type Params struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Response struct {
	Sum float64 `json:"sum"`
}

tool := agent.NewTool[Params, Response]().
	WithName("sum").
	WithDescription("Returns the sum of two numbers").
	WithSchema(H{
		"type": "object",
		"properties": {
			"a": {
				"type": "number"
			},
			"b": {
				"type": "number"
			}
		},
		"required": ["a", "b"]
	}).
	WithHandler(func(params Params) (Response, error) {
		return Response{
			Sum: params.A + params.B,
		}, nil
	})
*/

type H map[string]interface{}

type Tool[T any, R any] struct {
	Name        string
	Description string
	Schema      H
	Handler     func(T) (R, error)
}

func NewTool[T any, R any]() *Tool[T, R] {
	return &Tool[T, R]{}
}

func (t *Tool[T, R]) WithName(name string) *Tool[T, R] {
	t.Name = name
	return t
}

func (t *Tool[T, R]) WithDescription(description string) *Tool[T, R] {
	t.Description = description
	return t
}

func (t *Tool[T, R]) WithSchema(schema H) *Tool[T, R] {
	t.Schema = schema
	return t
}

func (t *Tool[T, R]) WithHandler(handler func(T) (R, error)) *Tool[T, R] {
	t.Handler = handler
	return t
}

func (t *Tool[T, R]) Run(params T) (R, error) {
	// TODO: validate params with some JSON-schema validator
	return t.Handler(params)
}

func (t *Tool[T, R]) String() string {
	return t.Name
}
