// Code generated by MockGen. DO NOT EDIT.
// Source: sharing/rabbit/connection (interfaces: ConnectionInterface,RConnectionInterface)

// Package mock_connection is a generated GoMock package.
package mock_connection

import (
	gomock "github.com/golang/mock/gomock"
	amqp "github.com/streadway/amqp"
	reflect "reflect"
	config "rabbit-mq-go/config"
)

// MockConnectionInterface is a mock of ConnectionInterface interface
type MockConnectionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionInterfaceMockRecorder
}

// MockConnectionInterfaceMockRecorder is the mock recorder for MockConnectionInterface
type MockConnectionInterfaceMockRecorder struct {
	mock *MockConnectionInterface
}

// NewMockConnectionInterface creates a new mock instance
func NewMockConnectionInterface(ctrl *gomock.Controller) *MockConnectionInterface {
	mock := &MockConnectionInterface{ctrl: ctrl}
	mock.recorder = &MockConnectionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnectionInterface) EXPECT() *MockConnectionInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockConnectionInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConnectionInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnectionInterface)(nil).Close))
}

// Connect mocks base method
func (m *MockConnectionInterface) Connect(arg0 config.Credentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockConnectionInterfaceMockRecorder) Connect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockConnectionInterface)(nil).Connect), arg0)
}

// MockRConnectionInterface is a mock of RConnectionInterface interface
type MockRConnectionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRConnectionInterfaceMockRecorder
}

// MockRConnectionInterfaceMockRecorder is the mock recorder for MockRConnectionInterface
type MockRConnectionInterfaceMockRecorder struct {
	mock *MockRConnectionInterface
}

// NewMockRConnectionInterface creates a new mock instance
func NewMockRConnectionInterface(ctrl *gomock.Controller) *MockRConnectionInterface {
	mock := &MockRConnectionInterface{ctrl: ctrl}
	mock.recorder = &MockRConnectionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRConnectionInterface) EXPECT() *MockRConnectionInterfaceMockRecorder {
	return m.recorder
}

// Channel mocks base method
func (m *MockRConnectionInterface) Channel() (*amqp.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Channel")
	ret0, _ := ret[0].(*amqp.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Channel indicates an expected call of Channel
func (mr *MockRConnectionInterfaceMockRecorder) Channel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Channel", reflect.TypeOf((*MockRConnectionInterface)(nil).Channel))
}

// Close mocks base method
func (m *MockRConnectionInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRConnectionInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRConnectionInterface)(nil).Close))
}
