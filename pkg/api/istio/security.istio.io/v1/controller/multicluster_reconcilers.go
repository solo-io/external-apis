// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	security_istio_io_v1 "istio.io/client-go/pkg/apis/security/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AuthorizationPolicy Resource across clusters.
// implemented by the user
type MulticlusterAuthorizationPolicyReconciler interface {
	ReconcileAuthorizationPolicy(clusterName string, obj *security_istio_io_v1.AuthorizationPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AuthorizationPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterAuthorizationPolicyDeletionReconciler interface {
	ReconcileAuthorizationPolicyDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterAuthorizationPolicyReconcilerFuncs struct {
	OnReconcileAuthorizationPolicy         func(clusterName string, obj *security_istio_io_v1.AuthorizationPolicy) (reconcile.Result, error)
	OnReconcileAuthorizationPolicyDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterAuthorizationPolicyReconcilerFuncs) ReconcileAuthorizationPolicy(clusterName string, obj *security_istio_io_v1.AuthorizationPolicy) (reconcile.Result, error) {
	if f.OnReconcileAuthorizationPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAuthorizationPolicy(clusterName, obj)
}

func (f *MulticlusterAuthorizationPolicyReconcilerFuncs) ReconcileAuthorizationPolicyDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileAuthorizationPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileAuthorizationPolicyDeletion(clusterName, req)
}

type MulticlusterAuthorizationPolicyReconcileLoop interface {
	// AddMulticlusterAuthorizationPolicyReconciler adds a MulticlusterAuthorizationPolicyReconciler to the MulticlusterAuthorizationPolicyReconcileLoop.
	AddMulticlusterAuthorizationPolicyReconciler(ctx context.Context, rec MulticlusterAuthorizationPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterAuthorizationPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterAuthorizationPolicyReconcileLoop) AddMulticlusterAuthorizationPolicyReconciler(ctx context.Context, rec MulticlusterAuthorizationPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericAuthorizationPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterAuthorizationPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterAuthorizationPolicyReconcileLoop {
	return &multiclusterAuthorizationPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_istio_io_v1.AuthorizationPolicy{}, options)}
}

type genericAuthorizationPolicyMulticlusterReconciler struct {
	reconciler MulticlusterAuthorizationPolicyReconciler
}

func (g genericAuthorizationPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterAuthorizationPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileAuthorizationPolicyDeletion(cluster, req)
	}
	return nil
}

func (g genericAuthorizationPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_istio_io_v1.AuthorizationPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileAuthorizationPolicy(cluster, obj)
}

// Reconcile Upsert events for the PeerAuthentication Resource across clusters.
// implemented by the user
type MulticlusterPeerAuthenticationReconciler interface {
	ReconcilePeerAuthentication(clusterName string, obj *security_istio_io_v1.PeerAuthentication) (reconcile.Result, error)
}

// Reconcile deletion events for the PeerAuthentication Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterPeerAuthenticationDeletionReconciler interface {
	ReconcilePeerAuthenticationDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterPeerAuthenticationReconcilerFuncs struct {
	OnReconcilePeerAuthentication         func(clusterName string, obj *security_istio_io_v1.PeerAuthentication) (reconcile.Result, error)
	OnReconcilePeerAuthenticationDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterPeerAuthenticationReconcilerFuncs) ReconcilePeerAuthentication(clusterName string, obj *security_istio_io_v1.PeerAuthentication) (reconcile.Result, error) {
	if f.OnReconcilePeerAuthentication == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcilePeerAuthentication(clusterName, obj)
}

func (f *MulticlusterPeerAuthenticationReconcilerFuncs) ReconcilePeerAuthenticationDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcilePeerAuthenticationDeletion == nil {
		return nil
	}
	return f.OnReconcilePeerAuthenticationDeletion(clusterName, req)
}

