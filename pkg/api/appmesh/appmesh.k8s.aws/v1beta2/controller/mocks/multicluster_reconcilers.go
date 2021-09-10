// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterMeshReconciler is a mock of MulticlusterMeshReconciler interface.
type MockMulticlusterMeshReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshReconcilerMockRecorder
}

// MockMulticlusterMeshReconcilerMockRecorder is the mock recorder for MockMulticlusterMeshReconciler.
type MockMulticlusterMeshReconcilerMockRecorder struct {
	mock *MockMulticlusterMeshReconciler
}

// NewMockMulticlusterMeshReconciler creates a new mock instance.
func NewMockMulticlusterMeshReconciler(ctrl *gomock.Controller) *MockMulticlusterMeshReconciler {
	mock := &MockMulticlusterMeshReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshReconciler) EXPECT() *MockMulticlusterMeshReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMesh mocks base method.
func (m *MockMulticlusterMeshReconciler) ReconcileMesh(clusterName string, obj *v1beta2.Mesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMesh", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMesh indicates an expected call of ReconcileMesh.
func (mr *MockMulticlusterMeshReconcilerMockRecorder) ReconcileMesh(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMesh", reflect.TypeOf((*MockMulticlusterMeshReconciler)(nil).ReconcileMesh), clusterName, obj)
}

// MockMulticlusterMeshDeletionReconciler is a mock of MulticlusterMeshDeletionReconciler interface.
type MockMulticlusterMeshDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshDeletionReconcilerMockRecorder
}

// MockMulticlusterMeshDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterMeshDeletionReconciler.
type MockMulticlusterMeshDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterMeshDeletionReconciler
}

// NewMockMulticlusterMeshDeletionReconciler creates a new mock instance.
func NewMockMulticlusterMeshDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterMeshDeletionReconciler {
	mock := &MockMulticlusterMeshDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshDeletionReconciler) EXPECT() *MockMulticlusterMeshDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshDeletion mocks base method.
func (m *MockMulticlusterMeshDeletionReconciler) ReconcileMeshDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMeshDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileMeshDeletion indicates an expected call of ReconcileMeshDeletion.
func (mr *MockMulticlusterMeshDeletionReconcilerMockRecorder) ReconcileMeshDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshDeletion", reflect.TypeOf((*MockMulticlusterMeshDeletionReconciler)(nil).ReconcileMeshDeletion), clusterName, req)
}

// MockMulticlusterMeshReconcileLoop is a mock of MulticlusterMeshReconcileLoop interface.
type MockMulticlusterMeshReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshReconcileLoopMockRecorder
}

// MockMulticlusterMeshReconcileLoopMockRecorder is the mock recorder for MockMulticlusterMeshReconcileLoop.
type MockMulticlusterMeshReconcileLoopMockRecorder struct {
	mock *MockMulticlusterMeshReconcileLoop
}

// NewMockMulticlusterMeshReconcileLoop creates a new mock instance.
func NewMockMulticlusterMeshReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterMeshReconcileLoop {
	mock := &MockMulticlusterMeshReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshReconcileLoop) EXPECT() *MockMulticlusterMeshReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterMeshReconciler mocks base method.
func (m *MockMulticlusterMeshReconcileLoop) AddMulticlusterMeshReconciler(ctx context.Context, rec controller.MulticlusterMeshReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterMeshReconciler", varargs...)
}

// AddMulticlusterMeshReconciler indicates an expected call of AddMulticlusterMeshReconciler.
func (mr *MockMulticlusterMeshReconcileLoopMockRecorder) AddMulticlusterMeshReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterMeshReconciler", reflect.TypeOf((*MockMulticlusterMeshReconcileLoop)(nil).AddMulticlusterMeshReconciler), varargs...)
}

// MockMulticlusterVirtualServiceReconciler is a mock of MulticlusterVirtualServiceReconciler interface.
type MockMulticlusterVirtualServiceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualServiceReconcilerMockRecorder
}

// MockMulticlusterVirtualServiceReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualServiceReconciler.
type MockMulticlusterVirtualServiceReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualServiceReconciler
}

