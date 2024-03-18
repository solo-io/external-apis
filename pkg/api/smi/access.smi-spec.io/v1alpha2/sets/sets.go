// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha2sets

import (
	access_smi_spec_io_v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type TrafficTargetSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*access_smi_spec_io_v1alpha2.TrafficTarget) bool) []*access_smi_spec_io_v1alpha2.TrafficTarget
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*access_smi_spec_io_v1alpha2.TrafficTarget) bool) []*access_smi_spec_io_v1alpha2.TrafficTarget
	// Return the Set as a map of key to resource.
	Map() map[string]*access_smi_spec_io_v1alpha2.TrafficTarget
	// Insert a resource into the set.
	Insert(trafficTarget ...*access_smi_spec_io_v1alpha2.TrafficTarget)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(trafficTargetSet TrafficTargetSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(trafficTarget ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(trafficTarget ezkube.ResourceId)
	// Return the union with the provided set
	Union(set TrafficTargetSet) TrafficTargetSet
	// Return the difference with the provided set
	Difference(set TrafficTargetSet) TrafficTargetSet
	// Return the intersection with the provided set
	Intersection(set TrafficTargetSet) TrafficTargetSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*access_smi_spec_io_v1alpha2.TrafficTarget, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another TrafficTargetSet
	Delta(newSet TrafficTargetSet) sksets.ResourceDelta
	// Create a deep copy of the current TrafficTargetSet
	Clone() TrafficTargetSet
}

func makeGenericTrafficTargetSet(trafficTargetList []*access_smi_spec_io_v1alpha2.TrafficTarget) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range trafficTargetList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type trafficTargetSet struct {
	set sksets.ResourceSet
}

func NewTrafficTargetSet(trafficTargetList ...*access_smi_spec_io_v1alpha2.TrafficTarget) TrafficTargetSet {
	return &trafficTargetSet{set: makeGenericTrafficTargetSet(trafficTargetList)}
}

func NewTrafficTargetSetFromList(trafficTargetList *access_smi_spec_io_v1alpha2.TrafficTargetList) TrafficTargetSet {
	list := make([]*access_smi_spec_io_v1alpha2.TrafficTarget, 0, len(trafficTargetList.Items))
	for idx := range trafficTargetList.Items {
		list = append(list, &trafficTargetList.Items[idx])
	}
	return &trafficTargetSet{set: makeGenericTrafficTargetSet(list)}
}

func (s *trafficTargetSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *trafficTargetSet) List(filterResource ...func(*access_smi_spec_io_v1alpha2.TrafficTarget) bool) []*access_smi_spec_io_v1alpha2.TrafficTarget {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
		})
	}

	objs := s.Generic().List(genericFilters...)
	trafficTargetList := make([]*access_smi_spec_io_v1alpha2.TrafficTarget, 0, len(objs))
	for _, obj := range objs {
		trafficTargetList = append(trafficTargetList, obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
	}
	return trafficTargetList
}

func (s *trafficTargetSet) UnsortedList(filterResource ...func(*access_smi_spec_io_v1alpha2.TrafficTarget) bool) []*access_smi_spec_io_v1alpha2.TrafficTarget {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
		})
	}

	var trafficTargetList []*access_smi_spec_io_v1alpha2.TrafficTarget
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		trafficTargetList = append(trafficTargetList, obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
	}
	return trafficTargetList
}

func (s *trafficTargetSet) Map() map[string]*access_smi_spec_io_v1alpha2.TrafficTarget {
	if s == nil {
		return nil
	}

	newMap := map[string]*access_smi_spec_io_v1alpha2.TrafficTarget{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*access_smi_spec_io_v1alpha2.TrafficTarget)
	}
	return newMap
}

func (s *trafficTargetSet) Insert(
	trafficTargetList ...*access_smi_spec_io_v1alpha2.TrafficTarget,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range trafficTargetList {
		s.Generic().Insert(obj)
	}
}

func (s *trafficTargetSet) Has(trafficTarget ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(trafficTarget)
}

func (s *trafficTargetSet) Equal(
	trafficTargetSet TrafficTargetSet,
) bool {
	if s == nil {
		return trafficTargetSet == nil
	}
	return s.Generic().Equal(trafficTargetSet.Generic())
}

func (s *trafficTargetSet) Delete(TrafficTarget ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(TrafficTarget)
}

func (s *trafficTargetSet) Union(set TrafficTargetSet) TrafficTargetSet {
	if s == nil {
		return set
	}
	return &trafficTargetMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *trafficTargetSet) Difference(set TrafficTargetSet) TrafficTargetSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &trafficTargetSet{set: newSet}
}

func (s *trafficTargetSet) Intersection(set TrafficTargetSet) TrafficTargetSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var trafficTargetList []*access_smi_spec_io_v1alpha2.TrafficTarget
	for _, obj := range newSet.List() {
		trafficTargetList = append(trafficTargetList, obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
	}
	return NewTrafficTargetSet(trafficTargetList...)
}

func (s *trafficTargetSet) Find(id ezkube.ResourceId) (*access_smi_spec_io_v1alpha2.TrafficTarget, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find TrafficTarget %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&access_smi_spec_io_v1alpha2.TrafficTarget{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*access_smi_spec_io_v1alpha2.TrafficTarget), nil
}

