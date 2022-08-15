package ports

// Downloader use for download project from url
type Downloader interface {
	Download(url string) (string, error)
}
