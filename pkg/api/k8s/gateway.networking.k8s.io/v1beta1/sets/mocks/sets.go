// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1beta1sets is a generated GoMock package.
package mock_v1beta1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1sets "github.com/solo-io/external-apis/pkg/api/k8s/gateway.networking.k8s.io/v1beta1/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
)

// MockGatewaySet is a mock of GatewaySet interface.
type MockGatewaySet struct {
	ctrl     *gomock.Controller
	recorder *MockGatewaySetMockRecorder
}

// MockGatewaySetMockRecorder is the mock recorder for MockGatewaySet.
type MockGatewaySetMockRecorder struct {
	mock *MockGatewaySet
}

// NewMockGatewaySet creates a new mock instance.
func NewMockGatewaySet(ctrl *gomock.Controller) *MockGatewaySet {
	mock := &MockGatewaySet{ctrl: ctrl}
	mock.recorder = &MockGatewaySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewaySet) EXPECT() *MockGatewaySetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockGatewaySet) Clone() v1beta1sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1beta1sets.GatewaySet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockGatewaySetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockGatewaySet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockGatewaySet) Delete(gateway ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", gateway)
}

// Delete indicates an expected call of Delete.
func (mr *MockGatewaySetMockRecorder) Delete(gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGatewaySet)(nil).Delete), gateway)
}

// Delta mocks base method.
func (m *MockGatewaySet) Delta(newSet v1beta1sets.GatewaySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockGatewaySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockGatewaySet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockGatewaySet) Difference(set v1beta1sets.GatewaySet) v1beta1sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1beta1sets.GatewaySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockGatewaySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockGatewaySet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockGatewaySet) Equal(gatewaySet v1beta1sets.GatewaySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", gatewaySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockGatewaySetMockRecorder) Equal(gatewaySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockGatewaySet)(nil).Equal), gatewaySet)
}

// Find mocks base method.
func (m *MockGatewaySet) Find(id ezkube.ResourceId) (*v1.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockGatewaySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGatewaySet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockGatewaySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockGatewaySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockGatewaySet)(nil).Generic))
}

// Has mocks base method.
func (m *MockGatewaySet) Has(gateway ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", gateway)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockGatewaySetMockRecorder) Has(gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockGatewaySet)(nil).Has), gateway)
}

// Insert mocks base method.
func (m *MockGatewaySet) Insert(gateway ...*v1.Gateway) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range gateway {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockGatewaySetMockRecorder) Insert(gateway ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockGatewaySet)(nil).Insert), gateway...)
}

// Intersection mocks base method.
func (m *MockGatewaySet) Intersection(set v1beta1sets.GatewaySet) v1beta1sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1beta1sets.GatewaySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockGatewaySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockGatewaySet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockGatewaySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockGatewaySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockGatewaySet)(nil).Keys))
}

// Length mocks base method.
func (m *MockGatewaySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockGatewaySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockGatewaySet)(nil).Length))
}

// List mocks base method.
func (m *MockGatewaySet) List(filterResource ...func(*v1.Gateway) bool) []*v1.Gateway {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.Gateway)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockGatewaySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGatewaySet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockGatewaySet) Map() map[string]*v1.Gateway {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Gateway)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockGatewaySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockGatewaySet)(nil).Map))
}

// Union mocks base method.
func (m *MockGatewaySet) Union(set v1beta1sets.GatewaySet) v1beta1sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1beta1sets.GatewaySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockGatewaySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockGatewaySet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockGatewaySet) UnsortedList(filterResource ...func(*v1.Gateway) bool) []*v1.Gateway {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.Gateway)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockGatewaySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockGatewaySet)(nil).UnsortedList), filterResource...)
}

// MockGatewayClassSet is a mock of GatewayClassSet interface.
type MockGatewayClassSet struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClassSetMockRecorder
}

// MockGatewayClassSetMockRecorder is the mock recorder for MockGatewayClassSet.
type MockGatewayClassSetMockRecorder struct {
	mock *MockGatewayClassSet
}

// NewMockGatewayClassSet creates a new mock instance.
func NewMockGatewayClassSet(ctrl *gomock.Controller) *MockGatewayClassSet {
	mock := &MockGatewayClassSet{ctrl: ctrl}
	mock.recorder = &MockGatewayClassSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayClassSet) EXPECT() *MockGatewayClassSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockGatewayClassSet) Clone() v1beta1sets.GatewayClassSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1beta1sets.GatewayClassSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockGatewayClassSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockGatewayClassSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockGatewayClassSet) Delete(gatewayClass ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", gatewayClass)
}

// Delete indicates an expected call of Delete.
func (mr *MockGatewayClassSetMockRecorder) Delete(gatewayClass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGatewayClassSet)(nil).Delete), gatewayClass)
}

// Delta mocks base method.
func (m *MockGatewayClassSet) Delta(newSet v1beta1sets.GatewayClassSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockGatewayClassSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockGatewayClassSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockGatewayClassSet) Difference(set v1beta1sets.GatewayClassSet) v1beta1sets.GatewayClassSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1beta1sets.GatewayClassSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockGatewayClassSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockGatewayClassSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockGatewayClassSet) Equal(gatewayClassSet v1beta1sets.GatewayClassSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", gatewayClassSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockGatewayClassSetMockRecorder) Equal(gatewayClassSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockGatewayClassSet)(nil).Equal), gatewayClassSet)
}

// Find mocks base method.
func (m *MockGatewayClassSet) Find(id ezkube.ResourceId) (*v1.GatewayClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.GatewayClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockGatewayClassSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGatewayClassSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockGatewayClassSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockGatewayClassSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockGatewayClassSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockGatewayClassSet) Has(gatewayClass ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", gatewayClass)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockGatewayClassSetMockRecorder) Has(gatewayClass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockGatewayClassSet)(nil).Has), gatewayClass)
}

