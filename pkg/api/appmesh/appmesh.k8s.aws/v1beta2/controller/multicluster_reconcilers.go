// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Mesh Resource across clusters.
// implemented by the user
type MulticlusterMeshReconciler interface {
	ReconcileMesh(clusterName string, obj *appmesh_k8s_aws_v1beta2.Mesh) (reconcile.Result, error)
}

// Reconcile deletion events for the Mesh Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterMeshDeletionReconciler interface {
	ReconcileMeshDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterMeshReconcilerFuncs struct {
	OnReconcileMesh         func(clusterName string, obj *appmesh_k8s_aws_v1beta2.Mesh) (reconcile.Result, error)
	OnReconcileMeshDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterMeshReconcilerFuncs) ReconcileMesh(clusterName string, obj *appmesh_k8s_aws_v1beta2.Mesh) (reconcile.Result, error) {
	if f.OnReconcileMesh == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMesh(clusterName, obj)
}

func (f *MulticlusterMeshReconcilerFuncs) ReconcileMeshDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileMeshDeletion == nil {
		return nil
	}
	return f.OnReconcileMeshDeletion(clusterName, req)
}

type MulticlusterMeshReconcileLoop interface {
	// AddMulticlusterMeshReconciler adds a MulticlusterMeshReconciler to the MulticlusterMeshReconcileLoop.
	AddMulticlusterMeshReconciler(ctx context.Context, rec MulticlusterMeshReconciler, predicates ...predicate.Predicate)
}

type multiclusterMeshReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterMeshReconcileLoop) AddMulticlusterMeshReconciler(ctx context.Context, rec MulticlusterMeshReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericMeshMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterMeshReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterMeshReconcileLoop {
	return &multiclusterMeshReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &appmesh_k8s_aws_v1beta2.Mesh{}, options)}
}

type genericMeshMulticlusterReconciler struct {
	reconciler MulticlusterMeshReconciler
}

func (g genericMeshMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterMeshDeletionReconciler); ok {
		return deletionReconciler.ReconcileMeshDeletion(cluster, req)
	}
	return nil
}

func (g genericMeshMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.Mesh)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return g.reconciler.ReconcileMesh(cluster, obj)
}

// Reconcile Upsert events for the VirtualService Resource across clusters.
// implemented by the user
type MulticlusterVirtualServiceReconciler interface {
	ReconcileVirtualService(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualService) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualService Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualServiceDeletionReconciler interface {
	ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualServiceReconcilerFuncs struct {
	OnReconcileVirtualService         func(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualService) (reconcile.Result, error)
	OnReconcileVirtualServiceDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualServiceReconcilerFuncs) ReconcileVirtualService(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualService) (reconcile.Result, error) {
	if f.OnReconcileVirtualService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualService(clusterName, obj)
}

func (f *MulticlusterVirtualServiceReconcilerFuncs) ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileVirtualServiceDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualServiceDeletion(clusterName, req)
}

type MulticlusterVirtualServiceReconcileLoop interface {
	// AddMulticlusterVirtualServiceReconciler adds a MulticlusterVirtualServiceReconciler to the MulticlusterVirtualServiceReconcileLoop.
	AddMulticlusterVirtualServiceReconciler(ctx context.Context, rec MulticlusterVirtualServiceReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualServiceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualServiceReconcileLoop) AddMulticlusterVirtualServiceReconciler(ctx context.Context, rec MulticlusterVirtualServiceReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualServiceMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualServiceReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterVirtualServiceReconcileLoop {
	return &multiclusterVirtualServiceReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &appmesh_k8s_aws_v1beta2.VirtualService{}, options)}
}

type genericVirtualServiceMulticlusterReconciler struct {
	reconciler MulticlusterVirtualServiceReconciler
}

func (g genericVirtualServiceMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualServiceDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualServiceDeletion(cluster, req)
	}
	return nil
}

func (g genericVirtualServiceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualService(cluster, obj)
}

