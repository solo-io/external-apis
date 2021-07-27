// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/controller"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockAuthorizationPolicyEventHandler is a mock of AuthorizationPolicyEventHandler interface
type MockAuthorizationPolicyEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationPolicyEventHandlerMockRecorder
}

// MockAuthorizationPolicyEventHandlerMockRecorder is the mock recorder for MockAuthorizationPolicyEventHandler
type MockAuthorizationPolicyEventHandlerMockRecorder struct {
	mock *MockAuthorizationPolicyEventHandler
}

// NewMockAuthorizationPolicyEventHandler creates a new mock instance
func NewMockAuthorizationPolicyEventHandler(ctrl *gomock.Controller) *MockAuthorizationPolicyEventHandler {
	mock := &MockAuthorizationPolicyEventHandler{ctrl: ctrl}
	mock.recorder = &MockAuthorizationPolicyEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizationPolicyEventHandler) EXPECT() *MockAuthorizationPolicyEventHandlerMockRecorder {
	return m.recorder
}

// CreateAuthorizationPolicy mocks base method
func (m *MockAuthorizationPolicyEventHandler) CreateAuthorizationPolicy(obj *v1beta1.AuthorizationPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthorizationPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAuthorizationPolicy indicates an expected call of CreateAuthorizationPolicy
func (mr *MockAuthorizationPolicyEventHandlerMockRecorder) CreateAuthorizationPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthorizationPolicy", reflect.TypeOf((*MockAuthorizationPolicyEventHandler)(nil).CreateAuthorizationPolicy), obj)
}

// UpdateAuthorizationPolicy mocks base method
func (m *MockAuthorizationPolicyEventHandler) UpdateAuthorizationPolicy(old, new *v1beta1.AuthorizationPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthorizationPolicy", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthorizationPolicy indicates an expected call of UpdateAuthorizationPolicy
func (mr *MockAuthorizationPolicyEventHandlerMockRecorder) UpdateAuthorizationPolicy(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthorizationPolicy", reflect.TypeOf((*MockAuthorizationPolicyEventHandler)(nil).UpdateAuthorizationPolicy), old, new)
}

// DeleteAuthorizationPolicy mocks base method
func (m *MockAuthorizationPolicyEventHandler) DeleteAuthorizationPolicy(obj *v1beta1.AuthorizationPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthorizationPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthorizationPolicy indicates an expected call of DeleteAuthorizationPolicy
func (mr *MockAuthorizationPolicyEventHandlerMockRecorder) DeleteAuthorizationPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthorizationPolicy", reflect.TypeOf((*MockAuthorizationPolicyEventHandler)(nil).DeleteAuthorizationPolicy), obj)
}

// GenericAuthorizationPolicy mocks base method
func (m *MockAuthorizationPolicyEventHandler) GenericAuthorizationPolicy(obj *v1beta1.AuthorizationPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericAuthorizationPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericAuthorizationPolicy indicates an expected call of GenericAuthorizationPolicy
func (mr *MockAuthorizationPolicyEventHandlerMockRecorder) GenericAuthorizationPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericAuthorizationPolicy", reflect.TypeOf((*MockAuthorizationPolicyEventHandler)(nil).GenericAuthorizationPolicy), obj)
}

// MockAuthorizationPolicyEventWatcher is a mock of AuthorizationPolicyEventWatcher interface
type MockAuthorizationPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationPolicyEventWatcherMockRecorder
}

// MockAuthorizationPolicyEventWatcherMockRecorder is the mock recorder for MockAuthorizationPolicyEventWatcher
type MockAuthorizationPolicyEventWatcherMockRecorder struct {
	mock *MockAuthorizationPolicyEventWatcher
}

// NewMockAuthorizationPolicyEventWatcher creates a new mock instance
func NewMockAuthorizationPolicyEventWatcher(ctrl *gomock.Controller) *MockAuthorizationPolicyEventWatcher {
	mock := &MockAuthorizationPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockAuthorizationPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizationPolicyEventWatcher) EXPECT() *MockAuthorizationPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockAuthorizationPolicyEventWatcher) AddEventHandler(ctx context.Context, h controller.AuthorizationPolicyEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockAuthorizationPolicyEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockAuthorizationPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockPeerAuthenticationEventHandler is a mock of PeerAuthenticationEventHandler interface
type MockPeerAuthenticationEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPeerAuthenticationEventHandlerMockRecorder
}

