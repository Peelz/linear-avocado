package ports

import (
	"context"
	"github.com/monopeelz/linear-avocado/internal/project/models"
)

type ProjectRepository interface {
	Create(ctx context.Context, project models.Project) (models.Project, error)
	Update(ctx context.Context, project models.UpdateProject) (models.Project, error)
	DeleteBytProjectID(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]models.Project, error)
	GetByID(ctx context.Context, id string) (models.Project, error)
	GetJobs(ctx context.Context, id string) ([]models.Job, error)
}
