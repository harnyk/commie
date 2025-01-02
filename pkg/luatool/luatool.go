package luatool

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/harnyk/gena"
)

type LuaToolOptions struct {
	Module   string
	Manifest string
}

func NewOptions() *LuaToolOptions {
	return &LuaToolOptions{}
}

func (o *LuaToolOptions) WithDir(dir string) *LuaToolOptions {
	return o.
		WithManifest(dir + "/tool.json").
		WithModule(dir + "/tool.lua")
}

func (o *LuaToolOptions) WithModule(module string) *LuaToolOptions {
	o.Module = module
	return o
}

func (o *LuaToolOptions) WithManifest(manifest string) *LuaToolOptions {
	o.Manifest = manifest
	return o
}

func NewTool(options *LuaToolOptions) (*gena.Tool, error) {
	handler := LuaHandler{
		Module: options.Module,
	}

	manifestBytes, err := os.ReadFile(options.Manifest)
	if err != nil {
		return nil, fmt.Errorf("could not read params schema: %w", err)
	}

	manifest := Manifest{}
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		return nil, fmt.Errorf("could not parse manifest: %w", err)
	}

	return gena.NewTool().
		WithName(manifest.Name).
		WithDescription(manifest.Description).
		WithSchema(manifest.Params).
		WithHandler(handler.Execute), nil
}

func MustNewTool(options *LuaToolOptions) *gena.Tool {
	tool, err := NewTool(options)
	if err != nil {
		panic(err)
	}
	return tool
}
