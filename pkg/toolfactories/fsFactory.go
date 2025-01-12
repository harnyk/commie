package toolfactories

import (
	"github.com/harnyk/commie/pkg/toolmw"
	"github.com/harnyk/commie/pkg/tools/filesystem"
	"github.com/harnyk/gena"
)

type FsToolFactory struct {
}

func NewFsToolFactory() *FsToolFactory {
	return &FsToolFactory{}
}

func (f *FsToolFactory) NewDump() *gena.Tool {
	consentMiddleware := toolmw.NewConsentMiddleware("Commie is about to write to the file **{{ .file }}**")
	return filesystem.NewDump().WithMiddleware(consentMiddleware)
}

func (f *FsToolFactory) NewList() *gena.Tool {
	return filesystem.NewList()
}

func (f *FsToolFactory) NewLs() *gena.Tool {
	return filesystem.NewLs()
}

func (f *FsToolFactory) NewMkdir() *gena.Tool {
	return filesystem.NewMkdir()
}

func (f *FsToolFactory) NewRealpath() *gena.Tool {
	return filesystem.NewRealpath()
}

func (f *FsToolFactory) NewRename() *gena.Tool {
	consentMiddleware := toolmw.NewConsentMiddleware("Commie is about to rename the file **{{ .old_path }}** to **{{ .new_path }}**")
	return filesystem.NewRename().WithMiddleware(consentMiddleware)
}

func (f *FsToolFactory) NewRm() *gena.Tool {
	consentMiddleware := toolmw.NewConsentMiddleware("Commie is about to delete the file **{{ .file }}**")
	return filesystem.NewRm().WithMiddleware(consentMiddleware)
}
