package models

import (
	"github.com/google/uuid"
	"github.com/monopeelz/linear-avocado/internal/project/entity"
	"github.com/uptrace/bun"
)

type Project struct {
	bun.BaseModel `bun:"table:projects,alias:proj"`
	ID            int       `bun:"id,pk,autoincrement" json:"-"`
	UUID          uuid.UUID `bun:"uuid" json:"uuid"`
	Name          string    `bun:"name" json:"name"`
	URL           string    `bun:"url" json:"url"`
	Jobs          []Job     `bun:"rel:has-many,join:id=project_id" json:"jobs,omitempty"`
}

func (p Project) Entity() entity.Project {
	return entity.Project{
		UUID: p.UUID.String(),
		URL:  p.URL,
		Name: p.Name,
	}
}

type UpdateProject struct {
	bun.BaseModel `bun:"table:projects,alias:proj"`
	ID            int    `bun:"id,pk,autoincrement" json:"-"`
	UUID          string `bun:"uuid" json:"uuid"`
	Name          string `bun:"name" json:"name"`
	URL           string `bun:"url" json:"url"`
	Jobs          []Job  `bun:"rel:has-many,join:id=project_id" json:"jobs,omitempty"`
}
