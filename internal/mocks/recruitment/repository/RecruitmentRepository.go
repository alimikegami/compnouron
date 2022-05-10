// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	entity "github.com/alimikegami/compnouron/internal/recruitment/entity"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RecruitmentRepository is an autogenerated mock type for the RecruitmentRepository type
type RecruitmentRepository struct {
	mock.Mock
}

// CreateRecruitment provides a mock function with given fields: recruitment
func (_m *RecruitmentRepository) CreateRecruitment(recruitment entity.Recruitment) error {
	ret := _m.Called(recruitment)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Recruitment) error); ok {
		r0 = rf(recruitment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRecruitmentApplication provides a mock function with given fields: recruitmentApplication
func (_m *RecruitmentRepository) CreateRecruitmentApplication(recruitmentApplication entity.RecruitmentApplication) error {
	ret := _m.Called(recruitmentApplication)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.RecruitmentApplication) error); ok {
		r0 = rf(recruitmentApplication)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRecruitment provides a mock function with given fields: recruitment
func (_m *RecruitmentRepository) UpdateRecruitment(recruitment entity.Recruitment) error {
	ret := _m.Called(recruitment)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Recruitment) error); ok {
		r0 = rf(recruitment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRecruitmentRepository creates a new instance of RecruitmentRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRecruitmentRepository(t testing.TB) *RecruitmentRepository {
	mock := &RecruitmentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
