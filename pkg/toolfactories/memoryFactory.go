package toolfactories

import (
	"github.com/harnyk/commie/pkg/toolmw"
	"github.com/harnyk/commie/pkg/tools/memory"
	"github.com/harnyk/gena"
)

type MemoryToolFactory struct {
	memoryRepo memory.MemoryRepo
}

func NewMemoryToolFactory(memoryRepo memory.MemoryRepo) *MemoryToolFactory {
	return &MemoryToolFactory{memoryRepo: memoryRepo}
}

func (f *MemoryToolFactory) NewGet() *gena.Tool {
	return memory.NewGet(f.memoryRepo)
}

func (f *MemoryToolFactory) NewSet() *gena.Tool {
	return memory.NewSet(f.memoryRepo)
}

func (f *MemoryToolFactory) NewDel() *gena.Tool {
	consentTemplate := `Commie is about to delete the memory item **{{ .id }}**`
	return memory.NewDel(f.memoryRepo).WithMiddleware(
		toolmw.NewConsentMiddleware(consentTemplate),
	)
}
