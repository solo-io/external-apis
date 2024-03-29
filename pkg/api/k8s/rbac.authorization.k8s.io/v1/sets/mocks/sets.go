// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/rbac.authorization.k8s.io/v1/sets"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1 "k8s.io/api/rbac/v1"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockRoleSet is a mock of RoleSet interface.
type MockRoleSet struct {
	ctrl     *gomock.Controller
	recorder *MockRoleSetMockRecorder
}

// MockRoleSetMockRecorder is the mock recorder for MockRoleSet.
type MockRoleSetMockRecorder struct {
	mock *MockRoleSet
}

// NewMockRoleSet creates a new mock instance.
func NewMockRoleSet(ctrl *gomock.Controller) *MockRoleSet {
	mock := &MockRoleSet{ctrl: ctrl}
	mock.recorder = &MockRoleSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleSet) EXPECT() *MockRoleSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockRoleSet) Clone() v1sets.RoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.RoleSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockRoleSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockRoleSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockRoleSet) Delete(role ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", role)
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleSetMockRecorder) Delete(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleSet)(nil).Delete), role)
}

// Delta mocks base method.
func (m *MockRoleSet) Delta(newSet v1sets.RoleSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockRoleSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockRoleSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockRoleSet) Difference(set v1sets.RoleSet) v1sets.RoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.RoleSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockRoleSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockRoleSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockRoleSet) Equal(roleSet v1sets.RoleSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", roleSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockRoleSetMockRecorder) Equal(roleSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockRoleSet)(nil).Equal), roleSet)
}

// Find mocks base method.
func (m *MockRoleSet) Find(id ezkube.ResourceId) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRoleSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRoleSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockRoleSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockRoleSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockRoleSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockRoleSet) Has(role ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", role)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockRoleSetMockRecorder) Has(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRoleSet)(nil).Has), role)
}

// Insert mocks base method.
func (m *MockRoleSet) Insert(role ...*v1.Role) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range role {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockRoleSetMockRecorder) Insert(role ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRoleSet)(nil).Insert), role...)
}

// Intersection mocks base method.
func (m *MockRoleSet) Intersection(set v1sets.RoleSet) v1sets.RoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.RoleSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockRoleSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockRoleSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockRoleSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockRoleSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRoleSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockRoleSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockRoleSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockRoleSet)(nil).Length))
}

// List mocks base method.
func (m *MockRoleSet) List(filterResource ...func(*v1.Role) bool) []*v1.Role {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.Role)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRoleSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockRoleSet) Map() map[string]*v1.Role {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Role)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRoleSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRoleSet)(nil).Map))
}

// Union mocks base method.
func (m *MockRoleSet) Union(set v1sets.RoleSet) v1sets.RoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.RoleSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockRoleSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockRoleSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockRoleSet) UnsortedList(filterResource ...func(*v1.Role) bool) []*v1.Role {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.Role)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockRoleSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockRoleSet)(nil).UnsortedList), filterResource...)
}

// MockRoleBindingSet is a mock of RoleBindingSet interface.
type MockRoleBindingSet struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingSetMockRecorder
}

// MockRoleBindingSetMockRecorder is the mock recorder for MockRoleBindingSet.
type MockRoleBindingSetMockRecorder struct {
	mock *MockRoleBindingSet
}

// NewMockRoleBindingSet creates a new mock instance.
func NewMockRoleBindingSet(ctrl *gomock.Controller) *MockRoleBindingSet {
	mock := &MockRoleBindingSet{ctrl: ctrl}
	mock.recorder = &MockRoleBindingSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleBindingSet) EXPECT() *MockRoleBindingSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockRoleBindingSet) Clone() v1sets.RoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.RoleBindingSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockRoleBindingSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockRoleBindingSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockRoleBindingSet) Delete(roleBinding ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", roleBinding)
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleBindingSetMockRecorder) Delete(roleBinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleBindingSet)(nil).Delete), roleBinding)
}

// Delta mocks base method.
func (m *MockRoleBindingSet) Delta(newSet v1sets.RoleBindingSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockRoleBindingSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockRoleBindingSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockRoleBindingSet) Difference(set v1sets.RoleBindingSet) v1sets.RoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.RoleBindingSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockRoleBindingSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockRoleBindingSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockRoleBindingSet) Equal(roleBindingSet v1sets.RoleBindingSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", roleBindingSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockRoleBindingSetMockRecorder) Equal(roleBindingSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockRoleBindingSet)(nil).Equal), roleBindingSet)
}

