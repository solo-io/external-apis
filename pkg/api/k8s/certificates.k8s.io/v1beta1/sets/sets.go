// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1beta1sets

import (
	certificates_k8s_io_v1beta1 "k8s.io/api/certificates/v1beta1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CertificateSigningRequestSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*certificates_k8s_io_v1beta1.CertificateSigningRequest) bool) []*certificates_k8s_io_v1beta1.CertificateSigningRequest
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*certificates_k8s_io_v1beta1.CertificateSigningRequest) bool) []*certificates_k8s_io_v1beta1.CertificateSigningRequest
	// Return the Set as a map of key to resource.
	Map() map[string]*certificates_k8s_io_v1beta1.CertificateSigningRequest
	// Insert a resource into the set.
	Insert(certificateSigningRequest ...*certificates_k8s_io_v1beta1.CertificateSigningRequest)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(certificateSigningRequestSet CertificateSigningRequestSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(certificateSigningRequest ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(certificateSigningRequest ezkube.ResourceId)
	// Return the union with the provided set
	Union(set CertificateSigningRequestSet) CertificateSigningRequestSet
	// Return the difference with the provided set
	Difference(set CertificateSigningRequestSet) CertificateSigningRequestSet
	// Return the intersection with the provided set
	Intersection(set CertificateSigningRequestSet) CertificateSigningRequestSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*certificates_k8s_io_v1beta1.CertificateSigningRequest, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another CertificateSigningRequestSet
	Delta(newSet CertificateSigningRequestSet) sksets.ResourceDelta
	// Create a deep copy of the current CertificateSigningRequestSet
	Clone() CertificateSigningRequestSet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
}

func makeGenericCertificateSigningRequestSet(
	sortFunc func(toInsert, existing client.Object) bool,
	certificateSigningRequestList []*certificates_k8s_io_v1beta1.CertificateSigningRequest,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range certificateSigningRequestList {
		genericResources = append(genericResources, obj)
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return sksets.NewResourceSet(genericSortFunc, genericResources...)
}

type certificateSigningRequestSet struct {
	set      sksets.ResourceSet
	sortFunc func(toInsert, existing client.Object) bool
}

func NewCertificateSigningRequestSet(
	sortFunc func(toInsert, existing client.Object) bool,
	certificateSigningRequestList ...*certificates_k8s_io_v1beta1.CertificateSigningRequest,
) CertificateSigningRequestSet {
	return &certificateSigningRequestSet{
		set:      makeGenericCertificateSigningRequestSet(sortFunc, certificateSigningRequestList),
		sortFunc: sortFunc,
	}
}

func NewCertificateSigningRequestSetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	certificateSigningRequestList *certificates_k8s_io_v1beta1.CertificateSigningRequestList,
) CertificateSigningRequestSet {
	list := make([]*certificates_k8s_io_v1beta1.CertificateSigningRequest, 0, len(certificateSigningRequestList.Items))
	for idx := range certificateSigningRequestList.Items {
		list = append(list, &certificateSigningRequestList.Items[idx])
	}
	return &certificateSigningRequestSet{
		set:      makeGenericCertificateSigningRequestSet(sortFunc, list),
		sortFunc: sortFunc,
	}
}

func (s *certificateSigningRequestSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *certificateSigningRequestSet) List(filterResource ...func(*certificates_k8s_io_v1beta1.CertificateSigningRequest) bool) []*certificates_k8s_io_v1beta1.CertificateSigningRequest {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*certificates_k8s_io_v1beta1.CertificateSigningRequest))
		})
	}

	objs := s.Generic().List(genericFilters...)
	certificateSigningRequestList := make([]*certificates_k8s_io_v1beta1.CertificateSigningRequest, 0, len(objs))
	for _, obj := range objs {
		certificateSigningRequestList = append(certificateSigningRequestList, obj.(*certificates_k8s_io_v1beta1.CertificateSigningRequest))
	}
	return certificateSigningRequestList
}

func (s *certificateSigningRequestSet) UnsortedList(filterResource ...func(*certificates_k8s_io_v1beta1.CertificateSigningRequest) bool) []*certificates_k8s_io_v1beta1.CertificateSigningRequest {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*certificates_k8s_io_v1beta1.CertificateSigningRequest))
		})
	}

	var certificateSigningRequestList []*certificates_k8s_io_v1beta1.CertificateSigningRequest
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		certificateSigningRequestList = append(certificateSigningRequestList, obj.(*certificates_k8s_io_v1beta1.CertificateSigningRequest))
	}
	return certificateSigningRequestList
}

func (s *certificateSigningRequestSet) Map() map[string]*certificates_k8s_io_v1beta1.CertificateSigningRequest {
	if s == nil {
		return nil
	}

	newMap := map[string]*certificates_k8s_io_v1beta1.CertificateSigningRequest{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*certificates_k8s_io_v1beta1.CertificateSigningRequest)
	}
	return newMap
}

func (s *certificateSigningRequestSet) Insert(
	certificateSigningRequestList ...*certificates_k8s_io_v1beta1.CertificateSigningRequest,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range certificateSigningRequestList {
		s.Generic().Insert(obj)
	}
}

func (s *certificateSigningRequestSet) Has(certificateSigningRequest ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(certificateSigningRequest)
}

func (s *certificateSigningRequestSet) Equal(
	certificateSigningRequestSet CertificateSigningRequestSet,
) bool {
	if s == nil {
		return certificateSigningRequestSet == nil
	}
	return s.Generic().Equal(certificateSigningRequestSet.Generic())
}

func (s *certificateSigningRequestSet) Delete(CertificateSigningRequest ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(CertificateSigningRequest)
}

func (s *certificateSigningRequestSet) Union(set CertificateSigningRequestSet) CertificateSigningRequestSet {
	if s == nil {
		return set
	}
	return NewCertificateSigningRequestSet(s.GetSortFunc(), append(s.List(), set.List()...)...)
}

func (s *certificateSigningRequestSet) Difference(set CertificateSigningRequestSet) CertificateSigningRequestSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &certificateSigningRequestSet{set: newSet}
}

func (s *certificateSigningRequestSet) Intersection(set CertificateSigningRequestSet) CertificateSigningRequestSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var certificateSigningRequestList []*certificates_k8s_io_v1beta1.CertificateSigningRequest
	for _, obj := range newSet.List() {
		certificateSigningRequestList = append(certificateSigningRequestList, obj.(*certificates_k8s_io_v1beta1.CertificateSigningRequest))
	}
	return NewCertificateSigningRequestSet(s.GetSortFunc(), certificateSigningRequestList...)
}

func (s *certificateSigningRequestSet) Find(id ezkube.ResourceId) (*certificates_k8s_io_v1beta1.CertificateSigningRequest, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find CertificateSigningRequest %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&certificates_k8s_io_v1beta1.CertificateSigningRequest{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*certificates_k8s_io_v1beta1.CertificateSigningRequest), nil
}

func (s *certificateSigningRequestSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *certificateSigningRequestSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *certificateSigningRequestSet) Delta(newSet CertificateSigningRequestSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *certificateSigningRequestSet) Clone() CertificateSigningRequestSet {
	if s == nil {
		return nil
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return &certificateSigningRequestSet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *certificateSigningRequestSet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}
