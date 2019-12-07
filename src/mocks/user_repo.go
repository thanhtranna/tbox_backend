package mocks

import (
	"tbox_backend/src/entity"

	mock "github.com/stretchr/testify/mock"
)

type IUserRepo struct {
	mock.Mock
}

func (_m *IUserRepo) FindByPhoneNumber(phoneNumber string) (entity.User, error) {
	ret := _m.Called(phoneNumber)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(phoneNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phoneNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *IUserRepo) FindByUserID(userID string) (entity.User, error) {
	ret := _m.Called(userID)
	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *IUserRepo) UpdateVerifyPhoneNumber(input entity.User) error {
	ret := _m.Called(input)
	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
