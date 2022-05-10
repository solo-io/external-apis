// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1beta1/controller"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockDestinationRuleEventHandler is a mock of DestinationRuleEventHandler interface.
type MockDestinationRuleEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationRuleEventHandlerMockRecorder
}

// MockDestinationRuleEventHandlerMockRecorder is the mock recorder for MockDestinationRuleEventHandler.
type MockDestinationRuleEventHandlerMockRecorder struct {
	mock *MockDestinationRuleEventHandler
}

// NewMockDestinationRuleEventHandler creates a new mock instance.
func NewMockDestinationRuleEventHandler(ctrl *gomock.Controller) *MockDestinationRuleEventHandler {
	mock := &MockDestinationRuleEventHandler{ctrl: ctrl}
	mock.recorder = &MockDestinationRuleEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationRuleEventHandler) EXPECT() *MockDestinationRuleEventHandlerMockRecorder {
	return m.recorder
}

// CreateDestinationRule mocks base method.
func (m *MockDestinationRuleEventHandler) CreateDestinationRule(obj *v1beta1.DestinationRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDestinationRule", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDestinationRule indicates an expected call of CreateDestinationRule.
func (mr *MockDestinationRuleEventHandlerMockRecorder) CreateDestinationRule(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDestinationRule", reflect.TypeOf((*MockDestinationRuleEventHandler)(nil).CreateDestinationRule), obj)
}

// DeleteDestinationRule mocks base method.
func (m *MockDestinationRuleEventHandler) DeleteDestinationRule(obj *v1beta1.DestinationRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDestinationRule", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDestinationRule indicates an expected call of DeleteDestinationRule.
func (mr *MockDestinationRuleEventHandlerMockRecorder) DeleteDestinationRule(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDestinationRule", reflect.TypeOf((*MockDestinationRuleEventHandler)(nil).DeleteDestinationRule), obj)
}

// GenericDestinationRule mocks base method.
func (m *MockDestinationRuleEventHandler) GenericDestinationRule(obj *v1beta1.DestinationRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericDestinationRule", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericDestinationRule indicates an expected call of GenericDestinationRule.
func (mr *MockDestinationRuleEventHandlerMockRecorder) GenericDestinationRule(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericDestinationRule", reflect.TypeOf((*MockDestinationRuleEventHandler)(nil).GenericDestinationRule), obj)
}

// UpdateDestinationRule mocks base method.
func (m *MockDestinationRuleEventHandler) UpdateDestinationRule(old, new *v1beta1.DestinationRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDestinationRule", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDestinationRule indicates an expected call of UpdateDestinationRule.
func (mr *MockDestinationRuleEventHandlerMockRecorder) UpdateDestinationRule(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDestinationRule", reflect.TypeOf((*MockDestinationRuleEventHandler)(nil).UpdateDestinationRule), old, new)
}

// MockDestinationRuleEventWatcher is a mock of DestinationRuleEventWatcher interface.
type MockDestinationRuleEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationRuleEventWatcherMockRecorder
}

// MockDestinationRuleEventWatcherMockRecorder is the mock recorder for MockDestinationRuleEventWatcher.
type MockDestinationRuleEventWatcherMockRecorder struct {
	mock *MockDestinationRuleEventWatcher
}

// NewMockDestinationRuleEventWatcher creates a new mock instance.
func NewMockDestinationRuleEventWatcher(ctrl *gomock.Controller) *MockDestinationRuleEventWatcher {
	mock := &MockDestinationRuleEventWatcher{ctrl: ctrl}
	mock.recorder = &MockDestinationRuleEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationRuleEventWatcher) EXPECT() *MockDestinationRuleEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockDestinationRuleEventWatcher) AddEventHandler(ctx context.Context, h controller.DestinationRuleEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockDestinationRuleEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockDestinationRuleEventWatcher)(nil).AddEventHandler), varargs...)
}

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

// MockServiceEntryEventHandler is a mock of ServiceEntryEventHandler interface.
type MockServiceEntryEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntryEventHandlerMockRecorder
}

// MockServiceEntryEventHandlerMockRecorder is the mock recorder for MockServiceEntryEventHandler.
type MockServiceEntryEventHandlerMockRecorder struct {
	mock *MockServiceEntryEventHandler
}

// NewMockServiceEntryEventHandler creates a new mock instance.
func NewMockServiceEntryEventHandler(ctrl *gomock.Controller) *MockServiceEntryEventHandler {
	mock := &MockServiceEntryEventHandler{ctrl: ctrl}
	mock.recorder = &MockServiceEntryEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEntryEventHandler) EXPECT() *MockServiceEntryEventHandlerMockRecorder {
	return m.recorder
}

