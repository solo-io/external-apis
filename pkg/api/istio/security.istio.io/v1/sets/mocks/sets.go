// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockAuthorizationPolicySet is a mock of AuthorizationPolicySet interface.
type MockAuthorizationPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationPolicySetMockRecorder
}

// MockAuthorizationPolicySetMockRecorder is the mock recorder for MockAuthorizationPolicySet.
type MockAuthorizationPolicySetMockRecorder struct {
	mock *MockAuthorizationPolicySet
}

// NewMockAuthorizationPolicySet creates a new mock instance.
func NewMockAuthorizationPolicySet(ctrl *gomock.Controller) *MockAuthorizationPolicySet {
	mock := &MockAuthorizationPolicySet{ctrl: ctrl}
	mock.recorder = &MockAuthorizationPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationPolicySet) EXPECT() *MockAuthorizationPolicySetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockAuthorizationPolicySet) Clone() v1sets.AuthorizationPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.AuthorizationPolicySet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockAuthorizationPolicySetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockAuthorizationPolicySet) Delete(authorizationPolicy ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", authorizationPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorizationPolicySetMockRecorder) Delete(authorizationPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Delete), authorizationPolicy)
}

// Delta mocks base method.
func (m *MockAuthorizationPolicySet) Delta(newSet v1sets.AuthorizationPolicySet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockAuthorizationPolicySetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockAuthorizationPolicySet) Difference(set v1sets.AuthorizationPolicySet) v1sets.AuthorizationPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.AuthorizationPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockAuthorizationPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockAuthorizationPolicySet) Equal(authorizationPolicySet v1sets.AuthorizationPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", authorizationPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockAuthorizationPolicySetMockRecorder) Equal(authorizationPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Equal), authorizationPolicySet)
}

// Find mocks base method.
func (m *MockAuthorizationPolicySet) Find(id ezkube.ResourceId) (*v1beta1.AuthorizationPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1beta1.AuthorizationPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAuthorizationPolicySetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockAuthorizationPolicySet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockAuthorizationPolicySetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Generic))
}

// Has mocks base method.
func (m *MockAuthorizationPolicySet) Has(authorizationPolicy ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", authorizationPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockAuthorizationPolicySetMockRecorder) Has(authorizationPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Has), authorizationPolicy)
}

// Insert mocks base method.
func (m *MockAuthorizationPolicySet) Insert(authorizationPolicy ...*v1beta1.AuthorizationPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range authorizationPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockAuthorizationPolicySetMockRecorder) Insert(authorizationPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Insert), authorizationPolicy...)
}

// Intersection mocks base method.
func (m *MockAuthorizationPolicySet) Intersection(set v1sets.AuthorizationPolicySet) v1sets.AuthorizationPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.AuthorizationPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockAuthorizationPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockAuthorizationPolicySet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockAuthorizationPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Keys))
}

// Length mocks base method.
func (m *MockAuthorizationPolicySet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockAuthorizationPolicySetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Length))
}

// List mocks base method.
func (m *MockAuthorizationPolicySet) List(filterResource ...func(*v1beta1.AuthorizationPolicy) bool) []*v1beta1.AuthorizationPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1beta1.AuthorizationPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockAuthorizationPolicySetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockAuthorizationPolicySet) Map() map[string]*v1beta1.AuthorizationPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1beta1.AuthorizationPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockAuthorizationPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Map))
}

// Union mocks base method.
func (m *MockAuthorizationPolicySet) Union(set v1sets.AuthorizationPolicySet) v1sets.AuthorizationPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.AuthorizationPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockAuthorizationPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockAuthorizationPolicySet) UnsortedList(filterResource ...func(*v1beta1.AuthorizationPolicy) bool) []*v1beta1.AuthorizationPolicy {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1beta1.AuthorizationPolicy)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockAuthorizationPolicySetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).UnsortedList), filterResource...)
}

