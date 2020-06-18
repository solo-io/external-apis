// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha3sets is a generated GoMock package.
package mock_v1alpha3sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockDestinationRuleSet is a mock of DestinationRuleSet interface.
type MockDestinationRuleSet struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationRuleSetMockRecorder
}

// MockDestinationRuleSetMockRecorder is the mock recorder for MockDestinationRuleSet.
type MockDestinationRuleSetMockRecorder struct {
	mock *MockDestinationRuleSet
}

// NewMockDestinationRuleSet creates a new mock instance.
func NewMockDestinationRuleSet(ctrl *gomock.Controller) *MockDestinationRuleSet {
	mock := &MockDestinationRuleSet{ctrl: ctrl}
	mock.recorder = &MockDestinationRuleSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationRuleSet) EXPECT() *MockDestinationRuleSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockDestinationRuleSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockDestinationRuleSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockDestinationRuleSet)(nil).Keys))
}

// List mocks base method.
func (m *MockDestinationRuleSet) List() []*v1alpha3.DestinationRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha3.DestinationRule)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockDestinationRuleSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDestinationRuleSet)(nil).List))
}

// Map mocks base method.
func (m *MockDestinationRuleSet) Map() map[string]*v1alpha3.DestinationRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha3.DestinationRule)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockDestinationRuleSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockDestinationRuleSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockDestinationRuleSet) Insert(destinationRule ...*v1alpha3.DestinationRule) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range destinationRule {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockDestinationRuleSetMockRecorder) Insert(destinationRule ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDestinationRuleSet)(nil).Insert), destinationRule...)
}

// Equal mocks base method.
func (m *MockDestinationRuleSet) Equal(destinationRuleSet v1alpha3sets.DestinationRuleSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", destinationRuleSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockDestinationRuleSetMockRecorder) Equal(destinationRuleSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockDestinationRuleSet)(nil).Equal), destinationRuleSet)
}

// Has mocks base method.
func (m *MockDestinationRuleSet) Has(destinationRule *v1alpha3.DestinationRule) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", destinationRule)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockDestinationRuleSetMockRecorder) Has(destinationRule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockDestinationRuleSet)(nil).Has), destinationRule)
}

// Delete mocks base method.
func (m *MockDestinationRuleSet) Delete(destinationRule *v1alpha3.DestinationRule) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", destinationRule)
}

// Delete indicates an expected call of Delete.
func (mr *MockDestinationRuleSetMockRecorder) Delete(destinationRule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDestinationRuleSet)(nil).Delete), destinationRule)
}

// Union mocks base method.
func (m *MockDestinationRuleSet) Union(set v1alpha3sets.DestinationRuleSet) v1alpha3sets.DestinationRuleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha3sets.DestinationRuleSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockDestinationRuleSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockDestinationRuleSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockDestinationRuleSet) Difference(set v1alpha3sets.DestinationRuleSet) v1alpha3sets.DestinationRuleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha3sets.DestinationRuleSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockDestinationRuleSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockDestinationRuleSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockDestinationRuleSet) Intersection(set v1alpha3sets.DestinationRuleSet) v1alpha3sets.DestinationRuleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha3sets.DestinationRuleSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockDestinationRuleSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockDestinationRuleSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockDestinationRuleSet) Find(id ezkube.ResourceId) (*v1alpha3.DestinationRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha3.DestinationRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockDestinationRuleSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDestinationRuleSet)(nil).Find), id)
}

// MockEnvoyFilterSet is a mock of EnvoyFilterSet interface.
type MockEnvoyFilterSet struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterSetMockRecorder
}

// MockEnvoyFilterSetMockRecorder is the mock recorder for MockEnvoyFilterSet.
type MockEnvoyFilterSetMockRecorder struct {
	mock *MockEnvoyFilterSet
}

// NewMockEnvoyFilterSet creates a new mock instance.
func NewMockEnvoyFilterSet(ctrl *gomock.Controller) *MockEnvoyFilterSet {
	mock := &MockEnvoyFilterSet{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterSet) EXPECT() *MockEnvoyFilterSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockEnvoyFilterSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockEnvoyFilterSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Keys))
}

// List mocks base method.
func (m *MockEnvoyFilterSet) List() []*v1alpha3.EnvoyFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha3.EnvoyFilter)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockEnvoyFilterSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEnvoyFilterSet)(nil).List))
}

// Map mocks base method.
func (m *MockEnvoyFilterSet) Map() map[string]*v1alpha3.EnvoyFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha3.EnvoyFilter)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockEnvoyFilterSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockEnvoyFilterSet) Insert(envoyFilter ...*v1alpha3.EnvoyFilter) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range envoyFilter {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockEnvoyFilterSetMockRecorder) Insert(envoyFilter ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Insert), envoyFilter...)
}

