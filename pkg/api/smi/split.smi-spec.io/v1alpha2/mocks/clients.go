// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha2 is a generated GoMock package.
package mock_v1alpha2

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"
	v1alpha20 "github.com/solo-io/external-apis/pkg/api/smi/split.smi-spec.io/v1alpha2"
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
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1alpha20.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha20.Clientset)
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

// TrafficSplits mocks base method.
func (m *MockClientset) TrafficSplits() v1alpha20.TrafficSplitClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficSplits")
	ret0, _ := ret[0].(v1alpha20.TrafficSplitClient)
	return ret0
}

// TrafficSplits indicates an expected call of TrafficSplits.
func (mr *MockClientsetMockRecorder) TrafficSplits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficSplits", reflect.TypeOf((*MockClientset)(nil).TrafficSplits))
}

// MockTrafficSplitReader is a mock of TrafficSplitReader interface.
type MockTrafficSplitReader struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitReaderMockRecorder
}

// MockTrafficSplitReaderMockRecorder is the mock recorder for MockTrafficSplitReader.
type MockTrafficSplitReaderMockRecorder struct {
	mock *MockTrafficSplitReader
}

// NewMockTrafficSplitReader creates a new mock instance.
func NewMockTrafficSplitReader(ctrl *gomock.Controller) *MockTrafficSplitReader {
	mock := &MockTrafficSplitReader{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitReader) EXPECT() *MockTrafficSplitReaderMockRecorder {
	return m.recorder
}

// GetTrafficSplit mocks base method.
func (m *MockTrafficSplitReader) GetTrafficSplit(ctx context.Context, key client.ObjectKey) (*v1alpha2.TrafficSplit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrafficSplit", ctx, key)
	ret0, _ := ret[0].(*v1alpha2.TrafficSplit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrafficSplit indicates an expected call of GetTrafficSplit.
func (mr *MockTrafficSplitReaderMockRecorder) GetTrafficSplit(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrafficSplit", reflect.TypeOf((*MockTrafficSplitReader)(nil).GetTrafficSplit), ctx, key)
}

// ListTrafficSplit mocks base method.
func (m *MockTrafficSplitReader) ListTrafficSplit(ctx context.Context, opts ...client.ListOption) (*v1alpha2.TrafficSplitList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTrafficSplit", varargs...)
	ret0, _ := ret[0].(*v1alpha2.TrafficSplitList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrafficSplit indicates an expected call of ListTrafficSplit.
func (mr *MockTrafficSplitReaderMockRecorder) ListTrafficSplit(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficSplit", reflect.TypeOf((*MockTrafficSplitReader)(nil).ListTrafficSplit), varargs...)
}

// MockTrafficSplitWriter is a mock of TrafficSplitWriter interface.
type MockTrafficSplitWriter struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitWriterMockRecorder
}

// MockTrafficSplitWriterMockRecorder is the mock recorder for MockTrafficSplitWriter.
type MockTrafficSplitWriterMockRecorder struct {
	mock *MockTrafficSplitWriter
}

// NewMockTrafficSplitWriter creates a new mock instance.
func NewMockTrafficSplitWriter(ctrl *gomock.Controller) *MockTrafficSplitWriter {
	mock := &MockTrafficSplitWriter{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitWriter) EXPECT() *MockTrafficSplitWriterMockRecorder {
	return m.recorder
}

// CreateTrafficSplit mocks base method.
func (m *MockTrafficSplitWriter) CreateTrafficSplit(ctx context.Context, obj *v1alpha2.TrafficSplit, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrafficSplit indicates an expected call of CreateTrafficSplit.
func (mr *MockTrafficSplitWriterMockRecorder) CreateTrafficSplit(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrafficSplit", reflect.TypeOf((*MockTrafficSplitWriter)(nil).CreateTrafficSplit), varargs...)
}

// DeleteAllOfTrafficSplit mocks base method.
func (m *MockTrafficSplitWriter) DeleteAllOfTrafficSplit(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfTrafficSplit indicates an expected call of DeleteAllOfTrafficSplit.
func (mr *MockTrafficSplitWriterMockRecorder) DeleteAllOfTrafficSplit(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfTrafficSplit", reflect.TypeOf((*MockTrafficSplitWriter)(nil).DeleteAllOfTrafficSplit), varargs...)
}

// DeleteTrafficSplit mocks base method.
func (m *MockTrafficSplitWriter) DeleteTrafficSplit(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrafficSplit indicates an expected call of DeleteTrafficSplit.
func (mr *MockTrafficSplitWriterMockRecorder) DeleteTrafficSplit(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrafficSplit", reflect.TypeOf((*MockTrafficSplitWriter)(nil).DeleteTrafficSplit), varargs...)
}

// PatchTrafficSplit mocks base method.
func (m *MockTrafficSplitWriter) PatchTrafficSplit(ctx context.Context, obj *v1alpha2.TrafficSplit, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchTrafficSplit indicates an expected call of PatchTrafficSplit.
func (mr *MockTrafficSplitWriterMockRecorder) PatchTrafficSplit(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchTrafficSplit", reflect.TypeOf((*MockTrafficSplitWriter)(nil).PatchTrafficSplit), varargs...)
}

// UpdateTrafficSplit mocks base method.
func (m *MockTrafficSplitWriter) UpdateTrafficSplit(ctx context.Context, obj *v1alpha2.TrafficSplit, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrafficSplit indicates an expected call of UpdateTrafficSplit.
func (mr *MockTrafficSplitWriterMockRecorder) UpdateTrafficSplit(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrafficSplit", reflect.TypeOf((*MockTrafficSplitWriter)(nil).UpdateTrafficSplit), varargs...)
}

// UpsertTrafficSplit mocks base method.
func (m *MockTrafficSplitWriter) UpsertTrafficSplit(ctx context.Context, obj *v1alpha2.TrafficSplit, transitionFuncs ...v1alpha20.TrafficSplitTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertTrafficSplit indicates an expected call of UpsertTrafficSplit.
func (mr *MockTrafficSplitWriterMockRecorder) UpsertTrafficSplit(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTrafficSplit", reflect.TypeOf((*MockTrafficSplitWriter)(nil).UpsertTrafficSplit), varargs...)
}

// MockTrafficSplitStatusWriter is a mock of TrafficSplitStatusWriter interface.
type MockTrafficSplitStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitStatusWriterMockRecorder
}

// MockTrafficSplitStatusWriterMockRecorder is the mock recorder for MockTrafficSplitStatusWriter.
type MockTrafficSplitStatusWriterMockRecorder struct {
	mock *MockTrafficSplitStatusWriter
}

// NewMockTrafficSplitStatusWriter creates a new mock instance.
func NewMockTrafficSplitStatusWriter(ctrl *gomock.Controller) *MockTrafficSplitStatusWriter {
	mock := &MockTrafficSplitStatusWriter{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitStatusWriter) EXPECT() *MockTrafficSplitStatusWriterMockRecorder {
	return m.recorder
}

// PatchTrafficSplitStatus mocks base method.
func (m *MockTrafficSplitStatusWriter) PatchTrafficSplitStatus(ctx context.Context, obj *v1alpha2.TrafficSplit, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchTrafficSplitStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchTrafficSplitStatus indicates an expected call of PatchTrafficSplitStatus.
func (mr *MockTrafficSplitStatusWriterMockRecorder) PatchTrafficSplitStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchTrafficSplitStatus", reflect.TypeOf((*MockTrafficSplitStatusWriter)(nil).PatchTrafficSplitStatus), varargs...)
}

// UpdateTrafficSplitStatus mocks base method.
func (m *MockTrafficSplitStatusWriter) UpdateTrafficSplitStatus(ctx context.Context, obj *v1alpha2.TrafficSplit, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTrafficSplitStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrafficSplitStatus indicates an expected call of UpdateTrafficSplitStatus.
func (mr *MockTrafficSplitStatusWriterMockRecorder) UpdateTrafficSplitStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrafficSplitStatus", reflect.TypeOf((*MockTrafficSplitStatusWriter)(nil).UpdateTrafficSplitStatus), varargs...)
}

// MockTrafficSplitClient is a mock of TrafficSplitClient interface.
type MockTrafficSplitClient struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitClientMockRecorder
}

// MockTrafficSplitClientMockRecorder is the mock recorder for MockTrafficSplitClient.
type MockTrafficSplitClientMockRecorder struct {
	mock *MockTrafficSplitClient
}

// NewMockTrafficSplitClient creates a new mock instance.
func NewMockTrafficSplitClient(ctrl *gomock.Controller) *MockTrafficSplitClient {
	mock := &MockTrafficSplitClient{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficSplitClient) EXPECT() *MockTrafficSplitClientMockRecorder {
	return m.recorder
}

// CreateTrafficSplit mocks base method.
func (m *MockTrafficSplitClient) CreateTrafficSplit(ctx context.Context, obj *v1alpha2.TrafficSplit, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrafficSplit indicates an expected call of CreateTrafficSplit.
func (mr *MockTrafficSplitClientMockRecorder) CreateTrafficSplit(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrafficSplit", reflect.TypeOf((*MockTrafficSplitClient)(nil).CreateTrafficSplit), varargs...)
}

// DeleteAllOfTrafficSplit mocks base method.
func (m *MockTrafficSplitClient) DeleteAllOfTrafficSplit(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfTrafficSplit indicates an expected call of DeleteAllOfTrafficSplit.
func (mr *MockTrafficSplitClientMockRecorder) DeleteAllOfTrafficSplit(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfTrafficSplit", reflect.TypeOf((*MockTrafficSplitClient)(nil).DeleteAllOfTrafficSplit), varargs...)
}

// DeleteTrafficSplit mocks base method.
func (m *MockTrafficSplitClient) DeleteTrafficSplit(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrafficSplit indicates an expected call of DeleteTrafficSplit.
func (mr *MockTrafficSplitClientMockRecorder) DeleteTrafficSplit(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrafficSplit", reflect.TypeOf((*MockTrafficSplitClient)(nil).DeleteTrafficSplit), varargs...)
}

// GetTrafficSplit mocks base method.
func (m *MockTrafficSplitClient) GetTrafficSplit(ctx context.Context, key client.ObjectKey) (*v1alpha2.TrafficSplit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrafficSplit", ctx, key)
	ret0, _ := ret[0].(*v1alpha2.TrafficSplit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrafficSplit indicates an expected call of GetTrafficSplit.
func (mr *MockTrafficSplitClientMockRecorder) GetTrafficSplit(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrafficSplit", reflect.TypeOf((*MockTrafficSplitClient)(nil).GetTrafficSplit), ctx, key)
}

// ListTrafficSplit mocks base method.
func (m *MockTrafficSplitClient) ListTrafficSplit(ctx context.Context, opts ...client.ListOption) (*v1alpha2.TrafficSplitList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTrafficSplit", varargs...)
	ret0, _ := ret[0].(*v1alpha2.TrafficSplitList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrafficSplit indicates an expected call of ListTrafficSplit.
func (mr *MockTrafficSplitClientMockRecorder) ListTrafficSplit(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficSplit", reflect.TypeOf((*MockTrafficSplitClient)(nil).ListTrafficSplit), varargs...)
}

// PatchTrafficSplit mocks base method.
func (m *MockTrafficSplitClient) PatchTrafficSplit(ctx context.Context, obj *v1alpha2.TrafficSplit, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchTrafficSplit indicates an expected call of PatchTrafficSplit.
func (mr *MockTrafficSplitClientMockRecorder) PatchTrafficSplit(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchTrafficSplit", reflect.TypeOf((*MockTrafficSplitClient)(nil).PatchTrafficSplit), varargs...)
}

// PatchTrafficSplitStatus mocks base method.
func (m *MockTrafficSplitClient) PatchTrafficSplitStatus(ctx context.Context, obj *v1alpha2.TrafficSplit, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchTrafficSplitStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchTrafficSplitStatus indicates an expected call of PatchTrafficSplitStatus.
func (mr *MockTrafficSplitClientMockRecorder) PatchTrafficSplitStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchTrafficSplitStatus", reflect.TypeOf((*MockTrafficSplitClient)(nil).PatchTrafficSplitStatus), varargs...)
}

// UpdateTrafficSplit mocks base method.
func (m *MockTrafficSplitClient) UpdateTrafficSplit(ctx context.Context, obj *v1alpha2.TrafficSplit, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrafficSplit indicates an expected call of UpdateTrafficSplit.
func (mr *MockTrafficSplitClientMockRecorder) UpdateTrafficSplit(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrafficSplit", reflect.TypeOf((*MockTrafficSplitClient)(nil).UpdateTrafficSplit), varargs...)
}

// UpdateTrafficSplitStatus mocks base method.
func (m *MockTrafficSplitClient) UpdateTrafficSplitStatus(ctx context.Context, obj *v1alpha2.TrafficSplit, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTrafficSplitStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrafficSplitStatus indicates an expected call of UpdateTrafficSplitStatus.
func (mr *MockTrafficSplitClientMockRecorder) UpdateTrafficSplitStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrafficSplitStatus", reflect.TypeOf((*MockTrafficSplitClient)(nil).UpdateTrafficSplitStatus), varargs...)
}

// UpsertTrafficSplit mocks base method.
func (m *MockTrafficSplitClient) UpsertTrafficSplit(ctx context.Context, obj *v1alpha2.TrafficSplit, transitionFuncs ...v1alpha20.TrafficSplitTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertTrafficSplit", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertTrafficSplit indicates an expected call of UpsertTrafficSplit.
func (mr *MockTrafficSplitClientMockRecorder) UpsertTrafficSplit(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTrafficSplit", reflect.TypeOf((*MockTrafficSplitClient)(nil).UpsertTrafficSplit), varargs...)
}

// MockMulticlusterTrafficSplitClient is a mock of MulticlusterTrafficSplitClient interface.
type MockMulticlusterTrafficSplitClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTrafficSplitClientMockRecorder
}

// MockMulticlusterTrafficSplitClientMockRecorder is the mock recorder for MockMulticlusterTrafficSplitClient.
type MockMulticlusterTrafficSplitClientMockRecorder struct {
	mock *MockMulticlusterTrafficSplitClient
}

// NewMockMulticlusterTrafficSplitClient creates a new mock instance.
func NewMockMulticlusterTrafficSplitClient(ctrl *gomock.Controller) *MockMulticlusterTrafficSplitClient {
	mock := &MockMulticlusterTrafficSplitClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTrafficSplitClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterTrafficSplitClient) EXPECT() *MockMulticlusterTrafficSplitClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockMulticlusterTrafficSplitClient) Cluster(cluster string) (v1alpha20.TrafficSplitClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha20.TrafficSplitClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockMulticlusterTrafficSplitClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterTrafficSplitClient)(nil).Cluster), cluster)
}