// NewMockMulticlusterVirtualServiceReconciler creates a new mock instance.
func NewMockMulticlusterVirtualServiceReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualServiceReconciler {
	mock := &MockMulticlusterVirtualServiceReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualServiceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualServiceReconciler) EXPECT() *MockMulticlusterVirtualServiceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualService mocks base method.
func (m *MockMulticlusterVirtualServiceReconciler) ReconcileVirtualService(clusterName string, obj *v1beta2.VirtualService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualService", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualService indicates an expected call of ReconcileVirtualService.
func (mr *MockMulticlusterVirtualServiceReconcilerMockRecorder) ReconcileVirtualService(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualService", reflect.TypeOf((*MockMulticlusterVirtualServiceReconciler)(nil).ReconcileVirtualService), clusterName, obj)
}

// MockMulticlusterVirtualServiceDeletionReconciler is a mock of MulticlusterVirtualServiceDeletionReconciler interface.
type MockMulticlusterVirtualServiceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder
}

// MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualServiceDeletionReconciler.
type MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualServiceDeletionReconciler
}

// NewMockMulticlusterVirtualServiceDeletionReconciler creates a new mock instance.
func NewMockMulticlusterVirtualServiceDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualServiceDeletionReconciler {
	mock := &MockMulticlusterVirtualServiceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualServiceDeletionReconciler) EXPECT() *MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualServiceDeletion mocks base method.
func (m *MockMulticlusterVirtualServiceDeletionReconciler) ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualServiceDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualServiceDeletion indicates an expected call of ReconcileVirtualServiceDeletion.
func (mr *MockMulticlusterVirtualServiceDeletionReconcilerMockRecorder) ReconcileVirtualServiceDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualServiceDeletion", reflect.TypeOf((*MockMulticlusterVirtualServiceDeletionReconciler)(nil).ReconcileVirtualServiceDeletion), clusterName, req)
}

// MockMulticlusterVirtualServiceReconcileLoop is a mock of MulticlusterVirtualServiceReconcileLoop interface.
type MockMulticlusterVirtualServiceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualServiceReconcileLoopMockRecorder
}

// MockMulticlusterVirtualServiceReconcileLoopMockRecorder is the mock recorder for MockMulticlusterVirtualServiceReconcileLoop.
type MockMulticlusterVirtualServiceReconcileLoopMockRecorder struct {
	mock *MockMulticlusterVirtualServiceReconcileLoop
}

// NewMockMulticlusterVirtualServiceReconcileLoop creates a new mock instance.
func NewMockMulticlusterVirtualServiceReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterVirtualServiceReconcileLoop {
	mock := &MockMulticlusterVirtualServiceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualServiceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualServiceReconcileLoop) EXPECT() *MockMulticlusterVirtualServiceReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterVirtualServiceReconciler mocks base method.
func (m *MockMulticlusterVirtualServiceReconcileLoop) AddMulticlusterVirtualServiceReconciler(ctx context.Context, rec controller.MulticlusterVirtualServiceReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterVirtualServiceReconciler", varargs...)
}

// AddMulticlusterVirtualServiceReconciler indicates an expected call of AddMulticlusterVirtualServiceReconciler.
func (mr *MockMulticlusterVirtualServiceReconcileLoopMockRecorder) AddMulticlusterVirtualServiceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterVirtualServiceReconciler", reflect.TypeOf((*MockMulticlusterVirtualServiceReconcileLoop)(nil).AddMulticlusterVirtualServiceReconciler), varargs...)
}

// MockMulticlusterVirtualNodeReconciler is a mock of MulticlusterVirtualNodeReconciler interface.
type MockMulticlusterVirtualNodeReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualNodeReconcilerMockRecorder
}

// MockMulticlusterVirtualNodeReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualNodeReconciler.
type MockMulticlusterVirtualNodeReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualNodeReconciler
}

// NewMockMulticlusterVirtualNodeReconciler creates a new mock instance.
func NewMockMulticlusterVirtualNodeReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualNodeReconciler {
	mock := &MockMulticlusterVirtualNodeReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualNodeReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualNodeReconciler) EXPECT() *MockMulticlusterVirtualNodeReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualNode mocks base method.
