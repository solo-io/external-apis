// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	controller "github.com/solo-io/external-apis/pkg/api/smi/specs.smi-spec.io/v1alpha3/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockHTTPRouteGroupReconciler is a mock of HTTPRouteGroupReconciler interface
type MockHTTPRouteGroupReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupReconcilerMockRecorder
}

// MockHTTPRouteGroupReconcilerMockRecorder is the mock recorder for MockHTTPRouteGroupReconciler
type MockHTTPRouteGroupReconcilerMockRecorder struct {
	mock *MockHTTPRouteGroupReconciler
}

// NewMockHTTPRouteGroupReconciler creates a new mock instance
func NewMockHTTPRouteGroupReconciler(ctrl *gomock.Controller) *MockHTTPRouteGroupReconciler {
	mock := &MockHTTPRouteGroupReconciler{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupReconciler) EXPECT() *MockHTTPRouteGroupReconcilerMockRecorder {
	return m.recorder
}

// ReconcileHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupReconciler) ReconcileHTTPRouteGroup(obj *v1alpha3.HTTPRouteGroup) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileHTTPRouteGroup", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileHTTPRouteGroup indicates an expected call of ReconcileHTTPRouteGroup
func (mr *MockHTTPRouteGroupReconcilerMockRecorder) ReconcileHTTPRouteGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupReconciler)(nil).ReconcileHTTPRouteGroup), obj)
}

// MockHTTPRouteGroupDeletionReconciler is a mock of HTTPRouteGroupDeletionReconciler interface
type MockHTTPRouteGroupDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupDeletionReconcilerMockRecorder
}

// MockHTTPRouteGroupDeletionReconcilerMockRecorder is the mock recorder for MockHTTPRouteGroupDeletionReconciler
type MockHTTPRouteGroupDeletionReconcilerMockRecorder struct {
	mock *MockHTTPRouteGroupDeletionReconciler
}

// NewMockHTTPRouteGroupDeletionReconciler creates a new mock instance
func NewMockHTTPRouteGroupDeletionReconciler(ctrl *gomock.Controller) *MockHTTPRouteGroupDeletionReconciler {
	mock := &MockHTTPRouteGroupDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupDeletionReconciler) EXPECT() *MockHTTPRouteGroupDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileHTTPRouteGroupDeletion mocks base method
func (m *MockHTTPRouteGroupDeletionReconciler) ReconcileHTTPRouteGroupDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileHTTPRouteGroupDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileHTTPRouteGroupDeletion indicates an expected call of ReconcileHTTPRouteGroupDeletion
func (mr *MockHTTPRouteGroupDeletionReconcilerMockRecorder) ReconcileHTTPRouteGroupDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileHTTPRouteGroupDeletion", reflect.TypeOf((*MockHTTPRouteGroupDeletionReconciler)(nil).ReconcileHTTPRouteGroupDeletion), req)
}

// MockHTTPRouteGroupFinalizer is a mock of HTTPRouteGroupFinalizer interface
type MockHTTPRouteGroupFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupFinalizerMockRecorder
}

// MockHTTPRouteGroupFinalizerMockRecorder is the mock recorder for MockHTTPRouteGroupFinalizer
type MockHTTPRouteGroupFinalizerMockRecorder struct {
	mock *MockHTTPRouteGroupFinalizer
}

// NewMockHTTPRouteGroupFinalizer creates a new mock instance
func NewMockHTTPRouteGroupFinalizer(ctrl *gomock.Controller) *MockHTTPRouteGroupFinalizer {
	mock := &MockHTTPRouteGroupFinalizer{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupFinalizer) EXPECT() *MockHTTPRouteGroupFinalizerMockRecorder {
	return m.recorder
}

// ReconcileHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupFinalizer) ReconcileHTTPRouteGroup(obj *v1alpha3.HTTPRouteGroup) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileHTTPRouteGroup", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileHTTPRouteGroup indicates an expected call of ReconcileHTTPRouteGroup
func (mr *MockHTTPRouteGroupFinalizerMockRecorder) ReconcileHTTPRouteGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupFinalizer)(nil).ReconcileHTTPRouteGroup), obj)
}

// HTTPRouteGroupFinalizerName mocks base method
func (m *MockHTTPRouteGroupFinalizer) HTTPRouteGroupFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPRouteGroupFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// HTTPRouteGroupFinalizerName indicates an expected call of HTTPRouteGroupFinalizerName
func (mr *MockHTTPRouteGroupFinalizerMockRecorder) HTTPRouteGroupFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPRouteGroupFinalizerName", reflect.TypeOf((*MockHTTPRouteGroupFinalizer)(nil).HTTPRouteGroupFinalizerName))
}

// FinalizeHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupFinalizer) FinalizeHTTPRouteGroup(obj *v1alpha3.HTTPRouteGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeHTTPRouteGroup", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeHTTPRouteGroup indicates an expected call of FinalizeHTTPRouteGroup
func (mr *MockHTTPRouteGroupFinalizerMockRecorder) FinalizeHTTPRouteGroup(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupFinalizer)(nil).FinalizeHTTPRouteGroup), obj)
}

// MockHTTPRouteGroupReconcileLoop is a mock of HTTPRouteGroupReconcileLoop interface
type MockHTTPRouteGroupReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupReconcileLoopMockRecorder
}

// MockHTTPRouteGroupReconcileLoopMockRecorder is the mock recorder for MockHTTPRouteGroupReconcileLoop
type MockHTTPRouteGroupReconcileLoopMockRecorder struct {
	mock *MockHTTPRouteGroupReconcileLoop
}

// NewMockHTTPRouteGroupReconcileLoop creates a new mock instance
func NewMockHTTPRouteGroupReconcileLoop(ctrl *gomock.Controller) *MockHTTPRouteGroupReconcileLoop {
	mock := &MockHTTPRouteGroupReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupReconcileLoop) EXPECT() *MockHTTPRouteGroupReconcileLoopMockRecorder {
	return m.recorder
}

// RunHTTPRouteGroupReconciler mocks base method
func (m *MockHTTPRouteGroupReconcileLoop) RunHTTPRouteGroupReconciler(ctx context.Context, rec controller.HTTPRouteGroupReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunHTTPRouteGroupReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunHTTPRouteGroupReconciler indicates an expected call of RunHTTPRouteGroupReconciler
func (mr *MockHTTPRouteGroupReconcileLoopMockRecorder) RunHTTPRouteGroupReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunHTTPRouteGroupReconciler", reflect.TypeOf((*MockHTTPRouteGroupReconcileLoop)(nil).RunHTTPRouteGroupReconciler), varargs...)
}
