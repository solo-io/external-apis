// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_k8s_io_v1 "k8s.io/api/networking/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the NetworkPolicy Resource.
// implemented by the user
type NetworkPolicyReconciler interface {
	ReconcileNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the NetworkPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type NetworkPolicyDeletionReconciler interface {
	ReconcileNetworkPolicyDeletion(req reconcile.Request) error
}

type NetworkPolicyReconcilerFuncs struct {
	OnReconcileNetworkPolicy         func(obj *networking_k8s_io_v1.NetworkPolicy) (reconcile.Result, error)
	OnReconcileNetworkPolicyDeletion func(req reconcile.Request) error
}

func (f *NetworkPolicyReconcilerFuncs) ReconcileNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) (reconcile.Result, error) {
	if f.OnReconcileNetworkPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileNetworkPolicy(obj)
}

func (f *NetworkPolicyReconcilerFuncs) ReconcileNetworkPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileNetworkPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileNetworkPolicyDeletion(req)
}

// Reconcile and finalize the NetworkPolicy Resource
// implemented by the user
type NetworkPolicyFinalizer interface {
	NetworkPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	NetworkPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) error
}

type NetworkPolicyReconcileLoop interface {
	RunNetworkPolicyReconciler(ctx context.Context, rec NetworkPolicyReconciler, predicates ...predicate.Predicate) error
}

type networkPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewNetworkPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) NetworkPolicyReconcileLoop {
	return &networkPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_k8s_io_v1.NetworkPolicy{}, options),
	}
}

func (c *networkPolicyReconcileLoop) RunNetworkPolicyReconciler(ctx context.Context, reconciler NetworkPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericNetworkPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(NetworkPolicyFinalizer); ok {
		reconcilerWrapper = genericNetworkPolicyFinalizer{
			genericNetworkPolicyReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericNetworkPolicyHandler implements a generic reconcile.Reconciler
type genericNetworkPolicyReconciler struct {
	reconciler NetworkPolicyReconciler
}

func (r genericNetworkPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_k8s_io_v1.NetworkPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: NetworkPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileNetworkPolicy(obj)
}

func (r genericNetworkPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(NetworkPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileNetworkPolicyDeletion(request)
	}
	return nil
}

// genericNetworkPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericNetworkPolicyFinalizer struct {
	genericNetworkPolicyReconciler
	finalizingReconciler NetworkPolicyFinalizer
}

func (r genericNetworkPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.NetworkPolicyFinalizerName()
}

func (r genericNetworkPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_k8s_io_v1.NetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: NetworkPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeNetworkPolicy(obj)
}
