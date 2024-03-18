// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha2sets

import (
	split_smi_spec_io_v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type TrafficSplitSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*split_smi_spec_io_v1alpha2.TrafficSplit) bool) []*split_smi_spec_io_v1alpha2.TrafficSplit
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*split_smi_spec_io_v1alpha2.TrafficSplit) bool) []*split_smi_spec_io_v1alpha2.TrafficSplit
	// Return the Set as a map of key to resource.
	Map() map[string]*split_smi_spec_io_v1alpha2.TrafficSplit
	// Insert a resource into the set.
	Insert(trafficSplit ...*split_smi_spec_io_v1alpha2.TrafficSplit)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(trafficSplitSet TrafficSplitSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(trafficSplit ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(trafficSplit ezkube.ResourceId)
	// Return the union with the provided set
	Union(set TrafficSplitSet) TrafficSplitSet
	// Return the difference with the provided set
	Difference(set TrafficSplitSet) TrafficSplitSet
	// Return the intersection with the provided set
	Intersection(set TrafficSplitSet) TrafficSplitSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*split_smi_spec_io_v1alpha2.TrafficSplit, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another TrafficSplitSet
	Delta(newSet TrafficSplitSet) sksets.ResourceDelta
	// Create a deep copy of the current TrafficSplitSet
	Clone() TrafficSplitSet
}

func makeGenericTrafficSplitSet(trafficSplitList []*split_smi_spec_io_v1alpha2.TrafficSplit) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range trafficSplitList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type trafficSplitSet struct {
	set sksets.ResourceSet
}

func NewTrafficSplitSet(trafficSplitList ...*split_smi_spec_io_v1alpha2.TrafficSplit) TrafficSplitSet {
	return &trafficSplitSet{set: makeGenericTrafficSplitSet(trafficSplitList)}
}

func NewTrafficSplitSetFromList(trafficSplitList *split_smi_spec_io_v1alpha2.TrafficSplitList) TrafficSplitSet {
	list := make([]*split_smi_spec_io_v1alpha2.TrafficSplit, 0, len(trafficSplitList.Items))
	for idx := range trafficSplitList.Items {
		list = append(list, &trafficSplitList.Items[idx])
	}
	return &trafficSplitSet{set: makeGenericTrafficSplitSet(list)}
}

func (s *trafficSplitSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *trafficSplitSet) List(filterResource ...func(*split_smi_spec_io_v1alpha2.TrafficSplit) bool) []*split_smi_spec_io_v1alpha2.TrafficSplit {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
		})
	}

	objs := s.Generic().List(genericFilters...)
	trafficSplitList := make([]*split_smi_spec_io_v1alpha2.TrafficSplit, 0, len(objs))
	for _, obj := range objs {
		trafficSplitList = append(trafficSplitList, obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
	}
	return trafficSplitList
}

func (s *trafficSplitSet) UnsortedList(filterResource ...func(*split_smi_spec_io_v1alpha2.TrafficSplit) bool) []*split_smi_spec_io_v1alpha2.TrafficSplit {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
		})
	}

	var trafficSplitList []*split_smi_spec_io_v1alpha2.TrafficSplit
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		trafficSplitList = append(trafficSplitList, obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
	}
	return trafficSplitList
}

func (s *trafficSplitSet) Map() map[string]*split_smi_spec_io_v1alpha2.TrafficSplit {
	if s == nil {
		return nil
	}

	newMap := map[string]*split_smi_spec_io_v1alpha2.TrafficSplit{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*split_smi_spec_io_v1alpha2.TrafficSplit)
	}
	return newMap
}

func (s *trafficSplitSet) Insert(
	trafficSplitList ...*split_smi_spec_io_v1alpha2.TrafficSplit,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range trafficSplitList {
		s.Generic().Insert(obj)
	}
}

func (s *trafficSplitSet) Has(trafficSplit ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(trafficSplit)
}

func (s *trafficSplitSet) Equal(
	trafficSplitSet TrafficSplitSet,
) bool {
	if s == nil {
		return trafficSplitSet == nil
	}
	return s.Generic().Equal(trafficSplitSet.Generic())
}

func (s *trafficSplitSet) Delete(TrafficSplit ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(TrafficSplit)
}

func (s *trafficSplitSet) Union(set TrafficSplitSet) TrafficSplitSet {
	if s == nil {
		return set
	}
	return &trafficSplitMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *trafficSplitSet) Difference(set TrafficSplitSet) TrafficSplitSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &trafficSplitSet{set: newSet}
}

func (s *trafficSplitSet) Intersection(set TrafficSplitSet) TrafficSplitSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var trafficSplitList []*split_smi_spec_io_v1alpha2.TrafficSplit
	for _, obj := range newSet.List() {
		trafficSplitList = append(trafficSplitList, obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
	}
	return NewTrafficSplitSet(trafficSplitList...)
}

func (s *trafficSplitSet) Find(id ezkube.ResourceId) (*split_smi_spec_io_v1alpha2.TrafficSplit, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find TrafficSplit %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&split_smi_spec_io_v1alpha2.TrafficSplit{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*split_smi_spec_io_v1alpha2.TrafficSplit), nil
}

