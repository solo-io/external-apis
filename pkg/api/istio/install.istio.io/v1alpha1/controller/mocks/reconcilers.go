// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/install.istio.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	apis "istio.io/istio/operator/pkg/apis"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockIstioOperatorReconciler is a mock of IstioOperatorReconciler interface.
type MockIstioOperatorReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockIstioOperatorReconcilerMockRecorder
}

// MockIstioOperatorReconcilerMockRecorder is the mock recorder for MockIstioOperatorReconciler.
type MockIstioOperatorReconcilerMockRecorder struct {
	mock *MockIstioOperatorReconciler
}

// NewMockIstioOperatorReconciler creates a new mock instance.
func NewMockIstioOperatorReconciler(ctrl *gomock.Controller) *MockIstioOperatorReconciler {
	mock := &MockIstioOperatorReconciler{ctrl: ctrl}
	mock.recorder = &MockIstioOperatorReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioOperatorReconciler) EXPECT() *MockIstioOperatorReconcilerMockRecorder {
	return m.recorder
}

// ReconcileIstioOperator mocks base method.
func (m *MockIstioOperatorReconciler) ReconcileIstioOperator(obj *apis.IstioOperator) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIstioOperator", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileIstioOperator indicates an expected call of ReconcileIstioOperator.
func (mr *MockIstioOperatorReconcilerMockRecorder) ReconcileIstioOperator(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIstioOperator", reflect.TypeOf((*MockIstioOperatorReconciler)(nil).ReconcileIstioOperator), obj)
}

// MockIstioOperatorDeletionReconciler is a mock of IstioOperatorDeletionReconciler interface.
type MockIstioOperatorDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockIstioOperatorDeletionReconcilerMockRecorder
}

// MockIstioOperatorDeletionReconcilerMockRecorder is the mock recorder for MockIstioOperatorDeletionReconciler.
type MockIstioOperatorDeletionReconcilerMockRecorder struct {
	mock *MockIstioOperatorDeletionReconciler
}

// NewMockIstioOperatorDeletionReconciler creates a new mock instance.
func NewMockIstioOperatorDeletionReconciler(ctrl *gomock.Controller) *MockIstioOperatorDeletionReconciler {
	mock := &MockIstioOperatorDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockIstioOperatorDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioOperatorDeletionReconciler) EXPECT() *MockIstioOperatorDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileIstioOperatorDeletion mocks base method.
func (m *MockIstioOperatorDeletionReconciler) ReconcileIstioOperatorDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIstioOperatorDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileIstioOperatorDeletion indicates an expected call of ReconcileIstioOperatorDeletion.
func (mr *MockIstioOperatorDeletionReconcilerMockRecorder) ReconcileIstioOperatorDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIstioOperatorDeletion", reflect.TypeOf((*MockIstioOperatorDeletionReconciler)(nil).ReconcileIstioOperatorDeletion), req)
}

// MockIstioOperatorFinalizer is a mock of IstioOperatorFinalizer interface.
type MockIstioOperatorFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockIstioOperatorFinalizerMockRecorder
}

// MockIstioOperatorFinalizerMockRecorder is the mock recorder for MockIstioOperatorFinalizer.
type MockIstioOperatorFinalizerMockRecorder struct {
	mock *MockIstioOperatorFinalizer
}

// NewMockIstioOperatorFinalizer creates a new mock instance.
func NewMockIstioOperatorFinalizer(ctrl *gomock.Controller) *MockIstioOperatorFinalizer {
	mock := &MockIstioOperatorFinalizer{ctrl: ctrl}
	mock.recorder = &MockIstioOperatorFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioOperatorFinalizer) EXPECT() *MockIstioOperatorFinalizerMockRecorder {
	return m.recorder
}

// FinalizeIstioOperator mocks base method.
func (m *MockIstioOperatorFinalizer) FinalizeIstioOperator(obj *apis.IstioOperator) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeIstioOperator", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeIstioOperator indicates an expected call of FinalizeIstioOperator.
func (mr *MockIstioOperatorFinalizerMockRecorder) FinalizeIstioOperator(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeIstioOperator", reflect.TypeOf((*MockIstioOperatorFinalizer)(nil).FinalizeIstioOperator), obj)
}

// IstioOperatorFinalizerName mocks base method.
func (m *MockIstioOperatorFinalizer) IstioOperatorFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IstioOperatorFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// IstioOperatorFinalizerName indicates an expected call of IstioOperatorFinalizerName.
func (mr *MockIstioOperatorFinalizerMockRecorder) IstioOperatorFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IstioOperatorFinalizerName", reflect.TypeOf((*MockIstioOperatorFinalizer)(nil).IstioOperatorFinalizerName))
}

// ReconcileIstioOperator mocks base method.
func (m *MockIstioOperatorFinalizer) ReconcileIstioOperator(obj *apis.IstioOperator) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIstioOperator", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileIstioOperator indicates an expected call of ReconcileIstioOperator.
func (mr *MockIstioOperatorFinalizerMockRecorder) ReconcileIstioOperator(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIstioOperator", reflect.TypeOf((*MockIstioOperatorFinalizer)(nil).ReconcileIstioOperator), obj)
}

// MockIstioOperatorReconcileLoop is a mock of IstioOperatorReconcileLoop interface.
type MockIstioOperatorReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockIstioOperatorReconcileLoopMockRecorder
}

// MockIstioOperatorReconcileLoopMockRecorder is the mock recorder for MockIstioOperatorReconcileLoop.
type MockIstioOperatorReconcileLoopMockRecorder struct {
	mock *MockIstioOperatorReconcileLoop
}

// NewMockIstioOperatorReconcileLoop creates a new mock instance.
func NewMockIstioOperatorReconcileLoop(ctrl *gomock.Controller) *MockIstioOperatorReconcileLoop {
	mock := &MockIstioOperatorReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockIstioOperatorReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioOperatorReconcileLoop) EXPECT() *MockIstioOperatorReconcileLoopMockRecorder {
	return m.recorder
}

// RunIstioOperatorReconciler mocks base method.
func (m *MockIstioOperatorReconcileLoop) RunIstioOperatorReconciler(ctx context.Context, rec controller.IstioOperatorReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunIstioOperatorReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunIstioOperatorReconciler indicates an expected call of RunIstioOperatorReconciler.
func (mr *MockIstioOperatorReconcileLoopMockRecorder) RunIstioOperatorReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunIstioOperatorReconciler", reflect.TypeOf((*MockIstioOperatorReconcileLoop)(nil).RunIstioOperatorReconciler), varargs...)
}
