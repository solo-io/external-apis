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
	Insert(httprouteGroup ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(httprouteGroupSet HTTPRouteGroupSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(httprouteGroup ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(httprouteGroup ezkube.ResourceId)
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

func makeGenericHTTPRouteGroupSet(httprouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range httprouteGroupList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type httprouteGroupSet struct {
	set sksets.ResourceSet
}

func NewHTTPRouteGroupSet(httprouteGroupList ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) HTTPRouteGroupSet {
	return &httprouteGroupSet{set: makeGenericHTTPRouteGroupSet(httprouteGroupList)}
}

func NewHTTPRouteGroupSetFromList(httprouteGroupList *specs_smi_spec_io_v1alpha3.HTTPRouteGroupList) HTTPRouteGroupSet {
	list := make([]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, 0, len(httprouteGroupList.Items))
	for idx := range httprouteGroupList.Items {
		list = append(list, &httprouteGroupList.Items[idx])
	}
	return &httprouteGroupSet{set: makeGenericHTTPRouteGroupSet(list)}
}

func (s *httprouteGroupSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *httprouteGroupSet) List(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
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
	httprouteGroupList := make([]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, 0, len(objs))
	for _, obj := range objs {
		httprouteGroupList = append(httprouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
	}
	return httprouteGroupList
}

func (s *httprouteGroupSet) UnsortedList(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
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

	var httprouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		httprouteGroupList = append(httprouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
	}
	return httprouteGroupList
}

func (s *httprouteGroupSet) Map() map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}

	newMap := map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	}
	return newMap
}

func (s *httprouteGroupSet) Insert(
	httprouteGroupList ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range httprouteGroupList {
		s.Generic().Insert(obj)
	}
}

func (s *httprouteGroupSet) Has(httprouteGroup ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(httprouteGroup)
}

func (s *httprouteGroupSet) Equal(
	httprouteGroupSet HTTPRouteGroupSet,
) bool {
	if s == nil {
		return httprouteGroupSet == nil
	}
	return s.Generic().Equal(httprouteGroupSet.Generic())
}

func (s *httprouteGroupSet) Delete(HTTPRouteGroup ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(HTTPRouteGroup)
}

func (s *httprouteGroupSet) Union(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return set
	}
	return NewHTTPRouteGroupSet(append(s.List(), set.List()...)...)
}

func (s *httprouteGroupSet) Difference(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &httprouteGroupSet{set: newSet}
}

func (s *httprouteGroupSet) Intersection(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var httprouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	for _, obj := range newSet.List() {
		httprouteGroupList = append(httprouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
	}
	return NewHTTPRouteGroupSet(httprouteGroupList...)
}

func (s *httprouteGroupSet) Find(id ezkube.ResourceId) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find HTTPRouteGroup %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup), nil
}

func (s *httprouteGroupSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *httprouteGroupSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *httprouteGroupSet) Delta(newSet HTTPRouteGroupSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *httprouteGroupSet) Clone() HTTPRouteGroupSet {
	if s == nil {
		return nil
	}
	return &httprouteGroupSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
