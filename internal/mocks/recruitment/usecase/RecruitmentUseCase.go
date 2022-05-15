// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	dto "github.com/alimikegami/compnouron/internal/recruitment/dto"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RecruitmentUseCase is an autogenerated mock type for the RecruitmentUseCase type
type RecruitmentUseCase struct {
	mock.Mock
}

// AcceptRecruitmentApplication provides a mock function with given fields: id
func (_m *RecruitmentUseCase) AcceptRecruitmentApplication(id uint) error {
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
func (_m *RecruitmentUseCase) CloseRecruitmentApplicationPeriod(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRecruitment provides a mock function with given fields: recruitmentRequest
func (_m *RecruitmentUseCase) CreateRecruitment(recruitmentRequest dto.RecruitmentRequest) error {
	ret := _m.Called(recruitmentRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.RecruitmentRequest) error); ok {
		r0 = rf(recruitmentRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRecruitmentApplication provides a mock function with given fields: recruitmentApplication, userID
func (_m *RecruitmentUseCase) CreateRecruitmentApplication(recruitmentApplication dto.RecruitmentApplicationRequest, userID uint) error {
	ret := _m.Called(recruitmentApplication, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.RecruitmentApplicationRequest, uint) error); ok {
		r0 = rf(recruitmentApplication, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRecruitmentByID provides a mock function with given fields: id
func (_m *RecruitmentUseCase) DeleteRecruitmentByID(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRecruitmentByID provides a mock function with given fields: id
func (_m *RecruitmentUseCase) GetRecruitmentByID(id uint) (dto.RecruitmentResponse, error) {
	ret := _m.Called(id)

	var r0 dto.RecruitmentResponse
	if rf, ok := ret.Get(0).(func(uint) dto.RecruitmentResponse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(dto.RecruitmentResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecruitmentByUserID provides a mock function with given fields: id
func (_m *RecruitmentUseCase) GetRecruitmentByUserID(id uint) (dto.RecruitmentsResponse, error) {
	ret := _m.Called(id)

	var r0 dto.RecruitmentsResponse
	if rf, ok := ret.Get(0).(func(uint) dto.RecruitmentsResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dto.RecruitmentsResponse)
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

// GetRecruitmentDetailsByID provides a mock function with given fields: id
func (_m *RecruitmentUseCase) GetRecruitmentDetailsByID(id uint) (dto.RecruitmentDetailsResponse, error) {
	ret := _m.Called(id)

	var r0 dto.RecruitmentDetailsResponse
	if rf, ok := ret.Get(0).(func(uint) dto.RecruitmentDetailsResponse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(dto.RecruitmentDetailsResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenRecruitmentApplicationPeriod provides a mock function with given fields: id
func (_m *RecruitmentUseCase) OpenRecruitmentApplicationPeriod(id uint) error {
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
func (_m *RecruitmentUseCase) RejectRecruitmentApplication(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRecruitment provides a mock function with given fields: recruitmentRequest, id
func (_m *RecruitmentUseCase) UpdateRecruitment(recruitmentRequest dto.RecruitmentRequest, id uint) error {
	ret := _m.Called(recruitmentRequest, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.RecruitmentRequest, uint) error); ok {
		r0 = rf(recruitmentRequest, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRecruitmentUseCase creates a new instance of RecruitmentUseCase. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRecruitmentUseCase(t testing.TB) *RecruitmentUseCase {
	mock := &RecruitmentUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