func (m *MockMulticlusterVirtualNodeReconciler) ReconcileVirtualNode(clusterName string, obj *v1beta2.VirtualNode) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualNode", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualNode indicates an expected call of ReconcileVirtualNode.
func (mr *MockMulticlusterVirtualNodeReconcilerMockRecorder) ReconcileVirtualNode(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualNode", reflect.TypeOf((*MockMulticlusterVirtualNodeReconciler)(nil).ReconcileVirtualNode), clusterName, obj)
}

// MockMulticlusterVirtualNodeDeletionReconciler is a mock of MulticlusterVirtualNodeDeletionReconciler interface.
type MockMulticlusterVirtualNodeDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualNodeDeletionReconcilerMockRecorder
}

// MockMulticlusterVirtualNodeDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualNodeDeletionReconciler.
type MockMulticlusterVirtualNodeDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualNodeDeletionReconciler
}

// NewMockMulticlusterVirtualNodeDeletionReconciler creates a new mock instance.
func NewMockMulticlusterVirtualNodeDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualNodeDeletionReconciler {
	mock := &MockMulticlusterVirtualNodeDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualNodeDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualNodeDeletionReconciler) EXPECT() *MockMulticlusterVirtualNodeDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualNodeDeletion mocks base method.
func (m *MockMulticlusterVirtualNodeDeletionReconciler) ReconcileVirtualNodeDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualNodeDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualNodeDeletion indicates an expected call of ReconcileVirtualNodeDeletion.
func (mr *MockMulticlusterVirtualNodeDeletionReconcilerMockRecorder) ReconcileVirtualNodeDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualNodeDeletion", reflect.TypeOf((*MockMulticlusterVirtualNodeDeletionReconciler)(nil).ReconcileVirtualNodeDeletion), clusterName, req)
}

// MockMulticlusterVirtualNodeReconcileLoop is a mock of MulticlusterVirtualNodeReconcileLoop interface.
type MockMulticlusterVirtualNodeReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualNodeReconcileLoopMockRecorder
}

// MockMulticlusterVirtualNodeReconcileLoopMockRecorder is the mock recorder for MockMulticlusterVirtualNodeReconcileLoop.
type MockMulticlusterVirtualNodeReconcileLoopMockRecorder struct {
	mock *MockMulticlusterVirtualNodeReconcileLoop
}

// NewMockMulticlusterVirtualNodeReconcileLoop creates a new mock instance.
func NewMockMulticlusterVirtualNodeReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterVirtualNodeReconcileLoop {
	mock := &MockMulticlusterVirtualNodeReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualNodeReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualNodeReconcileLoop) EXPECT() *MockMulticlusterVirtualNodeReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterVirtualNodeReconciler mocks base method.
func (m *MockMulticlusterVirtualNodeReconcileLoop) AddMulticlusterVirtualNodeReconciler(ctx context.Context, rec controller.MulticlusterVirtualNodeReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterVirtualNodeReconciler", varargs...)
}

// AddMulticlusterVirtualNodeReconciler indicates an expected call of AddMulticlusterVirtualNodeReconciler.
func (mr *MockMulticlusterVirtualNodeReconcileLoopMockRecorder) AddMulticlusterVirtualNodeReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterVirtualNodeReconciler", reflect.TypeOf((*MockMulticlusterVirtualNodeReconcileLoop)(nil).AddMulticlusterVirtualNodeReconciler), varargs...)
}

// MockMulticlusterVirtualRouterReconciler is a mock of MulticlusterVirtualRouterReconciler interface.
type MockMulticlusterVirtualRouterReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualRouterReconcilerMockRecorder
}

// MockMulticlusterVirtualRouterReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualRouterReconciler.
type MockMulticlusterVirtualRouterReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualRouterReconciler
}

// NewMockMulticlusterVirtualRouterReconciler creates a new mock instance.
func NewMockMulticlusterVirtualRouterReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualRouterReconciler {
	mock := &MockMulticlusterVirtualRouterReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualRouterReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualRouterReconciler) EXPECT() *MockMulticlusterVirtualRouterReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualRouter mocks base method.