// CreateServiceEntry mocks base method.
func (m *MockServiceEntryEventHandler) CreateServiceEntry(obj *v1beta1.ServiceEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceEntry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateServiceEntry indicates an expected call of CreateServiceEntry.
func (mr *MockServiceEntryEventHandlerMockRecorder) CreateServiceEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceEntry", reflect.TypeOf((*MockServiceEntryEventHandler)(nil).CreateServiceEntry), obj)
}

// DeleteServiceEntry mocks base method.
func (m *MockServiceEntryEventHandler) DeleteServiceEntry(obj *v1beta1.ServiceEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceEntry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceEntry indicates an expected call of DeleteServiceEntry.
func (mr *MockServiceEntryEventHandlerMockRecorder) DeleteServiceEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceEntry", reflect.TypeOf((*MockServiceEntryEventHandler)(nil).DeleteServiceEntry), obj)
}

// GenericServiceEntry mocks base method.
func (m *MockServiceEntryEventHandler) GenericServiceEntry(obj *v1beta1.ServiceEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericServiceEntry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericServiceEntry indicates an expected call of GenericServiceEntry.
func (mr *MockServiceEntryEventHandlerMockRecorder) GenericServiceEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericServiceEntry", reflect.TypeOf((*MockServiceEntryEventHandler)(nil).GenericServiceEntry), obj)
}

// UpdateServiceEntry mocks base method.
func (m *MockServiceEntryEventHandler) UpdateServiceEntry(old, new *v1beta1.ServiceEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceEntry", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateServiceEntry indicates an expected call of UpdateServiceEntry.
func (mr *MockServiceEntryEventHandlerMockRecorder) UpdateServiceEntry(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceEntry", reflect.TypeOf((*MockServiceEntryEventHandler)(nil).UpdateServiceEntry), old, new)
}

// MockServiceEntryEventWatcher is a mock of ServiceEntryEventWatcher interface.
type MockServiceEntryEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntryEventWatcherMockRecorder
}

// MockServiceEntryEventWatcherMockRecorder is the mock recorder for MockServiceEntryEventWatcher.
type MockServiceEntryEventWatcherMockRecorder struct {
	mock *MockServiceEntryEventWatcher
}

// NewMockServiceEntryEventWatcher creates a new mock instance.
func NewMockServiceEntryEventWatcher(ctrl *gomock.Controller) *MockServiceEntryEventWatcher {
	mock := &MockServiceEntryEventWatcher{ctrl: ctrl}
	mock.recorder = &MockServiceEntryEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEntryEventWatcher) EXPECT() *MockServiceEntryEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockServiceEntryEventWatcher) AddEventHandler(ctx context.Context, h controller.ServiceEntryEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockServiceEntryEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockServiceEntryEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockWorkloadEntryEventHandler is a mock of WorkloadEntryEventHandler interface.
type MockWorkloadEntryEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadEntryEventHandlerMockRecorder
}

// MockWorkloadEntryEventHandlerMockRecorder is the mock recorder for MockWorkloadEntryEventHandler.
type MockWorkloadEntryEventHandlerMockRecorder struct {
	mock *MockWorkloadEntryEventHandler
}

// NewMockWorkloadEntryEventHandler creates a new mock instance.
func NewMockWorkloadEntryEventHandler(ctrl *gomock.Controller) *MockWorkloadEntryEventHandler {
	mock := &MockWorkloadEntryEventHandler{ctrl: ctrl}
	mock.recorder = &MockWorkloadEntryEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadEntryEventHandler) EXPECT() *MockWorkloadEntryEventHandlerMockRecorder {
	return m.recorder
}

// CreateWorkloadEntry mocks base method.
func (m *MockWorkloadEntryEventHandler) CreateWorkloadEntry(obj *v1beta1.WorkloadEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkloadEntry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWorkloadEntry indicates an expected call of CreateWorkloadEntry.
func (mr *MockWorkloadEntryEventHandlerMockRecorder) CreateWorkloadEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkloadEntry", reflect.TypeOf((*MockWorkloadEntryEventHandler)(nil).CreateWorkloadEntry), obj)
}

