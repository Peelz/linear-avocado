package entity

import "time"

type Job struct {
	// ID: Job ID
	ID int
	// UUID: Global ID
	UUID string
	// ProjectID: Project ID use for RDBMS
	ProjectID uint
	// ProjectUUID:  Global ID use for expose public
	ProjectUUID string
	// CreatedAt: timestamp
	CreatedAt *time.Time
	// StartedAt: timestamp
	StartedAt *time.Time
	// FinishedAt: timestamp
	FinishedAt *time.Time
	// Status: current status and value is enum of "Queued" | "In Progress" | "Success" | "Failure"
	Status string `json:"status,omitempty"`

	Detail map[string]interface{} `bun:"type:json"`
}
