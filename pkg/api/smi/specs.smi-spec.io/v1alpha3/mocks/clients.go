// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha3 is a generated GoMock package.
package mock_v1alpha3

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	v1alpha30 "github.com/solo-io/external-apis/pkg/api/smi/specs.smi-spec.io/v1alpha3"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1alpha30.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha30.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// HttpRouteGroups mocks base method
func (m *MockClientset) HttpRouteGroups() v1alpha30.HttpRouteGroupClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HttpRouteGroups")
	ret0, _ := ret[0].(v1alpha30.HttpRouteGroupClient)
	return ret0
}

// HttpRouteGroups indicates an expected call of HttpRouteGroups
func (mr *MockClientsetMockRecorder) HttpRouteGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HttpRouteGroups", reflect.TypeOf((*MockClientset)(nil).HttpRouteGroups))
}

// MockHttpRouteGroupReader is a mock of HttpRouteGroupReader interface
type MockHttpRouteGroupReader struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRouteGroupReaderMockRecorder
}

// MockHttpRouteGroupReaderMockRecorder is the mock recorder for MockHttpRouteGroupReader
type MockHttpRouteGroupReaderMockRecorder struct {
	mock *MockHttpRouteGroupReader
}

// NewMockHttpRouteGroupReader creates a new mock instance
func NewMockHttpRouteGroupReader(ctrl *gomock.Controller) *MockHttpRouteGroupReader {
	mock := &MockHttpRouteGroupReader{ctrl: ctrl}
	mock.recorder = &MockHttpRouteGroupReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpRouteGroupReader) EXPECT() *MockHttpRouteGroupReaderMockRecorder {
	return m.recorder
}

// GetHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupReader) GetHttpRouteGroup(ctx context.Context, key client.ObjectKey) (*v1alpha3.HttpRouteGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHttpRouteGroup", ctx, key)
	ret0, _ := ret[0].(*v1alpha3.HttpRouteGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHttpRouteGroup indicates an expected call of GetHttpRouteGroup
func (mr *MockHttpRouteGroupReaderMockRecorder) GetHttpRouteGroup(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupReader)(nil).GetHttpRouteGroup), ctx, key)
}

// ListHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupReader) ListHttpRouteGroup(ctx context.Context, opts ...client.ListOption) (*v1alpha3.HttpRouteGroupList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(*v1alpha3.HttpRouteGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHttpRouteGroup indicates an expected call of ListHttpRouteGroup
func (mr *MockHttpRouteGroupReaderMockRecorder) ListHttpRouteGroup(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupReader)(nil).ListHttpRouteGroup), varargs...)
}

// MockHttpRouteGroupWriter is a mock of HttpRouteGroupWriter interface
type MockHttpRouteGroupWriter struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRouteGroupWriterMockRecorder
}

// MockHttpRouteGroupWriterMockRecorder is the mock recorder for MockHttpRouteGroupWriter
type MockHttpRouteGroupWriterMockRecorder struct {
	mock *MockHttpRouteGroupWriter
}

// NewMockHttpRouteGroupWriter creates a new mock instance
func NewMockHttpRouteGroupWriter(ctrl *gomock.Controller) *MockHttpRouteGroupWriter {
	mock := &MockHttpRouteGroupWriter{ctrl: ctrl}
	mock.recorder = &MockHttpRouteGroupWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpRouteGroupWriter) EXPECT() *MockHttpRouteGroupWriterMockRecorder {
	return m.recorder
}

// CreateHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupWriter) CreateHttpRouteGroup(ctx context.Context, obj *v1alpha3.HttpRouteGroup, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHttpRouteGroup indicates an expected call of CreateHttpRouteGroup
func (mr *MockHttpRouteGroupWriterMockRecorder) CreateHttpRouteGroup(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupWriter)(nil).CreateHttpRouteGroup), varargs...)
}

// DeleteHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupWriter) DeleteHttpRouteGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHttpRouteGroup indicates an expected call of DeleteHttpRouteGroup
func (mr *MockHttpRouteGroupWriterMockRecorder) DeleteHttpRouteGroup(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupWriter)(nil).DeleteHttpRouteGroup), varargs...)
}

// UpdateHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupWriter) UpdateHttpRouteGroup(ctx context.Context, obj *v1alpha3.HttpRouteGroup, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHttpRouteGroup indicates an expected call of UpdateHttpRouteGroup
func (mr *MockHttpRouteGroupWriterMockRecorder) UpdateHttpRouteGroup(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupWriter)(nil).UpdateHttpRouteGroup), varargs...)
}

// PatchHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupWriter) PatchHttpRouteGroup(ctx context.Context, obj *v1alpha3.HttpRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchHttpRouteGroup indicates an expected call of PatchHttpRouteGroup
func (mr *MockHttpRouteGroupWriterMockRecorder) PatchHttpRouteGroup(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupWriter)(nil).PatchHttpRouteGroup), varargs...)
}

// DeleteAllOfHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupWriter) DeleteAllOfHttpRouteGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfHttpRouteGroup indicates an expected call of DeleteAllOfHttpRouteGroup
func (mr *MockHttpRouteGroupWriterMockRecorder) DeleteAllOfHttpRouteGroup(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupWriter)(nil).DeleteAllOfHttpRouteGroup), varargs...)
}

// UpsertHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupWriter) UpsertHttpRouteGroup(ctx context.Context, obj *v1alpha3.HttpRouteGroup, transitionFuncs ...v1alpha30.HttpRouteGroupTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertHttpRouteGroup indicates an expected call of UpsertHttpRouteGroup
func (mr *MockHttpRouteGroupWriterMockRecorder) UpsertHttpRouteGroup(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupWriter)(nil).UpsertHttpRouteGroup), varargs...)
}

// MockHttpRouteGroupStatusWriter is a mock of HttpRouteGroupStatusWriter interface
type MockHttpRouteGroupStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRouteGroupStatusWriterMockRecorder
}

// MockHttpRouteGroupStatusWriterMockRecorder is the mock recorder for MockHttpRouteGroupStatusWriter
type MockHttpRouteGroupStatusWriterMockRecorder struct {
	mock *MockHttpRouteGroupStatusWriter
}

// NewMockHttpRouteGroupStatusWriter creates a new mock instance
func NewMockHttpRouteGroupStatusWriter(ctrl *gomock.Controller) *MockHttpRouteGroupStatusWriter {
	mock := &MockHttpRouteGroupStatusWriter{ctrl: ctrl}
	mock.recorder = &MockHttpRouteGroupStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpRouteGroupStatusWriter) EXPECT() *MockHttpRouteGroupStatusWriterMockRecorder {
	return m.recorder
}

// UpdateHttpRouteGroupStatus mocks base method
func (m *MockHttpRouteGroupStatusWriter) UpdateHttpRouteGroupStatus(ctx context.Context, obj *v1alpha3.HttpRouteGroup, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateHttpRouteGroupStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHttpRouteGroupStatus indicates an expected call of UpdateHttpRouteGroupStatus
func (mr *MockHttpRouteGroupStatusWriterMockRecorder) UpdateHttpRouteGroupStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHttpRouteGroupStatus", reflect.TypeOf((*MockHttpRouteGroupStatusWriter)(nil).UpdateHttpRouteGroupStatus), varargs...)
}

// PatchHttpRouteGroupStatus mocks base method
func (m *MockHttpRouteGroupStatusWriter) PatchHttpRouteGroupStatus(ctx context.Context, obj *v1alpha3.HttpRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchHttpRouteGroupStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchHttpRouteGroupStatus indicates an expected call of PatchHttpRouteGroupStatus
func (mr *MockHttpRouteGroupStatusWriterMockRecorder) PatchHttpRouteGroupStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchHttpRouteGroupStatus", reflect.TypeOf((*MockHttpRouteGroupStatusWriter)(nil).PatchHttpRouteGroupStatus), varargs...)
}

// MockHttpRouteGroupClient is a mock of HttpRouteGroupClient interface
type MockHttpRouteGroupClient struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRouteGroupClientMockRecorder
}

// MockHttpRouteGroupClientMockRecorder is the mock recorder for MockHttpRouteGroupClient
type MockHttpRouteGroupClientMockRecorder struct {
	mock *MockHttpRouteGroupClient
}

// NewMockHttpRouteGroupClient creates a new mock instance
func NewMockHttpRouteGroupClient(ctrl *gomock.Controller) *MockHttpRouteGroupClient {
	mock := &MockHttpRouteGroupClient{ctrl: ctrl}
	mock.recorder = &MockHttpRouteGroupClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpRouteGroupClient) EXPECT() *MockHttpRouteGroupClientMockRecorder {
	return m.recorder
}

// GetHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupClient) GetHttpRouteGroup(ctx context.Context, key client.ObjectKey) (*v1alpha3.HttpRouteGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHttpRouteGroup", ctx, key)
	ret0, _ := ret[0].(*v1alpha3.HttpRouteGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHttpRouteGroup indicates an expected call of GetHttpRouteGroup
func (mr *MockHttpRouteGroupClientMockRecorder) GetHttpRouteGroup(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).GetHttpRouteGroup), ctx, key)
}

// ListHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupClient) ListHttpRouteGroup(ctx context.Context, opts ...client.ListOption) (*v1alpha3.HttpRouteGroupList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(*v1alpha3.HttpRouteGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHttpRouteGroup indicates an expected call of ListHttpRouteGroup
func (mr *MockHttpRouteGroupClientMockRecorder) ListHttpRouteGroup(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).ListHttpRouteGroup), varargs...)
}

// CreateHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupClient) CreateHttpRouteGroup(ctx context.Context, obj *v1alpha3.HttpRouteGroup, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHttpRouteGroup indicates an expected call of CreateHttpRouteGroup
func (mr *MockHttpRouteGroupClientMockRecorder) CreateHttpRouteGroup(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).CreateHttpRouteGroup), varargs...)
}

// DeleteHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupClient) DeleteHttpRouteGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHttpRouteGroup indicates an expected call of DeleteHttpRouteGroup
func (mr *MockHttpRouteGroupClientMockRecorder) DeleteHttpRouteGroup(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).DeleteHttpRouteGroup), varargs...)
}

// UpdateHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupClient) UpdateHttpRouteGroup(ctx context.Context, obj *v1alpha3.HttpRouteGroup, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHttpRouteGroup indicates an expected call of UpdateHttpRouteGroup
func (mr *MockHttpRouteGroupClientMockRecorder) UpdateHttpRouteGroup(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).UpdateHttpRouteGroup), varargs...)
}

// PatchHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupClient) PatchHttpRouteGroup(ctx context.Context, obj *v1alpha3.HttpRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchHttpRouteGroup indicates an expected call of PatchHttpRouteGroup
func (mr *MockHttpRouteGroupClientMockRecorder) PatchHttpRouteGroup(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).PatchHttpRouteGroup), varargs...)
}

// DeleteAllOfHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupClient) DeleteAllOfHttpRouteGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfHttpRouteGroup indicates an expected call of DeleteAllOfHttpRouteGroup
func (mr *MockHttpRouteGroupClientMockRecorder) DeleteAllOfHttpRouteGroup(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).DeleteAllOfHttpRouteGroup), varargs...)
}

// UpsertHttpRouteGroup mocks base method
func (m *MockHttpRouteGroupClient) UpsertHttpRouteGroup(ctx context.Context, obj *v1alpha3.HttpRouteGroup, transitionFuncs ...v1alpha30.HttpRouteGroupTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertHttpRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertHttpRouteGroup indicates an expected call of UpsertHttpRouteGroup
func (mr *MockHttpRouteGroupClientMockRecorder) UpsertHttpRouteGroup(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertHttpRouteGroup", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).UpsertHttpRouteGroup), varargs...)
}

// UpdateHttpRouteGroupStatus mocks base method
func (m *MockHttpRouteGroupClient) UpdateHttpRouteGroupStatus(ctx context.Context, obj *v1alpha3.HttpRouteGroup, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateHttpRouteGroupStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHttpRouteGroupStatus indicates an expected call of UpdateHttpRouteGroupStatus
func (mr *MockHttpRouteGroupClientMockRecorder) UpdateHttpRouteGroupStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHttpRouteGroupStatus", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).UpdateHttpRouteGroupStatus), varargs...)
}

// PatchHttpRouteGroupStatus mocks base method
func (m *MockHttpRouteGroupClient) PatchHttpRouteGroupStatus(ctx context.Context, obj *v1alpha3.HttpRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchHttpRouteGroupStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchHttpRouteGroupStatus indicates an expected call of PatchHttpRouteGroupStatus
func (mr *MockHttpRouteGroupClientMockRecorder) PatchHttpRouteGroupStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchHttpRouteGroupStatus", reflect.TypeOf((*MockHttpRouteGroupClient)(nil).PatchHttpRouteGroupStatus), varargs...)
}

// MockMulticlusterHttpRouteGroupClient is a mock of MulticlusterHttpRouteGroupClient interface
type MockMulticlusterHttpRouteGroupClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterHttpRouteGroupClientMockRecorder
}

// MockMulticlusterHttpRouteGroupClientMockRecorder is the mock recorder for MockMulticlusterHttpRouteGroupClient
type MockMulticlusterHttpRouteGroupClientMockRecorder struct {
	mock *MockMulticlusterHttpRouteGroupClient
}

// NewMockMulticlusterHttpRouteGroupClient creates a new mock instance
func NewMockMulticlusterHttpRouteGroupClient(ctrl *gomock.Controller) *MockMulticlusterHttpRouteGroupClient {
	mock := &MockMulticlusterHttpRouteGroupClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterHttpRouteGroupClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterHttpRouteGroupClient) EXPECT() *MockMulticlusterHttpRouteGroupClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterHttpRouteGroupClient) Cluster(cluster string) (v1alpha30.HttpRouteGroupClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha30.HttpRouteGroupClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterHttpRouteGroupClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterHttpRouteGroupClient)(nil).Cluster), cluster)
}