// DeleteWorkloadEntry mocks base method.
func (m *MockWorkloadEntryEventHandler) DeleteWorkloadEntry(obj *v1beta1.WorkloadEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkloadEntry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkloadEntry indicates an expected call of DeleteWorkloadEntry.
func (mr *MockWorkloadEntryEventHandlerMockRecorder) DeleteWorkloadEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkloadEntry", reflect.TypeOf((*MockWorkloadEntryEventHandler)(nil).DeleteWorkloadEntry), obj)
}

// GenericWorkloadEntry mocks base method.
func (m *MockWorkloadEntryEventHandler) GenericWorkloadEntry(obj *v1beta1.WorkloadEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericWorkloadEntry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericWorkloadEntry indicates an expected call of GenericWorkloadEntry.
func (mr *MockWorkloadEntryEventHandlerMockRecorder) GenericWorkloadEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericWorkloadEntry", reflect.TypeOf((*MockWorkloadEntryEventHandler)(nil).GenericWorkloadEntry), obj)
}

// UpdateWorkloadEntry mocks base method.
func (m *MockWorkloadEntryEventHandler) UpdateWorkloadEntry(old, new *v1beta1.WorkloadEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkloadEntry", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkloadEntry indicates an expected call of UpdateWorkloadEntry.
func (mr *MockWorkloadEntryEventHandlerMockRecorder) UpdateWorkloadEntry(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkloadEntry", reflect.TypeOf((*MockWorkloadEntryEventHandler)(nil).UpdateWorkloadEntry), old, new)
}

// MockWorkloadEntryEventWatcher is a mock of WorkloadEntryEventWatcher interface.
type MockWorkloadEntryEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadEntryEventWatcherMockRecorder
}

// MockWorkloadEntryEventWatcherMockRecorder is the mock recorder for MockWorkloadEntryEventWatcher.
type MockWorkloadEntryEventWatcherMockRecorder struct {
	mock *MockWorkloadEntryEventWatcher
}

// NewMockWorkloadEntryEventWatcher creates a new mock instance.
func NewMockWorkloadEntryEventWatcher(ctrl *gomock.Controller) *MockWorkloadEntryEventWatcher {
	mock := &MockWorkloadEntryEventWatcher{ctrl: ctrl}
	mock.recorder = &MockWorkloadEntryEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadEntryEventWatcher) EXPECT() *MockWorkloadEntryEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockWorkloadEntryEventWatcher) AddEventHandler(ctx context.Context, h controller.WorkloadEntryEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockWorkloadEntryEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockWorkloadEntryEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockVirtualServiceEventHandler is a mock of VirtualServiceEventHandler interface.
type MockVirtualServiceEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceEventHandlerMockRecorder
}

// MockVirtualServiceEventHandlerMockRecorder is the mock recorder for MockVirtualServiceEventHandler.
type MockVirtualServiceEventHandlerMockRecorder struct {
	mock *MockVirtualServiceEventHandler
}

// NewMockVirtualServiceEventHandler creates a new mock instance.
func NewMockVirtualServiceEventHandler(ctrl *gomock.Controller) *MockVirtualServiceEventHandler {
	mock := &MockVirtualServiceEventHandler{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualServiceEventHandler) EXPECT() *MockVirtualServiceEventHandlerMockRecorder {
	return m.recorder
}

// CreateVirtualService mocks base method.
func (m *MockVirtualServiceEventHandler) CreateVirtualService(obj *v1beta1.VirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVirtualService indicates an expected call of CreateVirtualService.
func (mr *MockVirtualServiceEventHandlerMockRecorder) CreateVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualService", reflect.TypeOf((*MockVirtualServiceEventHandler)(nil).CreateVirtualService), obj)
}

// DeleteVirtualService mocks base method.
func (m *MockVirtualServiceEventHandler) DeleteVirtualService(obj *v1beta1.VirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVirtualService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualService indicates an expected call of DeleteVirtualService.
func (mr *MockVirtualServiceEventHandlerMockRecorder) DeleteVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualService", reflect.TypeOf((*MockVirtualServiceEventHandler)(nil).DeleteVirtualService), obj)
}

// GenericVirtualService mocks base method.
func (m *MockVirtualServiceEventHandler) GenericVirtualService(obj *v1beta1.VirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericVirtualService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericVirtualService indicates an expected call of GenericVirtualService.
func (mr *MockVirtualServiceEventHandlerMockRecorder) GenericVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericVirtualService", reflect.TypeOf((*MockVirtualServiceEventHandler)(nil).GenericVirtualService), obj)
}

// UpdateVirtualService mocks base method.
func (m *MockVirtualServiceEventHandler) UpdateVirtualService(old, new *v1beta1.VirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVirtualService", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVirtualService indicates an expected call of UpdateVirtualService.
func (mr *MockVirtualServiceEventHandlerMockRecorder) UpdateVirtualService(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVirtualService", reflect.TypeOf((*MockVirtualServiceEventHandler)(nil).UpdateVirtualService), old, new)
}

// MockVirtualServiceEventWatcher is a mock of VirtualServiceEventWatcher interface.
type MockVirtualServiceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceEventWatcherMockRecorder
}

// MockVirtualServiceEventWatcherMockRecorder is the mock recorder for MockVirtualServiceEventWatcher.
type MockVirtualServiceEventWatcherMockRecorder struct {
	mock *MockVirtualServiceEventWatcher
}

// NewMockVirtualServiceEventWatcher creates a new mock instance.
func NewMockVirtualServiceEventWatcher(ctrl *gomock.Controller) *MockVirtualServiceEventWatcher {
	mock := &MockVirtualServiceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualServiceEventWatcher) EXPECT() *MockVirtualServiceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockVirtualServiceEventWatcher) AddEventHandler(ctx context.Context, h controller.VirtualServiceEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockVirtualServiceEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockVirtualServiceEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockSidecarEventHandler is a mock of SidecarEventHandler interface.
type MockSidecarEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarEventHandlerMockRecorder
}

