package memory

import (
	"log/slog"
	"os"
	"path/filepath"
)

/*
Resolve strategy:

Let CWD is /home/user/alex/projects/foo.
Memory file is serached in the following order:
 - /home/alex/projects/foo/.commie/memory.yaml

 - /home/alex/projects/.commie/profiles/foo/memory.yaml
 - /home/alex/projects/.commie/memory.yaml

 - /home/alex/.commie/profiles/projects/memory.yaml
 - /home/alex/.commie/memory.yaml

 - /home/.commie/profiles/alex/memory.yaml
 - /home/.commie/memory.yaml

 - /.commie/profiles/home/memory.yaml
 - /.commie/memory.yaml

First found file will be used.

If the FS root (/ for unix or Disk:\ for windows) is reached, the search will be stopped, and the file will be created at the CDW/.commie/memory.yaml. In this example the file will be created at /home/alex/projects/foo/.commie/memory.yaml

*/

type MemoryFileResolver struct {
	log *slog.Logger
}

func NewMemoryFileResolver(log *slog.Logger) *MemoryFileResolver {
	return &MemoryFileResolver{
		log: log,
	}
}

func (m *MemoryFileResolver) File() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	current := cwd
	profileName := ""
	for {
		commiePath := filepath.Join(current, ".commie")

		paths := make([]string, 0, 2)
		if profileName != "" {
			paths = append(paths, filepath.Join(commiePath, "profiles", profileName, "memory.yaml"))
		}
		paths = append(paths, filepath.Join(commiePath, "memory.yaml"))

		for _, path := range paths {
			m.log.Debug("trying memory file", "path", path)
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
		}

		profileName = filepath.Base(current)
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	defaultPath := filepath.Join(cwd, ".commie", "memory.yaml")
	err = os.MkdirAll(filepath.Dir(defaultPath), 0755)
	if err != nil {
		return "", err
	}

	file, err := os.Create(defaultPath)
	if err != nil {
		return "", err
	}
	m.log.Debug("created memory file", "path", defaultPath)
	defer file.Close()

	return defaultPath, nil
}
