// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/cilium/cilium.io/v2/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockCiliumNetworkPolicyReconciler is a mock of CiliumNetworkPolicyReconciler interface.
type MockCiliumNetworkPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCiliumNetworkPolicyReconcilerMockRecorder
}

// MockCiliumNetworkPolicyReconcilerMockRecorder is the mock recorder for MockCiliumNetworkPolicyReconciler.
type MockCiliumNetworkPolicyReconcilerMockRecorder struct {
	mock *MockCiliumNetworkPolicyReconciler
}

// NewMockCiliumNetworkPolicyReconciler creates a new mock instance.
func NewMockCiliumNetworkPolicyReconciler(ctrl *gomock.Controller) *MockCiliumNetworkPolicyReconciler {
	mock := &MockCiliumNetworkPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockCiliumNetworkPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCiliumNetworkPolicyReconciler) EXPECT() *MockCiliumNetworkPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyReconciler) ReconcileCiliumNetworkPolicy(obj *v2.CiliumNetworkPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCiliumNetworkPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCiliumNetworkPolicy indicates an expected call of ReconcileCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyReconcilerMockRecorder) ReconcileCiliumNetworkPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyReconciler)(nil).ReconcileCiliumNetworkPolicy), obj)
}

// MockCiliumNetworkPolicyDeletionReconciler is a mock of CiliumNetworkPolicyDeletionReconciler interface.
type MockCiliumNetworkPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCiliumNetworkPolicyDeletionReconcilerMockRecorder
}

// MockCiliumNetworkPolicyDeletionReconcilerMockRecorder is the mock recorder for MockCiliumNetworkPolicyDeletionReconciler.
type MockCiliumNetworkPolicyDeletionReconcilerMockRecorder struct {
	mock *MockCiliumNetworkPolicyDeletionReconciler
}

// NewMockCiliumNetworkPolicyDeletionReconciler creates a new mock instance.
func NewMockCiliumNetworkPolicyDeletionReconciler(ctrl *gomock.Controller) *MockCiliumNetworkPolicyDeletionReconciler {
	mock := &MockCiliumNetworkPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockCiliumNetworkPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCiliumNetworkPolicyDeletionReconciler) EXPECT() *MockCiliumNetworkPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCiliumNetworkPolicyDeletion mocks base method.
func (m *MockCiliumNetworkPolicyDeletionReconciler) ReconcileCiliumNetworkPolicyDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCiliumNetworkPolicyDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileCiliumNetworkPolicyDeletion indicates an expected call of ReconcileCiliumNetworkPolicyDeletion.
func (mr *MockCiliumNetworkPolicyDeletionReconcilerMockRecorder) ReconcileCiliumNetworkPolicyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCiliumNetworkPolicyDeletion", reflect.TypeOf((*MockCiliumNetworkPolicyDeletionReconciler)(nil).ReconcileCiliumNetworkPolicyDeletion), req)
}

// MockCiliumNetworkPolicyFinalizer is a mock of CiliumNetworkPolicyFinalizer interface.
type MockCiliumNetworkPolicyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockCiliumNetworkPolicyFinalizerMockRecorder
}

// MockCiliumNetworkPolicyFinalizerMockRecorder is the mock recorder for MockCiliumNetworkPolicyFinalizer.
type MockCiliumNetworkPolicyFinalizerMockRecorder struct {
	mock *MockCiliumNetworkPolicyFinalizer
}

// NewMockCiliumNetworkPolicyFinalizer creates a new mock instance.
func NewMockCiliumNetworkPolicyFinalizer(ctrl *gomock.Controller) *MockCiliumNetworkPolicyFinalizer {
	mock := &MockCiliumNetworkPolicyFinalizer{ctrl: ctrl}
	mock.recorder = &MockCiliumNetworkPolicyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCiliumNetworkPolicyFinalizer) EXPECT() *MockCiliumNetworkPolicyFinalizerMockRecorder {
	return m.recorder
}

// CiliumNetworkPolicyFinalizerName mocks base method.
func (m *MockCiliumNetworkPolicyFinalizer) CiliumNetworkPolicyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CiliumNetworkPolicyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// CiliumNetworkPolicyFinalizerName indicates an expected call of CiliumNetworkPolicyFinalizerName.
func (mr *MockCiliumNetworkPolicyFinalizerMockRecorder) CiliumNetworkPolicyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CiliumNetworkPolicyFinalizerName", reflect.TypeOf((*MockCiliumNetworkPolicyFinalizer)(nil).CiliumNetworkPolicyFinalizerName))
}

// FinalizeCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyFinalizer) FinalizeCiliumNetworkPolicy(obj *v2.CiliumNetworkPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeCiliumNetworkPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeCiliumNetworkPolicy indicates an expected call of FinalizeCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyFinalizerMockRecorder) FinalizeCiliumNetworkPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyFinalizer)(nil).FinalizeCiliumNetworkPolicy), obj)
}

// ReconcileCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyFinalizer) ReconcileCiliumNetworkPolicy(obj *v2.CiliumNetworkPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCiliumNetworkPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCiliumNetworkPolicy indicates an expected call of ReconcileCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyFinalizerMockRecorder) ReconcileCiliumNetworkPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyFinalizer)(nil).ReconcileCiliumNetworkPolicy), obj)
}

// MockCiliumNetworkPolicyReconcileLoop is a mock of CiliumNetworkPolicyReconcileLoop interface.
type MockCiliumNetworkPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockCiliumNetworkPolicyReconcileLoopMockRecorder
}

// MockCiliumNetworkPolicyReconcileLoopMockRecorder is the mock recorder for MockCiliumNetworkPolicyReconcileLoop.
type MockCiliumNetworkPolicyReconcileLoopMockRecorder struct {
	mock *MockCiliumNetworkPolicyReconcileLoop
}

// NewMockCiliumNetworkPolicyReconcileLoop creates a new mock instance.
func NewMockCiliumNetworkPolicyReconcileLoop(ctrl *gomock.Controller) *MockCiliumNetworkPolicyReconcileLoop {
	mock := &MockCiliumNetworkPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockCiliumNetworkPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCiliumNetworkPolicyReconcileLoop) EXPECT() *MockCiliumNetworkPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// RunCiliumNetworkPolicyReconciler mocks base method.
func (m *MockCiliumNetworkPolicyReconcileLoop) RunCiliumNetworkPolicyReconciler(ctx context.Context, rec controller.CiliumNetworkPolicyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCiliumNetworkPolicyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCiliumNetworkPolicyReconciler indicates an expected call of RunCiliumNetworkPolicyReconciler.
func (mr *MockCiliumNetworkPolicyReconcileLoopMockRecorder) RunCiliumNetworkPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCiliumNetworkPolicyReconciler", reflect.TypeOf((*MockCiliumNetworkPolicyReconcileLoop)(nil).RunCiliumNetworkPolicyReconciler), varargs...)
}