func (m *MockMulticlusterVirtualRouterReconciler) ReconcileVirtualRouter(clusterName string, obj *v1beta2.VirtualRouter) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualRouter", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualRouter indicates an expected call of ReconcileVirtualRouter.
func (mr *MockMulticlusterVirtualRouterReconcilerMockRecorder) ReconcileVirtualRouter(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualRouter", reflect.TypeOf((*MockMulticlusterVirtualRouterReconciler)(nil).ReconcileVirtualRouter), clusterName, obj)
}

// MockMulticlusterVirtualRouterDeletionReconciler is a mock of MulticlusterVirtualRouterDeletionReconciler interface.
type MockMulticlusterVirtualRouterDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualRouterDeletionReconcilerMockRecorder
}

// MockMulticlusterVirtualRouterDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualRouterDeletionReconciler.
type MockMulticlusterVirtualRouterDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualRouterDeletionReconciler
}

// NewMockMulticlusterVirtualRouterDeletionReconciler creates a new mock instance.
func NewMockMulticlusterVirtualRouterDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualRouterDeletionReconciler {
	mock := &MockMulticlusterVirtualRouterDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualRouterDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualRouterDeletionReconciler) EXPECT() *MockMulticlusterVirtualRouterDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualRouterDeletion mocks base method.
func (m *MockMulticlusterVirtualRouterDeletionReconciler) ReconcileVirtualRouterDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualRouterDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualRouterDeletion indicates an expected call of ReconcileVirtualRouterDeletion.
func (mr *MockMulticlusterVirtualRouterDeletionReconcilerMockRecorder) ReconcileVirtualRouterDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualRouterDeletion", reflect.TypeOf((*MockMulticlusterVirtualRouterDeletionReconciler)(nil).ReconcileVirtualRouterDeletion), clusterName, req)
}

// MockMulticlusterVirtualRouterReconcileLoop is a mock of MulticlusterVirtualRouterReconcileLoop interface.
type MockMulticlusterVirtualRouterReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualRouterReconcileLoopMockRecorder
}

// MockMulticlusterVirtualRouterReconcileLoopMockRecorder is the mock recorder for MockMulticlusterVirtualRouterReconcileLoop.
type MockMulticlusterVirtualRouterReconcileLoopMockRecorder struct {
	mock *MockMulticlusterVirtualRouterReconcileLoop
}

// NewMockMulticlusterVirtualRouterReconcileLoop creates a new mock instance.
func NewMockMulticlusterVirtualRouterReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterVirtualRouterReconcileLoop {
	mock := &MockMulticlusterVirtualRouterReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualRouterReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualRouterReconcileLoop) EXPECT() *MockMulticlusterVirtualRouterReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterVirtualRouterReconciler mocks base method.
func (m *MockMulticlusterVirtualRouterReconcileLoop) AddMulticlusterVirtualRouterReconciler(ctx context.Context, rec controller.MulticlusterVirtualRouterReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterVirtualRouterReconciler", varargs...)
}

// AddMulticlusterVirtualRouterReconciler indicates an expected call of AddMulticlusterVirtualRouterReconciler.
func (mr *MockMulticlusterVirtualRouterReconcileLoopMockRecorder) AddMulticlusterVirtualRouterReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterVirtualRouterReconciler", reflect.TypeOf((*MockMulticlusterVirtualRouterReconcileLoop)(nil).AddMulticlusterVirtualRouterReconciler), varargs...)
}

// MockMulticlusterVirtualGatewayReconciler is a mock of MulticlusterVirtualGatewayReconciler interface.
type MockMulticlusterVirtualGatewayReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualGatewayReconcilerMockRecorder
}

// MockMulticlusterVirtualGatewayReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualGatewayReconciler.
type MockMulticlusterVirtualGatewayReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualGatewayReconciler
}

// NewMockMulticlusterVirtualGatewayReconciler creates a new mock instance.
func NewMockMulticlusterVirtualGatewayReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualGatewayReconciler {
	mock := &MockMulticlusterVirtualGatewayReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualGatewayReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualGatewayReconciler) EXPECT() *MockMulticlusterVirtualGatewayReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualGateway mocks base method.
