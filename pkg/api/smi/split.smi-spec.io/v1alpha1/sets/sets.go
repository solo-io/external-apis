// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	split_smi_spec_io_v1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type TrafficSplitSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*split_smi_spec_io_v1alpha1.TrafficSplit) bool) []*split_smi_spec_io_v1alpha1.TrafficSplit
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*split_smi_spec_io_v1alpha1.TrafficSplit) bool) []*split_smi_spec_io_v1alpha1.TrafficSplit
	// Return the Set as a map of key to resource.
	Map() map[string]*split_smi_spec_io_v1alpha1.TrafficSplit
	// Insert a resource into the set.
	Insert(trafficSplit ...*split_smi_spec_io_v1alpha1.TrafficSplit)
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
	Find(id ezkube.ResourceId) (*split_smi_spec_io_v1alpha1.TrafficSplit, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another TrafficSplitSet
	Delta(newSet TrafficSplitSet) sksets.ResourceDelta
	// Create a deep copy of the current TrafficSplitSet
	Clone() TrafficSplitSet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
	// Get the equality function used by the set
	GetEqualityFunc() func(a, b client.Object) bool
}

func makeGenericTrafficSplitSet(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	trafficSplitList []*split_smi_spec_io_v1alpha1.TrafficSplit,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range trafficSplitList {
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

type trafficSplitSet struct {
	set          sksets.ResourceSet
	sortFunc     func(toInsert, existing client.Object) bool
	equalityFunc func(a, b client.Object) bool
}

func NewTrafficSplitSet(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	trafficSplitList ...*split_smi_spec_io_v1alpha1.TrafficSplit,
) TrafficSplitSet {
	return &trafficSplitSet{
		set:          makeGenericTrafficSplitSet(sortFunc, equalityFunc, trafficSplitList),
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
	}
}

func NewTrafficSplitSetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	equalityFunc func(a, b client.Object) bool,
	trafficSplitList *split_smi_spec_io_v1alpha1.TrafficSplitList,
) TrafficSplitSet {
	list := make([]*split_smi_spec_io_v1alpha1.TrafficSplit, 0, len(trafficSplitList.Items))
	for idx := range trafficSplitList.Items {
		list = append(list, &trafficSplitList.Items[idx])
	}
	return &trafficSplitSet{
		set:          makeGenericTrafficSplitSet(sortFunc, equalityFunc, list),
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
	}
}

func (s *trafficSplitSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *trafficSplitSet) List(filterResource ...func(*split_smi_spec_io_v1alpha1.TrafficSplit) bool) []*split_smi_spec_io_v1alpha1.TrafficSplit {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*split_smi_spec_io_v1alpha1.TrafficSplit))
		})
	}

	objs := s.Generic().List(genericFilters...)
	trafficSplitList := make([]*split_smi_spec_io_v1alpha1.TrafficSplit, 0, len(objs))
	for _, obj := range objs {
		trafficSplitList = append(trafficSplitList, obj.(*split_smi_spec_io_v1alpha1.TrafficSplit))
	}
	return trafficSplitList
}

func (s *trafficSplitSet) UnsortedList(filterResource ...func(*split_smi_spec_io_v1alpha1.TrafficSplit) bool) []*split_smi_spec_io_v1alpha1.TrafficSplit {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*split_smi_spec_io_v1alpha1.TrafficSplit))
		})
	}

	var trafficSplitList []*split_smi_spec_io_v1alpha1.TrafficSplit
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		trafficSplitList = append(trafficSplitList, obj.(*split_smi_spec_io_v1alpha1.TrafficSplit))
	}
	return trafficSplitList
}

func (s *trafficSplitSet) Map() map[string]*split_smi_spec_io_v1alpha1.TrafficSplit {
	if s == nil {
		return nil
	}

	newMap := map[string]*split_smi_spec_io_v1alpha1.TrafficSplit{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*split_smi_spec_io_v1alpha1.TrafficSplit)
	}
	return newMap
}

func (s *trafficSplitSet) Insert(
	trafficSplitList ...*split_smi_spec_io_v1alpha1.TrafficSplit,
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
	return NewTrafficSplitSet(s.sortFunc, s.equalityFunc, append(s.List(), set.List()...)...)
}

func (s *trafficSplitSet) Difference(set TrafficSplitSet) TrafficSplitSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &trafficSplitSet{
		set:          newSet,
		sortFunc:     s.sortFunc,
		equalityFunc: s.equalityFunc,
	}
}

func (s *trafficSplitSet) Intersection(set TrafficSplitSet) TrafficSplitSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var trafficSplitList []*split_smi_spec_io_v1alpha1.TrafficSplit
	for _, obj := range newSet.List() {
		trafficSplitList = append(trafficSplitList, obj.(*split_smi_spec_io_v1alpha1.TrafficSplit))
	}
	return NewTrafficSplitSet(s.sortFunc, s.equalityFunc, trafficSplitList...)
}

func (s *trafficSplitSet) Find(id ezkube.ResourceId) (*split_smi_spec_io_v1alpha1.TrafficSplit, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find TrafficSplit %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&split_smi_spec_io_v1alpha1.TrafficSplit{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*split_smi_spec_io_v1alpha1.TrafficSplit), nil
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
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	genericEqualityFunc := func(a, b ezkube.ResourceId) bool {
		return s.equalityFunc(a.(client.Object), b.(client.Object))
	}
	return &trafficSplitSet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			genericEqualityFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *trafficSplitSet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}

func (s *trafficSplitSet) GetEqualityFunc() func(a, b client.Object) bool {
	return s.equalityFunc
}
