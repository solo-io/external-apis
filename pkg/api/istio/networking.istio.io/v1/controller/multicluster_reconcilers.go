// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	networking_istio_io_v1 "istio.io/client-go/pkg/apis/networking/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the DestinationRule Resource across clusters.
// implemented by the user
type MulticlusterDestinationRuleReconciler interface {
	ReconcileDestinationRule(clusterName string, obj *networking_istio_io_v1.DestinationRule) (reconcile.Result, error)
}

// Reconcile deletion events for the DestinationRule Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterDestinationRuleDeletionReconciler interface {
	ReconcileDestinationRuleDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterDestinationRuleReconcilerFuncs struct {
	OnReconcileDestinationRule         func(clusterName string, obj *networking_istio_io_v1.DestinationRule) (reconcile.Result, error)
	OnReconcileDestinationRuleDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterDestinationRuleReconcilerFuncs) ReconcileDestinationRule(clusterName string, obj *networking_istio_io_v1.DestinationRule) (reconcile.Result, error) {
	if f.OnReconcileDestinationRule == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileDestinationRule(clusterName, obj)
}

func (f *MulticlusterDestinationRuleReconcilerFuncs) ReconcileDestinationRuleDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileDestinationRuleDeletion == nil {
		return nil
	}
	return f.OnReconcileDestinationRuleDeletion(clusterName, req)
}

type MulticlusterDestinationRuleReconcileLoop interface {
	// AddMulticlusterDestinationRuleReconciler adds a MulticlusterDestinationRuleReconciler to the MulticlusterDestinationRuleReconcileLoop.
	AddMulticlusterDestinationRuleReconciler(ctx context.Context, rec MulticlusterDestinationRuleReconciler, predicates ...predicate.Predicate)
}

type multiclusterDestinationRuleReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterDestinationRuleReconcileLoop) AddMulticlusterDestinationRuleReconciler(ctx context.Context, rec MulticlusterDestinationRuleReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericDestinationRuleMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterDestinationRuleReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterDestinationRuleReconcileLoop {
	return &multiclusterDestinationRuleReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1.DestinationRule{}, options)}
}

type genericDestinationRuleMulticlusterReconciler struct {
	reconciler MulticlusterDestinationRuleReconciler
}

func (g genericDestinationRuleMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterDestinationRuleDeletionReconciler); ok {
		return deletionReconciler.ReconcileDestinationRuleDeletion(cluster, req)
	}
	return nil
}

func (g genericDestinationRuleMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1.DestinationRule)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: DestinationRule handler received event for %T", object)
	}
	return g.reconciler.ReconcileDestinationRule(cluster, obj)
}

// Reconcile Upsert events for the Gateway Resource across clusters.
// implemented by the user
type MulticlusterGatewayReconciler interface {
	ReconcileGateway(clusterName string, obj *networking_istio_io_v1.Gateway) (reconcile.Result, error)
}

// Reconcile deletion events for the Gateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGatewayDeletionReconciler interface {
	ReconcileGatewayDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGatewayReconcilerFuncs struct {
	OnReconcileGateway         func(clusterName string, obj *networking_istio_io_v1.Gateway) (reconcile.Result, error)
	OnReconcileGatewayDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGatewayReconcilerFuncs) ReconcileGateway(clusterName string, obj *networking_istio_io_v1.Gateway) (reconcile.Result, error) {
	if f.OnReconcileGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGateway(clusterName, obj)
}

func (f *MulticlusterGatewayReconcilerFuncs) ReconcileGatewayDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayDeletion(clusterName, req)
}

type MulticlusterGatewayReconcileLoop interface {
	// AddMulticlusterGatewayReconciler adds a MulticlusterGatewayReconciler to the MulticlusterGatewayReconcileLoop.
	AddMulticlusterGatewayReconciler(ctx context.Context, rec MulticlusterGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGatewayReconcileLoop) AddMulticlusterGatewayReconciler(ctx context.Context, rec MulticlusterGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGatewayReconcileLoop {
	return &multiclusterGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1.Gateway{}, options)}
}

type genericGatewayMulticlusterReconciler struct {
	reconciler MulticlusterGatewayReconciler
}

func (g genericGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayDeletion(cluster, req)
	}
	return nil
}

func (g genericGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1.Gateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileGateway(cluster, obj)
}

// Reconcile Upsert events for the ServiceEntry Resource across clusters.
// implemented by the user
type MulticlusterServiceEntryReconciler interface {
	ReconcileServiceEntry(clusterName string, obj *networking_istio_io_v1.ServiceEntry) (reconcile.Result, error)
}

// Reconcile deletion events for the ServiceEntry Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterServiceEntryDeletionReconciler interface {
	ReconcileServiceEntryDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterServiceEntryReconcilerFuncs struct {
	OnReconcileServiceEntry         func(clusterName string, obj *networking_istio_io_v1.ServiceEntry) (reconcile.Result, error)
	OnReconcileServiceEntryDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterServiceEntryReconcilerFuncs) ReconcileServiceEntry(clusterName string, obj *networking_istio_io_v1.ServiceEntry) (reconcile.Result, error) {
	if f.OnReconcileServiceEntry == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileServiceEntry(clusterName, obj)
}

