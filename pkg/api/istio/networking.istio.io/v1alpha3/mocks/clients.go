// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha3 is a generated GoMock package.
package mock_v1alpha3

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3"
	v1alpha30 "istio.io/client-go/pkg/apis/networking/v1alpha3"
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
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1alpha3.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha3.Clientset)
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

// EnvoyFilters mocks base method.
func (m *MockClientset) EnvoyFilters() v1alpha3.EnvoyFilterClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvoyFilters")
	ret0, _ := ret[0].(v1alpha3.EnvoyFilterClient)
	return ret0
}

// EnvoyFilters indicates an expected call of EnvoyFilters.
func (mr *MockClientsetMockRecorder) EnvoyFilters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvoyFilters", reflect.TypeOf((*MockClientset)(nil).EnvoyFilters))
}

// MockEnvoyFilterReader is a mock of EnvoyFilterReader interface.
type MockEnvoyFilterReader struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterReaderMockRecorder
}

// MockEnvoyFilterReaderMockRecorder is the mock recorder for MockEnvoyFilterReader.
type MockEnvoyFilterReaderMockRecorder struct {
	mock *MockEnvoyFilterReader
}

// NewMockEnvoyFilterReader creates a new mock instance.
func NewMockEnvoyFilterReader(ctrl *gomock.Controller) *MockEnvoyFilterReader {
	mock := &MockEnvoyFilterReader{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterReader) EXPECT() *MockEnvoyFilterReaderMockRecorder {
	return m.recorder
}

// GetEnvoyFilter mocks base method.
func (m *MockEnvoyFilterReader) GetEnvoyFilter(ctx context.Context, key client.ObjectKey) (*v1alpha30.EnvoyFilter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvoyFilter", ctx, key)
	ret0, _ := ret[0].(*v1alpha30.EnvoyFilter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvoyFilter indicates an expected call of GetEnvoyFilter.
func (mr *MockEnvoyFilterReaderMockRecorder) GetEnvoyFilter(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterReader)(nil).GetEnvoyFilter), ctx, key)
}

// ListEnvoyFilter mocks base method.
func (m *MockEnvoyFilterReader) ListEnvoyFilter(ctx context.Context, opts ...client.ListOption) (*v1alpha30.EnvoyFilterList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnvoyFilter", varargs...)
	ret0, _ := ret[0].(*v1alpha30.EnvoyFilterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvoyFilter indicates an expected call of ListEnvoyFilter.
func (mr *MockEnvoyFilterReaderMockRecorder) ListEnvoyFilter(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterReader)(nil).ListEnvoyFilter), varargs...)
}

// MockEnvoyFilterWriter is a mock of EnvoyFilterWriter interface.
type MockEnvoyFilterWriter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterWriterMockRecorder
}

// MockEnvoyFilterWriterMockRecorder is the mock recorder for MockEnvoyFilterWriter.
type MockEnvoyFilterWriterMockRecorder struct {
	mock *MockEnvoyFilterWriter
}

// NewMockEnvoyFilterWriter creates a new mock instance.
func NewMockEnvoyFilterWriter(ctrl *gomock.Controller) *MockEnvoyFilterWriter {
	mock := &MockEnvoyFilterWriter{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterWriter) EXPECT() *MockEnvoyFilterWriterMockRecorder {
	return m.recorder
}

// CreateEnvoyFilter mocks base method.
func (m *MockEnvoyFilterWriter) CreateEnvoyFilter(ctx context.Context, obj *v1alpha30.EnvoyFilter, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEnvoyFilter indicates an expected call of CreateEnvoyFilter.
func (mr *MockEnvoyFilterWriterMockRecorder) CreateEnvoyFilter(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterWriter)(nil).CreateEnvoyFilter), varargs...)
}

// DeleteAllOfEnvoyFilter mocks base method.
func (m *MockEnvoyFilterWriter) DeleteAllOfEnvoyFilter(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfEnvoyFilter indicates an expected call of DeleteAllOfEnvoyFilter.
func (mr *MockEnvoyFilterWriterMockRecorder) DeleteAllOfEnvoyFilter(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterWriter)(nil).DeleteAllOfEnvoyFilter), varargs...)
}

// DeleteEnvoyFilter mocks base method.
func (m *MockEnvoyFilterWriter) DeleteEnvoyFilter(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvoyFilter indicates an expected call of DeleteEnvoyFilter.
func (mr *MockEnvoyFilterWriterMockRecorder) DeleteEnvoyFilter(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterWriter)(nil).DeleteEnvoyFilter), varargs...)
}

// PatchEnvoyFilter mocks base method.
func (m *MockEnvoyFilterWriter) PatchEnvoyFilter(ctx context.Context, obj *v1alpha30.EnvoyFilter, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchEnvoyFilter indicates an expected call of PatchEnvoyFilter.
func (mr *MockEnvoyFilterWriterMockRecorder) PatchEnvoyFilter(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterWriter)(nil).PatchEnvoyFilter), varargs...)
}

