// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1beta2sets

import (
	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type MeshSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*appmesh_k8s_aws_v1beta2.Mesh) bool) []*appmesh_k8s_aws_v1beta2.Mesh
	// Return the Set as a map of key to resource.
	Map() map[string]*appmesh_k8s_aws_v1beta2.Mesh
	// Insert a resource into the set.
	Insert(mesh ...*appmesh_k8s_aws_v1beta2.Mesh)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(meshSet MeshSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(mesh ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(mesh ezkube.ResourceId)
	// Return the union with the provided set
	Union(set MeshSet) MeshSet
	// Return the difference with the provided set
	Difference(set MeshSet) MeshSet
	// Return the intersection with the provided set
	Intersection(set MeshSet) MeshSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.Mesh, error)
	// Get the length of the set
	Length() int
}

func makeGenericMeshSet(meshList []*appmesh_k8s_aws_v1beta2.Mesh) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range meshList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type meshSet struct {
	set sksets.ResourceSet
}

func NewMeshSet(meshList ...*appmesh_k8s_aws_v1beta2.Mesh) MeshSet {
	return &meshSet{set: makeGenericMeshSet(meshList)}
}

func NewMeshSetFromList(meshList *appmesh_k8s_aws_v1beta2.MeshList) MeshSet {
	list := make([]*appmesh_k8s_aws_v1beta2.Mesh, 0, len(meshList.Items))
	for idx := range meshList.Items {
		list = append(list, &meshList.Items[idx])
	}
	return &meshSet{set: makeGenericMeshSet(list)}
}

func (s *meshSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *meshSet) List(filterResource ...func(*appmesh_k8s_aws_v1beta2.Mesh) bool) []*appmesh_k8s_aws_v1beta2.Mesh {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*appmesh_k8s_aws_v1beta2.Mesh))
		})
	}

	var meshList []*appmesh_k8s_aws_v1beta2.Mesh
	for _, obj := range s.set.List(genericFilters...) {
		meshList = append(meshList, obj.(*appmesh_k8s_aws_v1beta2.Mesh))
	}
	return meshList
}

func (s *meshSet) Map() map[string]*appmesh_k8s_aws_v1beta2.Mesh {
	if s == nil {
		return nil
	}

	newMap := map[string]*appmesh_k8s_aws_v1beta2.Mesh{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*appmesh_k8s_aws_v1beta2.Mesh)
	}
	return newMap
}

func (s *meshSet) Insert(
	meshList ...*appmesh_k8s_aws_v1beta2.Mesh,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range meshList {
		s.set.Insert(obj)
	}
}

func (s *meshSet) Has(mesh ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(mesh)
}

func (s *meshSet) Equal(
	meshSet MeshSet,
) bool {
	if s == nil {
		return meshSet == nil
	}
	return s.set.Equal(makeGenericMeshSet(meshSet.List()))
}

func (s *meshSet) Delete(Mesh ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(Mesh)
}

func (s *meshSet) Union(set MeshSet) MeshSet {
	if s == nil {
		return set
	}
	return NewMeshSet(append(s.List(), set.List()...)...)
}

func (s *meshSet) Difference(set MeshSet) MeshSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericMeshSet(set.List()))
	return &meshSet{set: newSet}
}

func (s *meshSet) Intersection(set MeshSet) MeshSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericMeshSet(set.List()))
	var meshList []*appmesh_k8s_aws_v1beta2.Mesh
	for _, obj := range newSet.List() {
		meshList = append(meshList, obj.(*appmesh_k8s_aws_v1beta2.Mesh))
	}
	return NewMeshSet(meshList...)
}

func (s *meshSet) Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.Mesh, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Mesh %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&appmesh_k8s_aws_v1beta2.Mesh{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*appmesh_k8s_aws_v1beta2.Mesh), nil
}

func (s *meshSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}

type VirtualServiceSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*appmesh_k8s_aws_v1beta2.VirtualService) bool) []*appmesh_k8s_aws_v1beta2.VirtualService
	// Return the Set as a map of key to resource.
	Map() map[string]*appmesh_k8s_aws_v1beta2.VirtualService
	// Insert a resource into the set.
	Insert(virtualService ...*appmesh_k8s_aws_v1beta2.VirtualService)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualServiceSet VirtualServiceSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualService ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualService ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualServiceSet) VirtualServiceSet
	// Return the difference with the provided set
	Difference(set VirtualServiceSet) VirtualServiceSet
	// Return the intersection with the provided set
	Intersection(set VirtualServiceSet) VirtualServiceSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.VirtualService, error)
	// Get the length of the set
	Length() int
}

