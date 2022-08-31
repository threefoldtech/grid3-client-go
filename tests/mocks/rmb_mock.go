// Code generated by MockGen. DO NOT EDIT.
// Source: /home/alaa/zos/pkg/rmb/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)


// RMBMockClient is a mock of Client interface.
type RMBMockClient struct {
	ctrl     *gomock.Controller
	recorder *RMBMockClientMockRecorder
}

// RMBMockClientMockRecorder is the mock recorder for RMBMockClient.
type RMBMockClientMockRecorder struct {
	mock *RMBMockClient
}

// NewRMBMockClient creates a new mock instance.
func NewRMBMockClient(ctrl *gomock.Controller) *RMBMockClient {
	mock := &RMBMockClient{ctrl: ctrl}
	mock.recorder = &RMBMockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *RMBMockClient) EXPECT() *RMBMockClientMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *RMBMockClient) Call(ctx context.Context, twin uint32, fn string, data, result interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, twin, fn, data, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// Call indicates an expected call of Call.
func (mr *RMBMockClientMockRecorder) Call(ctx, twin, fn, data, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*RMBMockClient)(nil).Call), ctx, twin, fn, data, result)
}
