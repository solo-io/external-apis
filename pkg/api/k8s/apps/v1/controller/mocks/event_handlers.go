// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/controller"
	v1 "k8s.io/api/apps/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockDeploymentEventHandler is a mock of DeploymentEventHandler interface.
type MockDeploymentEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentEventHandlerMockRecorder
}

// MockDeploymentEventHandlerMockRecorder is the mock recorder for MockDeploymentEventHandler.
type MockDeploymentEventHandlerMockRecorder struct {
	mock *MockDeploymentEventHandler
}

// NewMockDeploymentEventHandler creates a new mock instance.
func NewMockDeploymentEventHandler(ctrl *gomock.Controller) *MockDeploymentEventHandler {
	mock := &MockDeploymentEventHandler{ctrl: ctrl}
	mock.recorder = &MockDeploymentEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentEventHandler) EXPECT() *MockDeploymentEventHandlerMockRecorder {
	return m.recorder
}

// CreateDeployment mocks base method.
func (m *MockDeploymentEventHandler) CreateDeployment(obj *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeployment", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDeployment indicates an expected call of CreateDeployment.
func (mr *MockDeploymentEventHandlerMockRecorder) CreateDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployment", reflect.TypeOf((*MockDeploymentEventHandler)(nil).CreateDeployment), obj)
}

// DeleteDeployment mocks base method.
func (m *MockDeploymentEventHandler) DeleteDeployment(obj *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeployment", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeployment indicates an expected call of DeleteDeployment.
func (mr *MockDeploymentEventHandlerMockRecorder) DeleteDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeployment", reflect.TypeOf((*MockDeploymentEventHandler)(nil).DeleteDeployment), obj)
}

// GenericDeployment mocks base method.
func (m *MockDeploymentEventHandler) GenericDeployment(obj *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericDeployment", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericDeployment indicates an expected call of GenericDeployment.
func (mr *MockDeploymentEventHandlerMockRecorder) GenericDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericDeployment", reflect.TypeOf((*MockDeploymentEventHandler)(nil).GenericDeployment), obj)
}

// UpdateDeployment mocks base method.
func (m *MockDeploymentEventHandler) UpdateDeployment(old, new *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeployment", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeployment indicates an expected call of UpdateDeployment.
func (mr *MockDeploymentEventHandlerMockRecorder) UpdateDeployment(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployment", reflect.TypeOf((*MockDeploymentEventHandler)(nil).UpdateDeployment), old, new)
}

// MockDeploymentEventWatcher is a mock of DeploymentEventWatcher interface.
type MockDeploymentEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentEventWatcherMockRecorder
}

// MockDeploymentEventWatcherMockRecorder is the mock recorder for MockDeploymentEventWatcher.
type MockDeploymentEventWatcherMockRecorder struct {
	mock *MockDeploymentEventWatcher
}

// NewMockDeploymentEventWatcher creates a new mock instance.
func NewMockDeploymentEventWatcher(ctrl *gomock.Controller) *MockDeploymentEventWatcher {
	mock := &MockDeploymentEventWatcher{ctrl: ctrl}
	mock.recorder = &MockDeploymentEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentEventWatcher) EXPECT() *MockDeploymentEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockDeploymentEventWatcher) AddEventHandler(ctx context.Context, h controller.DeploymentEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockDeploymentEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockDeploymentEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockReplicaSetEventHandler is a mock of ReplicaSetEventHandler interface.
type MockReplicaSetEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetEventHandlerMockRecorder
}

// MockReplicaSetEventHandlerMockRecorder is the mock recorder for MockReplicaSetEventHandler.
type MockReplicaSetEventHandlerMockRecorder struct {
	mock *MockReplicaSetEventHandler
}

// NewMockReplicaSetEventHandler creates a new mock instance.
func NewMockReplicaSetEventHandler(ctrl *gomock.Controller) *MockReplicaSetEventHandler {
	mock := &MockReplicaSetEventHandler{ctrl: ctrl}
	mock.recorder = &MockReplicaSetEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetEventHandler) EXPECT() *MockReplicaSetEventHandlerMockRecorder {
	return m.recorder
}

// CreateReplicaSet mocks base method.
func (m *MockReplicaSetEventHandler) CreateReplicaSet(obj *v1.ReplicaSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReplicaSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReplicaSet indicates an expected call of CreateReplicaSet.
func (mr *MockReplicaSetEventHandlerMockRecorder) CreateReplicaSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReplicaSet", reflect.TypeOf((*MockReplicaSetEventHandler)(nil).CreateReplicaSet), obj)
}

