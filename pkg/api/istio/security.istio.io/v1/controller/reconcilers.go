// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	security_istio_io_v1 "istio.io/client-go/pkg/apis/security/v1beta1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AuthorizationPolicy Resource.
// implemented by the user
type AuthorizationPolicyReconciler interface {
	ReconcileAuthorizationPolicy(obj *security_istio_io_v1.AuthorizationPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AuthorizationPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type AuthorizationPolicyDeletionReconciler interface {
	ReconcileAuthorizationPolicyDeletion(req reconcile.Request) error
}

type AuthorizationPolicyReconcilerFuncs struct {
	OnReconcileAuthorizationPolicy         func(obj *security_istio_io_v1.AuthorizationPolicy) (reconcile.Result, error)
	OnReconcileAuthorizationPolicyDeletion func(req reconcile.Request) error
}

func (f *AuthorizationPolicyReconcilerFuncs) ReconcileAuthorizationPolicy(obj *security_istio_io_v1.AuthorizationPolicy) (reconcile.Result, error) {
	if f.OnReconcileAuthorizationPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAuthorizationPolicy(obj)
}

func (f *AuthorizationPolicyReconcilerFuncs) ReconcileAuthorizationPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileAuthorizationPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileAuthorizationPolicyDeletion(req)
}

// Reconcile and finalize the AuthorizationPolicy Resource
// implemented by the user
type AuthorizationPolicyFinalizer interface {
	AuthorizationPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	AuthorizationPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeAuthorizationPolicy(obj *security_istio_io_v1.AuthorizationPolicy) error
}

type AuthorizationPolicyReconcileLoop interface {
	RunAuthorizationPolicyReconciler(ctx context.Context, rec AuthorizationPolicyReconciler, predicates ...predicate.Predicate) error
}

type authorizationPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewAuthorizationPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) AuthorizationPolicyReconcileLoop {
	return &authorizationPolicyReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_istio_io_v1.AuthorizationPolicy{}, options),
	}
}

func (c *authorizationPolicyReconcileLoop) RunAuthorizationPolicyReconciler(ctx context.Context, reconciler AuthorizationPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericAuthorizationPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(AuthorizationPolicyFinalizer); ok {
		reconcilerWrapper = genericAuthorizationPolicyFinalizer{
			genericAuthorizationPolicyReconciler: genericReconciler,
			finalizingReconciler:                 finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericAuthorizationPolicyHandler implements a generic reconcile.Reconciler
type genericAuthorizationPolicyReconciler struct {
	reconciler AuthorizationPolicyReconciler
}

func (r genericAuthorizationPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_istio_io_v1.AuthorizationPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileAuthorizationPolicy(obj)
}

func (r genericAuthorizationPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(AuthorizationPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileAuthorizationPolicyDeletion(request)
	}
	return nil
}

// genericAuthorizationPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericAuthorizationPolicyFinalizer struct {
	genericAuthorizationPolicyReconciler
	finalizingReconciler AuthorizationPolicyFinalizer
}

func (r genericAuthorizationPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.AuthorizationPolicyFinalizerName()
}

func (r genericAuthorizationPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_istio_io_v1.AuthorizationPolicy)
	if !ok {
		return errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeAuthorizationPolicy(obj)
}

// Reconcile Upsert events for the PeerAuthentication Resource.
// implemented by the user
type PeerAuthenticationReconciler interface {
	ReconcilePeerAuthentication(obj *security_istio_io_v1.PeerAuthentication) (reconcile.Result, error)
}

// Reconcile deletion events for the PeerAuthentication Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type PeerAuthenticationDeletionReconciler interface {
	ReconcilePeerAuthenticationDeletion(req reconcile.Request) error
}

type PeerAuthenticationReconcilerFuncs struct {
	OnReconcilePeerAuthentication         func(obj *security_istio_io_v1.PeerAuthentication) (reconcile.Result, error)
	OnReconcilePeerAuthenticationDeletion func(req reconcile.Request) error
}

func (f *PeerAuthenticationReconcilerFuncs) ReconcilePeerAuthentication(obj *security_istio_io_v1.PeerAuthentication) (reconcile.Result, error) {
	if f.OnReconcilePeerAuthentication == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcilePeerAuthentication(obj)
}

func (f *PeerAuthenticationReconcilerFuncs) ReconcilePeerAuthenticationDeletion(req reconcile.Request) error {
	if f.OnReconcilePeerAuthenticationDeletion == nil {
		return nil
	}
	return f.OnReconcilePeerAuthenticationDeletion(req)
}

// Reconcile and finalize the PeerAuthentication Resource
// implemented by the user
type PeerAuthenticationFinalizer interface {
	PeerAuthenticationReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	PeerAuthenticationFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizePeerAuthentication(obj *security_istio_io_v1.PeerAuthentication) error
}

type PeerAuthenticationReconcileLoop interface {
	RunPeerAuthenticationReconciler(ctx context.Context, rec PeerAuthenticationReconciler, predicates ...predicate.Predicate) error
}

type peerAuthenticationReconcileLoop struct {
	loop reconcile.Loop
}

func NewPeerAuthenticationReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) PeerAuthenticationReconcileLoop {
	return &peerAuthenticationReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_istio_io_v1.PeerAuthentication{}, options),
	}
}