func (s *trafficTargetSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *trafficTargetSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *trafficTargetSet) Delta(newSet TrafficTargetSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *trafficTargetSet) Clone() TrafficTargetSet {
	if s == nil {
		return nil
	}
	return &trafficTargetMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type trafficTargetMergedSet struct {
	sets []sksets.ResourceSet
}

func NewTrafficTargetMergedSet(trafficTargetList ...*access_smi_spec_io_v1alpha2.TrafficTarget) TrafficTargetSet {
	return &trafficTargetMergedSet{sets: []sksets.ResourceSet{makeGenericTrafficTargetSet(trafficTargetList)}}
}

func NewTrafficTargetMergedSetFromList(trafficTargetList *access_smi_spec_io_v1alpha2.TrafficTargetList) TrafficTargetSet {
	list := make([]*access_smi_spec_io_v1alpha2.TrafficTarget, 0, len(trafficTargetList.Items))
	for idx := range trafficTargetList.Items {
		list = append(list, &trafficTargetList.Items[idx])
	}
	return &trafficTargetMergedSet{sets: []sksets.ResourceSet{makeGenericTrafficTargetSet(list)}}
}

func (s *trafficTargetMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	toRet := sets.String{}
	for _, set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *trafficTargetMergedSet) List(filterResource ...func(*access_smi_spec_io_v1alpha2.TrafficTarget) bool) []*access_smi_spec_io_v1alpha2.TrafficTarget {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
		})
	}
	trafficTargetList := []*access_smi_spec_io_v1alpha2.TrafficTarget{}
	for _, set := range s.sets {
		for _, obj := range set.List(genericFilters...) {
			trafficTargetList = append(trafficTargetList, obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
		}
	}
	return trafficTargetList
}

func (s *trafficTargetMergedSet) UnsortedList(filterResource ...func(*access_smi_spec_io_v1alpha2.TrafficTarget) bool) []*access_smi_spec_io_v1alpha2.TrafficTarget {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
		})
	}

	trafficTargetList := []*access_smi_spec_io_v1alpha2.TrafficTarget{}
	for _, set := range s.sets {
		for _, obj := range set.UnsortedList(genericFilters...) {
			trafficTargetList = append(trafficTargetList, obj.(*access_smi_spec_io_v1alpha2.TrafficTarget))
		}
	}
	return trafficTargetList
}

func (s *trafficTargetMergedSet) Map() map[string]*access_smi_spec_io_v1alpha2.TrafficTarget {
	if s == nil {
		return nil
	}

	newMap := map[string]*access_smi_spec_io_v1alpha2.TrafficTarget{}
	for _, set := range s.sets {
		for k, v := range set.Map() {
			newMap[k] = v.(*access_smi_spec_io_v1alpha2.TrafficTarget)
		}
	}
	return newMap
}

func (s *trafficTargetMergedSet) Insert(
	trafficTargetList ...*access_smi_spec_io_v1alpha2.TrafficTarget,
) {
	if s == nil {
	}
	if len(s.sets) == 0 {
		s.sets = append(s.sets, makeGenericTrafficTargetSet(trafficTargetList))
	}
	for _, obj := range trafficTargetList {
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

func (s *trafficTargetMergedSet) Has(trafficTarget ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	for _, set := range s.sets {
		if set.Has(trafficTarget) {
			return true
		}
	}
	return false
}

func (s *trafficTargetMergedSet) Equal(
	trafficTargetSet TrafficTargetSet,
) bool {
	panic("unimplemented")
}

func (s *trafficTargetMergedSet) Delete(TrafficTarget ezkube.ResourceId) {
	for _, set := range s.sets {
		set.Delete(TrafficTarget)
	}
}

func (s *trafficTargetMergedSet) Union(set TrafficTargetSet) TrafficTargetSet {
	if s == nil {
		return set
	}
	return &trafficTargetMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *trafficTargetMergedSet) Difference(set TrafficTargetSet) TrafficTargetSet {
	panic("unimplemented")
}

func (s *trafficTargetMergedSet) Intersection(set TrafficTargetSet) TrafficTargetSet {
	panic("unimplemented")
}

func (s *trafficTargetMergedSet) Find(id ezkube.ResourceId) (*access_smi_spec_io_v1alpha2.TrafficTarget, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find TrafficTarget %v", sksets.Key(id))
	}

	var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&access_smi_spec_io_v1alpha2.TrafficTarget{}, id)
		if err == nil {
			return obj.(*access_smi_spec_io_v1alpha2.TrafficTarget), nil
		}
	}

	return nil, err
}

func (s *trafficTargetMergedSet) Length() int {
	if s == nil {
		return 0
	}
	totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
	return totalLen
}

func (s *trafficTargetMergedSet) Generic() sksets.ResourceSet {
	panic("unimplemented")
}

func (s *trafficTargetMergedSet) Delta(newSet TrafficTargetSet) sksets.ResourceDelta {
	panic("unimplemented")
}

func (s *trafficTargetMergedSet) Clone() TrafficTargetSet {
	if s == nil {
		return nil
	}
	return &trafficTargetMergedSet{sets: s.sets[:]}
}
