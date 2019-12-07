package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

type IMgoDataLayer struct {
	mock.Mock
}

func (_m *IMgoDataLayer) Insert(input ...interface{}) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *IMgoDataLayer) Update(condition, something interface{}) error {
	ret := _m.Called(condition, something)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(condition, something)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *IMgoDataLayer) FindOne(condition, something interface{}) error {
	ret := _m.Called(condition, something)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(condition, something)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *IMgoDataLayer) Delete(condition interface{}) error {
	ret := _m.Called(condition)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(condition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *IMgoDataLayer) Upsert(condition, something interface{}) error {
	ret := _m.Called(condition, something)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(condition, something)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
