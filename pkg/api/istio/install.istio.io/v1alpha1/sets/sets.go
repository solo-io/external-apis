// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	install_istio_io_v1alpha1 "istio.io/istio/operator/pkg/apis/istio/v1alpha1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type IstioOperatorSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*install_istio_io_v1alpha1.IstioOperator) bool) []*install_istio_io_v1alpha1.IstioOperator
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*install_istio_io_v1alpha1.IstioOperator) bool) []*install_istio_io_v1alpha1.IstioOperator
	// Return the Set as a map of key to resource.
	Map() map[string]*install_istio_io_v1alpha1.IstioOperator
	// Insert a resource into the set.
	Insert(istioOperator ...*install_istio_io_v1alpha1.IstioOperator)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(istioOperatorSet IstioOperatorSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(istioOperator ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(istioOperator ezkube.ResourceId)
	// Return the union with the provided set
	Union(set IstioOperatorSet) IstioOperatorSet
	// Return the difference with the provided set
	Difference(set IstioOperatorSet) IstioOperatorSet
	// Return the intersection with the provided set
	Intersection(set IstioOperatorSet) IstioOperatorSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*install_istio_io_v1alpha1.IstioOperator, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another IstioOperatorSet
	Delta(newSet IstioOperatorSet) sksets.ResourceDelta
	// Create a deep copy of the current IstioOperatorSet
	Clone() IstioOperatorSet
}

func makeGenericIstioOperatorSet(istioOperatorList []*install_istio_io_v1alpha1.IstioOperator) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range istioOperatorList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type istioOperatorSet struct {
	set sksets.ResourceSet
}

func NewIstioOperatorSet(istioOperatorList ...*install_istio_io_v1alpha1.IstioOperator) IstioOperatorSet {
	return &istioOperatorSet{set: makeGenericIstioOperatorSet(istioOperatorList)}
}

func NewIstioOperatorSetFromList(istioOperatorList *install_istio_io_v1alpha1.IstioOperatorList) IstioOperatorSet {
	list := make([]*install_istio_io_v1alpha1.IstioOperator, 0, len(istioOperatorList.Items))
	for idx := range istioOperatorList.Items {
		list = append(list, &istioOperatorList.Items[idx])
	}
	return &istioOperatorSet{set: makeGenericIstioOperatorSet(list)}
}

func (s *istioOperatorSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *istioOperatorSet) List(filterResource ...func(*install_istio_io_v1alpha1.IstioOperator) bool) []*install_istio_io_v1alpha1.IstioOperator {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*install_istio_io_v1alpha1.IstioOperator))
		})
	}

	objs := s.Generic().List(genericFilters...)
	istioOperatorList := make([]*install_istio_io_v1alpha1.IstioOperator, 0, len(objs))
	for _, obj := range objs {
		istioOperatorList = append(istioOperatorList, obj.(*install_istio_io_v1alpha1.IstioOperator))
	}
	return istioOperatorList
}

func (s *istioOperatorSet) UnsortedList(filterResource ...func(*install_istio_io_v1alpha1.IstioOperator) bool) []*install_istio_io_v1alpha1.IstioOperator {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*install_istio_io_v1alpha1.IstioOperator))
		})
	}

	var istioOperatorList []*install_istio_io_v1alpha1.IstioOperator
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		istioOperatorList = append(istioOperatorList, obj.(*install_istio_io_v1alpha1.IstioOperator))
	}
	return istioOperatorList
}

func (s *istioOperatorSet) Map() map[string]*install_istio_io_v1alpha1.IstioOperator {
	if s == nil {
		return nil
	}

	newMap := map[string]*install_istio_io_v1alpha1.IstioOperator{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*install_istio_io_v1alpha1.IstioOperator)
	}
	return newMap
}

func (s *istioOperatorSet) Insert(
	istioOperatorList ...*install_istio_io_v1alpha1.IstioOperator,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range istioOperatorList {
		s.Generic().Insert(obj)
	}
}

func (s *istioOperatorSet) Has(istioOperator ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(istioOperator)
}

func (s *istioOperatorSet) Equal(
	istioOperatorSet IstioOperatorSet,
) bool {
	if s == nil {
		return istioOperatorSet == nil
	}
	return s.Generic().Equal(istioOperatorSet.Generic())
}

func (s *istioOperatorSet) Delete(IstioOperator ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(IstioOperator)
}

func (s *istioOperatorSet) Union(set IstioOperatorSet) IstioOperatorSet {
	if s == nil {
		return set
	}
	return NewIstioOperatorSet(append(s.List(), set.List()...)...)
}

func (s *istioOperatorSet) Difference(set IstioOperatorSet) IstioOperatorSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &istioOperatorSet{set: newSet}
}

func (s *istioOperatorSet) Intersection(set IstioOperatorSet) IstioOperatorSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var istioOperatorList []*install_istio_io_v1alpha1.IstioOperator
	for _, obj := range newSet.List() {
		istioOperatorList = append(istioOperatorList, obj.(*install_istio_io_v1alpha1.IstioOperator))
	}
	return NewIstioOperatorSet(istioOperatorList...)
}

func (s *istioOperatorSet) Find(id ezkube.ResourceId) (*install_istio_io_v1alpha1.IstioOperator, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find IstioOperator %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&install_istio_io_v1alpha1.IstioOperator{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*install_istio_io_v1alpha1.IstioOperator), nil
}

