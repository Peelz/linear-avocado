package proj

import (
	"time"
)

// Project A entity for represent user repository that created
type Project struct {
	// Id: any type of unique id
	ID string
	// Status: "Queued" | "In Progress" | "Success" | "Failure"
	Status ScanningStatus
	// RepositoryName: string
	Name string
	// RepositoryUrl: string
	URL string
	// Findings: JSONB, see example
	Finding interface{}
	// QueuedAt: timestamp
	QueuedAt *time.Time
	// ScanningAt: timestamp
	ScanningAt *time.Time
	// FinishedAt: timestamp
	FinishedAt *time.Time
}
