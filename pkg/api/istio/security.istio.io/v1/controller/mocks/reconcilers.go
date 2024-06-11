// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockAuthorizationPolicyReconciler is a mock of AuthorizationPolicyReconciler interface.
type MockAuthorizationPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationPolicyReconcilerMockRecorder
}

// MockAuthorizationPolicyReconcilerMockRecorder is the mock recorder for MockAuthorizationPolicyReconciler.
type MockAuthorizationPolicyReconcilerMockRecorder struct {
	mock *MockAuthorizationPolicyReconciler
}

// NewMockAuthorizationPolicyReconciler creates a new mock instance.
func NewMockAuthorizationPolicyReconciler(ctrl *gomock.Controller) *MockAuthorizationPolicyReconciler {
	mock := &MockAuthorizationPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockAuthorizationPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationPolicyReconciler) EXPECT() *MockAuthorizationPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAuthorizationPolicy mocks base method.
func (m *MockAuthorizationPolicyReconciler) ReconcileAuthorizationPolicy(obj *v1beta1.AuthorizationPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthorizationPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAuthorizationPolicy indicates an expected call of ReconcileAuthorizationPolicy.
func (mr *MockAuthorizationPolicyReconcilerMockRecorder) ReconcileAuthorizationPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthorizationPolicy", reflect.TypeOf((*MockAuthorizationPolicyReconciler)(nil).ReconcileAuthorizationPolicy), obj)
}

// MockAuthorizationPolicyDeletionReconciler is a mock of AuthorizationPolicyDeletionReconciler interface.
type MockAuthorizationPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationPolicyDeletionReconcilerMockRecorder
}

// MockAuthorizationPolicyDeletionReconcilerMockRecorder is the mock recorder for MockAuthorizationPolicyDeletionReconciler.
type MockAuthorizationPolicyDeletionReconcilerMockRecorder struct {
	mock *MockAuthorizationPolicyDeletionReconciler
}

// NewMockAuthorizationPolicyDeletionReconciler creates a new mock instance.
func NewMockAuthorizationPolicyDeletionReconciler(ctrl *gomock.Controller) *MockAuthorizationPolicyDeletionReconciler {
	mock := &MockAuthorizationPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockAuthorizationPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationPolicyDeletionReconciler) EXPECT() *MockAuthorizationPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAuthorizationPolicyDeletion mocks base method.
func (m *MockAuthorizationPolicyDeletionReconciler) ReconcileAuthorizationPolicyDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthorizationPolicyDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAuthorizationPolicyDeletion indicates an expected call of ReconcileAuthorizationPolicyDeletion.
func (mr *MockAuthorizationPolicyDeletionReconcilerMockRecorder) ReconcileAuthorizationPolicyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthorizationPolicyDeletion", reflect.TypeOf((*MockAuthorizationPolicyDeletionReconciler)(nil).ReconcileAuthorizationPolicyDeletion), req)
}

// MockAuthorizationPolicyFinalizer is a mock of AuthorizationPolicyFinalizer interface.
type MockAuthorizationPolicyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationPolicyFinalizerMockRecorder
}

// MockAuthorizationPolicyFinalizerMockRecorder is the mock recorder for MockAuthorizationPolicyFinalizer.
type MockAuthorizationPolicyFinalizerMockRecorder struct {
	mock *MockAuthorizationPolicyFinalizer
}

// NewMockAuthorizationPolicyFinalizer creates a new mock instance.
func NewMockAuthorizationPolicyFinalizer(ctrl *gomock.Controller) *MockAuthorizationPolicyFinalizer {
	mock := &MockAuthorizationPolicyFinalizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizationPolicyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationPolicyFinalizer) EXPECT() *MockAuthorizationPolicyFinalizerMockRecorder {
	return m.recorder
}

// AuthorizationPolicyFinalizerName mocks base method.
func (m *MockAuthorizationPolicyFinalizer) AuthorizationPolicyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationPolicyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthorizationPolicyFinalizerName indicates an expected call of AuthorizationPolicyFinalizerName.
func (mr *MockAuthorizationPolicyFinalizerMockRecorder) AuthorizationPolicyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationPolicyFinalizerName", reflect.TypeOf((*MockAuthorizationPolicyFinalizer)(nil).AuthorizationPolicyFinalizerName))
}

