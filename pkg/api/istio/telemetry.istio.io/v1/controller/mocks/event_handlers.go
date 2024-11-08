// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/telemetry.istio.io/v1/controller"
	v1 "istio.io/client-go/pkg/apis/telemetry/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockTelemetryEventHandler is a mock of TelemetryEventHandler interface.
type MockTelemetryEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTelemetryEventHandlerMockRecorder
}

// MockTelemetryEventHandlerMockRecorder is the mock recorder for MockTelemetryEventHandler.
type MockTelemetryEventHandlerMockRecorder struct {
	mock *MockTelemetryEventHandler
}

// NewMockTelemetryEventHandler creates a new mock instance.
func NewMockTelemetryEventHandler(ctrl *gomock.Controller) *MockTelemetryEventHandler {
	mock := &MockTelemetryEventHandler{ctrl: ctrl}
	mock.recorder = &MockTelemetryEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelemetryEventHandler) EXPECT() *MockTelemetryEventHandlerMockRecorder {
	return m.recorder
}

// CreateTelemetry mocks base method.
func (m *MockTelemetryEventHandler) CreateTelemetry(obj *v1.Telemetry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTelemetry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTelemetry indicates an expected call of CreateTelemetry.
func (mr *MockTelemetryEventHandlerMockRecorder) CreateTelemetry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTelemetry", reflect.TypeOf((*MockTelemetryEventHandler)(nil).CreateTelemetry), obj)
}

// DeleteTelemetry mocks base method.
func (m *MockTelemetryEventHandler) DeleteTelemetry(obj *v1.Telemetry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTelemetry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTelemetry indicates an expected call of DeleteTelemetry.
func (mr *MockTelemetryEventHandlerMockRecorder) DeleteTelemetry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTelemetry", reflect.TypeOf((*MockTelemetryEventHandler)(nil).DeleteTelemetry), obj)
}

// GenericTelemetry mocks base method.
func (m *MockTelemetryEventHandler) GenericTelemetry(obj *v1.Telemetry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericTelemetry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericTelemetry indicates an expected call of GenericTelemetry.
func (mr *MockTelemetryEventHandlerMockRecorder) GenericTelemetry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericTelemetry", reflect.TypeOf((*MockTelemetryEventHandler)(nil).GenericTelemetry), obj)
}

// UpdateTelemetry mocks base method.
func (m *MockTelemetryEventHandler) UpdateTelemetry(old, new *v1.Telemetry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTelemetry", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTelemetry indicates an expected call of UpdateTelemetry.
func (mr *MockTelemetryEventHandlerMockRecorder) UpdateTelemetry(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTelemetry", reflect.TypeOf((*MockTelemetryEventHandler)(nil).UpdateTelemetry), old, new)
}

// MockTelemetryEventWatcher is a mock of TelemetryEventWatcher interface.
type MockTelemetryEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockTelemetryEventWatcherMockRecorder
}

// MockTelemetryEventWatcherMockRecorder is the mock recorder for MockTelemetryEventWatcher.
type MockTelemetryEventWatcherMockRecorder struct {
	mock *MockTelemetryEventWatcher
}

// NewMockTelemetryEventWatcher creates a new mock instance.
func NewMockTelemetryEventWatcher(ctrl *gomock.Controller) *MockTelemetryEventWatcher {
	mock := &MockTelemetryEventWatcher{ctrl: ctrl}
	mock.recorder = &MockTelemetryEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelemetryEventWatcher) EXPECT() *MockTelemetryEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockTelemetryEventWatcher) AddEventHandler(ctx context.Context, h controller.TelemetryEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockTelemetryEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockTelemetryEventWatcher)(nil).AddEventHandler), varargs...)
}