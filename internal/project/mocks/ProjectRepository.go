// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/monopeelz/linear-avocado/internal/project/models"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ProjectRepository is an autogenerated mock type for the ProjectRepository type
type ProjectRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, project
func (_m *ProjectRepository) Create(ctx context.Context, project models.Project) (models.Project, error) {
	ret := _m.Called(ctx, project)

	var r0 models.Project
	if rf, ok := ret.Get(0).(func(context.Context, models.Project) models.Project); ok {
		r0 = rf(ctx, project)
	} else {
		r0 = ret.Get(0).(models.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Project) error); ok {
		r1 = rf(ctx, project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBytProjectID provides a mock function with given fields: ctx, id
func (_m *ProjectRepository) DeleteBytProjectID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *ProjectRepository) GetAll(ctx context.Context) ([]models.Project, error) {
	ret := _m.Called(ctx)

	var r0 []models.Project
	if rf, ok := ret.Get(0).(func(context.Context) []models.Project); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ProjectRepository) GetByID(ctx context.Context, id string) (models.Project, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Project
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Project); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobs provides a mock function with given fields: ctx, id
func (_m *ProjectRepository) GetJobs(ctx context.Context, id string) ([]models.Job, error) {
	ret := _m.Called(ctx, id)

	var r0 []models.Job
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Job); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, project
func (_m *ProjectRepository) Update(ctx context.Context, project models.UpdateProject) (models.Project, error) {
	ret := _m.Called(ctx, project)

	var r0 models.Project
	if rf, ok := ret.Get(0).(func(context.Context, models.Project) models.Project); ok {
		r0 = rf(ctx, project)
	} else {
		r0 = ret.Get(0).(models.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Project) error); ok {
		r1 = rf(ctx, project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProjectRepository creates a new instance of ProjectRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewProjectRepository(t testing.TB) *ProjectRepository {
	mock := &ProjectRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
