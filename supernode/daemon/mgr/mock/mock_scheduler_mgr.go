// Code generated by MockGen. DO NOT EDIT.
// Source: supernode/daemon/mgr/scheduler_mgr.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	mgr "github.com/Dragonfly/supernode/daemon/mgr"
)

// MockSchedulerMgr is a mock of SchedulerMgr interface
type MockSchedulerMgr struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMgrMockRecorder
}

// MockSchedulerMgrMockRecorder is the mock recorder for MockSchedulerMgr
type MockSchedulerMgrMockRecorder struct {
	mock *MockSchedulerMgr
}

// NewMockSchedulerMgr creates a new mock instance
func NewMockSchedulerMgr(ctrl *gomock.Controller) *MockSchedulerMgr {
	mock := &MockSchedulerMgr{ctrl: ctrl}
	mock.recorder = &MockSchedulerMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSchedulerMgr) EXPECT() *MockSchedulerMgrMockRecorder {
	return m.recorder
}

// Schedule mocks base method
func (m *MockSchedulerMgr) Schedule(ctx context.Context, taskID, clientID, peerID string) ([]*mgr.PieceResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schedule", ctx, taskID, clientID, peerID)
	ret0, _ := ret[0].([]*mgr.PieceResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Schedule indicates an expected call of Schedule
func (mr *MockSchedulerMgrMockRecorder) Schedule(ctx, taskID, clientID, peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockSchedulerMgr)(nil).Schedule), ctx, taskID, clientID, peerID)
}
