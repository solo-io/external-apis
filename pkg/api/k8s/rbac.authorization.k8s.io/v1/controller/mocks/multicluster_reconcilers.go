// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/rbac.authorization.k8s.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "k8s.io/api/rbac/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterRoleReconciler is a mock of MulticlusterRoleReconciler interface
type MockMulticlusterRoleReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleReconcilerMockRecorder
}

// MockMulticlusterRoleReconcilerMockRecorder is the mock recorder for MockMulticlusterRoleReconciler
type MockMulticlusterRoleReconcilerMockRecorder struct {
	mock *MockMulticlusterRoleReconciler
}

// NewMockMulticlusterRoleReconciler creates a new mock instance
func NewMockMulticlusterRoleReconciler(ctrl *gomock.Controller) *MockMulticlusterRoleReconciler {
	mock := &MockMulticlusterRoleReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRoleReconciler) EXPECT() *MockMulticlusterRoleReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRole mocks base method
func (m *MockMulticlusterRoleReconciler) ReconcileRole(clusterName string, obj *v1.Role) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRole", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRole indicates an expected call of ReconcileRole
func (mr *MockMulticlusterRoleReconcilerMockRecorder) ReconcileRole(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRole", reflect.TypeOf((*MockMulticlusterRoleReconciler)(nil).ReconcileRole), clusterName, obj)
}

// MockMulticlusterRoleDeletionReconciler is a mock of MulticlusterRoleDeletionReconciler interface
type MockMulticlusterRoleDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleDeletionReconcilerMockRecorder
}

// MockMulticlusterRoleDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterRoleDeletionReconciler
type MockMulticlusterRoleDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterRoleDeletionReconciler
}

// NewMockMulticlusterRoleDeletionReconciler creates a new mock instance
func NewMockMulticlusterRoleDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterRoleDeletionReconciler {
	mock := &MockMulticlusterRoleDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRoleDeletionReconciler) EXPECT() *MockMulticlusterRoleDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleDeletion mocks base method
func (m *MockMulticlusterRoleDeletionReconciler) ReconcileRoleDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRoleDeletion indicates an expected call of ReconcileRoleDeletion
func (mr *MockMulticlusterRoleDeletionReconcilerMockRecorder) ReconcileRoleDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleDeletion", reflect.TypeOf((*MockMulticlusterRoleDeletionReconciler)(nil).ReconcileRoleDeletion), clusterName, req)
}

// MockMulticlusterRoleReconcileLoop is a mock of MulticlusterRoleReconcileLoop interface
type MockMulticlusterRoleReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleReconcileLoopMockRecorder
}

// MockMulticlusterRoleReconcileLoopMockRecorder is the mock recorder for MockMulticlusterRoleReconcileLoop
type MockMulticlusterRoleReconcileLoopMockRecorder struct {
	mock *MockMulticlusterRoleReconcileLoop
}

// NewMockMulticlusterRoleReconcileLoop creates a new mock instance
func NewMockMulticlusterRoleReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterRoleReconcileLoop {
	mock := &MockMulticlusterRoleReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRoleReconcileLoop) EXPECT() *MockMulticlusterRoleReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterRoleReconciler mocks base method
func (m *MockMulticlusterRoleReconcileLoop) AddMulticlusterRoleReconciler(ctx context.Context, rec controller.MulticlusterRoleReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterRoleReconciler", varargs...)
}

// AddMulticlusterRoleReconciler indicates an expected call of AddMulticlusterRoleReconciler
func (mr *MockMulticlusterRoleReconcileLoopMockRecorder) AddMulticlusterRoleReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterRoleReconciler", reflect.TypeOf((*MockMulticlusterRoleReconcileLoop)(nil).AddMulticlusterRoleReconciler), varargs...)
}

// MockMulticlusterRoleBindingReconciler is a mock of MulticlusterRoleBindingReconciler interface
type MockMulticlusterRoleBindingReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleBindingReconcilerMockRecorder
}

// MockMulticlusterRoleBindingReconcilerMockRecorder is the mock recorder for MockMulticlusterRoleBindingReconciler
type MockMulticlusterRoleBindingReconcilerMockRecorder struct {
	mock *MockMulticlusterRoleBindingReconciler
}