// FinalizeAuthorizationPolicy mocks base method.
func (m *MockAuthorizationPolicyFinalizer) FinalizeAuthorizationPolicy(obj *v1beta1.AuthorizationPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAuthorizationPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeAuthorizationPolicy indicates an expected call of FinalizeAuthorizationPolicy.
func (mr *MockAuthorizationPolicyFinalizerMockRecorder) FinalizeAuthorizationPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAuthorizationPolicy", reflect.TypeOf((*MockAuthorizationPolicyFinalizer)(nil).FinalizeAuthorizationPolicy), obj)
}

// ReconcileAuthorizationPolicy mocks base method.
func (m *MockAuthorizationPolicyFinalizer) ReconcileAuthorizationPolicy(obj *v1beta1.AuthorizationPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthorizationPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAuthorizationPolicy indicates an expected call of ReconcileAuthorizationPolicy.
func (mr *MockAuthorizationPolicyFinalizerMockRecorder) ReconcileAuthorizationPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthorizationPolicy", reflect.TypeOf((*MockAuthorizationPolicyFinalizer)(nil).ReconcileAuthorizationPolicy), obj)
}

// MockAuthorizationPolicyReconcileLoop is a mock of AuthorizationPolicyReconcileLoop interface.
type MockAuthorizationPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationPolicyReconcileLoopMockRecorder
}

// MockAuthorizationPolicyReconcileLoopMockRecorder is the mock recorder for MockAuthorizationPolicyReconcileLoop.
type MockAuthorizationPolicyReconcileLoopMockRecorder struct {
	mock *MockAuthorizationPolicyReconcileLoop
}

// NewMockAuthorizationPolicyReconcileLoop creates a new mock instance.
func NewMockAuthorizationPolicyReconcileLoop(ctrl *gomock.Controller) *MockAuthorizationPolicyReconcileLoop {
	mock := &MockAuthorizationPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockAuthorizationPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationPolicyReconcileLoop) EXPECT() *MockAuthorizationPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// RunAuthorizationPolicyReconciler mocks base method.
func (m *MockAuthorizationPolicyReconcileLoop) RunAuthorizationPolicyReconciler(ctx context.Context, rec controller.AuthorizationPolicyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunAuthorizationPolicyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunAuthorizationPolicyReconciler indicates an expected call of RunAuthorizationPolicyReconciler.
func (mr *MockAuthorizationPolicyReconcileLoopMockRecorder) RunAuthorizationPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAuthorizationPolicyReconciler", reflect.TypeOf((*MockAuthorizationPolicyReconcileLoop)(nil).RunAuthorizationPolicyReconciler), varargs...)
}

// MockPeerAuthenticationReconciler is a mock of PeerAuthenticationReconciler interface.
type MockPeerAuthenticationReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockPeerAuthenticationReconcilerMockRecorder
}

// MockPeerAuthenticationReconcilerMockRecorder is the mock recorder for MockPeerAuthenticationReconciler.
type MockPeerAuthenticationReconcilerMockRecorder struct {
	mock *MockPeerAuthenticationReconciler
}

// NewMockPeerAuthenticationReconciler creates a new mock instance.
func NewMockPeerAuthenticationReconciler(ctrl *gomock.Controller) *MockPeerAuthenticationReconciler {
	mock := &MockPeerAuthenticationReconciler{ctrl: ctrl}
	mock.recorder = &MockPeerAuthenticationReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerAuthenticationReconciler) EXPECT() *MockPeerAuthenticationReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePeerAuthentication mocks base method.
func (m *MockPeerAuthenticationReconciler) ReconcilePeerAuthentication(obj *v1beta1.PeerAuthentication) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePeerAuthentication", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePeerAuthentication indicates an expected call of ReconcilePeerAuthentication.
func (mr *MockPeerAuthenticationReconcilerMockRecorder) ReconcilePeerAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePeerAuthentication", reflect.TypeOf((*MockPeerAuthenticationReconciler)(nil).ReconcilePeerAuthentication), obj)
}

