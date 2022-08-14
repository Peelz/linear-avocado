package usecase

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/monopeelz/linear-avocado/internal/project/models"
	"github.com/monopeelz/linear-avocado/internal/project/ports"
)

// ProjectUseCase Interface for Project business logic
type ProjectUseCase interface {
	Create(ctx context.Context, i models.Project) (models.Project, error)
	Update(ctx context.Context, i models.UpdateProject) (models.Project, error)
	Delete(ctx context.Context, i string) error
	All(ctx context.Context) ([]models.Project, error)
	Get(ctx context.Context, i string) (models.Project, error)
	Scan(ctx context.Context, i string) error
}

type projectUseCase struct {
	rp  ports.ProjectRepository
	msg ports.MessagePublisher
}

func (p projectUseCase) Scan(ctx context.Context, i string) error {
	proj, err := p.rp.GetByID(ctx, i)
	if err != nil {
		return err
	}
	b, err := json.Marshal(proj)
	if err != nil {
		return err
	}
	err = p.msg.Publish(ctx, "scan-project", b)
	if err != nil {
		return err
	}

	return nil
}

func (p projectUseCase) Create(ctx context.Context, i models.Project) (models.Project, error) {
	i.UUID, _ = uuid.NewUUID()
	return p.rp.Create(ctx, i)
}

func (p projectUseCase) Update(ctx context.Context, i models.UpdateProject) (models.Project, error) {
	return p.rp.Update(ctx, i)
}

func (p projectUseCase) Delete(ctx context.Context, i string) error {
	return p.rp.DeleteBytProjectID(ctx, i)
}

func (p projectUseCase) All(ctx context.Context) ([]models.Project, error) {
	return p.rp.GetAll(ctx)
}

func (p projectUseCase) Get(ctx context.Context, i string) (models.Project, error) {
	return p.rp.GetByID(ctx, i)
}

func (p projectUseCase) GetJob(ctx context.Context, i string) ([]models.Job, error) {
	return p.rp.GetJobs(ctx, i)
}

func NewProjectUseCase(rp ports.ProjectRepository, msg ports.MessagePublisher) ProjectUseCase {
	return &projectUseCase{
		rp,
		msg,
	}
}
