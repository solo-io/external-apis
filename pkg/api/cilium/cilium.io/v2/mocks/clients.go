// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v2 is a generated GoMock package.
package mock_v2

import (
	context "context"
	reflect "reflect"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	gomock "github.com/golang/mock/gomock"
	v20 "github.com/solo-io/external-apis/pkg/api/cilium/cilium.io/v2"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface.
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset.
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance.
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterClientset) Cluster(cluster string) (v20.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v20.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface.
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset.
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance.
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// CiliumNetworkPolicies mocks base method.
func (m *MockClientset) CiliumNetworkPolicies() v20.CiliumNetworkPolicyClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CiliumNetworkPolicies")
	ret0, _ := ret[0].(v20.CiliumNetworkPolicyClient)
	return ret0
}

// CiliumNetworkPolicies indicates an expected call of CiliumNetworkPolicies.
func (mr *MockClientsetMockRecorder) CiliumNetworkPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CiliumNetworkPolicies", reflect.TypeOf((*MockClientset)(nil).CiliumNetworkPolicies))
}

// MockCiliumNetworkPolicyReader is a mock of CiliumNetworkPolicyReader interface.
type MockCiliumNetworkPolicyReader struct {
	ctrl     *gomock.Controller
	recorder *MockCiliumNetworkPolicyReaderMockRecorder
}

// MockCiliumNetworkPolicyReaderMockRecorder is the mock recorder for MockCiliumNetworkPolicyReader.
type MockCiliumNetworkPolicyReaderMockRecorder struct {
	mock *MockCiliumNetworkPolicyReader
}

// NewMockCiliumNetworkPolicyReader creates a new mock instance.
func NewMockCiliumNetworkPolicyReader(ctrl *gomock.Controller) *MockCiliumNetworkPolicyReader {
	mock := &MockCiliumNetworkPolicyReader{ctrl: ctrl}
	mock.recorder = &MockCiliumNetworkPolicyReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCiliumNetworkPolicyReader) EXPECT() *MockCiliumNetworkPolicyReaderMockRecorder {
	return m.recorder
}

// GetCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyReader) GetCiliumNetworkPolicy(ctx context.Context, key client.ObjectKey) (*v2.CiliumNetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCiliumNetworkPolicy", ctx, key)
	ret0, _ := ret[0].(*v2.CiliumNetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCiliumNetworkPolicy indicates an expected call of GetCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyReaderMockRecorder) GetCiliumNetworkPolicy(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyReader)(nil).GetCiliumNetworkPolicy), ctx, key)
}

// ListCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyReader) ListCiliumNetworkPolicy(ctx context.Context, opts ...client.ListOption) (*v2.CiliumNetworkPolicyList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(*v2.CiliumNetworkPolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCiliumNetworkPolicy indicates an expected call of ListCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyReaderMockRecorder) ListCiliumNetworkPolicy(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyReader)(nil).ListCiliumNetworkPolicy), varargs...)
}

// MockCiliumNetworkPolicyWriter is a mock of CiliumNetworkPolicyWriter interface.
type MockCiliumNetworkPolicyWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCiliumNetworkPolicyWriterMockRecorder
}

// MockCiliumNetworkPolicyWriterMockRecorder is the mock recorder for MockCiliumNetworkPolicyWriter.
type MockCiliumNetworkPolicyWriterMockRecorder struct {
	mock *MockCiliumNetworkPolicyWriter
}

// NewMockCiliumNetworkPolicyWriter creates a new mock instance.
func NewMockCiliumNetworkPolicyWriter(ctrl *gomock.Controller) *MockCiliumNetworkPolicyWriter {
	mock := &MockCiliumNetworkPolicyWriter{ctrl: ctrl}
	mock.recorder = &MockCiliumNetworkPolicyWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCiliumNetworkPolicyWriter) EXPECT() *MockCiliumNetworkPolicyWriterMockRecorder {
	return m.recorder
}

// CreateCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyWriter) CreateCiliumNetworkPolicy(ctx context.Context, obj *v2.CiliumNetworkPolicy, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCiliumNetworkPolicy indicates an expected call of CreateCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyWriterMockRecorder) CreateCiliumNetworkPolicy(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyWriter)(nil).CreateCiliumNetworkPolicy), varargs...)
}

// DeleteAllOfCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyWriter) DeleteAllOfCiliumNetworkPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfCiliumNetworkPolicy indicates an expected call of DeleteAllOfCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyWriterMockRecorder) DeleteAllOfCiliumNetworkPolicy(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyWriter)(nil).DeleteAllOfCiliumNetworkPolicy), varargs...)
}

// DeleteCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyWriter) DeleteCiliumNetworkPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCiliumNetworkPolicy indicates an expected call of DeleteCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyWriterMockRecorder) DeleteCiliumNetworkPolicy(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyWriter)(nil).DeleteCiliumNetworkPolicy), varargs...)
}

// PatchCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyWriter) PatchCiliumNetworkPolicy(ctx context.Context, obj *v2.CiliumNetworkPolicy, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCiliumNetworkPolicy indicates an expected call of PatchCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyWriterMockRecorder) PatchCiliumNetworkPolicy(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyWriter)(nil).PatchCiliumNetworkPolicy), varargs...)
}

// UpdateCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyWriter) UpdateCiliumNetworkPolicy(ctx context.Context, obj *v2.CiliumNetworkPolicy, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCiliumNetworkPolicy indicates an expected call of UpdateCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyWriterMockRecorder) UpdateCiliumNetworkPolicy(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyWriter)(nil).UpdateCiliumNetworkPolicy), varargs...)
}

// UpsertCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyWriter) UpsertCiliumNetworkPolicy(ctx context.Context, obj *v2.CiliumNetworkPolicy, transitionFuncs ...v20.CiliumNetworkPolicyTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertCiliumNetworkPolicy indicates an expected call of UpsertCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyWriterMockRecorder) UpsertCiliumNetworkPolicy(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyWriter)(nil).UpsertCiliumNetworkPolicy), varargs...)
}

// MockCiliumNetworkPolicyStatusWriter is a mock of CiliumNetworkPolicyStatusWriter interface.
type MockCiliumNetworkPolicyStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCiliumNetworkPolicyStatusWriterMockRecorder
}

// MockCiliumNetworkPolicyStatusWriterMockRecorder is the mock recorder for MockCiliumNetworkPolicyStatusWriter.
type MockCiliumNetworkPolicyStatusWriterMockRecorder struct {
	mock *MockCiliumNetworkPolicyStatusWriter
}

// NewMockCiliumNetworkPolicyStatusWriter creates a new mock instance.
func NewMockCiliumNetworkPolicyStatusWriter(ctrl *gomock.Controller) *MockCiliumNetworkPolicyStatusWriter {
	mock := &MockCiliumNetworkPolicyStatusWriter{ctrl: ctrl}
	mock.recorder = &MockCiliumNetworkPolicyStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCiliumNetworkPolicyStatusWriter) EXPECT() *MockCiliumNetworkPolicyStatusWriterMockRecorder {
	return m.recorder
}

// PatchCiliumNetworkPolicyStatus mocks base method.
func (m *MockCiliumNetworkPolicyStatusWriter) PatchCiliumNetworkPolicyStatus(ctx context.Context, obj *v2.CiliumNetworkPolicy, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCiliumNetworkPolicyStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCiliumNetworkPolicyStatus indicates an expected call of PatchCiliumNetworkPolicyStatus.
func (mr *MockCiliumNetworkPolicyStatusWriterMockRecorder) PatchCiliumNetworkPolicyStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCiliumNetworkPolicyStatus", reflect.TypeOf((*MockCiliumNetworkPolicyStatusWriter)(nil).PatchCiliumNetworkPolicyStatus), varargs...)
}

// UpdateCiliumNetworkPolicyStatus mocks base method.
func (m *MockCiliumNetworkPolicyStatusWriter) UpdateCiliumNetworkPolicyStatus(ctx context.Context, obj *v2.CiliumNetworkPolicy, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCiliumNetworkPolicyStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCiliumNetworkPolicyStatus indicates an expected call of UpdateCiliumNetworkPolicyStatus.
func (mr *MockCiliumNetworkPolicyStatusWriterMockRecorder) UpdateCiliumNetworkPolicyStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCiliumNetworkPolicyStatus", reflect.TypeOf((*MockCiliumNetworkPolicyStatusWriter)(nil).UpdateCiliumNetworkPolicyStatus), varargs...)
}

// MockCiliumNetworkPolicyClient is a mock of CiliumNetworkPolicyClient interface.
type MockCiliumNetworkPolicyClient struct {
	ctrl     *gomock.Controller
	recorder *MockCiliumNetworkPolicyClientMockRecorder
}

// MockCiliumNetworkPolicyClientMockRecorder is the mock recorder for MockCiliumNetworkPolicyClient.
type MockCiliumNetworkPolicyClientMockRecorder struct {
	mock *MockCiliumNetworkPolicyClient
}

// NewMockCiliumNetworkPolicyClient creates a new mock instance.
func NewMockCiliumNetworkPolicyClient(ctrl *gomock.Controller) *MockCiliumNetworkPolicyClient {
	mock := &MockCiliumNetworkPolicyClient{ctrl: ctrl}
	mock.recorder = &MockCiliumNetworkPolicyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCiliumNetworkPolicyClient) EXPECT() *MockCiliumNetworkPolicyClientMockRecorder {
	return m.recorder
}

// CreateCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyClient) CreateCiliumNetworkPolicy(ctx context.Context, obj *v2.CiliumNetworkPolicy, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCiliumNetworkPolicy indicates an expected call of CreateCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) CreateCiliumNetworkPolicy(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).CreateCiliumNetworkPolicy), varargs...)
}

// DeleteAllOfCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyClient) DeleteAllOfCiliumNetworkPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfCiliumNetworkPolicy indicates an expected call of DeleteAllOfCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) DeleteAllOfCiliumNetworkPolicy(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).DeleteAllOfCiliumNetworkPolicy), varargs...)
}

// DeleteCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyClient) DeleteCiliumNetworkPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCiliumNetworkPolicy indicates an expected call of DeleteCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) DeleteCiliumNetworkPolicy(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).DeleteCiliumNetworkPolicy), varargs...)
}

// GetCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyClient) GetCiliumNetworkPolicy(ctx context.Context, key client.ObjectKey) (*v2.CiliumNetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCiliumNetworkPolicy", ctx, key)
	ret0, _ := ret[0].(*v2.CiliumNetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCiliumNetworkPolicy indicates an expected call of GetCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) GetCiliumNetworkPolicy(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).GetCiliumNetworkPolicy), ctx, key)
}

// ListCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyClient) ListCiliumNetworkPolicy(ctx context.Context, opts ...client.ListOption) (*v2.CiliumNetworkPolicyList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(*v2.CiliumNetworkPolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCiliumNetworkPolicy indicates an expected call of ListCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) ListCiliumNetworkPolicy(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).ListCiliumNetworkPolicy), varargs...)
}

// PatchCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyClient) PatchCiliumNetworkPolicy(ctx context.Context, obj *v2.CiliumNetworkPolicy, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCiliumNetworkPolicy indicates an expected call of PatchCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) PatchCiliumNetworkPolicy(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).PatchCiliumNetworkPolicy), varargs...)
}

// PatchCiliumNetworkPolicyStatus mocks base method.
func (m *MockCiliumNetworkPolicyClient) PatchCiliumNetworkPolicyStatus(ctx context.Context, obj *v2.CiliumNetworkPolicy, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCiliumNetworkPolicyStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCiliumNetworkPolicyStatus indicates an expected call of PatchCiliumNetworkPolicyStatus.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) PatchCiliumNetworkPolicyStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCiliumNetworkPolicyStatus", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).PatchCiliumNetworkPolicyStatus), varargs...)
}

// UpdateCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyClient) UpdateCiliumNetworkPolicy(ctx context.Context, obj *v2.CiliumNetworkPolicy, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCiliumNetworkPolicy indicates an expected call of UpdateCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) UpdateCiliumNetworkPolicy(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).UpdateCiliumNetworkPolicy), varargs...)
}

// UpdateCiliumNetworkPolicyStatus mocks base method.
func (m *MockCiliumNetworkPolicyClient) UpdateCiliumNetworkPolicyStatus(ctx context.Context, obj *v2.CiliumNetworkPolicy, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCiliumNetworkPolicyStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCiliumNetworkPolicyStatus indicates an expected call of UpdateCiliumNetworkPolicyStatus.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) UpdateCiliumNetworkPolicyStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCiliumNetworkPolicyStatus", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).UpdateCiliumNetworkPolicyStatus), varargs...)
}

// UpsertCiliumNetworkPolicy mocks base method.
func (m *MockCiliumNetworkPolicyClient) UpsertCiliumNetworkPolicy(ctx context.Context, obj *v2.CiliumNetworkPolicy, transitionFuncs ...v20.CiliumNetworkPolicyTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertCiliumNetworkPolicy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertCiliumNetworkPolicy indicates an expected call of UpsertCiliumNetworkPolicy.
func (mr *MockCiliumNetworkPolicyClientMockRecorder) UpsertCiliumNetworkPolicy(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertCiliumNetworkPolicy", reflect.TypeOf((*MockCiliumNetworkPolicyClient)(nil).UpsertCiliumNetworkPolicy), varargs...)
}

// MockMulticlusterCiliumNetworkPolicyClient is a mock of MulticlusterCiliumNetworkPolicyClient interface.
type MockMulticlusterCiliumNetworkPolicyClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCiliumNetworkPolicyClientMockRecorder
}

// MockMulticlusterCiliumNetworkPolicyClientMockRecorder is the mock recorder for MockMulticlusterCiliumNetworkPolicyClient.
type MockMulticlusterCiliumNetworkPolicyClientMockRecorder struct {
	mock *MockMulticlusterCiliumNetworkPolicyClient
}

// NewMockMulticlusterCiliumNetworkPolicyClient creates a new mock instance.
func NewMockMulticlusterCiliumNetworkPolicyClient(ctrl *gomock.Controller) *MockMulticlusterCiliumNetworkPolicyClient {
	mock := &MockMulticlusterCiliumNetworkPolicyClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCiliumNetworkPolicyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterCiliumNetworkPolicyClient) EXPECT() *MockMulticlusterCiliumNetworkPolicyClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterCiliumNetworkPolicyClient) Cluster(cluster string) (v20.CiliumNetworkPolicyClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v20.CiliumNetworkPolicyClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterCiliumNetworkPolicyClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterCiliumNetworkPolicyClient)(nil).Cluster), cluster)
}