// MockPeerAuthenticationDeletionReconciler is a mock of PeerAuthenticationDeletionReconciler interface.
type MockPeerAuthenticationDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockPeerAuthenticationDeletionReconcilerMockRecorder
}

// MockPeerAuthenticationDeletionReconcilerMockRecorder is the mock recorder for MockPeerAuthenticationDeletionReconciler.
type MockPeerAuthenticationDeletionReconcilerMockRecorder struct {
	mock *MockPeerAuthenticationDeletionReconciler
}

// NewMockPeerAuthenticationDeletionReconciler creates a new mock instance.
func NewMockPeerAuthenticationDeletionReconciler(ctrl *gomock.Controller) *MockPeerAuthenticationDeletionReconciler {
	mock := &MockPeerAuthenticationDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockPeerAuthenticationDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerAuthenticationDeletionReconciler) EXPECT() *MockPeerAuthenticationDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePeerAuthenticationDeletion mocks base method.
func (m *MockPeerAuthenticationDeletionReconciler) ReconcilePeerAuthenticationDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePeerAuthenticationDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcilePeerAuthenticationDeletion indicates an expected call of ReconcilePeerAuthenticationDeletion.
func (mr *MockPeerAuthenticationDeletionReconcilerMockRecorder) ReconcilePeerAuthenticationDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePeerAuthenticationDeletion", reflect.TypeOf((*MockPeerAuthenticationDeletionReconciler)(nil).ReconcilePeerAuthenticationDeletion), req)
}

// MockPeerAuthenticationFinalizer is a mock of PeerAuthenticationFinalizer interface.
type MockPeerAuthenticationFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockPeerAuthenticationFinalizerMockRecorder
}

// MockPeerAuthenticationFinalizerMockRecorder is the mock recorder for MockPeerAuthenticationFinalizer.
type MockPeerAuthenticationFinalizerMockRecorder struct {
	mock *MockPeerAuthenticationFinalizer
}

// NewMockPeerAuthenticationFinalizer creates a new mock instance.
func NewMockPeerAuthenticationFinalizer(ctrl *gomock.Controller) *MockPeerAuthenticationFinalizer {
	mock := &MockPeerAuthenticationFinalizer{ctrl: ctrl}
	mock.recorder = &MockPeerAuthenticationFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerAuthenticationFinalizer) EXPECT() *MockPeerAuthenticationFinalizerMockRecorder {
	return m.recorder
}

// FinalizePeerAuthentication mocks base method.
func (m *MockPeerAuthenticationFinalizer) FinalizePeerAuthentication(obj *v1beta1.PeerAuthentication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizePeerAuthentication", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizePeerAuthentication indicates an expected call of FinalizePeerAuthentication.
func (mr *MockPeerAuthenticationFinalizerMockRecorder) FinalizePeerAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizePeerAuthentication", reflect.TypeOf((*MockPeerAuthenticationFinalizer)(nil).FinalizePeerAuthentication), obj)
}

// PeerAuthenticationFinalizerName mocks base method.
func (m *MockPeerAuthenticationFinalizer) PeerAuthenticationFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerAuthenticationFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// PeerAuthenticationFinalizerName indicates an expected call of PeerAuthenticationFinalizerName.
func (mr *MockPeerAuthenticationFinalizerMockRecorder) PeerAuthenticationFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerAuthenticationFinalizerName", reflect.TypeOf((*MockPeerAuthenticationFinalizer)(nil).PeerAuthenticationFinalizerName))
}

// ReconcilePeerAuthentication mocks base method.
func (m *MockPeerAuthenticationFinalizer) ReconcilePeerAuthentication(obj *v1beta1.PeerAuthentication) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePeerAuthentication", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePeerAuthentication indicates an expected call of ReconcilePeerAuthentication.
func (mr *MockPeerAuthenticationFinalizerMockRecorder) ReconcilePeerAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePeerAuthentication", reflect.TypeOf((*MockPeerAuthenticationFinalizer)(nil).ReconcilePeerAuthentication), obj)
}

// MockPeerAuthenticationReconcileLoop is a mock of PeerAuthenticationReconcileLoop interface.
type MockPeerAuthenticationReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockPeerAuthenticationReconcileLoopMockRecorder
}

