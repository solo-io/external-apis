// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	apps_v1 "k8s.io/api/apps/v1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeploymentSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*apps_v1.Deployment) bool) []*apps_v1.Deployment
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*apps_v1.Deployment) bool) []*apps_v1.Deployment
	// Return the Set as a map of key to resource.
	Map() map[string]*apps_v1.Deployment
	// Insert a resource into the set.
	Insert(deployment ...*apps_v1.Deployment)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(deploymentSet DeploymentSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(deployment ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(deployment ezkube.ResourceId)
	// Return the union with the provided set
	Union(set DeploymentSet) DeploymentSet
	// Return the difference with the provided set
	Difference(set DeploymentSet) DeploymentSet
	// Return the intersection with the provided set
	Intersection(set DeploymentSet) DeploymentSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*apps_v1.Deployment, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another DeploymentSet
	Delta(newSet DeploymentSet) sksets.ResourceDelta
	// Create a deep copy of the current DeploymentSet
	Clone() DeploymentSet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
}

func makeGenericDeploymentSet(
	sortFunc func(toInsert, existing client.Object) bool,
	deploymentList []*apps_v1.Deployment,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range deploymentList {
		genericResources = append(genericResources, obj)
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return sksets.NewResourceSet(genericSortFunc, genericResources...)
}

type deploymentSet struct {
	set      sksets.ResourceSet
	sortFunc func(toInsert, existing client.Object) bool
}

func NewDeploymentSet(
	sortFunc func(toInsert, existing client.Object) bool,
	deploymentList ...*apps_v1.Deployment,
) DeploymentSet {
	return &deploymentSet{
		set:      makeGenericDeploymentSet(sortFunc, deploymentList),
		sortFunc: sortFunc,
	}
}

func NewDeploymentSetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	deploymentList *apps_v1.DeploymentList,
) DeploymentSet {
	list := make([]*apps_v1.Deployment, 0, len(deploymentList.Items))
	for idx := range deploymentList.Items {
		list = append(list, &deploymentList.Items[idx])
	}
	return &deploymentSet{
		set:      makeGenericDeploymentSet(sortFunc, list),
		sortFunc: sortFunc,
	}
}

func (s *deploymentSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *deploymentSet) List(filterResource ...func(*apps_v1.Deployment) bool) []*apps_v1.Deployment {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apps_v1.Deployment))
		})
	}

	objs := s.Generic().List(genericFilters...)
	deploymentList := make([]*apps_v1.Deployment, 0, len(objs))
	for _, obj := range objs {
		deploymentList = append(deploymentList, obj.(*apps_v1.Deployment))
	}
	return deploymentList
}

func (s *deploymentSet) UnsortedList(filterResource ...func(*apps_v1.Deployment) bool) []*apps_v1.Deployment {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apps_v1.Deployment))
		})
	}

	var deploymentList []*apps_v1.Deployment
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		deploymentList = append(deploymentList, obj.(*apps_v1.Deployment))
	}
	return deploymentList
}

func (s *deploymentSet) Map() map[string]*apps_v1.Deployment {
	if s == nil {
		return nil
	}

	newMap := map[string]*apps_v1.Deployment{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*apps_v1.Deployment)
	}
	return newMap
}

func (s *deploymentSet) Insert(
	deploymentList ...*apps_v1.Deployment,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range deploymentList {
		s.Generic().Insert(obj)
	}
}

func (s *deploymentSet) Has(deployment ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(deployment)
}

func (s *deploymentSet) Equal(
	deploymentSet DeploymentSet,
) bool {
	if s == nil {
		return deploymentSet == nil
	}
	return s.Generic().Equal(deploymentSet.Generic())
}

func (s *deploymentSet) Delete(Deployment ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Deployment)
}

func (s *deploymentSet) Union(set DeploymentSet) DeploymentSet {
	if s == nil {
		return set
	}
	return NewDeploymentSet(s.GetSortFunc(), append(s.List(), set.List()...)...)
}

func (s *deploymentSet) Difference(set DeploymentSet) DeploymentSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &deploymentSet{set: newSet}
}