// MockPeerAuthenticationSet is a mock of PeerAuthenticationSet interface.
type MockPeerAuthenticationSet struct {
	ctrl     *gomock.Controller
	recorder *MockPeerAuthenticationSetMockRecorder
}

// MockPeerAuthenticationSetMockRecorder is the mock recorder for MockPeerAuthenticationSet.
type MockPeerAuthenticationSetMockRecorder struct {
	mock *MockPeerAuthenticationSet
}

// NewMockPeerAuthenticationSet creates a new mock instance.
func NewMockPeerAuthenticationSet(ctrl *gomock.Controller) *MockPeerAuthenticationSet {
	mock := &MockPeerAuthenticationSet{ctrl: ctrl}
	mock.recorder = &MockPeerAuthenticationSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerAuthenticationSet) EXPECT() *MockPeerAuthenticationSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockPeerAuthenticationSet) Clone() v1sets.PeerAuthenticationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.PeerAuthenticationSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockPeerAuthenticationSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockPeerAuthenticationSet) Delete(peerAuthentication ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", peerAuthentication)
}

// Delete indicates an expected call of Delete.
func (mr *MockPeerAuthenticationSetMockRecorder) Delete(peerAuthentication interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Delete), peerAuthentication)
}

// Delta mocks base method.
func (m *MockPeerAuthenticationSet) Delta(newSet v1sets.PeerAuthenticationSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockPeerAuthenticationSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockPeerAuthenticationSet) Difference(set v1sets.PeerAuthenticationSet) v1sets.PeerAuthenticationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.PeerAuthenticationSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockPeerAuthenticationSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockPeerAuthenticationSet) Equal(peerAuthenticationSet v1sets.PeerAuthenticationSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", peerAuthenticationSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockPeerAuthenticationSetMockRecorder) Equal(peerAuthenticationSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Equal), peerAuthenticationSet)
}

// Find mocks base method.
func (m *MockPeerAuthenticationSet) Find(id ezkube.ResourceId) (*v1beta1.PeerAuthentication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1beta1.PeerAuthentication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockPeerAuthenticationSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockPeerAuthenticationSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockPeerAuthenticationSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockPeerAuthenticationSet) Has(peerAuthentication ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", peerAuthentication)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockPeerAuthenticationSetMockRecorder) Has(peerAuthentication interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Has), peerAuthentication)
}

// Insert mocks base method.
func (m *MockPeerAuthenticationSet) Insert(peerAuthentication ...*v1beta1.PeerAuthentication) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range peerAuthentication {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockPeerAuthenticationSetMockRecorder) Insert(peerAuthentication ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Insert), peerAuthentication...)
}

// Intersection mocks base method.
func (m *MockPeerAuthenticationSet) Intersection(set v1sets.PeerAuthenticationSet) v1sets.PeerAuthenticationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.PeerAuthenticationSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockPeerAuthenticationSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockPeerAuthenticationSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockPeerAuthenticationSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockPeerAuthenticationSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockPeerAuthenticationSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Length))
}

// List mocks base method.
func (m *MockPeerAuthenticationSet) List(filterResource ...func(*v1beta1.PeerAuthentication) bool) []*v1beta1.PeerAuthentication {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1beta1.PeerAuthentication)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockPeerAuthenticationSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockPeerAuthenticationSet) Map() map[string]*v1beta1.PeerAuthentication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1beta1.PeerAuthentication)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockPeerAuthenticationSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Map))
}

// Union mocks base method.
func (m *MockPeerAuthenticationSet) Union(set v1sets.PeerAuthenticationSet) v1sets.PeerAuthenticationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.PeerAuthenticationSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockPeerAuthenticationSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockPeerAuthenticationSet) UnsortedList(filterResource ...func(*v1beta1.PeerAuthentication) bool) []*v1beta1.PeerAuthentication {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1beta1.PeerAuthentication)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockPeerAuthenticationSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockPeerAuthenticationSet)(nil).UnsortedList), filterResource...)
}

