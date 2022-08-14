package usecase

import "os"

// ProjectExplorer Explore project to collect files
type ProjectExplorer interface {
	Explore(root string) []os.File
}

type projectExplore struct {
}

func (p projectExplore) Explore(root string) []os.File {
	// TODO implement me
	panic("implement me")
}

func NewProjectExplore() ProjectExplorer {
	return &projectExplore{}
}