func (m *MockMulticlusterVirtualGatewayReconciler) ReconcileVirtualGateway(clusterName string, obj *v1beta2.VirtualGateway) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualGateway", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualGateway indicates an expected call of ReconcileVirtualGateway.
func (mr *MockMulticlusterVirtualGatewayReconcilerMockRecorder) ReconcileVirtualGateway(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualGateway", reflect.TypeOf((*MockMulticlusterVirtualGatewayReconciler)(nil).ReconcileVirtualGateway), clusterName, obj)
}

// MockMulticlusterVirtualGatewayDeletionReconciler is a mock of MulticlusterVirtualGatewayDeletionReconciler interface.
type MockMulticlusterVirtualGatewayDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualGatewayDeletionReconcilerMockRecorder
}

// MockMulticlusterVirtualGatewayDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterVirtualGatewayDeletionReconciler.
type MockMulticlusterVirtualGatewayDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterVirtualGatewayDeletionReconciler
}

// NewMockMulticlusterVirtualGatewayDeletionReconciler creates a new mock instance.
func NewMockMulticlusterVirtualGatewayDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterVirtualGatewayDeletionReconciler {
	mock := &MockMulticlusterVirtualGatewayDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualGatewayDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualGatewayDeletionReconciler) EXPECT() *MockMulticlusterVirtualGatewayDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualGatewayDeletion mocks base method.
func (m *MockMulticlusterVirtualGatewayDeletionReconciler) ReconcileVirtualGatewayDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualGatewayDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualGatewayDeletion indicates an expected call of ReconcileVirtualGatewayDeletion.
func (mr *MockMulticlusterVirtualGatewayDeletionReconcilerMockRecorder) ReconcileVirtualGatewayDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualGatewayDeletion", reflect.TypeOf((*MockMulticlusterVirtualGatewayDeletionReconciler)(nil).ReconcileVirtualGatewayDeletion), clusterName, req)
}

// MockMulticlusterVirtualGatewayReconcileLoop is a mock of MulticlusterVirtualGatewayReconcileLoop interface.
type MockMulticlusterVirtualGatewayReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterVirtualGatewayReconcileLoopMockRecorder
}

// MockMulticlusterVirtualGatewayReconcileLoopMockRecorder is the mock recorder for MockMulticlusterVirtualGatewayReconcileLoop.
type MockMulticlusterVirtualGatewayReconcileLoopMockRecorder struct {
	mock *MockMulticlusterVirtualGatewayReconcileLoop
}

// NewMockMulticlusterVirtualGatewayReconcileLoop creates a new mock instance.
func NewMockMulticlusterVirtualGatewayReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterVirtualGatewayReconcileLoop {
	mock := &MockMulticlusterVirtualGatewayReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterVirtualGatewayReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterVirtualGatewayReconcileLoop) EXPECT() *MockMulticlusterVirtualGatewayReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterVirtualGatewayReconciler mocks base method.
func (m *MockMulticlusterVirtualGatewayReconcileLoop) AddMulticlusterVirtualGatewayReconciler(ctx context.Context, rec controller.MulticlusterVirtualGatewayReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterVirtualGatewayReconciler", varargs...)
}

// AddMulticlusterVirtualGatewayReconciler indicates an expected call of AddMulticlusterVirtualGatewayReconciler.
func (mr *MockMulticlusterVirtualGatewayReconcileLoopMockRecorder) AddMulticlusterVirtualGatewayReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterVirtualGatewayReconciler", reflect.TypeOf((*MockMulticlusterVirtualGatewayReconcileLoop)(nil).AddMulticlusterVirtualGatewayReconciler), varargs...)
}

// MockMulticlusterGatewayRouteReconciler is a mock of MulticlusterGatewayRouteReconciler interface.
type MockMulticlusterGatewayRouteReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGatewayRouteReconcilerMockRecorder
}