// Reconcile Upsert events for the VirtualNode Resource across clusters.
// implemented by the user
type MulticlusterVirtualNodeReconciler interface {
	ReconcileVirtualNode(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualNode) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualNode Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualNodeDeletionReconciler interface {
	ReconcileVirtualNodeDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualNodeReconcilerFuncs struct {
	OnReconcileVirtualNode         func(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualNode) (reconcile.Result, error)
	OnReconcileVirtualNodeDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualNodeReconcilerFuncs) ReconcileVirtualNode(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualNode) (reconcile.Result, error) {
	if f.OnReconcileVirtualNode == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualNode(clusterName, obj)
}

func (f *MulticlusterVirtualNodeReconcilerFuncs) ReconcileVirtualNodeDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileVirtualNodeDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualNodeDeletion(clusterName, req)
}

type MulticlusterVirtualNodeReconcileLoop interface {
	// AddMulticlusterVirtualNodeReconciler adds a MulticlusterVirtualNodeReconciler to the MulticlusterVirtualNodeReconcileLoop.
	AddMulticlusterVirtualNodeReconciler(ctx context.Context, rec MulticlusterVirtualNodeReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualNodeReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualNodeReconcileLoop) AddMulticlusterVirtualNodeReconciler(ctx context.Context, rec MulticlusterVirtualNodeReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualNodeMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualNodeReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterVirtualNodeReconcileLoop {
	return &multiclusterVirtualNodeReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &appmesh_k8s_aws_v1beta2.VirtualNode{}, options)}
}

type genericVirtualNodeMulticlusterReconciler struct {
	reconciler MulticlusterVirtualNodeReconciler
}

func (g genericVirtualNodeMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualNodeDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualNodeDeletion(cluster, req)
	}
	return nil
}

func (g genericVirtualNodeMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualNode handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualNode(cluster, obj)
}

// Reconcile Upsert events for the VirtualRouter Resource across clusters.
// implemented by the user
type MulticlusterVirtualRouterReconciler interface {
	ReconcileVirtualRouter(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualRouter) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualRouter Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualRouterDeletionReconciler interface {
	ReconcileVirtualRouterDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualRouterReconcilerFuncs struct {
	OnReconcileVirtualRouter         func(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualRouter) (reconcile.Result, error)
	OnReconcileVirtualRouterDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualRouterReconcilerFuncs) ReconcileVirtualRouter(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualRouter) (reconcile.Result, error) {
	if f.OnReconcileVirtualRouter == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualRouter(clusterName, obj)
}

func (f *MulticlusterVirtualRouterReconcilerFuncs) ReconcileVirtualRouterDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileVirtualRouterDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualRouterDeletion(clusterName, req)
}

type MulticlusterVirtualRouterReconcileLoop interface {
	// AddMulticlusterVirtualRouterReconciler adds a MulticlusterVirtualRouterReconciler to the MulticlusterVirtualRouterReconcileLoop.
	AddMulticlusterVirtualRouterReconciler(ctx context.Context, rec MulticlusterVirtualRouterReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualRouterReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualRouterReconcileLoop) AddMulticlusterVirtualRouterReconciler(ctx context.Context, rec MulticlusterVirtualRouterReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualRouterMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualRouterReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterVirtualRouterReconcileLoop {
	return &multiclusterVirtualRouterReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &appmesh_k8s_aws_v1beta2.VirtualRouter{}, options)}
}

type genericVirtualRouterMulticlusterReconciler struct {
	reconciler MulticlusterVirtualRouterReconciler
}

func (g genericVirtualRouterMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualRouterDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualRouterDeletion(cluster, req)
	}
	return nil
}

func (g genericVirtualRouterMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualRouter handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualRouter(cluster, obj)
}