type MulticlusterPeerAuthenticationReconcileLoop interface {
	// AddMulticlusterPeerAuthenticationReconciler adds a MulticlusterPeerAuthenticationReconciler to the MulticlusterPeerAuthenticationReconcileLoop.
	AddMulticlusterPeerAuthenticationReconciler(ctx context.Context, rec MulticlusterPeerAuthenticationReconciler, predicates ...predicate.Predicate)
}

type multiclusterPeerAuthenticationReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterPeerAuthenticationReconcileLoop) AddMulticlusterPeerAuthenticationReconciler(ctx context.Context, rec MulticlusterPeerAuthenticationReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericPeerAuthenticationMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterPeerAuthenticationReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterPeerAuthenticationReconcileLoop {
	return &multiclusterPeerAuthenticationReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_istio_io_v1.PeerAuthentication{}, options)}
}

type genericPeerAuthenticationMulticlusterReconciler struct {
	reconciler MulticlusterPeerAuthenticationReconciler
}

func (g genericPeerAuthenticationMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterPeerAuthenticationDeletionReconciler); ok {
		return deletionReconciler.ReconcilePeerAuthenticationDeletion(cluster, req)
	}
	return nil
}

func (g genericPeerAuthenticationMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_istio_io_v1.PeerAuthentication)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: PeerAuthentication handler received event for %T", object)
	}
	return g.reconciler.ReconcilePeerAuthentication(cluster, obj)
}

// Reconcile Upsert events for the RequestAuthentication Resource across clusters.
// implemented by the user
type MulticlusterRequestAuthenticationReconciler interface {
	ReconcileRequestAuthentication(clusterName string, obj *security_istio_io_v1.RequestAuthentication) (reconcile.Result, error)
}

// Reconcile deletion events for the RequestAuthentication Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterRequestAuthenticationDeletionReconciler interface {
	ReconcileRequestAuthenticationDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterRequestAuthenticationReconcilerFuncs struct {
	OnReconcileRequestAuthentication         func(clusterName string, obj *security_istio_io_v1.RequestAuthentication) (reconcile.Result, error)
	OnReconcileRequestAuthenticationDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterRequestAuthenticationReconcilerFuncs) ReconcileRequestAuthentication(clusterName string, obj *security_istio_io_v1.RequestAuthentication) (reconcile.Result, error) {
	if f.OnReconcileRequestAuthentication == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileRequestAuthentication(clusterName, obj)
}

func (f *MulticlusterRequestAuthenticationReconcilerFuncs) ReconcileRequestAuthenticationDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileRequestAuthenticationDeletion == nil {
		return nil
	}
	return f.OnReconcileRequestAuthenticationDeletion(clusterName, req)
}

type MulticlusterRequestAuthenticationReconcileLoop interface {
	// AddMulticlusterRequestAuthenticationReconciler adds a MulticlusterRequestAuthenticationReconciler to the MulticlusterRequestAuthenticationReconcileLoop.
	AddMulticlusterRequestAuthenticationReconciler(ctx context.Context, rec MulticlusterRequestAuthenticationReconciler, predicates ...predicate.Predicate)
}

type multiclusterRequestAuthenticationReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterRequestAuthenticationReconcileLoop) AddMulticlusterRequestAuthenticationReconciler(ctx context.Context, rec MulticlusterRequestAuthenticationReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericRequestAuthenticationMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterRequestAuthenticationReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterRequestAuthenticationReconcileLoop {
	return &multiclusterRequestAuthenticationReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &security_istio_io_v1.RequestAuthentication{}, options)}
}

type genericRequestAuthenticationMulticlusterReconciler struct {
	reconciler MulticlusterRequestAuthenticationReconciler
}

func (g genericRequestAuthenticationMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterRequestAuthenticationDeletionReconciler); ok {
		return deletionReconciler.ReconcileRequestAuthenticationDeletion(cluster, req)
	}
	return nil
}

func (g genericRequestAuthenticationMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_istio_io_v1.RequestAuthentication)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: RequestAuthentication handler received event for %T", object)
	}
	return g.reconciler.ReconcileRequestAuthentication(cluster, obj)
}
