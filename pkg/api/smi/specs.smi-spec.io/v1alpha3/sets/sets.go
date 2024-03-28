// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha3sets

import (
	specs_smi_spec_io_v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type HTTPRouteGroupSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	// Return the Set as a map of key to resource.
	Map() map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	// Insert a resource into the set.
	Insert(hTTPRouteGroup ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(hTTPRouteGroupSet HTTPRouteGroupSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(hTTPRouteGroup ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(hTTPRouteGroup ezkube.ResourceId)
	// Return the union with the provided set
	Union(set HTTPRouteGroupSet) HTTPRouteGroupSet
	// Return the difference with the provided set
	Difference(set HTTPRouteGroupSet) HTTPRouteGroupSet
	// Return the intersection with the provided set
	Intersection(set HTTPRouteGroupSet) HTTPRouteGroupSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another HTTPRouteGroupSet
	Delta(newSet HTTPRouteGroupSet) sksets.ResourceDelta
	// Create a deep copy of the current HTTPRouteGroupSet
	Clone() HTTPRouteGroupSet
}

func makeGenericHTTPRouteGroupSet(hTTPRouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range hTTPRouteGroupList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type hTTPRouteGroupSet struct {
	set sksets.ResourceSet
}

func NewHTTPRouteGroupSet(hTTPRouteGroupList ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) HTTPRouteGroupSet {
	return &hTTPRouteGroupSet{set: makeGenericHTTPRouteGroupSet(hTTPRouteGroupList)}
}

func NewHTTPRouteGroupSetFromList(hTTPRouteGroupList *specs_smi_spec_io_v1alpha3.HTTPRouteGroupList) HTTPRouteGroupSet {
	list := make([]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, 0, len(hTTPRouteGroupList.Items))
	for idx := range hTTPRouteGroupList.Items {
		list = append(list, &hTTPRouteGroupList.Items[idx])
	}
	return &hTTPRouteGroupSet{set: makeGenericHTTPRouteGroupSet(list)}
}

func (s *hTTPRouteGroupSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *hTTPRouteGroupSet) List(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
		})
	}

	objs := s.Generic().List(genericFilters...)
	hTTPRouteGroupList := make([]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, 0, len(objs))
	for _, obj := range objs {
		hTTPRouteGroupList = append(hTTPRouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
	}
	return hTTPRouteGroupList
}

func (s *hTTPRouteGroupSet) UnsortedList(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
		})
	}

	var hTTPRouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		hTTPRouteGroupList = append(hTTPRouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
	}
	return hTTPRouteGroupList
}

func (s *hTTPRouteGroupSet) Map() map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}

	newMap := map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	}
	return newMap
}

func (s *hTTPRouteGroupSet) Insert(
	hTTPRouteGroupList ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range hTTPRouteGroupList {
		s.Generic().Insert(obj)
	}
}

func (s *hTTPRouteGroupSet) Has(hTTPRouteGroup ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(hTTPRouteGroup)
}

func (s *hTTPRouteGroupSet) Equal(
	hTTPRouteGroupSet HTTPRouteGroupSet,
) bool {
	if s == nil {
		return hTTPRouteGroupSet == nil
	}
	return s.Generic().Equal(hTTPRouteGroupSet.Generic())
}

func (s *hTTPRouteGroupSet) Delete(HTTPRouteGroup ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(HTTPRouteGroup)
}

func (s *hTTPRouteGroupSet) Union(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return set
	}
	return NewHTTPRouteGroupSet(append(s.List(), set.List()...)...)
}

func (s *hTTPRouteGroupSet) Difference(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &hTTPRouteGroupSet{set: newSet}
}

func (s *hTTPRouteGroupSet) Intersection(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var hTTPRouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	for _, obj := range newSet.List() {
		hTTPRouteGroupList = append(hTTPRouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
	}
	return NewHTTPRouteGroupSet(hTTPRouteGroupList...)
}

func (s *hTTPRouteGroupSet) Find(id ezkube.ResourceId) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find HTTPRouteGroup %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup), nil
}

