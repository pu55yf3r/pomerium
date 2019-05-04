// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pomerium/pomerium/proto/authorize (interfaces: AuthorizerClient)

// Package mock_authorize is a generated GoMock package.
package mock_authorize

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	authorize "github.com/pomerium/pomerium/proto/authorize"
	grpc "google.golang.org/grpc"
)

// MockAuthorizerClient is a mock of AuthorizerClient interface
type MockAuthorizerClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerClientMockRecorder
}

// MockAuthorizerClientMockRecorder is the mock recorder for MockAuthorizerClient
type MockAuthorizerClientMockRecorder struct {
	mock *MockAuthorizerClient
}

// NewMockAuthorizerClient creates a new mock instance
func NewMockAuthorizerClient(ctrl *gomock.Controller) *MockAuthorizerClient {
	mock := &MockAuthorizerClient{ctrl: ctrl}
	mock.recorder = &MockAuthorizerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizerClient) EXPECT() *MockAuthorizerClientMockRecorder {
	return m.recorder
}

// Authorize mocks base method
func (m *MockAuthorizerClient) Authorize(arg0 context.Context, arg1 *authorize.AuthorizeRequest, arg2 ...grpc.CallOption) (*authorize.AuthorizeReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Authorize", varargs...)
	ret0, _ := ret[0].(*authorize.AuthorizeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize
func (mr *MockAuthorizerClientMockRecorder) Authorize(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizerClient)(nil).Authorize), varargs...)
}