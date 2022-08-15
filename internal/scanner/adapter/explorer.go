package adapter

import (
	"fmt"
	"github.com/monopeelz/linear-avocado/internal/scanner/ports"
	"os"
	"path/filepath"
)

type projectExplore struct {
}

func (p projectExplore) Explore(root string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fmt.Println(path, info.Size())
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func NewProjectExplore() ports.ProjectExplorer {
	return &projectExplore{}
}
