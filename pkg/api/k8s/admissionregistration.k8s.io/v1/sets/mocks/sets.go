// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/admissionregistration.k8s.io/v1/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1 "k8s.io/api/admissionregistration/v1"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockValidatingWebhookConfigurationSet is a mock of ValidatingWebhookConfigurationSet interface.
type MockValidatingWebhookConfigurationSet struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationSetMockRecorder
}

// MockValidatingWebhookConfigurationSetMockRecorder is the mock recorder for MockValidatingWebhookConfigurationSet.
type MockValidatingWebhookConfigurationSetMockRecorder struct {
	mock *MockValidatingWebhookConfigurationSet
}

// NewMockValidatingWebhookConfigurationSet creates a new mock instance.
func NewMockValidatingWebhookConfigurationSet(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationSet {
	mock := &MockValidatingWebhookConfigurationSet{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationSet) EXPECT() *MockValidatingWebhookConfigurationSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Clone() v1sets.ValidatingWebhookConfigurationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.ValidatingWebhookConfigurationSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Delete(validatingWebhookConfiguration ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", validatingWebhookConfiguration)
}

// Delete indicates an expected call of Delete.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Delete(validatingWebhookConfiguration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Delete), validatingWebhookConfiguration)
}

// Delta mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Delta(newSet v1sets.ValidatingWebhookConfigurationSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Difference(set v1sets.ValidatingWebhookConfigurationSet) v1sets.ValidatingWebhookConfigurationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.ValidatingWebhookConfigurationSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Equal(validatingWebhookConfigurationSet v1sets.ValidatingWebhookConfigurationSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", validatingWebhookConfigurationSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Equal(validatingWebhookConfigurationSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Equal), validatingWebhookConfigurationSet)
}

// Find mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Find(id ezkube.ResourceId) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Has(validatingWebhookConfiguration ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", validatingWebhookConfiguration)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Has(validatingWebhookConfiguration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Has), validatingWebhookConfiguration)
}

// Insert mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Insert(validatingWebhookConfiguration ...*v1.ValidatingWebhookConfiguration) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range validatingWebhookConfiguration {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Insert(validatingWebhookConfiguration ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Insert), validatingWebhookConfiguration...)
}

// Intersection mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Intersection(set v1sets.ValidatingWebhookConfigurationSet) v1sets.ValidatingWebhookConfigurationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.ValidatingWebhookConfigurationSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Length))
}

// List mocks base method.
func (m *MockValidatingWebhookConfigurationSet) List(filterResource ...func(*v1.ValidatingWebhookConfiguration) bool) []*v1.ValidatingWebhookConfiguration {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.ValidatingWebhookConfiguration)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Map() map[string]*v1.ValidatingWebhookConfiguration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.ValidatingWebhookConfiguration)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Map))
}

// Union mocks base method.
func (m *MockValidatingWebhookConfigurationSet) Union(set v1sets.ValidatingWebhookConfigurationSet) v1sets.ValidatingWebhookConfigurationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.ValidatingWebhookConfigurationSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockValidatingWebhookConfigurationSet) UnsortedList(filterResource ...func(*v1.ValidatingWebhookConfiguration) bool) []*v1.ValidatingWebhookConfiguration {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.ValidatingWebhookConfiguration)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockValidatingWebhookConfigurationSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockValidatingWebhookConfigurationSet)(nil).UnsortedList), filterResource...)
}
