// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/admissionregistration.k8s.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "k8s.io/api/admissionregistration/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockValidatingWebhookConfigurationReconciler is a mock of ValidatingWebhookConfigurationReconciler interface.
type MockValidatingWebhookConfigurationReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationReconcilerMockRecorder
}

// MockValidatingWebhookConfigurationReconcilerMockRecorder is the mock recorder for MockValidatingWebhookConfigurationReconciler.
type MockValidatingWebhookConfigurationReconcilerMockRecorder struct {
	mock *MockValidatingWebhookConfigurationReconciler
}

// NewMockValidatingWebhookConfigurationReconciler creates a new mock instance.
func NewMockValidatingWebhookConfigurationReconciler(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationReconciler {
	mock := &MockValidatingWebhookConfigurationReconciler{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationReconciler) EXPECT() *MockValidatingWebhookConfigurationReconcilerMockRecorder {
	return m.recorder
}

// ReconcileValidatingWebhookConfiguration mocks base method.
func (m *MockValidatingWebhookConfigurationReconciler) ReconcileValidatingWebhookConfiguration(obj *v1.ValidatingWebhookConfiguration) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileValidatingWebhookConfiguration", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileValidatingWebhookConfiguration indicates an expected call of ReconcileValidatingWebhookConfiguration.
func (mr *MockValidatingWebhookConfigurationReconcilerMockRecorder) ReconcileValidatingWebhookConfiguration(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileValidatingWebhookConfiguration", reflect.TypeOf((*MockValidatingWebhookConfigurationReconciler)(nil).ReconcileValidatingWebhookConfiguration), obj)
}

// MockValidatingWebhookConfigurationDeletionReconciler is a mock of ValidatingWebhookConfigurationDeletionReconciler interface.
type MockValidatingWebhookConfigurationDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationDeletionReconcilerMockRecorder
}

// MockValidatingWebhookConfigurationDeletionReconcilerMockRecorder is the mock recorder for MockValidatingWebhookConfigurationDeletionReconciler.
type MockValidatingWebhookConfigurationDeletionReconcilerMockRecorder struct {
	mock *MockValidatingWebhookConfigurationDeletionReconciler
}

// NewMockValidatingWebhookConfigurationDeletionReconciler creates a new mock instance.
func NewMockValidatingWebhookConfigurationDeletionReconciler(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationDeletionReconciler {
	mock := &MockValidatingWebhookConfigurationDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationDeletionReconciler) EXPECT() *MockValidatingWebhookConfigurationDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileValidatingWebhookConfigurationDeletion mocks base method.
func (m *MockValidatingWebhookConfigurationDeletionReconciler) ReconcileValidatingWebhookConfigurationDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileValidatingWebhookConfigurationDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileValidatingWebhookConfigurationDeletion indicates an expected call of ReconcileValidatingWebhookConfigurationDeletion.
func (mr *MockValidatingWebhookConfigurationDeletionReconcilerMockRecorder) ReconcileValidatingWebhookConfigurationDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileValidatingWebhookConfigurationDeletion", reflect.TypeOf((*MockValidatingWebhookConfigurationDeletionReconciler)(nil).ReconcileValidatingWebhookConfigurationDeletion), req)
}

// MockValidatingWebhookConfigurationFinalizer is a mock of ValidatingWebhookConfigurationFinalizer interface.
type MockValidatingWebhookConfigurationFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationFinalizerMockRecorder
}

// MockValidatingWebhookConfigurationFinalizerMockRecorder is the mock recorder for MockValidatingWebhookConfigurationFinalizer.
type MockValidatingWebhookConfigurationFinalizerMockRecorder struct {
	mock *MockValidatingWebhookConfigurationFinalizer
}

