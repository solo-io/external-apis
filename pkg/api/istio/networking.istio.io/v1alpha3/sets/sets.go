// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha3sets

import (
	networking_istio_io_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type DestinationRuleSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_istio_io_v1alpha3.DestinationRule) bool) []*networking_istio_io_v1alpha3.DestinationRule
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_istio_io_v1alpha3.DestinationRule
	// Insert a resource into the set.
	Insert(destinationRule ...*networking_istio_io_v1alpha3.DestinationRule)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(destinationRuleSet DestinationRuleSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(destinationRule ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(destinationRule ezkube.ResourceId)
	// Return the union with the provided set
	Union(set DestinationRuleSet) DestinationRuleSet
	// Return the difference with the provided set
	Difference(set DestinationRuleSet) DestinationRuleSet
	// Return the intersection with the provided set
	Intersection(set DestinationRuleSet) DestinationRuleSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.DestinationRule, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another DestinationRuleSet
	Delta(newSet DestinationRuleSet) sksets.ResourceDelta
}

func makeGenericDestinationRuleSet(destinationRuleList []*networking_istio_io_v1alpha3.DestinationRule) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range destinationRuleList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type destinationRuleSet struct {
	set sksets.ResourceSet
}

func NewDestinationRuleSet(destinationRuleList ...*networking_istio_io_v1alpha3.DestinationRule) DestinationRuleSet {
	return &destinationRuleSet{set: makeGenericDestinationRuleSet(destinationRuleList)}
}

func NewDestinationRuleSetFromList(destinationRuleList *networking_istio_io_v1alpha3.DestinationRuleList) DestinationRuleSet {
	list := make([]*networking_istio_io_v1alpha3.DestinationRule, 0, len(destinationRuleList.Items))
	for idx := range destinationRuleList.Items {
		list = append(list, &destinationRuleList.Items[idx])
	}
	return &destinationRuleSet{set: makeGenericDestinationRuleSet(list)}
}

func (s *destinationRuleSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *destinationRuleSet) List(filterResource ...func(*networking_istio_io_v1alpha3.DestinationRule) bool) []*networking_istio_io_v1alpha3.DestinationRule {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_istio_io_v1alpha3.DestinationRule))
		})
	}

	var destinationRuleList []*networking_istio_io_v1alpha3.DestinationRule
	for _, obj := range s.Generic().List(genericFilters...) {
		destinationRuleList = append(destinationRuleList, obj.(*networking_istio_io_v1alpha3.DestinationRule))
	}
	return destinationRuleList
}

func (s *destinationRuleSet) Map() map[string]*networking_istio_io_v1alpha3.DestinationRule {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_istio_io_v1alpha3.DestinationRule{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_istio_io_v1alpha3.DestinationRule)
	}
	return newMap
}

func (s *destinationRuleSet) Insert(
	destinationRuleList ...*networking_istio_io_v1alpha3.DestinationRule,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range destinationRuleList {
		s.Generic().Insert(obj)
	}
}

func (s *destinationRuleSet) Has(destinationRule ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(destinationRule)
}

func (s *destinationRuleSet) Equal(
	destinationRuleSet DestinationRuleSet,
) bool {
	if s == nil {
		return destinationRuleSet == nil
	}
	return s.Generic().Equal(destinationRuleSet.Generic())
}

func (s *destinationRuleSet) Delete(DestinationRule ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(DestinationRule)
}

func (s *destinationRuleSet) Union(set DestinationRuleSet) DestinationRuleSet {
	if s == nil {
		return set
	}
	return NewDestinationRuleSet(append(s.List(), set.List()...)...)
}

func (s *destinationRuleSet) Difference(set DestinationRuleSet) DestinationRuleSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &destinationRuleSet{set: newSet}
}

func (s *destinationRuleSet) Intersection(set DestinationRuleSet) DestinationRuleSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var destinationRuleList []*networking_istio_io_v1alpha3.DestinationRule
	for _, obj := range newSet.List() {
		destinationRuleList = append(destinationRuleList, obj.(*networking_istio_io_v1alpha3.DestinationRule))
	}
	return NewDestinationRuleSet(destinationRuleList...)
}

