// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1beta1sets

import (
	security_istio_io_v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AuthorizationPolicySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*security_istio_io_v1beta1.AuthorizationPolicy) bool) []*security_istio_io_v1beta1.AuthorizationPolicy
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*security_istio_io_v1beta1.AuthorizationPolicy) bool) []*security_istio_io_v1beta1.AuthorizationPolicy
	// Return the Set as a map of key to resource.
	Map() map[string]*security_istio_io_v1beta1.AuthorizationPolicy
	// Insert a resource into the set.
	Insert(authorizationPolicy ...*security_istio_io_v1beta1.AuthorizationPolicy)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(authorizationPolicySet AuthorizationPolicySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(authorizationPolicy ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(authorizationPolicy ezkube.ResourceId)
	// Return the union with the provided set
	Union(set AuthorizationPolicySet) AuthorizationPolicySet
	// Return the difference with the provided set
	Difference(set AuthorizationPolicySet) AuthorizationPolicySet
	// Return the intersection with the provided set
	Intersection(set AuthorizationPolicySet) AuthorizationPolicySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*security_istio_io_v1beta1.AuthorizationPolicy, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another AuthorizationPolicySet
	Delta(newSet AuthorizationPolicySet) sksets.ResourceDelta
	// Create a deep copy of the current AuthorizationPolicySet
	Clone() AuthorizationPolicySet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
	// Get the equality function used by the set
	GetEqualityFunc() func(a, b client.Object) bool
}

func makeGenericAuthorizationPolicySet(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	authorizationPolicyList []*security_istio_io_v1beta1.AuthorizationPolicy,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range authorizationPolicyList {
		genericResources = append(genericResources, obj)
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	genericEqualityFunc := func(a, b ezkube.ResourceId) bool {
		return equalityFunc(a.(client.Object), b.(client.Object))
	}
	return sksets.NewResourceSet(genericSortFunc, genericEqualityFunc, genericResources...)
}

type authorizationPolicySet struct {
	set          sksets.ResourceSet
	sortFunc     func(toInsert, existing client.Object) bool
	equalityFunc func(a, b client.Object) bool
}

func NewAuthorizationPolicySet(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	authorizationPolicyList ...*security_istio_io_v1beta1.AuthorizationPolicy,
) AuthorizationPolicySet {
	return &authorizationPolicySet{
		set:          makeGenericAuthorizationPolicySet(sortFunc, equalityFunc, authorizationPolicyList),
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
	}
}

func NewAuthorizationPolicySetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	authorizationPolicyList *security_istio_io_v1beta1.AuthorizationPolicyList,
) AuthorizationPolicySet {
	list := make([]*security_istio_io_v1beta1.AuthorizationPolicy, 0, len(authorizationPolicyList.Items))
	for idx := range authorizationPolicyList.Items {
		list = append(list, authorizationPolicyList.Items[idx])
	}
	return &authorizationPolicySet{
		set:          makeGenericAuthorizationPolicySet(sortFunc, equalityFunc, list),
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
	}
}

func (s *authorizationPolicySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *authorizationPolicySet) List(filterResource ...func(*security_istio_io_v1beta1.AuthorizationPolicy) bool) []*security_istio_io_v1beta1.AuthorizationPolicy {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*security_istio_io_v1beta1.AuthorizationPolicy))
		})
	}

	objs := s.Generic().List(genericFilters...)
	authorizationPolicyList := make([]*security_istio_io_v1beta1.AuthorizationPolicy, 0, len(objs))
	for _, obj := range objs {
		authorizationPolicyList = append(authorizationPolicyList, obj.(*security_istio_io_v1beta1.AuthorizationPolicy))
	}
	return authorizationPolicyList
}

func (s *authorizationPolicySet) UnsortedList(filterResource ...func(*security_istio_io_v1beta1.AuthorizationPolicy) bool) []*security_istio_io_v1beta1.AuthorizationPolicy {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*security_istio_io_v1beta1.AuthorizationPolicy))
		})
	}

	var authorizationPolicyList []*security_istio_io_v1beta1.AuthorizationPolicy
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		authorizationPolicyList = append(authorizationPolicyList, obj.(*security_istio_io_v1beta1.AuthorizationPolicy))
	}
	return authorizationPolicyList
}

