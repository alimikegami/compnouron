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

// AcceptRecruitmentApplication provides a mock function with given fields: id
func (_m *RecruitmentRepository) AcceptRecruitmentApplication(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseRecruitmentApplicationPeriod provides a mock function with given fields: id
func (_m *RecruitmentRepository) CloseRecruitmentApplicationPeriod(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// DeleteRecruitmentByID provides a mock function with given fields: id
func (_m *RecruitmentRepository) DeleteRecruitmentByID(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRecruitmentApplicationByID provides a mock function with given fields: id
func (_m *RecruitmentRepository) GetRecruitmentApplicationByID(id uint) (entity.RecruitmentApplication, error) {
	ret := _m.Called(id)

	var r0 entity.RecruitmentApplication
	if rf, ok := ret.Get(0).(func(uint) entity.RecruitmentApplication); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.RecruitmentApplication)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecruitmentApplicationByRecruitmentID provides a mock function with given fields: id
func (_m *RecruitmentRepository) GetRecruitmentApplicationByRecruitmentID(id uint) ([]entity.RecruitmentApplication, error) {
	ret := _m.Called(id)

	var r0 []entity.RecruitmentApplication
	if rf, ok := ret.Get(0).(func(uint) []entity.RecruitmentApplication); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.RecruitmentApplication)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecruitmentByID provides a mock function with given fields: id
func (_m *RecruitmentRepository) GetRecruitmentByID(id uint) (entity.Recruitment, error) {
	ret := _m.Called(id)

	var r0 entity.Recruitment
	if rf, ok := ret.Get(0).(func(uint) entity.Recruitment); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Recruitment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecruitmentByTeamID provides a mock function with given fields: id
func (_m *RecruitmentRepository) GetRecruitmentByTeamID(id uint) ([]entity.Recruitment, error) {
	ret := _m.Called(id)

	var r0 []entity.Recruitment
	if rf, ok := ret.Get(0).(func(uint) []entity.Recruitment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Recruitment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecruitments provides a mock function with given fields: limit, offset
func (_m *RecruitmentRepository) GetRecruitments(limit int, offset int) ([]entity.Recruitment, error) {
	ret := _m.Called(limit, offset)

	var r0 []entity.Recruitment
	if rf, ok := ret.Get(0).(func(int, int) []entity.Recruitment); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Recruitment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenRecruitmentApplicationPeriod provides a mock function with given fields: id
func (_m *RecruitmentRepository) OpenRecruitmentApplicationPeriod(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RejectRecruitmentApplication provides a mock function with given fields: id
func (_m *RecruitmentRepository) RejectRecruitmentApplication(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchRecruitment provides a mock function with given fields: limit, offset, keyword
func (_m *RecruitmentRepository) SearchRecruitment(limit int, offset int, keyword string) ([]entity.Recruitment, error) {
	ret := _m.Called(limit, offset, keyword)

	var r0 []entity.Recruitment
	if rf, ok := ret.Get(0).(func(int, int, string) []entity.Recruitment); ok {
		r0 = rf(limit, offset, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Recruitment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(limit, offset, keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
