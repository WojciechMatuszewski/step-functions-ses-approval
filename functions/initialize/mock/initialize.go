// Code generated by MockGen. DO NOT EDIT.
// Source: ../initialize.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMachineStarter is a mock of MachineStarter interface
type MockMachineStarter struct {
	ctrl     *gomock.Controller
	recorder *MockMachineStarterMockRecorder
}

// MockMachineStarterMockRecorder is the mock recorder for MockMachineStarter
type MockMachineStarterMockRecorder struct {
	mock *MockMachineStarter
}

// NewMockMachineStarter creates a new mock instance
func NewMockMachineStarter(ctrl *gomock.Controller) *MockMachineStarter {
	mock := &MockMachineStarter{ctrl: ctrl}
	mock.recorder = &MockMachineStarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachineStarter) EXPECT() *MockMachineStarterMockRecorder {
	return m.recorder
}

// Kickoff mocks base method
func (m *MockMachineStarter) Kickoff(ctx context.Context, input, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kickoff", ctx, input, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Kickoff indicates an expected call of Kickoff
func (mr *MockMachineStarterMockRecorder) Kickoff(ctx, input, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kickoff", reflect.TypeOf((*MockMachineStarter)(nil).Kickoff), ctx, input, name)
}