func (s *destinationRuleSet) Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.DestinationRule, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find DestinationRule %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_istio_io_v1alpha3.DestinationRule{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_istio_io_v1alpha3.DestinationRule), nil
}

func (s *destinationRuleSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *destinationRuleSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *destinationRuleSet) Delta(newSet DestinationRuleSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type EnvoyFilterSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_istio_io_v1alpha3.EnvoyFilter) bool) []*networking_istio_io_v1alpha3.EnvoyFilter
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_istio_io_v1alpha3.EnvoyFilter
	// Insert a resource into the set.
	Insert(envoyFilter ...*networking_istio_io_v1alpha3.EnvoyFilter)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(envoyFilterSet EnvoyFilterSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(envoyFilter ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(envoyFilter ezkube.ResourceId)
	// Return the union with the provided set
	Union(set EnvoyFilterSet) EnvoyFilterSet
	// Return the difference with the provided set
	Difference(set EnvoyFilterSet) EnvoyFilterSet
	// Return the intersection with the provided set
	Intersection(set EnvoyFilterSet) EnvoyFilterSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.EnvoyFilter, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another EnvoyFilterSet
	Delta(newSet EnvoyFilterSet) sksets.ResourceDelta
}

func makeGenericEnvoyFilterSet(envoyFilterList []*networking_istio_io_v1alpha3.EnvoyFilter) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range envoyFilterList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type envoyFilterSet struct {
	set sksets.ResourceSet
}

func NewEnvoyFilterSet(envoyFilterList ...*networking_istio_io_v1alpha3.EnvoyFilter) EnvoyFilterSet {
	return &envoyFilterSet{set: makeGenericEnvoyFilterSet(envoyFilterList)}
}

func NewEnvoyFilterSetFromList(envoyFilterList *networking_istio_io_v1alpha3.EnvoyFilterList) EnvoyFilterSet {
	list := make([]*networking_istio_io_v1alpha3.EnvoyFilter, 0, len(envoyFilterList.Items))
	for idx := range envoyFilterList.Items {
		list = append(list, &envoyFilterList.Items[idx])
	}
	return &envoyFilterSet{set: makeGenericEnvoyFilterSet(list)}
}

func (s *envoyFilterSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *envoyFilterSet) List(filterResource ...func(*networking_istio_io_v1alpha3.EnvoyFilter) bool) []*networking_istio_io_v1alpha3.EnvoyFilter {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_istio_io_v1alpha3.EnvoyFilter))
		})
	}

	var envoyFilterList []*networking_istio_io_v1alpha3.EnvoyFilter
	for _, obj := range s.Generic().List(genericFilters...) {
		envoyFilterList = append(envoyFilterList, obj.(*networking_istio_io_v1alpha3.EnvoyFilter))
	}
	return envoyFilterList
}

func (s *envoyFilterSet) Map() map[string]*networking_istio_io_v1alpha3.EnvoyFilter {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_istio_io_v1alpha3.EnvoyFilter{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_istio_io_v1alpha3.EnvoyFilter)
	}
	return newMap
}

func (s *envoyFilterSet) Insert(
	envoyFilterList ...*networking_istio_io_v1alpha3.EnvoyFilter,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range envoyFilterList {
		s.Generic().Insert(obj)
	}
}

func (s *envoyFilterSet) Has(envoyFilter ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(envoyFilter)
}

func (s *envoyFilterSet) Equal(
	envoyFilterSet EnvoyFilterSet,
) bool {
	if s == nil {
		return envoyFilterSet == nil
	}
	return s.Generic().Equal(envoyFilterSet.Generic())
}

func (s *envoyFilterSet) Delete(EnvoyFilter ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(EnvoyFilter)
}

func (s *envoyFilterSet) Union(set EnvoyFilterSet) EnvoyFilterSet {
	if s == nil {
		return set
	}
	return NewEnvoyFilterSet(append(s.List(), set.List()...)...)
}

func (s *envoyFilterSet) Difference(set EnvoyFilterSet) EnvoyFilterSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &envoyFilterSet{set: newSet}
}

