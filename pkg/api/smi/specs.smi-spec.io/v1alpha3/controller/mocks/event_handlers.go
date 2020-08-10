// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	controller "github.com/solo-io/external-apis/pkg/api/smi/specs.smi-spec.io/v1alpha3/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockHTTPRouteGroupEventHandler is a mock of HTTPRouteGroupEventHandler interface
type MockHTTPRouteGroupEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupEventHandlerMockRecorder
}

// MockHTTPRouteGroupEventHandlerMockRecorder is the mock recorder for MockHTTPRouteGroupEventHandler
type MockHTTPRouteGroupEventHandlerMockRecorder struct {
	mock *MockHTTPRouteGroupEventHandler
}

// NewMockHTTPRouteGroupEventHandler creates a new mock instance
func NewMockHTTPRouteGroupEventHandler(ctrl *gomock.Controller) *MockHTTPRouteGroupEventHandler {
	mock := &MockHTTPRouteGroupEventHandler{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupEventHandler) EXPECT() *MockHTTPRouteGroupEventHandlerMockRecorder {
	return m.recorder
}

// CreateHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupEventHandler) CreateHTTPRouteGroup(obj *v1alpha3.HTTPRouteGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHTTPRouteGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHTTPRouteGroup indicates an expected call of CreateHTTPRouteGroup
func (mr *MockHTTPRouteGroupEventHandlerMockRecorder) CreateHTTPRouteGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupEventHandler)(nil).CreateHTTPRouteGroup), obj)
}

// UpdateHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupEventHandler) UpdateHTTPRouteGroup(old, new *v1alpha3.HTTPRouteGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHTTPRouteGroup", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHTTPRouteGroup indicates an expected call of UpdateHTTPRouteGroup
func (mr *MockHTTPRouteGroupEventHandlerMockRecorder) UpdateHTTPRouteGroup(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupEventHandler)(nil).UpdateHTTPRouteGroup), old, new)
}

// DeleteHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupEventHandler) DeleteHTTPRouteGroup(obj *v1alpha3.HTTPRouteGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHTTPRouteGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHTTPRouteGroup indicates an expected call of DeleteHTTPRouteGroup
func (mr *MockHTTPRouteGroupEventHandlerMockRecorder) DeleteHTTPRouteGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupEventHandler)(nil).DeleteHTTPRouteGroup), obj)
}

// GenericHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupEventHandler) GenericHTTPRouteGroup(obj *v1alpha3.HTTPRouteGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericHTTPRouteGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericHTTPRouteGroup indicates an expected call of GenericHTTPRouteGroup
func (mr *MockHTTPRouteGroupEventHandlerMockRecorder) GenericHTTPRouteGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupEventHandler)(nil).GenericHTTPRouteGroup), obj)
}

// MockHTTPRouteGroupEventWatcher is a mock of HTTPRouteGroupEventWatcher interface
type MockHTTPRouteGroupEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupEventWatcherMockRecorder
}

// MockHTTPRouteGroupEventWatcherMockRecorder is the mock recorder for MockHTTPRouteGroupEventWatcher
type MockHTTPRouteGroupEventWatcherMockRecorder struct {
	mock *MockHTTPRouteGroupEventWatcher
}

// NewMockHTTPRouteGroupEventWatcher creates a new mock instance
func NewMockHTTPRouteGroupEventWatcher(ctrl *gomock.Controller) *MockHTTPRouteGroupEventWatcher {
	mock := &MockHTTPRouteGroupEventWatcher{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupEventWatcher) EXPECT() *MockHTTPRouteGroupEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockHTTPRouteGroupEventWatcher) AddEventHandler(ctx context.Context, h controller.HTTPRouteGroupEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockHTTPRouteGroupEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockHTTPRouteGroupEventWatcher)(nil).AddEventHandler), varargs...)
}