func (c *peerAuthenticationReconcileLoop) RunPeerAuthenticationReconciler(ctx context.Context, reconciler PeerAuthenticationReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericPeerAuthenticationReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(PeerAuthenticationFinalizer); ok {
		reconcilerWrapper = genericPeerAuthenticationFinalizer{
			genericPeerAuthenticationReconciler: genericReconciler,
			finalizingReconciler:                finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericPeerAuthenticationHandler implements a generic reconcile.Reconciler
type genericPeerAuthenticationReconciler struct {
	reconciler PeerAuthenticationReconciler
}

func (r genericPeerAuthenticationReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_istio_io_v1.PeerAuthentication)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: PeerAuthentication handler received event for %T", object)
	}
	return r.reconciler.ReconcilePeerAuthentication(obj)
}

func (r genericPeerAuthenticationReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(PeerAuthenticationDeletionReconciler); ok {
		return deletionReconciler.ReconcilePeerAuthenticationDeletion(request)
	}
	return nil
}

// genericPeerAuthenticationFinalizer implements a generic reconcile.FinalizingReconciler
type genericPeerAuthenticationFinalizer struct {
	genericPeerAuthenticationReconciler
	finalizingReconciler PeerAuthenticationFinalizer
}

func (r genericPeerAuthenticationFinalizer) FinalizerName() string {
	return r.finalizingReconciler.PeerAuthenticationFinalizerName()
}

func (r genericPeerAuthenticationFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_istio_io_v1.PeerAuthentication)
	if !ok {
		return errors.Errorf("internal error: PeerAuthentication handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizePeerAuthentication(obj)
}

// Reconcile Upsert events for the RequestAuthentication Resource.
// implemented by the user
type RequestAuthenticationReconciler interface {
	ReconcileRequestAuthentication(obj *security_istio_io_v1.RequestAuthentication) (reconcile.Result, error)
}

// Reconcile deletion events for the RequestAuthentication Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type RequestAuthenticationDeletionReconciler interface {
	ReconcileRequestAuthenticationDeletion(req reconcile.Request) error
}

type RequestAuthenticationReconcilerFuncs struct {
	OnReconcileRequestAuthentication         func(obj *security_istio_io_v1.RequestAuthentication) (reconcile.Result, error)
	OnReconcileRequestAuthenticationDeletion func(req reconcile.Request) error
}

func (f *RequestAuthenticationReconcilerFuncs) ReconcileRequestAuthentication(obj *security_istio_io_v1.RequestAuthentication) (reconcile.Result, error) {
	if f.OnReconcileRequestAuthentication == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileRequestAuthentication(obj)
}

func (f *RequestAuthenticationReconcilerFuncs) ReconcileRequestAuthenticationDeletion(req reconcile.Request) error {
	if f.OnReconcileRequestAuthenticationDeletion == nil {
		return nil
	}
	return f.OnReconcileRequestAuthenticationDeletion(req)
}

// Reconcile and finalize the RequestAuthentication Resource
// implemented by the user
type RequestAuthenticationFinalizer interface {
	RequestAuthenticationReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	RequestAuthenticationFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeRequestAuthentication(obj *security_istio_io_v1.RequestAuthentication) error
}

type RequestAuthenticationReconcileLoop interface {
	RunRequestAuthenticationReconciler(ctx context.Context, rec RequestAuthenticationReconciler, predicates ...predicate.Predicate) error
}

type requestAuthenticationReconcileLoop struct {
	loop reconcile.Loop
}

func NewRequestAuthenticationReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) RequestAuthenticationReconcileLoop {
	return &requestAuthenticationReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &security_istio_io_v1.RequestAuthentication{}, options),
	}
}

func (c *requestAuthenticationReconcileLoop) RunRequestAuthenticationReconciler(ctx context.Context, reconciler RequestAuthenticationReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericRequestAuthenticationReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(RequestAuthenticationFinalizer); ok {
		reconcilerWrapper = genericRequestAuthenticationFinalizer{
			genericRequestAuthenticationReconciler: genericReconciler,
			finalizingReconciler:                   finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericRequestAuthenticationHandler implements a generic reconcile.Reconciler
type genericRequestAuthenticationReconciler struct {
	reconciler RequestAuthenticationReconciler
}

func (r genericRequestAuthenticationReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_istio_io_v1.RequestAuthentication)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: RequestAuthentication handler received event for %T", object)
	}
	return r.reconciler.ReconcileRequestAuthentication(obj)
}

func (r genericRequestAuthenticationReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(RequestAuthenticationDeletionReconciler); ok {
		return deletionReconciler.ReconcileRequestAuthenticationDeletion(request)
	}
	return nil
}

// genericRequestAuthenticationFinalizer implements a generic reconcile.FinalizingReconciler
type genericRequestAuthenticationFinalizer struct {
	genericRequestAuthenticationReconciler
	finalizingReconciler RequestAuthenticationFinalizer
}

func (r genericRequestAuthenticationFinalizer) FinalizerName() string {
	return r.finalizingReconciler.RequestAuthenticationFinalizerName()
}

func (r genericRequestAuthenticationFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_istio_io_v1.RequestAuthentication)
	if !ok {
		return errors.Errorf("internal error: RequestAuthentication handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeRequestAuthentication(obj)
}