func makeGenericVirtualServiceSet(virtualServiceList []*appmesh_k8s_aws_v1beta2.VirtualService) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualServiceList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualServiceSet struct {
	set sksets.ResourceSet
}

func NewVirtualServiceSet(virtualServiceList ...*appmesh_k8s_aws_v1beta2.VirtualService) VirtualServiceSet {
	return &virtualServiceSet{set: makeGenericVirtualServiceSet(virtualServiceList)}
}

func NewVirtualServiceSetFromList(virtualServiceList *appmesh_k8s_aws_v1beta2.VirtualServiceList) VirtualServiceSet {
	list := make([]*appmesh_k8s_aws_v1beta2.VirtualService, 0, len(virtualServiceList.Items))
	for idx := range virtualServiceList.Items {
		list = append(list, &virtualServiceList.Items[idx])
	}
	return &virtualServiceSet{set: makeGenericVirtualServiceSet(list)}
}

func (s *virtualServiceSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *virtualServiceSet) List(filterResource ...func(*appmesh_k8s_aws_v1beta2.VirtualService) bool) []*appmesh_k8s_aws_v1beta2.VirtualService {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*appmesh_k8s_aws_v1beta2.VirtualService))
		})
	}

	var virtualServiceList []*appmesh_k8s_aws_v1beta2.VirtualService
	for _, obj := range s.set.List(genericFilters...) {
		virtualServiceList = append(virtualServiceList, obj.(*appmesh_k8s_aws_v1beta2.VirtualService))
	}
	return virtualServiceList
}

func (s *virtualServiceSet) Map() map[string]*appmesh_k8s_aws_v1beta2.VirtualService {
	if s == nil {
		return nil
	}

	newMap := map[string]*appmesh_k8s_aws_v1beta2.VirtualService{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*appmesh_k8s_aws_v1beta2.VirtualService)
	}
	return newMap
}

func (s *virtualServiceSet) Insert(
	virtualServiceList ...*appmesh_k8s_aws_v1beta2.VirtualService,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualServiceList {
		s.set.Insert(obj)
	}
}

func (s *virtualServiceSet) Has(virtualService ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(virtualService)
}

func (s *virtualServiceSet) Equal(
	virtualServiceSet VirtualServiceSet,
) bool {
	if s == nil {
		return virtualServiceSet == nil
	}
	return s.set.Equal(makeGenericVirtualServiceSet(virtualServiceSet.List()))
}

func (s *virtualServiceSet) Delete(VirtualService ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(VirtualService)
}

func (s *virtualServiceSet) Union(set VirtualServiceSet) VirtualServiceSet {
	if s == nil {
		return set
	}
	return NewVirtualServiceSet(append(s.List(), set.List()...)...)
}

func (s *virtualServiceSet) Difference(set VirtualServiceSet) VirtualServiceSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericVirtualServiceSet(set.List()))
	return &virtualServiceSet{set: newSet}
}

func (s *virtualServiceSet) Intersection(set VirtualServiceSet) VirtualServiceSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericVirtualServiceSet(set.List()))
	var virtualServiceList []*appmesh_k8s_aws_v1beta2.VirtualService
	for _, obj := range newSet.List() {
		virtualServiceList = append(virtualServiceList, obj.(*appmesh_k8s_aws_v1beta2.VirtualService))
	}
	return NewVirtualServiceSet(virtualServiceList...)
}

func (s *virtualServiceSet) Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.VirtualService, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualService %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&appmesh_k8s_aws_v1beta2.VirtualService{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*appmesh_k8s_aws_v1beta2.VirtualService), nil
}

func (s *virtualServiceSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}

type VirtualNodeSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*appmesh_k8s_aws_v1beta2.VirtualNode) bool) []*appmesh_k8s_aws_v1beta2.VirtualNode
	// Return the Set as a map of key to resource.
	Map() map[string]*appmesh_k8s_aws_v1beta2.VirtualNode
	// Insert a resource into the set.
	Insert(virtualNode ...*appmesh_k8s_aws_v1beta2.VirtualNode)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualNodeSet VirtualNodeSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualNode ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualNode ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualNodeSet) VirtualNodeSet
	// Return the difference with the provided set
	Difference(set VirtualNodeSet) VirtualNodeSet
	// Return the intersection with the provided set
	Intersection(set VirtualNodeSet) VirtualNodeSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.VirtualNode, error)
	// Get the length of the set
	Length() int
}