// MockPeerAuthenticationReconcileLoopMockRecorder is the mock recorder for MockPeerAuthenticationReconcileLoop.
type MockPeerAuthenticationReconcileLoopMockRecorder struct {
	mock *MockPeerAuthenticationReconcileLoop
}

// NewMockPeerAuthenticationReconcileLoop creates a new mock instance.
func NewMockPeerAuthenticationReconcileLoop(ctrl *gomock.Controller) *MockPeerAuthenticationReconcileLoop {
	mock := &MockPeerAuthenticationReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockPeerAuthenticationReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerAuthenticationReconcileLoop) EXPECT() *MockPeerAuthenticationReconcileLoopMockRecorder {
	return m.recorder
}

// RunPeerAuthenticationReconciler mocks base method.
func (m *MockPeerAuthenticationReconcileLoop) RunPeerAuthenticationReconciler(ctx context.Context, rec controller.PeerAuthenticationReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunPeerAuthenticationReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunPeerAuthenticationReconciler indicates an expected call of RunPeerAuthenticationReconciler.
func (mr *MockPeerAuthenticationReconcileLoopMockRecorder) RunPeerAuthenticationReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPeerAuthenticationReconciler", reflect.TypeOf((*MockPeerAuthenticationReconcileLoop)(nil).RunPeerAuthenticationReconciler), varargs...)
}

// MockRequestAuthenticationReconciler is a mock of RequestAuthenticationReconciler interface.
type MockRequestAuthenticationReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockRequestAuthenticationReconcilerMockRecorder
}

// MockRequestAuthenticationReconcilerMockRecorder is the mock recorder for MockRequestAuthenticationReconciler.
type MockRequestAuthenticationReconcilerMockRecorder struct {
	mock *MockRequestAuthenticationReconciler
}

// NewMockRequestAuthenticationReconciler creates a new mock instance.
func NewMockRequestAuthenticationReconciler(ctrl *gomock.Controller) *MockRequestAuthenticationReconciler {
	mock := &MockRequestAuthenticationReconciler{ctrl: ctrl}
	mock.recorder = &MockRequestAuthenticationReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestAuthenticationReconciler) EXPECT() *MockRequestAuthenticationReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRequestAuthentication mocks base method.
func (m *MockRequestAuthenticationReconciler) ReconcileRequestAuthentication(obj *v1beta1.RequestAuthentication) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRequestAuthentication", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRequestAuthentication indicates an expected call of ReconcileRequestAuthentication.
func (mr *MockRequestAuthenticationReconcilerMockRecorder) ReconcileRequestAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRequestAuthentication", reflect.TypeOf((*MockRequestAuthenticationReconciler)(nil).ReconcileRequestAuthentication), obj)
}

// MockRequestAuthenticationDeletionReconciler is a mock of RequestAuthenticationDeletionReconciler interface.
type MockRequestAuthenticationDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockRequestAuthenticationDeletionReconcilerMockRecorder
}

// MockRequestAuthenticationDeletionReconcilerMockRecorder is the mock recorder for MockRequestAuthenticationDeletionReconciler.
type MockRequestAuthenticationDeletionReconcilerMockRecorder struct {
	mock *MockRequestAuthenticationDeletionReconciler
}

// NewMockRequestAuthenticationDeletionReconciler creates a new mock instance.
func NewMockRequestAuthenticationDeletionReconciler(ctrl *gomock.Controller) *MockRequestAuthenticationDeletionReconciler {
	mock := &MockRequestAuthenticationDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockRequestAuthenticationDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestAuthenticationDeletionReconciler) EXPECT() *MockRequestAuthenticationDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRequestAuthenticationDeletion mocks base method.
func (m *MockRequestAuthenticationDeletionReconciler) ReconcileRequestAuthenticationDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRequestAuthenticationDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRequestAuthenticationDeletion indicates an expected call of ReconcileRequestAuthenticationDeletion.
func (mr *MockRequestAuthenticationDeletionReconcilerMockRecorder) ReconcileRequestAuthenticationDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRequestAuthenticationDeletion", reflect.TypeOf((*MockRequestAuthenticationDeletionReconciler)(nil).ReconcileRequestAuthenticationDeletion), req)
}

