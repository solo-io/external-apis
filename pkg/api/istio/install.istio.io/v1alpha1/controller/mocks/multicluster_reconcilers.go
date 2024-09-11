// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/install.istio.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1alpha1 "istio.io/istio/operator/pkg/apis"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterIstioOperatorReconciler is a mock of MulticlusterIstioOperatorReconciler interface.
type MockMulticlusterIstioOperatorReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterIstioOperatorReconcilerMockRecorder
}

// MockMulticlusterIstioOperatorReconcilerMockRecorder is the mock recorder for MockMulticlusterIstioOperatorReconciler.
type MockMulticlusterIstioOperatorReconcilerMockRecorder struct {
	mock *MockMulticlusterIstioOperatorReconciler
}

// NewMockMulticlusterIstioOperatorReconciler creates a new mock instance.
func NewMockMulticlusterIstioOperatorReconciler(ctrl *gomock.Controller) *MockMulticlusterIstioOperatorReconciler {
	mock := &MockMulticlusterIstioOperatorReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterIstioOperatorReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterIstioOperatorReconciler) EXPECT() *MockMulticlusterIstioOperatorReconcilerMockRecorder {
	return m.recorder
}

// ReconcileIstioOperator mocks base method.
func (m *MockMulticlusterIstioOperatorReconciler) ReconcileIstioOperator(clusterName string, obj *v1alpha1.IstioOperator) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIstioOperator", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileIstioOperator indicates an expected call of ReconcileIstioOperator.
func (mr *MockMulticlusterIstioOperatorReconcilerMockRecorder) ReconcileIstioOperator(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIstioOperator", reflect.TypeOf((*MockMulticlusterIstioOperatorReconciler)(nil).ReconcileIstioOperator), clusterName, obj)
}

// MockMulticlusterIstioOperatorDeletionReconciler is a mock of MulticlusterIstioOperatorDeletionReconciler interface.
type MockMulticlusterIstioOperatorDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterIstioOperatorDeletionReconcilerMockRecorder
}

// MockMulticlusterIstioOperatorDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterIstioOperatorDeletionReconciler.
type MockMulticlusterIstioOperatorDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterIstioOperatorDeletionReconciler
}

// NewMockMulticlusterIstioOperatorDeletionReconciler creates a new mock instance.
func NewMockMulticlusterIstioOperatorDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterIstioOperatorDeletionReconciler {
	mock := &MockMulticlusterIstioOperatorDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterIstioOperatorDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterIstioOperatorDeletionReconciler) EXPECT() *MockMulticlusterIstioOperatorDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileIstioOperatorDeletion mocks base method.
func (m *MockMulticlusterIstioOperatorDeletionReconciler) ReconcileIstioOperatorDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIstioOperatorDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileIstioOperatorDeletion indicates an expected call of ReconcileIstioOperatorDeletion.
func (mr *MockMulticlusterIstioOperatorDeletionReconcilerMockRecorder) ReconcileIstioOperatorDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIstioOperatorDeletion", reflect.TypeOf((*MockMulticlusterIstioOperatorDeletionReconciler)(nil).ReconcileIstioOperatorDeletion), clusterName, req)
}

// MockMulticlusterIstioOperatorReconcileLoop is a mock of MulticlusterIstioOperatorReconcileLoop interface.
type MockMulticlusterIstioOperatorReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterIstioOperatorReconcileLoopMockRecorder
}

// MockMulticlusterIstioOperatorReconcileLoopMockRecorder is the mock recorder for MockMulticlusterIstioOperatorReconcileLoop.
type MockMulticlusterIstioOperatorReconcileLoopMockRecorder struct {
	mock *MockMulticlusterIstioOperatorReconcileLoop
}

// NewMockMulticlusterIstioOperatorReconcileLoop creates a new mock instance.
func NewMockMulticlusterIstioOperatorReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterIstioOperatorReconcileLoop {
	mock := &MockMulticlusterIstioOperatorReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterIstioOperatorReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterIstioOperatorReconcileLoop) EXPECT() *MockMulticlusterIstioOperatorReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterIstioOperatorReconciler mocks base method.
func (m *MockMulticlusterIstioOperatorReconcileLoop) AddMulticlusterIstioOperatorReconciler(ctx context.Context, rec controller.MulticlusterIstioOperatorReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterIstioOperatorReconciler", varargs...)
}

// AddMulticlusterIstioOperatorReconciler indicates an expected call of AddMulticlusterIstioOperatorReconciler.
func (mr *MockMulticlusterIstioOperatorReconcileLoopMockRecorder) AddMulticlusterIstioOperatorReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterIstioOperatorReconciler", reflect.TypeOf((*MockMulticlusterIstioOperatorReconcileLoop)(nil).AddMulticlusterIstioOperatorReconciler), varargs...)
}