func makeGenericVirtualNodeSet(virtualNodeList []*appmesh_k8s_aws_v1beta2.VirtualNode) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualNodeList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualNodeSet struct {
	set sksets.ResourceSet
}

func NewVirtualNodeSet(virtualNodeList ...*appmesh_k8s_aws_v1beta2.VirtualNode) VirtualNodeSet {
	return &virtualNodeSet{set: makeGenericVirtualNodeSet(virtualNodeList)}
}

func NewVirtualNodeSetFromList(virtualNodeList *appmesh_k8s_aws_v1beta2.VirtualNodeList) VirtualNodeSet {
	list := make([]*appmesh_k8s_aws_v1beta2.VirtualNode, 0, len(virtualNodeList.Items))
	for idx := range virtualNodeList.Items {
		list = append(list, &virtualNodeList.Items[idx])
	}
	return &virtualNodeSet{set: makeGenericVirtualNodeSet(list)}
}

func (s *virtualNodeSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *virtualNodeSet) List(filterResource ...func(*appmesh_k8s_aws_v1beta2.VirtualNode) bool) []*appmesh_k8s_aws_v1beta2.VirtualNode {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*appmesh_k8s_aws_v1beta2.VirtualNode))
		})
	}

	var virtualNodeList []*appmesh_k8s_aws_v1beta2.VirtualNode
	for _, obj := range s.set.List(genericFilters...) {
		virtualNodeList = append(virtualNodeList, obj.(*appmesh_k8s_aws_v1beta2.VirtualNode))
	}
	return virtualNodeList
}

func (s *virtualNodeSet) Map() map[string]*appmesh_k8s_aws_v1beta2.VirtualNode {
	if s == nil {
		return nil
	}

	newMap := map[string]*appmesh_k8s_aws_v1beta2.VirtualNode{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	}
	return newMap
}

func (s *virtualNodeSet) Insert(
	virtualNodeList ...*appmesh_k8s_aws_v1beta2.VirtualNode,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualNodeList {
		s.set.Insert(obj)
	}
}

func (s *virtualNodeSet) Has(virtualNode ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(virtualNode)
}

func (s *virtualNodeSet) Equal(
	virtualNodeSet VirtualNodeSet,
) bool {
	if s == nil {
		return virtualNodeSet == nil
	}
	return s.set.Equal(makeGenericVirtualNodeSet(virtualNodeSet.List()))
}

func (s *virtualNodeSet) Delete(VirtualNode ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(VirtualNode)
}

func (s *virtualNodeSet) Union(set VirtualNodeSet) VirtualNodeSet {
	if s == nil {
		return set
	}
	return NewVirtualNodeSet(append(s.List(), set.List()...)...)
}

func (s *virtualNodeSet) Difference(set VirtualNodeSet) VirtualNodeSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericVirtualNodeSet(set.List()))
	return &virtualNodeSet{set: newSet}
}

func (s *virtualNodeSet) Intersection(set VirtualNodeSet) VirtualNodeSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericVirtualNodeSet(set.List()))
	var virtualNodeList []*appmesh_k8s_aws_v1beta2.VirtualNode
	for _, obj := range newSet.List() {
		virtualNodeList = append(virtualNodeList, obj.(*appmesh_k8s_aws_v1beta2.VirtualNode))
	}
	return NewVirtualNodeSet(virtualNodeList...)
}

func (s *virtualNodeSet) Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.VirtualNode, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualNode %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&appmesh_k8s_aws_v1beta2.VirtualNode{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*appmesh_k8s_aws_v1beta2.VirtualNode), nil
}

func (s *virtualNodeSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}

type VirtualRouterSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*appmesh_k8s_aws_v1beta2.VirtualRouter) bool) []*appmesh_k8s_aws_v1beta2.VirtualRouter
	// Return the Set as a map of key to resource.
	Map() map[string]*appmesh_k8s_aws_v1beta2.VirtualRouter
	// Insert a resource into the set.
	Insert(virtualRouter ...*appmesh_k8s_aws_v1beta2.VirtualRouter)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualRouterSet VirtualRouterSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualRouter ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualRouter ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualRouterSet) VirtualRouterSet
	// Return the difference with the provided set
	Difference(set VirtualRouterSet) VirtualRouterSet
	// Return the intersection with the provided set
	Intersection(set VirtualRouterSet) VirtualRouterSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.VirtualRouter, error)
	// Get the length of the set
	Length() int
}

