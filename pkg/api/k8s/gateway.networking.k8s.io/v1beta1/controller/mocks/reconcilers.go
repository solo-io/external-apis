// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/gateway.networking.k8s.io/v1beta1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// MockGatewayReconciler is a mock of GatewayReconciler interface.
type MockGatewayReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayReconcilerMockRecorder
}

// MockGatewayReconcilerMockRecorder is the mock recorder for MockGatewayReconciler.
type MockGatewayReconcilerMockRecorder struct {
	mock *MockGatewayReconciler
}

// NewMockGatewayReconciler creates a new mock instance.
func NewMockGatewayReconciler(ctrl *gomock.Controller) *MockGatewayReconciler {
	mock := &MockGatewayReconciler{ctrl: ctrl}
	mock.recorder = &MockGatewayReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayReconciler) EXPECT() *MockGatewayReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGateway mocks base method.
func (m *MockGatewayReconciler) ReconcileGateway(obj *v1beta1.Gateway) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGateway", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGateway indicates an expected call of ReconcileGateway.
func (mr *MockGatewayReconcilerMockRecorder) ReconcileGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGateway", reflect.TypeOf((*MockGatewayReconciler)(nil).ReconcileGateway), obj)
}

// MockGatewayDeletionReconciler is a mock of GatewayDeletionReconciler interface.
type MockGatewayDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayDeletionReconcilerMockRecorder
}

// MockGatewayDeletionReconcilerMockRecorder is the mock recorder for MockGatewayDeletionReconciler.
type MockGatewayDeletionReconcilerMockRecorder struct {
	mock *MockGatewayDeletionReconciler
}

// NewMockGatewayDeletionReconciler creates a new mock instance.
func NewMockGatewayDeletionReconciler(ctrl *gomock.Controller) *MockGatewayDeletionReconciler {
	mock := &MockGatewayDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockGatewayDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayDeletionReconciler) EXPECT() *MockGatewayDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayDeletion mocks base method.
func (m *MockGatewayDeletionReconciler) ReconcileGatewayDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGatewayDeletion indicates an expected call of ReconcileGatewayDeletion.
func (mr *MockGatewayDeletionReconcilerMockRecorder) ReconcileGatewayDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayDeletion", reflect.TypeOf((*MockGatewayDeletionReconciler)(nil).ReconcileGatewayDeletion), req)
}

// MockGatewayFinalizer is a mock of GatewayFinalizer interface.
type MockGatewayFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayFinalizerMockRecorder
}

// MockGatewayFinalizerMockRecorder is the mock recorder for MockGatewayFinalizer.
type MockGatewayFinalizerMockRecorder struct {
	mock *MockGatewayFinalizer
}

// NewMockGatewayFinalizer creates a new mock instance.
func NewMockGatewayFinalizer(ctrl *gomock.Controller) *MockGatewayFinalizer {
	mock := &MockGatewayFinalizer{ctrl: ctrl}
	mock.recorder = &MockGatewayFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayFinalizer) EXPECT() *MockGatewayFinalizerMockRecorder {
	return m.recorder
}

// FinalizeGateway mocks base method.
func (m *MockGatewayFinalizer) FinalizeGateway(obj *v1beta1.Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeGateway indicates an expected call of FinalizeGateway.
func (mr *MockGatewayFinalizerMockRecorder) FinalizeGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeGateway", reflect.TypeOf((*MockGatewayFinalizer)(nil).FinalizeGateway), obj)
}

// GatewayFinalizerName mocks base method.
func (m *MockGatewayFinalizer) GatewayFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GatewayFinalizerName indicates an expected call of GatewayFinalizerName.
func (mr *MockGatewayFinalizerMockRecorder) GatewayFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayFinalizerName", reflect.TypeOf((*MockGatewayFinalizer)(nil).GatewayFinalizerName))
}

