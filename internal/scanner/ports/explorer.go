package ports

// ProjectExplorer Explore project to collect files
type ProjectExplorer interface {
	Explore(root string) ([]string, error)
}