func makeGenericVirtualRouterSet(virtualRouterList []*appmesh_k8s_aws_v1beta2.VirtualRouter) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualRouterList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualRouterSet struct {
	set sksets.ResourceSet
}

func NewVirtualRouterSet(virtualRouterList ...*appmesh_k8s_aws_v1beta2.VirtualRouter) VirtualRouterSet {
	return &virtualRouterSet{set: makeGenericVirtualRouterSet(virtualRouterList)}
}

func NewVirtualRouterSetFromList(virtualRouterList *appmesh_k8s_aws_v1beta2.VirtualRouterList) VirtualRouterSet {
	list := make([]*appmesh_k8s_aws_v1beta2.VirtualRouter, 0, len(virtualRouterList.Items))
	for idx := range virtualRouterList.Items {
		list = append(list, &virtualRouterList.Items[idx])
	}
	return &virtualRouterSet{set: makeGenericVirtualRouterSet(list)}
}

func (s *virtualRouterSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *virtualRouterSet) List(filterResource ...func(*appmesh_k8s_aws_v1beta2.VirtualRouter) bool) []*appmesh_k8s_aws_v1beta2.VirtualRouter {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*appmesh_k8s_aws_v1beta2.VirtualRouter))
		})
	}

	var virtualRouterList []*appmesh_k8s_aws_v1beta2.VirtualRouter
	for _, obj := range s.set.List(genericFilters...) {
		virtualRouterList = append(virtualRouterList, obj.(*appmesh_k8s_aws_v1beta2.VirtualRouter))
	}
	return virtualRouterList
}

func (s *virtualRouterSet) Map() map[string]*appmesh_k8s_aws_v1beta2.VirtualRouter {
	if s == nil {
		return nil
	}

	newMap := map[string]*appmesh_k8s_aws_v1beta2.VirtualRouter{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	}
	return newMap
}

func (s *virtualRouterSet) Insert(
	virtualRouterList ...*appmesh_k8s_aws_v1beta2.VirtualRouter,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualRouterList {
		s.set.Insert(obj)
	}
}

func (s *virtualRouterSet) Has(virtualRouter ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(virtualRouter)
}

func (s *virtualRouterSet) Equal(
	virtualRouterSet VirtualRouterSet,
) bool {
	if s == nil {
		return virtualRouterSet == nil
	}
	return s.set.Equal(makeGenericVirtualRouterSet(virtualRouterSet.List()))
}

func (s *virtualRouterSet) Delete(VirtualRouter ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(VirtualRouter)
}

func (s *virtualRouterSet) Union(set VirtualRouterSet) VirtualRouterSet {
	if s == nil {
		return set
	}
	return NewVirtualRouterSet(append(s.List(), set.List()...)...)
}

func (s *virtualRouterSet) Difference(set VirtualRouterSet) VirtualRouterSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericVirtualRouterSet(set.List()))
	return &virtualRouterSet{set: newSet}
}

func (s *virtualRouterSet) Intersection(set VirtualRouterSet) VirtualRouterSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericVirtualRouterSet(set.List()))
	var virtualRouterList []*appmesh_k8s_aws_v1beta2.VirtualRouter
	for _, obj := range newSet.List() {
		virtualRouterList = append(virtualRouterList, obj.(*appmesh_k8s_aws_v1beta2.VirtualRouter))
	}
	return NewVirtualRouterSet(virtualRouterList...)
}

func (s *virtualRouterSet) Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.VirtualRouter, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualRouter %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&appmesh_k8s_aws_v1beta2.VirtualRouter{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*appmesh_k8s_aws_v1beta2.VirtualRouter), nil
}

func (s *virtualRouterSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}

type VirtualGatewaySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*appmesh_k8s_aws_v1beta2.VirtualGateway) bool) []*appmesh_k8s_aws_v1beta2.VirtualGateway
	// Return the Set as a map of key to resource.
	Map() map[string]*appmesh_k8s_aws_v1beta2.VirtualGateway
	// Insert a resource into the set.
	Insert(virtualGateway ...*appmesh_k8s_aws_v1beta2.VirtualGateway)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualGatewaySet VirtualGatewaySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualGateway ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualGateway ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualGatewaySet) VirtualGatewaySet
	// Return the difference with the provided set
	Difference(set VirtualGatewaySet) VirtualGatewaySet
	// Return the intersection with the provided set
	Intersection(set VirtualGatewaySet) VirtualGatewaySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.VirtualGateway, error)
	// Get the length of the set
	Length() int
}

