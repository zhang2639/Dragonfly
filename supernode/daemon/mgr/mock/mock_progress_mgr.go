// Code generated by MockGen. DO NOT EDIT.
// Source: supernode/daemon/mgr/progress_mgr.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/Dragonfly/apis/types"
	syncmap "github.com/Dragonfly/pkg/syncmap"
	mgr "github.com/Dragonfly/supernode/daemon/mgr"
)

// MockProgressMgr is a mock of ProgressMgr interface
type MockProgressMgr struct {
	ctrl     *gomock.Controller
	recorder *MockProgressMgrMockRecorder
}

// MockProgressMgrMockRecorder is the mock recorder for MockProgressMgr
type MockProgressMgrMockRecorder struct {
	mock *MockProgressMgr
}

// NewMockProgressMgr creates a new mock instance
func NewMockProgressMgr(ctrl *gomock.Controller) *MockProgressMgr {
	mock := &MockProgressMgr{ctrl: ctrl}
	mock.recorder = &MockProgressMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProgressMgr) EXPECT() *MockProgressMgrMockRecorder {
	return m.recorder
}

// InitProgress mocks base method
func (m *MockProgressMgr) InitProgress(ctx context.Context, taskID, peerID, clientID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitProgress", ctx, taskID, peerID, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitProgress indicates an expected call of InitProgress
func (mr *MockProgressMgrMockRecorder) InitProgress(ctx, taskID, peerID, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitProgress", reflect.TypeOf((*MockProgressMgr)(nil).InitProgress), ctx, taskID, peerID, clientID)
}

// UpdateProgress mocks base method
func (m *MockProgressMgr) UpdateProgress(ctx context.Context, taskID, srcCID, srcPID, dstPID string, pieceNum, pieceStatus int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProgress", ctx, taskID, srcCID, srcPID, dstPID, pieceNum, pieceStatus)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProgress indicates an expected call of UpdateProgress
func (mr *MockProgressMgrMockRecorder) UpdateProgress(ctx, taskID, srcCID, srcPID, dstPID, pieceNum, pieceStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProgress", reflect.TypeOf((*MockProgressMgr)(nil).UpdateProgress), ctx, taskID, srcCID, srcPID, dstPID, pieceNum, pieceStatus)
}

// UpdateClientProgress mocks base method
func (m *MockProgressMgr) UpdateClientProgress(ctx context.Context, taskID, srcCID, dstPID string, pieceNum, pieceStatus int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClientProgress", ctx, taskID, srcCID, dstPID, pieceNum, pieceStatus)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClientProgress indicates an expected call of UpdateClientProgress
func (mr *MockProgressMgrMockRecorder) UpdateClientProgress(ctx, taskID, srcCID, dstPID, pieceNum, pieceStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClientProgress", reflect.TypeOf((*MockProgressMgr)(nil).UpdateClientProgress), ctx, taskID, srcCID, dstPID, pieceNum, pieceStatus)
}

// GetPieceProgressByCID mocks base method
func (m *MockProgressMgr) GetPieceProgressByCID(ctx context.Context, taskID, clientID, filter string) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPieceProgressByCID", ctx, taskID, clientID, filter)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieceProgressByCID indicates an expected call of GetPieceProgressByCID
func (mr *MockProgressMgrMockRecorder) GetPieceProgressByCID(ctx, taskID, clientID, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieceProgressByCID", reflect.TypeOf((*MockProgressMgr)(nil).GetPieceProgressByCID), ctx, taskID, clientID, filter)
}

// GetPeerIDsByPieceNum mocks base method
func (m *MockProgressMgr) GetPeerIDsByPieceNum(ctx context.Context, taskID string, pieceNum int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerIDsByPieceNum", ctx, taskID, pieceNum)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeerIDsByPieceNum indicates an expected call of GetPeerIDsByPieceNum
func (mr *MockProgressMgrMockRecorder) GetPeerIDsByPieceNum(ctx, taskID, pieceNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerIDsByPieceNum", reflect.TypeOf((*MockProgressMgr)(nil).GetPeerIDsByPieceNum), ctx, taskID, pieceNum)
}

// DeletePeerIDByPieceNum mocks base method
func (m *MockProgressMgr) DeletePeerIDByPieceNum(ctx context.Context, taskID string, pieceNum int, peerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePeerIDByPieceNum", ctx, taskID, pieceNum, peerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePeerIDByPieceNum indicates an expected call of DeletePeerIDByPieceNum
func (mr *MockProgressMgrMockRecorder) DeletePeerIDByPieceNum(ctx, taskID, pieceNum, peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePeerIDByPieceNum", reflect.TypeOf((*MockProgressMgr)(nil).DeletePeerIDByPieceNum), ctx, taskID, pieceNum, peerID)
}

// GetPeerStateByPeerID mocks base method
func (m *MockProgressMgr) GetPeerStateByPeerID(ctx context.Context, peerID string) (*mgr.PeerState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerStateByPeerID", ctx, peerID)
	ret0, _ := ret[0].(*mgr.PeerState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeerStateByPeerID indicates an expected call of GetPeerStateByPeerID
func (mr *MockProgressMgrMockRecorder) GetPeerStateByPeerID(ctx, peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerStateByPeerID", reflect.TypeOf((*MockProgressMgr)(nil).GetPeerStateByPeerID), ctx, peerID)
}

// UpdatePeerServiceDown mocks base method
func (m *MockProgressMgr) UpdatePeerServiceDown(ctx context.Context, peerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePeerServiceDown", ctx, peerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePeerServiceDown indicates an expected call of UpdatePeerServiceDown
func (mr *MockProgressMgrMockRecorder) UpdatePeerServiceDown(ctx, peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePeerServiceDown", reflect.TypeOf((*MockProgressMgr)(nil).UpdatePeerServiceDown), ctx, peerID)
}

// GetPeersByTaskID mocks base method
func (m *MockProgressMgr) GetPeersByTaskID(ctx context.Context, taskID string) ([]*types.PeerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeersByTaskID", ctx, taskID)
	ret0, _ := ret[0].([]*types.PeerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeersByTaskID indicates an expected call of GetPeersByTaskID
func (mr *MockProgressMgrMockRecorder) GetPeersByTaskID(ctx, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeersByTaskID", reflect.TypeOf((*MockProgressMgr)(nil).GetPeersByTaskID), ctx, taskID)
}

// GetBlackInfoByPeerID mocks base method
func (m *MockProgressMgr) GetBlackInfoByPeerID(ctx context.Context, peerID string) (*syncmap.SyncMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlackInfoByPeerID", ctx, peerID)
	ret0, _ := ret[0].(*syncmap.SyncMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlackInfoByPeerID indicates an expected call of GetBlackInfoByPeerID
func (mr *MockProgressMgrMockRecorder) GetBlackInfoByPeerID(ctx, peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlackInfoByPeerID", reflect.TypeOf((*MockProgressMgr)(nil).GetBlackInfoByPeerID), ctx, peerID)
}

// UpdateSuperLoad mocks base method
func (m *MockProgressMgr) UpdateSuperLoad(ctx context.Context, taskID string, delta, limit int32) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSuperLoad", ctx, taskID, delta, limit)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSuperLoad indicates an expected call of UpdateSuperLoad
func (mr *MockProgressMgrMockRecorder) UpdateSuperLoad(ctx, taskID, delta, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSuperLoad", reflect.TypeOf((*MockProgressMgr)(nil).UpdateSuperLoad), ctx, taskID, delta, limit)
}

// DeleteTaskID mocks base method
func (m *MockProgressMgr) DeleteTaskID(ctx context.Context, taskID string, pieceTotal int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTaskID", ctx, taskID, pieceTotal)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTaskID indicates an expected call of DeleteTaskID
func (mr *MockProgressMgrMockRecorder) DeleteTaskID(ctx, taskID, pieceTotal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTaskID", reflect.TypeOf((*MockProgressMgr)(nil).DeleteTaskID), ctx, taskID, pieceTotal)
}

// DeleteCID mocks base method
func (m *MockProgressMgr) DeleteCID(ctx context.Context, clientID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCID", ctx, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCID indicates an expected call of DeleteCID
func (mr *MockProgressMgrMockRecorder) DeleteCID(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCID", reflect.TypeOf((*MockProgressMgr)(nil).DeleteCID), ctx, clientID)
}

// DeletePeerID mocks base method
func (m *MockProgressMgr) DeletePeerID(ctx context.Context, peerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePeerID", ctx, peerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePeerID indicates an expected call of DeletePeerID
func (mr *MockProgressMgrMockRecorder) DeletePeerID(ctx, peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePeerID", reflect.TypeOf((*MockProgressMgr)(nil).DeletePeerID), ctx, peerID)
}