// ReconcileGateway mocks base method.
func (m *MockGatewayFinalizer) ReconcileGateway(obj *v1beta1.Gateway) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGateway", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGateway indicates an expected call of ReconcileGateway.
func (mr *MockGatewayFinalizerMockRecorder) ReconcileGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGateway", reflect.TypeOf((*MockGatewayFinalizer)(nil).ReconcileGateway), obj)
}

// MockGatewayReconcileLoop is a mock of GatewayReconcileLoop interface.
type MockGatewayReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayReconcileLoopMockRecorder
}

// MockGatewayReconcileLoopMockRecorder is the mock recorder for MockGatewayReconcileLoop.
type MockGatewayReconcileLoopMockRecorder struct {
	mock *MockGatewayReconcileLoop
}

// NewMockGatewayReconcileLoop creates a new mock instance.
func NewMockGatewayReconcileLoop(ctrl *gomock.Controller) *MockGatewayReconcileLoop {
	mock := &MockGatewayReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockGatewayReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayReconcileLoop) EXPECT() *MockGatewayReconcileLoopMockRecorder {
	return m.recorder
}

// RunGatewayReconciler mocks base method.
func (m *MockGatewayReconcileLoop) RunGatewayReconciler(ctx context.Context, rec controller.GatewayReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunGatewayReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunGatewayReconciler indicates an expected call of RunGatewayReconciler.
func (mr *MockGatewayReconcileLoopMockRecorder) RunGatewayReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunGatewayReconciler", reflect.TypeOf((*MockGatewayReconcileLoop)(nil).RunGatewayReconciler), varargs...)
}

// MockGatewayClassReconciler is a mock of GatewayClassReconciler interface.
type MockGatewayClassReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClassReconcilerMockRecorder
}

// MockGatewayClassReconcilerMockRecorder is the mock recorder for MockGatewayClassReconciler.
type MockGatewayClassReconcilerMockRecorder struct {
	mock *MockGatewayClassReconciler
}

// NewMockGatewayClassReconciler creates a new mock instance.
func NewMockGatewayClassReconciler(ctrl *gomock.Controller) *MockGatewayClassReconciler {
	mock := &MockGatewayClassReconciler{ctrl: ctrl}
	mock.recorder = &MockGatewayClassReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayClassReconciler) EXPECT() *MockGatewayClassReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayClass mocks base method.
func (m *MockGatewayClassReconciler) ReconcileGatewayClass(obj *v1beta1.GatewayClass) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayClass", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGatewayClass indicates an expected call of ReconcileGatewayClass.
func (mr *MockGatewayClassReconcilerMockRecorder) ReconcileGatewayClass(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayClass", reflect.TypeOf((*MockGatewayClassReconciler)(nil).ReconcileGatewayClass), obj)
}

// MockGatewayClassDeletionReconciler is a mock of GatewayClassDeletionReconciler interface.
type MockGatewayClassDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClassDeletionReconcilerMockRecorder
}

// MockGatewayClassDeletionReconcilerMockRecorder is the mock recorder for MockGatewayClassDeletionReconciler.
type MockGatewayClassDeletionReconcilerMockRecorder struct {
	mock *MockGatewayClassDeletionReconciler
}

// NewMockGatewayClassDeletionReconciler creates a new mock instance.
func NewMockGatewayClassDeletionReconciler(ctrl *gomock.Controller) *MockGatewayClassDeletionReconciler {
	mock := &MockGatewayClassDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockGatewayClassDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayClassDeletionReconciler) EXPECT() *MockGatewayClassDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayClassDeletion mocks base method.
func (m *MockGatewayClassDeletionReconciler) ReconcileGatewayClassDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayClassDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGatewayClassDeletion indicates an expected call of ReconcileGatewayClassDeletion.
func (mr *MockGatewayClassDeletionReconcilerMockRecorder) ReconcileGatewayClassDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayClassDeletion", reflect.TypeOf((*MockGatewayClassDeletionReconciler)(nil).ReconcileGatewayClassDeletion), req)
}

// MockGatewayClassFinalizer is a mock of GatewayClassFinalizer interface.
type MockGatewayClassFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClassFinalizerMockRecorder
}