func (s *deploymentSet) Intersection(set DeploymentSet) DeploymentSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var deploymentList []*apps_v1.Deployment
	for _, obj := range newSet.List() {
		deploymentList = append(deploymentList, obj.(*apps_v1.Deployment))
	}
	return NewDeploymentSet(s.GetSortFunc(), deploymentList...)
}

func (s *deploymentSet) Find(id ezkube.ResourceId) (*apps_v1.Deployment, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Deployment %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&apps_v1.Deployment{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*apps_v1.Deployment), nil
}

func (s *deploymentSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *deploymentSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *deploymentSet) Delta(newSet DeploymentSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *deploymentSet) Clone() DeploymentSet {
	if s == nil {
		return nil
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return &deploymentSet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *deploymentSet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}

type ReplicaSetSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*apps_v1.ReplicaSet) bool) []*apps_v1.ReplicaSet
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*apps_v1.ReplicaSet) bool) []*apps_v1.ReplicaSet
	// Return the Set as a map of key to resource.
	Map() map[string]*apps_v1.ReplicaSet
	// Insert a resource into the set.
	Insert(replicaSet ...*apps_v1.ReplicaSet)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(replicaSetSet ReplicaSetSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(replicaSet ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(replicaSet ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ReplicaSetSet) ReplicaSetSet
	// Return the difference with the provided set
	Difference(set ReplicaSetSet) ReplicaSetSet
	// Return the intersection with the provided set
	Intersection(set ReplicaSetSet) ReplicaSetSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*apps_v1.ReplicaSet, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ReplicaSetSet
	Delta(newSet ReplicaSetSet) sksets.ResourceDelta
	// Create a deep copy of the current ReplicaSetSet
	Clone() ReplicaSetSet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
}

func makeGenericReplicaSetSet(
	sortFunc func(toInsert, existing client.Object) bool,
	replicaSetList []*apps_v1.ReplicaSet,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range replicaSetList {
		genericResources = append(genericResources, obj)
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return sksets.NewResourceSet(genericSortFunc, genericResources...)
}

type replicaSetSet struct {
	set      sksets.ResourceSet
	sortFunc func(toInsert, existing client.Object) bool
}

func NewReplicaSetSet(
	sortFunc func(toInsert, existing client.Object) bool,
	replicaSetList ...*apps_v1.ReplicaSet,
) ReplicaSetSet {
	return &replicaSetSet{
		set:      makeGenericReplicaSetSet(sortFunc, replicaSetList),
		sortFunc: sortFunc,
	}
}

func NewReplicaSetSetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	replicaSetList *apps_v1.ReplicaSetList,
) ReplicaSetSet {
	list := make([]*apps_v1.ReplicaSet, 0, len(replicaSetList.Items))
	for idx := range replicaSetList.Items {
		list = append(list, &replicaSetList.Items[idx])
	}
	return &replicaSetSet{
		set:      makeGenericReplicaSetSet(sortFunc, list),
		sortFunc: sortFunc,
	}
}

func (s *replicaSetSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *replicaSetSet) List(filterResource ...func(*apps_v1.ReplicaSet) bool) []*apps_v1.ReplicaSet {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apps_v1.ReplicaSet))
		})
	}

	objs := s.Generic().List(genericFilters...)
	replicaSetList := make([]*apps_v1.ReplicaSet, 0, len(objs))
	for _, obj := range objs {
		replicaSetList = append(replicaSetList, obj.(*apps_v1.ReplicaSet))
	}
	return replicaSetList
}

func (s *replicaSetSet) UnsortedList(filterResource ...func(*apps_v1.ReplicaSet) bool) []*apps_v1.ReplicaSet {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apps_v1.ReplicaSet))
		})
	}

	var replicaSetList []*apps_v1.ReplicaSet
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		replicaSetList = append(replicaSetList, obj.(*apps_v1.ReplicaSet))
	}
	return replicaSetList
}

func (s *replicaSetSet) Map() map[string]*apps_v1.ReplicaSet {
	if s == nil {
		return nil
	}

	newMap := map[string]*apps_v1.ReplicaSet{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*apps_v1.ReplicaSet)
	}
	return newMap
}

func (s *replicaSetSet) Insert(
	replicaSetList ...*apps_v1.ReplicaSet,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range replicaSetList {
		s.Generic().Insert(obj)
	}
}

