// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIDer is a mock of IDer interface
type MockIDer struct {
	ctrl     *gomock.Controller
	recorder *MockIDerMockRecorder
}

// MockIDerMockRecorder is the mock recorder for MockIDer
type MockIDerMockRecorder struct {
	mock *MockIDer
}

// NewMockIDer creates a new mock instance
func NewMockIDer(ctrl *gomock.Controller) *MockIDer {
	mock := &MockIDer{ctrl: ctrl}
	mock.recorder = &MockIDerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDer) EXPECT() *MockIDerMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockIDer) ID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockIDerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockIDer)(nil).ID))
}
