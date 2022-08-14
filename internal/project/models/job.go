package models

import (
	"github.com/monopeelz/linear-avocado/internal/pkg/proj"
	"github.com/monopeelz/linear-avocado/internal/project/entity"
	"time"
)

type Job struct {
	// ID: Job ID
	ID int `bun:"id,pk,autoincrement" json:"-"`
	// UUID: Global ID
	UUID string `bun:"uuid" json:"uuid"`
	// ProjectID: Project ID use for RDBMS
	ProjectID int `bun:"project_id" json:"-"`
	// CreatedAt: timestamp
	CreatedAt *time.Time
	// StartedAt: timestamp
	StartedAt *time.Time
	// FinishedAt: timestamp
	FinishedAt *time.Time
	// Detail: output from processing
	Detail map[string]interface{} `bun:"type:json"`
	// StatusID: "Queued" | "In Progress" | "Success" | "Failure"
	StatusID proj.ScanningStatus
}

func (j Job) Entity() entity.Job {
	return entity.Job{
		ID:         j.ID,
		UUID:       j.UUID,
		CreatedAt:  j.CreatedAt,
		StartedAt:  j.CreatedAt,
		FinishedAt: j.FinishedAt,
		Status:     j.StatusID.String(),
	}
}