func (f *MulticlusterServiceEntryReconcilerFuncs) ReconcileServiceEntryDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileServiceEntryDeletion == nil {
		return nil
	}
	return f.OnReconcileServiceEntryDeletion(clusterName, req)
}

type MulticlusterServiceEntryReconcileLoop interface {
	// AddMulticlusterServiceEntryReconciler adds a MulticlusterServiceEntryReconciler to the MulticlusterServiceEntryReconcileLoop.
	AddMulticlusterServiceEntryReconciler(ctx context.Context, rec MulticlusterServiceEntryReconciler, predicates ...predicate.Predicate)
}

type multiclusterServiceEntryReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterServiceEntryReconcileLoop) AddMulticlusterServiceEntryReconciler(ctx context.Context, rec MulticlusterServiceEntryReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericServiceEntryMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterServiceEntryReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterServiceEntryReconcileLoop {
	return &multiclusterServiceEntryReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1.ServiceEntry{}, options)}
}

type genericServiceEntryMulticlusterReconciler struct {
	reconciler MulticlusterServiceEntryReconciler
}

func (g genericServiceEntryMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterServiceEntryDeletionReconciler); ok {
		return deletionReconciler.ReconcileServiceEntryDeletion(cluster, req)
	}
	return nil
}

func (g genericServiceEntryMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1.ServiceEntry)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ServiceEntry handler received event for %T", object)
	}
	return g.reconciler.ReconcileServiceEntry(cluster, obj)
}

// Reconcile Upsert events for the WorkloadEntry Resource across clusters.
// implemented by the user
type MulticlusterWorkloadEntryReconciler interface {
	ReconcileWorkloadEntry(clusterName string, obj *networking_istio_io_v1.WorkloadEntry) (reconcile.Result, error)
}

// Reconcile deletion events for the WorkloadEntry Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterWorkloadEntryDeletionReconciler interface {
	ReconcileWorkloadEntryDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterWorkloadEntryReconcilerFuncs struct {
	OnReconcileWorkloadEntry         func(clusterName string, obj *networking_istio_io_v1.WorkloadEntry) (reconcile.Result, error)
	OnReconcileWorkloadEntryDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterWorkloadEntryReconcilerFuncs) ReconcileWorkloadEntry(clusterName string, obj *networking_istio_io_v1.WorkloadEntry) (reconcile.Result, error) {
	if f.OnReconcileWorkloadEntry == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWorkloadEntry(clusterName, obj)
}

func (f *MulticlusterWorkloadEntryReconcilerFuncs) ReconcileWorkloadEntryDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileWorkloadEntryDeletion == nil {
		return nil
	}
	return f.OnReconcileWorkloadEntryDeletion(clusterName, req)
}

type MulticlusterWorkloadEntryReconcileLoop interface {
	// AddMulticlusterWorkloadEntryReconciler adds a MulticlusterWorkloadEntryReconciler to the MulticlusterWorkloadEntryReconcileLoop.
	AddMulticlusterWorkloadEntryReconciler(ctx context.Context, rec MulticlusterWorkloadEntryReconciler, predicates ...predicate.Predicate)
}

type multiclusterWorkloadEntryReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterWorkloadEntryReconcileLoop) AddMulticlusterWorkloadEntryReconciler(ctx context.Context, rec MulticlusterWorkloadEntryReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericWorkloadEntryMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterWorkloadEntryReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterWorkloadEntryReconcileLoop {
	return &multiclusterWorkloadEntryReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1.WorkloadEntry{}, options)}
}

type genericWorkloadEntryMulticlusterReconciler struct {
	reconciler MulticlusterWorkloadEntryReconciler
}

func (g genericWorkloadEntryMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterWorkloadEntryDeletionReconciler); ok {
		return deletionReconciler.ReconcileWorkloadEntryDeletion(cluster, req)
	}
	return nil
}

func (g genericWorkloadEntryMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1.WorkloadEntry)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WorkloadEntry handler received event for %T", object)
	}
	return g.reconciler.ReconcileWorkloadEntry(cluster, obj)
}

// Reconcile Upsert events for the WorkloadGroup Resource across clusters.
// implemented by the user
type MulticlusterWorkloadGroupReconciler interface {
	ReconcileWorkloadGroup(clusterName string, obj *networking_istio_io_v1.WorkloadGroup) (reconcile.Result, error)
}

// Reconcile deletion events for the WorkloadGroup Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterWorkloadGroupDeletionReconciler interface {
	ReconcileWorkloadGroupDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterWorkloadGroupReconcilerFuncs struct {
	OnReconcileWorkloadGroup         func(clusterName string, obj *networking_istio_io_v1.WorkloadGroup) (reconcile.Result, error)
	OnReconcileWorkloadGroupDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterWorkloadGroupReconcilerFuncs) ReconcileWorkloadGroup(clusterName string, obj *networking_istio_io_v1.WorkloadGroup) (reconcile.Result, error) {
	if f.OnReconcileWorkloadGroup == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWorkloadGroup(clusterName, obj)
}

