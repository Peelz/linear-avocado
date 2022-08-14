package adapter

import (
	"context"
	"database/sql"
	"github.com/jackc/pgerrcode"
	"github.com/monopeelz/linear-avocado/internal/project/models"
	"github.com/monopeelz/linear-avocado/internal/project/ports"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
)

func pgError(err error) error {
	if err, ok := err.(pgdriver.Error); ok && err.IntegrityViolation() {
		return err
	} else if err.Field('C') == pgerrcode.InvalidTransactionState {
		return err
	}
	return nil
}

type projectRepository struct {
	db     *sql.DB
	bun    *bun.DB
	logger *zap.Logger
}

func (p projectRepository) GetJobs(ctx context.Context, id string) ([]models.Job, error) {
	var projID int
	var jobs []models.Job
	_ = p.bun.QueryRow("SELECT id FROM projects where uuid = $1", id).Scan(projID)
	rows, _ := p.bun.Query("SELECT * from jobs WHERE project_id = $1", projID)
	rows.Scan(&jobs)
	return jobs, nil
}

func (p projectRepository) Create(ctx context.Context, i models.Project) (models.Project, error) {
	var proj models.Project
	_, err := p.bun.NewInsert().Model(&i).
		Returning("id, uuid, name, url").
		Exec(ctx, &proj)
	if pgError(err) != nil {
		p.logger.Error("", zap.Error(err))
	}
	return proj, pgError(err)
}

func (p projectRepository) Update(ctx context.Context, i models.UpdateProject) (models.Project, error) {
	var proj models.Project
	_, err := p.bun.NewUpdate().
		Model(&i).
		Column("name", "url").
		Where("uuid = ?", i.UUID).
		Returning("id, uuid, name, url").
		Exec(ctx, &proj)
	if err != nil {
		p.logger.Error("", zap.Error(err))
	}
	return proj, err
}

func (p projectRepository) DeleteBytProjectID(ctx context.Context, id string) error {
	sql := `
    DELETE from projects
 	WHERE uuid = $1
	`
	_, err := p.bun.Exec(sql, id)
	return err
}

func (p projectRepository) GetAll(ctx context.Context) ([]models.Project, error) {
	var proj []models.Project
	sqlCmd := `
    SELECT uuid, name, url  from projects
	`
	rows, err := p.bun.QueryContext(ctx, sqlCmd)
	if err != nil {
		p.logger.Error("Error: ", zap.String("error", err.Error()))
	}
	err = p.bun.ScanRows(ctx, rows, &proj)
	if err != nil {
		p.logger.Error("Error: ", zap.String("error", err.Error()))
	}
	return proj, err
}

func (p projectRepository) GetByID(ctx context.Context, id string) (models.Project, error) {
	var proj models.Project
	err := p.bun.NewSelect().
		Model(&proj).
		Where("uuid = ?", id).
		Relation("Jobs").
		Scan(ctx)
	if err != nil {
		p.logger.Error("", zap.Error(err))
	}
	return proj, err
}

func NewProjectRepository(db *sql.DB, logger *zap.Logger) ports.ProjectRepository {
	bunx := bun.NewDB(db, pgdialect.New())
	return &projectRepository{
		db,
		bunx,
		logger,
	}
}
