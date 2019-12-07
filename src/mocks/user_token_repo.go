package mocks

import (
	"tbox_backend/src/entity"

	mock "github.com/stretchr/testify/mock"
)

type IUserTokenRepo struct {
	mock.Mock
}

func (_m *IUserTokenRepo) Save(input entity.UserToken) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.UserToken) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *IUserTokenRepo) GetTokenByUserID(userID string) (entity.UserToken, error) {
	ret := _m.Called(userID)

	var r0 entity.UserToken
	if rf, ok := ret.Get(0).(func(string) entity.UserToken); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.UserToken)
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
