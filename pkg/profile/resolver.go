package profile

import (
	"log/slog"
	"os"
	"path/filepath"
)

/*
Resolve strategy:

Let CWD is /home/user/alex/projects/foo/ .
Profile directory is serached in the following order:
 - /home/alex/projects/foo/.commie/

 - /home/alex/projects/.commie/profiles/foo/
 - /home/alex/projects/.commie/

 - /home/alex/.commie/profiles/projects/
 - /home/alex/.commie/

 - /home/.commie/profiles/alex/
 - /home/.commie/

 - /.commie/profiles/home/
 - /.commie/

First found directory will be used.

If the FS root (/ for unix or Disk:\\ for windows) is reached, the search will be stopped, and the directory will be created at the CDW/.commie/profiles. In this example the directory will be created at /home/alex/projects/foo/.commie/profiles
*/

type ProfileResolver struct {
	log *slog.Logger
}

func New(log *slog.Logger) *ProfileResolver {
	return &ProfileResolver{
		log: log,
	}
}

func (m *ProfileResolver) Get() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	current := cwd
	prev := current
	for {
		paths := make([]string, 0, 2)

		if cwd != current {
			paths = append(
				paths,
				filepath.Join(current, ".commie", "profiles", filepath.Base(prev)),
			)
		}
		paths = append(
			paths,
			filepath.Join(current, ".commie"),
		)

		for _, path := range paths {
			m.log.Debug("trying profile", "path", path)
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
		}

		prev = current
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	defaultPath := filepath.Join(cwd, ".commie", "profiles")
	err = os.MkdirAll(defaultPath, 0755)
	if err != nil {
		return "", err
	}

	m.log.Debug("created profile", "path", defaultPath)
	return defaultPath, nil
}