func (s *istioOperatorSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *istioOperatorSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *istioOperatorSet) Delta(newSet IstioOperatorSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *istioOperatorSet) Clone() IstioOperatorSet {
	if s == nil {
		return nil
	}
	return &istioOperatorMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type istioOperatorMergedSet struct {
	sets []sksets.ResourceSet
}

func NewIstioOperatorMergedSet(istioOperatorList ...*install_istio_io_v1alpha1.IstioOperator) IstioOperatorSet {
	return &istioOperatorMergedSet{sets: []sksets.ResourceSet{makeGenericIstioOperatorSet(istioOperatorList)}}
}

func NewIstioOperatorMergedSetFromList(istioOperatorList *install_istio_io_v1alpha1.IstioOperatorList) IstioOperatorSet {
	list := make([]*install_istio_io_v1alpha1.IstioOperator, 0, len(istioOperatorList.Items))
	for idx := range istioOperatorList.Items {
		list = append(list, &istioOperatorList.Items[idx])
	}
	return &istioOperatorMergedSet{sets: []sksets.ResourceSet{makeGenericIstioOperatorSet(list)}}
}

func (s *istioOperatorMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	toRet := sets.String{}
	for _, set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *istioOperatorMergedSet) List(filterResource ...func(*install_istio_io_v1alpha1.IstioOperator) bool) []*install_istio_io_v1alpha1.IstioOperator {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*install_istio_io_v1alpha1.IstioOperator))
		})
	}
	istioOperatorList := []*install_istio_io_v1alpha1.IstioOperator{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.List(genericFilters...) {
			if tracker[obj] {
				continue
			}
			tracker[obj] = true
			istioOperatorList = append(istioOperatorList, obj.(*install_istio_io_v1alpha1.IstioOperator))
		}
	}
	return istioOperatorList
}

func (s *istioOperatorMergedSet) UnsortedList(filterResource ...func(*install_istio_io_v1alpha1.IstioOperator) bool) []*install_istio_io_v1alpha1.IstioOperator {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*install_istio_io_v1alpha1.IstioOperator))
		})
	}
	istioOperatorList := []*install_istio_io_v1alpha1.IstioOperator{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.UnsortedList(genericFilters...) {
			if tracker[obj] {
				continue
			}
			tracker[obj] = true
			istioOperatorList = append(istioOperatorList, obj.(*install_istio_io_v1alpha1.IstioOperator))
		}
	}
	return istioOperatorList
}

func (s *istioOperatorMergedSet) Map() map[string]*install_istio_io_v1alpha1.IstioOperator {
	if s == nil {
		return nil
	}

	newMap := map[string]*install_istio_io_v1alpha1.IstioOperator{}
	for _, set := range s.sets {
		for k, v := range set.Map() {
			newMap[k] = v.(*install_istio_io_v1alpha1.IstioOperator)
		}
	}
	return newMap
}

func (s *istioOperatorMergedSet) Insert(
	istioOperatorList ...*install_istio_io_v1alpha1.IstioOperator,
) {
	if s == nil {
	}
	if len(s.sets) == 0 {
		s.sets = append(s.sets, makeGenericIstioOperatorSet(istioOperatorList))
	}
	for _, obj := range istioOperatorList {
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

func (s *istioOperatorMergedSet) Has(istioOperator ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	for _, set := range s.sets {
		if set.Has(istioOperator) {
			return true
		}
	}
	return false
}

func (s *istioOperatorMergedSet) Equal(
	istioOperatorSet IstioOperatorSet,
) bool {
	panic("unimplemented")
}

func (s *istioOperatorMergedSet) Delete(IstioOperator ezkube.ResourceId) {
	for _, set := range s.sets {
		set.Delete(IstioOperator)
	}
}

func (s *istioOperatorMergedSet) Union(set IstioOperatorSet) IstioOperatorSet {
	if s == nil {
		return set
	}
	return &istioOperatorMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *istioOperatorMergedSet) Difference(set IstioOperatorSet) IstioOperatorSet {
	panic("unimplemented")
}

func (s *istioOperatorMergedSet) Intersection(set IstioOperatorSet) IstioOperatorSet {
	panic("unimplemented")
}

func (s *istioOperatorMergedSet) Find(id ezkube.ResourceId) (*install_istio_io_v1alpha1.IstioOperator, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find IstioOperator %v", sksets.Key(id))
	}

	var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&install_istio_io_v1alpha1.IstioOperator{}, id)
		if err == nil {
			return obj.(*install_istio_io_v1alpha1.IstioOperator), nil
		}
	}

	return nil, err
}

func (s *istioOperatorMergedSet) Length() int {
	if s == nil {
		return 0
	}
	totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
	return totalLen
}

func (s *istioOperatorMergedSet) Generic() sksets.ResourceSet {
	res := make([]ezkube.ResourceId, s.Length())
	for _, thing := range s.List() {
		res = append(res, thing)
	}
	return sksets.NewResourceSet(res...)
}

func (s *istioOperatorMergedSet) Delta(newSet IstioOperatorSet) sksets.ResourceDelta {
	panic("unimplemented")
}

func (s *istioOperatorMergedSet) Clone() IstioOperatorSet {
	if s == nil {
		return nil
	}
	return &istioOperatorMergedSet{sets: s.sets[:]}
}
