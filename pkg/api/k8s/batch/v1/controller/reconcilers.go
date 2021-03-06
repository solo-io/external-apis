// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	batch_v1 "k8s.io/api/batch/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Job Resource.
// implemented by the user
type JobReconciler interface {
	ReconcileJob(obj *batch_v1.Job) (reconcile.Result, error)
}

// Reconcile deletion events for the Job Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type JobDeletionReconciler interface {
	ReconcileJobDeletion(req reconcile.Request) error
}

type JobReconcilerFuncs struct {
	OnReconcileJob         func(obj *batch_v1.Job) (reconcile.Result, error)
	OnReconcileJobDeletion func(req reconcile.Request) error
}

func (f *JobReconcilerFuncs) ReconcileJob(obj *batch_v1.Job) (reconcile.Result, error) {
	if f.OnReconcileJob == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileJob(obj)
}

func (f *JobReconcilerFuncs) ReconcileJobDeletion(req reconcile.Request) error {
	if f.OnReconcileJobDeletion == nil {
		return nil
	}
	return f.OnReconcileJobDeletion(req)
}

// Reconcile and finalize the Job Resource
// implemented by the user
type JobFinalizer interface {
	JobReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	JobFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeJob(obj *batch_v1.Job) error
}

type JobReconcileLoop interface {
	RunJobReconciler(ctx context.Context, rec JobReconciler, predicates ...predicate.Predicate) error
}

type jobReconcileLoop struct {
	loop reconcile.Loop
}

func NewJobReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) JobReconcileLoop {
	return &jobReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &batch_v1.Job{}, options),
	}
}

func (c *jobReconcileLoop) RunJobReconciler(ctx context.Context, reconciler JobReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericJobReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(JobFinalizer); ok {
		reconcilerWrapper = genericJobFinalizer{
			genericJobReconciler: genericReconciler,
			finalizingReconciler: finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericJobHandler implements a generic reconcile.Reconciler
type genericJobReconciler struct {
	reconciler JobReconciler
}

func (r genericJobReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*batch_v1.Job)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Job handler received event for %T", object)
	}
	return r.reconciler.ReconcileJob(obj)
}

func (r genericJobReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(JobDeletionReconciler); ok {
		return deletionReconciler.ReconcileJobDeletion(request)
	}
	return nil
}

// genericJobFinalizer implements a generic reconcile.FinalizingReconciler
type genericJobFinalizer struct {
	genericJobReconciler
	finalizingReconciler JobFinalizer
}

func (r genericJobFinalizer) FinalizerName() string {
	return r.finalizingReconciler.JobFinalizerName()
}

func (r genericJobFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*batch_v1.Job)
	if !ok {
		return errors.Errorf("internal error: Job handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeJob(obj)
}
