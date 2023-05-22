// Code generated by MockGen. DO NOT EDIT.
// Source: clients/client.go

// Package mock_clients is a generated GoMock package.
package mock_clients

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockClient) GetData() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetData indicates an expected call of GetData.
func (mr *MockClientMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockClient)(nil).GetData))
}