// MockGatewayClassFinalizerMockRecorder is the mock recorder for MockGatewayClassFinalizer.
type MockGatewayClassFinalizerMockRecorder struct {
	mock *MockGatewayClassFinalizer
}

// NewMockGatewayClassFinalizer creates a new mock instance.
func NewMockGatewayClassFinalizer(ctrl *gomock.Controller) *MockGatewayClassFinalizer {
	mock := &MockGatewayClassFinalizer{ctrl: ctrl}
	mock.recorder = &MockGatewayClassFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayClassFinalizer) EXPECT() *MockGatewayClassFinalizerMockRecorder {
	return m.recorder
}

// FinalizeGatewayClass mocks base method.
func (m *MockGatewayClassFinalizer) FinalizeGatewayClass(obj *v1beta1.GatewayClass) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeGatewayClass", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeGatewayClass indicates an expected call of FinalizeGatewayClass.
func (mr *MockGatewayClassFinalizerMockRecorder) FinalizeGatewayClass(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeGatewayClass", reflect.TypeOf((*MockGatewayClassFinalizer)(nil).FinalizeGatewayClass), obj)
}

// GatewayClassFinalizerName mocks base method.
func (m *MockGatewayClassFinalizer) GatewayClassFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayClassFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GatewayClassFinalizerName indicates an expected call of GatewayClassFinalizerName.
func (mr *MockGatewayClassFinalizerMockRecorder) GatewayClassFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayClassFinalizerName", reflect.TypeOf((*MockGatewayClassFinalizer)(nil).GatewayClassFinalizerName))
}

// ReconcileGatewayClass mocks base method.
func (m *MockGatewayClassFinalizer) ReconcileGatewayClass(obj *v1beta1.GatewayClass) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayClass", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGatewayClass indicates an expected call of ReconcileGatewayClass.
func (mr *MockGatewayClassFinalizerMockRecorder) ReconcileGatewayClass(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayClass", reflect.TypeOf((*MockGatewayClassFinalizer)(nil).ReconcileGatewayClass), obj)
}

// MockGatewayClassReconcileLoop is a mock of GatewayClassReconcileLoop interface.
type MockGatewayClassReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClassReconcileLoopMockRecorder
}

// MockGatewayClassReconcileLoopMockRecorder is the mock recorder for MockGatewayClassReconcileLoop.
type MockGatewayClassReconcileLoopMockRecorder struct {
	mock *MockGatewayClassReconcileLoop
}

// NewMockGatewayClassReconcileLoop creates a new mock instance.
func NewMockGatewayClassReconcileLoop(ctrl *gomock.Controller) *MockGatewayClassReconcileLoop {
	mock := &MockGatewayClassReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockGatewayClassReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayClassReconcileLoop) EXPECT() *MockGatewayClassReconcileLoopMockRecorder {
	return m.recorder
}

// RunGatewayClassReconciler mocks base method.
func (m *MockGatewayClassReconcileLoop) RunGatewayClassReconciler(ctx context.Context, rec controller.GatewayClassReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunGatewayClassReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunGatewayClassReconciler indicates an expected call of RunGatewayClassReconciler.
func (mr *MockGatewayClassReconcileLoopMockRecorder) RunGatewayClassReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunGatewayClassReconciler", reflect.TypeOf((*MockGatewayClassReconcileLoop)(nil).RunGatewayClassReconciler), varargs...)
}

// MockHTTPRouteReconciler is a mock of HTTPRouteReconciler interface.
type MockHTTPRouteReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteReconcilerMockRecorder
}

// MockHTTPRouteReconcilerMockRecorder is the mock recorder for MockHTTPRouteReconciler.
type MockHTTPRouteReconcilerMockRecorder struct {
	mock *MockHTTPRouteReconciler
}