// Equal mocks base method.
func (m *MockEnvoyFilterSet) Equal(envoyFilterSet v1alpha3sets.EnvoyFilterSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", envoyFilterSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockEnvoyFilterSetMockRecorder) Equal(envoyFilterSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Equal), envoyFilterSet)
}

// Has mocks base method.
func (m *MockEnvoyFilterSet) Has(envoyFilter *v1alpha3.EnvoyFilter) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", envoyFilter)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockEnvoyFilterSetMockRecorder) Has(envoyFilter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Has), envoyFilter)
}

// Delete mocks base method.
func (m *MockEnvoyFilterSet) Delete(envoyFilter *v1alpha3.EnvoyFilter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", envoyFilter)
}

// Delete indicates an expected call of Delete.
func (mr *MockEnvoyFilterSetMockRecorder) Delete(envoyFilter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Delete), envoyFilter)
}

// Union mocks base method.
func (m *MockEnvoyFilterSet) Union(set v1alpha3sets.EnvoyFilterSet) v1alpha3sets.EnvoyFilterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha3sets.EnvoyFilterSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockEnvoyFilterSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockEnvoyFilterSet) Difference(set v1alpha3sets.EnvoyFilterSet) v1alpha3sets.EnvoyFilterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha3sets.EnvoyFilterSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockEnvoyFilterSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockEnvoyFilterSet) Intersection(set v1alpha3sets.EnvoyFilterSet) v1alpha3sets.EnvoyFilterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha3sets.EnvoyFilterSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockEnvoyFilterSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockEnvoyFilterSet) Find(id ezkube.ResourceId) (*v1alpha3.EnvoyFilter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha3.EnvoyFilter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockEnvoyFilterSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockEnvoyFilterSet)(nil).Find), id)
}

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

// Keys mocks base method.
func (m *MockGatewaySet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockGatewaySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockGatewaySet)(nil).Keys))
}

// List mocks base method.
func (m *MockGatewaySet) List() []*v1alpha3.Gateway {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha3.Gateway)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockGatewaySetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGatewaySet)(nil).List))
}

// Map mocks base method.
func (m *MockGatewaySet) Map() map[string]*v1alpha3.Gateway {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha3.Gateway)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockGatewaySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockGatewaySet)(nil).Map))
}

// Insert mocks base method.
func (m *MockGatewaySet) Insert(gateway ...*v1alpha3.Gateway) {
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

// Equal mocks base method.
func (m *MockGatewaySet) Equal(gatewaySet v1alpha3sets.GatewaySet) bool {
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

// Has mocks base method.
func (m *MockGatewaySet) Has(gateway *v1alpha3.Gateway) bool {
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

// Delete mocks base method.
func (m *MockGatewaySet) Delete(gateway *v1alpha3.Gateway) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", gateway)
}

// Delete indicates an expected call of Delete.
func (mr *MockGatewaySetMockRecorder) Delete(gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGatewaySet)(nil).Delete), gateway)
}

// Union mocks base method.
func (m *MockGatewaySet) Union(set v1alpha3sets.GatewaySet) v1alpha3sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha3sets.GatewaySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockGatewaySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockGatewaySet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockGatewaySet) Difference(set v1alpha3sets.GatewaySet) v1alpha3sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha3sets.GatewaySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockGatewaySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockGatewaySet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockGatewaySet) Intersection(set v1alpha3sets.GatewaySet) v1alpha3sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha3sets.GatewaySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockGatewaySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockGatewaySet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockGatewaySet) Find(id ezkube.ResourceId) (*v1alpha3.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha3.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockGatewaySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGatewaySet)(nil).Find), id)
}

// MockServiceEntrySet is a mock of ServiceEntrySet interface.
type MockServiceEntrySet struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntrySetMockRecorder
}

// MockServiceEntrySetMockRecorder is the mock recorder for MockServiceEntrySet.
type MockServiceEntrySetMockRecorder struct {
	mock *MockServiceEntrySet
}

// NewMockServiceEntrySet creates a new mock instance.
func NewMockServiceEntrySet(ctrl *gomock.Controller) *MockServiceEntrySet {
	mock := &MockServiceEntrySet{ctrl: ctrl}
	mock.recorder = &MockServiceEntrySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEntrySet) EXPECT() *MockServiceEntrySetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockServiceEntrySet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockServiceEntrySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockServiceEntrySet)(nil).Keys))
}

// List mocks base method.
func (m *MockServiceEntrySet) List() []*v1alpha3.ServiceEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha3.ServiceEntry)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockServiceEntrySetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceEntrySet)(nil).List))
}

// Map mocks base method.
func (m *MockServiceEntrySet) Map() map[string]*v1alpha3.ServiceEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha3.ServiceEntry)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockServiceEntrySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockServiceEntrySet)(nil).Map))
}

// Insert mocks base method.
func (m *MockServiceEntrySet) Insert(serviceEntry ...*v1alpha3.ServiceEntry) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range serviceEntry {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceEntrySetMockRecorder) Insert(serviceEntry ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceEntrySet)(nil).Insert), serviceEntry...)
}

