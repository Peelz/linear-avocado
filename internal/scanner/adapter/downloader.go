package adapter

import (
	"github.com/monopeelz/linear-avocado/internal/scanner/ports"
	"os"
)
import "github.com/go-git/go-git/v5"

type gitDownloadAdapter struct {
}

func (d gitDownloadAdapter) Download(url string) (string, error) {
	dest := os.TempDir()
	_, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return "", err
	}
	return dest, nil
}

func NewDownloader() ports.Downloader {
	return &gitDownloadAdapter{}
}