func (s *replicaSetSet) Has(replicaSet ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(replicaSet)
}

func (s *replicaSetSet) Equal(
	replicaSetSet ReplicaSetSet,
) bool {
	if s == nil {
		return replicaSetSet == nil
	}
	return s.Generic().Equal(replicaSetSet.Generic())
}

func (s *replicaSetSet) Delete(ReplicaSet ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ReplicaSet)
}

func (s *replicaSetSet) Union(set ReplicaSetSet) ReplicaSetSet {
	if s == nil {
		return set
	}
	return NewReplicaSetSet(s.GetSortFunc(), append(s.List(), set.List()...)...)
}

func (s *replicaSetSet) Difference(set ReplicaSetSet) ReplicaSetSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &replicaSetSet{set: newSet}
}

func (s *replicaSetSet) Intersection(set ReplicaSetSet) ReplicaSetSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var replicaSetList []*apps_v1.ReplicaSet
	for _, obj := range newSet.List() {
		replicaSetList = append(replicaSetList, obj.(*apps_v1.ReplicaSet))
	}
	return NewReplicaSetSet(s.GetSortFunc(), replicaSetList...)
}

func (s *replicaSetSet) Find(id ezkube.ResourceId) (*apps_v1.ReplicaSet, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ReplicaSet %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&apps_v1.ReplicaSet{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*apps_v1.ReplicaSet), nil
}

func (s *replicaSetSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *replicaSetSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *replicaSetSet) Delta(newSet ReplicaSetSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *replicaSetSet) Clone() ReplicaSetSet {
	if s == nil {
		return nil
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return &replicaSetSet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *replicaSetSet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}

type DaemonSetSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*apps_v1.DaemonSet) bool) []*apps_v1.DaemonSet
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*apps_v1.DaemonSet) bool) []*apps_v1.DaemonSet
	// Return the Set as a map of key to resource.
	Map() map[string]*apps_v1.DaemonSet
	// Insert a resource into the set.
	Insert(daemonSet ...*apps_v1.DaemonSet)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(daemonSetSet DaemonSetSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(daemonSet ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(daemonSet ezkube.ResourceId)
	// Return the union with the provided set
	Union(set DaemonSetSet) DaemonSetSet
	// Return the difference with the provided set
	Difference(set DaemonSetSet) DaemonSetSet
	// Return the intersection with the provided set
	Intersection(set DaemonSetSet) DaemonSetSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*apps_v1.DaemonSet, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another DaemonSetSet
	Delta(newSet DaemonSetSet) sksets.ResourceDelta
	// Create a deep copy of the current DaemonSetSet
	Clone() DaemonSetSet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
}

func makeGenericDaemonSetSet(
	sortFunc func(toInsert, existing client.Object) bool,
	daemonSetList []*apps_v1.DaemonSet,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range daemonSetList {
		genericResources = append(genericResources, obj)
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return sksets.NewResourceSet(genericSortFunc, genericResources...)
}

type daemonSetSet struct {
	set      sksets.ResourceSet
	sortFunc func(toInsert, existing client.Object) bool
}

func NewDaemonSetSet(
	sortFunc func(toInsert, existing client.Object) bool,
	daemonSetList ...*apps_v1.DaemonSet,
) DaemonSetSet {
	return &daemonSetSet{
		set:      makeGenericDaemonSetSet(sortFunc, daemonSetList),
		sortFunc: sortFunc,
	}
}

func NewDaemonSetSetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	daemonSetList *apps_v1.DaemonSetList,
) DaemonSetSet {
	list := make([]*apps_v1.DaemonSet, 0, len(daemonSetList.Items))
	for idx := range daemonSetList.Items {
		list = append(list, &daemonSetList.Items[idx])
	}
	return &daemonSetSet{
		set:      makeGenericDaemonSetSet(sortFunc, list),
		sortFunc: sortFunc,
	}
}

func (s *daemonSetSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *daemonSetSet) List(filterResource ...func(*apps_v1.DaemonSet) bool) []*apps_v1.DaemonSet {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apps_v1.DaemonSet))
		})
	}

	objs := s.Generic().List(genericFilters...)
	daemonSetList := make([]*apps_v1.DaemonSet, 0, len(objs))
	for _, obj := range objs {
		daemonSetList = append(daemonSetList, obj.(*apps_v1.DaemonSet))
	}
	return daemonSetList
}

