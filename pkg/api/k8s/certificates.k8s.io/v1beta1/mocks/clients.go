// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1beta1 is a generated GoMock package.
package mock_v1beta1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/solo-io/external-apis/pkg/api/k8s/certificates.k8s.io/v1beta1"
	v1beta10 "k8s.io/api/certificates/v1beta1"
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
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1beta1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1beta1.Clientset)
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

// CertificateSigningRequests mocks base method
func (m *MockClientset) CertificateSigningRequests() v1beta1.CertificateSigningRequestClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificateSigningRequests")
	ret0, _ := ret[0].(v1beta1.CertificateSigningRequestClient)
	return ret0
}

// CertificateSigningRequests indicates an expected call of CertificateSigningRequests
func (mr *MockClientsetMockRecorder) CertificateSigningRequests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificateSigningRequests", reflect.TypeOf((*MockClientset)(nil).CertificateSigningRequests))
}

// MockCertificateSigningRequestReader is a mock of CertificateSigningRequestReader interface
type MockCertificateSigningRequestReader struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestReaderMockRecorder
}

// MockCertificateSigningRequestReaderMockRecorder is the mock recorder for MockCertificateSigningRequestReader
type MockCertificateSigningRequestReaderMockRecorder struct {
	mock *MockCertificateSigningRequestReader
}

// NewMockCertificateSigningRequestReader creates a new mock instance
func NewMockCertificateSigningRequestReader(ctrl *gomock.Controller) *MockCertificateSigningRequestReader {
	mock := &MockCertificateSigningRequestReader{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateSigningRequestReader) EXPECT() *MockCertificateSigningRequestReaderMockRecorder {
	return m.recorder
}

// GetCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestReader) GetCertificateSigningRequest(ctx context.Context, key client.ObjectKey) (*v1beta10.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateSigningRequest", ctx, key)
	ret0, _ := ret[0].(*v1beta10.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificateSigningRequest indicates an expected call of GetCertificateSigningRequest
func (mr *MockCertificateSigningRequestReaderMockRecorder) GetCertificateSigningRequest(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestReader)(nil).GetCertificateSigningRequest), ctx, key)
}

// ListCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestReader) ListCertificateSigningRequest(ctx context.Context, opts ...client.ListOption) (*v1beta10.CertificateSigningRequestList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(*v1beta10.CertificateSigningRequestList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificateSigningRequest indicates an expected call of ListCertificateSigningRequest
func (mr *MockCertificateSigningRequestReaderMockRecorder) ListCertificateSigningRequest(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestReader)(nil).ListCertificateSigningRequest), varargs...)
}

// MockCertificateSigningRequestWriter is a mock of CertificateSigningRequestWriter interface
type MockCertificateSigningRequestWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestWriterMockRecorder
}

// MockCertificateSigningRequestWriterMockRecorder is the mock recorder for MockCertificateSigningRequestWriter
type MockCertificateSigningRequestWriterMockRecorder struct {
	mock *MockCertificateSigningRequestWriter
}

// NewMockCertificateSigningRequestWriter creates a new mock instance
func NewMockCertificateSigningRequestWriter(ctrl *gomock.Controller) *MockCertificateSigningRequestWriter {
	mock := &MockCertificateSigningRequestWriter{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateSigningRequestWriter) EXPECT() *MockCertificateSigningRequestWriterMockRecorder {
	return m.recorder
}

// CreateCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestWriter) CreateCertificateSigningRequest(ctx context.Context, obj *v1beta10.CertificateSigningRequest, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCertificateSigningRequest indicates an expected call of CreateCertificateSigningRequest
func (mr *MockCertificateSigningRequestWriterMockRecorder) CreateCertificateSigningRequest(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestWriter)(nil).CreateCertificateSigningRequest), varargs...)
}

// DeleteCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestWriter) DeleteCertificateSigningRequest(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCertificateSigningRequest indicates an expected call of DeleteCertificateSigningRequest
func (mr *MockCertificateSigningRequestWriterMockRecorder) DeleteCertificateSigningRequest(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestWriter)(nil).DeleteCertificateSigningRequest), varargs...)
}

// UpdateCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestWriter) UpdateCertificateSigningRequest(ctx context.Context, obj *v1beta10.CertificateSigningRequest, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCertificateSigningRequest indicates an expected call of UpdateCertificateSigningRequest
func (mr *MockCertificateSigningRequestWriterMockRecorder) UpdateCertificateSigningRequest(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestWriter)(nil).UpdateCertificateSigningRequest), varargs...)
}

// PatchCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestWriter) PatchCertificateSigningRequest(ctx context.Context, obj *v1beta10.CertificateSigningRequest, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCertificateSigningRequest indicates an expected call of PatchCertificateSigningRequest
func (mr *MockCertificateSigningRequestWriterMockRecorder) PatchCertificateSigningRequest(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestWriter)(nil).PatchCertificateSigningRequest), varargs...)
}

// DeleteAllOfCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestWriter) DeleteAllOfCertificateSigningRequest(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfCertificateSigningRequest indicates an expected call of DeleteAllOfCertificateSigningRequest
func (mr *MockCertificateSigningRequestWriterMockRecorder) DeleteAllOfCertificateSigningRequest(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestWriter)(nil).DeleteAllOfCertificateSigningRequest), varargs...)
}

// UpsertCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestWriter) UpsertCertificateSigningRequest(ctx context.Context, obj *v1beta10.CertificateSigningRequest, transitionFuncs ...v1beta1.CertificateSigningRequestTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertCertificateSigningRequest indicates an expected call of UpsertCertificateSigningRequest
func (mr *MockCertificateSigningRequestWriterMockRecorder) UpsertCertificateSigningRequest(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestWriter)(nil).UpsertCertificateSigningRequest), varargs...)
}

// MockCertificateSigningRequestStatusWriter is a mock of CertificateSigningRequestStatusWriter interface
type MockCertificateSigningRequestStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestStatusWriterMockRecorder
}

// MockCertificateSigningRequestStatusWriterMockRecorder is the mock recorder for MockCertificateSigningRequestStatusWriter
type MockCertificateSigningRequestStatusWriterMockRecorder struct {
	mock *MockCertificateSigningRequestStatusWriter
}

// NewMockCertificateSigningRequestStatusWriter creates a new mock instance
func NewMockCertificateSigningRequestStatusWriter(ctrl *gomock.Controller) *MockCertificateSigningRequestStatusWriter {
	mock := &MockCertificateSigningRequestStatusWriter{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateSigningRequestStatusWriter) EXPECT() *MockCertificateSigningRequestStatusWriterMockRecorder {
	return m.recorder
}

// UpdateCertificateSigningRequestStatus mocks base method
func (m *MockCertificateSigningRequestStatusWriter) UpdateCertificateSigningRequestStatus(ctx context.Context, obj *v1beta10.CertificateSigningRequest, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCertificateSigningRequestStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCertificateSigningRequestStatus indicates an expected call of UpdateCertificateSigningRequestStatus
func (mr *MockCertificateSigningRequestStatusWriterMockRecorder) UpdateCertificateSigningRequestStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateSigningRequestStatus", reflect.TypeOf((*MockCertificateSigningRequestStatusWriter)(nil).UpdateCertificateSigningRequestStatus), varargs...)
}

// PatchCertificateSigningRequestStatus mocks base method
func (m *MockCertificateSigningRequestStatusWriter) PatchCertificateSigningRequestStatus(ctx context.Context, obj *v1beta10.CertificateSigningRequest, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCertificateSigningRequestStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCertificateSigningRequestStatus indicates an expected call of PatchCertificateSigningRequestStatus
func (mr *MockCertificateSigningRequestStatusWriterMockRecorder) PatchCertificateSigningRequestStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCertificateSigningRequestStatus", reflect.TypeOf((*MockCertificateSigningRequestStatusWriter)(nil).PatchCertificateSigningRequestStatus), varargs...)
}

// MockCertificateSigningRequestClient is a mock of CertificateSigningRequestClient interface
type MockCertificateSigningRequestClient struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestClientMockRecorder
}

// MockCertificateSigningRequestClientMockRecorder is the mock recorder for MockCertificateSigningRequestClient
type MockCertificateSigningRequestClientMockRecorder struct {
	mock *MockCertificateSigningRequestClient
}

// NewMockCertificateSigningRequestClient creates a new mock instance
func NewMockCertificateSigningRequestClient(ctrl *gomock.Controller) *MockCertificateSigningRequestClient {
	mock := &MockCertificateSigningRequestClient{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateSigningRequestClient) EXPECT() *MockCertificateSigningRequestClientMockRecorder {
	return m.recorder
}

// GetCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestClient) GetCertificateSigningRequest(ctx context.Context, key client.ObjectKey) (*v1beta10.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateSigningRequest", ctx, key)
	ret0, _ := ret[0].(*v1beta10.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificateSigningRequest indicates an expected call of GetCertificateSigningRequest
func (mr *MockCertificateSigningRequestClientMockRecorder) GetCertificateSigningRequest(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).GetCertificateSigningRequest), ctx, key)
}

// ListCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestClient) ListCertificateSigningRequest(ctx context.Context, opts ...client.ListOption) (*v1beta10.CertificateSigningRequestList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(*v1beta10.CertificateSigningRequestList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificateSigningRequest indicates an expected call of ListCertificateSigningRequest
func (mr *MockCertificateSigningRequestClientMockRecorder) ListCertificateSigningRequest(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).ListCertificateSigningRequest), varargs...)
}

// CreateCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestClient) CreateCertificateSigningRequest(ctx context.Context, obj *v1beta10.CertificateSigningRequest, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCertificateSigningRequest indicates an expected call of CreateCertificateSigningRequest
func (mr *MockCertificateSigningRequestClientMockRecorder) CreateCertificateSigningRequest(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).CreateCertificateSigningRequest), varargs...)
}

// DeleteCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestClient) DeleteCertificateSigningRequest(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCertificateSigningRequest indicates an expected call of DeleteCertificateSigningRequest
func (mr *MockCertificateSigningRequestClientMockRecorder) DeleteCertificateSigningRequest(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).DeleteCertificateSigningRequest), varargs...)
}

// UpdateCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestClient) UpdateCertificateSigningRequest(ctx context.Context, obj *v1beta10.CertificateSigningRequest, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCertificateSigningRequest indicates an expected call of UpdateCertificateSigningRequest
func (mr *MockCertificateSigningRequestClientMockRecorder) UpdateCertificateSigningRequest(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).UpdateCertificateSigningRequest), varargs...)
}

// PatchCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestClient) PatchCertificateSigningRequest(ctx context.Context, obj *v1beta10.CertificateSigningRequest, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCertificateSigningRequest indicates an expected call of PatchCertificateSigningRequest
func (mr *MockCertificateSigningRequestClientMockRecorder) PatchCertificateSigningRequest(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).PatchCertificateSigningRequest), varargs...)
}

// DeleteAllOfCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestClient) DeleteAllOfCertificateSigningRequest(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfCertificateSigningRequest indicates an expected call of DeleteAllOfCertificateSigningRequest
func (mr *MockCertificateSigningRequestClientMockRecorder) DeleteAllOfCertificateSigningRequest(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).DeleteAllOfCertificateSigningRequest), varargs...)
}

// UpsertCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestClient) UpsertCertificateSigningRequest(ctx context.Context, obj *v1beta10.CertificateSigningRequest, transitionFuncs ...v1beta1.CertificateSigningRequestTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertCertificateSigningRequest indicates an expected call of UpsertCertificateSigningRequest
func (mr *MockCertificateSigningRequestClientMockRecorder) UpsertCertificateSigningRequest(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).UpsertCertificateSigningRequest), varargs...)
}

// UpdateCertificateSigningRequestStatus mocks base method
func (m *MockCertificateSigningRequestClient) UpdateCertificateSigningRequestStatus(ctx context.Context, obj *v1beta10.CertificateSigningRequest, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCertificateSigningRequestStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCertificateSigningRequestStatus indicates an expected call of UpdateCertificateSigningRequestStatus
func (mr *MockCertificateSigningRequestClientMockRecorder) UpdateCertificateSigningRequestStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateSigningRequestStatus", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).UpdateCertificateSigningRequestStatus), varargs...)
}

// PatchCertificateSigningRequestStatus mocks base method
func (m *MockCertificateSigningRequestClient) PatchCertificateSigningRequestStatus(ctx context.Context, obj *v1beta10.CertificateSigningRequest, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCertificateSigningRequestStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCertificateSigningRequestStatus indicates an expected call of PatchCertificateSigningRequestStatus
func (mr *MockCertificateSigningRequestClientMockRecorder) PatchCertificateSigningRequestStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCertificateSigningRequestStatus", reflect.TypeOf((*MockCertificateSigningRequestClient)(nil).PatchCertificateSigningRequestStatus), varargs...)
}

// MockMulticlusterCertificateSigningRequestClient is a mock of MulticlusterCertificateSigningRequestClient interface
type MockMulticlusterCertificateSigningRequestClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCertificateSigningRequestClientMockRecorder
}

// MockMulticlusterCertificateSigningRequestClientMockRecorder is the mock recorder for MockMulticlusterCertificateSigningRequestClient
type MockMulticlusterCertificateSigningRequestClientMockRecorder struct {
	mock *MockMulticlusterCertificateSigningRequestClient
}

// NewMockMulticlusterCertificateSigningRequestClient creates a new mock instance
func NewMockMulticlusterCertificateSigningRequestClient(ctrl *gomock.Controller) *MockMulticlusterCertificateSigningRequestClient {
	mock := &MockMulticlusterCertificateSigningRequestClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCertificateSigningRequestClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCertificateSigningRequestClient) EXPECT() *MockMulticlusterCertificateSigningRequestClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterCertificateSigningRequestClient) Cluster(cluster string) (v1beta1.CertificateSigningRequestClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1beta1.CertificateSigningRequestClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterCertificateSigningRequestClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterCertificateSigningRequestClient)(nil).Cluster), cluster)
}
