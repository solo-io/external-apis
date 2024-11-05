// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	install_istio_io_v1alpha1 "github.com/jehawley/istio/operator/pkg/apis/istio/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the IstioOperator Resource across clusters.
// implemented by the user
type MulticlusterIstioOperatorReconciler interface {
	ReconcileIstioOperator(clusterName string, obj *install_istio_io_v1alpha1.IstioOperator) (reconcile.Result, error)
}

// Reconcile deletion events for the IstioOperator Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterIstioOperatorDeletionReconciler interface {
	ReconcileIstioOperatorDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterIstioOperatorReconcilerFuncs struct {
	OnReconcileIstioOperator         func(clusterName string, obj *install_istio_io_v1alpha1.IstioOperator) (reconcile.Result, error)
	OnReconcileIstioOperatorDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterIstioOperatorReconcilerFuncs) ReconcileIstioOperator(clusterName string, obj *install_istio_io_v1alpha1.IstioOperator) (reconcile.Result, error) {
	if f.OnReconcileIstioOperator == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileIstioOperator(clusterName, obj)
}

func (f *MulticlusterIstioOperatorReconcilerFuncs) ReconcileIstioOperatorDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileIstioOperatorDeletion == nil {
		return nil
	}
	return f.OnReconcileIstioOperatorDeletion(clusterName, req)
}

type MulticlusterIstioOperatorReconcileLoop interface {
	// AddMulticlusterIstioOperatorReconciler adds a MulticlusterIstioOperatorReconciler to the MulticlusterIstioOperatorReconcileLoop.
	AddMulticlusterIstioOperatorReconciler(ctx context.Context, rec MulticlusterIstioOperatorReconciler, predicates ...predicate.Predicate)
}

type multiclusterIstioOperatorReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterIstioOperatorReconcileLoop) AddMulticlusterIstioOperatorReconciler(ctx context.Context, rec MulticlusterIstioOperatorReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericIstioOperatorMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterIstioOperatorReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterIstioOperatorReconcileLoop {
	return &multiclusterIstioOperatorReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &install_istio_io_v1alpha1.IstioOperator{}, options)}
}

type genericIstioOperatorMulticlusterReconciler struct {
	reconciler MulticlusterIstioOperatorReconciler
}

func (g genericIstioOperatorMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterIstioOperatorDeletionReconciler); ok {
		return deletionReconciler.ReconcileIstioOperatorDeletion(cluster, req)
	}
	return nil
}

func (g genericIstioOperatorMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*install_istio_io_v1alpha1.IstioOperator)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: IstioOperator handler received event for %T", object)
	}
	return g.reconciler.ReconcileIstioOperator(cluster, obj)
}