// MockMulticlusterGatewayRouteReconcilerMockRecorder is the mock recorder for MockMulticlusterGatewayRouteReconciler.
type MockMulticlusterGatewayRouteReconcilerMockRecorder struct {
	mock *MockMulticlusterGatewayRouteReconciler
}

// NewMockMulticlusterGatewayRouteReconciler creates a new mock instance.
func NewMockMulticlusterGatewayRouteReconciler(ctrl *gomock.Controller) *MockMulticlusterGatewayRouteReconciler {
	mock := &MockMulticlusterGatewayRouteReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGatewayRouteReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGatewayRouteReconciler) EXPECT() *MockMulticlusterGatewayRouteReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayRoute mocks base method.
func (m *MockMulticlusterGatewayRouteReconciler) ReconcileGatewayRoute(clusterName string, obj *v1beta2.GatewayRoute) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayRoute", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGatewayRoute indicates an expected call of ReconcileGatewayRoute.
func (mr *MockMulticlusterGatewayRouteReconcilerMockRecorder) ReconcileGatewayRoute(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayRoute", reflect.TypeOf((*MockMulticlusterGatewayRouteReconciler)(nil).ReconcileGatewayRoute), clusterName, obj)
}

// MockMulticlusterGatewayRouteDeletionReconciler is a mock of MulticlusterGatewayRouteDeletionReconciler interface.
type MockMulticlusterGatewayRouteDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGatewayRouteDeletionReconcilerMockRecorder
}

// MockMulticlusterGatewayRouteDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterGatewayRouteDeletionReconciler.
type MockMulticlusterGatewayRouteDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterGatewayRouteDeletionReconciler
}

// NewMockMulticlusterGatewayRouteDeletionReconciler creates a new mock instance.
func NewMockMulticlusterGatewayRouteDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterGatewayRouteDeletionReconciler {
	mock := &MockMulticlusterGatewayRouteDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGatewayRouteDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGatewayRouteDeletionReconciler) EXPECT() *MockMulticlusterGatewayRouteDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayRouteDeletion mocks base method.
func (m *MockMulticlusterGatewayRouteDeletionReconciler) ReconcileGatewayRouteDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGatewayRouteDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileGatewayRouteDeletion indicates an expected call of ReconcileGatewayRouteDeletion.
func (mr *MockMulticlusterGatewayRouteDeletionReconcilerMockRecorder) ReconcileGatewayRouteDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayRouteDeletion", reflect.TypeOf((*MockMulticlusterGatewayRouteDeletionReconciler)(nil).ReconcileGatewayRouteDeletion), clusterName, req)
}

// MockMulticlusterGatewayRouteReconcileLoop is a mock of MulticlusterGatewayRouteReconcileLoop interface.
type MockMulticlusterGatewayRouteReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterGatewayRouteReconcileLoopMockRecorder
}

// MockMulticlusterGatewayRouteReconcileLoopMockRecorder is the mock recorder for MockMulticlusterGatewayRouteReconcileLoop.
type MockMulticlusterGatewayRouteReconcileLoopMockRecorder struct {
	mock *MockMulticlusterGatewayRouteReconcileLoop
}

// NewMockMulticlusterGatewayRouteReconcileLoop creates a new mock instance.
func NewMockMulticlusterGatewayRouteReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterGatewayRouteReconcileLoop {
	mock := &MockMulticlusterGatewayRouteReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterGatewayRouteReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterGatewayRouteReconcileLoop) EXPECT() *MockMulticlusterGatewayRouteReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterGatewayRouteReconciler mocks base method.
func (m *MockMulticlusterGatewayRouteReconcileLoop) AddMulticlusterGatewayRouteReconciler(ctx context.Context, rec controller.MulticlusterGatewayRouteReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterGatewayRouteReconciler", varargs...)
}

// AddMulticlusterGatewayRouteReconciler indicates an expected call of AddMulticlusterGatewayRouteReconciler.
func (mr *MockMulticlusterGatewayRouteReconcileLoopMockRecorder) AddMulticlusterGatewayRouteReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterGatewayRouteReconciler", reflect.TypeOf((*MockMulticlusterGatewayRouteReconcileLoop)(nil).AddMulticlusterGatewayRouteReconciler), varargs...)
}
