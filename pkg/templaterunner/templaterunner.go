package templaterunner

import (
	"path/filepath"
	"strings"
	"text/template"

	"github.com/harnyk/commie/pkg/shell"
)

type TemplateRunner struct {
	commandRunner *shell.CommandRunner
}

func New(
	commandRunner *shell.CommandRunner,
) *TemplateRunner {
	return &TemplateRunner{
		commandRunner: commandRunner,
	}
}

func (t *TemplateRunner) Run(
	templatePath string,
) (string, error) {
	tmpl, err := template.
		New(filepath.Base(templatePath)).
		Funcs(t.templateFuncs()).
		ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	output := &strings.Builder{}

	err = tmpl.Execute(output, nil)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (t *TemplateRunner) templateFuncs() template.FuncMap {
	return template.FuncMap{
		"shell": t.ShellFn,
	}
}

func (t *TemplateRunner) ShellFn(command string) (string, error) {
	return t.commandRunner.RunString(command)
}
