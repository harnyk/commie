package userscript

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/commie/pkg/templaterunner"
)

type ScriptType int

const (
	ScriptTypeUnknown    ScriptType = iota
	ScriptTypeShell      ScriptType = iota
	ScriptTypeMarkdown   ScriptType = iota
	ScriptTypeGoTemplate ScriptType = iota
)

type Runner struct {
	templateRunner *templaterunner.TemplateRunner
	shellRunner    *shell.CommandRunner
}

func New(templateRunner *templaterunner.TemplateRunner, shellRunner *shell.CommandRunner) *Runner {
	return &Runner{
		templateRunner: templateRunner,
		shellRunner:    shell.NewCommandRunner(),
	}
}

func (r *Runner) Run(scriptPath string) (string, error) {
	scriptType, err := r.detectScriptType(scriptPath)
	if err != nil {
		return "", err
	}
	switch scriptType {
	case ScriptTypeShell:
		return r.runShellScript(scriptPath)
	case ScriptTypeMarkdown:
		return r.runMarkdownScript(scriptPath)
	case ScriptTypeGoTemplate:
		return r.runTemplateScript(scriptPath)
	default:
		return "", errors.New("unknown script type")
	}
}

func (r *Runner) runShellScript(scriptPath string) (string, error) {
	return r.shellRunner.RunString(scriptPath)
}

func (r *Runner) runTemplateScript(scriptPath string) (string, error) {
	return r.templateRunner.Run(scriptPath)
}

func (r *Runner) runMarkdownScript(scriptPath string) (string, error) {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (r *Runner) detectScriptType(scriptPath string) (ScriptType, error) {
	fileInfo, err := os.Stat(scriptPath)
	if err != nil {
		return ScriptTypeUnknown, err
	}
	fileExtension := filepath.Ext(fileInfo.Name())
	switch fileExtension {
	case ".sh":
		return ScriptTypeShell, nil
	case ".md", ".markdown":
		return ScriptTypeMarkdown, nil
	case ".gotmpl", ".gotpl", ".tpl":
		return ScriptTypeGoTemplate, nil
	default:
		return ScriptTypeUnknown, errors.New("unknown script type")
	}
}
