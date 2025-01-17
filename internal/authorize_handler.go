// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: AuthorizeEndpointHandler)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockAuthorizeEndpointHandler is a mock of AuthorizeEndpointHandler interface.
type MockAuthorizeEndpointHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizeEndpointHandlerMockRecorder
}

// MockAuthorizeEndpointHandlerMockRecorder is the mock recorder for MockAuthorizeEndpointHandler.
type MockAuthorizeEndpointHandlerMockRecorder struct {
	mock *MockAuthorizeEndpointHandler
}

// NewMockAuthorizeEndpointHandler creates a new mock instance.
func NewMockAuthorizeEndpointHandler(ctrl *gomock.Controller) *MockAuthorizeEndpointHandler {
	mock := &MockAuthorizeEndpointHandler{ctrl: ctrl}
	mock.recorder = &MockAuthorizeEndpointHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizeEndpointHandler) EXPECT() *MockAuthorizeEndpointHandlerMockRecorder {
	return m.recorder
}

// HandleAuthorizeEndpointRequest mocks base method.
func (m *MockAuthorizeEndpointHandler) HandleAuthorizeEndpointRequest(arg0 context.Context, arg1 fosite.AuthorizeRequester, arg2 fosite.AuthorizeResponder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleAuthorizeEndpointRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleAuthorizeEndpointRequest indicates an expected call of HandleAuthorizeEndpointRequest.
func (mr *MockAuthorizeEndpointHandlerMockRecorder) HandleAuthorizeEndpointRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAuthorizeEndpointRequest", reflect.TypeOf((*MockAuthorizeEndpointHandler)(nil).HandleAuthorizeEndpointRequest), arg0, arg1, arg2)
}