func (f *MulticlusterWorkloadGroupReconcilerFuncs) ReconcileWorkloadGroupDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileWorkloadGroupDeletion == nil {
		return nil
	}
	return f.OnReconcileWorkloadGroupDeletion(clusterName, req)
}

type MulticlusterWorkloadGroupReconcileLoop interface {
	// AddMulticlusterWorkloadGroupReconciler adds a MulticlusterWorkloadGroupReconciler to the MulticlusterWorkloadGroupReconcileLoop.
	AddMulticlusterWorkloadGroupReconciler(ctx context.Context, rec MulticlusterWorkloadGroupReconciler, predicates ...predicate.Predicate)
}

type multiclusterWorkloadGroupReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterWorkloadGroupReconcileLoop) AddMulticlusterWorkloadGroupReconciler(ctx context.Context, rec MulticlusterWorkloadGroupReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericWorkloadGroupMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterWorkloadGroupReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterWorkloadGroupReconcileLoop {
	return &multiclusterWorkloadGroupReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1.WorkloadGroup{}, options)}
}

type genericWorkloadGroupMulticlusterReconciler struct {
	reconciler MulticlusterWorkloadGroupReconciler
}

func (g genericWorkloadGroupMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterWorkloadGroupDeletionReconciler); ok {
		return deletionReconciler.ReconcileWorkloadGroupDeletion(cluster, req)
	}
	return nil
}

func (g genericWorkloadGroupMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1.WorkloadGroup)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WorkloadGroup handler received event for %T", object)
	}
	return g.reconciler.ReconcileWorkloadGroup(cluster, obj)
}

// Reconcile Upsert events for the VirtualService Resource across clusters.
// implemented by the user
type MulticlusterVirtualServiceReconciler interface {
	ReconcileVirtualService(clusterName string, obj *networking_istio_io_v1.VirtualService) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualService Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualServiceDeletionReconciler interface {
	ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterVirtualServiceReconcilerFuncs struct {
	OnReconcileVirtualService         func(clusterName string, obj *networking_istio_io_v1.VirtualService) (reconcile.Result, error)
	OnReconcileVirtualServiceDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterVirtualServiceReconcilerFuncs) ReconcileVirtualService(clusterName string, obj *networking_istio_io_v1.VirtualService) (reconcile.Result, error) {
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
	return &multiclusterVirtualServiceReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1.VirtualService{}, options)}
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
	obj, ok := object.(*networking_istio_io_v1.VirtualService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualService(cluster, obj)
}

// Reconcile Upsert events for the Sidecar Resource across clusters.
// implemented by the user
type MulticlusterSidecarReconciler interface {
	ReconcileSidecar(clusterName string, obj *networking_istio_io_v1.Sidecar) (reconcile.Result, error)
}

// Reconcile deletion events for the Sidecar Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterSidecarDeletionReconciler interface {
	ReconcileSidecarDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterSidecarReconcilerFuncs struct {
	OnReconcileSidecar         func(clusterName string, obj *networking_istio_io_v1.Sidecar) (reconcile.Result, error)
	OnReconcileSidecarDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterSidecarReconcilerFuncs) ReconcileSidecar(clusterName string, obj *networking_istio_io_v1.Sidecar) (reconcile.Result, error) {
	if f.OnReconcileSidecar == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSidecar(clusterName, obj)
}

func (f *MulticlusterSidecarReconcilerFuncs) ReconcileSidecarDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileSidecarDeletion == nil {
		return nil
	}
	return f.OnReconcileSidecarDeletion(clusterName, req)
}

type MulticlusterSidecarReconcileLoop interface {
	// AddMulticlusterSidecarReconciler adds a MulticlusterSidecarReconciler to the MulticlusterSidecarReconcileLoop.
	AddMulticlusterSidecarReconciler(ctx context.Context, rec MulticlusterSidecarReconciler, predicates ...predicate.Predicate)
}

type multiclusterSidecarReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterSidecarReconcileLoop) AddMulticlusterSidecarReconciler(ctx context.Context, rec MulticlusterSidecarReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericSidecarMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterSidecarReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterSidecarReconcileLoop {
	return &multiclusterSidecarReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1.Sidecar{}, options)}
}

type genericSidecarMulticlusterReconciler struct {
	reconciler MulticlusterSidecarReconciler
}

func (g genericSidecarMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterSidecarDeletionReconciler); ok {
		return deletionReconciler.ReconcileSidecarDeletion(cluster, req)
	}
	return nil
}

func (g genericSidecarMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1.Sidecar)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Sidecar handler received event for %T", object)
	}
	return g.reconciler.ReconcileSidecar(cluster, obj)
}