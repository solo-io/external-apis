// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/certificates.k8s.io/v1beta1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1beta1 "k8s.io/api/certificates/v1beta1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterCertificateSigningRequestReconciler is a mock of MulticlusterCertificateSigningRequestReconciler interface
type MockMulticlusterCertificateSigningRequestReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCertificateSigningRequestReconcilerMockRecorder
}

// MockMulticlusterCertificateSigningRequestReconcilerMockRecorder is the mock recorder for MockMulticlusterCertificateSigningRequestReconciler
type MockMulticlusterCertificateSigningRequestReconcilerMockRecorder struct {
	mock *MockMulticlusterCertificateSigningRequestReconciler
}

// NewMockMulticlusterCertificateSigningRequestReconciler creates a new mock instance
func NewMockMulticlusterCertificateSigningRequestReconciler(ctrl *gomock.Controller) *MockMulticlusterCertificateSigningRequestReconciler {
	mock := &MockMulticlusterCertificateSigningRequestReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCertificateSigningRequestReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCertificateSigningRequestReconciler) EXPECT() *MockMulticlusterCertificateSigningRequestReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCertificateSigningRequest mocks base method
func (m *MockMulticlusterCertificateSigningRequestReconciler) ReconcileCertificateSigningRequest(clusterName string, obj *v1beta1.CertificateSigningRequest) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCertificateSigningRequest", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCertificateSigningRequest indicates an expected call of ReconcileCertificateSigningRequest
func (mr *MockMulticlusterCertificateSigningRequestReconcilerMockRecorder) ReconcileCertificateSigningRequest(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCertificateSigningRequest", reflect.TypeOf((*MockMulticlusterCertificateSigningRequestReconciler)(nil).ReconcileCertificateSigningRequest), clusterName, obj)
}

// MockMulticlusterCertificateSigningRequestDeletionReconciler is a mock of MulticlusterCertificateSigningRequestDeletionReconciler interface
type MockMulticlusterCertificateSigningRequestDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCertificateSigningRequestDeletionReconcilerMockRecorder
}

// MockMulticlusterCertificateSigningRequestDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterCertificateSigningRequestDeletionReconciler
type MockMulticlusterCertificateSigningRequestDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterCertificateSigningRequestDeletionReconciler
}

// NewMockMulticlusterCertificateSigningRequestDeletionReconciler creates a new mock instance
func NewMockMulticlusterCertificateSigningRequestDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterCertificateSigningRequestDeletionReconciler {
	mock := &MockMulticlusterCertificateSigningRequestDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCertificateSigningRequestDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCertificateSigningRequestDeletionReconciler) EXPECT() *MockMulticlusterCertificateSigningRequestDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCertificateSigningRequestDeletion mocks base method
func (m *MockMulticlusterCertificateSigningRequestDeletionReconciler) ReconcileCertificateSigningRequestDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCertificateSigningRequestDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileCertificateSigningRequestDeletion indicates an expected call of ReconcileCertificateSigningRequestDeletion
func (mr *MockMulticlusterCertificateSigningRequestDeletionReconcilerMockRecorder) ReconcileCertificateSigningRequestDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCertificateSigningRequestDeletion", reflect.TypeOf((*MockMulticlusterCertificateSigningRequestDeletionReconciler)(nil).ReconcileCertificateSigningRequestDeletion), clusterName, req)
}

// MockMulticlusterCertificateSigningRequestReconcileLoop is a mock of MulticlusterCertificateSigningRequestReconcileLoop interface
type MockMulticlusterCertificateSigningRequestReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCertificateSigningRequestReconcileLoopMockRecorder
}

// MockMulticlusterCertificateSigningRequestReconcileLoopMockRecorder is the mock recorder for MockMulticlusterCertificateSigningRequestReconcileLoop
type MockMulticlusterCertificateSigningRequestReconcileLoopMockRecorder struct {
	mock *MockMulticlusterCertificateSigningRequestReconcileLoop
}

// NewMockMulticlusterCertificateSigningRequestReconcileLoop creates a new mock instance
func NewMockMulticlusterCertificateSigningRequestReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterCertificateSigningRequestReconcileLoop {
	mock := &MockMulticlusterCertificateSigningRequestReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCertificateSigningRequestReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCertificateSigningRequestReconcileLoop) EXPECT() *MockMulticlusterCertificateSigningRequestReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterCertificateSigningRequestReconciler mocks base method
func (m *MockMulticlusterCertificateSigningRequestReconcileLoop) AddMulticlusterCertificateSigningRequestReconciler(ctx context.Context, rec controller.MulticlusterCertificateSigningRequestReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterCertificateSigningRequestReconciler", varargs...)
}

// AddMulticlusterCertificateSigningRequestReconciler indicates an expected call of AddMulticlusterCertificateSigningRequestReconciler
func (mr *MockMulticlusterCertificateSigningRequestReconcileLoopMockRecorder) AddMulticlusterCertificateSigningRequestReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterCertificateSigningRequestReconciler", reflect.TypeOf((*MockMulticlusterCertificateSigningRequestReconcileLoop)(nil).AddMulticlusterCertificateSigningRequestReconciler), varargs...)
}