// NewMockHTTPRouteReconciler creates a new mock instance.
func NewMockHTTPRouteReconciler(ctrl *gomock.Controller) *MockHTTPRouteReconciler {
	mock := &MockHTTPRouteReconciler{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPRouteReconciler) EXPECT() *MockHTTPRouteReconcilerMockRecorder {
	return m.recorder
}

// ReconcileHTTPRoute mocks base method.
func (m *MockHTTPRouteReconciler) ReconcileHTTPRoute(obj *v1beta1.HTTPRoute) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileHTTPRoute", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileHTTPRoute indicates an expected call of ReconcileHTTPRoute.
func (mr *MockHTTPRouteReconcilerMockRecorder) ReconcileHTTPRoute(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileHTTPRoute", reflect.TypeOf((*MockHTTPRouteReconciler)(nil).ReconcileHTTPRoute), obj)
}

// MockHTTPRouteDeletionReconciler is a mock of HTTPRouteDeletionReconciler interface.
type MockHTTPRouteDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteDeletionReconcilerMockRecorder
}

// MockHTTPRouteDeletionReconcilerMockRecorder is the mock recorder for MockHTTPRouteDeletionReconciler.
type MockHTTPRouteDeletionReconcilerMockRecorder struct {
	mock *MockHTTPRouteDeletionReconciler
}

// NewMockHTTPRouteDeletionReconciler creates a new mock instance.
func NewMockHTTPRouteDeletionReconciler(ctrl *gomock.Controller) *MockHTTPRouteDeletionReconciler {
	mock := &MockHTTPRouteDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPRouteDeletionReconciler) EXPECT() *MockHTTPRouteDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileHTTPRouteDeletion mocks base method.
func (m *MockHTTPRouteDeletionReconciler) ReconcileHTTPRouteDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileHTTPRouteDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileHTTPRouteDeletion indicates an expected call of ReconcileHTTPRouteDeletion.
func (mr *MockHTTPRouteDeletionReconcilerMockRecorder) ReconcileHTTPRouteDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileHTTPRouteDeletion", reflect.TypeOf((*MockHTTPRouteDeletionReconciler)(nil).ReconcileHTTPRouteDeletion), req)
}

// MockHTTPRouteFinalizer is a mock of HTTPRouteFinalizer interface.
type MockHTTPRouteFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteFinalizerMockRecorder
}

// MockHTTPRouteFinalizerMockRecorder is the mock recorder for MockHTTPRouteFinalizer.
type MockHTTPRouteFinalizerMockRecorder struct {
	mock *MockHTTPRouteFinalizer
}

// NewMockHTTPRouteFinalizer creates a new mock instance.
func NewMockHTTPRouteFinalizer(ctrl *gomock.Controller) *MockHTTPRouteFinalizer {
	mock := &MockHTTPRouteFinalizer{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPRouteFinalizer) EXPECT() *MockHTTPRouteFinalizerMockRecorder {
	return m.recorder
}

// FinalizeHTTPRoute mocks base method.
func (m *MockHTTPRouteFinalizer) FinalizeHTTPRoute(obj *v1beta1.HTTPRoute) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeHTTPRoute", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeHTTPRoute indicates an expected call of FinalizeHTTPRoute.
func (mr *MockHTTPRouteFinalizerMockRecorder) FinalizeHTTPRoute(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeHTTPRoute", reflect.TypeOf((*MockHTTPRouteFinalizer)(nil).FinalizeHTTPRoute), obj)
}

// HTTPRouteFinalizerName mocks base method.
func (m *MockHTTPRouteFinalizer) HTTPRouteFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPRouteFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// HTTPRouteFinalizerName indicates an expected call of HTTPRouteFinalizerName.
func (mr *MockHTTPRouteFinalizerMockRecorder) HTTPRouteFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPRouteFinalizerName", reflect.TypeOf((*MockHTTPRouteFinalizer)(nil).HTTPRouteFinalizerName))
}

// ReconcileHTTPRoute mocks base method.
func (m *MockHTTPRouteFinalizer) ReconcileHTTPRoute(obj *v1beta1.HTTPRoute) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileHTTPRoute", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileHTTPRoute indicates an expected call of ReconcileHTTPRoute.
func (mr *MockHTTPRouteFinalizerMockRecorder) ReconcileHTTPRoute(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileHTTPRoute", reflect.TypeOf((*MockHTTPRouteFinalizer)(nil).ReconcileHTTPRoute), obj)
}

