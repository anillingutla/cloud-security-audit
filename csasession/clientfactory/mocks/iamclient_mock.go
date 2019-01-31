// Code generated by MockGen. DO NOT EDIT.
// Source: ./csasession/clientfactory/iamclient.go

// Package mocks is a generated GoMock package.
package mocks

import (
	iam "github.com/aws/aws-sdk-go/service/iam"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAMClient is a mock of IAMClient interface
type MockIAMClient struct {
	ctrl     *gomock.Controller
	recorder *MockIAMClientMockRecorder
}

// MockIAMClientMockRecorder is the mock recorder for MockIAMClient
type MockIAMClientMockRecorder struct {
	mock *MockIAMClient
}

// NewMockIAMClient creates a new mock instance
func NewMockIAMClient(ctrl *gomock.Controller) *MockIAMClient {
	mock := &MockIAMClient{ctrl: ctrl}
	mock.recorder = &MockIAMClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAMClient) EXPECT() *MockIAMClientMockRecorder {
	return m.recorder
}

// ListUsers mocks base method
func (m *MockIAMClient) ListUsers(input *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", input)
	ret0, _ := ret[0].(*iam.GetAccountAuthorizationDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockIAMClientMockRecorder) ListUsers(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockIAMClient)(nil).ListUsers), input)
}

// ListAccessKeys mocks base method
func (m *MockIAMClient) ListAccessKeys(input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessKeys", input)
	ret0, _ := ret[0].(*iam.ListAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessKeys indicates an expected call of ListAccessKeys
func (mr *MockIAMClientMockRecorder) ListAccessKeys(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessKeys", reflect.TypeOf((*MockIAMClient)(nil).ListAccessKeys), input)
}
