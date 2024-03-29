// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	admissionregistration_k8s_io_v1 "k8s.io/api/admissionregistration/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the ValidatingWebhookConfiguration Resource.
// implemented by the user
type ValidatingWebhookConfigurationReconciler interface {
	ReconcileValidatingWebhookConfiguration(obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) (reconcile.Result, error)
}

// Reconcile deletion events for the ValidatingWebhookConfiguration Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ValidatingWebhookConfigurationDeletionReconciler interface {
	ReconcileValidatingWebhookConfigurationDeletion(req reconcile.Request) error
}

type ValidatingWebhookConfigurationReconcilerFuncs struct {
	OnReconcileValidatingWebhookConfiguration         func(obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) (reconcile.Result, error)
	OnReconcileValidatingWebhookConfigurationDeletion func(req reconcile.Request) error
}

func (f *ValidatingWebhookConfigurationReconcilerFuncs) ReconcileValidatingWebhookConfiguration(obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) (reconcile.Result, error) {
	if f.OnReconcileValidatingWebhookConfiguration == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileValidatingWebhookConfiguration(obj)
}

func (f *ValidatingWebhookConfigurationReconcilerFuncs) ReconcileValidatingWebhookConfigurationDeletion(req reconcile.Request) error {
	if f.OnReconcileValidatingWebhookConfigurationDeletion == nil {
		return nil
	}
	return f.OnReconcileValidatingWebhookConfigurationDeletion(req)
}

// Reconcile and finalize the ValidatingWebhookConfiguration Resource
// implemented by the user
type ValidatingWebhookConfigurationFinalizer interface {
	ValidatingWebhookConfigurationReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ValidatingWebhookConfigurationFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeValidatingWebhookConfiguration(obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) error
}

type ValidatingWebhookConfigurationReconcileLoop interface {
	RunValidatingWebhookConfigurationReconciler(ctx context.Context, rec ValidatingWebhookConfigurationReconciler, predicates ...predicate.Predicate) error
}

type validatingWebhookConfigurationReconcileLoop struct {
	loop reconcile.Loop
}

func NewValidatingWebhookConfigurationReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) ValidatingWebhookConfigurationReconcileLoop {
	return &validatingWebhookConfigurationReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration{}, options),
	}
}

