// Code generated by MockGen. DO NOT EDIT.
// Source: ./subi/substrate_manager.go

// Package mock_subi is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	subi "github.com/threefoldtech/grid3-go/subi"
	"github.com/threefoldtech/substrate-client"
)

// MockManagerInterface is a mock of ManagerInterface interface.
type MockManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockManagerInterfaceMockRecorder
}

// MockManagerInterfaceMockRecorder is the mock recorder for MockManagerInterface.
type MockManagerInterfaceMockRecorder struct {
	mock *MockManagerInterface
}

// NewMockManagerInterface creates a new mock instance.
func NewMockManagerInterface(ctrl *gomock.Controller) *MockManagerInterface {
	mock := &MockManagerInterface{ctrl: ctrl}
	mock.recorder = &MockManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagerInterface) EXPECT() *MockManagerInterfaceMockRecorder {
	return m.recorder
}

// Raw mocks base method.
func (m *MockManagerInterface) Raw() (substrate.Conn, substrate.Meta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(substrate.Conn)
	ret1, _ := ret[1].(substrate.Meta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Raw indicates an expected call of Raw.
func (mr *MockManagerInterfaceMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockManagerInterface)(nil).Raw))
}

// Substrate mocks base method.
func (m *MockManagerInterface) Substrate() (*substrate.Substrate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Substrate")
	ret0, _ := ret[0].(*substrate.Substrate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Substrate indicates an expected call of Substrate.
func (mr *MockManagerInterfaceMockRecorder) Substrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Substrate", reflect.TypeOf((*MockManagerInterface)(nil).Substrate))
}

// SubstrateExt mocks base method.
func (m *MockManagerInterface) SubstrateExt() (subi.SubstrateExt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubstrateExt")
	ret0, _ := ret[0].(subi.SubstrateExt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubstrateExt indicates an expected call of SubstrateExt.
func (mr *MockManagerInterfaceMockRecorder) SubstrateExt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubstrateExt", reflect.TypeOf((*MockManagerInterface)(nil).SubstrateExt))
}

// MockSubstrate is a mock of Substrate interface.
type MockSubstrate struct {
	ctrl     *gomock.Controller
	recorder *MockSubstrateMockRecorder
}

// MockSubstrateMockRecorder is the mock recorder for MockSubstrate.
type MockSubstrateMockRecorder struct {
	mock *MockSubstrate
}

// NewMockSubstrate creates a new mock instance.
func NewMockSubstrate(ctrl *gomock.Controller) *MockSubstrate {
	mock := &MockSubstrate{ctrl: ctrl}
	mock.recorder = &MockSubstrateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubstrate) EXPECT() *MockSubstrateMockRecorder {
	return m.recorder
}

// CancelContract mocks base method.
func (m *MockSubstrate) CancelContract(identity substrate.Identity, contractID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelContract", identity, contractID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelContract indicates an expected call of CancelContract.
func (mr *MockSubstrateMockRecorder) CancelContract(identity, contractID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelContract", reflect.TypeOf((*MockSubstrate)(nil).CancelContract), identity, contractID)
}

// Close mocks base method.
func (m *MockSubstrate) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSubstrateMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSubstrate)(nil).Close))
}