func (s *authorizationPolicySet) Map() map[string]*security_istio_io_v1beta1.AuthorizationPolicy {
	if s == nil {
		return nil
	}

	newMap := map[string]*security_istio_io_v1beta1.AuthorizationPolicy{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*security_istio_io_v1beta1.AuthorizationPolicy)
	}
	return newMap
}

func (s *authorizationPolicySet) Insert(
	authorizationPolicyList ...*security_istio_io_v1beta1.AuthorizationPolicy,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range authorizationPolicyList {
		s.Generic().Insert(obj)
	}
}

func (s *authorizationPolicySet) Has(authorizationPolicy ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(authorizationPolicy)
}

func (s *authorizationPolicySet) Equal(
	authorizationPolicySet AuthorizationPolicySet,
) bool {
	if s == nil {
		return authorizationPolicySet == nil
	}
	return s.Generic().Equal(authorizationPolicySet.Generic())
}

func (s *authorizationPolicySet) Delete(AuthorizationPolicy ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(AuthorizationPolicy)
}

func (s *authorizationPolicySet) Union(set AuthorizationPolicySet) AuthorizationPolicySet {
	if s == nil {
		return set
	}
	return NewAuthorizationPolicySet(s.sortFunc, s.equalityFunc, append(s.List(), set.List()...)...)
}

func (s *authorizationPolicySet) Difference(set AuthorizationPolicySet) AuthorizationPolicySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &authorizationPolicySet{
		set:          newSet,
		sortFunc:     s.sortFunc,
		equalityFunc: s.equalityFunc,
	}
}

func (s *authorizationPolicySet) Intersection(set AuthorizationPolicySet) AuthorizationPolicySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var authorizationPolicyList []*security_istio_io_v1beta1.AuthorizationPolicy
	for _, obj := range newSet.List() {
		authorizationPolicyList = append(authorizationPolicyList, obj.(*security_istio_io_v1beta1.AuthorizationPolicy))
	}
	return NewAuthorizationPolicySet(s.sortFunc, s.equalityFunc, authorizationPolicyList...)
}

func (s *authorizationPolicySet) Find(id ezkube.ResourceId) (*security_istio_io_v1beta1.AuthorizationPolicy, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find AuthorizationPolicy %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&security_istio_io_v1beta1.AuthorizationPolicy{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*security_istio_io_v1beta1.AuthorizationPolicy), nil
}

