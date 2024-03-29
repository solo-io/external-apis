// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

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

// MockMulticlusterCiliumNetworkPolicyReconciler is a mock of MulticlusterCiliumNetworkPolicyReconciler interface.
type MockMulticlusterCiliumNetworkPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCiliumNetworkPolicyReconcilerMockRecorder
}

// MockMulticlusterCiliumNetworkPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterCiliumNetworkPolicyReconciler.
type MockMulticlusterCiliumNetworkPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterCiliumNetworkPolicyReconciler
}

// NewMockMulticlusterCiliumNetworkPolicyReconciler creates a new mock instance.
func NewMockMulticlusterCiliumNetworkPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterCiliumNetworkPolicyReconciler {
	mock := &MockMulticlusterCiliumNetworkPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCiliumNetworkPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterCiliumNetworkPolicyReconciler) EXPECT() *MockMulticlusterCiliumNetworkPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCiliumNetworkPolicy mocks base method.
func (m *MockMulticlusterCiliumNetworkPolicyReconciler) ReconcileCiliumNetworkPolicy(clusterName string, obj *v2.CiliumNetworkPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCiliumNetworkPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCiliumNetworkPolicy indicates an expected call of ReconcileCiliumNetworkPolicy.
func (mr *MockMulticlusterCiliumNetworkPolicyReconcilerMockRecorder) ReconcileCiliumNetworkPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCiliumNetworkPolicy", reflect.TypeOf((*MockMulticlusterCiliumNetworkPolicyReconciler)(nil).ReconcileCiliumNetworkPolicy), clusterName, obj)
}

// MockMulticlusterCiliumNetworkPolicyDeletionReconciler is a mock of MulticlusterCiliumNetworkPolicyDeletionReconciler interface.
type MockMulticlusterCiliumNetworkPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCiliumNetworkPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterCiliumNetworkPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterCiliumNetworkPolicyDeletionReconciler.
type MockMulticlusterCiliumNetworkPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterCiliumNetworkPolicyDeletionReconciler
}

// NewMockMulticlusterCiliumNetworkPolicyDeletionReconciler creates a new mock instance.
func NewMockMulticlusterCiliumNetworkPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterCiliumNetworkPolicyDeletionReconciler {
	mock := &MockMulticlusterCiliumNetworkPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCiliumNetworkPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterCiliumNetworkPolicyDeletionReconciler) EXPECT() *MockMulticlusterCiliumNetworkPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCiliumNetworkPolicyDeletion mocks base method.
func (m *MockMulticlusterCiliumNetworkPolicyDeletionReconciler) ReconcileCiliumNetworkPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCiliumNetworkPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileCiliumNetworkPolicyDeletion indicates an expected call of ReconcileCiliumNetworkPolicyDeletion.
func (mr *MockMulticlusterCiliumNetworkPolicyDeletionReconcilerMockRecorder) ReconcileCiliumNetworkPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCiliumNetworkPolicyDeletion", reflect.TypeOf((*MockMulticlusterCiliumNetworkPolicyDeletionReconciler)(nil).ReconcileCiliumNetworkPolicyDeletion), clusterName, req)
}

// MockMulticlusterCiliumNetworkPolicyReconcileLoop is a mock of MulticlusterCiliumNetworkPolicyReconcileLoop interface.
type MockMulticlusterCiliumNetworkPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCiliumNetworkPolicyReconcileLoopMockRecorder
}

// MockMulticlusterCiliumNetworkPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterCiliumNetworkPolicyReconcileLoop.
type MockMulticlusterCiliumNetworkPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterCiliumNetworkPolicyReconcileLoop
}

// NewMockMulticlusterCiliumNetworkPolicyReconcileLoop creates a new mock instance.
func NewMockMulticlusterCiliumNetworkPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterCiliumNetworkPolicyReconcileLoop {
	mock := &MockMulticlusterCiliumNetworkPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCiliumNetworkPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterCiliumNetworkPolicyReconcileLoop) EXPECT() *MockMulticlusterCiliumNetworkPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterCiliumNetworkPolicyReconciler mocks base method.
func (m *MockMulticlusterCiliumNetworkPolicyReconcileLoop) AddMulticlusterCiliumNetworkPolicyReconciler(ctx context.Context, rec controller.MulticlusterCiliumNetworkPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterCiliumNetworkPolicyReconciler", varargs...)
}

// AddMulticlusterCiliumNetworkPolicyReconciler indicates an expected call of AddMulticlusterCiliumNetworkPolicyReconciler.
func (mr *MockMulticlusterCiliumNetworkPolicyReconcileLoopMockRecorder) AddMulticlusterCiliumNetworkPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterCiliumNetworkPolicyReconciler", reflect.TypeOf((*MockMulticlusterCiliumNetworkPolicyReconcileLoop)(nil).AddMulticlusterCiliumNetworkPolicyReconciler), varargs...)
}
