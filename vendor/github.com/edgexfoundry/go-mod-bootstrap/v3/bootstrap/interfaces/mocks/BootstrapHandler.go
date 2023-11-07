// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	di "github.com/edgexfoundry/go-mod-bootstrap/v3/di"

	mock "github.com/stretchr/testify/mock"

	startup "github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/startup"

	sync "sync"
)

// BootstrapHandler is an autogenerated mock type for the BootstrapHandler type
type BootstrapHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, wg, startupTimer, dic
func (_m *BootstrapHandler) Execute(ctx context.Context, wg *sync.WaitGroup, startupTimer startup.Timer, dic *di.Container) bool {
	ret := _m.Called(ctx, wg, startupTimer, dic)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *sync.WaitGroup, startup.Timer, *di.Container) bool); ok {
		r0 = rf(ctx, wg, startupTimer, dic)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewBootstrapHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewBootstrapHandler creates a new instance of BootstrapHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBootstrapHandler(t mockConstructorTestingTNewBootstrapHandler) *BootstrapHandler {
	mock := &BootstrapHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
