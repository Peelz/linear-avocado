package models

import (
	"github.com/google/uuid"
	"time"
)

type JobStatus string

func (j JobStatus) String() string {
	return string(j)
}

const (
	Queued     JobStatus = "Queued"
	InProgress JobStatus = "InProgress"
	Success    JobStatus = "Success"
	Failure    JobStatus = "Failure"
)

type JobDetail struct {
	Findings []map[string]interface{} `json:"findings"`
	Errors   []map[string]interface{} `json:"errors"`
}

type Job struct {
	// ID: Job ID
	ID int `bun:"id,pk,autoincrement" json:"-"`
	// UUID: Global ID
	UUID uuid.UUID `bun:"uuid,scanonly" json:"uuid"`
	// ProjectID: Project ID use for RDBMS
	ProjectID int `json:"projectId"`
	// CreatedAt: timestamp
	CreatedAt *time.Time `json:"createdAt"`
	// StartedAt: timestamp
	StartedAt *time.Time `json:"startedAt"`
	// FinishedAt: timestamp
	FinishedAt *time.Time `json:"finishedAt"`
	// Detail: output from processing
	Detail map[string]interface{} `bun:"type:json" json:"detail"`
	// StatusID: "Queued" | "In InProgress" | "Success" | "Failure"
	Status  JobStatus `bun:"type:string" json:"status"`
	Project Project   `json:"project"`
}