// NewMockValidatingWebhookConfigurationFinalizer creates a new mock instance.
func NewMockValidatingWebhookConfigurationFinalizer(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationFinalizer {
	mock := &MockValidatingWebhookConfigurationFinalizer{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationFinalizer) EXPECT() *MockValidatingWebhookConfigurationFinalizerMockRecorder {
	return m.recorder
}

// FinalizeValidatingWebhookConfiguration mocks base method.
func (m *MockValidatingWebhookConfigurationFinalizer) FinalizeValidatingWebhookConfiguration(obj *v1.ValidatingWebhookConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeValidatingWebhookConfiguration", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeValidatingWebhookConfiguration indicates an expected call of FinalizeValidatingWebhookConfiguration.
func (mr *MockValidatingWebhookConfigurationFinalizerMockRecorder) FinalizeValidatingWebhookConfiguration(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeValidatingWebhookConfiguration", reflect.TypeOf((*MockValidatingWebhookConfigurationFinalizer)(nil).FinalizeValidatingWebhookConfiguration), obj)
}

// ReconcileValidatingWebhookConfiguration mocks base method.
func (m *MockValidatingWebhookConfigurationFinalizer) ReconcileValidatingWebhookConfiguration(obj *v1.ValidatingWebhookConfiguration) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileValidatingWebhookConfiguration", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileValidatingWebhookConfiguration indicates an expected call of ReconcileValidatingWebhookConfiguration.
func (mr *MockValidatingWebhookConfigurationFinalizerMockRecorder) ReconcileValidatingWebhookConfiguration(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileValidatingWebhookConfiguration", reflect.TypeOf((*MockValidatingWebhookConfigurationFinalizer)(nil).ReconcileValidatingWebhookConfiguration), obj)
}

// ValidatingWebhookConfigurationFinalizerName mocks base method.
func (m *MockValidatingWebhookConfigurationFinalizer) ValidatingWebhookConfigurationFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatingWebhookConfigurationFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ValidatingWebhookConfigurationFinalizerName indicates an expected call of ValidatingWebhookConfigurationFinalizerName.
func (mr *MockValidatingWebhookConfigurationFinalizerMockRecorder) ValidatingWebhookConfigurationFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatingWebhookConfigurationFinalizerName", reflect.TypeOf((*MockValidatingWebhookConfigurationFinalizer)(nil).ValidatingWebhookConfigurationFinalizerName))
}

// MockValidatingWebhookConfigurationReconcileLoop is a mock of ValidatingWebhookConfigurationReconcileLoop interface.
type MockValidatingWebhookConfigurationReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationReconcileLoopMockRecorder
}

// MockValidatingWebhookConfigurationReconcileLoopMockRecorder is the mock recorder for MockValidatingWebhookConfigurationReconcileLoop.
type MockValidatingWebhookConfigurationReconcileLoopMockRecorder struct {
	mock *MockValidatingWebhookConfigurationReconcileLoop
}

// NewMockValidatingWebhookConfigurationReconcileLoop creates a new mock instance.
func NewMockValidatingWebhookConfigurationReconcileLoop(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationReconcileLoop {
	mock := &MockValidatingWebhookConfigurationReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationReconcileLoop) EXPECT() *MockValidatingWebhookConfigurationReconcileLoopMockRecorder {
	return m.recorder
}

// RunValidatingWebhookConfigurationReconciler mocks base method.
func (m *MockValidatingWebhookConfigurationReconcileLoop) RunValidatingWebhookConfigurationReconciler(ctx context.Context, rec controller.ValidatingWebhookConfigurationReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunValidatingWebhookConfigurationReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunValidatingWebhookConfigurationReconciler indicates an expected call of RunValidatingWebhookConfigurationReconciler.
func (mr *MockValidatingWebhookConfigurationReconcileLoopMockRecorder) RunValidatingWebhookConfigurationReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunValidatingWebhookConfigurationReconciler", reflect.TypeOf((*MockValidatingWebhookConfigurationReconcileLoop)(nil).RunValidatingWebhookConfigurationReconciler), varargs...)
}
