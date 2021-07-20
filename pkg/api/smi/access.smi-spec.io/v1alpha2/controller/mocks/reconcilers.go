// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	controller "github.com/solo-io/external-apis/pkg/api/smi/access.smi-spec.io/v1alpha2/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockTrafficTargetReconciler is a mock of TrafficTargetReconciler interface.
type MockTrafficTargetReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficTargetReconcilerMockRecorder
}

// MockTrafficTargetReconcilerMockRecorder is the mock recorder for MockTrafficTargetReconciler.
type MockTrafficTargetReconcilerMockRecorder struct {
	mock *MockTrafficTargetReconciler
}

// NewMockTrafficTargetReconciler creates a new mock instance.
func NewMockTrafficTargetReconciler(ctrl *gomock.Controller) *MockTrafficTargetReconciler {
	mock := &MockTrafficTargetReconciler{ctrl: ctrl}
	mock.recorder = &MockTrafficTargetReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficTargetReconciler) EXPECT() *MockTrafficTargetReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficTarget mocks base method.
func (m *MockTrafficTargetReconciler) ReconcileTrafficTarget(obj *v1alpha2.TrafficTarget) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficTarget", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficTarget indicates an expected call of ReconcileTrafficTarget.
func (mr *MockTrafficTargetReconcilerMockRecorder) ReconcileTrafficTarget(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficTarget", reflect.TypeOf((*MockTrafficTargetReconciler)(nil).ReconcileTrafficTarget), obj)
}

// MockTrafficTargetDeletionReconciler is a mock of TrafficTargetDeletionReconciler interface.
type MockTrafficTargetDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficTargetDeletionReconcilerMockRecorder
}

// MockTrafficTargetDeletionReconcilerMockRecorder is the mock recorder for MockTrafficTargetDeletionReconciler.
type MockTrafficTargetDeletionReconcilerMockRecorder struct {
	mock *MockTrafficTargetDeletionReconciler
}

// NewMockTrafficTargetDeletionReconciler creates a new mock instance.
func NewMockTrafficTargetDeletionReconciler(ctrl *gomock.Controller) *MockTrafficTargetDeletionReconciler {
	mock := &MockTrafficTargetDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockTrafficTargetDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficTargetDeletionReconciler) EXPECT() *MockTrafficTargetDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficTargetDeletion mocks base method.
func (m *MockTrafficTargetDeletionReconciler) ReconcileTrafficTargetDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficTargetDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileTrafficTargetDeletion indicates an expected call of ReconcileTrafficTargetDeletion.
func (mr *MockTrafficTargetDeletionReconcilerMockRecorder) ReconcileTrafficTargetDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficTargetDeletion", reflect.TypeOf((*MockTrafficTargetDeletionReconciler)(nil).ReconcileTrafficTargetDeletion), req)
}

// MockTrafficTargetFinalizer is a mock of TrafficTargetFinalizer interface.
type MockTrafficTargetFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficTargetFinalizerMockRecorder
}

// MockTrafficTargetFinalizerMockRecorder is the mock recorder for MockTrafficTargetFinalizer.
type MockTrafficTargetFinalizerMockRecorder struct {
	mock *MockTrafficTargetFinalizer
}

// NewMockTrafficTargetFinalizer creates a new mock instance.
func NewMockTrafficTargetFinalizer(ctrl *gomock.Controller) *MockTrafficTargetFinalizer {
	mock := &MockTrafficTargetFinalizer{ctrl: ctrl}
	mock.recorder = &MockTrafficTargetFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficTargetFinalizer) EXPECT() *MockTrafficTargetFinalizerMockRecorder {
	return m.recorder
}

// FinalizeTrafficTarget mocks base method.
func (m *MockTrafficTargetFinalizer) FinalizeTrafficTarget(obj *v1alpha2.TrafficTarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeTrafficTarget", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeTrafficTarget indicates an expected call of FinalizeTrafficTarget.
func (mr *MockTrafficTargetFinalizerMockRecorder) FinalizeTrafficTarget(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeTrafficTarget", reflect.TypeOf((*MockTrafficTargetFinalizer)(nil).FinalizeTrafficTarget), obj)
}

// ReconcileTrafficTarget mocks base method.
func (m *MockTrafficTargetFinalizer) ReconcileTrafficTarget(obj *v1alpha2.TrafficTarget) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficTarget", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficTarget indicates an expected call of ReconcileTrafficTarget.
func (mr *MockTrafficTargetFinalizerMockRecorder) ReconcileTrafficTarget(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficTarget", reflect.TypeOf((*MockTrafficTargetFinalizer)(nil).ReconcileTrafficTarget), obj)
}

// TrafficTargetFinalizerName mocks base method.
func (m *MockTrafficTargetFinalizer) TrafficTargetFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficTargetFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TrafficTargetFinalizerName indicates an expected call of TrafficTargetFinalizerName.
func (mr *MockTrafficTargetFinalizerMockRecorder) TrafficTargetFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficTargetFinalizerName", reflect.TypeOf((*MockTrafficTargetFinalizer)(nil).TrafficTargetFinalizerName))
}

// MockTrafficTargetReconcileLoop is a mock of TrafficTargetReconcileLoop interface.
type MockTrafficTargetReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficTargetReconcileLoopMockRecorder
}

// MockTrafficTargetReconcileLoopMockRecorder is the mock recorder for MockTrafficTargetReconcileLoop.
type MockTrafficTargetReconcileLoopMockRecorder struct {
	mock *MockTrafficTargetReconcileLoop
}

// NewMockTrafficTargetReconcileLoop creates a new mock instance.
func NewMockTrafficTargetReconcileLoop(ctrl *gomock.Controller) *MockTrafficTargetReconcileLoop {
	mock := &MockTrafficTargetReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockTrafficTargetReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficTargetReconcileLoop) EXPECT() *MockTrafficTargetReconcileLoopMockRecorder {
	return m.recorder
}

// RunTrafficTargetReconciler mocks base method.
func (m *MockTrafficTargetReconcileLoop) RunTrafficTargetReconciler(ctx context.Context, rec controller.TrafficTargetReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunTrafficTargetReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTrafficTargetReconciler indicates an expected call of RunTrafficTargetReconciler.
func (mr *MockTrafficTargetReconcileLoopMockRecorder) RunTrafficTargetReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTrafficTargetReconciler", reflect.TypeOf((*MockTrafficTargetReconcileLoop)(nil).RunTrafficTargetReconciler), varargs...)
}