func makeGenericVirtualGatewaySet(virtualGatewayList []*appmesh_k8s_aws_v1beta2.VirtualGateway) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualGatewayList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualGatewaySet struct {
	set sksets.ResourceSet
}

func NewVirtualGatewaySet(virtualGatewayList ...*appmesh_k8s_aws_v1beta2.VirtualGateway) VirtualGatewaySet {
	return &virtualGatewaySet{set: makeGenericVirtualGatewaySet(virtualGatewayList)}
}

func NewVirtualGatewaySetFromList(virtualGatewayList *appmesh_k8s_aws_v1beta2.VirtualGatewayList) VirtualGatewaySet {
	list := make([]*appmesh_k8s_aws_v1beta2.VirtualGateway, 0, len(virtualGatewayList.Items))
	for idx := range virtualGatewayList.Items {
		list = append(list, &virtualGatewayList.Items[idx])
	}
	return &virtualGatewaySet{set: makeGenericVirtualGatewaySet(list)}
}

func (s *virtualGatewaySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *virtualGatewaySet) List(filterResource ...func(*appmesh_k8s_aws_v1beta2.VirtualGateway) bool) []*appmesh_k8s_aws_v1beta2.VirtualGateway {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*appmesh_k8s_aws_v1beta2.VirtualGateway))
		})
	}

	var virtualGatewayList []*appmesh_k8s_aws_v1beta2.VirtualGateway
	for _, obj := range s.set.List(genericFilters...) {
		virtualGatewayList = append(virtualGatewayList, obj.(*appmesh_k8s_aws_v1beta2.VirtualGateway))
	}
	return virtualGatewayList
}

func (s *virtualGatewaySet) Map() map[string]*appmesh_k8s_aws_v1beta2.VirtualGateway {
	if s == nil {
		return nil
	}

	newMap := map[string]*appmesh_k8s_aws_v1beta2.VirtualGateway{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	}
	return newMap
}

func (s *virtualGatewaySet) Insert(
	virtualGatewayList ...*appmesh_k8s_aws_v1beta2.VirtualGateway,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualGatewayList {
		s.set.Insert(obj)
	}
}

func (s *virtualGatewaySet) Has(virtualGateway ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(virtualGateway)
}

func (s *virtualGatewaySet) Equal(
	virtualGatewaySet VirtualGatewaySet,
) bool {
	if s == nil {
		return virtualGatewaySet == nil
	}
	return s.set.Equal(makeGenericVirtualGatewaySet(virtualGatewaySet.List()))
}

func (s *virtualGatewaySet) Delete(VirtualGateway ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(VirtualGateway)
}

func (s *virtualGatewaySet) Union(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return set
	}
	return NewVirtualGatewaySet(append(s.List(), set.List()...)...)
}

func (s *virtualGatewaySet) Difference(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericVirtualGatewaySet(set.List()))
	return &virtualGatewaySet{set: newSet}
}

func (s *virtualGatewaySet) Intersection(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericVirtualGatewaySet(set.List()))
	var virtualGatewayList []*appmesh_k8s_aws_v1beta2.VirtualGateway
	for _, obj := range newSet.List() {
		virtualGatewayList = append(virtualGatewayList, obj.(*appmesh_k8s_aws_v1beta2.VirtualGateway))
	}
	return NewVirtualGatewaySet(virtualGatewayList...)
}

func (s *virtualGatewaySet) Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.VirtualGateway, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualGateway %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&appmesh_k8s_aws_v1beta2.VirtualGateway{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*appmesh_k8s_aws_v1beta2.VirtualGateway), nil
}

func (s *virtualGatewaySet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}

type GatewayRouteSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*appmesh_k8s_aws_v1beta2.GatewayRoute) bool) []*appmesh_k8s_aws_v1beta2.GatewayRoute
	// Return the Set as a map of key to resource.
	Map() map[string]*appmesh_k8s_aws_v1beta2.GatewayRoute
	// Insert a resource into the set.
	Insert(gatewayRoute ...*appmesh_k8s_aws_v1beta2.GatewayRoute)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(gatewayRouteSet GatewayRouteSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(gatewayRoute ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(gatewayRoute ezkube.ResourceId)
	// Return the union with the provided set
	Union(set GatewayRouteSet) GatewayRouteSet
	// Return the difference with the provided set
	Difference(set GatewayRouteSet) GatewayRouteSet
	// Return the intersection with the provided set
	Intersection(set GatewayRouteSet) GatewayRouteSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.GatewayRoute, error)
	// Get the length of the set
	Length() int
}

func makeGenericGatewayRouteSet(gatewayRouteList []*appmesh_k8s_aws_v1beta2.GatewayRoute) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range gatewayRouteList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type gatewayRouteSet struct {
	set sksets.ResourceSet
}

func NewGatewayRouteSet(gatewayRouteList ...*appmesh_k8s_aws_v1beta2.GatewayRoute) GatewayRouteSet {
	return &gatewayRouteSet{set: makeGenericGatewayRouteSet(gatewayRouteList)}
}

func NewGatewayRouteSetFromList(gatewayRouteList *appmesh_k8s_aws_v1beta2.GatewayRouteList) GatewayRouteSet {
	list := make([]*appmesh_k8s_aws_v1beta2.GatewayRoute, 0, len(gatewayRouteList.Items))
	for idx := range gatewayRouteList.Items {
		list = append(list, &gatewayRouteList.Items[idx])
	}
	return &gatewayRouteSet{set: makeGenericGatewayRouteSet(list)}
}

func (s *gatewayRouteSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *gatewayRouteSet) List(filterResource ...func(*appmesh_k8s_aws_v1beta2.GatewayRoute) bool) []*appmesh_k8s_aws_v1beta2.GatewayRoute {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*appmesh_k8s_aws_v1beta2.GatewayRoute))
		})
	}

	var gatewayRouteList []*appmesh_k8s_aws_v1beta2.GatewayRoute
	for _, obj := range s.set.List(genericFilters...) {
		gatewayRouteList = append(gatewayRouteList, obj.(*appmesh_k8s_aws_v1beta2.GatewayRoute))
	}
	return gatewayRouteList
}

func (s *gatewayRouteSet) Map() map[string]*appmesh_k8s_aws_v1beta2.GatewayRoute {
	if s == nil {
		return nil
	}

	newMap := map[string]*appmesh_k8s_aws_v1beta2.GatewayRoute{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	}
	return newMap
}

func (s *gatewayRouteSet) Insert(
	gatewayRouteList ...*appmesh_k8s_aws_v1beta2.GatewayRoute,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range gatewayRouteList {
		s.set.Insert(obj)
	}
}

func (s *gatewayRouteSet) Has(gatewayRoute ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(gatewayRoute)
}

func (s *gatewayRouteSet) Equal(
	gatewayRouteSet GatewayRouteSet,
) bool {
	if s == nil {
		return gatewayRouteSet == nil
	}
	return s.set.Equal(makeGenericGatewayRouteSet(gatewayRouteSet.List()))
}

func (s *gatewayRouteSet) Delete(GatewayRoute ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(GatewayRoute)
}

func (s *gatewayRouteSet) Union(set GatewayRouteSet) GatewayRouteSet {
	if s == nil {
		return set
	}
	return NewGatewayRouteSet(append(s.List(), set.List()...)...)
}

func (s *gatewayRouteSet) Difference(set GatewayRouteSet) GatewayRouteSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericGatewayRouteSet(set.List()))
	return &gatewayRouteSet{set: newSet}
}

func (s *gatewayRouteSet) Intersection(set GatewayRouteSet) GatewayRouteSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericGatewayRouteSet(set.List()))
	var gatewayRouteList []*appmesh_k8s_aws_v1beta2.GatewayRoute
	for _, obj := range newSet.List() {
		gatewayRouteList = append(gatewayRouteList, obj.(*appmesh_k8s_aws_v1beta2.GatewayRoute))
	}
	return NewGatewayRouteSet(gatewayRouteList...)
}

func (s *gatewayRouteSet) Find(id ezkube.ResourceId) (*appmesh_k8s_aws_v1beta2.GatewayRoute, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find GatewayRoute %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&appmesh_k8s_aws_v1beta2.GatewayRoute{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*appmesh_k8s_aws_v1beta2.GatewayRoute), nil
}

func (s *gatewayRouteSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}
