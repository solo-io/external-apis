// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	extensions_istio_io_v1alpha1 "istio.io/client-go/pkg/apis/extensions/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the WasmPlugin Resource across clusters.
// implemented by the user
type MulticlusterWasmPluginReconciler interface {
	ReconcileWasmPlugin(clusterName string, obj *extensions_istio_io_v1alpha1.WasmPlugin) (reconcile.Result, error)
}

// Reconcile deletion events for the WasmPlugin Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterWasmPluginDeletionReconciler interface {
	ReconcileWasmPluginDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterWasmPluginReconcilerFuncs struct {
	OnReconcileWasmPlugin         func(clusterName string, obj *extensions_istio_io_v1alpha1.WasmPlugin) (reconcile.Result, error)
	OnReconcileWasmPluginDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterWasmPluginReconcilerFuncs) ReconcileWasmPlugin(clusterName string, obj *extensions_istio_io_v1alpha1.WasmPlugin) (reconcile.Result, error) {
	if f.OnReconcileWasmPlugin == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWasmPlugin(clusterName, obj)
}

func (f *MulticlusterWasmPluginReconcilerFuncs) ReconcileWasmPluginDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileWasmPluginDeletion == nil {
		return nil
	}
	return f.OnReconcileWasmPluginDeletion(clusterName, req)
}

type MulticlusterWasmPluginReconcileLoop interface {
	// AddMulticlusterWasmPluginReconciler adds a MulticlusterWasmPluginReconciler to the MulticlusterWasmPluginReconcileLoop.
	AddMulticlusterWasmPluginReconciler(ctx context.Context, rec MulticlusterWasmPluginReconciler, predicates ...predicate.Predicate)
}

type multiclusterWasmPluginReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterWasmPluginReconcileLoop) AddMulticlusterWasmPluginReconciler(ctx context.Context, rec MulticlusterWasmPluginReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericWasmPluginMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterWasmPluginReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterWasmPluginReconcileLoop {
	return &multiclusterWasmPluginReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &extensions_istio_io_v1alpha1.WasmPlugin{}, options)}
}

type genericWasmPluginMulticlusterReconciler struct {
	reconciler MulticlusterWasmPluginReconciler
}

func (g genericWasmPluginMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterWasmPluginDeletionReconciler); ok {
		return deletionReconciler.ReconcileWasmPluginDeletion(cluster, req)
	}
	return nil
}

func (g genericWasmPluginMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*extensions_istio_io_v1alpha1.WasmPlugin)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WasmPlugin handler received event for %T", object)
	}
	return g.reconciler.ReconcileWasmPlugin(cluster, obj)
}