// MockRequestAuthenticationSet is a mock of RequestAuthenticationSet interface.
type MockRequestAuthenticationSet struct {
	ctrl     *gomock.Controller
	recorder *MockRequestAuthenticationSetMockRecorder
}

// MockRequestAuthenticationSetMockRecorder is the mock recorder for MockRequestAuthenticationSet.
type MockRequestAuthenticationSetMockRecorder struct {
	mock *MockRequestAuthenticationSet
}

// NewMockRequestAuthenticationSet creates a new mock instance.
func NewMockRequestAuthenticationSet(ctrl *gomock.Controller) *MockRequestAuthenticationSet {
	mock := &MockRequestAuthenticationSet{ctrl: ctrl}
	mock.recorder = &MockRequestAuthenticationSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestAuthenticationSet) EXPECT() *MockRequestAuthenticationSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockRequestAuthenticationSet) Clone() v1sets.RequestAuthenticationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.RequestAuthenticationSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockRequestAuthenticationSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockRequestAuthenticationSet) Delete(requestAuthentication ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", requestAuthentication)
}

// Delete indicates an expected call of Delete.
func (mr *MockRequestAuthenticationSetMockRecorder) Delete(requestAuthentication interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Delete), requestAuthentication)
}

// Delta mocks base method.
func (m *MockRequestAuthenticationSet) Delta(newSet v1sets.RequestAuthenticationSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockRequestAuthenticationSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockRequestAuthenticationSet) Difference(set v1sets.RequestAuthenticationSet) v1sets.RequestAuthenticationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.RequestAuthenticationSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockRequestAuthenticationSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockRequestAuthenticationSet) Equal(requestAuthenticationSet v1sets.RequestAuthenticationSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", requestAuthenticationSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockRequestAuthenticationSetMockRecorder) Equal(requestAuthenticationSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Equal), requestAuthenticationSet)
}

// Find mocks base method.
func (m *MockRequestAuthenticationSet) Find(id ezkube.ResourceId) (*v1beta1.RequestAuthentication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1beta1.RequestAuthentication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRequestAuthenticationSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockRequestAuthenticationSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockRequestAuthenticationSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockRequestAuthenticationSet) Has(requestAuthentication ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", requestAuthentication)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockRequestAuthenticationSetMockRecorder) Has(requestAuthentication interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Has), requestAuthentication)
}

// Insert mocks base method.
func (m *MockRequestAuthenticationSet) Insert(requestAuthentication ...*v1beta1.RequestAuthentication) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range requestAuthentication {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockRequestAuthenticationSetMockRecorder) Insert(requestAuthentication ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Insert), requestAuthentication...)
}

// Intersection mocks base method.
func (m *MockRequestAuthenticationSet) Intersection(set v1sets.RequestAuthenticationSet) v1sets.RequestAuthenticationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.RequestAuthenticationSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockRequestAuthenticationSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockRequestAuthenticationSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockRequestAuthenticationSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockRequestAuthenticationSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockRequestAuthenticationSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Length))
}

// List mocks base method.
func (m *MockRequestAuthenticationSet) List(filterResource ...func(*v1beta1.RequestAuthentication) bool) []*v1beta1.RequestAuthentication {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1beta1.RequestAuthentication)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRequestAuthenticationSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockRequestAuthenticationSet) Map() map[string]*v1beta1.RequestAuthentication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1beta1.RequestAuthentication)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRequestAuthenticationSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Map))
}

// Union mocks base method.
func (m *MockRequestAuthenticationSet) Union(set v1sets.RequestAuthenticationSet) v1sets.RequestAuthenticationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.RequestAuthenticationSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockRequestAuthenticationSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockRequestAuthenticationSet) UnsortedList(filterResource ...func(*v1beta1.RequestAuthentication) bool) []*v1beta1.RequestAuthentication {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1beta1.RequestAuthentication)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockRequestAuthenticationSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockRequestAuthenticationSet)(nil).UnsortedList), filterResource...)
}
