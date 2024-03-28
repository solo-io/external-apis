// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1beta1sets

import (
	apiextensions_k8s_io_v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type CustomResourceDefinitionSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) bool) []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) bool) []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition
	// Return the Set as a map of key to resource.
	Map() map[string]*apiextensions_k8s_io_v1beta1.CustomResourceDefinition
	// Insert a resource into the set.
	Insert(customResourceDefinition ...*apiextensions_k8s_io_v1beta1.CustomResourceDefinition)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(customResourceDefinitionSet CustomResourceDefinitionSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(customResourceDefinition ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(customResourceDefinition ezkube.ResourceId)
	// Return the union with the provided set
	Union(set CustomResourceDefinitionSet) CustomResourceDefinitionSet
	// Return the difference with the provided set
	Difference(set CustomResourceDefinitionSet) CustomResourceDefinitionSet
	// Return the intersection with the provided set
	Intersection(set CustomResourceDefinitionSet) CustomResourceDefinitionSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*apiextensions_k8s_io_v1beta1.CustomResourceDefinition, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another CustomResourceDefinitionSet
	Delta(newSet CustomResourceDefinitionSet) sksets.ResourceDelta
	// Create a deep copy of the current CustomResourceDefinitionSet
	Clone() CustomResourceDefinitionSet
}

func makeGenericCustomResourceDefinitionSet(customResourceDefinitionList []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range customResourceDefinitionList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type customResourceDefinitionSet struct {
	set sksets.ResourceSet
}

func NewCustomResourceDefinitionSet(customResourceDefinitionList ...*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) CustomResourceDefinitionSet {
	return &customResourceDefinitionSet{set: makeGenericCustomResourceDefinitionSet(customResourceDefinitionList)}
}

func NewCustomResourceDefinitionSetFromList(customResourceDefinitionList *apiextensions_k8s_io_v1beta1.CustomResourceDefinitionList) CustomResourceDefinitionSet {
	list := make([]*apiextensions_k8s_io_v1beta1.CustomResourceDefinition, 0, len(customResourceDefinitionList.Items))
	for idx := range customResourceDefinitionList.Items {
		list = append(list, &customResourceDefinitionList.Items[idx])
	}
	return &customResourceDefinitionSet{set: makeGenericCustomResourceDefinitionSet(list)}
}

func (s *customResourceDefinitionSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *customResourceDefinitionSet) List(filterResource ...func(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) bool) []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
		})
	}

	objs := s.Generic().List(genericFilters...)
	customResourceDefinitionList := make([]*apiextensions_k8s_io_v1beta1.CustomResourceDefinition, 0, len(objs))
	for _, obj := range objs {
		customResourceDefinitionList = append(customResourceDefinitionList, obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
	}
	return customResourceDefinitionList
}

func (s *customResourceDefinitionSet) UnsortedList(filterResource ...func(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) bool) []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
		})
	}

	var customResourceDefinitionList []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		customResourceDefinitionList = append(customResourceDefinitionList, obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
	}
	return customResourceDefinitionList
}

func (s *customResourceDefinitionSet) Map() map[string]*apiextensions_k8s_io_v1beta1.CustomResourceDefinition {
	if s == nil {
		return nil
	}

	newMap := map[string]*apiextensions_k8s_io_v1beta1.CustomResourceDefinition{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition)
	}
	return newMap
}

func (s *customResourceDefinitionSet) Insert(
	customResourceDefinitionList ...*apiextensions_k8s_io_v1beta1.CustomResourceDefinition,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range customResourceDefinitionList {
		s.Generic().Insert(obj)
	}
}

func (s *customResourceDefinitionSet) Has(customResourceDefinition ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(customResourceDefinition)
}

func (s *customResourceDefinitionSet) Equal(
	customResourceDefinitionSet CustomResourceDefinitionSet,
) bool {
	if s == nil {
		return customResourceDefinitionSet == nil
	}
	return s.Generic().Equal(customResourceDefinitionSet.Generic())
}

func (s *customResourceDefinitionSet) Delete(CustomResourceDefinition ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(CustomResourceDefinition)
}

func (s *customResourceDefinitionSet) Union(set CustomResourceDefinitionSet) CustomResourceDefinitionSet {
	if s == nil {
		return set
	}
	return NewCustomResourceDefinitionSet(append(s.List(), set.List()...)...)
}

func (s *customResourceDefinitionSet) Difference(set CustomResourceDefinitionSet) CustomResourceDefinitionSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &customResourceDefinitionSet{set: newSet}
}

func (s *customResourceDefinitionSet) Intersection(set CustomResourceDefinitionSet) CustomResourceDefinitionSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var customResourceDefinitionList []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition
	for _, obj := range newSet.List() {
		customResourceDefinitionList = append(customResourceDefinitionList, obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
	}
	return NewCustomResourceDefinitionSet(customResourceDefinitionList...)
}

func (s *customResourceDefinitionSet) Find(id ezkube.ResourceId) (*apiextensions_k8s_io_v1beta1.CustomResourceDefinition, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find CustomResourceDefinition %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&apiextensions_k8s_io_v1beta1.CustomResourceDefinition{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition), nil
}

