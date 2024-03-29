// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha1sets is a generated GoMock package.
package mock_v1alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1sets "github.com/solo-io/external-apis/pkg/api/istio/extensions.istio.io/v1alpha1/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1alpha1 "istio.io/client-go/pkg/apis/extensions/v1alpha1"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockWasmPluginSet is a mock of WasmPluginSet interface.
type MockWasmPluginSet struct {
	ctrl     *gomock.Controller
	recorder *MockWasmPluginSetMockRecorder
}

// MockWasmPluginSetMockRecorder is the mock recorder for MockWasmPluginSet.
type MockWasmPluginSetMockRecorder struct {
	mock *MockWasmPluginSet
}

// NewMockWasmPluginSet creates a new mock instance.
func NewMockWasmPluginSet(ctrl *gomock.Controller) *MockWasmPluginSet {
	mock := &MockWasmPluginSet{ctrl: ctrl}
	mock.recorder = &MockWasmPluginSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmPluginSet) EXPECT() *MockWasmPluginSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockWasmPluginSet) Clone() v1alpha1sets.WasmPluginSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1alpha1sets.WasmPluginSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockWasmPluginSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockWasmPluginSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockWasmPluginSet) Delete(wasmPlugin ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", wasmPlugin)
}

// Delete indicates an expected call of Delete.
func (mr *MockWasmPluginSetMockRecorder) Delete(wasmPlugin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWasmPluginSet)(nil).Delete), wasmPlugin)
}

// Delta mocks base method.
func (m *MockWasmPluginSet) Delta(newSet v1alpha1sets.WasmPluginSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockWasmPluginSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockWasmPluginSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockWasmPluginSet) Difference(set v1alpha1sets.WasmPluginSet) v1alpha1sets.WasmPluginSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.WasmPluginSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockWasmPluginSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockWasmPluginSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockWasmPluginSet) Equal(wasmPluginSet v1alpha1sets.WasmPluginSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", wasmPluginSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockWasmPluginSetMockRecorder) Equal(wasmPluginSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockWasmPluginSet)(nil).Equal), wasmPluginSet)
}

// Find mocks base method.
func (m *MockWasmPluginSet) Find(id ezkube.ResourceId) (*v1alpha1.WasmPlugin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.WasmPlugin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockWasmPluginSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockWasmPluginSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockWasmPluginSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockWasmPluginSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockWasmPluginSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockWasmPluginSet) Has(wasmPlugin ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", wasmPlugin)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockWasmPluginSetMockRecorder) Has(wasmPlugin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockWasmPluginSet)(nil).Has), wasmPlugin)
}

// Insert mocks base method.
func (m *MockWasmPluginSet) Insert(wasmPlugin ...*v1alpha1.WasmPlugin) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range wasmPlugin {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockWasmPluginSetMockRecorder) Insert(wasmPlugin ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockWasmPluginSet)(nil).Insert), wasmPlugin...)
}

// Intersection mocks base method.
func (m *MockWasmPluginSet) Intersection(set v1alpha1sets.WasmPluginSet) v1alpha1sets.WasmPluginSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.WasmPluginSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockWasmPluginSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockWasmPluginSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockWasmPluginSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockWasmPluginSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockWasmPluginSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockWasmPluginSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockWasmPluginSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockWasmPluginSet)(nil).Length))
}

// List mocks base method.
func (m *MockWasmPluginSet) List(filterResource ...func(*v1alpha1.WasmPlugin) bool) []*v1alpha1.WasmPlugin {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1alpha1.WasmPlugin)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockWasmPluginSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWasmPluginSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockWasmPluginSet) Map() map[string]*v1alpha1.WasmPlugin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.WasmPlugin)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockWasmPluginSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockWasmPluginSet)(nil).Map))
}

// Union mocks base method.
func (m *MockWasmPluginSet) Union(set v1alpha1sets.WasmPluginSet) v1alpha1sets.WasmPluginSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.WasmPluginSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockWasmPluginSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockWasmPluginSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockWasmPluginSet) UnsortedList(filterResource ...func(*v1alpha1.WasmPlugin) bool) []*v1alpha1.WasmPlugin {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1alpha1.WasmPlugin)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockWasmPluginSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockWasmPluginSet)(nil).UnsortedList), filterResource...)
}
