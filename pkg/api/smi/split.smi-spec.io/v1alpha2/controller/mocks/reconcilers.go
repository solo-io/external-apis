// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"
	controller "github.com/solo-io/external-apis/pkg/api/smi/split.smi-spec.io/v1alpha2/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockTrafficSplitReconciler is a mock of TrafficSplitReconciler interface.
type MockTrafficSplitReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitReconcilerMockRecorder
}

// MockTrafficSplitReconcilerMockRecorder is the mock recorder for MockTrafficSplitReconciler.
type MockTrafficSplitReconcilerMockRecorder struct {
	mock *MockTrafficSplitReconciler
}

// NewMockTrafficSplitReconciler creates a new mock instance.
func NewMockTrafficSplitReconciler(ctrl *gomock.Controller) *MockTrafficSplitReconciler {
	mock := &MockTrafficSplitReconciler{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitReconciler) EXPECT() *MockTrafficSplitReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficSplit mocks base method.
func (m *MockTrafficSplitReconciler) ReconcileTrafficSplit(obj *v1alpha2.TrafficSplit) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficSplit", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficSplit indicates an expected call of ReconcileTrafficSplit.
func (mr *MockTrafficSplitReconcilerMockRecorder) ReconcileTrafficSplit(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficSplit", reflect.TypeOf((*MockTrafficSplitReconciler)(nil).ReconcileTrafficSplit), obj)
}

// MockTrafficSplitDeletionReconciler is a mock of TrafficSplitDeletionReconciler interface.
type MockTrafficSplitDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitDeletionReconcilerMockRecorder
}

// MockTrafficSplitDeletionReconcilerMockRecorder is the mock recorder for MockTrafficSplitDeletionReconciler.
type MockTrafficSplitDeletionReconcilerMockRecorder struct {
	mock *MockTrafficSplitDeletionReconciler
}

// NewMockTrafficSplitDeletionReconciler creates a new mock instance.
func NewMockTrafficSplitDeletionReconciler(ctrl *gomock.Controller) *MockTrafficSplitDeletionReconciler {
	mock := &MockTrafficSplitDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitDeletionReconciler) EXPECT() *MockTrafficSplitDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficSplitDeletion mocks base method.
func (m *MockTrafficSplitDeletionReconciler) ReconcileTrafficSplitDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficSplitDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileTrafficSplitDeletion indicates an expected call of ReconcileTrafficSplitDeletion.
func (mr *MockTrafficSplitDeletionReconcilerMockRecorder) ReconcileTrafficSplitDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficSplitDeletion", reflect.TypeOf((*MockTrafficSplitDeletionReconciler)(nil).ReconcileTrafficSplitDeletion), req)
}

// MockTrafficSplitFinalizer is a mock of TrafficSplitFinalizer interface.
type MockTrafficSplitFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitFinalizerMockRecorder
}

// MockTrafficSplitFinalizerMockRecorder is the mock recorder for MockTrafficSplitFinalizer.
type MockTrafficSplitFinalizerMockRecorder struct {
	mock *MockTrafficSplitFinalizer
}

// NewMockTrafficSplitFinalizer creates a new mock instance.
func NewMockTrafficSplitFinalizer(ctrl *gomock.Controller) *MockTrafficSplitFinalizer {
	mock := &MockTrafficSplitFinalizer{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitFinalizer) EXPECT() *MockTrafficSplitFinalizerMockRecorder {
	return m.recorder
}

// FinalizeTrafficSplit mocks base method.
func (m *MockTrafficSplitFinalizer) FinalizeTrafficSplit(obj *v1alpha2.TrafficSplit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeTrafficSplit", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeTrafficSplit indicates an expected call of FinalizeTrafficSplit.
func (mr *MockTrafficSplitFinalizerMockRecorder) FinalizeTrafficSplit(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeTrafficSplit", reflect.TypeOf((*MockTrafficSplitFinalizer)(nil).FinalizeTrafficSplit), obj)
}

// ReconcileTrafficSplit mocks base method.
func (m *MockTrafficSplitFinalizer) ReconcileTrafficSplit(obj *v1alpha2.TrafficSplit) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficSplit", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficSplit indicates an expected call of ReconcileTrafficSplit.
func (mr *MockTrafficSplitFinalizerMockRecorder) ReconcileTrafficSplit(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficSplit", reflect.TypeOf((*MockTrafficSplitFinalizer)(nil).ReconcileTrafficSplit), obj)
}

// TrafficSplitFinalizerName mocks base method.
func (m *MockTrafficSplitFinalizer) TrafficSplitFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficSplitFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TrafficSplitFinalizerName indicates an expected call of TrafficSplitFinalizerName.
func (mr *MockTrafficSplitFinalizerMockRecorder) TrafficSplitFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficSplitFinalizerName", reflect.TypeOf((*MockTrafficSplitFinalizer)(nil).TrafficSplitFinalizerName))
}

// MockTrafficSplitReconcileLoop is a mock of TrafficSplitReconcileLoop interface.
type MockTrafficSplitReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitReconcileLoopMockRecorder
}

// MockTrafficSplitReconcileLoopMockRecorder is the mock recorder for MockTrafficSplitReconcileLoop.
type MockTrafficSplitReconcileLoopMockRecorder struct {
	mock *MockTrafficSplitReconcileLoop
}

// NewMockTrafficSplitReconcileLoop creates a new mock instance.
func NewMockTrafficSplitReconcileLoop(ctrl *gomock.Controller) *MockTrafficSplitReconcileLoop {
	mock := &MockTrafficSplitReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitReconcileLoop) EXPECT() *MockTrafficSplitReconcileLoopMockRecorder {
	return m.recorder
}

// RunTrafficSplitReconciler mocks base method.
func (m *MockTrafficSplitReconcileLoop) RunTrafficSplitReconciler(ctx context.Context, rec controller.TrafficSplitReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunTrafficSplitReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTrafficSplitReconciler indicates an expected call of RunTrafficSplitReconciler.
func (mr *MockTrafficSplitReconcileLoopMockRecorder) RunTrafficSplitReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTrafficSplitReconciler", reflect.TypeOf((*MockTrafficSplitReconcileLoop)(nil).RunTrafficSplitReconciler), varargs...)
}
