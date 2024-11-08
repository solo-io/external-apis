// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	specs_smi_spec_io_v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the HTTPRouteGroup Resource.
// implemented by the user
type HTTPRouteGroupReconciler interface {
	ReconcileHTTPRouteGroup(obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup) (reconcile.Result, error)
}

// Reconcile deletion events for the HTTPRouteGroup Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type HTTPRouteGroupDeletionReconciler interface {
	ReconcileHTTPRouteGroupDeletion(req reconcile.Request) error
}

type HTTPRouteGroupReconcilerFuncs struct {
	OnReconcileHTTPRouteGroup         func(obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup) (reconcile.Result, error)
	OnReconcileHTTPRouteGroupDeletion func(req reconcile.Request) error
}

func (f *HTTPRouteGroupReconcilerFuncs) ReconcileHTTPRouteGroup(obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup) (reconcile.Result, error) {
	if f.OnReconcileHTTPRouteGroup == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileHTTPRouteGroup(obj)
}

func (f *HTTPRouteGroupReconcilerFuncs) ReconcileHTTPRouteGroupDeletion(req reconcile.Request) error {
	if f.OnReconcileHTTPRouteGroupDeletion == nil {
		return nil
	}
	return f.OnReconcileHTTPRouteGroupDeletion(req)
}

// Reconcile and finalize the HTTPRouteGroup Resource
// implemented by the user
type HTTPRouteGroupFinalizer interface {
	HTTPRouteGroupReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	HTTPRouteGroupFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeHTTPRouteGroup(obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup) error
}

type HTTPRouteGroupReconcileLoop interface {
	RunHTTPRouteGroupReconciler(ctx context.Context, rec HTTPRouteGroupReconciler, predicates ...predicate.Predicate) error
}

type httprouteGroupReconcileLoop struct {
	loop reconcile.Loop
}

func NewHTTPRouteGroupReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) HTTPRouteGroupReconcileLoop {
	return &httprouteGroupReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}, options),
	}
}

func (c *httprouteGroupReconcileLoop) RunHTTPRouteGroupReconciler(ctx context.Context, reconciler HTTPRouteGroupReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericHTTPRouteGroupReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(HTTPRouteGroupFinalizer); ok {
		reconcilerWrapper = genericHTTPRouteGroupFinalizer{
			genericHTTPRouteGroupReconciler: genericReconciler,
			finalizingReconciler:            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericHTTPRouteGroupHandler implements a generic reconcile.Reconciler
type genericHTTPRouteGroupReconciler struct {
	reconciler HTTPRouteGroupReconciler
}

func (r genericHTTPRouteGroupReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: HTTPRouteGroup handler received event for %T", object)
	}
	return r.reconciler.ReconcileHTTPRouteGroup(obj)
}

func (r genericHTTPRouteGroupReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(HTTPRouteGroupDeletionReconciler); ok {
		return deletionReconciler.ReconcileHTTPRouteGroupDeletion(request)
	}
	return nil
}

// genericHTTPRouteGroupFinalizer implements a generic reconcile.FinalizingReconciler
type genericHTTPRouteGroupFinalizer struct {
	genericHTTPRouteGroupReconciler
	finalizingReconciler HTTPRouteGroupFinalizer
}

func (r genericHTTPRouteGroupFinalizer) FinalizerName() string {
	return r.finalizingReconciler.HTTPRouteGroupFinalizerName()
}

func (r genericHTTPRouteGroupFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	if !ok {
		return errors.Errorf("internal error: HTTPRouteGroup handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeHTTPRouteGroup(obj)
}