// DeleteReplicaSet mocks base method.
func (m *MockReplicaSetEventHandler) DeleteReplicaSet(obj *v1.ReplicaSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReplicaSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReplicaSet indicates an expected call of DeleteReplicaSet.
func (mr *MockReplicaSetEventHandlerMockRecorder) DeleteReplicaSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReplicaSet", reflect.TypeOf((*MockReplicaSetEventHandler)(nil).DeleteReplicaSet), obj)
}

// GenericReplicaSet mocks base method.
func (m *MockReplicaSetEventHandler) GenericReplicaSet(obj *v1.ReplicaSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericReplicaSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericReplicaSet indicates an expected call of GenericReplicaSet.
func (mr *MockReplicaSetEventHandlerMockRecorder) GenericReplicaSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericReplicaSet", reflect.TypeOf((*MockReplicaSetEventHandler)(nil).GenericReplicaSet), obj)
}

// UpdateReplicaSet mocks base method.
func (m *MockReplicaSetEventHandler) UpdateReplicaSet(old, new *v1.ReplicaSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReplicaSet", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReplicaSet indicates an expected call of UpdateReplicaSet.
func (mr *MockReplicaSetEventHandlerMockRecorder) UpdateReplicaSet(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReplicaSet", reflect.TypeOf((*MockReplicaSetEventHandler)(nil).UpdateReplicaSet), old, new)
}

// MockReplicaSetEventWatcher is a mock of ReplicaSetEventWatcher interface.
type MockReplicaSetEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetEventWatcherMockRecorder
}

// MockReplicaSetEventWatcherMockRecorder is the mock recorder for MockReplicaSetEventWatcher.
type MockReplicaSetEventWatcherMockRecorder struct {
	mock *MockReplicaSetEventWatcher
}

// NewMockReplicaSetEventWatcher creates a new mock instance.
func NewMockReplicaSetEventWatcher(ctrl *gomock.Controller) *MockReplicaSetEventWatcher {
	mock := &MockReplicaSetEventWatcher{ctrl: ctrl}
	mock.recorder = &MockReplicaSetEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetEventWatcher) EXPECT() *MockReplicaSetEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockReplicaSetEventWatcher) AddEventHandler(ctx context.Context, h controller.ReplicaSetEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockReplicaSetEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockReplicaSetEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockDaemonSetEventHandler is a mock of DaemonSetEventHandler interface.
type MockDaemonSetEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetEventHandlerMockRecorder
}

// MockDaemonSetEventHandlerMockRecorder is the mock recorder for MockDaemonSetEventHandler.
type MockDaemonSetEventHandlerMockRecorder struct {
	mock *MockDaemonSetEventHandler
}

// NewMockDaemonSetEventHandler creates a new mock instance.
func NewMockDaemonSetEventHandler(ctrl *gomock.Controller) *MockDaemonSetEventHandler {
	mock := &MockDaemonSetEventHandler{ctrl: ctrl}
	mock.recorder = &MockDaemonSetEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetEventHandler) EXPECT() *MockDaemonSetEventHandlerMockRecorder {
	return m.recorder
}

// CreateDaemonSet mocks base method.
func (m *MockDaemonSetEventHandler) CreateDaemonSet(obj *v1.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDaemonSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDaemonSet indicates an expected call of CreateDaemonSet.
func (mr *MockDaemonSetEventHandlerMockRecorder) CreateDaemonSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDaemonSet", reflect.TypeOf((*MockDaemonSetEventHandler)(nil).CreateDaemonSet), obj)
}

// DeleteDaemonSet mocks base method.
func (m *MockDaemonSetEventHandler) DeleteDaemonSet(obj *v1.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDaemonSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDaemonSet indicates an expected call of DeleteDaemonSet.
func (mr *MockDaemonSetEventHandlerMockRecorder) DeleteDaemonSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDaemonSet", reflect.TypeOf((*MockDaemonSetEventHandler)(nil).DeleteDaemonSet), obj)
}

// GenericDaemonSet mocks base method.
func (m *MockDaemonSetEventHandler) GenericDaemonSet(obj *v1.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericDaemonSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericDaemonSet indicates an expected call of GenericDaemonSet.
func (mr *MockDaemonSetEventHandlerMockRecorder) GenericDaemonSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericDaemonSet", reflect.TypeOf((*MockDaemonSetEventHandler)(nil).GenericDaemonSet), obj)
}

// UpdateDaemonSet mocks base method.
func (m *MockDaemonSetEventHandler) UpdateDaemonSet(old, new *v1.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDaemonSet", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDaemonSet indicates an expected call of UpdateDaemonSet.
func (mr *MockDaemonSetEventHandlerMockRecorder) UpdateDaemonSet(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDaemonSet", reflect.TypeOf((*MockDaemonSetEventHandler)(nil).UpdateDaemonSet), old, new)
}

// MockDaemonSetEventWatcher is a mock of DaemonSetEventWatcher interface.
type MockDaemonSetEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetEventWatcherMockRecorder
}