func (s *authorizationPolicySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *authorizationPolicySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *authorizationPolicySet) Delta(newSet AuthorizationPolicySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *authorizationPolicySet) Clone() AuthorizationPolicySet {
	if s == nil {
		return nil
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	genericEqualityFunc := func(a, b ezkube.ResourceId) bool {
		return s.equalityFunc(a.(client.Object), b.(client.Object))
	}
	return &authorizationPolicySet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			genericEqualityFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *authorizationPolicySet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}

func (s *authorizationPolicySet) GetEqualityFunc() func(a, b client.Object) bool {
	return s.equalityFunc
}

type PeerAuthenticationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*security_istio_io_v1beta1.PeerAuthentication) bool) []*security_istio_io_v1beta1.PeerAuthentication
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*security_istio_io_v1beta1.PeerAuthentication) bool) []*security_istio_io_v1beta1.PeerAuthentication
	// Return the Set as a map of key to resource.
	Map() map[string]*security_istio_io_v1beta1.PeerAuthentication
	// Insert a resource into the set.
	Insert(peerAuthentication ...*security_istio_io_v1beta1.PeerAuthentication)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(peerAuthenticationSet PeerAuthenticationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(peerAuthentication ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(peerAuthentication ezkube.ResourceId)
	// Return the union with the provided set
	Union(set PeerAuthenticationSet) PeerAuthenticationSet
	// Return the difference with the provided set
	Difference(set PeerAuthenticationSet) PeerAuthenticationSet
	// Return the intersection with the provided set
	Intersection(set PeerAuthenticationSet) PeerAuthenticationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*security_istio_io_v1beta1.PeerAuthentication, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another PeerAuthenticationSet
	Delta(newSet PeerAuthenticationSet) sksets.ResourceDelta
	// Create a deep copy of the current PeerAuthenticationSet
	Clone() PeerAuthenticationSet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
	// Get the equality function used by the set
	GetEqualityFunc() func(a, b client.Object) bool
}

func makeGenericPeerAuthenticationSet(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	peerAuthenticationList []*security_istio_io_v1beta1.PeerAuthentication,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range peerAuthenticationList {
		genericResources = append(genericResources, obj)
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	genericEqualityFunc := func(a, b ezkube.ResourceId) bool {
		return equalityFunc(a.(client.Object), b.(client.Object))
	}
	return sksets.NewResourceSet(genericSortFunc, genericEqualityFunc, genericResources...)
}

type peerAuthenticationSet struct {
	set          sksets.ResourceSet
	sortFunc     func(toInsert, existing client.Object) bool
	equalityFunc func(a, b client.Object) bool
}

func NewPeerAuthenticationSet(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	peerAuthenticationList ...*security_istio_io_v1beta1.PeerAuthentication,
) PeerAuthenticationSet {
	return &peerAuthenticationSet{
		set:          makeGenericPeerAuthenticationSet(sortFunc, equalityFunc, peerAuthenticationList),
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
	}
}

func NewPeerAuthenticationSetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	peerAuthenticationList *security_istio_io_v1beta1.PeerAuthenticationList,
) PeerAuthenticationSet {
	list := make([]*security_istio_io_v1beta1.PeerAuthentication, 0, len(peerAuthenticationList.Items))
	for idx := range peerAuthenticationList.Items {
		list = append(list, peerAuthenticationList.Items[idx])
	}
	return &peerAuthenticationSet{
		set:          makeGenericPeerAuthenticationSet(sortFunc, equalityFunc, list),
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
	}
}

func (s *peerAuthenticationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *peerAuthenticationSet) List(filterResource ...func(*security_istio_io_v1beta1.PeerAuthentication) bool) []*security_istio_io_v1beta1.PeerAuthentication {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*security_istio_io_v1beta1.PeerAuthentication))
		})
	}

	objs := s.Generic().List(genericFilters...)
	peerAuthenticationList := make([]*security_istio_io_v1beta1.PeerAuthentication, 0, len(objs))
	for _, obj := range objs {
		peerAuthenticationList = append(peerAuthenticationList, obj.(*security_istio_io_v1beta1.PeerAuthentication))
	}
	return peerAuthenticationList
}

func (s *peerAuthenticationSet) UnsortedList(filterResource ...func(*security_istio_io_v1beta1.PeerAuthentication) bool) []*security_istio_io_v1beta1.PeerAuthentication {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*security_istio_io_v1beta1.PeerAuthentication))
		})
	}

	var peerAuthenticationList []*security_istio_io_v1beta1.PeerAuthentication
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		peerAuthenticationList = append(peerAuthenticationList, obj.(*security_istio_io_v1beta1.PeerAuthentication))
	}
	return peerAuthenticationList
}

func (s *peerAuthenticationSet) Map() map[string]*security_istio_io_v1beta1.PeerAuthentication {
	if s == nil {
		return nil
	}

	newMap := map[string]*security_istio_io_v1beta1.PeerAuthentication{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*security_istio_io_v1beta1.PeerAuthentication)
	}
	return newMap
}

func (s *peerAuthenticationSet) Insert(
	peerAuthenticationList ...*security_istio_io_v1beta1.PeerAuthentication,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range peerAuthenticationList {
		s.Generic().Insert(obj)
	}
}

func (s *peerAuthenticationSet) Has(peerAuthentication ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(peerAuthentication)
}

func (s *peerAuthenticationSet) Equal(
	peerAuthenticationSet PeerAuthenticationSet,
) bool {
	if s == nil {
		return peerAuthenticationSet == nil
	}
	return s.Generic().Equal(peerAuthenticationSet.Generic())
}

func (s *peerAuthenticationSet) Delete(PeerAuthentication ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(PeerAuthentication)
}

func (s *peerAuthenticationSet) Union(set PeerAuthenticationSet) PeerAuthenticationSet {
	if s == nil {
		return set
	}
	return NewPeerAuthenticationSet(s.sortFunc, s.equalityFunc, append(s.List(), set.List()...)...)
}

func (s *peerAuthenticationSet) Difference(set PeerAuthenticationSet) PeerAuthenticationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &peerAuthenticationSet{
		set:          newSet,
		sortFunc:     s.sortFunc,
		equalityFunc: s.equalityFunc,
	}
}

