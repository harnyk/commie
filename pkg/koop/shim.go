package koop

import "github.com/harnyk/gena"

type KoopHandlerShim struct {
	koop     *Koop
	toolName string
}

var _ gena.ToolHandler = (*KoopHandlerShim)(nil)

func NewKoopHandlerShim(
	koop *Koop,
	toolName string,
) *KoopHandlerShim {
	return &KoopHandlerShim{
		koop:     koop,
		toolName: toolName,
	}
}

func (s *KoopHandlerShim) Execute(params gena.H) (any, error) {
	return s.koop.CallTool(s.toolName, params)
}
