// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha2sets is a generated GoMock package.
package mock_v1alpha2sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	v1alpha2sets "github.com/solo-io/external-apis/pkg/api/smi/access.smi-spec.io/v1alpha2/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockTrafficTargetSet is a mock of TrafficTargetSet interface.
type MockTrafficTargetSet struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficTargetSetMockRecorder
}

// MockTrafficTargetSetMockRecorder is the mock recorder for MockTrafficTargetSet.
type MockTrafficTargetSetMockRecorder struct {
	mock *MockTrafficTargetSet
}

// NewMockTrafficTargetSet creates a new mock instance.
func NewMockTrafficTargetSet(ctrl *gomock.Controller) *MockTrafficTargetSet {
	mock := &MockTrafficTargetSet{ctrl: ctrl}
	mock.recorder = &MockTrafficTargetSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficTargetSet) EXPECT() *MockTrafficTargetSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockTrafficTargetSet) Clone() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockTrafficTargetSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockTrafficTargetSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockTrafficTargetSet) Delete(trafficTarget ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", trafficTarget)
}

// Delete indicates an expected call of Delete.
func (mr *MockTrafficTargetSetMockRecorder) Delete(trafficTarget interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrafficTargetSet)(nil).Delete), trafficTarget)
}

// Delta mocks base method.
func (m *MockTrafficTargetSet) Delta(newSet v1alpha2sets.TrafficTargetSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockTrafficTargetSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockTrafficTargetSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockTrafficTargetSet) Difference(set v1alpha2sets.TrafficTargetSet) v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockTrafficTargetSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockTrafficTargetSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockTrafficTargetSet) Equal(trafficTargetSet v1alpha2sets.TrafficTargetSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", trafficTargetSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockTrafficTargetSetMockRecorder) Equal(trafficTargetSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockTrafficTargetSet)(nil).Equal), trafficTargetSet)
}

// Find mocks base method.
func (m *MockTrafficTargetSet) Find(id ezkube.ResourceId) (*v1alpha2.TrafficTarget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha2.TrafficTarget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockTrafficTargetSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTrafficTargetSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockTrafficTargetSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockTrafficTargetSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockTrafficTargetSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockTrafficTargetSet) Has(trafficTarget ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", trafficTarget)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockTrafficTargetSetMockRecorder) Has(trafficTarget interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockTrafficTargetSet)(nil).Has), trafficTarget)
}

// Insert mocks base method.
func (m *MockTrafficTargetSet) Insert(trafficTarget ...*v1alpha2.TrafficTarget) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range trafficTarget {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockTrafficTargetSetMockRecorder) Insert(trafficTarget ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTrafficTargetSet)(nil).Insert), trafficTarget...)
}

// Intersection mocks base method.
func (m *MockTrafficTargetSet) Intersection(set v1alpha2sets.TrafficTargetSet) v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockTrafficTargetSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockTrafficTargetSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockTrafficTargetSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockTrafficTargetSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockTrafficTargetSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockTrafficTargetSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockTrafficTargetSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockTrafficTargetSet)(nil).Length))
}

// List mocks base method.
func (m *MockTrafficTargetSet) List(filterResource ...func(*v1alpha2.TrafficTarget) bool) []*v1alpha2.TrafficTarget {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha2.TrafficTarget)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockTrafficTargetSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTrafficTargetSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockTrafficTargetSet) Map() map[string]*v1alpha2.TrafficTarget {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha2.TrafficTarget)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockTrafficTargetSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockTrafficTargetSet)(nil).Map))
}

// Union mocks base method.
func (m *MockTrafficTargetSet) Union(set v1alpha2sets.TrafficTargetSet) v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockTrafficTargetSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockTrafficTargetSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockTrafficTargetSet) UnsortedList(filterResource ...func(*v1alpha2.TrafficTarget) bool) []*v1alpha2.TrafficTarget {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1alpha2.TrafficTarget)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockTrafficTargetSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockTrafficTargetSet)(nil).UnsortedList), filterResource...)
}