func (s *peerAuthenticationSet) Intersection(set PeerAuthenticationSet) PeerAuthenticationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var peerAuthenticationList []*security_istio_io_v1beta1.PeerAuthentication
	for _, obj := range newSet.List() {
		peerAuthenticationList = append(peerAuthenticationList, obj.(*security_istio_io_v1beta1.PeerAuthentication))
	}
	return NewPeerAuthenticationSet(s.sortFunc, s.equalityFunc, peerAuthenticationList...)
}

func (s *peerAuthenticationSet) Find(id ezkube.ResourceId) (*security_istio_io_v1beta1.PeerAuthentication, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find PeerAuthentication %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&security_istio_io_v1beta1.PeerAuthentication{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*security_istio_io_v1beta1.PeerAuthentication), nil
}

func (s *peerAuthenticationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *peerAuthenticationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *peerAuthenticationSet) Delta(newSet PeerAuthenticationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *peerAuthenticationSet) Clone() PeerAuthenticationSet {
	if s == nil {
		return nil
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	genericEqualityFunc := func(a, b ezkube.ResourceId) bool {
		return s.equalityFunc(a.(client.Object), b.(client.Object))
	}
	return &peerAuthenticationSet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			genericEqualityFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *peerAuthenticationSet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}

func (s *peerAuthenticationSet) GetEqualityFunc() func(a, b client.Object) bool {
	return s.equalityFunc
}

type RequestAuthenticationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*security_istio_io_v1beta1.RequestAuthentication) bool) []*security_istio_io_v1beta1.RequestAuthentication
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*security_istio_io_v1beta1.RequestAuthentication) bool) []*security_istio_io_v1beta1.RequestAuthentication
	// Return the Set as a map of key to resource.
	Map() map[string]*security_istio_io_v1beta1.RequestAuthentication
	// Insert a resource into the set.
	Insert(requestAuthentication ...*security_istio_io_v1beta1.RequestAuthentication)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(requestAuthenticationSet RequestAuthenticationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(requestAuthentication ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(requestAuthentication ezkube.ResourceId)
	// Return the union with the provided set
	Union(set RequestAuthenticationSet) RequestAuthenticationSet
	// Return the difference with the provided set
	Difference(set RequestAuthenticationSet) RequestAuthenticationSet
	// Return the intersection with the provided set
	Intersection(set RequestAuthenticationSet) RequestAuthenticationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*security_istio_io_v1beta1.RequestAuthentication, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another RequestAuthenticationSet
	Delta(newSet RequestAuthenticationSet) sksets.ResourceDelta
	// Create a deep copy of the current RequestAuthenticationSet
	Clone() RequestAuthenticationSet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
	// Get the equality function used by the set
	GetEqualityFunc() func(a, b client.Object) bool
}

func makeGenericRequestAuthenticationSet(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	requestAuthenticationList []*security_istio_io_v1beta1.RequestAuthentication,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range requestAuthenticationList {
		genericResources = append(genericResources, obj)
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	genericEqualityFunc := func(a, b ezkube.ResourceId) bool {
		return equalityFunc(a.(client.Object), b.(client.Object))
	}
	return sksets.NewResourceSet(genericSortFunc, genericEqualityFunc, genericResources...)
}

type requestAuthenticationSet struct {
	set          sksets.ResourceSet
	sortFunc     func(toInsert, existing client.Object) bool
	equalityFunc func(a, b client.Object) bool
}

func NewRequestAuthenticationSet(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	requestAuthenticationList ...*security_istio_io_v1beta1.RequestAuthentication,
) RequestAuthenticationSet {
	return &requestAuthenticationSet{
		set:          makeGenericRequestAuthenticationSet(sortFunc, equalityFunc, requestAuthenticationList),
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
	}
}

func NewRequestAuthenticationSetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	requestAuthenticationList *security_istio_io_v1beta1.RequestAuthenticationList,
) RequestAuthenticationSet {
	list := make([]*security_istio_io_v1beta1.RequestAuthentication, 0, len(requestAuthenticationList.Items))
	for idx := range requestAuthenticationList.Items {
		list = append(list, requestAuthenticationList.Items[idx])
	}
	return &requestAuthenticationSet{
		set:          makeGenericRequestAuthenticationSet(sortFunc, equalityFunc, list),
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
	}
}

func (s *requestAuthenticationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *requestAuthenticationSet) List(filterResource ...func(*security_istio_io_v1beta1.RequestAuthentication) bool) []*security_istio_io_v1beta1.RequestAuthentication {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*security_istio_io_v1beta1.RequestAuthentication))
		})
	}

	objs := s.Generic().List(genericFilters...)
	requestAuthenticationList := make([]*security_istio_io_v1beta1.RequestAuthentication, 0, len(objs))
	for _, obj := range objs {
		requestAuthenticationList = append(requestAuthenticationList, obj.(*security_istio_io_v1beta1.RequestAuthentication))
	}
	return requestAuthenticationList
}

