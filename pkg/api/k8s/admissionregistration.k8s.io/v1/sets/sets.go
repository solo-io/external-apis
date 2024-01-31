// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	admissionregistration_k8s_io_v1 "k8s.io/api/admissionregistration/v1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type ValidatingWebhookConfigurationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) bool) []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) bool) []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	// Return the Set as a map of key to resource.
	Map() map[string]*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	// Insert a resource into the set.
	Insert(validatingWebhookConfiguration ...*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(validatingWebhookConfigurationSet ValidatingWebhookConfigurationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(validatingWebhookConfiguration ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(validatingWebhookConfiguration ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet
	// Return the difference with the provided set
	Difference(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet
	// Return the intersection with the provided set
	Intersection(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ValidatingWebhookConfigurationSet
	Delta(newSet ValidatingWebhookConfigurationSet) sksets.ResourceDelta
	// Create a deep copy of the current ValidatingWebhookConfigurationSet
	Clone() ValidatingWebhookConfigurationSet
}

func makeGenericValidatingWebhookConfigurationSet(validatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range validatingWebhookConfigurationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type validatingWebhookConfigurationSet struct {
	set sksets.ResourceSet
}

func NewValidatingWebhookConfigurationSet(validatingWebhookConfigurationList ...*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) ValidatingWebhookConfigurationSet {
	return &validatingWebhookConfigurationSet{set: makeGenericValidatingWebhookConfigurationSet(validatingWebhookConfigurationList)}
}

func NewValidatingWebhookConfigurationSetFromList(validatingWebhookConfigurationList *admissionregistration_k8s_io_v1.ValidatingWebhookConfigurationList) ValidatingWebhookConfigurationSet {
	list := make([]*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, 0, len(validatingWebhookConfigurationList.Items))
	for idx := range validatingWebhookConfigurationList.Items {
		list = append(list, &validatingWebhookConfigurationList.Items[idx])
	}
	return &validatingWebhookConfigurationSet{set: makeGenericValidatingWebhookConfigurationSet(list)}
}

func (s *validatingWebhookConfigurationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *validatingWebhookConfigurationSet) List(filterResource ...func(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) bool) []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration))
		})
	}

	objs := s.Generic().List(genericFilters...)
	validatingWebhookConfigurationList := make([]*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, 0, len(objs))
	for _, obj := range objs {
		validatingWebhookConfigurationList = append(validatingWebhookConfigurationList, obj.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration))
	}
	return validatingWebhookConfigurationList
}

func (s *validatingWebhookConfigurationSet) UnsortedList(filterResource ...func(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) bool) []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration))
		})
	}

	var validatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		validatingWebhookConfigurationList = append(validatingWebhookConfigurationList, obj.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration))
	}
	return validatingWebhookConfigurationList
}

func (s *validatingWebhookConfigurationSet) Map() map[string]*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration {
	if s == nil {
		return nil
	}

	newMap := map[string]*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)
	}
	return newMap
}

func (s *validatingWebhookConfigurationSet) Insert(
	validatingWebhookConfigurationList ...*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range validatingWebhookConfigurationList {
		s.Generic().Insert(obj)
	}
}

func (s *validatingWebhookConfigurationSet) Has(validatingWebhookConfiguration ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(validatingWebhookConfiguration)
}

func (s *validatingWebhookConfigurationSet) Equal(
	validatingWebhookConfigurationSet ValidatingWebhookConfigurationSet,
) bool {
	if s == nil {
		return validatingWebhookConfigurationSet == nil
	}
	return s.Generic().Equal(validatingWebhookConfigurationSet.Generic())
}

func (s *validatingWebhookConfigurationSet) Delete(ValidatingWebhookConfiguration ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ValidatingWebhookConfiguration)
}

func (s *validatingWebhookConfigurationSet) Union(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet {
	if s == nil {
		return set
	}
	return NewValidatingWebhookConfigurationSet(append(s.List(), set.List()...)...)
}

func (s *validatingWebhookConfigurationSet) Difference(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &validatingWebhookConfigurationSet{set: newSet}
}

func (s *validatingWebhookConfigurationSet) Intersection(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var validatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	for _, obj := range newSet.List() {
		validatingWebhookConfigurationList = append(validatingWebhookConfigurationList, obj.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration))
	}
	return NewValidatingWebhookConfigurationSet(validatingWebhookConfigurationList...)
}

func (s *validatingWebhookConfigurationSet) Find(id ezkube.ResourceId) (*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ValidatingWebhookConfiguration %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration), nil
}

