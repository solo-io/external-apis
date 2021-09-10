// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"
	controller "github.com/solo-io/external-apis/pkg/api/smi/split.smi-spec.io/v1alpha2/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockTrafficSplitEventHandler is a mock of TrafficSplitEventHandler interface.
type MockTrafficSplitEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitEventHandlerMockRecorder
}

// MockTrafficSplitEventHandlerMockRecorder is the mock recorder for MockTrafficSplitEventHandler.
type MockTrafficSplitEventHandlerMockRecorder struct {
	mock *MockTrafficSplitEventHandler
}

// NewMockTrafficSplitEventHandler creates a new mock instance.
func NewMockTrafficSplitEventHandler(ctrl *gomock.Controller) *MockTrafficSplitEventHandler {
	mock := &MockTrafficSplitEventHandler{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitEventHandler) EXPECT() *MockTrafficSplitEventHandlerMockRecorder {
	return m.recorder
}

// CreateTrafficSplit mocks base method.
func (m *MockTrafficSplitEventHandler) CreateTrafficSplit(obj *v1alpha2.TrafficSplit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrafficSplit", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrafficSplit indicates an expected call of CreateTrafficSplit.
func (mr *MockTrafficSplitEventHandlerMockRecorder) CreateTrafficSplit(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrafficSplit", reflect.TypeOf((*MockTrafficSplitEventHandler)(nil).CreateTrafficSplit), obj)
}

// DeleteTrafficSplit mocks base method.
func (m *MockTrafficSplitEventHandler) DeleteTrafficSplit(obj *v1alpha2.TrafficSplit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrafficSplit", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrafficSplit indicates an expected call of DeleteTrafficSplit.
func (mr *MockTrafficSplitEventHandlerMockRecorder) DeleteTrafficSplit(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrafficSplit", reflect.TypeOf((*MockTrafficSplitEventHandler)(nil).DeleteTrafficSplit), obj)
}

// GenericTrafficSplit mocks base method.
func (m *MockTrafficSplitEventHandler) GenericTrafficSplit(obj *v1alpha2.TrafficSplit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericTrafficSplit", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericTrafficSplit indicates an expected call of GenericTrafficSplit.
func (mr *MockTrafficSplitEventHandlerMockRecorder) GenericTrafficSplit(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericTrafficSplit", reflect.TypeOf((*MockTrafficSplitEventHandler)(nil).GenericTrafficSplit), obj)
}

// UpdateTrafficSplit mocks base method.
func (m *MockTrafficSplitEventHandler) UpdateTrafficSplit(old, new *v1alpha2.TrafficSplit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrafficSplit", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrafficSplit indicates an expected call of UpdateTrafficSplit.
func (mr *MockTrafficSplitEventHandlerMockRecorder) UpdateTrafficSplit(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrafficSplit", reflect.TypeOf((*MockTrafficSplitEventHandler)(nil).UpdateTrafficSplit), old, new)
}

// MockTrafficSplitEventWatcher is a mock of TrafficSplitEventWatcher interface.
type MockTrafficSplitEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitEventWatcherMockRecorder
}

// MockTrafficSplitEventWatcherMockRecorder is the mock recorder for MockTrafficSplitEventWatcher.
type MockTrafficSplitEventWatcherMockRecorder struct {
	mock *MockTrafficSplitEventWatcher
}

// NewMockTrafficSplitEventWatcher creates a new mock instance.
func NewMockTrafficSplitEventWatcher(ctrl *gomock.Controller) *MockTrafficSplitEventWatcher {
	mock := &MockTrafficSplitEventWatcher{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitEventWatcher) EXPECT() *MockTrafficSplitEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockTrafficSplitEventWatcher) AddEventHandler(ctx context.Context, h controller.TrafficSplitEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockTrafficSplitEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockTrafficSplitEventWatcher)(nil).AddEventHandler), varargs...)
}