func (s *daemonSetSet) UnsortedList(filterResource ...func(*apps_v1.DaemonSet) bool) []*apps_v1.DaemonSet {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apps_v1.DaemonSet))
		})
	}

	var daemonSetList []*apps_v1.DaemonSet
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		daemonSetList = append(daemonSetList, obj.(*apps_v1.DaemonSet))
	}
	return daemonSetList
}

func (s *daemonSetSet) Map() map[string]*apps_v1.DaemonSet {
	if s == nil {
		return nil
	}

	newMap := map[string]*apps_v1.DaemonSet{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*apps_v1.DaemonSet)
	}
	return newMap
}

func (s *daemonSetSet) Insert(
	daemonSetList ...*apps_v1.DaemonSet,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range daemonSetList {
		s.Generic().Insert(obj)
	}
}

func (s *daemonSetSet) Has(daemonSet ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(daemonSet)
}

func (s *daemonSetSet) Equal(
	daemonSetSet DaemonSetSet,
) bool {
	if s == nil {
		return daemonSetSet == nil
	}
	return s.Generic().Equal(daemonSetSet.Generic())
}

func (s *daemonSetSet) Delete(DaemonSet ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(DaemonSet)
}

func (s *daemonSetSet) Union(set DaemonSetSet) DaemonSetSet {
	if s == nil {
		return set
	}
	return NewDaemonSetSet(s.GetSortFunc(), append(s.List(), set.List()...)...)
}

func (s *daemonSetSet) Difference(set DaemonSetSet) DaemonSetSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &daemonSetSet{set: newSet}
}

func (s *daemonSetSet) Intersection(set DaemonSetSet) DaemonSetSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var daemonSetList []*apps_v1.DaemonSet
	for _, obj := range newSet.List() {
		daemonSetList = append(daemonSetList, obj.(*apps_v1.DaemonSet))
	}
	return NewDaemonSetSet(s.GetSortFunc(), daemonSetList...)
}

func (s *daemonSetSet) Find(id ezkube.ResourceId) (*apps_v1.DaemonSet, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find DaemonSet %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&apps_v1.DaemonSet{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*apps_v1.DaemonSet), nil
}

func (s *daemonSetSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *daemonSetSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *daemonSetSet) Delta(newSet DaemonSetSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *daemonSetSet) Clone() DaemonSetSet {
	if s == nil {
		return nil
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return &daemonSetSet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *daemonSetSet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}

type StatefulSetSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*apps_v1.StatefulSet) bool) []*apps_v1.StatefulSet
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*apps_v1.StatefulSet) bool) []*apps_v1.StatefulSet
	// Return the Set as a map of key to resource.
	Map() map[string]*apps_v1.StatefulSet
	// Insert a resource into the set.
	Insert(statefulSet ...*apps_v1.StatefulSet)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(statefulSetSet StatefulSetSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(statefulSet ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(statefulSet ezkube.ResourceId)
	// Return the union with the provided set
	Union(set StatefulSetSet) StatefulSetSet
	// Return the difference with the provided set
	Difference(set StatefulSetSet) StatefulSetSet
	// Return the intersection with the provided set
	Intersection(set StatefulSetSet) StatefulSetSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*apps_v1.StatefulSet, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another StatefulSetSet
	Delta(newSet StatefulSetSet) sksets.ResourceDelta
	// Create a deep copy of the current StatefulSetSet
	Clone() StatefulSetSet
	// Get the sort function used by the set
	GetSortFunc() func(toInsert, existing client.Object) bool
}

func makeGenericStatefulSetSet(
	sortFunc func(toInsert, existing client.Object) bool,
	statefulSetList []*apps_v1.StatefulSet,
) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range statefulSetList {
		genericResources = append(genericResources, obj)
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return sksets.NewResourceSet(genericSortFunc, genericResources...)
}

type statefulSetSet struct {
	set      sksets.ResourceSet
	sortFunc func(toInsert, existing client.Object) bool
}

func NewStatefulSetSet(
	sortFunc func(toInsert, existing client.Object) bool,
	statefulSetList ...*apps_v1.StatefulSet,
) StatefulSetSet {
	return &statefulSetSet{
		set:      makeGenericStatefulSetSet(sortFunc, statefulSetList),
		sortFunc: sortFunc,
	}
}

func NewStatefulSetSetFromList(
	sortFunc func(toInsert, existing client.Object) bool,
	statefulSetList *apps_v1.StatefulSetList,
) StatefulSetSet {
	list := make([]*apps_v1.StatefulSet, 0, len(statefulSetList.Items))
	for idx := range statefulSetList.Items {
		list = append(list, &statefulSetList.Items[idx])
	}
	return &statefulSetSet{
		set:      makeGenericStatefulSetSet(sortFunc, list),
		sortFunc: sortFunc,
	}
}

func (s *statefulSetSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *statefulSetSet) List(filterResource ...func(*apps_v1.StatefulSet) bool) []*apps_v1.StatefulSet {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apps_v1.StatefulSet))
		})
	}

	objs := s.Generic().List(genericFilters...)
	statefulSetList := make([]*apps_v1.StatefulSet, 0, len(objs))
	for _, obj := range objs {
		statefulSetList = append(statefulSetList, obj.(*apps_v1.StatefulSet))
	}
	return statefulSetList
}