// Equal mocks base method.
func (m *MockServiceEntrySet) Equal(serviceEntrySet v1alpha3sets.ServiceEntrySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", serviceEntrySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockServiceEntrySetMockRecorder) Equal(serviceEntrySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockServiceEntrySet)(nil).Equal), serviceEntrySet)
}

// Has mocks base method.
func (m *MockServiceEntrySet) Has(serviceEntry *v1alpha3.ServiceEntry) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", serviceEntry)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockServiceEntrySetMockRecorder) Has(serviceEntry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockServiceEntrySet)(nil).Has), serviceEntry)
}

// Delete mocks base method.
func (m *MockServiceEntrySet) Delete(serviceEntry *v1alpha3.ServiceEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", serviceEntry)
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceEntrySetMockRecorder) Delete(serviceEntry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceEntrySet)(nil).Delete), serviceEntry)
}

// Union mocks base method.
func (m *MockServiceEntrySet) Union(set v1alpha3sets.ServiceEntrySet) v1alpha3sets.ServiceEntrySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha3sets.ServiceEntrySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockServiceEntrySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockServiceEntrySet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockServiceEntrySet) Difference(set v1alpha3sets.ServiceEntrySet) v1alpha3sets.ServiceEntrySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha3sets.ServiceEntrySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockServiceEntrySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockServiceEntrySet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockServiceEntrySet) Intersection(set v1alpha3sets.ServiceEntrySet) v1alpha3sets.ServiceEntrySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha3sets.ServiceEntrySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockServiceEntrySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockServiceEntrySet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockServiceEntrySet) Find(id ezkube.ResourceId) (*v1alpha3.ServiceEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha3.ServiceEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockServiceEntrySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockServiceEntrySet)(nil).Find), id)
}

// MockVirtualServiceSet is a mock of VirtualServiceSet interface.
type MockVirtualServiceSet struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceSetMockRecorder
}

// MockVirtualServiceSetMockRecorder is the mock recorder for MockVirtualServiceSet.
type MockVirtualServiceSetMockRecorder struct {
	mock *MockVirtualServiceSet
}

// NewMockVirtualServiceSet creates a new mock instance.
func NewMockVirtualServiceSet(ctrl *gomock.Controller) *MockVirtualServiceSet {
	mock := &MockVirtualServiceSet{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualServiceSet) EXPECT() *MockVirtualServiceSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockVirtualServiceSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockVirtualServiceSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockVirtualServiceSet)(nil).Keys))
}

// List mocks base method.
func (m *MockVirtualServiceSet) List() []*v1alpha3.VirtualService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha3.VirtualService)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockVirtualServiceSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualServiceSet)(nil).List))
}

// Map mocks base method.
func (m *MockVirtualServiceSet) Map() map[string]*v1alpha3.VirtualService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha3.VirtualService)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockVirtualServiceSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockVirtualServiceSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockVirtualServiceSet) Insert(virtualService ...*v1alpha3.VirtualService) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range virtualService {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockVirtualServiceSetMockRecorder) Insert(virtualService ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockVirtualServiceSet)(nil).Insert), virtualService...)
}

// Equal mocks base method.
func (m *MockVirtualServiceSet) Equal(virtualServiceSet v1alpha3sets.VirtualServiceSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", virtualServiceSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockVirtualServiceSetMockRecorder) Equal(virtualServiceSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockVirtualServiceSet)(nil).Equal), virtualServiceSet)
}

// Has mocks base method.
func (m *MockVirtualServiceSet) Has(virtualService *v1alpha3.VirtualService) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", virtualService)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockVirtualServiceSetMockRecorder) Has(virtualService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockVirtualServiceSet)(nil).Has), virtualService)
}

// Delete mocks base method.
func (m *MockVirtualServiceSet) Delete(virtualService *v1alpha3.VirtualService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", virtualService)
}

// Delete indicates an expected call of Delete.
func (mr *MockVirtualServiceSetMockRecorder) Delete(virtualService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualServiceSet)(nil).Delete), virtualService)
}

// Union mocks base method.
func (m *MockVirtualServiceSet) Union(set v1alpha3sets.VirtualServiceSet) v1alpha3sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha3sets.VirtualServiceSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockVirtualServiceSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockVirtualServiceSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockVirtualServiceSet) Difference(set v1alpha3sets.VirtualServiceSet) v1alpha3sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha3sets.VirtualServiceSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockVirtualServiceSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockVirtualServiceSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockVirtualServiceSet) Intersection(set v1alpha3sets.VirtualServiceSet) v1alpha3sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha3sets.VirtualServiceSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockVirtualServiceSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockVirtualServiceSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockVirtualServiceSet) Find(id ezkube.ResourceId) (*v1alpha3.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha3.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockVirtualServiceSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockVirtualServiceSet)(nil).Find), id)
}
