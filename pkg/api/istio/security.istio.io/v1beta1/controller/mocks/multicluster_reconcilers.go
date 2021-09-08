// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterAuthorizationPolicyReconciler is a mock of MulticlusterAuthorizationPolicyReconciler interface
type MockMulticlusterAuthorizationPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAuthorizationPolicyReconcilerMockRecorder
}

// MockMulticlusterAuthorizationPolicyReconcilerMockRecorder is the mock recorder for MockMulticlusterAuthorizationPolicyReconciler
type MockMulticlusterAuthorizationPolicyReconcilerMockRecorder struct {
	mock *MockMulticlusterAuthorizationPolicyReconciler
}

// NewMockMulticlusterAuthorizationPolicyReconciler creates a new mock instance
func NewMockMulticlusterAuthorizationPolicyReconciler(ctrl *gomock.Controller) *MockMulticlusterAuthorizationPolicyReconciler {
	mock := &MockMulticlusterAuthorizationPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAuthorizationPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterAuthorizationPolicyReconciler) EXPECT() *MockMulticlusterAuthorizationPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAuthorizationPolicy mocks base method
func (m *MockMulticlusterAuthorizationPolicyReconciler) ReconcileAuthorizationPolicy(clusterName string, obj *v1beta1.AuthorizationPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthorizationPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAuthorizationPolicy indicates an expected call of ReconcileAuthorizationPolicy
func (mr *MockMulticlusterAuthorizationPolicyReconcilerMockRecorder) ReconcileAuthorizationPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthorizationPolicy", reflect.TypeOf((*MockMulticlusterAuthorizationPolicyReconciler)(nil).ReconcileAuthorizationPolicy), clusterName, obj)
}

// MockMulticlusterAuthorizationPolicyDeletionReconciler is a mock of MulticlusterAuthorizationPolicyDeletionReconciler interface
type MockMulticlusterAuthorizationPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAuthorizationPolicyDeletionReconcilerMockRecorder
}

// MockMulticlusterAuthorizationPolicyDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterAuthorizationPolicyDeletionReconciler
type MockMulticlusterAuthorizationPolicyDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterAuthorizationPolicyDeletionReconciler
}

// NewMockMulticlusterAuthorizationPolicyDeletionReconciler creates a new mock instance
func NewMockMulticlusterAuthorizationPolicyDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterAuthorizationPolicyDeletionReconciler {
	mock := &MockMulticlusterAuthorizationPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAuthorizationPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterAuthorizationPolicyDeletionReconciler) EXPECT() *MockMulticlusterAuthorizationPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAuthorizationPolicyDeletion mocks base method
func (m *MockMulticlusterAuthorizationPolicyDeletionReconciler) ReconcileAuthorizationPolicyDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthorizationPolicyDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAuthorizationPolicyDeletion indicates an expected call of ReconcileAuthorizationPolicyDeletion
func (mr *MockMulticlusterAuthorizationPolicyDeletionReconcilerMockRecorder) ReconcileAuthorizationPolicyDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthorizationPolicyDeletion", reflect.TypeOf((*MockMulticlusterAuthorizationPolicyDeletionReconciler)(nil).ReconcileAuthorizationPolicyDeletion), clusterName, req)
}

// MockMulticlusterAuthorizationPolicyReconcileLoop is a mock of MulticlusterAuthorizationPolicyReconcileLoop interface
type MockMulticlusterAuthorizationPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterAuthorizationPolicyReconcileLoopMockRecorder
}

// MockMulticlusterAuthorizationPolicyReconcileLoopMockRecorder is the mock recorder for MockMulticlusterAuthorizationPolicyReconcileLoop
type MockMulticlusterAuthorizationPolicyReconcileLoopMockRecorder struct {
	mock *MockMulticlusterAuthorizationPolicyReconcileLoop
}

// NewMockMulticlusterAuthorizationPolicyReconcileLoop creates a new mock instance
func NewMockMulticlusterAuthorizationPolicyReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterAuthorizationPolicyReconcileLoop {
	mock := &MockMulticlusterAuthorizationPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterAuthorizationPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterAuthorizationPolicyReconcileLoop) EXPECT() *MockMulticlusterAuthorizationPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterAuthorizationPolicyReconciler mocks base method
func (m *MockMulticlusterAuthorizationPolicyReconcileLoop) AddMulticlusterAuthorizationPolicyReconciler(ctx context.Context, rec controller.MulticlusterAuthorizationPolicyReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterAuthorizationPolicyReconciler", varargs...)
}

// AddMulticlusterAuthorizationPolicyReconciler indicates an expected call of AddMulticlusterAuthorizationPolicyReconciler
func (mr *MockMulticlusterAuthorizationPolicyReconcileLoopMockRecorder) AddMulticlusterAuthorizationPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterAuthorizationPolicyReconciler", reflect.TypeOf((*MockMulticlusterAuthorizationPolicyReconcileLoop)(nil).AddMulticlusterAuthorizationPolicyReconciler), varargs...)
}