// MockHTTPRouteReconcileLoop is a mock of HTTPRouteReconcileLoop interface.
type MockHTTPRouteReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteReconcileLoopMockRecorder
}

// MockHTTPRouteReconcileLoopMockRecorder is the mock recorder for MockHTTPRouteReconcileLoop.
type MockHTTPRouteReconcileLoopMockRecorder struct {
	mock *MockHTTPRouteReconcileLoop
}

// NewMockHTTPRouteReconcileLoop creates a new mock instance.
func NewMockHTTPRouteReconcileLoop(ctrl *gomock.Controller) *MockHTTPRouteReconcileLoop {
	mock := &MockHTTPRouteReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPRouteReconcileLoop) EXPECT() *MockHTTPRouteReconcileLoopMockRecorder {
	return m.recorder
}

// RunHTTPRouteReconciler mocks base method.
func (m *MockHTTPRouteReconcileLoop) RunHTTPRouteReconciler(ctx context.Context, rec controller.HTTPRouteReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunHTTPRouteReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunHTTPRouteReconciler indicates an expected call of RunHTTPRouteReconciler.
func (mr *MockHTTPRouteReconcileLoopMockRecorder) RunHTTPRouteReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunHTTPRouteReconciler", reflect.TypeOf((*MockHTTPRouteReconcileLoop)(nil).RunHTTPRouteReconciler), varargs...)
}

// MockReferenceGrantReconciler is a mock of ReferenceGrantReconciler interface.
type MockReferenceGrantReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceGrantReconcilerMockRecorder
}

// MockReferenceGrantReconcilerMockRecorder is the mock recorder for MockReferenceGrantReconciler.
type MockReferenceGrantReconcilerMockRecorder struct {
	mock *MockReferenceGrantReconciler
}

// NewMockReferenceGrantReconciler creates a new mock instance.
func NewMockReferenceGrantReconciler(ctrl *gomock.Controller) *MockReferenceGrantReconciler {
	mock := &MockReferenceGrantReconciler{ctrl: ctrl}
	mock.recorder = &MockReferenceGrantReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReferenceGrantReconciler) EXPECT() *MockReferenceGrantReconcilerMockRecorder {
	return m.recorder
}

// ReconcileReferenceGrant mocks base method.
func (m *MockReferenceGrantReconciler) ReconcileReferenceGrant(obj *v1beta1.ReferenceGrant) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileReferenceGrant", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileReferenceGrant indicates an expected call of ReconcileReferenceGrant.
func (mr *MockReferenceGrantReconcilerMockRecorder) ReconcileReferenceGrant(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileReferenceGrant", reflect.TypeOf((*MockReferenceGrantReconciler)(nil).ReconcileReferenceGrant), obj)
}

// MockReferenceGrantDeletionReconciler is a mock of ReferenceGrantDeletionReconciler interface.
type MockReferenceGrantDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceGrantDeletionReconcilerMockRecorder
}

// MockReferenceGrantDeletionReconcilerMockRecorder is the mock recorder for MockReferenceGrantDeletionReconciler.
type MockReferenceGrantDeletionReconcilerMockRecorder struct {
	mock *MockReferenceGrantDeletionReconciler
}

// NewMockReferenceGrantDeletionReconciler creates a new mock instance.
func NewMockReferenceGrantDeletionReconciler(ctrl *gomock.Controller) *MockReferenceGrantDeletionReconciler {
	mock := &MockReferenceGrantDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockReferenceGrantDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReferenceGrantDeletionReconciler) EXPECT() *MockReferenceGrantDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileReferenceGrantDeletion mocks base method.
func (m *MockReferenceGrantDeletionReconciler) ReconcileReferenceGrantDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileReferenceGrantDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileReferenceGrantDeletion indicates an expected call of ReconcileReferenceGrantDeletion.
func (mr *MockReferenceGrantDeletionReconcilerMockRecorder) ReconcileReferenceGrantDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileReferenceGrantDeletion", reflect.TypeOf((*MockReferenceGrantDeletionReconciler)(nil).ReconcileReferenceGrantDeletion), req)
}