func (s *hTTPRouteGroupSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *hTTPRouteGroupSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *hTTPRouteGroupSet) Delta(newSet HTTPRouteGroupSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *hTTPRouteGroupSet) Clone() HTTPRouteGroupSet {
	if s == nil {
		return nil
	}
	return &hTTPRouteGroupMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type hTTPRouteGroupMergedSet struct {
	sets []sksets.ResourceSet
}

func NewHTTPRouteGroupMergedSet(hTTPRouteGroupList ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) HTTPRouteGroupSet {
	return &hTTPRouteGroupMergedSet{sets: []sksets.ResourceSet{makeGenericHTTPRouteGroupSet(hTTPRouteGroupList)}}
}

func NewHTTPRouteGroupMergedSetFromList(hTTPRouteGroupList *specs_smi_spec_io_v1alpha3.HTTPRouteGroupList) HTTPRouteGroupSet {
	list := make([]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, 0, len(hTTPRouteGroupList.Items))
	for idx := range hTTPRouteGroupList.Items {
		list = append(list, &hTTPRouteGroupList.Items[idx])
	}
	return &hTTPRouteGroupMergedSet{sets: []sksets.ResourceSet{makeGenericHTTPRouteGroupSet(list)}}
}

func (s *hTTPRouteGroupMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	toRet := sets.String{}
	for _, set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *hTTPRouteGroupMergedSet) List(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
		})
	}
	hTTPRouteGroupList := []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.List(genericFilters...) {
			if tracker[obj] {
				continue
			}
			tracker[obj] = true
			hTTPRouteGroupList = append(hTTPRouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
		}
	}
	return hTTPRouteGroupList
}

func (s *hTTPRouteGroupMergedSet) UnsortedList(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
		})
	}
	hTTPRouteGroupList := []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	tracker := map[ezkube.ResourceId]bool{}
	for i := len(s.sets) - 1; i >= 0; i-- {
		set := s.sets[i]
		for _, obj := range set.UnsortedList(genericFilters...) {
			if tracker[obj] {
				continue
			}
			tracker[obj] = true
			hTTPRouteGroupList = append(hTTPRouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
		}
	}
	return hTTPRouteGroupList
}

func (s *hTTPRouteGroupMergedSet) Map() map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}

	newMap := map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	for _, set := range s.sets {
		for k, v := range set.Map() {
			newMap[k] = v.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
		}
	}
	return newMap
}

func (s *hTTPRouteGroupMergedSet) Insert(
	hTTPRouteGroupList ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup,
) {
	if s == nil {
	}
	if len(s.sets) == 0 {
		s.sets = append(s.sets, makeGenericHTTPRouteGroupSet(hTTPRouteGroupList))
	}
	for _, obj := range hTTPRouteGroupList {
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

func (s *hTTPRouteGroupMergedSet) Has(hTTPRouteGroup ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	for _, set := range s.sets {
		if set.Has(hTTPRouteGroup) {
			return true
		}
	}
	return false
}

func (s *hTTPRouteGroupMergedSet) Equal(
	hTTPRouteGroupSet HTTPRouteGroupSet,
) bool {
	panic("unimplemented")
}

func (s *hTTPRouteGroupMergedSet) Delete(HTTPRouteGroup ezkube.ResourceId) {
	for _, set := range s.sets {
		set.Delete(HTTPRouteGroup)
	}
}

func (s *hTTPRouteGroupMergedSet) Union(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return set
	}
	return &hTTPRouteGroupMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *hTTPRouteGroupMergedSet) Difference(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	panic("unimplemented")
}

func (s *hTTPRouteGroupMergedSet) Intersection(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	panic("unimplemented")
}

func (s *hTTPRouteGroupMergedSet) Find(id ezkube.ResourceId) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find HTTPRouteGroup %v", sksets.Key(id))
	}

	var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}, id)
		if err == nil {
			return obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup), nil
		}
	}

	return nil, err
}

func (s *hTTPRouteGroupMergedSet) Length() int {
	if s == nil {
		return 0
	}
	totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
	return totalLen
}

func (s *hTTPRouteGroupMergedSet) Generic() sksets.ResourceSet {
	res := make([]ezkube.ResourceId, s.Length())
	for _, thing := range s.List() {
		res = append(res, thing)
	}
	return sksets.NewResourceSet(res...)
}

func (s *hTTPRouteGroupMergedSet) Delta(newSet HTTPRouteGroupSet) sksets.ResourceDelta {
	panic("unimplemented")
}

func (s *hTTPRouteGroupMergedSet) Clone() HTTPRouteGroupSet {
	if s == nil {
		return nil
	}
	return &hTTPRouteGroupMergedSet{sets: s.sets[:]}
}