func (s *validatingWebhookConfigurationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *validatingWebhookConfigurationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *validatingWebhookConfigurationSet) Delta(newSet ValidatingWebhookConfigurationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *validatingWebhookConfigurationSet) Clone() ValidatingWebhookConfigurationSet {
	if s == nil {
		return nil
	}
	return &validatingWebhookConfigurationSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type MutatingWebhookConfigurationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) bool) []*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) bool) []*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration
	// Return the Set as a map of key to resource.
	Map() map[string]*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration
	// Insert a resource into the set.
	Insert(mutatingWebhookConfiguration ...*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(mutatingWebhookConfigurationSet MutatingWebhookConfigurationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(mutatingWebhookConfiguration ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(mutatingWebhookConfiguration ezkube.ResourceId)
	// Return the union with the provided set
	Union(set MutatingWebhookConfigurationSet) MutatingWebhookConfigurationSet
	// Return the difference with the provided set
	Difference(set MutatingWebhookConfigurationSet) MutatingWebhookConfigurationSet
	// Return the intersection with the provided set
	Intersection(set MutatingWebhookConfigurationSet) MutatingWebhookConfigurationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another MutatingWebhookConfigurationSet
	Delta(newSet MutatingWebhookConfigurationSet) sksets.ResourceDelta
	// Create a deep copy of the current MutatingWebhookConfigurationSet
	Clone() MutatingWebhookConfigurationSet
}

func makeGenericMutatingWebhookConfigurationSet(mutatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range mutatingWebhookConfigurationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type mutatingWebhookConfigurationSet struct {
	set sksets.ResourceSet
}

func NewMutatingWebhookConfigurationSet(mutatingWebhookConfigurationList ...*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) MutatingWebhookConfigurationSet {
	return &mutatingWebhookConfigurationSet{set: makeGenericMutatingWebhookConfigurationSet(mutatingWebhookConfigurationList)}
}

func NewMutatingWebhookConfigurationSetFromList(mutatingWebhookConfigurationList *admissionregistration_k8s_io_v1.MutatingWebhookConfigurationList) MutatingWebhookConfigurationSet {
	list := make([]*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration, 0, len(mutatingWebhookConfigurationList.Items))
	for idx := range mutatingWebhookConfigurationList.Items {
		list = append(list, &mutatingWebhookConfigurationList.Items[idx])
	}
	return &mutatingWebhookConfigurationSet{set: makeGenericMutatingWebhookConfigurationSet(list)}
}

func (s *mutatingWebhookConfigurationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *mutatingWebhookConfigurationSet) List(filterResource ...func(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) bool) []*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration))
		})
	}

	objs := s.Generic().List(genericFilters...)
	mutatingWebhookConfigurationList := make([]*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration, 0, len(objs))
	for _, obj := range objs {
		mutatingWebhookConfigurationList = append(mutatingWebhookConfigurationList, obj.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration))
	}
	return mutatingWebhookConfigurationList
}

func (s *mutatingWebhookConfigurationSet) UnsortedList(filterResource ...func(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) bool) []*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration))
		})
	}

	var mutatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		mutatingWebhookConfigurationList = append(mutatingWebhookConfigurationList, obj.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration))
	}
	return mutatingWebhookConfigurationList
}

func (s *mutatingWebhookConfigurationSet) Map() map[string]*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration {
	if s == nil {
		return nil
	}

	newMap := map[string]*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration)
	}
	return newMap
}

func (s *mutatingWebhookConfigurationSet) Insert(
	mutatingWebhookConfigurationList ...*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range mutatingWebhookConfigurationList {
		s.Generic().Insert(obj)
	}
}

func (s *mutatingWebhookConfigurationSet) Has(mutatingWebhookConfiguration ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(mutatingWebhookConfiguration)
}

func (s *mutatingWebhookConfigurationSet) Equal(
	mutatingWebhookConfigurationSet MutatingWebhookConfigurationSet,
) bool {
	if s == nil {
		return mutatingWebhookConfigurationSet == nil
	}
	return s.Generic().Equal(mutatingWebhookConfigurationSet.Generic())
}

func (s *mutatingWebhookConfigurationSet) Delete(MutatingWebhookConfiguration ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(MutatingWebhookConfiguration)
}

func (s *mutatingWebhookConfigurationSet) Union(set MutatingWebhookConfigurationSet) MutatingWebhookConfigurationSet {
	if s == nil {
		return set
	}
	return NewMutatingWebhookConfigurationSet(append(s.List(), set.List()...)...)
}

func (s *mutatingWebhookConfigurationSet) Difference(set MutatingWebhookConfigurationSet) MutatingWebhookConfigurationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &mutatingWebhookConfigurationSet{set: newSet}
}

func (s *mutatingWebhookConfigurationSet) Intersection(set MutatingWebhookConfigurationSet) MutatingWebhookConfigurationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var mutatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration
	for _, obj := range newSet.List() {
		mutatingWebhookConfigurationList = append(mutatingWebhookConfigurationList, obj.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration))
	}
	return NewMutatingWebhookConfigurationSet(mutatingWebhookConfigurationList...)
}

func (s *mutatingWebhookConfigurationSet) Find(id ezkube.ResourceId) (*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find MutatingWebhookConfiguration %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&admissionregistration_k8s_io_v1.MutatingWebhookConfiguration{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration), nil
}

func (s *mutatingWebhookConfigurationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *mutatingWebhookConfigurationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *mutatingWebhookConfigurationSet) Delta(newSet MutatingWebhookConfigurationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *mutatingWebhookConfigurationSet) Clone() MutatingWebhookConfigurationSet {
	if s == nil {
		return nil
	}
	return &mutatingWebhookConfigurationSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