// MockRequestAuthenticationFinalizer is a mock of RequestAuthenticationFinalizer interface.
type MockRequestAuthenticationFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockRequestAuthenticationFinalizerMockRecorder
}

// MockRequestAuthenticationFinalizerMockRecorder is the mock recorder for MockRequestAuthenticationFinalizer.
type MockRequestAuthenticationFinalizerMockRecorder struct {
	mock *MockRequestAuthenticationFinalizer
}

// NewMockRequestAuthenticationFinalizer creates a new mock instance.
func NewMockRequestAuthenticationFinalizer(ctrl *gomock.Controller) *MockRequestAuthenticationFinalizer {
	mock := &MockRequestAuthenticationFinalizer{ctrl: ctrl}
	mock.recorder = &MockRequestAuthenticationFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestAuthenticationFinalizer) EXPECT() *MockRequestAuthenticationFinalizerMockRecorder {
	return m.recorder
}

// FinalizeRequestAuthentication mocks base method.
func (m *MockRequestAuthenticationFinalizer) FinalizeRequestAuthentication(obj *v1beta1.RequestAuthentication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeRequestAuthentication", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeRequestAuthentication indicates an expected call of FinalizeRequestAuthentication.
func (mr *MockRequestAuthenticationFinalizerMockRecorder) FinalizeRequestAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeRequestAuthentication", reflect.TypeOf((*MockRequestAuthenticationFinalizer)(nil).FinalizeRequestAuthentication), obj)
}

// ReconcileRequestAuthentication mocks base method.
func (m *MockRequestAuthenticationFinalizer) ReconcileRequestAuthentication(obj *v1beta1.RequestAuthentication) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRequestAuthentication", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRequestAuthentication indicates an expected call of ReconcileRequestAuthentication.
func (mr *MockRequestAuthenticationFinalizerMockRecorder) ReconcileRequestAuthentication(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRequestAuthentication", reflect.TypeOf((*MockRequestAuthenticationFinalizer)(nil).ReconcileRequestAuthentication), obj)
}

// RequestAuthenticationFinalizerName mocks base method.
func (m *MockRequestAuthenticationFinalizer) RequestAuthenticationFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestAuthenticationFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// RequestAuthenticationFinalizerName indicates an expected call of RequestAuthenticationFinalizerName.
func (mr *MockRequestAuthenticationFinalizerMockRecorder) RequestAuthenticationFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAuthenticationFinalizerName", reflect.TypeOf((*MockRequestAuthenticationFinalizer)(nil).RequestAuthenticationFinalizerName))
}

// MockRequestAuthenticationReconcileLoop is a mock of RequestAuthenticationReconcileLoop interface.
type MockRequestAuthenticationReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockRequestAuthenticationReconcileLoopMockRecorder
}

// MockRequestAuthenticationReconcileLoopMockRecorder is the mock recorder for MockRequestAuthenticationReconcileLoop.
type MockRequestAuthenticationReconcileLoopMockRecorder struct {
	mock *MockRequestAuthenticationReconcileLoop
}

// NewMockRequestAuthenticationReconcileLoop creates a new mock instance.
func NewMockRequestAuthenticationReconcileLoop(ctrl *gomock.Controller) *MockRequestAuthenticationReconcileLoop {
	mock := &MockRequestAuthenticationReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockRequestAuthenticationReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestAuthenticationReconcileLoop) EXPECT() *MockRequestAuthenticationReconcileLoopMockRecorder {
	return m.recorder
}

// RunRequestAuthenticationReconciler mocks base method.
func (m *MockRequestAuthenticationReconcileLoop) RunRequestAuthenticationReconciler(ctx context.Context, rec controller.RequestAuthenticationReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunRequestAuthenticationReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunRequestAuthenticationReconciler indicates an expected call of RunRequestAuthenticationReconciler.
func (mr *MockRequestAuthenticationReconcileLoopMockRecorder) RunRequestAuthenticationReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunRequestAuthenticationReconciler", reflect.TypeOf((*MockRequestAuthenticationReconcileLoop)(nil).RunRequestAuthenticationReconciler), varargs...)
}
