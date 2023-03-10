// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	mentees "immersiveApp/features/mentees"

	mock "github.com/stretchr/testify/mock"
)

// MenteeDataInterface is an autogenerated mock type for the MenteeDataInterface type
type MenteeDataInterface struct {
	mock.Mock
}

// Destroy provides a mock function with given fields: id
func (_m *MenteeDataInterface) Destroy(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: menteeEntity, id
func (_m *MenteeDataInterface) Edit(menteeEntity mentees.MenteeEntity, id uint) error {
	ret := _m.Called(menteeEntity, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(mentees.MenteeEntity, uint) error); ok {
		r0 = rf(menteeEntity, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields:
func (_m *MenteeDataInterface) SelectAll() ([]mentees.MenteeEntity, error) {
	ret := _m.Called()

	var r0 []mentees.MenteeEntity
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]mentees.MenteeEntity, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []mentees.MenteeEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentees.MenteeEntity)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *MenteeDataInterface) SelectById(id uint) (mentees.MenteeEntity, error) {
	ret := _m.Called(id)

	var r0 mentees.MenteeEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (mentees.MenteeEntity, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) mentees.MenteeEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(mentees.MenteeEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectFeedbackById provides a mock function with given fields: id
func (_m *MenteeDataInterface) SelectFeedbackById(id uint) (interface{}, error) {
	ret := _m.Called(id)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (interface{}, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) interface{}); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: menteeEntity
func (_m *MenteeDataInterface) Store(menteeEntity mentees.MenteeEntity) (uint, error) {
	ret := _m.Called(menteeEntity)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(mentees.MenteeEntity) (uint, error)); ok {
		return rf(menteeEntity)
	}
	if rf, ok := ret.Get(0).(func(mentees.MenteeEntity) uint); ok {
		r0 = rf(menteeEntity)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(mentees.MenteeEntity) error); ok {
		r1 = rf(menteeEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMenteeDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenteeDataInterface creates a new instance of MenteeDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenteeDataInterface(t mockConstructorTestingTNewMenteeDataInterface) *MenteeDataInterface {
	mock := &MenteeDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