func (s *envoyFilterSet) Intersection(set EnvoyFilterSet) EnvoyFilterSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var envoyFilterList []*networking_istio_io_v1alpha3.EnvoyFilter
	for _, obj := range newSet.List() {
		envoyFilterList = append(envoyFilterList, obj.(*networking_istio_io_v1alpha3.EnvoyFilter))
	}
	return NewEnvoyFilterSet(envoyFilterList...)
}

func (s *envoyFilterSet) Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.EnvoyFilter, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find EnvoyFilter %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_istio_io_v1alpha3.EnvoyFilter{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_istio_io_v1alpha3.EnvoyFilter), nil
}

func (s *envoyFilterSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *envoyFilterSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *envoyFilterSet) Delta(newSet EnvoyFilterSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type GatewaySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_istio_io_v1alpha3.Gateway) bool) []*networking_istio_io_v1alpha3.Gateway
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_istio_io_v1alpha3.Gateway
	// Insert a resource into the set.
	Insert(gateway ...*networking_istio_io_v1alpha3.Gateway)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(gatewaySet GatewaySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(gateway ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(gateway ezkube.ResourceId)
	// Return the union with the provided set
	Union(set GatewaySet) GatewaySet
	// Return the difference with the provided set
	Difference(set GatewaySet) GatewaySet
	// Return the intersection with the provided set
	Intersection(set GatewaySet) GatewaySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.Gateway, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another GatewaySet
	Delta(newSet GatewaySet) sksets.ResourceDelta
}

func makeGenericGatewaySet(gatewayList []*networking_istio_io_v1alpha3.Gateway) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range gatewayList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type gatewaySet struct {
	set sksets.ResourceSet
}

func NewGatewaySet(gatewayList ...*networking_istio_io_v1alpha3.Gateway) GatewaySet {
	return &gatewaySet{set: makeGenericGatewaySet(gatewayList)}
}

func NewGatewaySetFromList(gatewayList *networking_istio_io_v1alpha3.GatewayList) GatewaySet {
	list := make([]*networking_istio_io_v1alpha3.Gateway, 0, len(gatewayList.Items))
	for idx := range gatewayList.Items {
		list = append(list, &gatewayList.Items[idx])
	}
	return &gatewaySet{set: makeGenericGatewaySet(list)}
}

func (s *gatewaySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *gatewaySet) List(filterResource ...func(*networking_istio_io_v1alpha3.Gateway) bool) []*networking_istio_io_v1alpha3.Gateway {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_istio_io_v1alpha3.Gateway))
		})
	}

	var gatewayList []*networking_istio_io_v1alpha3.Gateway
	for _, obj := range s.Generic().List(genericFilters...) {
		gatewayList = append(gatewayList, obj.(*networking_istio_io_v1alpha3.Gateway))
	}
	return gatewayList
}

func (s *gatewaySet) Map() map[string]*networking_istio_io_v1alpha3.Gateway {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_istio_io_v1alpha3.Gateway{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_istio_io_v1alpha3.Gateway)
	}
	return newMap
}

func (s *gatewaySet) Insert(
	gatewayList ...*networking_istio_io_v1alpha3.Gateway,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range gatewayList {
		s.Generic().Insert(obj)
	}
}

func (s *gatewaySet) Has(gateway ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(gateway)
}

func (s *gatewaySet) Equal(
	gatewaySet GatewaySet,
) bool {
	if s == nil {
		return gatewaySet == nil
	}
	return s.Generic().Equal(gatewaySet.Generic())
}

func (s *gatewaySet) Delete(Gateway ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Gateway)
}

func (s *gatewaySet) Union(set GatewaySet) GatewaySet {
	if s == nil {
		return set
	}
	return NewGatewaySet(append(s.List(), set.List()...)...)
}

func (s *gatewaySet) Difference(set GatewaySet) GatewaySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &gatewaySet{set: newSet}
}

func (s *gatewaySet) Intersection(set GatewaySet) GatewaySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var gatewayList []*networking_istio_io_v1alpha3.Gateway
	for _, obj := range newSet.List() {
		gatewayList = append(gatewayList, obj.(*networking_istio_io_v1alpha3.Gateway))
	}
	return NewGatewaySet(gatewayList...)
}

