// Code generated by MockGen. DO NOT EDIT.
// Source: iwf/communication.go

// Package iwftest is a generated GoMock package.
package iwftest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	iwfidl "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

// MockCommunication is a mock of Communication interface.
type MockCommunication struct {
	ctrl     *gomock.Controller
	recorder *MockCommunicationMockRecorder
}

// MockCommunicationMockRecorder is the mock recorder for MockCommunication.
type MockCommunicationMockRecorder struct {
	mock *MockCommunication
}

// NewMockCommunication creates a new mock instance.
func NewMockCommunication(ctrl *gomock.Controller) *MockCommunication {
	mock := &MockCommunication{ctrl: ctrl}
	mock.recorder = &MockCommunicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommunication) EXPECT() *MockCommunicationMockRecorder {
	return m.recorder
}

// PublishInterstateChannel mocks base method.
func (m *MockCommunication) PublishInterstateChannel(channelName string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishInterstateChannel", channelName, value)
}

// PublishInterstateChannel indicates an expected call of PublishInterstateChannel.
func (mr *MockCommunicationMockRecorder) PublishInterstateChannel(channelName, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishInterstateChannel", reflect.TypeOf((*MockCommunication)(nil).PublishInterstateChannel), channelName, value)
}

// getToPublishInterStateChannel mocks base method.
func (m *MockCommunication) getToPublishInterStateChannel() map[string][]iwfidl.EncodedObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getToPublishInterStateChannel")
	ret0, _ := ret[0].(map[string][]iwfidl.EncodedObject)
	return ret0
}

// getToPublishInterStateChannel indicates an expected call of getToPublishInterStateChannel.
func (mr *MockCommunicationMockRecorder) getToPublishInterStateChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getToPublishInterStateChannel", reflect.TypeOf((*MockCommunication)(nil).getToPublishInterStateChannel))
}