// MockDaemonSetEventWatcherMockRecorder is the mock recorder for MockDaemonSetEventWatcher.
type MockDaemonSetEventWatcherMockRecorder struct {
	mock *MockDaemonSetEventWatcher
}

// NewMockDaemonSetEventWatcher creates a new mock instance.
func NewMockDaemonSetEventWatcher(ctrl *gomock.Controller) *MockDaemonSetEventWatcher {
	mock := &MockDaemonSetEventWatcher{ctrl: ctrl}
	mock.recorder = &MockDaemonSetEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetEventWatcher) EXPECT() *MockDaemonSetEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockDaemonSetEventWatcher) AddEventHandler(ctx context.Context, h controller.DaemonSetEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockDaemonSetEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockDaemonSetEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockStatefulSetEventHandler is a mock of StatefulSetEventHandler interface.
type MockStatefulSetEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockStatefulSetEventHandlerMockRecorder
}

// MockStatefulSetEventHandlerMockRecorder is the mock recorder for MockStatefulSetEventHandler.
type MockStatefulSetEventHandlerMockRecorder struct {
	mock *MockStatefulSetEventHandler
}

// NewMockStatefulSetEventHandler creates a new mock instance.
func NewMockStatefulSetEventHandler(ctrl *gomock.Controller) *MockStatefulSetEventHandler {
	mock := &MockStatefulSetEventHandler{ctrl: ctrl}
	mock.recorder = &MockStatefulSetEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatefulSetEventHandler) EXPECT() *MockStatefulSetEventHandlerMockRecorder {
	return m.recorder
}

// CreateStatefulSet mocks base method.
func (m *MockStatefulSetEventHandler) CreateStatefulSet(obj *v1.StatefulSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStatefulSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStatefulSet indicates an expected call of CreateStatefulSet.
func (mr *MockStatefulSetEventHandlerMockRecorder) CreateStatefulSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStatefulSet", reflect.TypeOf((*MockStatefulSetEventHandler)(nil).CreateStatefulSet), obj)
}

// DeleteStatefulSet mocks base method.
func (m *MockStatefulSetEventHandler) DeleteStatefulSet(obj *v1.StatefulSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStatefulSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStatefulSet indicates an expected call of DeleteStatefulSet.
func (mr *MockStatefulSetEventHandlerMockRecorder) DeleteStatefulSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStatefulSet", reflect.TypeOf((*MockStatefulSetEventHandler)(nil).DeleteStatefulSet), obj)
}

// GenericStatefulSet mocks base method.
func (m *MockStatefulSetEventHandler) GenericStatefulSet(obj *v1.StatefulSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericStatefulSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericStatefulSet indicates an expected call of GenericStatefulSet.
func (mr *MockStatefulSetEventHandlerMockRecorder) GenericStatefulSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericStatefulSet", reflect.TypeOf((*MockStatefulSetEventHandler)(nil).GenericStatefulSet), obj)
}

// UpdateStatefulSet mocks base method.
func (m *MockStatefulSetEventHandler) UpdateStatefulSet(old, new *v1.StatefulSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatefulSet", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatefulSet indicates an expected call of UpdateStatefulSet.
func (mr *MockStatefulSetEventHandlerMockRecorder) UpdateStatefulSet(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatefulSet", reflect.TypeOf((*MockStatefulSetEventHandler)(nil).UpdateStatefulSet), old, new)
}

// MockStatefulSetEventWatcher is a mock of StatefulSetEventWatcher interface.
type MockStatefulSetEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockStatefulSetEventWatcherMockRecorder
}

// MockStatefulSetEventWatcherMockRecorder is the mock recorder for MockStatefulSetEventWatcher.
type MockStatefulSetEventWatcherMockRecorder struct {
	mock *MockStatefulSetEventWatcher
}

// NewMockStatefulSetEventWatcher creates a new mock instance.
func NewMockStatefulSetEventWatcher(ctrl *gomock.Controller) *MockStatefulSetEventWatcher {
	mock := &MockStatefulSetEventWatcher{ctrl: ctrl}
	mock.recorder = &MockStatefulSetEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatefulSetEventWatcher) EXPECT() *MockStatefulSetEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockStatefulSetEventWatcher) AddEventHandler(ctx context.Context, h controller.StatefulSetEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockStatefulSetEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockStatefulSetEventWatcher)(nil).AddEventHandler), varargs...)
}