func (s *gatewaySet) Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.Gateway, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Gateway %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_istio_io_v1alpha3.Gateway{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_istio_io_v1alpha3.Gateway), nil
}

func (s *gatewaySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *gatewaySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *gatewaySet) Delta(newSet GatewaySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type ServiceEntrySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_istio_io_v1alpha3.ServiceEntry) bool) []*networking_istio_io_v1alpha3.ServiceEntry
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_istio_io_v1alpha3.ServiceEntry
	// Insert a resource into the set.
	Insert(serviceEntry ...*networking_istio_io_v1alpha3.ServiceEntry)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(serviceEntrySet ServiceEntrySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(serviceEntry ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(serviceEntry ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ServiceEntrySet) ServiceEntrySet
	// Return the difference with the provided set
	Difference(set ServiceEntrySet) ServiceEntrySet
	// Return the intersection with the provided set
	Intersection(set ServiceEntrySet) ServiceEntrySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.ServiceEntry, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ServiceEntrySet
	Delta(newSet ServiceEntrySet) sksets.ResourceDelta
}

func makeGenericServiceEntrySet(serviceEntryList []*networking_istio_io_v1alpha3.ServiceEntry) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range serviceEntryList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type serviceEntrySet struct {
	set sksets.ResourceSet
}

func NewServiceEntrySet(serviceEntryList ...*networking_istio_io_v1alpha3.ServiceEntry) ServiceEntrySet {
	return &serviceEntrySet{set: makeGenericServiceEntrySet(serviceEntryList)}
}

func NewServiceEntrySetFromList(serviceEntryList *networking_istio_io_v1alpha3.ServiceEntryList) ServiceEntrySet {
	list := make([]*networking_istio_io_v1alpha3.ServiceEntry, 0, len(serviceEntryList.Items))
	for idx := range serviceEntryList.Items {
		list = append(list, &serviceEntryList.Items[idx])
	}
	return &serviceEntrySet{set: makeGenericServiceEntrySet(list)}
}

func (s *serviceEntrySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *serviceEntrySet) List(filterResource ...func(*networking_istio_io_v1alpha3.ServiceEntry) bool) []*networking_istio_io_v1alpha3.ServiceEntry {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_istio_io_v1alpha3.ServiceEntry))
		})
	}

	var serviceEntryList []*networking_istio_io_v1alpha3.ServiceEntry
	for _, obj := range s.Generic().List(genericFilters...) {
		serviceEntryList = append(serviceEntryList, obj.(*networking_istio_io_v1alpha3.ServiceEntry))
	}
	return serviceEntryList
}

func (s *serviceEntrySet) Map() map[string]*networking_istio_io_v1alpha3.ServiceEntry {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_istio_io_v1alpha3.ServiceEntry{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_istio_io_v1alpha3.ServiceEntry)
	}
	return newMap
}

func (s *serviceEntrySet) Insert(
	serviceEntryList ...*networking_istio_io_v1alpha3.ServiceEntry,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range serviceEntryList {
		s.Generic().Insert(obj)
	}
}

func (s *serviceEntrySet) Has(serviceEntry ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(serviceEntry)
}

func (s *serviceEntrySet) Equal(
	serviceEntrySet ServiceEntrySet,
) bool {
	if s == nil {
		return serviceEntrySet == nil
	}
	return s.Generic().Equal(serviceEntrySet.Generic())
}

func (s *serviceEntrySet) Delete(ServiceEntry ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ServiceEntry)
}

func (s *serviceEntrySet) Union(set ServiceEntrySet) ServiceEntrySet {
	if s == nil {
		return set
	}
	return NewServiceEntrySet(append(s.List(), set.List()...)...)
}

func (s *serviceEntrySet) Difference(set ServiceEntrySet) ServiceEntrySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &serviceEntrySet{set: newSet}
}

func (s *serviceEntrySet) Intersection(set ServiceEntrySet) ServiceEntrySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var serviceEntryList []*networking_istio_io_v1alpha3.ServiceEntry
	for _, obj := range newSet.List() {
		serviceEntryList = append(serviceEntryList, obj.(*networking_istio_io_v1alpha3.ServiceEntry))
	}
	return NewServiceEntrySet(serviceEntryList...)
}

