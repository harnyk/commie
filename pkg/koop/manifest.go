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
	Prompts     map[string]Prompt `yaml:"prompts"`
	Tools       []Tool            `yaml:"tools" validate:"required,dive"`
}

type Prompt struct {
	Text       string   `yaml:"text"`
	SelfInvoke bool     `yaml:"selfInvoke"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args"`
}

// Tool represents an executable tool defined in the Koop manifest.
type Tool struct {
	Name        string     `yaml:"name"`
	Raw         bool       `yaml:"raw"`
	Command     string     `yaml:"command"`
	SelfInvoke  bool       `yaml:"selfInvoke"`
	Args        []string   `yaml:"args"`
	Description string     `yaml:"description" validate:"required"`
	Parameters  Parameters `yaml:"parameters"`
	Docker      Docker     `yaml:"docker"`
}

type Docker struct {
	Image   string   `yaml:"image"`
	Volumes []string `yaml:"volumes"`
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

// Implementations of CallableEntity

func (t *Tool) GetSelfInvoke() bool {
	return t.SelfInvoke
}

func (t *Tool) IsRaw() bool {
	return t.Raw
}

func (t *Tool) GetCommand() string {
	if t.Docker.Image != "" {
		return "docker"
	}
	return t.Command
}

func (t *Tool) GetArgs() []string {
	if t.Docker.Image != "" {
		dockerArgs := []string{
			"run",
			"--rm",
			"-e",
			"commie.koop.tool.parameters",
			"-e",
			"commie_koop_tool_parameters",
		}

		for _, volume := range t.Docker.Volumes {
			dockerArgs = append(dockerArgs, "-v", volume)
		}

		dockerArgs = append(dockerArgs, t.Docker.Image)
		dockerArgs = append(dockerArgs, t.Command)
		dockerArgs = append(dockerArgs, t.Args...)

		return dockerArgs
	}

	return t.Args
}

// ----

func (p *Prompt) GetSelfInvoke() bool {
	return p.SelfInvoke
}

func (p *Prompt) GetCommand() string {
	return p.Command
}

func (p *Prompt) GetArgs() []string {
	return p.Args
}

func (p *Prompt) IsRaw() bool {
	return true
}
