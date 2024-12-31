package memory

// Memory is stored in the YAML file:
//
// items:
//   - id: git commit style
//     tags: [git, commit, style, guide]
//     content: |
//       Prefer Conventional Commits.
//       Include issue number where possible.
//       Infer it from branch.
//       Include component if possible.
//       Infe it from content of commit.
//       Example: "feat(aboutRoute): ABC-123 improve parser"
//   - id: user name
//     tags: [personalization]
//     content: |
//       Mark
//   - id: user email
//     tags: [personalization]
//     content: |
//       2F6lO@example.com

type MemoryItem struct {
	ID      string   `json:"id"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type Memory struct {
	Items []MemoryItem `json:"memoryItems"`
}

type MemoryRepo interface {
	GetById(id string) (*MemoryItem, error)
	GetByTag(tag string) ([]MemoryItem, error)
	GetTags() ([]string, error)
	GetTOC() ([]string, error)
	Save(item *MemoryItem) error
	Delete(id string) error
}
