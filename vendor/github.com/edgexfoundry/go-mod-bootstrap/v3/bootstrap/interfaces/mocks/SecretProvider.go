// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// SecretProvider is an autogenerated mock type for the SecretProvider type
type SecretProvider struct {
	mock.Mock
}

// DeregisterSecretUpdatedCallback provides a mock function with given fields: secretName
func (_m *SecretProvider) DeregisterSecretUpdatedCallback(secretName string) {
	_m.Called(secretName)
}

// GetSecret provides a mock function with given fields: secretName, keys
func (_m *SecretProvider) GetSecret(secretName string, keys ...string) (map[string]string, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, secretName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...string) (map[string]string, error)); ok {
		return rf(secretName, keys...)
	}
	if rf, ok := ret.Get(0).(func(string, ...string) map[string]string); ok {
		r0 = rf(secretName, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(secretName, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasSecret provides a mock function with given fields: secretName
func (_m *SecretProvider) HasSecret(secretName string) (bool, error) {
	ret := _m.Called(secretName)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(secretName)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(secretName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(secretName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecretNames provides a mock function with given fields:
func (_m *SecretProvider) ListSecretNames() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterSecretUpdatedCallback provides a mock function with given fields: secretName, callback
func (_m *SecretProvider) RegisterSecretUpdatedCallback(secretName string, callback func(string)) error {
	ret := _m.Called(secretName, callback)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(string)) error); ok {
		r0 = rf(secretName, callback)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SecretsLastUpdated provides a mock function with given fields:
func (_m *SecretProvider) SecretsLastUpdated() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// StoreSecret provides a mock function with given fields: secretName, secrets
func (_m *SecretProvider) StoreSecret(secretName string, secrets map[string]string) error {
	ret := _m.Called(secretName, secrets)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]string) error); ok {
		r0 = rf(secretName, secrets)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSecretProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewSecretProvider creates a new instance of SecretProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSecretProvider(t mockConstructorTestingTNewSecretProvider) *SecretProvider {
	mock := &SecretProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
