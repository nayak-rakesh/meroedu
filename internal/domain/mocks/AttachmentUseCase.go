// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/meroedu/meroedu/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// AttachmentUseCase is an autogenerated mock type for the AttachmentUseCase type
type AttachmentUseCase struct {
	mock.Mock
}

// CreateAttachment provides a mock function with given fields: ctx, attachment
func (_m *AttachmentUseCase) CreateAttachment(ctx context.Context, attachment domain.Attachment) (*domain.Attachment, error) {
	ret := _m.Called(ctx, attachment)

	var r0 *domain.Attachment
	if rf, ok := ret.Get(0).(func(context.Context, domain.Attachment) *domain.Attachment); ok {
		r0 = rf(ctx, attachment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Attachment) error); ok {
		r1 = rf(ctx, attachment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadAttachment provides a mock function with given fields: ctx, fileName
func (_m *AttachmentUseCase) DownloadAttachment(ctx context.Context, fileName string) (string, error) {
	ret := _m.Called(ctx, fileName)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, fileName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttachmentByCourse provides a mock function with given fields: ctx, courseID
func (_m *AttachmentUseCase) GetAttachmentByCourse(ctx context.Context, courseID int64) ([]domain.Attachment, error) {
	ret := _m.Called(ctx, courseID)

	var r0 []domain.Attachment
	if rf, ok := ret.Get(0).(func(context.Context, int64) []domain.Attachment); ok {
		r0 = rf(ctx, courseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, courseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
