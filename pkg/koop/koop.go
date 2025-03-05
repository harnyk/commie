package koop

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/harnyk/gena"
	"gopkg.in/yaml.v3"
)

// Koop represents a modular unit in the Commie system.
type Koop struct {
	Manifest        Manifest
	manifestCommand string
	workDir         string
}

// NewKoop creates a new instance of Koop.
func NewKoop() *Koop {
	return &Koop{}
}

func (k *Koop) LoadFromExecutable(command string) error {
	manifestExecutor := &ManifestExecutor{
		Command: command,
		Args:    []string{},
	}

	k.manifestCommand = command

	return k.executeCallableEntity(manifestExecutor, "", os.Environ(), &k.Manifest)
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
func (k *Koop) GetPrompt(name string) (string, error) {
	prompt, exists := k.Manifest.Prompts[name]
	if !exists {
		return "", fmt.Errorf("prompt %q not found", name)
	}

	if prompt.Text != "" {
		return prompt.Text, nil
	}

	if !prompt.SelfInvoke && prompt.Command == "" {
		return "", fmt.Errorf("prompt %q has no content", name)
	}

	var result string

	if err := k.executeCallableEntity(&prompt, k.manifestCommand, os.Environ(), &result); err != nil {
		return "", err
	}

	return result, nil
}

// GetDefaultPrompt returns the content of the default prompt.
func (k *Koop) GetDefaultPrompt() (string, error) {
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

	bParamsJson, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}
	paramsJson := string(bParamsJson)

	var result any

	envs := os.Environ()
	envs = append(envs, "commie.koop.tool.parameters="+paramsJson, "commie_koop_tool_parameters="+paramsJson)

	if err := k.executeCallableEntity(tool, k.manifestCommand, envs, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (k *Koop) executeCallableEntity(
	callableEntity CallableEntity,
	parentCommand string,
	envs []string,
	result any,
) error {
	var command string
	if callableEntity.GetSelfInvoke() {
		if parentCommand == "" {
			// TODO: add it to manifest validation
			return errors.New("self-invoke is not allowed without a parent command")
		}
		if callableEntity.GetCommand() != "" {
			// TODO: add it to manifest validation
			return errors.New("self-invoke is not allowed with a command")
		}
		command = parentCommand
	} else {
		command = callableEntity.GetCommand()
	}

	cmd := exec.Command(command, callableEntity.GetArgs()...)
	cmd.Dir = k.WorkDir()
	cmd.Env = envs

	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("execution error: %w, stderr: %s", err, string(exitError.Stderr))
		}
		return fmt.Errorf("execution error: %w", err)
	}

	if callableEntity.IsRaw() {
		return setPointerValue(result, string(output))
	} else {
		return json.Unmarshal(output, result)
	}
}

type rawOutputWrapped struct {
	Output string `json:"output"`
}

func wrapRawOutput(output []byte) rawOutputWrapped {
	return rawOutputWrapped{
		Output: string(output),
	}
}

func setPointerValue(ptr any, value any) error {
	v := reflect.ValueOf(ptr)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("expected a non-nil pointer")
	}

	elem := v.Elem()

	if !elem.CanSet() {
		return errors.New("cannot set value")
	}

	val := reflect.ValueOf(value)
	if elem.Type() != val.Type() {
		return fmt.Errorf("type mismatch: expected %s but got %s", elem.Type(), val.Type())
	}

	elem.Set(val)
	return nil
}
