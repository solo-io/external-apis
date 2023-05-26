// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/gateway.networking.k8s.io/v1beta1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// MockGatewayEventHandler is a mock of GatewayEventHandler interface.
type MockGatewayEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayEventHandlerMockRecorder
}

// MockGatewayEventHandlerMockRecorder is the mock recorder for MockGatewayEventHandler.
type MockGatewayEventHandlerMockRecorder struct {
	mock *MockGatewayEventHandler
}

// NewMockGatewayEventHandler creates a new mock instance.
func NewMockGatewayEventHandler(ctrl *gomock.Controller) *MockGatewayEventHandler {
	mock := &MockGatewayEventHandler{ctrl: ctrl}
	mock.recorder = &MockGatewayEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayEventHandler) EXPECT() *MockGatewayEventHandlerMockRecorder {
	return m.recorder
}

// CreateGateway mocks base method.
func (m *MockGatewayEventHandler) CreateGateway(obj *v1beta1.Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGateway indicates an expected call of CreateGateway.
func (mr *MockGatewayEventHandlerMockRecorder) CreateGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGateway", reflect.TypeOf((*MockGatewayEventHandler)(nil).CreateGateway), obj)
}

// DeleteGateway mocks base method.
func (m *MockGatewayEventHandler) DeleteGateway(obj *v1beta1.Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGateway indicates an expected call of DeleteGateway.
func (mr *MockGatewayEventHandlerMockRecorder) DeleteGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGateway", reflect.TypeOf((*MockGatewayEventHandler)(nil).DeleteGateway), obj)
}

// GenericGateway mocks base method.
func (m *MockGatewayEventHandler) GenericGateway(obj *v1beta1.Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericGateway indicates an expected call of GenericGateway.
func (mr *MockGatewayEventHandlerMockRecorder) GenericGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericGateway", reflect.TypeOf((*MockGatewayEventHandler)(nil).GenericGateway), obj)
}

// UpdateGateway mocks base method.
func (m *MockGatewayEventHandler) UpdateGateway(old, new *v1beta1.Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGateway", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGateway indicates an expected call of UpdateGateway.
func (mr *MockGatewayEventHandlerMockRecorder) UpdateGateway(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGateway", reflect.TypeOf((*MockGatewayEventHandler)(nil).UpdateGateway), old, new)
}

// MockGatewayEventWatcher is a mock of GatewayEventWatcher interface.
type MockGatewayEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayEventWatcherMockRecorder
}

// MockGatewayEventWatcherMockRecorder is the mock recorder for MockGatewayEventWatcher.
type MockGatewayEventWatcherMockRecorder struct {
	mock *MockGatewayEventWatcher
}

// NewMockGatewayEventWatcher creates a new mock instance.
func NewMockGatewayEventWatcher(ctrl *gomock.Controller) *MockGatewayEventWatcher {
	mock := &MockGatewayEventWatcher{ctrl: ctrl}
	mock.recorder = &MockGatewayEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayEventWatcher) EXPECT() *MockGatewayEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockGatewayEventWatcher) AddEventHandler(ctx context.Context, h controller.GatewayEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockGatewayEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockGatewayEventWatcher)(nil).AddEventHandler), varargs...)
}
