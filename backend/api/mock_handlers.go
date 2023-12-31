// Code generated by MockGen. DO NOT EDIT.
// Source: handlers/handlers.go

// Package api is a generated GoMock package.
package api

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGptHandler is a mock of GptHandler interface.
type MockGptHandler struct {
	ctrl     *gomock.Controller
	recorder *MockGptHandlerMockRecorder
}

// MockGptHandlerMockRecorder is the mock recorder for MockGptHandler.
type MockGptHandlerMockRecorder struct {
	mock *MockGptHandler
}

// NewMockGptHandler creates a new mock instance.
func NewMockGptHandler(ctrl *gomock.Controller) *MockGptHandler {
	mock := &MockGptHandler{ctrl: ctrl}
	mock.recorder = &MockGptHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGptHandler) EXPECT() *MockGptHandlerMockRecorder {
	return m.recorder
}

// FindSectors mocks base method.
func (m *MockGptHandler) FindSectors(tickers []string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSectors", tickers)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSectors indicates an expected call of FindSectors.
func (mr *MockGptHandlerMockRecorder) FindSectors(tickers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSectors", reflect.TypeOf((*MockGptHandler)(nil).FindSectors), tickers)
}

// MockRedisHandler is a mock of RedisHandler interface.
type MockRedisHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRedisHandlerMockRecorder
}

// MockRedisHandlerMockRecorder is the mock recorder for MockRedisHandler.
type MockRedisHandlerMockRecorder struct {
	mock *MockRedisHandler
}

// NewMockRedisHandler creates a new mock instance.
func NewMockRedisHandler(ctrl *gomock.Controller) *MockRedisHandler {
	mock := &MockRedisHandler{ctrl: ctrl}
	mock.recorder = &MockRedisHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisHandler) EXPECT() *MockRedisHandlerMockRecorder {
	return m.recorder
}

// FindSectors mocks base method.
func (m *MockRedisHandler) FindSectors(tickers map[string]bool) (map[string]string, map[string]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSectors", tickers)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(map[string]bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindSectors indicates an expected call of FindSectors.
func (mr *MockRedisHandlerMockRecorder) FindSectors(tickers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSectors", reflect.TypeOf((*MockRedisHandler)(nil).FindSectors), tickers)
}

// SetTicker mocks base method.
func (m *MockRedisHandler) SetTicker(tickers map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTicker", tickers)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTicker indicates an expected call of SetTicker.
func (mr *MockRedisHandlerMockRecorder) SetTicker(tickers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTicker", reflect.TypeOf((*MockRedisHandler)(nil).SetTicker), tickers)
}