func (s *trafficSplitSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *trafficSplitSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *trafficSplitSet) Delta(newSet TrafficSplitSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *trafficSplitSet) Clone() TrafficSplitSet {
	if s == nil {
		return nil
	}
	return &trafficSplitMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type trafficSplitMergedSet struct {
	sets []sksets.ResourceSet
}

func NewTrafficSplitMergedSet(trafficSplitList ...*split_smi_spec_io_v1alpha2.TrafficSplit) TrafficSplitSet {
	return &trafficSplitMergedSet{sets: []sksets.ResourceSet{makeGenericTrafficSplitSet(trafficSplitList)}}
}

func NewTrafficSplitMergedSetFromList(trafficSplitList *split_smi_spec_io_v1alpha2.TrafficSplitList) TrafficSplitSet {
	list := make([]*split_smi_spec_io_v1alpha2.TrafficSplit, 0, len(trafficSplitList.Items))
	for idx := range trafficSplitList.Items {
		list = append(list, &trafficSplitList.Items[idx])
	}
	return &trafficSplitMergedSet{sets: []sksets.ResourceSet{makeGenericTrafficSplitSet(list)}}
}

func (s *trafficSplitMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	toRet := sets.String{}
	for _, set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *trafficSplitMergedSet) List(filterResource ...func(*split_smi_spec_io_v1alpha2.TrafficSplit) bool) []*split_smi_spec_io_v1alpha2.TrafficSplit {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
		})
	}
	trafficSplitList := []*split_smi_spec_io_v1alpha2.TrafficSplit{}
	for _, set := range s.sets {
		for _, obj := range set.List(genericFilters...) {
			trafficSplitList = append(trafficSplitList, obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
		}
	}
	return trafficSplitList
}

func (s *trafficSplitMergedSet) UnsortedList(filterResource ...func(*split_smi_spec_io_v1alpha2.TrafficSplit) bool) []*split_smi_spec_io_v1alpha2.TrafficSplit {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
		})
	}

	trafficSplitList := []*split_smi_spec_io_v1alpha2.TrafficSplit{}
	for _, set := range s.sets {
		for _, obj := range set.UnsortedList(genericFilters...) {
			trafficSplitList = append(trafficSplitList, obj.(*split_smi_spec_io_v1alpha2.TrafficSplit))
		}
	}
	return trafficSplitList
}

func (s *trafficSplitMergedSet) Map() map[string]*split_smi_spec_io_v1alpha2.TrafficSplit {
	if s == nil {
		return nil
	}

	newMap := map[string]*split_smi_spec_io_v1alpha2.TrafficSplit{}
	for _, set := range s.sets {
		for k, v := range set.Map() {
			newMap[k] = v.(*split_smi_spec_io_v1alpha2.TrafficSplit)
		}
	}
	return newMap
}

func (s *trafficSplitMergedSet) Insert(
	trafficSplitList ...*split_smi_spec_io_v1alpha2.TrafficSplit,
) {
	if s == nil {
	}
	if len(s.sets) == 0 {
		s.sets = append(s.sets, makeGenericTrafficSplitSet(trafficSplitList))
	}
	for _, obj := range trafficSplitList {
		s.sets[0].Insert(obj)
	}
}

func (s *trafficSplitMergedSet) Has(trafficSplit ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	for _, set := range s.sets {
		if set.Has(trafficSplit) {
			return true
		}
	}
	return false
}

func (s *trafficSplitMergedSet) Equal(
	trafficSplitSet TrafficSplitSet,
) bool {
	panic("unimplemented")
}

func (s *trafficSplitMergedSet) Delete(TrafficSplit ezkube.ResourceId) {
	for _, set := range s.sets {
		set.Delete(TrafficSplit)
	}
}

func (s *trafficSplitMergedSet) Union(set TrafficSplitSet) TrafficSplitSet {
	if s == nil {
		return set
	}
	return &trafficSplitMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *trafficSplitMergedSet) Difference(set TrafficSplitSet) TrafficSplitSet {
	panic("unimplemented")
}

func (s *trafficSplitMergedSet) Intersection(set TrafficSplitSet) TrafficSplitSet {
	panic("unimplemented")
}

func (s *trafficSplitMergedSet) Find(id ezkube.ResourceId) (*split_smi_spec_io_v1alpha2.TrafficSplit, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find TrafficSplit %v", sksets.Key(id))
	}

	var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&split_smi_spec_io_v1alpha2.TrafficSplit{}, id)
		if err == nil {
			return obj.(*split_smi_spec_io_v1alpha2.TrafficSplit), nil
		}
	}

	return nil, err
}

func (s *trafficSplitMergedSet) Length() int {
	if s == nil {
		return 0
	}
	totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
	return totalLen
}

func (s *trafficSplitMergedSet) Generic() sksets.ResourceSet {
	panic("unimplemented")
}

func (s *trafficSplitMergedSet) Delta(newSet TrafficSplitSet) sksets.ResourceDelta {
	panic("unimplemented")
}

func (s *trafficSplitMergedSet) Clone() TrafficSplitSet {
	if s == nil {
		return nil
	}
	return &trafficSplitMergedSet{sets: s.sets[:]}
}