// NewMockMulticlusterRoleBindingReconciler creates a new mock instance
func NewMockMulticlusterRoleBindingReconciler(ctrl *gomock.Controller) *MockMulticlusterRoleBindingReconciler {
	mock := &MockMulticlusterRoleBindingReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleBindingReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRoleBindingReconciler) EXPECT() *MockMulticlusterRoleBindingReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleBinding mocks base method
func (m *MockMulticlusterRoleBindingReconciler) ReconcileRoleBinding(clusterName string, obj *v1.RoleBinding) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleBinding", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRoleBinding indicates an expected call of ReconcileRoleBinding
func (mr *MockMulticlusterRoleBindingReconcilerMockRecorder) ReconcileRoleBinding(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleBinding", reflect.TypeOf((*MockMulticlusterRoleBindingReconciler)(nil).ReconcileRoleBinding), clusterName, obj)
}

// MockMulticlusterRoleBindingDeletionReconciler is a mock of MulticlusterRoleBindingDeletionReconciler interface
type MockMulticlusterRoleBindingDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleBindingDeletionReconcilerMockRecorder
}

// MockMulticlusterRoleBindingDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterRoleBindingDeletionReconciler
type MockMulticlusterRoleBindingDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterRoleBindingDeletionReconciler
}

// NewMockMulticlusterRoleBindingDeletionReconciler creates a new mock instance
func NewMockMulticlusterRoleBindingDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterRoleBindingDeletionReconciler {
	mock := &MockMulticlusterRoleBindingDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleBindingDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRoleBindingDeletionReconciler) EXPECT() *MockMulticlusterRoleBindingDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleBindingDeletion mocks base method
func (m *MockMulticlusterRoleBindingDeletionReconciler) ReconcileRoleBindingDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleBindingDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRoleBindingDeletion indicates an expected call of ReconcileRoleBindingDeletion
func (mr *MockMulticlusterRoleBindingDeletionReconcilerMockRecorder) ReconcileRoleBindingDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleBindingDeletion", reflect.TypeOf((*MockMulticlusterRoleBindingDeletionReconciler)(nil).ReconcileRoleBindingDeletion), clusterName, req)
}

// MockMulticlusterRoleBindingReconcileLoop is a mock of MulticlusterRoleBindingReconcileLoop interface
type MockMulticlusterRoleBindingReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleBindingReconcileLoopMockRecorder
}

// MockMulticlusterRoleBindingReconcileLoopMockRecorder is the mock recorder for MockMulticlusterRoleBindingReconcileLoop
type MockMulticlusterRoleBindingReconcileLoopMockRecorder struct {
	mock *MockMulticlusterRoleBindingReconcileLoop
}

// NewMockMulticlusterRoleBindingReconcileLoop creates a new mock instance
func NewMockMulticlusterRoleBindingReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterRoleBindingReconcileLoop {
	mock := &MockMulticlusterRoleBindingReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleBindingReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterRoleBindingReconcileLoop) EXPECT() *MockMulticlusterRoleBindingReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterRoleBindingReconciler mocks base method
func (m *MockMulticlusterRoleBindingReconcileLoop) AddMulticlusterRoleBindingReconciler(ctx context.Context, rec controller.MulticlusterRoleBindingReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterRoleBindingReconciler", varargs...)
}

// AddMulticlusterRoleBindingReconciler indicates an expected call of AddMulticlusterRoleBindingReconciler
func (mr *MockMulticlusterRoleBindingReconcileLoopMockRecorder) AddMulticlusterRoleBindingReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterRoleBindingReconciler", reflect.TypeOf((*MockMulticlusterRoleBindingReconcileLoop)(nil).AddMulticlusterRoleBindingReconciler), varargs...)
}

// MockMulticlusterClusterRoleReconciler is a mock of MulticlusterClusterRoleReconciler interface
type MockMulticlusterClusterRoleReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterRoleReconcilerMockRecorder
}

// MockMulticlusterClusterRoleReconcilerMockRecorder is the mock recorder for MockMulticlusterClusterRoleReconciler
type MockMulticlusterClusterRoleReconcilerMockRecorder struct {
	mock *MockMulticlusterClusterRoleReconciler
}

