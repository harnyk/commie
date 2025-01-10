package templaterunner

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateRunner struct {
}

func New() *TemplateRunner {
	return &TemplateRunner{}
}

func (t *TemplateRunner) Run(
	templatePath string,
) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	tmpl.Funcs(t.templateFuncs())

	var output string
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		return "", err
	}

	return output, nil
}

func (t *TemplateRunner) templateFuncs() template.FuncMap {
	return template.FuncMap{
		"ReadTextFile": ReadTextFile,
		"GlobFiles":    GlobFiles,
	}
}

func ReadTextFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	return string(bytes), err
}

func GlobFiles(glob string) (string, error) {
	files, err := filepath.Glob(glob)
	if err != nil {
		return "", err
	}
	result := strings.Builder{}
	for _, file := range files {
		fstat, err := os.Stat(file)
		if err != nil {
			return "", err
		}

		if fstat.IsDir() {
			result.WriteString(file + "/")
		} else {
			result.WriteString(file)
		}
		result.WriteString("\n")
	}
	return result.String(), nil
}