// MockReferenceGrantFinalizer is a mock of ReferenceGrantFinalizer interface.
type MockReferenceGrantFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceGrantFinalizerMockRecorder
}

// MockReferenceGrantFinalizerMockRecorder is the mock recorder for MockReferenceGrantFinalizer.
type MockReferenceGrantFinalizerMockRecorder struct {
	mock *MockReferenceGrantFinalizer
}

// NewMockReferenceGrantFinalizer creates a new mock instance.
func NewMockReferenceGrantFinalizer(ctrl *gomock.Controller) *MockReferenceGrantFinalizer {
	mock := &MockReferenceGrantFinalizer{ctrl: ctrl}
	mock.recorder = &MockReferenceGrantFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReferenceGrantFinalizer) EXPECT() *MockReferenceGrantFinalizerMockRecorder {
	return m.recorder
}

// FinalizeReferenceGrant mocks base method.
func (m *MockReferenceGrantFinalizer) FinalizeReferenceGrant(obj *v1beta1.ReferenceGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeReferenceGrant", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeReferenceGrant indicates an expected call of FinalizeReferenceGrant.
func (mr *MockReferenceGrantFinalizerMockRecorder) FinalizeReferenceGrant(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeReferenceGrant", reflect.TypeOf((*MockReferenceGrantFinalizer)(nil).FinalizeReferenceGrant), obj)
}

// ReconcileReferenceGrant mocks base method.
func (m *MockReferenceGrantFinalizer) ReconcileReferenceGrant(obj *v1beta1.ReferenceGrant) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileReferenceGrant", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileReferenceGrant indicates an expected call of ReconcileReferenceGrant.
func (mr *MockReferenceGrantFinalizerMockRecorder) ReconcileReferenceGrant(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileReferenceGrant", reflect.TypeOf((*MockReferenceGrantFinalizer)(nil).ReconcileReferenceGrant), obj)
}

// ReferenceGrantFinalizerName mocks base method.
func (m *MockReferenceGrantFinalizer) ReferenceGrantFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReferenceGrantFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ReferenceGrantFinalizerName indicates an expected call of ReferenceGrantFinalizerName.
func (mr *MockReferenceGrantFinalizerMockRecorder) ReferenceGrantFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReferenceGrantFinalizerName", reflect.TypeOf((*MockReferenceGrantFinalizer)(nil).ReferenceGrantFinalizerName))
}

// MockReferenceGrantReconcileLoop is a mock of ReferenceGrantReconcileLoop interface.
type MockReferenceGrantReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceGrantReconcileLoopMockRecorder
}

// MockReferenceGrantReconcileLoopMockRecorder is the mock recorder for MockReferenceGrantReconcileLoop.
type MockReferenceGrantReconcileLoopMockRecorder struct {
	mock *MockReferenceGrantReconcileLoop
}

// NewMockReferenceGrantReconcileLoop creates a new mock instance.
func NewMockReferenceGrantReconcileLoop(ctrl *gomock.Controller) *MockReferenceGrantReconcileLoop {
	mock := &MockReferenceGrantReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockReferenceGrantReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReferenceGrantReconcileLoop) EXPECT() *MockReferenceGrantReconcileLoopMockRecorder {
	return m.recorder
}

// RunReferenceGrantReconciler mocks base method.
func (m *MockReferenceGrantReconcileLoop) RunReferenceGrantReconciler(ctx context.Context, rec controller.ReferenceGrantReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunReferenceGrantReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunReferenceGrantReconciler indicates an expected call of RunReferenceGrantReconciler.
func (mr *MockReferenceGrantReconcileLoopMockRecorder) RunReferenceGrantReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunReferenceGrantReconciler", reflect.TypeOf((*MockReferenceGrantReconcileLoop)(nil).RunReferenceGrantReconciler), varargs...)
}
