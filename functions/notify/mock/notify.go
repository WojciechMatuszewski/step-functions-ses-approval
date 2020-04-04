// Code generated by MockGen. DO NOT EDIT.
// Source: ../notify.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	emails "step-functions/platform/emails"
)

// MockSender is a mock of Sender interface
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockSender) Send(ctx context.Context, template emails.EmailTemplate, recipient string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, template, recipient)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockSenderMockRecorder) Send(ctx, template, recipient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSender)(nil).Send), ctx, template, recipient)
}

// MockEndpointer is a mock of Endpointer interface
type MockEndpointer struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointerMockRecorder
}

// MockEndpointerMockRecorder is the mock recorder for MockEndpointer
type MockEndpointerMockRecorder struct {
	mock *MockEndpointer
}

// NewMockEndpointer creates a new mock instance
func NewMockEndpointer(ctrl *gomock.Controller) *MockEndpointer {
	mock := &MockEndpointer{ctrl: ctrl}
	mock.recorder = &MockEndpointerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEndpointer) EXPECT() *MockEndpointerMockRecorder {
	return m.recorder
}

// NewDenyEndpoint mocks base method
func (m *MockEndpointer) NewDenyEndpoint(rootURL, token, candidateEmail string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDenyEndpoint", rootURL, token, candidateEmail)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDenyEndpoint indicates an expected call of NewDenyEndpoint
func (mr *MockEndpointerMockRecorder) NewDenyEndpoint(rootURL, token, candidateEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDenyEndpoint", reflect.TypeOf((*MockEndpointer)(nil).NewDenyEndpoint), rootURL, token, candidateEmail)
}

// NewApproveEndpoint mocks base method
func (m *MockEndpointer) NewApproveEndpoint(rootURL, token, candidateEmail string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewApproveEndpoint", rootURL, token, candidateEmail)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewApproveEndpoint indicates an expected call of NewApproveEndpoint
func (mr *MockEndpointerMockRecorder) NewApproveEndpoint(rootURL, token, candidateEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewApproveEndpoint", reflect.TypeOf((*MockEndpointer)(nil).NewApproveEndpoint), rootURL, token, candidateEmail)
}