// CreateNodeContract mocks base method.
func (m *MockSubstrate) CreateNodeContract(identity substrate.Identity, node uint32, body, hash string, publicIPs uint32, solutionProviderID *uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNodeContract", identity, node, body, hash, publicIPs, solutionProviderID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNodeContract indicates an expected call of CreateNodeContract.
func (mr *MockSubstrateMockRecorder) CreateNodeContract(identity, node, body, hash, publicIPs, solutionProviderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNodeContract", reflect.TypeOf((*MockSubstrate)(nil).CreateNodeContract), identity, node, body, hash, publicIPs, solutionProviderID)
}

// GetTwinByPubKey mocks base method.
func (m *MockSubstrate) GetTwinByPubKey(pk []byte) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTwinByPubKey", pk)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTwinByPubKey indicates an expected call of GetTwinByPubKey.
func (mr *MockSubstrateMockRecorder) GetTwinByPubKey(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTwinByPubKey", reflect.TypeOf((*MockSubstrate)(nil).GetTwinByPubKey), pk)
}

// UpdateNodeContract mocks base method.
func (m *MockSubstrate) UpdateNodeContract(identity substrate.Identity, contract uint64, body, hash string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNodeContract", identity, contract, body, hash)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNodeContract indicates an expected call of UpdateNodeContract.
func (mr *MockSubstrateMockRecorder) UpdateNodeContract(identity, contract, body, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNodeContract", reflect.TypeOf((*MockSubstrate)(nil).UpdateNodeContract), identity, contract, body, hash)
}

// MockSubstrateExt is a mock of SubstrateExt interface.
type MockSubstrateExt struct {
	ctrl     *gomock.Controller
	recorder *MockSubstrateExtMockRecorder
}

// MockSubstrateExtMockRecorder is the mock recorder for MockSubstrateExt.
type MockSubstrateExtMockRecorder struct {
	mock *MockSubstrateExt
}

// NewMockSubstrateExt creates a new mock instance.
func NewMockSubstrateExt(ctrl *gomock.Controller) *MockSubstrateExt {
	mock := &MockSubstrateExt{ctrl: ctrl}
	mock.recorder = &MockSubstrateExtMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubstrateExt) EXPECT() *MockSubstrateExtMockRecorder {
	return m.recorder
}

// CancelContract mocks base method.
func (m *MockSubstrateExt) CancelContract(identity substrate.Identity, contractID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelContract", identity, contractID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelContract indicates an expected call of CancelContract.
func (mr *MockSubstrateExtMockRecorder) CancelContract(identity, contractID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelContract", reflect.TypeOf((*MockSubstrateExt)(nil).CancelContract), identity, contractID)
}

// Close mocks base method.
func (m *MockSubstrateExt) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSubstrateExtMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSubstrateExt)(nil).Close))
}

// CreateNameContract mocks base method.
func (m *MockSubstrateExt) CreateNameContract(identity substrate.Identity, name string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNameContract", identity, name)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNameContract indicates an expected call of CreateNameContract.
func (mr *MockSubstrateExtMockRecorder) CreateNameContract(identity, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNameContract", reflect.TypeOf((*MockSubstrateExt)(nil).CreateNameContract), identity, name)
}

// CreateNodeContract mocks base method.
func (m *MockSubstrateExt) CreateNodeContract(identity substrate.Identity, node uint32, body, hash string, publicIPs uint32, solutionProviderID *uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNodeContract", identity, node, body, hash, publicIPs, solutionProviderID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNodeContract indicates an expected call of CreateNodeContract.
func (mr *MockSubstrateExtMockRecorder) CreateNodeContract(identity, node, body, hash, publicIPs, solutionProviderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNodeContract", reflect.TypeOf((*MockSubstrateExt)(nil).CreateNodeContract), identity, node, body, hash, publicIPs, solutionProviderID)
}

// DeleteInvalidContracts mocks base method.
func (m *MockSubstrateExt) DeleteInvalidContracts(contracts map[uint32]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInvalidContracts", contracts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInvalidContracts indicates an expected call of DeleteInvalidContracts.
func (mr *MockSubstrateExtMockRecorder) DeleteInvalidContracts(contracts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInvalidContracts", reflect.TypeOf((*MockSubstrateExt)(nil).DeleteInvalidContracts), contracts)
}

// EnsureContractCanceled mocks base method.
func (m *MockSubstrateExt) EnsureContractCanceled(identity substrate.Identity, contractID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureContractCanceled", identity, contractID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureContractCanceled indicates an expected call of EnsureContractCanceled.
func (mr *MockSubstrateExtMockRecorder) EnsureContractCanceled(identity, contractID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureContractCanceled", reflect.TypeOf((*MockSubstrateExt)(nil).EnsureContractCanceled), identity, contractID)
}

// GetAccount mocks base method.
func (m *MockSubstrateExt) GetAccount(identity substrate.Identity) (substrate.AccountInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", identity)
	ret0, _ := ret[0].(substrate.AccountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockSubstrateExtMockRecorder) GetAccount(identity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockSubstrateExt)(nil).GetAccount), identity)
}

// GetBalance mocks base method.
func (m *MockSubstrateExt) GetBalance(identity substrate.Identity) (balance substrate.Balance, err error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", identity)
	ret0, _ := ret[0].(substrate.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockSubstrateExtMockRecorder) GetBalance(identity substrate.Identity) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockSubstrateExt)(nil).GetBalance), identity)
}

