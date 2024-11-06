// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	telemetry_istio_io_v1 "istio.io/client-go/pkg/apis/telemetry/v1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type TelemetrySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*telemetry_istio_io_v1.Telemetry) bool) []*telemetry_istio_io_v1.Telemetry
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*telemetry_istio_io_v1.Telemetry) bool) []*telemetry_istio_io_v1.Telemetry
	// Return the Set as a map of key to resource.
	Map() map[string]*telemetry_istio_io_v1.Telemetry
	// Insert a resource into the set.
	Insert(telemetry ...*telemetry_istio_io_v1.Telemetry)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(telemetrySet TelemetrySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(telemetry ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(telemetry ezkube.ResourceId)
	// Return the union with the provided set
	Union(set TelemetrySet) TelemetrySet
	// Return the difference with the provided set
	Difference(set TelemetrySet) TelemetrySet
	// Return the intersection with the provided set
	Intersection(set TelemetrySet) TelemetrySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*telemetry_istio_io_v1.Telemetry, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another TelemetrySet
	Delta(newSet TelemetrySet) sksets.ResourceDelta
	// Create a deep copy of the current TelemetrySet
	Clone() TelemetrySet
}

func makeGenericTelemetrySet(telemetryList []*telemetry_istio_io_v1.Telemetry) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range telemetryList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type telemetrySet struct {
	set sksets.ResourceSet
}

func NewTelemetrySet(telemetryList ...*telemetry_istio_io_v1.Telemetry) TelemetrySet {
	return &telemetrySet{set: makeGenericTelemetrySet(telemetryList)}
}

func NewTelemetrySetFromList(telemetryList *telemetry_istio_io_v1.TelemetryList) TelemetrySet {
	list := make([]*telemetry_istio_io_v1.Telemetry, 0, len(telemetryList.Items))
	for idx := range telemetryList.Items {
		list = append(list, telemetryList.Items[idx])
	}
	return &telemetrySet{set: makeGenericTelemetrySet(list)}
}

func (s *telemetrySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *telemetrySet) List(filterResource ...func(*telemetry_istio_io_v1.Telemetry) bool) []*telemetry_istio_io_v1.Telemetry {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*telemetry_istio_io_v1.Telemetry))
		})
	}

	objs := s.Generic().List(genericFilters...)
	telemetryList := make([]*telemetry_istio_io_v1.Telemetry, 0, len(objs))
	for _, obj := range objs {
		telemetryList = append(telemetryList, obj.(*telemetry_istio_io_v1.Telemetry))
	}
	return telemetryList
}

func (s *telemetrySet) UnsortedList(filterResource ...func(*telemetry_istio_io_v1.Telemetry) bool) []*telemetry_istio_io_v1.Telemetry {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*telemetry_istio_io_v1.Telemetry))
		})
	}

	var telemetryList []*telemetry_istio_io_v1.Telemetry
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		telemetryList = append(telemetryList, obj.(*telemetry_istio_io_v1.Telemetry))
	}
	return telemetryList
}

func (s *telemetrySet) Map() map[string]*telemetry_istio_io_v1.Telemetry {
	if s == nil {
		return nil
	}

	newMap := map[string]*telemetry_istio_io_v1.Telemetry{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*telemetry_istio_io_v1.Telemetry)
	}
	return newMap
}

func (s *telemetrySet) Insert(
	telemetryList ...*telemetry_istio_io_v1.Telemetry,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range telemetryList {
		s.Generic().Insert(obj)
	}
}

func (s *telemetrySet) Has(telemetry ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(telemetry)
}

func (s *telemetrySet) Equal(
	telemetrySet TelemetrySet,
) bool {
	if s == nil {
		return telemetrySet == nil
	}
	return s.Generic().Equal(telemetrySet.Generic())
}

func (s *telemetrySet) Delete(Telemetry ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Telemetry)
}

func (s *telemetrySet) Union(set TelemetrySet) TelemetrySet {
	if s == nil {
		return set
	}
	return NewTelemetrySet(append(s.List(), set.List()...)...)
}

func (s *telemetrySet) Difference(set TelemetrySet) TelemetrySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &telemetrySet{set: newSet}
}

func (s *telemetrySet) Intersection(set TelemetrySet) TelemetrySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var telemetryList []*telemetry_istio_io_v1.Telemetry
	for _, obj := range newSet.List() {
		telemetryList = append(telemetryList, obj.(*telemetry_istio_io_v1.Telemetry))
	}
	return NewTelemetrySet(telemetryList...)
}

func (s *telemetrySet) Find(id ezkube.ResourceId) (*telemetry_istio_io_v1.Telemetry, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Telemetry %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&telemetry_istio_io_v1.Telemetry{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*telemetry_istio_io_v1.Telemetry), nil
}

func (s *telemetrySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *telemetrySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *telemetrySet) Delta(newSet TelemetrySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *telemetrySet) Clone() TelemetrySet {
	if s == nil {
		return nil
	}
	return &telemetrySet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
