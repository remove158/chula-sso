// Code generated by MockGen. DO NOT EDIT.
// Source: ./auth.go

// Package handlers is a generated GoMock package.
package handlers

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIAuthHandler is a mock of IAuthHandler interface.
type MockIAuthHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthHandlerMockRecorder
}

// MockIAuthHandlerMockRecorder is the mock recorder for MockIAuthHandler.
type MockIAuthHandlerMockRecorder struct {
	mock *MockIAuthHandler
}

// NewMockIAuthHandler creates a new mock instance.
func NewMockIAuthHandler(ctrl *gomock.Controller) *MockIAuthHandler {
	mock := &MockIAuthHandler{ctrl: ctrl}
	mock.recorder = &MockIAuthHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthHandler) EXPECT() *MockIAuthHandlerMockRecorder {
	return m.recorder
}

// GetLogin mocks base method.
func (m *MockIAuthHandler) GetLogin(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetLogin", ctx)
}

// GetLogin indicates an expected call of GetLogin.
func (mr *MockIAuthHandlerMockRecorder) GetLogin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogin", reflect.TypeOf((*MockIAuthHandler)(nil).GetLogin), ctx)
}

// PostLogin mocks base method.
func (m *MockIAuthHandler) PostLogin(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostLogin", ctx)
}

// PostLogin indicates an expected call of PostLogin.
func (mr *MockIAuthHandlerMockRecorder) PostLogin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostLogin", reflect.TypeOf((*MockIAuthHandler)(nil).PostLogin), ctx)
}

// ServiceValidation mocks base method.
func (m *MockIAuthHandler) ServiceValidation(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServiceValidation", ctx)
}

// ServiceValidation indicates an expected call of ServiceValidation.
func (mr *MockIAuthHandlerMockRecorder) ServiceValidation(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceValidation", reflect.TypeOf((*MockIAuthHandler)(nil).ServiceValidation), ctx)
}
