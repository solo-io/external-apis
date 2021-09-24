// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/apiextensions.k8s.io/v1/controller"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockCustomResourceDefinitionEventHandler is a mock of CustomResourceDefinitionEventHandler interface.
type MockCustomResourceDefinitionEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionEventHandlerMockRecorder
}

// MockCustomResourceDefinitionEventHandlerMockRecorder is the mock recorder for MockCustomResourceDefinitionEventHandler.
type MockCustomResourceDefinitionEventHandlerMockRecorder struct {
	mock *MockCustomResourceDefinitionEventHandler
}

// NewMockCustomResourceDefinitionEventHandler creates a new mock instance.
func NewMockCustomResourceDefinitionEventHandler(ctrl *gomock.Controller) *MockCustomResourceDefinitionEventHandler {
	mock := &MockCustomResourceDefinitionEventHandler{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResourceDefinitionEventHandler) EXPECT() *MockCustomResourceDefinitionEventHandlerMockRecorder {
	return m.recorder
}

// CreateCustomResourceDefinition mocks base method.
func (m *MockCustomResourceDefinitionEventHandler) CreateCustomResourceDefinition(obj *v1.CustomResourceDefinition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomResourceDefinition", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomResourceDefinition indicates an expected call of CreateCustomResourceDefinition.
func (mr *MockCustomResourceDefinitionEventHandlerMockRecorder) CreateCustomResourceDefinition(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionEventHandler)(nil).CreateCustomResourceDefinition), obj)
}

// DeleteCustomResourceDefinition mocks base method.
func (m *MockCustomResourceDefinitionEventHandler) DeleteCustomResourceDefinition(obj *v1.CustomResourceDefinition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCustomResourceDefinition", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCustomResourceDefinition indicates an expected call of DeleteCustomResourceDefinition.
func (mr *MockCustomResourceDefinitionEventHandlerMockRecorder) DeleteCustomResourceDefinition(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionEventHandler)(nil).DeleteCustomResourceDefinition), obj)
}

// GenericCustomResourceDefinition mocks base method.
func (m *MockCustomResourceDefinitionEventHandler) GenericCustomResourceDefinition(obj *v1.CustomResourceDefinition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericCustomResourceDefinition", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericCustomResourceDefinition indicates an expected call of GenericCustomResourceDefinition.
func (mr *MockCustomResourceDefinitionEventHandlerMockRecorder) GenericCustomResourceDefinition(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionEventHandler)(nil).GenericCustomResourceDefinition), obj)
}

// UpdateCustomResourceDefinition mocks base method.
func (m *MockCustomResourceDefinitionEventHandler) UpdateCustomResourceDefinition(old, new *v1.CustomResourceDefinition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomResourceDefinition", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomResourceDefinition indicates an expected call of UpdateCustomResourceDefinition.
func (mr *MockCustomResourceDefinitionEventHandlerMockRecorder) UpdateCustomResourceDefinition(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionEventHandler)(nil).UpdateCustomResourceDefinition), old, new)
}

// MockCustomResourceDefinitionEventWatcher is a mock of CustomResourceDefinitionEventWatcher interface.
type MockCustomResourceDefinitionEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionEventWatcherMockRecorder
}

// MockCustomResourceDefinitionEventWatcherMockRecorder is the mock recorder for MockCustomResourceDefinitionEventWatcher.
type MockCustomResourceDefinitionEventWatcherMockRecorder struct {
	mock *MockCustomResourceDefinitionEventWatcher
}

// NewMockCustomResourceDefinitionEventWatcher creates a new mock instance.
func NewMockCustomResourceDefinitionEventWatcher(ctrl *gomock.Controller) *MockCustomResourceDefinitionEventWatcher {
	mock := &MockCustomResourceDefinitionEventWatcher{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResourceDefinitionEventWatcher) EXPECT() *MockCustomResourceDefinitionEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockCustomResourceDefinitionEventWatcher) AddEventHandler(ctx context.Context, h controller.CustomResourceDefinitionEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockCustomResourceDefinitionEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockCustomResourceDefinitionEventWatcher)(nil).AddEventHandler), varargs...)
}
