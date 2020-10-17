// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/meroedu/meroedu/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// TagRepository is an autogenerated mock type for the TagRepository type
type TagRepository struct {
	mock.Mock
}

// CreateCourseTag provides a mock function with given fields: ctx, tagID, courseID
func (_m *TagRepository) CreateCourseTag(ctx context.Context, tagID int64, courseID int64) error {
	ret := _m.Called(ctx, tagID, courseID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, tagID, courseID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateLessonTag provides a mock function with given fields: ctx, tagID, lessonID
func (_m *TagRepository) CreateLessonTag(ctx context.Context, tagID int64, lessonID int64) error {
	ret := _m.Called(ctx, tagID, lessonID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, tagID, lessonID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTag provides a mock function with given fields: ctx, Tag
func (_m *TagRepository) CreateTag(ctx context.Context, Tag *domain.Tag) error {
	ret := _m.Called(ctx, Tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Tag) error); ok {
		r0 = rf(ctx, Tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCourseTag provides a mock function with given fields: ctx, tagID, courseID
func (_m *TagRepository) DeleteCourseTag(ctx context.Context, tagID int64, courseID int64) error {
	ret := _m.Called(ctx, tagID, courseID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, tagID, courseID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteLessonTag provides a mock function with given fields: ctx, tagID, lessonID
func (_m *TagRepository) DeleteLessonTag(ctx context.Context, tagID int64, lessonID int64) error {
	ret := _m.Called(ctx, tagID, lessonID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, tagID, lessonID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTag provides a mock function with given fields: ctx, id
func (_m *TagRepository) DeleteTag(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, searchQuery, start, limit
func (_m *TagRepository) GetAll(ctx context.Context, searchQuery string, start int, limit int) ([]domain.Tag, error) {
	ret := _m.Called(ctx, searchQuery, start, limit)

	var r0 []domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []domain.Tag); ok {
		r0 = rf(ctx, searchQuery, start, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) error); ok {
		r1 = rf(ctx, searchQuery, start, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TagRepository) GetByID(ctx context.Context, id int64) (*domain.Tag, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Tag); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *TagRepository) GetByName(ctx context.Context, name string) (*domain.Tag, error) {
	ret := _m.Called(ctx, name)

	var r0 *domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Tag); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCourseTags provides a mock function with given fields: ctx, courseID
func (_m *TagRepository) GetCourseTags(ctx context.Context, courseID int64) ([]domain.Tag, error) {
	ret := _m.Called(ctx, courseID)

	var r0 []domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, int64) []domain.Tag); ok {
		r0 = rf(ctx, courseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Tag)
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

// GetLessonTags provides a mock function with given fields: ctx, lessonID
func (_m *TagRepository) GetLessonTags(ctx context.Context, lessonID int64) ([]domain.Tag, error) {
	ret := _m.Called(ctx, lessonID)

	var r0 []domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, int64) []domain.Tag); ok {
		r0 = rf(ctx, lessonID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, lessonID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTag provides a mock function with given fields: ctx, Tag
func (_m *TagRepository) UpdateTag(ctx context.Context, Tag *domain.Tag) error {
	ret := _m.Called(ctx, Tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Tag) error); ok {
		r0 = rf(ctx, Tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