// Find mocks base method.
func (m *MockRoleBindingSet) Find(id ezkube.ResourceId) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRoleBindingSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRoleBindingSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockRoleBindingSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockRoleBindingSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockRoleBindingSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockRoleBindingSet) Has(roleBinding ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", roleBinding)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockRoleBindingSetMockRecorder) Has(roleBinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRoleBindingSet)(nil).Has), roleBinding)
}

// Insert mocks base method.
func (m *MockRoleBindingSet) Insert(roleBinding ...*v1.RoleBinding) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range roleBinding {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockRoleBindingSetMockRecorder) Insert(roleBinding ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRoleBindingSet)(nil).Insert), roleBinding...)
}

// Intersection mocks base method.
func (m *MockRoleBindingSet) Intersection(set v1sets.RoleBindingSet) v1sets.RoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.RoleBindingSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockRoleBindingSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockRoleBindingSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockRoleBindingSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockRoleBindingSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRoleBindingSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockRoleBindingSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockRoleBindingSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockRoleBindingSet)(nil).Length))
}

// List mocks base method.
func (m *MockRoleBindingSet) List(filterResource ...func(*v1.RoleBinding) bool) []*v1.RoleBinding {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.RoleBinding)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRoleBindingSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleBindingSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockRoleBindingSet) Map() map[string]*v1.RoleBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.RoleBinding)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRoleBindingSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRoleBindingSet)(nil).Map))
}

// Union mocks base method.
func (m *MockRoleBindingSet) Union(set v1sets.RoleBindingSet) v1sets.RoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.RoleBindingSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockRoleBindingSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockRoleBindingSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockRoleBindingSet) UnsortedList(filterResource ...func(*v1.RoleBinding) bool) []*v1.RoleBinding {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.RoleBinding)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockRoleBindingSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockRoleBindingSet)(nil).UnsortedList), filterResource...)
}

// MockClusterRoleSet is a mock of ClusterRoleSet interface.
type MockClusterRoleSet struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleSetMockRecorder
}

// MockClusterRoleSetMockRecorder is the mock recorder for MockClusterRoleSet.
type MockClusterRoleSetMockRecorder struct {
	mock *MockClusterRoleSet
}

// NewMockClusterRoleSet creates a new mock instance.
func NewMockClusterRoleSet(ctrl *gomock.Controller) *MockClusterRoleSet {
	mock := &MockClusterRoleSet{ctrl: ctrl}
	mock.recorder = &MockClusterRoleSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRoleSet) EXPECT() *MockClusterRoleSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockClusterRoleSet) Clone() v1sets.ClusterRoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.ClusterRoleSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockClusterRoleSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockClusterRoleSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockClusterRoleSet) Delete(clusterRole ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", clusterRole)
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterRoleSetMockRecorder) Delete(clusterRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterRoleSet)(nil).Delete), clusterRole)
}

// Delta mocks base method.
func (m *MockClusterRoleSet) Delta(newSet v1sets.ClusterRoleSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockClusterRoleSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockClusterRoleSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockClusterRoleSet) Difference(set v1sets.ClusterRoleSet) v1sets.ClusterRoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockClusterRoleSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockClusterRoleSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockClusterRoleSet) Equal(clusterRoleSet v1sets.ClusterRoleSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", clusterRoleSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockClusterRoleSetMockRecorder) Equal(clusterRoleSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockClusterRoleSet)(nil).Equal), clusterRoleSet)
}

// Find mocks base method.
func (m *MockClusterRoleSet) Find(id ezkube.ResourceId) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockClusterRoleSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockClusterRoleSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockClusterRoleSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockClusterRoleSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockClusterRoleSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockClusterRoleSet) Has(clusterRole ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", clusterRole)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockClusterRoleSetMockRecorder) Has(clusterRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockClusterRoleSet)(nil).Has), clusterRole)
}

// Insert mocks base method.
func (m *MockClusterRoleSet) Insert(clusterRole ...*v1.ClusterRole) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range clusterRole {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockClusterRoleSetMockRecorder) Insert(clusterRole ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockClusterRoleSet)(nil).Insert), clusterRole...)
}

// Intersection mocks base method.
func (m *MockClusterRoleSet) Intersection(set v1sets.ClusterRoleSet) v1sets.ClusterRoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockClusterRoleSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockClusterRoleSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockClusterRoleSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockClusterRoleSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockClusterRoleSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockClusterRoleSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockClusterRoleSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockClusterRoleSet)(nil).Length))
}

// List mocks base method.
func (m *MockClusterRoleSet) List(filterResource ...func(*v1.ClusterRole) bool) []*v1.ClusterRole {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.ClusterRole)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockClusterRoleSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRoleSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockClusterRoleSet) Map() map[string]*v1.ClusterRole {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.ClusterRole)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockClusterRoleSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockClusterRoleSet)(nil).Map))
}