// UpdateEnvoyFilter mocks base method.
func (m *MockEnvoyFilterWriter) UpdateEnvoyFilter(ctx context.Context, obj *v1alpha30.EnvoyFilter, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEnvoyFilter indicates an expected call of UpdateEnvoyFilter.
func (mr *MockEnvoyFilterWriterMockRecorder) UpdateEnvoyFilter(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterWriter)(nil).UpdateEnvoyFilter), varargs...)
}

// UpsertEnvoyFilter mocks base method.
func (m *MockEnvoyFilterWriter) UpsertEnvoyFilter(ctx context.Context, obj *v1alpha30.EnvoyFilter, transitionFuncs ...v1alpha3.EnvoyFilterTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertEnvoyFilter indicates an expected call of UpsertEnvoyFilter.
func (mr *MockEnvoyFilterWriterMockRecorder) UpsertEnvoyFilter(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterWriter)(nil).UpsertEnvoyFilter), varargs...)
}

// MockEnvoyFilterStatusWriter is a mock of EnvoyFilterStatusWriter interface.
type MockEnvoyFilterStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterStatusWriterMockRecorder
}

// MockEnvoyFilterStatusWriterMockRecorder is the mock recorder for MockEnvoyFilterStatusWriter.
type MockEnvoyFilterStatusWriterMockRecorder struct {
	mock *MockEnvoyFilterStatusWriter
}

// NewMockEnvoyFilterStatusWriter creates a new mock instance.
func NewMockEnvoyFilterStatusWriter(ctrl *gomock.Controller) *MockEnvoyFilterStatusWriter {
	mock := &MockEnvoyFilterStatusWriter{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterStatusWriter) EXPECT() *MockEnvoyFilterStatusWriterMockRecorder {
	return m.recorder
}

// PatchEnvoyFilterStatus mocks base method.
func (m *MockEnvoyFilterStatusWriter) PatchEnvoyFilterStatus(ctx context.Context, obj *v1alpha30.EnvoyFilter, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchEnvoyFilterStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchEnvoyFilterStatus indicates an expected call of PatchEnvoyFilterStatus.
func (mr *MockEnvoyFilterStatusWriterMockRecorder) PatchEnvoyFilterStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchEnvoyFilterStatus", reflect.TypeOf((*MockEnvoyFilterStatusWriter)(nil).PatchEnvoyFilterStatus), varargs...)
}

// UpdateEnvoyFilterStatus mocks base method.
func (m *MockEnvoyFilterStatusWriter) UpdateEnvoyFilterStatus(ctx context.Context, obj *v1alpha30.EnvoyFilter, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEnvoyFilterStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEnvoyFilterStatus indicates an expected call of UpdateEnvoyFilterStatus.
func (mr *MockEnvoyFilterStatusWriterMockRecorder) UpdateEnvoyFilterStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvoyFilterStatus", reflect.TypeOf((*MockEnvoyFilterStatusWriter)(nil).UpdateEnvoyFilterStatus), varargs...)
}

// MockEnvoyFilterClient is a mock of EnvoyFilterClient interface.
type MockEnvoyFilterClient struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterClientMockRecorder
}

// MockEnvoyFilterClientMockRecorder is the mock recorder for MockEnvoyFilterClient.
type MockEnvoyFilterClientMockRecorder struct {
	mock *MockEnvoyFilterClient
}

// NewMockEnvoyFilterClient creates a new mock instance.
func NewMockEnvoyFilterClient(ctrl *gomock.Controller) *MockEnvoyFilterClient {
	mock := &MockEnvoyFilterClient{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterClient) EXPECT() *MockEnvoyFilterClientMockRecorder {
	return m.recorder
}

// CreateEnvoyFilter mocks base method.
func (m *MockEnvoyFilterClient) CreateEnvoyFilter(ctx context.Context, obj *v1alpha30.EnvoyFilter, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEnvoyFilter indicates an expected call of CreateEnvoyFilter.
func (mr *MockEnvoyFilterClientMockRecorder) CreateEnvoyFilter(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterClient)(nil).CreateEnvoyFilter), varargs...)
}

// DeleteAllOfEnvoyFilter mocks base method.
func (m *MockEnvoyFilterClient) DeleteAllOfEnvoyFilter(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfEnvoyFilter indicates an expected call of DeleteAllOfEnvoyFilter.
func (mr *MockEnvoyFilterClientMockRecorder) DeleteAllOfEnvoyFilter(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterClient)(nil).DeleteAllOfEnvoyFilter), varargs...)
}

// DeleteEnvoyFilter mocks base method.
func (m *MockEnvoyFilterClient) DeleteEnvoyFilter(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvoyFilter indicates an expected call of DeleteEnvoyFilter.
func (mr *MockEnvoyFilterClientMockRecorder) DeleteEnvoyFilter(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterClient)(nil).DeleteEnvoyFilter), varargs...)
}

