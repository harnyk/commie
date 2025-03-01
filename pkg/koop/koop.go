package koop

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/harnyk/gena"
	"gopkg.in/yaml.v3"
)

// Koop represents a modular unit in the Commie system.
type Koop struct {
	Manifest Manifest
	workDir  string
}

// NewKoop creates a new instance of Koop.
func NewKoop() *Koop {
	return &Koop{}
}

// LoadFromFile loads the Koop configuration from a YAML file.
func (k *Koop) LoadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&k.Manifest); err != nil {
		return err
	}

	// Validate the manifest after loading
	if err := k.Manifest.ValidateManifest(); err != nil {
		return err
	}

	// Validate tools
	if err := k.validateTools(); err != nil {
		return err
	}

	// Set the work directory
	k.workDir = filepath.Dir(filePath)

	return nil
}

func (k *Koop) validateTools() error {
	tools := k.Manifest.Tools

	toolNames := make(map[string]struct{})

	for _, tool := range tools {
		if _, ok := toolNames[tool.Name]; ok {
			return errors.New("duplicate tool name")
		}
		toolNames[tool.Name] = struct{}{}
	}
	return nil
}

// WorkDir returns the directory from which the manifest was loaded.
func (k *Koop) WorkDir() string {
	return k.workDir
}

// ListPrompts returns a list of all prompt names available in the Koop.
func (k *Koop) ListPrompts() []string {
	promptNames := make([]string, 0, len(k.Manifest.Prompts))
	for name := range k.Manifest.Prompts {
		promptNames = append(promptNames, name)
	}
	return promptNames
}

// GetPrompt returns the content of the prompt with the given name.
func (k *Koop) GetPrompt(name string) (string, bool) {
	prompt, exists := k.Manifest.Prompts[name]
	return prompt, exists
}

// GetDefaultPrompt returns the content of the default prompt.
func (k *Koop) GetDefaultPrompt() (string, bool) {
	return k.GetPrompt("default")
}

// ListTools returns a list of all tool names available in the Koop.
func (k *Koop) ListTools() []string {
	toolNames := make([]string, 0, len(k.Manifest.Tools))
	for _, tool := range k.Manifest.Tools {
		toolNames = append(toolNames, tool.Name)
	}
	return toolNames
}

// GetTool returns the Tool structure for the given tool name.
func (k *Koop) GetTool(name string) (*Tool, bool) {
	for _, tool := range k.Manifest.Tools {
		if tool.Name == name {
			return &tool, true
		}
	}
	return nil, false
}

// CallTool executes a tool with the given name and parameters.
func (k *Koop) CallTool(toolName string, parameters gena.H) (any, error) {
	tool, ok := k.GetTool(toolName)
	if !ok {
		return nil, errors.New("tool not found")
	}

	cmd := exec.Command(tool.Command, tool.Args...)
	cmd.Dir = k.WorkDir()

	bParamsJson, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}
	paramsJson := string(bParamsJson)

	// TODO: maybe whitelisting/blacklisting?
	cmd.Env = os.Environ()

	cmd.Env = append(cmd.Env, "KOOP_TOOL_PARAMETERS="+paramsJson)

	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("tool execution error: %w, stderr: %s", err, string(exitError.Stderr))
		}
		return nil, fmt.Errorf("tool execution error: %w", err)
	}

	if tool.Raw {
		return wrapRawOputput(output), nil
	}

	var response any
	err = json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type rawOutputWrapped struct {
	Output string `json:"output"`
}

func wrapRawOputput(output []byte) rawOutputWrapped {
	return rawOutputWrapped{
		Output: string(output),
	}
}