// Insert mocks base method.
func (m *MockGatewayClassSet) Insert(gatewayClass ...*v1.GatewayClass) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range gatewayClass {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockGatewayClassSetMockRecorder) Insert(gatewayClass ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockGatewayClassSet)(nil).Insert), gatewayClass...)
}

// Intersection mocks base method.
func (m *MockGatewayClassSet) Intersection(set v1beta1sets.GatewayClassSet) v1beta1sets.GatewayClassSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1beta1sets.GatewayClassSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockGatewayClassSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockGatewayClassSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockGatewayClassSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockGatewayClassSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockGatewayClassSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockGatewayClassSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockGatewayClassSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockGatewayClassSet)(nil).Length))
}

// List mocks base method.
func (m *MockGatewayClassSet) List(filterResource ...func(*v1.GatewayClass) bool) []*v1.GatewayClass {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.GatewayClass)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockGatewayClassSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGatewayClassSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockGatewayClassSet) Map() map[string]*v1.GatewayClass {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.GatewayClass)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockGatewayClassSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockGatewayClassSet)(nil).Map))
}

// Union mocks base method.
func (m *MockGatewayClassSet) Union(set v1beta1sets.GatewayClassSet) v1beta1sets.GatewayClassSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1beta1sets.GatewayClassSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockGatewayClassSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockGatewayClassSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockGatewayClassSet) UnsortedList(filterResource ...func(*v1.GatewayClass) bool) []*v1.GatewayClass {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.GatewayClass)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockGatewayClassSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockGatewayClassSet)(nil).UnsortedList), filterResource...)
}

// MockHTTPRouteSet is a mock of HTTPRouteSet interface.
type MockHTTPRouteSet struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteSetMockRecorder
}

// MockHTTPRouteSetMockRecorder is the mock recorder for MockHTTPRouteSet.
type MockHTTPRouteSetMockRecorder struct {
	mock *MockHTTPRouteSet
}

// NewMockHTTPRouteSet creates a new mock instance.
func NewMockHTTPRouteSet(ctrl *gomock.Controller) *MockHTTPRouteSet {
	mock := &MockHTTPRouteSet{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPRouteSet) EXPECT() *MockHTTPRouteSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockHTTPRouteSet) Clone() v1beta1sets.HTTPRouteSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1beta1sets.HTTPRouteSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockHTTPRouteSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockHTTPRouteSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockHTTPRouteSet) Delete(hTTPRoute ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", hTTPRoute)
}

// Delete indicates an expected call of Delete.
func (mr *MockHTTPRouteSetMockRecorder) Delete(hTTPRoute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHTTPRouteSet)(nil).Delete), hTTPRoute)
}

// Delta mocks base method.
func (m *MockHTTPRouteSet) Delta(newSet v1beta1sets.HTTPRouteSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockHTTPRouteSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockHTTPRouteSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockHTTPRouteSet) Difference(set v1beta1sets.HTTPRouteSet) v1beta1sets.HTTPRouteSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1beta1sets.HTTPRouteSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockHTTPRouteSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockHTTPRouteSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockHTTPRouteSet) Equal(hTTPRouteSet v1beta1sets.HTTPRouteSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", hTTPRouteSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockHTTPRouteSetMockRecorder) Equal(hTTPRouteSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockHTTPRouteSet)(nil).Equal), hTTPRouteSet)
}

// Find mocks base method.
func (m *MockHTTPRouteSet) Find(id ezkube.ResourceId) (*v1.HTTPRoute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.HTTPRoute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockHTTPRouteSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockHTTPRouteSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockHTTPRouteSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockHTTPRouteSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockHTTPRouteSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockHTTPRouteSet) Has(hTTPRoute ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", hTTPRoute)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockHTTPRouteSetMockRecorder) Has(hTTPRoute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockHTTPRouteSet)(nil).Has), hTTPRoute)
}

// Insert mocks base method.
func (m *MockHTTPRouteSet) Insert(hTTPRoute ...*v1.HTTPRoute) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range hTTPRoute {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockHTTPRouteSetMockRecorder) Insert(hTTPRoute ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockHTTPRouteSet)(nil).Insert), hTTPRoute...)
}

// Intersection mocks base method.
func (m *MockHTTPRouteSet) Intersection(set v1beta1sets.HTTPRouteSet) v1beta1sets.HTTPRouteSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1beta1sets.HTTPRouteSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockHTTPRouteSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockHTTPRouteSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockHTTPRouteSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockHTTPRouteSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockHTTPRouteSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockHTTPRouteSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockHTTPRouteSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockHTTPRouteSet)(nil).Length))
}

// List mocks base method.
func (m *MockHTTPRouteSet) List(filterResource ...func(*v1.HTTPRoute) bool) []*v1.HTTPRoute {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.HTTPRoute)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockHTTPRouteSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockHTTPRouteSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockHTTPRouteSet) Map() map[string]*v1.HTTPRoute {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.HTTPRoute)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockHTTPRouteSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockHTTPRouteSet)(nil).Map))
}

// Union mocks base method.
func (m *MockHTTPRouteSet) Union(set v1beta1sets.HTTPRouteSet) v1beta1sets.HTTPRouteSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1beta1sets.HTTPRouteSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockHTTPRouteSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockHTTPRouteSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockHTTPRouteSet) UnsortedList(filterResource ...func(*v1.HTTPRoute) bool) []*v1.HTTPRoute {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.HTTPRoute)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockHTTPRouteSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockHTTPRouteSet)(nil).UnsortedList), filterResource...)
}
