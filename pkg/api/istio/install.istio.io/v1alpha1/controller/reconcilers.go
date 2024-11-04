// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	install_istio_io_v1alpha1 "istio.io/istio/operator/pkg/apis/istio/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the IstioOperator Resource.
// implemented by the user
type IstioOperatorReconciler interface {
	ReconcileIstioOperator(obj *install_istio_io_v1alpha1.IstioOperator) (reconcile.Result, error)
}

// Reconcile deletion events for the IstioOperator Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type IstioOperatorDeletionReconciler interface {
	ReconcileIstioOperatorDeletion(req reconcile.Request) error
}

type IstioOperatorReconcilerFuncs struct {
	OnReconcileIstioOperator         func(obj *install_istio_io_v1alpha1.IstioOperator) (reconcile.Result, error)
	OnReconcileIstioOperatorDeletion func(req reconcile.Request) error
}

func (f *IstioOperatorReconcilerFuncs) ReconcileIstioOperator(obj *install_istio_io_v1alpha1.IstioOperator) (reconcile.Result, error) {
	if f.OnReconcileIstioOperator == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileIstioOperator(obj)
}

func (f *IstioOperatorReconcilerFuncs) ReconcileIstioOperatorDeletion(req reconcile.Request) error {
	if f.OnReconcileIstioOperatorDeletion == nil {
		return nil
	}
	return f.OnReconcileIstioOperatorDeletion(req)
}

// Reconcile and finalize the IstioOperator Resource
// implemented by the user
type IstioOperatorFinalizer interface {
	IstioOperatorReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	IstioOperatorFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeIstioOperator(obj *install_istio_io_v1alpha1.IstioOperator) error
}

type IstioOperatorReconcileLoop interface {
	RunIstioOperatorReconciler(ctx context.Context, rec IstioOperatorReconciler, predicates ...predicate.Predicate) error
}

type istioOperatorReconcileLoop struct {
	loop reconcile.Loop
}

func NewIstioOperatorReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) IstioOperatorReconcileLoop {
	return &istioOperatorReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &install_istio_io_v1alpha1.IstioOperator{}, options),
	}
}

func (c *istioOperatorReconcileLoop) RunIstioOperatorReconciler(ctx context.Context, reconciler IstioOperatorReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericIstioOperatorReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(IstioOperatorFinalizer); ok {
		reconcilerWrapper = genericIstioOperatorFinalizer{
			genericIstioOperatorReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericIstioOperatorHandler implements a generic reconcile.Reconciler
type genericIstioOperatorReconciler struct {
	reconciler IstioOperatorReconciler
}

func (r genericIstioOperatorReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*install_istio_io_v1alpha1.IstioOperator)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: IstioOperator handler received event for %T", object)
	}
	return r.reconciler.ReconcileIstioOperator(obj)
}

func (r genericIstioOperatorReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(IstioOperatorDeletionReconciler); ok {
		return deletionReconciler.ReconcileIstioOperatorDeletion(request)
	}
	return nil
}

// genericIstioOperatorFinalizer implements a generic reconcile.FinalizingReconciler
type genericIstioOperatorFinalizer struct {
	genericIstioOperatorReconciler
	finalizingReconciler IstioOperatorFinalizer
}

func (r genericIstioOperatorFinalizer) FinalizerName() string {
	return r.finalizingReconciler.IstioOperatorFinalizerName()
}

func (r genericIstioOperatorFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*install_istio_io_v1alpha1.IstioOperator)
	if !ok {
		return errors.Errorf("internal error: IstioOperator handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeIstioOperator(obj)
}
