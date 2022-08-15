package adapter

import (
	"github.com/monopeelz/linear-avocado/internal/scanner/ports"
	"io/ioutil"
	"os"
)
import "github.com/go-git/go-git/v5"

// GitAdapter implement ports.Download to use git clone instead direct download blob
type GitAdapter struct {
}

func (d GitAdapter) Download(url string) (string, error) {
	dest, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}
	_, err = git.PlainClone(dest, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return "", err
	}
	// Remove .git after clone
	if err := os.RemoveAll(dest + "/.git"); err != nil {
		return "", err
	}

	return dest, nil
}

func NewGitAdapter() ports.Downloader {
	return &GitAdapter{}
}
