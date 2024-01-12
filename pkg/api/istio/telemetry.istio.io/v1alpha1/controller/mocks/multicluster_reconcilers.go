// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/telemetry.istio.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1alpha1 "istio.io/client-go/pkg/apis/telemetry/v1alpha1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterTelemetryReconciler is a mock of MulticlusterTelemetryReconciler interface.
type MockMulticlusterTelemetryReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTelemetryReconcilerMockRecorder
}

// MockMulticlusterTelemetryReconcilerMockRecorder is the mock recorder for MockMulticlusterTelemetryReconciler.
type MockMulticlusterTelemetryReconcilerMockRecorder struct {
	mock *MockMulticlusterTelemetryReconciler
}

// NewMockMulticlusterTelemetryReconciler creates a new mock instance.
func NewMockMulticlusterTelemetryReconciler(ctrl *gomock.Controller) *MockMulticlusterTelemetryReconciler {
	mock := &MockMulticlusterTelemetryReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTelemetryReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterTelemetryReconciler) EXPECT() *MockMulticlusterTelemetryReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTelemetry mocks base method.
func (m *MockMulticlusterTelemetryReconciler) ReconcileTelemetry(clusterName string, obj *v1alpha1.Telemetry) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTelemetry", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTelemetry indicates an expected call of ReconcileTelemetry.
func (mr *MockMulticlusterTelemetryReconcilerMockRecorder) ReconcileTelemetry(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTelemetry", reflect.TypeOf((*MockMulticlusterTelemetryReconciler)(nil).ReconcileTelemetry), clusterName, obj)
}

// MockMulticlusterTelemetryDeletionReconciler is a mock of MulticlusterTelemetryDeletionReconciler interface.
type MockMulticlusterTelemetryDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTelemetryDeletionReconcilerMockRecorder
}

// MockMulticlusterTelemetryDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterTelemetryDeletionReconciler.
type MockMulticlusterTelemetryDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterTelemetryDeletionReconciler
}

// NewMockMulticlusterTelemetryDeletionReconciler creates a new mock instance.
func NewMockMulticlusterTelemetryDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterTelemetryDeletionReconciler {
	mock := &MockMulticlusterTelemetryDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTelemetryDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterTelemetryDeletionReconciler) EXPECT() *MockMulticlusterTelemetryDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTelemetryDeletion mocks base method.
func (m *MockMulticlusterTelemetryDeletionReconciler) ReconcileTelemetryDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTelemetryDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileTelemetryDeletion indicates an expected call of ReconcileTelemetryDeletion.
func (mr *MockMulticlusterTelemetryDeletionReconcilerMockRecorder) ReconcileTelemetryDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTelemetryDeletion", reflect.TypeOf((*MockMulticlusterTelemetryDeletionReconciler)(nil).ReconcileTelemetryDeletion), clusterName, req)
}

// MockMulticlusterTelemetryReconcileLoop is a mock of MulticlusterTelemetryReconcileLoop interface.
type MockMulticlusterTelemetryReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTelemetryReconcileLoopMockRecorder
}

// MockMulticlusterTelemetryReconcileLoopMockRecorder is the mock recorder for MockMulticlusterTelemetryReconcileLoop.
type MockMulticlusterTelemetryReconcileLoopMockRecorder struct {
	mock *MockMulticlusterTelemetryReconcileLoop
}

// NewMockMulticlusterTelemetryReconcileLoop creates a new mock instance.
func NewMockMulticlusterTelemetryReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterTelemetryReconcileLoop {
	mock := &MockMulticlusterTelemetryReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTelemetryReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterTelemetryReconcileLoop) EXPECT() *MockMulticlusterTelemetryReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterTelemetryReconciler mocks base method.
func (m *MockMulticlusterTelemetryReconcileLoop) AddMulticlusterTelemetryReconciler(ctx context.Context, rec controller.MulticlusterTelemetryReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterTelemetryReconciler", varargs...)
}

// AddMulticlusterTelemetryReconciler indicates an expected call of AddMulticlusterTelemetryReconciler.
func (mr *MockMulticlusterTelemetryReconcileLoopMockRecorder) AddMulticlusterTelemetryReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterTelemetryReconciler", reflect.TypeOf((*MockMulticlusterTelemetryReconcileLoop)(nil).AddMulticlusterTelemetryReconciler), varargs...)
}
