package memory

import (
	"os"

	"gopkg.in/yaml.v2"
)

type MemoryRepoYAMLFile struct {
	fileName string
	memory   *Memory
}

func NewMemoryRepoYAMLFile(fileName string) *MemoryRepoYAMLFile {
	return &MemoryRepoYAMLFile{
		fileName: fileName,
	}
}

func initEmptyFile(fileName string) error {
	memory := &Memory{}
	data, err := yaml.Marshal(memory)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func (m *MemoryRepoYAMLFile) load() error {
	if !exists(m.fileName) {
		if err := initEmptyFile(m.fileName); err != nil {
			return err
		}
	}

	data, err := os.ReadFile(m.fileName)
	if err != nil {
		return err
	}
	memory := &Memory{}
	if err := yaml.Unmarshal(data, memory); err != nil {
		return err
	}
	m.memory = memory
	return nil
}

func (m *MemoryRepoYAMLFile) dump() error {
	memo := m.memory
	data, err := yaml.Marshal(memo)
	if err != nil {
		return err
	}
	return os.WriteFile(m.fileName, data, 0644)
}

func (m *MemoryRepoYAMLFile) GetByTag(tag string) ([]MemoryItem, error) {
	if err := m.load(); err != nil {
		return nil, err
	}
	items := []MemoryItem{}
	for _, item := range m.memory.Items {
		for _, itemTag := range item.Tags {
			if itemTag == tag {
				items = append(items, item)
				break
			}
		}
	}
	return items, nil
}

func (m *MemoryRepoYAMLFile) GetById(id string) (*MemoryItem, error) {
	if err := m.load(); err != nil {
		return nil, err
	}
	for _, item := range m.memory.Items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, nil
}

func (m *MemoryRepoYAMLFile) GetTags() ([]string, error) {
	if err := m.load(); err != nil {
		return nil, err
	}
	tagsMap := map[string]struct{}{}
	for _, item := range m.memory.Items {
		for _, tag := range item.Tags {
			tagsMap[tag] = struct{}{}
		}
	}
	tags := []string{}
	for tag := range tagsMap {
		tags = append(tags, tag)
	}
	return tags, nil
}

func (m *MemoryRepoYAMLFile) GetTOC() ([]string, error) {
	if err := m.load(); err != nil {
		return nil, err
	}
	toc := []string{}
	for _, item := range m.memory.Items {
		toc = append(toc, item.ID)
	}
	return toc, nil
}

func (m *MemoryRepoYAMLFile) Save(item *MemoryItem) error {
	if err := m.load(); err != nil {
		return err
	}
	m.upsert(*item)
	return m.dump()
}

func (m *MemoryRepoYAMLFile) Delete(id string) error {
	if err := m.load(); err != nil {
		return err
	}
	m.deleteByID(id)
	return m.dump()
}

func (m *MemoryRepoYAMLFile) deleteByID(id string) {
	for i, item := range m.memory.Items {
		if item.ID == id {
			m.memory.Items = append(m.memory.Items[:i], m.memory.Items[i+1:]...)
		}
	}
}

func (m *MemoryRepoYAMLFile) upsert(item MemoryItem) {
	// if exists, insert in place
	for i, existing := range m.memory.Items {
		if existing.ID == item.ID {
			m.memory.Items[i] = item
			return
		}
	}
	// else prepend
	m.memory.Items = append([]MemoryItem{item}, m.memory.Items...)
}