func (s *serviceEntrySet) Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.ServiceEntry, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ServiceEntry %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_istio_io_v1alpha3.ServiceEntry{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_istio_io_v1alpha3.ServiceEntry), nil
}

func (s *serviceEntrySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *serviceEntrySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *serviceEntrySet) Delta(newSet ServiceEntrySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type VirtualServiceSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_istio_io_v1alpha3.VirtualService) bool) []*networking_istio_io_v1alpha3.VirtualService
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_istio_io_v1alpha3.VirtualService
	// Insert a resource into the set.
	Insert(virtualService ...*networking_istio_io_v1alpha3.VirtualService)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualServiceSet VirtualServiceSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualService ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualService ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualServiceSet) VirtualServiceSet
	// Return the difference with the provided set
	Difference(set VirtualServiceSet) VirtualServiceSet
	// Return the intersection with the provided set
	Intersection(set VirtualServiceSet) VirtualServiceSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.VirtualService, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualServiceSet
	Delta(newSet VirtualServiceSet) sksets.ResourceDelta
}

func makeGenericVirtualServiceSet(virtualServiceList []*networking_istio_io_v1alpha3.VirtualService) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualServiceList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualServiceSet struct {
	set sksets.ResourceSet
}

func NewVirtualServiceSet(virtualServiceList ...*networking_istio_io_v1alpha3.VirtualService) VirtualServiceSet {
	return &virtualServiceSet{set: makeGenericVirtualServiceSet(virtualServiceList)}
}

func NewVirtualServiceSetFromList(virtualServiceList *networking_istio_io_v1alpha3.VirtualServiceList) VirtualServiceSet {
	list := make([]*networking_istio_io_v1alpha3.VirtualService, 0, len(virtualServiceList.Items))
	for idx := range virtualServiceList.Items {
		list = append(list, &virtualServiceList.Items[idx])
	}
	return &virtualServiceSet{set: makeGenericVirtualServiceSet(list)}
}

func (s *virtualServiceSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualServiceSet) List(filterResource ...func(*networking_istio_io_v1alpha3.VirtualService) bool) []*networking_istio_io_v1alpha3.VirtualService {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_istio_io_v1alpha3.VirtualService))
		})
	}

	var virtualServiceList []*networking_istio_io_v1alpha3.VirtualService
	for _, obj := range s.Generic().List(genericFilters...) {
		virtualServiceList = append(virtualServiceList, obj.(*networking_istio_io_v1alpha3.VirtualService))
	}
	return virtualServiceList
}

func (s *virtualServiceSet) Map() map[string]*networking_istio_io_v1alpha3.VirtualService {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_istio_io_v1alpha3.VirtualService{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_istio_io_v1alpha3.VirtualService)
	}
	return newMap
}

func (s *virtualServiceSet) Insert(
	virtualServiceList ...*networking_istio_io_v1alpha3.VirtualService,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualServiceList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualServiceSet) Has(virtualService ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualService)
}

func (s *virtualServiceSet) Equal(
	virtualServiceSet VirtualServiceSet,
) bool {
	if s == nil {
		return virtualServiceSet == nil
	}
	return s.Generic().Equal(virtualServiceSet.Generic())
}

func (s *virtualServiceSet) Delete(VirtualService ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualService)
}

func (s *virtualServiceSet) Union(set VirtualServiceSet) VirtualServiceSet {
	if s == nil {
		return set
	}
	return NewVirtualServiceSet(append(s.List(), set.List()...)...)
}

func (s *virtualServiceSet) Difference(set VirtualServiceSet) VirtualServiceSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualServiceSet{set: newSet}
}

func (s *virtualServiceSet) Intersection(set VirtualServiceSet) VirtualServiceSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualServiceList []*networking_istio_io_v1alpha3.VirtualService
	for _, obj := range newSet.List() {
		virtualServiceList = append(virtualServiceList, obj.(*networking_istio_io_v1alpha3.VirtualService))
	}
	return NewVirtualServiceSet(virtualServiceList...)
}

