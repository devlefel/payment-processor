// Code generated by MockGen. DO NOT EDIT.
// Source: processor/payment/gateway.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	models "processor/payment/models"
	reflect "reflect"
)

// MockGateway is a mock of Gateway interface
type MockGateway struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayMockRecorder
}

// MockGatewayMockRecorder is the mock recorder for MockGateway
type MockGatewayMockRecorder struct {
	mock *MockGateway
}

// NewMockGateway creates a new mock instance
func NewMockGateway(ctrl *gomock.Controller) *MockGateway {
	mock := &MockGateway{ctrl: ctrl}
	mock.recorder = &MockGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGateway) EXPECT() *MockGatewayMockRecorder {
	return m.recorder
}

// ProcessPayment mocks base method
func (m *MockGateway) ProcessPayment(card models.CardData, process models.Process, acquirerID int64, errors *models.Error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPayment", card, process, acquirerID, errors)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ProcessPayment indicates an expected call of ProcessPayment
func (mr *MockGatewayMockRecorder) ProcessPayment(card, process, acquirerID, errors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPayment", reflect.TypeOf((*MockGateway)(nil).ProcessPayment), card, process, acquirerID, errors)
}