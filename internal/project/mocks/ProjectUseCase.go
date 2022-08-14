// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/monopeelz/linear-avocado/internal/project/models"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ProjectUseCase is an autogenerated mock type for the ProjectUseCase type
type ProjectUseCase struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *ProjectUseCase) All(ctx context.Context) ([]models.Project, error) {
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

// Create provides a mock function with given fields: ctx, i
func (_m *ProjectUseCase) Create(ctx context.Context, i models.Project) (models.Project, error) {
	ret := _m.Called(ctx, i)

	var r0 models.Project
	if rf, ok := ret.Get(0).(func(context.Context, models.Project) models.Project); ok {
		r0 = rf(ctx, i)
	} else {
		r0 = ret.Get(0).(models.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Project) error); ok {
		r1 = rf(ctx, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, i
func (_m *ProjectUseCase) Delete(ctx context.Context, i string) error {
	ret := _m.Called(ctx, i)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, i
func (_m *ProjectUseCase) Get(ctx context.Context, i string) (models.Project, error) {
	ret := _m.Called(ctx, i)

	var r0 models.Project
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Project); ok {
		r0 = rf(ctx, i)
	} else {
		r0 = ret.Get(0).(models.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJob provides a mock function with given fields: ctx, i
func (_m *ProjectUseCase) GetJob(ctx context.Context, i string) ([]models.Job, error) {
	ret := _m.Called(ctx, i)

	var r0 []models.Job
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Job); ok {
		r0 = rf(ctx, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scan provides a mock function with given fields: ctx, i
func (_m *ProjectUseCase) Scan(ctx context.Context, i string) error {
	ret := _m.Called(ctx, i)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, i
func (_m *ProjectUseCase) Update(ctx context.Context, i models.UpdateProject) (models.Project, error) {
	ret := _m.Called(ctx, i)

	var r0 models.Project
	if rf, ok := ret.Get(0).(func(context.Context, models.Project) models.Project); ok {
		r0 = rf(ctx, i)
	} else {
		r0 = ret.Get(0).(models.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Project) error); ok {
		r1 = rf(ctx, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProjectUseCase creates a new instance of ProjectUseCase. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewProjectUseCase(t testing.TB) *ProjectUseCase {
	mock := &ProjectUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
