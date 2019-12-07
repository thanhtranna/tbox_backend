package mocks

import (
	"tbox_backend/src/entity"

	mock "github.com/stretchr/testify/mock"
)

type IUserOTPRepo struct {
	mock.Mock
}

func (_m *IUserOTPRepo) GetOTPByUserID(userID string) (entity.UserOTP, error) {
	ret := _m.Called(userID)

	var r0 entity.UserOTP
	if rf, ok := ret.Get(0).(func(string) entity.UserOTP); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.UserOTP)
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

func (_m *IUserOTPRepo) CheckSendedOTP(userID string) (string, error) {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
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

func (_m *IUserOTPRepo) Save(input entity.UserOTP) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.UserOTP) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *IUserOTPRepo) CacheOTPWithUser(userID string, OTP string) error {
	ret := _m.Called(userID, OTP)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, OTP)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
