package koop

type ManifestExecutor struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
}

func (m *ManifestExecutor) GetCommand() string {
	return m.Command
}

func (m *ManifestExecutor) GetArgs() []string {
	return m.Args
}

func (m *ManifestExecutor) GetSelfInvoke() bool {
	return false
}

func (m *ManifestExecutor) IsRaw() bool {
	return false
}