func (s *customResourceDefinitionSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *customResourceDefinitionSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *customResourceDefinitionSet) Delta(newSet CustomResourceDefinitionSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *customResourceDefinitionSet) Clone() CustomResourceDefinitionSet {
	if s == nil {
		return nil
	}
	return &customResourceDefinitionMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type customResourceDefinitionMergedSet struct {
	sets []sksets.ResourceSet
}

func NewCustomResourceDefinitionMergedSet(customResourceDefinitionList ...*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) CustomResourceDefinitionSet {
	return &customResourceDefinitionMergedSet{sets: []sksets.ResourceSet{makeGenericCustomResourceDefinitionSet(customResourceDefinitionList)}}
}

func NewCustomResourceDefinitionMergedSetFromList(customResourceDefinitionList *apiextensions_k8s_io_v1beta1.CustomResourceDefinitionList) CustomResourceDefinitionSet {
	list := make([]*apiextensions_k8s_io_v1beta1.CustomResourceDefinition, 0, len(customResourceDefinitionList.Items))
	for idx := range customResourceDefinitionList.Items {
		list = append(list, &customResourceDefinitionList.Items[idx])
	}
	return &customResourceDefinitionMergedSet{sets: []sksets.ResourceSet{makeGenericCustomResourceDefinitionSet(list)}}
}

func (s *customResourceDefinitionMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	toRet := sets.String{}
	for _, set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *customResourceDefinitionMergedSet) List(filterResource ...func(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) bool) []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
		})
	}
	customResourceDefinitionList := []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.List(genericFilters...) {
			if tracker[obj] {
				continue
			}
			tracker[obj] = true
			customResourceDefinitionList = append(customResourceDefinitionList, obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
		}
	}
	return customResourceDefinitionList
}

func (s *customResourceDefinitionMergedSet) UnsortedList(filterResource ...func(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition) bool) []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
		})
	}
	customResourceDefinitionList := []*apiextensions_k8s_io_v1beta1.CustomResourceDefinition{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.UnsortedList(genericFilters...) {
			if tracker[obj] {
				continue
			}
			tracker[obj] = true
			customResourceDefinitionList = append(customResourceDefinitionList, obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition))
		}
	}
	return customResourceDefinitionList
}

func (s *customResourceDefinitionMergedSet) Map() map[string]*apiextensions_k8s_io_v1beta1.CustomResourceDefinition {
	if s == nil {
		return nil
	}

	newMap := map[string]*apiextensions_k8s_io_v1beta1.CustomResourceDefinition{}
	for _, set := range s.sets {
		for k, v := range set.Map() {
			newMap[k] = v.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition)
		}
	}
	return newMap
}

func (s *customResourceDefinitionMergedSet) Insert(
	customResourceDefinitionList ...*apiextensions_k8s_io_v1beta1.CustomResourceDefinition,
) {
	if s == nil {
	}
	if len(s.sets) == 0 {
		s.sets = append(s.sets, makeGenericCustomResourceDefinitionSet(customResourceDefinitionList))
	}
	for _, obj := range customResourceDefinitionList {
		inserted := false
		for _, set := range s.sets {
			if set.Has(obj) {
				set.Insert(obj)
				inserted = true
				break
			}
		}
		if !inserted {
			s.sets[0].Insert(obj)
		}
	}
}

func (s *customResourceDefinitionMergedSet) Has(customResourceDefinition ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	for _, set := range s.sets {
		if set.Has(customResourceDefinition) {
			return true
		}
	}
	return false
}

func (s *customResourceDefinitionMergedSet) Equal(
	customResourceDefinitionSet CustomResourceDefinitionSet,
) bool {
	panic("unimplemented")
}

func (s *customResourceDefinitionMergedSet) Delete(CustomResourceDefinition ezkube.ResourceId) {
	for _, set := range s.sets {
		set.Delete(CustomResourceDefinition)
	}
}

func (s *customResourceDefinitionMergedSet) Union(set CustomResourceDefinitionSet) CustomResourceDefinitionSet {
	if s == nil {
		return set
	}
	return &customResourceDefinitionMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *customResourceDefinitionMergedSet) Difference(set CustomResourceDefinitionSet) CustomResourceDefinitionSet {
	panic("unimplemented")
}

func (s *customResourceDefinitionMergedSet) Intersection(set CustomResourceDefinitionSet) CustomResourceDefinitionSet {
	panic("unimplemented")
}

func (s *customResourceDefinitionMergedSet) Find(id ezkube.ResourceId) (*apiextensions_k8s_io_v1beta1.CustomResourceDefinition, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find CustomResourceDefinition %v", sksets.Key(id))
	}

	var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&apiextensions_k8s_io_v1beta1.CustomResourceDefinition{}, id)
		if err == nil {
			return obj.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition), nil
		}
	}

	return nil, err
}

func (s *customResourceDefinitionMergedSet) Length() int {
	if s == nil {
		return 0
	}
	totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
	return totalLen
}

func (s *customResourceDefinitionMergedSet) Generic() sksets.ResourceSet {
	panic("unimplemented")
}

func (s *customResourceDefinitionMergedSet) Delta(newSet CustomResourceDefinitionSet) sksets.ResourceDelta {
	panic("unimplemented")
}

func (s *customResourceDefinitionMergedSet) Clone() CustomResourceDefinitionSet {
	if s == nil {
		return nil
	}
	return &customResourceDefinitionMergedSet{sets: s.sets[:]}
}
