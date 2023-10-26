// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	gateway_networking_k8s_io_v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Gateway Resource across clusters.
// implemented by the user
type MulticlusterGatewayReconciler interface {
	ReconcileGateway(clusterName string, obj *gateway_networking_k8s_io_v1beta1.Gateway) (reconcile.Result, error)
}

// Reconcile deletion events for the Gateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGatewayDeletionReconciler interface {
	ReconcileGatewayDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGatewayReconcilerFuncs struct {
	OnReconcileGateway         func(clusterName string, obj *gateway_networking_k8s_io_v1beta1.Gateway) (reconcile.Result, error)
	OnReconcileGatewayDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGatewayReconcilerFuncs) ReconcileGateway(clusterName string, obj *gateway_networking_k8s_io_v1beta1.Gateway) (reconcile.Result, error) {
	if f.OnReconcileGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGateway(clusterName, obj)
}

func (f *MulticlusterGatewayReconcilerFuncs) ReconcileGatewayDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayDeletion(clusterName, req)
}

type MulticlusterGatewayReconcileLoop interface {
	// AddMulticlusterGatewayReconciler adds a MulticlusterGatewayReconciler to the MulticlusterGatewayReconcileLoop.
	AddMulticlusterGatewayReconciler(ctx context.Context, rec MulticlusterGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGatewayReconcileLoop) AddMulticlusterGatewayReconciler(ctx context.Context, rec MulticlusterGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGatewayReconcileLoop {
	return &multiclusterGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gateway_networking_k8s_io_v1beta1.Gateway{}, options)}
}

type genericGatewayMulticlusterReconciler struct {
	reconciler MulticlusterGatewayReconciler
}

func (g genericGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayDeletion(cluster, req)
	}
	return nil
}

func (g genericGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_networking_k8s_io_v1beta1.Gateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileGateway(cluster, obj)
}

// Reconcile Upsert events for the GatewayClass Resource across clusters.
// implemented by the user
type MulticlusterGatewayClassReconciler interface {
	ReconcileGatewayClass(clusterName string, obj *gateway_networking_k8s_io_v1beta1.GatewayClass) (reconcile.Result, error)
}

// Reconcile deletion events for the GatewayClass Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGatewayClassDeletionReconciler interface {
	ReconcileGatewayClassDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterGatewayClassReconcilerFuncs struct {
	OnReconcileGatewayClass         func(clusterName string, obj *gateway_networking_k8s_io_v1beta1.GatewayClass) (reconcile.Result, error)
	OnReconcileGatewayClassDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterGatewayClassReconcilerFuncs) ReconcileGatewayClass(clusterName string, obj *gateway_networking_k8s_io_v1beta1.GatewayClass) (reconcile.Result, error) {
	if f.OnReconcileGatewayClass == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGatewayClass(clusterName, obj)
}

func (f *MulticlusterGatewayClassReconcilerFuncs) ReconcileGatewayClassDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileGatewayClassDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayClassDeletion(clusterName, req)
}

type MulticlusterGatewayClassReconcileLoop interface {
	// AddMulticlusterGatewayClassReconciler adds a MulticlusterGatewayClassReconciler to the MulticlusterGatewayClassReconcileLoop.
	AddMulticlusterGatewayClassReconciler(ctx context.Context, rec MulticlusterGatewayClassReconciler, predicates ...predicate.Predicate)
}

type multiclusterGatewayClassReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGatewayClassReconcileLoop) AddMulticlusterGatewayClassReconciler(ctx context.Context, rec MulticlusterGatewayClassReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGatewayClassMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGatewayClassReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterGatewayClassReconcileLoop {
	return &multiclusterGatewayClassReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gateway_networking_k8s_io_v1beta1.GatewayClass{}, options)}
}

type genericGatewayClassMulticlusterReconciler struct {
	reconciler MulticlusterGatewayClassReconciler
}

func (g genericGatewayClassMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGatewayClassDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayClassDeletion(cluster, req)
	}
	return nil
}

func (g genericGatewayClassMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_networking_k8s_io_v1beta1.GatewayClass)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GatewayClass handler received event for %T", object)
	}
	return g.reconciler.ReconcileGatewayClass(cluster, obj)
}

// Reconcile Upsert events for the HTTPRoute Resource across clusters.
// implemented by the user
type MulticlusterHTTPRouteReconciler interface {
	ReconcileHTTPRoute(clusterName string, obj *gateway_networking_k8s_io_v1beta1.HTTPRoute) (reconcile.Result, error)
}