// NewMockMulticlusterClusterRoleReconciler creates a new mock instance
func NewMockMulticlusterClusterRoleReconciler(ctrl *gomock.Controller) *MockMulticlusterClusterRoleReconciler {
	mock := &MockMulticlusterClusterRoleReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterRoleReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterRoleReconciler) EXPECT() *MockMulticlusterClusterRoleReconcilerMockRecorder {
	return m.recorder
}

// ReconcileClusterRole mocks base method
func (m *MockMulticlusterClusterRoleReconciler) ReconcileClusterRole(clusterName string, obj *v1.ClusterRole) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterRole", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileClusterRole indicates an expected call of ReconcileClusterRole
func (mr *MockMulticlusterClusterRoleReconcilerMockRecorder) ReconcileClusterRole(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterRole", reflect.TypeOf((*MockMulticlusterClusterRoleReconciler)(nil).ReconcileClusterRole), clusterName, obj)
}

// MockMulticlusterClusterRoleDeletionReconciler is a mock of MulticlusterClusterRoleDeletionReconciler interface
type MockMulticlusterClusterRoleDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterRoleDeletionReconcilerMockRecorder
}

// MockMulticlusterClusterRoleDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterClusterRoleDeletionReconciler
type MockMulticlusterClusterRoleDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterClusterRoleDeletionReconciler
}

// NewMockMulticlusterClusterRoleDeletionReconciler creates a new mock instance
func NewMockMulticlusterClusterRoleDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterClusterRoleDeletionReconciler {
	mock := &MockMulticlusterClusterRoleDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterRoleDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterRoleDeletionReconciler) EXPECT() *MockMulticlusterClusterRoleDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileClusterRoleDeletion mocks base method
func (m *MockMulticlusterClusterRoleDeletionReconciler) ReconcileClusterRoleDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterRoleDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileClusterRoleDeletion indicates an expected call of ReconcileClusterRoleDeletion
func (mr *MockMulticlusterClusterRoleDeletionReconcilerMockRecorder) ReconcileClusterRoleDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterRoleDeletion", reflect.TypeOf((*MockMulticlusterClusterRoleDeletionReconciler)(nil).ReconcileClusterRoleDeletion), clusterName, req)
}

// MockMulticlusterClusterRoleReconcileLoop is a mock of MulticlusterClusterRoleReconcileLoop interface
type MockMulticlusterClusterRoleReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterRoleReconcileLoopMockRecorder
}

// MockMulticlusterClusterRoleReconcileLoopMockRecorder is the mock recorder for MockMulticlusterClusterRoleReconcileLoop
type MockMulticlusterClusterRoleReconcileLoopMockRecorder struct {
	mock *MockMulticlusterClusterRoleReconcileLoop
}

// NewMockMulticlusterClusterRoleReconcileLoop creates a new mock instance
func NewMockMulticlusterClusterRoleReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterClusterRoleReconcileLoop {
	mock := &MockMulticlusterClusterRoleReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterRoleReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterRoleReconcileLoop) EXPECT() *MockMulticlusterClusterRoleReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterClusterRoleReconciler mocks base method
func (m *MockMulticlusterClusterRoleReconcileLoop) AddMulticlusterClusterRoleReconciler(ctx context.Context, rec controller.MulticlusterClusterRoleReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterClusterRoleReconciler", varargs...)
}

// AddMulticlusterClusterRoleReconciler indicates an expected call of AddMulticlusterClusterRoleReconciler
func (mr *MockMulticlusterClusterRoleReconcileLoopMockRecorder) AddMulticlusterClusterRoleReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterClusterRoleReconciler", reflect.TypeOf((*MockMulticlusterClusterRoleReconcileLoop)(nil).AddMulticlusterClusterRoleReconciler), varargs...)
}

// MockMulticlusterClusterRoleBindingReconciler is a mock of MulticlusterClusterRoleBindingReconciler interface
type MockMulticlusterClusterRoleBindingReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterRoleBindingReconcilerMockRecorder
}

// MockMulticlusterClusterRoleBindingReconcilerMockRecorder is the mock recorder for MockMulticlusterClusterRoleBindingReconciler
type MockMulticlusterClusterRoleBindingReconcilerMockRecorder struct {
	mock *MockMulticlusterClusterRoleBindingReconciler
}

