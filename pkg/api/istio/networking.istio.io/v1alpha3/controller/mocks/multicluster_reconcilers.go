// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterEnvoyFilterReconciler is a mock of MulticlusterEnvoyFilterReconciler interface.
type MockMulticlusterEnvoyFilterReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterEnvoyFilterReconcilerMockRecorder
}

// MockMulticlusterEnvoyFilterReconcilerMockRecorder is the mock recorder for MockMulticlusterEnvoyFilterReconciler.
type MockMulticlusterEnvoyFilterReconcilerMockRecorder struct {
	mock *MockMulticlusterEnvoyFilterReconciler
}

// NewMockMulticlusterEnvoyFilterReconciler creates a new mock instance.
func NewMockMulticlusterEnvoyFilterReconciler(ctrl *gomock.Controller) *MockMulticlusterEnvoyFilterReconciler {
	mock := &MockMulticlusterEnvoyFilterReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterEnvoyFilterReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterEnvoyFilterReconciler) EXPECT() *MockMulticlusterEnvoyFilterReconcilerMockRecorder {
	return m.recorder
}

// ReconcileEnvoyFilter mocks base method.
func (m *MockMulticlusterEnvoyFilterReconciler) ReconcileEnvoyFilter(clusterName string, obj *v1alpha3.EnvoyFilter) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileEnvoyFilter", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileEnvoyFilter indicates an expected call of ReconcileEnvoyFilter.
func (mr *MockMulticlusterEnvoyFilterReconcilerMockRecorder) ReconcileEnvoyFilter(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileEnvoyFilter", reflect.TypeOf((*MockMulticlusterEnvoyFilterReconciler)(nil).ReconcileEnvoyFilter), clusterName, obj)
}

// MockMulticlusterEnvoyFilterDeletionReconciler is a mock of MulticlusterEnvoyFilterDeletionReconciler interface.
type MockMulticlusterEnvoyFilterDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterEnvoyFilterDeletionReconcilerMockRecorder
}

// MockMulticlusterEnvoyFilterDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterEnvoyFilterDeletionReconciler.
type MockMulticlusterEnvoyFilterDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterEnvoyFilterDeletionReconciler
}

// NewMockMulticlusterEnvoyFilterDeletionReconciler creates a new mock instance.
func NewMockMulticlusterEnvoyFilterDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterEnvoyFilterDeletionReconciler {
	mock := &MockMulticlusterEnvoyFilterDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterEnvoyFilterDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterEnvoyFilterDeletionReconciler) EXPECT() *MockMulticlusterEnvoyFilterDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileEnvoyFilterDeletion mocks base method.
func (m *MockMulticlusterEnvoyFilterDeletionReconciler) ReconcileEnvoyFilterDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileEnvoyFilterDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileEnvoyFilterDeletion indicates an expected call of ReconcileEnvoyFilterDeletion.
func (mr *MockMulticlusterEnvoyFilterDeletionReconcilerMockRecorder) ReconcileEnvoyFilterDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileEnvoyFilterDeletion", reflect.TypeOf((*MockMulticlusterEnvoyFilterDeletionReconciler)(nil).ReconcileEnvoyFilterDeletion), clusterName, req)
}

// MockMulticlusterEnvoyFilterReconcileLoop is a mock of MulticlusterEnvoyFilterReconcileLoop interface.
type MockMulticlusterEnvoyFilterReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterEnvoyFilterReconcileLoopMockRecorder
}

// MockMulticlusterEnvoyFilterReconcileLoopMockRecorder is the mock recorder for MockMulticlusterEnvoyFilterReconcileLoop.
type MockMulticlusterEnvoyFilterReconcileLoopMockRecorder struct {
	mock *MockMulticlusterEnvoyFilterReconcileLoop
}

// NewMockMulticlusterEnvoyFilterReconcileLoop creates a new mock instance.
func NewMockMulticlusterEnvoyFilterReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterEnvoyFilterReconcileLoop {
	mock := &MockMulticlusterEnvoyFilterReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterEnvoyFilterReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterEnvoyFilterReconcileLoop) EXPECT() *MockMulticlusterEnvoyFilterReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterEnvoyFilterReconciler mocks base method.
func (m *MockMulticlusterEnvoyFilterReconcileLoop) AddMulticlusterEnvoyFilterReconciler(ctx context.Context, rec controller.MulticlusterEnvoyFilterReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterEnvoyFilterReconciler", varargs...)
}

// AddMulticlusterEnvoyFilterReconciler indicates an expected call of AddMulticlusterEnvoyFilterReconciler.
func (mr *MockMulticlusterEnvoyFilterReconcileLoopMockRecorder) AddMulticlusterEnvoyFilterReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterEnvoyFilterReconciler", reflect.TypeOf((*MockMulticlusterEnvoyFilterReconcileLoop)(nil).AddMulticlusterEnvoyFilterReconciler), varargs...)
}