func (s *statefulSetSet) UnsortedList(filterResource ...func(*apps_v1.StatefulSet) bool) []*apps_v1.StatefulSet {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*apps_v1.StatefulSet))
		})
	}

	var statefulSetList []*apps_v1.StatefulSet
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		statefulSetList = append(statefulSetList, obj.(*apps_v1.StatefulSet))
	}
	return statefulSetList
}

func (s *statefulSetSet) Map() map[string]*apps_v1.StatefulSet {
	if s == nil {
		return nil
	}

	newMap := map[string]*apps_v1.StatefulSet{}
	for k, v := range s.Generic().Map().Map() {
		newMap[k] = v.(*apps_v1.StatefulSet)
	}
	return newMap
}

func (s *statefulSetSet) Insert(
	statefulSetList ...*apps_v1.StatefulSet,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range statefulSetList {
		s.Generic().Insert(obj)
	}
}

func (s *statefulSetSet) Has(statefulSet ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(statefulSet)
}

func (s *statefulSetSet) Equal(
	statefulSetSet StatefulSetSet,
) bool {
	if s == nil {
		return statefulSetSet == nil
	}
	return s.Generic().Equal(statefulSetSet.Generic())
}

func (s *statefulSetSet) Delete(StatefulSet ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(StatefulSet)
}

func (s *statefulSetSet) Union(set StatefulSetSet) StatefulSetSet {
	if s == nil {
		return set
	}
	return NewStatefulSetSet(s.GetSortFunc(), append(s.List(), set.List()...)...)
}

func (s *statefulSetSet) Difference(set StatefulSetSet) StatefulSetSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &statefulSetSet{set: newSet}
}

func (s *statefulSetSet) Intersection(set StatefulSetSet) StatefulSetSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var statefulSetList []*apps_v1.StatefulSet
	for _, obj := range newSet.List() {
		statefulSetList = append(statefulSetList, obj.(*apps_v1.StatefulSet))
	}
	return NewStatefulSetSet(s.GetSortFunc(), statefulSetList...)
}

func (s *statefulSetSet) Find(id ezkube.ResourceId) (*apps_v1.StatefulSet, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find StatefulSet %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&apps_v1.StatefulSet{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*apps_v1.StatefulSet), nil
}

func (s *statefulSetSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *statefulSetSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *statefulSetSet) Delta(newSet StatefulSetSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *statefulSetSet) Clone() StatefulSetSet {
	if s == nil {
		return nil
	}
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(client.Object), existing.(client.Object))
	}
	return &statefulSetSet{
		set: sksets.NewResourceSet(
			genericSortFunc,
			s.Generic().Clone().List()...,
		),
	}
}

func (s *statefulSetSet) GetSortFunc() func(toInsert, existing client.Object) bool {
	return s.sortFunc
}
