package scanner

// Scanner file finding interface
type Scanner interface {
	ScanFile(path string) ([]Finding, error)
	Rules() []Rule
}