// Reconcile Upsert events for the VirtualGateway Resource across clusters.
// implemented by the user
type MulticlusterVirtualGatewayReconciler interface {
	ReconcileVirtualGateway(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualGateway) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualGateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualGatewayDeletionReconciler interface {
	ReconcileVirtualGatewayDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualGatewayReconcilerFuncs struct {
	OnReconcileVirtualGateway         func(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualGateway) (reconcile.Result, error)
	OnReconcileVirtualGatewayDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualGatewayReconcilerFuncs) ReconcileVirtualGateway(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualGateway) (reconcile.Result, error) {
	if f.OnReconcileVirtualGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualGateway(clusterName, obj)
}

func (f *MulticlusterVirtualGatewayReconcilerFuncs) ReconcileVirtualGatewayDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileVirtualGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualGatewayDeletion(clusterName, req)
}

type MulticlusterVirtualGatewayReconcileLoop interface {
	// AddMulticlusterVirtualGatewayReconciler adds a MulticlusterVirtualGatewayReconciler to the MulticlusterVirtualGatewayReconcileLoop.
	AddMulticlusterVirtualGatewayReconciler(ctx context.Context, rec MulticlusterVirtualGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualGatewayReconcileLoop) AddMulticlusterVirtualGatewayReconciler(ctx context.Context, rec MulticlusterVirtualGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterVirtualGatewayReconcileLoop {
	return &multiclusterVirtualGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &appmesh_k8s_aws_v1beta2.VirtualGateway{}, options)}
}

type genericVirtualGatewayMulticlusterReconciler struct {
	reconciler MulticlusterVirtualGatewayReconciler
}

func (g genericVirtualGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualGatewayDeletion(cluster, req)
	}
	return nil
}

func (g genericVirtualGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualGateway(cluster, obj)
}

// Reconcile Upsert events for the GatewayRoute Resource across clusters.
// implemented by the user
type MulticlusterGatewayRouteReconciler interface {
	ReconcileGatewayRoute(clusterName string, obj *appmesh_k8s_aws_v1beta2.GatewayRoute) (reconcile.Result, error)
}

// Reconcile deletion events for the GatewayRoute Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGatewayRouteDeletionReconciler interface {
	ReconcileGatewayRouteDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGatewayRouteReconcilerFuncs struct {
	OnReconcileGatewayRoute         func(clusterName string, obj *appmesh_k8s_aws_v1beta2.GatewayRoute) (reconcile.Result, error)
	OnReconcileGatewayRouteDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGatewayRouteReconcilerFuncs) ReconcileGatewayRoute(clusterName string, obj *appmesh_k8s_aws_v1beta2.GatewayRoute) (reconcile.Result, error) {
	if f.OnReconcileGatewayRoute == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGatewayRoute(clusterName, obj)
}

func (f *MulticlusterGatewayRouteReconcilerFuncs) ReconcileGatewayRouteDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGatewayRouteDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayRouteDeletion(clusterName, req)
}

type MulticlusterGatewayRouteReconcileLoop interface {
	// AddMulticlusterGatewayRouteReconciler adds a MulticlusterGatewayRouteReconciler to the MulticlusterGatewayRouteReconcileLoop.
	AddMulticlusterGatewayRouteReconciler(ctx context.Context, rec MulticlusterGatewayRouteReconciler, predicates ...predicate.Predicate)
}

type multiclusterGatewayRouteReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGatewayRouteReconcileLoop) AddMulticlusterGatewayRouteReconciler(ctx context.Context, rec MulticlusterGatewayRouteReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGatewayRouteMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGatewayRouteReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGatewayRouteReconcileLoop {
	return &multiclusterGatewayRouteReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &appmesh_k8s_aws_v1beta2.GatewayRoute{}, options)}
}

type genericGatewayRouteMulticlusterReconciler struct {
	reconciler MulticlusterGatewayRouteReconciler
}

func (g genericGatewayRouteMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGatewayRouteDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayRouteDeletion(cluster, req)
	}
	return nil
}

func (g genericGatewayRouteMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GatewayRoute handler received event for %T", object)
	}
	return g.reconciler.ReconcileGatewayRoute(cluster, obj)
}