func (c *validatingWebhookConfigurationReconcileLoop) RunValidatingWebhookConfigurationReconciler(ctx context.Context, reconciler ValidatingWebhookConfigurationReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericValidatingWebhookConfigurationReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ValidatingWebhookConfigurationFinalizer); ok {
		reconcilerWrapper = genericValidatingWebhookConfigurationFinalizer{
			genericValidatingWebhookConfigurationReconciler: genericReconciler,
			finalizingReconciler:                            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericValidatingWebhookConfigurationHandler implements a generic reconcile.Reconciler
type genericValidatingWebhookConfigurationReconciler struct {
	reconciler ValidatingWebhookConfigurationReconciler
}

func (r genericValidatingWebhookConfigurationReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ValidatingWebhookConfiguration handler received event for %T", object)
	}
	return r.reconciler.ReconcileValidatingWebhookConfiguration(obj)
}

func (r genericValidatingWebhookConfigurationReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(ValidatingWebhookConfigurationDeletionReconciler); ok {
		return deletionReconciler.ReconcileValidatingWebhookConfigurationDeletion(request)
	}
	return nil
}

// genericValidatingWebhookConfigurationFinalizer implements a generic reconcile.FinalizingReconciler
type genericValidatingWebhookConfigurationFinalizer struct {
	genericValidatingWebhookConfigurationReconciler
	finalizingReconciler ValidatingWebhookConfigurationFinalizer
}

func (r genericValidatingWebhookConfigurationFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ValidatingWebhookConfigurationFinalizerName()
}

func (r genericValidatingWebhookConfigurationFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)
	if !ok {
		return errors.Errorf("internal error: ValidatingWebhookConfiguration handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeValidatingWebhookConfiguration(obj)
}

// Reconcile Upsert events for the MutatingWebhookConfiguration Resource.
// implemented by the user
type MutatingWebhookConfigurationReconciler interface {
	ReconcileMutatingWebhookConfiguration(obj *admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) (reconcile.Result, error)
}

// Reconcile deletion events for the MutatingWebhookConfiguration Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MutatingWebhookConfigurationDeletionReconciler interface {
	ReconcileMutatingWebhookConfigurationDeletion(req reconcile.Request) error
}

type MutatingWebhookConfigurationReconcilerFuncs struct {
	OnReconcileMutatingWebhookConfiguration         func(obj *admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) (reconcile.Result, error)
	OnReconcileMutatingWebhookConfigurationDeletion func(req reconcile.Request) error
}

func (f *MutatingWebhookConfigurationReconcilerFuncs) ReconcileMutatingWebhookConfiguration(obj *admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) (reconcile.Result, error) {
	if f.OnReconcileMutatingWebhookConfiguration == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMutatingWebhookConfiguration(obj)
}

func (f *MutatingWebhookConfigurationReconcilerFuncs) ReconcileMutatingWebhookConfigurationDeletion(req reconcile.Request) error {
	if f.OnReconcileMutatingWebhookConfigurationDeletion == nil {
		return nil
	}
	return f.OnReconcileMutatingWebhookConfigurationDeletion(req)
}

// Reconcile and finalize the MutatingWebhookConfiguration Resource
// implemented by the user
type MutatingWebhookConfigurationFinalizer interface {
	MutatingWebhookConfigurationReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	MutatingWebhookConfigurationFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeMutatingWebhookConfiguration(obj *admissionregistration_k8s_io_v1.MutatingWebhookConfiguration) error
}

type MutatingWebhookConfigurationReconcileLoop interface {
	RunMutatingWebhookConfigurationReconciler(ctx context.Context, rec MutatingWebhookConfigurationReconciler, predicates ...predicate.Predicate) error
}

type mutatingWebhookConfigurationReconcileLoop struct {
	loop reconcile.Loop
}

func NewMutatingWebhookConfigurationReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) MutatingWebhookConfigurationReconcileLoop {
	return &mutatingWebhookConfigurationReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &admissionregistration_k8s_io_v1.MutatingWebhookConfiguration{}, options),
	}
}

func (c *mutatingWebhookConfigurationReconcileLoop) RunMutatingWebhookConfigurationReconciler(ctx context.Context, reconciler MutatingWebhookConfigurationReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericMutatingWebhookConfigurationReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(MutatingWebhookConfigurationFinalizer); ok {
		reconcilerWrapper = genericMutatingWebhookConfigurationFinalizer{
			genericMutatingWebhookConfigurationReconciler: genericReconciler,
			finalizingReconciler:                          finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericMutatingWebhookConfigurationHandler implements a generic reconcile.Reconciler
type genericMutatingWebhookConfigurationReconciler struct {
	reconciler MutatingWebhookConfigurationReconciler
}

func (r genericMutatingWebhookConfigurationReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: MutatingWebhookConfiguration handler received event for %T", object)
	}
	return r.reconciler.ReconcileMutatingWebhookConfiguration(obj)
}

func (r genericMutatingWebhookConfigurationReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(MutatingWebhookConfigurationDeletionReconciler); ok {
		return deletionReconciler.ReconcileMutatingWebhookConfigurationDeletion(request)
	}
	return nil
}

// genericMutatingWebhookConfigurationFinalizer implements a generic reconcile.FinalizingReconciler
type genericMutatingWebhookConfigurationFinalizer struct {
	genericMutatingWebhookConfigurationReconciler
	finalizingReconciler MutatingWebhookConfigurationFinalizer
}

func (r genericMutatingWebhookConfigurationFinalizer) FinalizerName() string {
	return r.finalizingReconciler.MutatingWebhookConfigurationFinalizerName()
}

func (r genericMutatingWebhookConfigurationFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*admissionregistration_k8s_io_v1.MutatingWebhookConfiguration)
	if !ok {
		return errors.Errorf("internal error: MutatingWebhookConfiguration handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeMutatingWebhookConfiguration(obj)
}