func (s *virtualServiceSet) Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.VirtualService, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualService %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_istio_io_v1alpha3.VirtualService{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_istio_io_v1alpha3.VirtualService), nil
}

func (s *virtualServiceSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualServiceSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualServiceSet) Delta(newSet VirtualServiceSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type SidecarSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_istio_io_v1alpha3.Sidecar) bool) []*networking_istio_io_v1alpha3.Sidecar
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_istio_io_v1alpha3.Sidecar
	// Insert a resource into the set.
	Insert(sidecar ...*networking_istio_io_v1alpha3.Sidecar)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(sidecarSet SidecarSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(sidecar ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(sidecar ezkube.ResourceId)
	// Return the union with the provided set
	Union(set SidecarSet) SidecarSet
	// Return the difference with the provided set
	Difference(set SidecarSet) SidecarSet
	// Return the intersection with the provided set
	Intersection(set SidecarSet) SidecarSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.Sidecar, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another SidecarSet
	Delta(newSet SidecarSet) sksets.ResourceDelta
}

func makeGenericSidecarSet(sidecarList []*networking_istio_io_v1alpha3.Sidecar) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range sidecarList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type sidecarSet struct {
	set sksets.ResourceSet
}

func NewSidecarSet(sidecarList ...*networking_istio_io_v1alpha3.Sidecar) SidecarSet {
	return &sidecarSet{set: makeGenericSidecarSet(sidecarList)}
}

func NewSidecarSetFromList(sidecarList *networking_istio_io_v1alpha3.SidecarList) SidecarSet {
	list := make([]*networking_istio_io_v1alpha3.Sidecar, 0, len(sidecarList.Items))
	for idx := range sidecarList.Items {
		list = append(list, &sidecarList.Items[idx])
	}
	return &sidecarSet{set: makeGenericSidecarSet(list)}
}

func (s *sidecarSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *sidecarSet) List(filterResource ...func(*networking_istio_io_v1alpha3.Sidecar) bool) []*networking_istio_io_v1alpha3.Sidecar {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_istio_io_v1alpha3.Sidecar))
		})
	}

	var sidecarList []*networking_istio_io_v1alpha3.Sidecar
	for _, obj := range s.Generic().List(genericFilters...) {
		sidecarList = append(sidecarList, obj.(*networking_istio_io_v1alpha3.Sidecar))
	}
	return sidecarList
}

func (s *sidecarSet) Map() map[string]*networking_istio_io_v1alpha3.Sidecar {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_istio_io_v1alpha3.Sidecar{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_istio_io_v1alpha3.Sidecar)
	}
	return newMap
}

func (s *sidecarSet) Insert(
	sidecarList ...*networking_istio_io_v1alpha3.Sidecar,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range sidecarList {
		s.Generic().Insert(obj)
	}
}

func (s *sidecarSet) Has(sidecar ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(sidecar)
}

func (s *sidecarSet) Equal(
	sidecarSet SidecarSet,
) bool {
	if s == nil {
		return sidecarSet == nil
	}
	return s.Generic().Equal(sidecarSet.Generic())
}

func (s *sidecarSet) Delete(Sidecar ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Sidecar)
}

func (s *sidecarSet) Union(set SidecarSet) SidecarSet {
	if s == nil {
		return set
	}
	return NewSidecarSet(append(s.List(), set.List()...)...)
}

func (s *sidecarSet) Difference(set SidecarSet) SidecarSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &sidecarSet{set: newSet}
}

func (s *sidecarSet) Intersection(set SidecarSet) SidecarSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var sidecarList []*networking_istio_io_v1alpha3.Sidecar
	for _, obj := range newSet.List() {
		sidecarList = append(sidecarList, obj.(*networking_istio_io_v1alpha3.Sidecar))
	}
	return NewSidecarSet(sidecarList...)
}

func (s *sidecarSet) Find(id ezkube.ResourceId) (*networking_istio_io_v1alpha3.Sidecar, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Sidecar %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_istio_io_v1alpha3.Sidecar{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_istio_io_v1alpha3.Sidecar), nil
}

func (s *sidecarSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *sidecarSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *sidecarSet) Delta(newSet SidecarSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}
