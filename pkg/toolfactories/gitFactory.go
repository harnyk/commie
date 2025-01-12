package toolfactories

import (
	"github.com/harnyk/commie/pkg/shell"
	"github.com/harnyk/commie/pkg/toolmw"
	"github.com/harnyk/commie/pkg/tools/git"
	"github.com/harnyk/gena"
)

type GitToolFactory struct {
	cmdRunner *shell.CommandRunner
}

func NewGitToolFactory(cmdRunner *shell.CommandRunner) *GitToolFactory {
	return &GitToolFactory{cmdRunner: cmdRunner}
}

func (f *GitToolFactory) NewListTags() *gena.Tool {
	return git.NewListTags(f.cmdRunner)
}

func (f *GitToolFactory) NewCommit() *gena.Tool {
	return git.NewCommit(f.cmdRunner)
}

func (f *GitToolFactory) NewPush() *gena.Tool {
	consentTemplate := `Commie is about to push the branch **{{ or .branch "(default)" }}** to the remote repository **{{ or .remote "(default)" }}**`

	return git.NewPush(f.cmdRunner).WithMiddleware(
		toolmw.NewConsentMiddleware(consentTemplate),
	)
}

func (f *GitToolFactory) NewStatus() *gena.Tool {
	return git.NewStatus(f.cmdRunner)
}

func (f *GitToolFactory) NewAdd() *gena.Tool {
	return git.NewAdd(f.cmdRunner)
}

func (f *GitToolFactory) NewLog() *gena.Tool {
	return git.NewLog(f.cmdRunner)
}

func (f *GitToolFactory) NewDiff() *gena.Tool {
	return git.NewDiff(f.cmdRunner)
}

func (f *GitToolFactory) NewPRDiff() *gena.Tool {
	return git.NewDiff(f.cmdRunner)
}