// MockSidecarEventHandlerMockRecorder is the mock recorder for MockSidecarEventHandler.
type MockSidecarEventHandlerMockRecorder struct {
	mock *MockSidecarEventHandler
}

// NewMockSidecarEventHandler creates a new mock instance.
func NewMockSidecarEventHandler(ctrl *gomock.Controller) *MockSidecarEventHandler {
	mock := &MockSidecarEventHandler{ctrl: ctrl}
	mock.recorder = &MockSidecarEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarEventHandler) EXPECT() *MockSidecarEventHandlerMockRecorder {
	return m.recorder
}

// CreateSidecar mocks base method.
func (m *MockSidecarEventHandler) CreateSidecar(obj *v1beta1.Sidecar) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSidecar", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSidecar indicates an expected call of CreateSidecar.
func (mr *MockSidecarEventHandlerMockRecorder) CreateSidecar(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSidecar", reflect.TypeOf((*MockSidecarEventHandler)(nil).CreateSidecar), obj)
}

// DeleteSidecar mocks base method.
func (m *MockSidecarEventHandler) DeleteSidecar(obj *v1beta1.Sidecar) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSidecar", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSidecar indicates an expected call of DeleteSidecar.
func (mr *MockSidecarEventHandlerMockRecorder) DeleteSidecar(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSidecar", reflect.TypeOf((*MockSidecarEventHandler)(nil).DeleteSidecar), obj)
}

// GenericSidecar mocks base method.
func (m *MockSidecarEventHandler) GenericSidecar(obj *v1beta1.Sidecar) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericSidecar", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericSidecar indicates an expected call of GenericSidecar.
func (mr *MockSidecarEventHandlerMockRecorder) GenericSidecar(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericSidecar", reflect.TypeOf((*MockSidecarEventHandler)(nil).GenericSidecar), obj)
}

// UpdateSidecar mocks base method.
func (m *MockSidecarEventHandler) UpdateSidecar(old, new *v1beta1.Sidecar) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSidecar", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSidecar indicates an expected call of UpdateSidecar.
func (mr *MockSidecarEventHandlerMockRecorder) UpdateSidecar(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSidecar", reflect.TypeOf((*MockSidecarEventHandler)(nil).UpdateSidecar), old, new)
}

// MockSidecarEventWatcher is a mock of SidecarEventWatcher interface.
type MockSidecarEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarEventWatcherMockRecorder
}

// MockSidecarEventWatcherMockRecorder is the mock recorder for MockSidecarEventWatcher.
type MockSidecarEventWatcherMockRecorder struct {
	mock *MockSidecarEventWatcher
}

// NewMockSidecarEventWatcher creates a new mock instance.
func NewMockSidecarEventWatcher(ctrl *gomock.Controller) *MockSidecarEventWatcher {
	mock := &MockSidecarEventWatcher{ctrl: ctrl}
	mock.recorder = &MockSidecarEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarEventWatcher) EXPECT() *MockSidecarEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockSidecarEventWatcher) AddEventHandler(ctx context.Context, h controller.SidecarEventHandler, predicates ...predicate.Predicate) error {
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
func (mr *MockSidecarEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockSidecarEventWatcher)(nil).AddEventHandler), varargs...)
}