// GetContract mocks base method.
func (m *MockSubstrateExt) GetContract(id uint64) (subi.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract", id)
	ret0, _ := ret[0].(subi.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract.
func (mr *MockSubstrateExtMockRecorder) GetContract(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockSubstrateExt)(nil).GetContract), id)
}

// GetContractIDByNameRegistration mocks base method.
func (m *MockSubstrateExt) GetContractIDByNameRegistration(name string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractIDByNameRegistration", name)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractIDByNameRegistration indicates an expected call of GetContractIDByNameRegistration.
func (mr *MockSubstrateExtMockRecorder) GetContractIDByNameRegistration(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractIDByNameRegistration", reflect.TypeOf((*MockSubstrateExt)(nil).GetContractIDByNameRegistration), name)
}

// GetNodeTwin mocks base method.
func (m *MockSubstrateExt) GetNodeTwin(id uint32) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeTwin", id)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeTwin indicates an expected call of GetNodeTwin.
func (mr *MockSubstrateExtMockRecorder) GetNodeTwin(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeTwin", reflect.TypeOf((*MockSubstrateExt)(nil).GetNodeTwin), id)
}

// GetTwinByPubKey mocks base method.
func (m *MockSubstrateExt) GetTwinByPubKey(pk []byte) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTwinByPubKey", pk)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTwinByPubKey indicates an expected call of GetTwinByPubKey.
func (mr *MockSubstrateExtMockRecorder) GetTwinByPubKey(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTwinByPubKey", reflect.TypeOf((*MockSubstrateExt)(nil).GetTwinByPubKey), pk)
}

// GetTwinPK mocks base method.
func (m *MockSubstrateExt) GetTwinPK(twinID uint32) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTwinPK", twinID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTwinPK indicates an expected call of GetTwinPK.
func (mr *MockSubstrateExtMockRecorder) GetTwinPK(twinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTwinPK", reflect.TypeOf((*MockSubstrateExt)(nil).GetTwinPK), twinID)
}

// InvalidateNameContract mocks base method.
func (m *MockSubstrateExt) InvalidateNameContract(ctx context.Context, identity substrate.Identity, contractID uint64, name string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateNameContract", ctx, identity, contractID, name)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvalidateNameContract indicates an expected call of InvalidateNameContract.
func (mr *MockSubstrateExtMockRecorder) InvalidateNameContract(ctx, identity, contractID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateNameContract", reflect.TypeOf((*MockSubstrateExt)(nil).InvalidateNameContract), ctx, identity, contractID, name)
}

// IsValidContract mocks base method.
func (m *MockSubstrateExt) IsValidContract(contractID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidContract", contractID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValidContract indicates an expected call of IsValidContract.
func (mr *MockSubstrateExtMockRecorder) IsValidContract(contractID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidContract", reflect.TypeOf((*MockSubstrateExt)(nil).IsValidContract), contractID)
}

// UpdateNodeContract mocks base method.
func (m *MockSubstrateExt) UpdateNodeContract(identity substrate.Identity, contract uint64, body, hash string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNodeContract", identity, contract, body, hash)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNodeContract indicates an expected call of UpdateNodeContract.
func (mr *MockSubstrateExtMockRecorder) UpdateNodeContract(identity, contract, body, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNodeContract", reflect.TypeOf((*MockSubstrateExt)(nil).UpdateNodeContract), identity, contract, body, hash)
}
