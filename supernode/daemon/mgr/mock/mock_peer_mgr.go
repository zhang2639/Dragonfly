// Code generated by MockGen. DO NOT EDIT.
// Source: supernode/daemon/mgr/peer_mgr.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/Dragonfly/apis/types"
	util "github.com/Dragonfly/supernode/daemon/util"
)

// MockPeerMgr is a mock of PeerMgr interface
type MockPeerMgr struct {
	ctrl     *gomock.Controller
	recorder *MockPeerMgrMockRecorder
}

// MockPeerMgrMockRecorder is the mock recorder for MockPeerMgr
type MockPeerMgrMockRecorder struct {
	mock *MockPeerMgr
}

// NewMockPeerMgr creates a new mock instance
func NewMockPeerMgr(ctrl *gomock.Controller) *MockPeerMgr {
	mock := &MockPeerMgr{ctrl: ctrl}
	mock.recorder = &MockPeerMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeerMgr) EXPECT() *MockPeerMgrMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockPeerMgr) Register(ctx context.Context, peerCreateRequest *types.PeerCreateRequest) (*types.PeerCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, peerCreateRequest)
	ret0, _ := ret[0].(*types.PeerCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockPeerMgrMockRecorder) Register(ctx, peerCreateRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPeerMgr)(nil).Register), ctx, peerCreateRequest)
}

// DeRegister mocks base method
func (m *MockPeerMgr) DeRegister(ctx context.Context, peerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeRegister", ctx, peerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeRegister indicates an expected call of DeRegister
func (mr *MockPeerMgrMockRecorder) DeRegister(ctx, peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeRegister", reflect.TypeOf((*MockPeerMgr)(nil).DeRegister), ctx, peerID)
}

// Get mocks base method
func (m *MockPeerMgr) Get(ctx context.Context, peerID string) (*types.PeerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, peerID)
	ret0, _ := ret[0].(*types.PeerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPeerMgrMockRecorder) Get(ctx, peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPeerMgr)(nil).Get), ctx, peerID)
}

// GetAllPeerIDs mocks base method
func (m *MockPeerMgr) GetAllPeerIDs(ctx context.Context) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPeerIDs", ctx)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetAllPeerIDs indicates an expected call of GetAllPeerIDs
func (mr *MockPeerMgrMockRecorder) GetAllPeerIDs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPeerIDs", reflect.TypeOf((*MockPeerMgr)(nil).GetAllPeerIDs), ctx)
}

// List mocks base method
func (m *MockPeerMgr) List(ctx context.Context, filter *util.PageFilter) ([]*types.PeerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, filter)
	ret0, _ := ret[0].([]*types.PeerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPeerMgrMockRecorder) List(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPeerMgr)(nil).List), ctx, filter)
}