// Reconcile deletion events for the HTTPRoute Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterHTTPRouteDeletionReconciler interface {
	ReconcileHTTPRouteDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterHTTPRouteReconcilerFuncs struct {
	OnReconcileHTTPRoute         func(clusterName string, obj *gateway_networking_k8s_io_v1beta1.HTTPRoute) (reconcile.Result, error)
	OnReconcileHTTPRouteDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterHTTPRouteReconcilerFuncs) ReconcileHTTPRoute(clusterName string, obj *gateway_networking_k8s_io_v1beta1.HTTPRoute) (reconcile.Result, error) {
	if f.OnReconcileHTTPRoute == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileHTTPRoute(clusterName, obj)
}

func (f *MulticlusterHTTPRouteReconcilerFuncs) ReconcileHTTPRouteDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileHTTPRouteDeletion == nil {
		return nil
	}
	return f.OnReconcileHTTPRouteDeletion(clusterName, req)
}

type MulticlusterHTTPRouteReconcileLoop interface {
	// AddMulticlusterHTTPRouteReconciler adds a MulticlusterHTTPRouteReconciler to the MulticlusterHTTPRouteReconcileLoop.
	AddMulticlusterHTTPRouteReconciler(ctx context.Context, rec MulticlusterHTTPRouteReconciler, predicates ...predicate.Predicate)
}

type multiclusterHTTPRouteReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterHTTPRouteReconcileLoop) AddMulticlusterHTTPRouteReconciler(ctx context.Context, rec MulticlusterHTTPRouteReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericHTTPRouteMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterHTTPRouteReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterHTTPRouteReconcileLoop {
	return &multiclusterHTTPRouteReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gateway_networking_k8s_io_v1beta1.HTTPRoute{}, options)}
}

type genericHTTPRouteMulticlusterReconciler struct {
	reconciler MulticlusterHTTPRouteReconciler
}

func (g genericHTTPRouteMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterHTTPRouteDeletionReconciler); ok {
		return deletionReconciler.ReconcileHTTPRouteDeletion(cluster, req)
	}
	return nil
}

func (g genericHTTPRouteMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_networking_k8s_io_v1beta1.HTTPRoute)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: HTTPRoute handler received event for %T", object)
	}
	return g.reconciler.ReconcileHTTPRoute(cluster, obj)
}

// Reconcile Upsert events for the ReferenceGrant Resource across clusters.
// implemented by the user
type MulticlusterReferenceGrantReconciler interface {
	ReconcileReferenceGrant(clusterName string, obj *gateway_networking_k8s_io_v1beta1.ReferenceGrant) (reconcile.Result, error)
}

// Reconcile deletion events for the ReferenceGrant Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterReferenceGrantDeletionReconciler interface {
	ReconcileReferenceGrantDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterReferenceGrantReconcilerFuncs struct {
	OnReconcileReferenceGrant         func(clusterName string, obj *gateway_networking_k8s_io_v1beta1.ReferenceGrant) (reconcile.Result, error)
	OnReconcileReferenceGrantDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterReferenceGrantReconcilerFuncs) ReconcileReferenceGrant(clusterName string, obj *gateway_networking_k8s_io_v1beta1.ReferenceGrant) (reconcile.Result, error) {
	if f.OnReconcileReferenceGrant == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileReferenceGrant(clusterName, obj)
}

func (f *MulticlusterReferenceGrantReconcilerFuncs) ReconcileReferenceGrantDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileReferenceGrantDeletion == nil {
		return nil
	}
	return f.OnReconcileReferenceGrantDeletion(clusterName, req)
}

type MulticlusterReferenceGrantReconcileLoop interface {
	// AddMulticlusterReferenceGrantReconciler adds a MulticlusterReferenceGrantReconciler to the MulticlusterReferenceGrantReconcileLoop.
	AddMulticlusterReferenceGrantReconciler(ctx context.Context, rec MulticlusterReferenceGrantReconciler, predicates ...predicate.Predicate)
}

type multiclusterReferenceGrantReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterReferenceGrantReconcileLoop) AddMulticlusterReferenceGrantReconciler(ctx context.Context, rec MulticlusterReferenceGrantReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericReferenceGrantMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterReferenceGrantReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterReferenceGrantReconcileLoop {
	return &multiclusterReferenceGrantReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &gateway_networking_k8s_io_v1beta1.ReferenceGrant{}, options)}
}

type genericReferenceGrantMulticlusterReconciler struct {
	reconciler MulticlusterReferenceGrantReconciler
}

func (g genericReferenceGrantMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterReferenceGrantDeletionReconciler); ok {
		return deletionReconciler.ReconcileReferenceGrantDeletion(cluster, req)
	}
	return nil
}

func (g genericReferenceGrantMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gateway_networking_k8s_io_v1beta1.ReferenceGrant)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ReferenceGrant handler received event for %T", object)
	}
	return g.reconciler.ReconcileReferenceGrant(cluster, obj)
}
