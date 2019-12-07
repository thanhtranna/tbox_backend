package mocks

import mock "github.com/stretchr/testify/mock"

type IRedisDataLayer struct {
	mock.Mock
}

func (_m *IRedisDataLayer) GetString(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(0)
	}

	return r0, r1
}

func (_m *IRedisDataLayer) SetString(key string, expire int, value string) error {
	ret := _m.Called(key, expire, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, string) error); ok {
		r0 = rf(key, expire, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *IRedisDataLayer) DeleteKey(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
