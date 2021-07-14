// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha3sets is a generated GoMock package.
package mock_v1alpha3sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha3"
	v1alpha3sets "github.com/solo-io/external-apis/pkg/api/smi/split.smi-spec.io/v1alpha3/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockTrafficSplitSet is a mock of TrafficSplitSet interface
type MockTrafficSplitSet struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficSplitSetMockRecorder
}

// MockTrafficSplitSetMockRecorder is the mock recorder for MockTrafficSplitSet
type MockTrafficSplitSetMockRecorder struct {
	mock *MockTrafficSplitSet
}

// NewMockTrafficSplitSet creates a new mock instance
func NewMockTrafficSplitSet(ctrl *gomock.Controller) *MockTrafficSplitSet {
	mock := &MockTrafficSplitSet{ctrl: ctrl}
	mock.recorder = &MockTrafficSplitSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrafficSplitSet) EXPECT() *MockTrafficSplitSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockTrafficSplitSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockTrafficSplitSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockTrafficSplitSet)(nil).Keys))
}

// List mocks base method
func (m *MockTrafficSplitSet) List(filterResource ...func(*v1alpha3.TrafficSplit) bool) []*v1alpha3.TrafficSplit {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha3.TrafficSplit)
	return ret0
}

// List indicates an expected call of List
func (mr *MockTrafficSplitSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTrafficSplitSet)(nil).List), filterResource...)
}

// UnsortedList mocks base method
func (m *MockTrafficSplitSet) UnsortedList(filterResource ...func(*v1alpha3.TrafficSplit) bool) []*v1alpha3.TrafficSplit {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1alpha3.TrafficSplit)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList
func (mr *MockTrafficSplitSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockTrafficSplitSet)(nil).UnsortedList), filterResource...)
}

// Map mocks base method
func (m *MockTrafficSplitSet) Map() map[string]*v1alpha3.TrafficSplit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha3.TrafficSplit)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockTrafficSplitSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockTrafficSplitSet)(nil).Map))
}

// Insert mocks base method
func (m *MockTrafficSplitSet) Insert(trafficSplit ...*v1alpha3.TrafficSplit) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range trafficSplit {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockTrafficSplitSetMockRecorder) Insert(trafficSplit ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTrafficSplitSet)(nil).Insert), trafficSplit...)
}

// Equal mocks base method
func (m *MockTrafficSplitSet) Equal(trafficSplitSet v1alpha3sets.TrafficSplitSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", trafficSplitSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockTrafficSplitSetMockRecorder) Equal(trafficSplitSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockTrafficSplitSet)(nil).Equal), trafficSplitSet)
}

// Has mocks base method
func (m *MockTrafficSplitSet) Has(trafficSplit ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", trafficSplit)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockTrafficSplitSetMockRecorder) Has(trafficSplit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockTrafficSplitSet)(nil).Has), trafficSplit)
}

// Delete mocks base method
func (m *MockTrafficSplitSet) Delete(trafficSplit ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", trafficSplit)
}

// Delete indicates an expected call of Delete
func (mr *MockTrafficSplitSetMockRecorder) Delete(trafficSplit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrafficSplitSet)(nil).Delete), trafficSplit)
}

// Union mocks base method
func (m *MockTrafficSplitSet) Union(set v1alpha3sets.TrafficSplitSet) v1alpha3sets.TrafficSplitSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha3sets.TrafficSplitSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockTrafficSplitSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockTrafficSplitSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockTrafficSplitSet) Difference(set v1alpha3sets.TrafficSplitSet) v1alpha3sets.TrafficSplitSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha3sets.TrafficSplitSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockTrafficSplitSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockTrafficSplitSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockTrafficSplitSet) Intersection(set v1alpha3sets.TrafficSplitSet) v1alpha3sets.TrafficSplitSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha3sets.TrafficSplitSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockTrafficSplitSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockTrafficSplitSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockTrafficSplitSet) Find(id ezkube.ResourceId) (*v1alpha3.TrafficSplit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha3.TrafficSplit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockTrafficSplitSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTrafficSplitSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockTrafficSplitSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockTrafficSplitSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockTrafficSplitSet)(nil).Length))
}

// Generic mocks base method
func (m *MockTrafficSplitSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockTrafficSplitSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockTrafficSplitSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockTrafficSplitSet) Delta(newSet v1alpha3sets.TrafficSplitSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockTrafficSplitSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockTrafficSplitSet)(nil).Delta), newSet)
}

// Clone mocks base method
func (m *MockTrafficSplitSet) Clone() v1alpha3sets.TrafficSplitSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1alpha3sets.TrafficSplitSet)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockTrafficSplitSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockTrafficSplitSet)(nil).Clone))
}