// GetEnvoyFilter mocks base method.
func (m *MockEnvoyFilterClient) GetEnvoyFilter(ctx context.Context, key client.ObjectKey) (*v1alpha30.EnvoyFilter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvoyFilter", ctx, key)
	ret0, _ := ret[0].(*v1alpha30.EnvoyFilter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvoyFilter indicates an expected call of GetEnvoyFilter.
func (mr *MockEnvoyFilterClientMockRecorder) GetEnvoyFilter(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterClient)(nil).GetEnvoyFilter), ctx, key)
}

// ListEnvoyFilter mocks base method.
func (m *MockEnvoyFilterClient) ListEnvoyFilter(ctx context.Context, opts ...client.ListOption) (*v1alpha30.EnvoyFilterList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnvoyFilter", varargs...)
	ret0, _ := ret[0].(*v1alpha30.EnvoyFilterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvoyFilter indicates an expected call of ListEnvoyFilter.
func (mr *MockEnvoyFilterClientMockRecorder) ListEnvoyFilter(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterClient)(nil).ListEnvoyFilter), varargs...)
}

// PatchEnvoyFilter mocks base method.
func (m *MockEnvoyFilterClient) PatchEnvoyFilter(ctx context.Context, obj *v1alpha30.EnvoyFilter, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchEnvoyFilter indicates an expected call of PatchEnvoyFilter.
func (mr *MockEnvoyFilterClientMockRecorder) PatchEnvoyFilter(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterClient)(nil).PatchEnvoyFilter), varargs...)
}

// PatchEnvoyFilterStatus mocks base method.
func (m *MockEnvoyFilterClient) PatchEnvoyFilterStatus(ctx context.Context, obj *v1alpha30.EnvoyFilter, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchEnvoyFilterStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchEnvoyFilterStatus indicates an expected call of PatchEnvoyFilterStatus.
func (mr *MockEnvoyFilterClientMockRecorder) PatchEnvoyFilterStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchEnvoyFilterStatus", reflect.TypeOf((*MockEnvoyFilterClient)(nil).PatchEnvoyFilterStatus), varargs...)
}

// UpdateEnvoyFilter mocks base method.
func (m *MockEnvoyFilterClient) UpdateEnvoyFilter(ctx context.Context, obj *v1alpha30.EnvoyFilter, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEnvoyFilter indicates an expected call of UpdateEnvoyFilter.
func (mr *MockEnvoyFilterClientMockRecorder) UpdateEnvoyFilter(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterClient)(nil).UpdateEnvoyFilter), varargs...)
}

// UpdateEnvoyFilterStatus mocks base method.
func (m *MockEnvoyFilterClient) UpdateEnvoyFilterStatus(ctx context.Context, obj *v1alpha30.EnvoyFilter, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEnvoyFilterStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEnvoyFilterStatus indicates an expected call of UpdateEnvoyFilterStatus.
func (mr *MockEnvoyFilterClientMockRecorder) UpdateEnvoyFilterStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvoyFilterStatus", reflect.TypeOf((*MockEnvoyFilterClient)(nil).UpdateEnvoyFilterStatus), varargs...)
}

// UpsertEnvoyFilter mocks base method.
func (m *MockEnvoyFilterClient) UpsertEnvoyFilter(ctx context.Context, obj *v1alpha30.EnvoyFilter, transitionFuncs ...v1alpha3.EnvoyFilterTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertEnvoyFilter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertEnvoyFilter indicates an expected call of UpsertEnvoyFilter.
func (mr *MockEnvoyFilterClientMockRecorder) UpsertEnvoyFilter(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterClient)(nil).UpsertEnvoyFilter), varargs...)
}

// MockMulticlusterEnvoyFilterClient is a mock of MulticlusterEnvoyFilterClient interface.
type MockMulticlusterEnvoyFilterClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterEnvoyFilterClientMockRecorder
}

// MockMulticlusterEnvoyFilterClientMockRecorder is the mock recorder for MockMulticlusterEnvoyFilterClient.
type MockMulticlusterEnvoyFilterClientMockRecorder struct {
	mock *MockMulticlusterEnvoyFilterClient
}

// NewMockMulticlusterEnvoyFilterClient creates a new mock instance.
func NewMockMulticlusterEnvoyFilterClient(ctrl *gomock.Controller) *MockMulticlusterEnvoyFilterClient {
	mock := &MockMulticlusterEnvoyFilterClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterEnvoyFilterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterEnvoyFilterClient) EXPECT() *MockMulticlusterEnvoyFilterClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterEnvoyFilterClient) Cluster(cluster string) (v1alpha3.EnvoyFilterClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha3.EnvoyFilterClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterEnvoyFilterClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterEnvoyFilterClient)(nil).Cluster), cluster)
}
