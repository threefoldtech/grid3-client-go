// Code generated by MockGen. DO NOT EDIT.
// Source: deployer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	subi "github.com/threefoldtech/grid3-go/subi"
	gridtypes "github.com/threefoldtech/zos/pkg/gridtypes"
)

// MockDeployer is a mock of Deployer interface.
type MockDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockDeployerMockRecorder
}

// MockDeployerMockRecorder is the mock recorder for MockDeployer.
type MockDeployerMockRecorder struct {
	mock *MockDeployer
}

// NewMockDeployer creates a new mock instance.
func NewMockDeployer(ctrl *gomock.Controller) *MockDeployer {
	mock := &MockDeployer{ctrl: ctrl}
	mock.recorder = &MockDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployer) EXPECT() *MockDeployerMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockDeployer) Deploy(ctx context.Context, sub subi.SubstrateExt, oldDeployments map[uint32]uint64, newDeployments map[uint32]gridtypes.Deployment) (map[uint32]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", ctx, sub, oldDeployments, newDeployments)
	ret0, _ := ret[0].(map[uint32]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deploy indicates an expected call of Deploy.
func (mr *MockDeployerMockRecorder) Deploy(ctx, sub, oldDeployments, newDeployments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDeployer)(nil).Deploy), ctx, sub, oldDeployments, newDeployments)
}

// GetDeploymentObjects mocks base method.
func (m *MockDeployer) GetDeploymentObjects(ctx context.Context, sub subi.SubstrateExt, dls map[uint32]uint64) (map[uint32]gridtypes.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentObjects", ctx, sub, dls)
	ret0, _ := ret[0].(map[uint32]gridtypes.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentObjects indicates an expected call of GetDeploymentObjects.
func (mr *MockDeployerMockRecorder) GetDeploymentObjects(ctx, sub, dls interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentObjects", reflect.TypeOf((*MockDeployer)(nil).GetDeploymentObjects), ctx, sub, dls)
}