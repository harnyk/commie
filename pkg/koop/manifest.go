package koop

import (
	"github.com/go-playground/validator/v10"
	"github.com/harnyk/gena"
)

// Manifest represents the structure of a Koop manifest file.
type Manifest struct {
	Version     string            `yaml:"version" validate:"required"`
	Name        string            `yaml:"name" validate:"required"`
	Description string            `yaml:"description" validate:"required"`
	Prompts     map[string]string `yaml:"prompts"`
	Tools       []Tool            `yaml:"tools" validate:"required,dive"`
}

// Tool represents an executable tool defined in the Koop manifest.
type Tool struct {
	Name        string     `yaml:"name"`
	Raw         bool       `yaml:"raw"`
	Command     string     `yaml:"command" validate:"required"`
	Args        []string   `yaml:"args"`
	Description string     `yaml:"description" validate:"required"`
	Parameters  Parameters `yaml:"parameters"`
}

// Parameters represents the JSON schema for tool parameters.
type Parameters struct {
	Type       string   `yaml:"type" validate:"required"`
	Required   []string `yaml:"required"`
	Properties gena.H   `yaml:"properties"`
}

// ValidateManifest validates the Manifest structure
func (m *Manifest) ValidateManifest() error {
	validate := validator.New()
	return validate.Struct(m)
}
