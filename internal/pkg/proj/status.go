package proj

type ScanningStatus int

const (
	Queued ScanningStatus = iota
	Progress
	Failure
	Unknown
)

func (s ScanningStatus) String() string {
	switch s {
	case Queued:
		return "Queued"
	case Progress:
		return "Progress"
	case Failure:
		return "Failure"
	default:
		return "Unknown"
	}
}