// NewMockMulticlusterClusterRoleBindingReconciler creates a new mock instance
func NewMockMulticlusterClusterRoleBindingReconciler(ctrl *gomock.Controller) *MockMulticlusterClusterRoleBindingReconciler {
	mock := &MockMulticlusterClusterRoleBindingReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterRoleBindingReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterRoleBindingReconciler) EXPECT() *MockMulticlusterClusterRoleBindingReconcilerMockRecorder {
	return m.recorder
}

// ReconcileClusterRoleBinding mocks base method
func (m *MockMulticlusterClusterRoleBindingReconciler) ReconcileClusterRoleBinding(clusterName string, obj *v1.ClusterRoleBinding) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterRoleBinding", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileClusterRoleBinding indicates an expected call of ReconcileClusterRoleBinding
func (mr *MockMulticlusterClusterRoleBindingReconcilerMockRecorder) ReconcileClusterRoleBinding(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterRoleBinding", reflect.TypeOf((*MockMulticlusterClusterRoleBindingReconciler)(nil).ReconcileClusterRoleBinding), clusterName, obj)
}

// MockMulticlusterClusterRoleBindingDeletionReconciler is a mock of MulticlusterClusterRoleBindingDeletionReconciler interface
type MockMulticlusterClusterRoleBindingDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterRoleBindingDeletionReconcilerMockRecorder
}

// MockMulticlusterClusterRoleBindingDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterClusterRoleBindingDeletionReconciler
type MockMulticlusterClusterRoleBindingDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterClusterRoleBindingDeletionReconciler
}

// NewMockMulticlusterClusterRoleBindingDeletionReconciler creates a new mock instance
func NewMockMulticlusterClusterRoleBindingDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterClusterRoleBindingDeletionReconciler {
	mock := &MockMulticlusterClusterRoleBindingDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterRoleBindingDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterRoleBindingDeletionReconciler) EXPECT() *MockMulticlusterClusterRoleBindingDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileClusterRoleBindingDeletion mocks base method
func (m *MockMulticlusterClusterRoleBindingDeletionReconciler) ReconcileClusterRoleBindingDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterRoleBindingDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileClusterRoleBindingDeletion indicates an expected call of ReconcileClusterRoleBindingDeletion
func (mr *MockMulticlusterClusterRoleBindingDeletionReconcilerMockRecorder) ReconcileClusterRoleBindingDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterRoleBindingDeletion", reflect.TypeOf((*MockMulticlusterClusterRoleBindingDeletionReconciler)(nil).ReconcileClusterRoleBindingDeletion), clusterName, req)
}

// MockMulticlusterClusterRoleBindingReconcileLoop is a mock of MulticlusterClusterRoleBindingReconcileLoop interface
type MockMulticlusterClusterRoleBindingReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClusterRoleBindingReconcileLoopMockRecorder
}

// MockMulticlusterClusterRoleBindingReconcileLoopMockRecorder is the mock recorder for MockMulticlusterClusterRoleBindingReconcileLoop
type MockMulticlusterClusterRoleBindingReconcileLoopMockRecorder struct {
	mock *MockMulticlusterClusterRoleBindingReconcileLoop
}

// NewMockMulticlusterClusterRoleBindingReconcileLoop creates a new mock instance
func NewMockMulticlusterClusterRoleBindingReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterClusterRoleBindingReconcileLoop {
	mock := &MockMulticlusterClusterRoleBindingReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClusterRoleBindingReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClusterRoleBindingReconcileLoop) EXPECT() *MockMulticlusterClusterRoleBindingReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterClusterRoleBindingReconciler mocks base method
func (m *MockMulticlusterClusterRoleBindingReconcileLoop) AddMulticlusterClusterRoleBindingReconciler(ctx context.Context, rec controller.MulticlusterClusterRoleBindingReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterClusterRoleBindingReconciler", varargs...)
}

// AddMulticlusterClusterRoleBindingReconciler indicates an expected call of AddMulticlusterClusterRoleBindingReconciler
func (mr *MockMulticlusterClusterRoleBindingReconcileLoopMockRecorder) AddMulticlusterClusterRoleBindingReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterClusterRoleBindingReconciler", reflect.TypeOf((*MockMulticlusterClusterRoleBindingReconcileLoop)(nil).AddMulticlusterClusterRoleBindingReconciler), varargs...)
}
