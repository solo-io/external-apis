// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	telemetry_istio_io_v1 "istio.io/client-go/pkg/apis/telemetry/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Telemetry Resource.
// implemented by the user
type TelemetryReconciler interface {
	ReconcileTelemetry(obj *telemetry_istio_io_v1.Telemetry) (reconcile.Result, error)
}

// Reconcile deletion events for the Telemetry Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type TelemetryDeletionReconciler interface {
	ReconcileTelemetryDeletion(req reconcile.Request) error
}

type TelemetryReconcilerFuncs struct {
	OnReconcileTelemetry         func(obj *telemetry_istio_io_v1.Telemetry) (reconcile.Result, error)
	OnReconcileTelemetryDeletion func(req reconcile.Request) error
}

func (f *TelemetryReconcilerFuncs) ReconcileTelemetry(obj *telemetry_istio_io_v1.Telemetry) (reconcile.Result, error) {
	if f.OnReconcileTelemetry == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileTelemetry(obj)
}

func (f *TelemetryReconcilerFuncs) ReconcileTelemetryDeletion(req reconcile.Request) error {
	if f.OnReconcileTelemetryDeletion == nil {
		return nil
	}
	return f.OnReconcileTelemetryDeletion(req)
}

// Reconcile and finalize the Telemetry Resource
// implemented by the user
type TelemetryFinalizer interface {
	TelemetryReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	TelemetryFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeTelemetry(obj *telemetry_istio_io_v1.Telemetry) error
}

type TelemetryReconcileLoop interface {
	RunTelemetryReconciler(ctx context.Context, rec TelemetryReconciler, predicates ...predicate.Predicate) error
}

type telemetryReconcileLoop struct {
	loop reconcile.Loop
}

func NewTelemetryReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) TelemetryReconcileLoop {
	return &telemetryReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &telemetry_istio_io_v1.Telemetry{}, options),
	}
}

func (c *telemetryReconcileLoop) RunTelemetryReconciler(ctx context.Context, reconciler TelemetryReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericTelemetryReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(TelemetryFinalizer); ok {
		reconcilerWrapper = genericTelemetryFinalizer{
			genericTelemetryReconciler: genericReconciler,
			finalizingReconciler:       finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericTelemetryHandler implements a generic reconcile.Reconciler
type genericTelemetryReconciler struct {
	reconciler TelemetryReconciler
}

func (r genericTelemetryReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*telemetry_istio_io_v1.Telemetry)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Telemetry handler received event for %T", object)
	}
	return r.reconciler.ReconcileTelemetry(obj)
}

func (r genericTelemetryReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(TelemetryDeletionReconciler); ok {
		return deletionReconciler.ReconcileTelemetryDeletion(request)
	}
	return nil
}

// genericTelemetryFinalizer implements a generic reconcile.FinalizingReconciler
type genericTelemetryFinalizer struct {
	genericTelemetryReconciler
	finalizingReconciler TelemetryFinalizer
}

func (r genericTelemetryFinalizer) FinalizerName() string {
	return r.finalizingReconciler.TelemetryFinalizerName()
}

func (r genericTelemetryFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*telemetry_istio_io_v1.Telemetry)
	if !ok {
		return errors.Errorf("internal error: Telemetry handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeTelemetry(obj)
}