// Union mocks base method.
func (m *MockClusterRoleSet) Union(set v1sets.ClusterRoleSet) v1sets.ClusterRoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockClusterRoleSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockClusterRoleSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockClusterRoleSet) UnsortedList(filterResource ...func(*v1.ClusterRole) bool) []*v1.ClusterRole {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.ClusterRole)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockClusterRoleSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockClusterRoleSet)(nil).UnsortedList), filterResource...)
}

// MockClusterRoleBindingSet is a mock of ClusterRoleBindingSet interface.
type MockClusterRoleBindingSet struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleBindingSetMockRecorder
}

// MockClusterRoleBindingSetMockRecorder is the mock recorder for MockClusterRoleBindingSet.
type MockClusterRoleBindingSetMockRecorder struct {
	mock *MockClusterRoleBindingSet
}

// NewMockClusterRoleBindingSet creates a new mock instance.
func NewMockClusterRoleBindingSet(ctrl *gomock.Controller) *MockClusterRoleBindingSet {
	mock := &MockClusterRoleBindingSet{ctrl: ctrl}
	mock.recorder = &MockClusterRoleBindingSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRoleBindingSet) EXPECT() *MockClusterRoleBindingSetMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockClusterRoleBindingSet) Clone() v1sets.ClusterRoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.ClusterRoleBindingSet)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockClusterRoleBindingSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Clone))
}

// Delete mocks base method.
func (m *MockClusterRoleBindingSet) Delete(clusterRoleBinding ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", clusterRoleBinding)
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterRoleBindingSetMockRecorder) Delete(clusterRoleBinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Delete), clusterRoleBinding)
}

// Delta mocks base method.
func (m *MockClusterRoleBindingSet) Delta(newSet v1sets.ClusterRoleBindingSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockClusterRoleBindingSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Delta), newSet)
}

// Difference mocks base method.
func (m *MockClusterRoleBindingSet) Difference(set v1sets.ClusterRoleBindingSet) v1sets.ClusterRoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleBindingSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockClusterRoleBindingSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Difference), set)
}

// Equal mocks base method.
func (m *MockClusterRoleBindingSet) Equal(clusterRoleBindingSet v1sets.ClusterRoleBindingSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", clusterRoleBindingSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockClusterRoleBindingSetMockRecorder) Equal(clusterRoleBindingSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Equal), clusterRoleBindingSet)
}

// Find mocks base method.
func (m *MockClusterRoleBindingSet) Find(id ezkube.ResourceId) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockClusterRoleBindingSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Find), id)
}

// Generic mocks base method.
func (m *MockClusterRoleBindingSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockClusterRoleBindingSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Generic))
}

// Has mocks base method.
func (m *MockClusterRoleBindingSet) Has(clusterRoleBinding ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", clusterRoleBinding)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockClusterRoleBindingSetMockRecorder) Has(clusterRoleBinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Has), clusterRoleBinding)
}

// Insert mocks base method.
func (m *MockClusterRoleBindingSet) Insert(clusterRoleBinding ...*v1.ClusterRoleBinding) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range clusterRoleBinding {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockClusterRoleBindingSetMockRecorder) Insert(clusterRoleBinding ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Insert), clusterRoleBinding...)
}

// Intersection mocks base method.
func (m *MockClusterRoleBindingSet) Intersection(set v1sets.ClusterRoleBindingSet) v1sets.ClusterRoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleBindingSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockClusterRoleBindingSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Intersection), set)
}

// Keys mocks base method.
func (m *MockClusterRoleBindingSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockClusterRoleBindingSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Keys))
}

// Length mocks base method.
func (m *MockClusterRoleBindingSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockClusterRoleBindingSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Length))
}

// List mocks base method.
func (m *MockClusterRoleBindingSet) List(filterResource ...func(*v1.ClusterRoleBinding) bool) []*v1.ClusterRoleBinding {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.ClusterRoleBinding)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockClusterRoleBindingSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockClusterRoleBindingSet) Map() map[string]*v1.ClusterRoleBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.ClusterRoleBinding)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockClusterRoleBindingSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Map))
}

// Union mocks base method.
func (m *MockClusterRoleBindingSet) Union(set v1sets.ClusterRoleBindingSet) v1sets.ClusterRoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleBindingSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockClusterRoleBindingSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Union), set)
}

// UnsortedList mocks base method.
func (m *MockClusterRoleBindingSet) UnsortedList(filterResource ...func(*v1.ClusterRoleBinding) bool) []*v1.ClusterRoleBinding {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.ClusterRoleBinding)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList.
func (mr *MockClusterRoleBindingSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).UnsortedList), filterResource...)
}
