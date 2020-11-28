// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIDGen is a mock of IDGen interface
type MockIDGen struct {
	ctrl     *gomock.Controller
	recorder *MockIDGenMockRecorder
}

// MockIDGenMockRecorder is the mock recorder for MockIDGen
type MockIDGenMockRecorder struct {
	mock *MockIDGen
}

// NewMockIDGen creates a new mock instance
func NewMockIDGen(ctrl *gomock.Controller) *MockIDGen {
	mock := &MockIDGen{ctrl: ctrl}
	mock.recorder = &MockIDGenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDGen) EXPECT() *MockIDGenMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockIDGen) New() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// New indicates an expected call of New
func (mr *MockIDGenMockRecorder) New() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockIDGen)(nil).New))
}
