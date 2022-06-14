// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	cilium_io_v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the CiliumNetworkPolicy Resource.
// implemented by the user
type CiliumNetworkPolicyReconciler interface {
	ReconcileCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the CiliumNetworkPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type CiliumNetworkPolicyDeletionReconciler interface {
	ReconcileCiliumNetworkPolicyDeletion(req reconcile.Request) error
}

type CiliumNetworkPolicyReconcilerFuncs struct {
	OnReconcileCiliumNetworkPolicy         func(obj *cilium_io_v2.CiliumNetworkPolicy) (reconcile.Result, error)
	OnReconcileCiliumNetworkPolicyDeletion func(req reconcile.Request) error
}

func (f *CiliumNetworkPolicyReconcilerFuncs) ReconcileCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) (reconcile.Result, error) {
	if f.OnReconcileCiliumNetworkPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCiliumNetworkPolicy(obj)
}

func (f *CiliumNetworkPolicyReconcilerFuncs) ReconcileCiliumNetworkPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileCiliumNetworkPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileCiliumNetworkPolicyDeletion(req)
}

// Reconcile and finalize the CiliumNetworkPolicy Resource
// implemented by the user
type CiliumNetworkPolicyFinalizer interface {
	CiliumNetworkPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	CiliumNetworkPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) error
}

type CiliumNetworkPolicyReconcileLoop interface {
	RunCiliumNetworkPolicyReconciler(ctx context.Context, rec CiliumNetworkPolicyReconciler, predicates ...predicate.Predicate) error
}

type ciliumNetworkPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewCiliumNetworkPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) CiliumNetworkPolicyReconcileLoop {
	return &ciliumNetworkPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &cilium_io_v2.CiliumNetworkPolicy{}, options),
	}
}

func (c *ciliumNetworkPolicyReconcileLoop) RunCiliumNetworkPolicyReconciler(ctx context.Context, reconciler CiliumNetworkPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericCiliumNetworkPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(CiliumNetworkPolicyFinalizer); ok {
		reconcilerWrapper = genericCiliumNetworkPolicyFinalizer{
			genericCiliumNetworkPolicyReconciler: genericReconciler,
			finalizingReconciler:                 finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericCiliumNetworkPolicyHandler implements a generic reconcile.Reconciler
type genericCiliumNetworkPolicyReconciler struct {
	reconciler CiliumNetworkPolicyReconciler
}

func (r genericCiliumNetworkPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*cilium_io_v2.CiliumNetworkPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CiliumNetworkPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileCiliumNetworkPolicy(obj)
}

func (r genericCiliumNetworkPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(CiliumNetworkPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileCiliumNetworkPolicyDeletion(request)
	}
	return nil
}

// genericCiliumNetworkPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericCiliumNetworkPolicyFinalizer struct {
	genericCiliumNetworkPolicyReconciler
	finalizingReconciler CiliumNetworkPolicyFinalizer
}

func (r genericCiliumNetworkPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.CiliumNetworkPolicyFinalizerName()
}

func (r genericCiliumNetworkPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*cilium_io_v2.CiliumNetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: CiliumNetworkPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeCiliumNetworkPolicy(obj)
}