// MockMulticlusterPeerAuthenticationReconciler is a mock of MulticlusterPeerAuthenticationReconciler interface
type MockMulticlusterPeerAuthenticationReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPeerAuthenticationReconcilerMockRecorder
}

// MockMulticlusterPeerAuthenticationReconcilerMockRecorder is the mock recorder for MockMulticlusterPeerAuthenticationReconciler
type MockMulticlusterPeerAuthenticationReconcilerMockRecorder struct {
	mock *MockMulticlusterPeerAuthenticationReconciler
}

// NewMockMulticlusterPeerAuthenticationReconciler creates a new mock instance
func NewMockMulticlusterPeerAuthenticationReconciler(ctrl *gomock.Controller) *MockMulticlusterPeerAuthenticationReconciler {
	mock := &MockMulticlusterPeerAuthenticationReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPeerAuthenticationReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPeerAuthenticationReconciler) EXPECT() *MockMulticlusterPeerAuthenticationReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePeerAuthentication mocks base method
func (m *MockMulticlusterPeerAuthenticationReconciler) ReconcilePeerAuthentication(clusterName string, obj *v1beta1.PeerAuthentication) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePeerAuthentication", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePeerAuthentication indicates an expected call of ReconcilePeerAuthentication
func (mr *MockMulticlusterPeerAuthenticationReconcilerMockRecorder) ReconcilePeerAuthentication(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePeerAuthentication", reflect.TypeOf((*MockMulticlusterPeerAuthenticationReconciler)(nil).ReconcilePeerAuthentication), clusterName, obj)
}

// MockMulticlusterPeerAuthenticationDeletionReconciler is a mock of MulticlusterPeerAuthenticationDeletionReconciler interface
type MockMulticlusterPeerAuthenticationDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPeerAuthenticationDeletionReconcilerMockRecorder
}

// MockMulticlusterPeerAuthenticationDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterPeerAuthenticationDeletionReconciler
type MockMulticlusterPeerAuthenticationDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterPeerAuthenticationDeletionReconciler
}

// NewMockMulticlusterPeerAuthenticationDeletionReconciler creates a new mock instance
func NewMockMulticlusterPeerAuthenticationDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterPeerAuthenticationDeletionReconciler {
	mock := &MockMulticlusterPeerAuthenticationDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPeerAuthenticationDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPeerAuthenticationDeletionReconciler) EXPECT() *MockMulticlusterPeerAuthenticationDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePeerAuthenticationDeletion mocks base method
func (m *MockMulticlusterPeerAuthenticationDeletionReconciler) ReconcilePeerAuthenticationDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePeerAuthenticationDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcilePeerAuthenticationDeletion indicates an expected call of ReconcilePeerAuthenticationDeletion
func (mr *MockMulticlusterPeerAuthenticationDeletionReconcilerMockRecorder) ReconcilePeerAuthenticationDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePeerAuthenticationDeletion", reflect.TypeOf((*MockMulticlusterPeerAuthenticationDeletionReconciler)(nil).ReconcilePeerAuthenticationDeletion), clusterName, req)
}

// MockMulticlusterPeerAuthenticationReconcileLoop is a mock of MulticlusterPeerAuthenticationReconcileLoop interface
type MockMulticlusterPeerAuthenticationReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPeerAuthenticationReconcileLoopMockRecorder
}

// MockMulticlusterPeerAuthenticationReconcileLoopMockRecorder is the mock recorder for MockMulticlusterPeerAuthenticationReconcileLoop
type MockMulticlusterPeerAuthenticationReconcileLoopMockRecorder struct {
	mock *MockMulticlusterPeerAuthenticationReconcileLoop
}

// NewMockMulticlusterPeerAuthenticationReconcileLoop creates a new mock instance
func NewMockMulticlusterPeerAuthenticationReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterPeerAuthenticationReconcileLoop {
	mock := &MockMulticlusterPeerAuthenticationReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPeerAuthenticationReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterPeerAuthenticationReconcileLoop) EXPECT() *MockMulticlusterPeerAuthenticationReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterPeerAuthenticationReconciler mocks base method
func (m *MockMulticlusterPeerAuthenticationReconcileLoop) AddMulticlusterPeerAuthenticationReconciler(ctx context.Context, rec controller.MulticlusterPeerAuthenticationReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterPeerAuthenticationReconciler", varargs...)
}

// AddMulticlusterPeerAuthenticationReconciler indicates an expected call of AddMulticlusterPeerAuthenticationReconciler
func (mr *MockMulticlusterPeerAuthenticationReconcileLoopMockRecorder) AddMulticlusterPeerAuthenticationReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterPeerAuthenticationReconciler", reflect.TypeOf((*MockMulticlusterPeerAuthenticationReconcileLoop)(nil).AddMulticlusterPeerAuthenticationReconciler), varargs...)
}
