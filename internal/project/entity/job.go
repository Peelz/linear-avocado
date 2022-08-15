package entity

import "time"

type Job struct {
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

	Status string `json:"status,omitempty"`

	Project Project `json:"project"`

	Detail map[string]interface{} `bun:"type:json"`
}
