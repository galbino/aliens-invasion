// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/galbino/alien-assignment/internal/domain (interfaces: Alien)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/galbino/alien-assignment/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockAlien is a mock of Alien interface.
type MockAlien struct {
	ctrl     *gomock.Controller
	recorder *MockAlienMockRecorder
}

// MockAlienMockRecorder is the mock recorder for MockAlien.
type MockAlienMockRecorder struct {
	mock *MockAlien
}

// NewMockAlien creates a new mock instance.
func NewMockAlien(ctrl *gomock.Controller) *MockAlien {
	mock := &MockAlien{ctrl: ctrl}
	mock.recorder = &MockAlienMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlien) EXPECT() *MockAlienMockRecorder {
	return m.recorder
}

// Location mocks base method.
func (m *MockAlien) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockAlienMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockAlien)(nil).Location))
}

// String mocks base method.
func (m *MockAlien) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockAlienMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockAlien)(nil).String))
}

// Walk mocks base method.
func (m *MockAlien) Walk(arg0 domain.Cities) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Walk", arg0)
}

// Walk indicates an expected call of Walk.
func (mr *MockAlienMockRecorder) Walk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockAlien)(nil).Walk), arg0)
}
