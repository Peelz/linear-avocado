package entity

import "time"

type Job struct {
	// UUID: Global ID
	UUID string `json:"uuid"`
	// CreatedAt: timestamp
	CreatedAt *time.Time `json:"createdAt"`
	// StartedAt: timestamp
	StartedAt *time.Time `json:"startedAt"`
	// FinishedAt: timestamp
	FinishedAt *time.Time `json:"finishedAt"`

	Status string `json:"status,omitempty"`

	Project Project `json:"project"`

	Detail map[string]interface{} `bun:"type:json"`
}
