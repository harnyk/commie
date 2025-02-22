package pathresolver

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var FileNotFoundError = errors.New("file not found")

type Pathresolver struct {
	pathString string
	pathItems  []string
	extensions []string
}

func New(pathString string) *Pathresolver {
	return &Pathresolver{
		pathString: pathString,
		pathItems:  filepath.SplitList(pathString),
		extensions: []string{""},
	}
}

func (p *Pathresolver) AddExtensions(extensions []string) *Pathresolver {
	p.extensions = append(p.extensions, extensions...)
	return p
}

func (p *Pathresolver) PrependPath(path string) *Pathresolver {
	p.pathItems = append([]string{path}, p.pathItems...)
	return p
}

func (p *Pathresolver) AppendPath(path string) *Pathresolver {
	p.pathItems = append(p.pathItems, path)
	return p
}

func (p *Pathresolver) ResolveFileName(shortName string) (string, error) {
	for _, dir := range p.pathItems {
		for _, extension := range p.extensions {
			name := shortName
			if extension != "" {
				name += "." + extension
			}
			name = filepath.Join(dir, name)
			stat, err := os.Stat(name)
			if err == nil && stat.Mode().IsRegular() {
				return name, nil
			}
		}
	}
	return "", fmt.Errorf("file: '%s' not found in path '%s': %w", shortName, p.pathString, FileNotFoundError)
}