func (s *requestAuthenticationSet) UnsortedList(filterResource ...func(*security_istio_io_v1beta1.RequestAuthentication) bool) []*security_istio_io_v1beta1.RequestAuthentication {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*security_istio_io_v1beta1.RequestAuthentication))
		})
	}

	var requestAuthenticationList []*security_istio_io_v1beta1.RequestAuthentication
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		requestAuthenticationList = append(requestAuthenticationList, obj.(*security_istio_io_v1beta1.RequestAuthentication))
	}
	return requestAuthenticationList
}

func (s *requestAuthenticationSet) Map() map[string]*security_istio_io_v1beta1.RequestAuthentication {
	if s == nil {
		return nil
	}

	newMap := map[string]*security_istio_io_v1beta1.RequestAuthentication{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*security_istio_io_v1beta1.RequestAuthentication)
	}
	return newMap
}

func (s *requestAuthenticationSet) Insert(
	requestAuthenticationList ...*security_istio_io_v1beta1.RequestAuthentication,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range requestAuthenticationList {
		s.Generic().Insert(obj)
	}
}

func (s *requestAuthenticationSet) Has(requestAuthentication ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(requestAuthentication)
}

func (s *requestAuthenticationSet) Equal(
	requestAuthenticationSet RequestAuthenticationSet,
) bool {
	if s == nil {
		return requestAuthenticationSet == nil
	}
	return s.Generic().Equal(requestAuthenticationSet.Generic())
}

func (s *requestAuthenticationSet) Delete(RequestAuthentication ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(RequestAuthentication)
}

func (s *requestAuthenticationSet) Union(set RequestAuthenticationSet) RequestAuthenticationSet {
	if s == nil {
		return set
	}
	return NewRequestAuthenticationSet(s.sortFunc, s.equalityFunc, append(s.List(), set.List()...)...)
}

func (s *requestAuthenticationSet) Difference(set RequestAuthenticationSet) RequestAuthenticationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &requestAuthenticationSet{
		set:          newSet,
		sortFunc:     s.sortFunc,
		equalityFunc: s.equalityFunc,
	}
}

func (s *requestAuthenticationSet) Intersection(set RequestAuthenticationSet) RequestAuthenticationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var requestAuthenticationList []*security_istio_io_v1beta1.RequestAuthentication
	for _, obj := range newSet.List() {
		requestAuthenticationList = append(requestAuthenticationList, obj.(*security_istio_io_v1beta1.RequestAuthentication))
	}
	return NewRequestAuthenticationSet(s.sortFunc, s.equalityFunc, requestAuthenticationList...)
}

func (s *requestAuthenticationSet) Find(id ezkube.ResourceId) (*security_istio_io_v1beta1.RequestAuthentication, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find RequestAuthentication %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&security_istio_io_v1beta1.RequestAuthentication{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*security_istio_io_v1beta1.RequestAuthentication), nil
}

func (s *requestAuthenticationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *requestAuthenticationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *requestAuthenticationSet) Delta(newSet RequestAuthenticationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *requestAuthenticationSet) Clone() RequestAuthenticationSet {
	if s == nil {
		return nil
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	genericEqualityFunc := func(a, b ezkube.ResourceId) bool {
		return s.equalityFunc(a.(client.Object), b.(client.Object))
	}
	return &requestAuthenticationSet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			genericEqualityFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *requestAuthenticationSet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}

func (s *requestAuthenticationSet) GetEqualityFunc() func(a, b client.Object) bool {
	return s.equalityFunc
}
