package usecase

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/monopeelz/linear-avocado/internal/project/entity"
	"github.com/monopeelz/linear-avocado/internal/project/models"
	"github.com/monopeelz/linear-avocado/internal/project/ports"
	"go.uber.org/zap"
)

// ProjectUseCase Interface for Project business logic
type ProjectUseCase interface {
	Create(ctx context.Context, i models.Project) (models.Project, error)
	Update(ctx context.Context, i models.UpdateProject) (models.Project, error)
	Delete(ctx context.Context, i string) error
	All(ctx context.Context) ([]models.Project, error)
	Get(ctx context.Context, i string) (models.Project, error)
	Scan(ctx context.Context, i string) (entity.Job, error)
}

type projectUseCase struct {
	rp     ports.ProjectRepository
	msg    ports.MessagePublisher
	logger *zap.Logger
}

func (p *projectUseCase) Scan(ctx context.Context, i string) (entity.Job, error) {
	proj, err := p.rp.GetByID(ctx, i)
	if err != nil {
		return entity.Job{}, err
	}

	_job, err := p.rp.CreateJobFromProject(ctx, proj)
	job := _job.Entity()
	job.Project = proj.Entity()
	if err != nil {
		return entity.Job{}, err
	}
	b, err := json.Marshal(job)
	if err != nil {
		p.logger.Error("", zap.Error(err))
		return entity.Job{}, err
	}

	go func() {
		err = p.msg.Publish(ctx, "scan-project", b)
		if err != nil {
			p.logger.Error("", zap.Error(err))
		}

	}()

	return job, nil
}

func (p *projectUseCase) Create(ctx context.Context, i models.Project) (models.Project, error) {
	i.UUID, _ = uuid.NewUUID()
	return p.rp.Create(ctx, i)
}

func (p *projectUseCase) Update(ctx context.Context, i models.UpdateProject) (models.Project, error) {
	return p.rp.Update(ctx, i)
}

func (p *projectUseCase) Delete(ctx context.Context, i string) error {
	return p.rp.DeleteBytProjectID(ctx, i)
}

func (p *projectUseCase) All(ctx context.Context) ([]models.Project, error) {
	return p.rp.GetAll(ctx)
}

func (p *projectUseCase) Get(ctx context.Context, i string) (models.Project, error) {
	return p.rp.GetByID(ctx, i)
}

func (p *projectUseCase) GetJob(ctx context.Context, i string) ([]models.Job, error) {
	return p.rp.GetJobs(ctx, i)
}

func NewProjectUseCase(
	rp ports.ProjectRepository,
	msg ports.MessagePublisher,
	logger *zap.Logger,
) ProjectUseCase {
	return &projectUseCase{
		rp,
		msg,
		logger,
	}
}