// MockPeerAuthenticationEventHandlerMockRecorder is the mock recorder for MockPeerAuthenticationEventHandler
type MockPeerAuthenticationEventHandlerMockRecorder struct {
	mock *MockPeerAuthenticationEventHandler
}

// NewMockPeerAuthenticationEventHandler creates a new mock instance
func NewMockPeerAuthenticationEventHandler(ctrl *gomock.Controller) *MockPeerAuthenticationEventHandler {
	mock := &MockPeerAuthenticationEventHandler{ctrl: ctrl}
	mock.recorder = &MockPeerAuthenticationEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeerAuthenticationEventHandler) EXPECT() *MockPeerAuthenticationEventHandlerMockRecorder {
	return m.recorder
}

// CreatePeerAuthentication mocks base method
func (m *MockPeerAuthenticationEventHandler) CreatePeerAuthentication(obj *v1beta1.PeerAuthentication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeerAuthentication", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePeerAuthentication indicates an expected call of CreatePeerAuthentication
func (mr *MockPeerAuthenticationEventHandlerMockRecorder) CreatePeerAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeerAuthentication", reflect.TypeOf((*MockPeerAuthenticationEventHandler)(nil).CreatePeerAuthentication), obj)
}

// UpdatePeerAuthentication mocks base method
func (m *MockPeerAuthenticationEventHandler) UpdatePeerAuthentication(old, new *v1beta1.PeerAuthentication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePeerAuthentication", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePeerAuthentication indicates an expected call of UpdatePeerAuthentication
func (mr *MockPeerAuthenticationEventHandlerMockRecorder) UpdatePeerAuthentication(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePeerAuthentication", reflect.TypeOf((*MockPeerAuthenticationEventHandler)(nil).UpdatePeerAuthentication), old, new)
}

// DeletePeerAuthentication mocks base method
func (m *MockPeerAuthenticationEventHandler) DeletePeerAuthentication(obj *v1beta1.PeerAuthentication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePeerAuthentication", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePeerAuthentication indicates an expected call of DeletePeerAuthentication
func (mr *MockPeerAuthenticationEventHandlerMockRecorder) DeletePeerAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePeerAuthentication", reflect.TypeOf((*MockPeerAuthenticationEventHandler)(nil).DeletePeerAuthentication), obj)
}

// GenericPeerAuthentication mocks base method
func (m *MockPeerAuthenticationEventHandler) GenericPeerAuthentication(obj *v1beta1.PeerAuthentication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericPeerAuthentication", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericPeerAuthentication indicates an expected call of GenericPeerAuthentication
func (mr *MockPeerAuthenticationEventHandlerMockRecorder) GenericPeerAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericPeerAuthentication", reflect.TypeOf((*MockPeerAuthenticationEventHandler)(nil).GenericPeerAuthentication), obj)
}

// MockPeerAuthenticationEventWatcher is a mock of PeerAuthenticationEventWatcher interface
type MockPeerAuthenticationEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockPeerAuthenticationEventWatcherMockRecorder
}

// MockPeerAuthenticationEventWatcherMockRecorder is the mock recorder for MockPeerAuthenticationEventWatcher
type MockPeerAuthenticationEventWatcherMockRecorder struct {
	mock *MockPeerAuthenticationEventWatcher
}

// NewMockPeerAuthenticationEventWatcher creates a new mock instance
func NewMockPeerAuthenticationEventWatcher(ctrl *gomock.Controller) *MockPeerAuthenticationEventWatcher {
	mock := &MockPeerAuthenticationEventWatcher{ctrl: ctrl}
	mock.recorder = &MockPeerAuthenticationEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeerAuthenticationEventWatcher) EXPECT() *MockPeerAuthenticationEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockPeerAuthenticationEventWatcher) AddEventHandler(ctx context.Context, h controller.PeerAuthenticationEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockPeerAuthenticationEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockPeerAuthenticationEventWatcher)(nil).AddEventHandler), varargs...)
}
