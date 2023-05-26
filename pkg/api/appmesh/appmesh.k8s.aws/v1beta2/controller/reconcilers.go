// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Mesh Resource.
// implemented by the user
type MeshReconciler interface {
	ReconcileMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) (reconcile.Result, error)
}

// Reconcile deletion events for the Mesh Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MeshDeletionReconciler interface {
	ReconcileMeshDeletion(req reconcile.Request) error
}

type MeshReconcilerFuncs struct {
	OnReconcileMesh         func(obj *appmesh_k8s_aws_v1beta2.Mesh) (reconcile.Result, error)
	OnReconcileMeshDeletion func(req reconcile.Request) error
}

func (f *MeshReconcilerFuncs) ReconcileMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) (reconcile.Result, error) {
	if f.OnReconcileMesh == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMesh(obj)
}

func (f *MeshReconcilerFuncs) ReconcileMeshDeletion(req reconcile.Request) error {
	if f.OnReconcileMeshDeletion == nil {
		return nil
	}
	return f.OnReconcileMeshDeletion(req)
}

// Reconcile and finalize the Mesh Resource
// implemented by the user
type MeshFinalizer interface {
	MeshReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	MeshFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) error
}

type MeshReconcileLoop interface {
	RunMeshReconciler(ctx context.Context, rec MeshReconciler, predicates ...predicate.Predicate) error
}

type meshReconcileLoop struct {
	loop reconcile.Loop
}

func NewMeshReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) MeshReconcileLoop {
	return &meshReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &appmesh_k8s_aws_v1beta2.Mesh{}, options),
	}
}

func (c *meshReconcileLoop) RunMeshReconciler(ctx context.Context, reconciler MeshReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericMeshReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(MeshFinalizer); ok {
		reconcilerWrapper = genericMeshFinalizer{
			genericMeshReconciler: genericReconciler,
			finalizingReconciler:  finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericMeshHandler implements a generic reconcile.Reconciler
type genericMeshReconciler struct {
	reconciler MeshReconciler
}

func (r genericMeshReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.Mesh)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return r.reconciler.ReconcileMesh(obj)
}

func (r genericMeshReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(MeshDeletionReconciler); ok {
		return deletionReconciler.ReconcileMeshDeletion(request)
	}
	return nil
}

// genericMeshFinalizer implements a generic reconcile.FinalizingReconciler
type genericMeshFinalizer struct {
	genericMeshReconciler
	finalizingReconciler MeshFinalizer
}

func (r genericMeshFinalizer) FinalizerName() string {
	return r.finalizingReconciler.MeshFinalizerName()
}

func (r genericMeshFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.Mesh)
	if !ok {
		return errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeMesh(obj)
}

// Reconcile Upsert events for the VirtualService Resource.
// implemented by the user
type VirtualServiceReconciler interface {
	ReconcileVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualService Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualServiceDeletionReconciler interface {
	ReconcileVirtualServiceDeletion(req reconcile.Request) error
}

type VirtualServiceReconcilerFuncs struct {
	OnReconcileVirtualService         func(obj *appmesh_k8s_aws_v1beta2.VirtualService) (reconcile.Result, error)
	OnReconcileVirtualServiceDeletion func(req reconcile.Request) error
}

func (f *VirtualServiceReconcilerFuncs) ReconcileVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) (reconcile.Result, error) {
	if f.OnReconcileVirtualService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualService(obj)
}

func (f *VirtualServiceReconcilerFuncs) ReconcileVirtualServiceDeletion(req reconcile.Request) error {
	if f.OnReconcileVirtualServiceDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualServiceDeletion(req)
}

// Reconcile and finalize the VirtualService Resource
// implemented by the user
type VirtualServiceFinalizer interface {
	VirtualServiceReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	VirtualServiceFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) error
}

type VirtualServiceReconcileLoop interface {
	RunVirtualServiceReconciler(ctx context.Context, rec VirtualServiceReconciler, predicates ...predicate.Predicate) error
}

type virtualServiceReconcileLoop struct {
	loop reconcile.Loop
}

func NewVirtualServiceReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) VirtualServiceReconcileLoop {
	return &virtualServiceReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &appmesh_k8s_aws_v1beta2.VirtualService{}, options),
	}
}

func (c *virtualServiceReconcileLoop) RunVirtualServiceReconciler(ctx context.Context, reconciler VirtualServiceReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericVirtualServiceReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(VirtualServiceFinalizer); ok {
		reconcilerWrapper = genericVirtualServiceFinalizer{
			genericVirtualServiceReconciler: genericReconciler,
			finalizingReconciler:            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericVirtualServiceHandler implements a generic reconcile.Reconciler
type genericVirtualServiceReconciler struct {
	reconciler VirtualServiceReconciler
}

func (r genericVirtualServiceReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return r.reconciler.ReconcileVirtualService(obj)
}

func (r genericVirtualServiceReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(VirtualServiceDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualServiceDeletion(request)
	}
	return nil
}

// genericVirtualServiceFinalizer implements a generic reconcile.FinalizingReconciler
type genericVirtualServiceFinalizer struct {
	genericVirtualServiceReconciler
	finalizingReconciler VirtualServiceFinalizer
}

func (r genericVirtualServiceFinalizer) FinalizerName() string {
	return r.finalizingReconciler.VirtualServiceFinalizerName()
}

func (r genericVirtualServiceFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualService(obj)
}

// Reconcile Upsert events for the VirtualNode Resource.
// implemented by the user
type VirtualNodeReconciler interface {
	ReconcileVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualNode Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualNodeDeletionReconciler interface {
	ReconcileVirtualNodeDeletion(req reconcile.Request) error
}

type VirtualNodeReconcilerFuncs struct {
	OnReconcileVirtualNode         func(obj *appmesh_k8s_aws_v1beta2.VirtualNode) (reconcile.Result, error)
	OnReconcileVirtualNodeDeletion func(req reconcile.Request) error
}

func (f *VirtualNodeReconcilerFuncs) ReconcileVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) (reconcile.Result, error) {
	if f.OnReconcileVirtualNode == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualNode(obj)
}

func (f *VirtualNodeReconcilerFuncs) ReconcileVirtualNodeDeletion(req reconcile.Request) error {
	if f.OnReconcileVirtualNodeDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualNodeDeletion(req)
}

// Reconcile and finalize the VirtualNode Resource
// implemented by the user
type VirtualNodeFinalizer interface {
	VirtualNodeReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	VirtualNodeFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error
}

type VirtualNodeReconcileLoop interface {
	RunVirtualNodeReconciler(ctx context.Context, rec VirtualNodeReconciler, predicates ...predicate.Predicate) error
}

type virtualNodeReconcileLoop struct {
	loop reconcile.Loop
}

func NewVirtualNodeReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) VirtualNodeReconcileLoop {
	return &virtualNodeReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &appmesh_k8s_aws_v1beta2.VirtualNode{}, options),
	}
}

func (c *virtualNodeReconcileLoop) RunVirtualNodeReconciler(ctx context.Context, reconciler VirtualNodeReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericVirtualNodeReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(VirtualNodeFinalizer); ok {
		reconcilerWrapper = genericVirtualNodeFinalizer{
			genericVirtualNodeReconciler: genericReconciler,
			finalizingReconciler:         finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericVirtualNodeHandler implements a generic reconcile.Reconciler
type genericVirtualNodeReconciler struct {
	reconciler VirtualNodeReconciler
}

func (r genericVirtualNodeReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualNode handler received event for %T", object)
	}
	return r.reconciler.ReconcileVirtualNode(obj)
}

func (r genericVirtualNodeReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(VirtualNodeDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualNodeDeletion(request)
	}
	return nil
}

// genericVirtualNodeFinalizer implements a generic reconcile.FinalizingReconciler
type genericVirtualNodeFinalizer struct {
	genericVirtualNodeReconciler
	finalizingReconciler VirtualNodeFinalizer
}

func (r genericVirtualNodeFinalizer) FinalizerName() string {
	return r.finalizingReconciler.VirtualNodeFinalizerName()
}

func (r genericVirtualNodeFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	if !ok {
		return errors.Errorf("internal error: VirtualNode handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualNode(obj)
}

// Reconcile Upsert events for the VirtualRouter Resource.
// implemented by the user
type VirtualRouterReconciler interface {
	ReconcileVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualRouter Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualRouterDeletionReconciler interface {
	ReconcileVirtualRouterDeletion(req reconcile.Request) error
}

type VirtualRouterReconcilerFuncs struct {
	OnReconcileVirtualRouter         func(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) (reconcile.Result, error)
	OnReconcileVirtualRouterDeletion func(req reconcile.Request) error
}

func (f *VirtualRouterReconcilerFuncs) ReconcileVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) (reconcile.Result, error) {
	if f.OnReconcileVirtualRouter == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualRouter(obj)
}

func (f *VirtualRouterReconcilerFuncs) ReconcileVirtualRouterDeletion(req reconcile.Request) error {
	if f.OnReconcileVirtualRouterDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualRouterDeletion(req)
}

// Reconcile and finalize the VirtualRouter Resource
// implemented by the user
type VirtualRouterFinalizer interface {
	VirtualRouterReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	VirtualRouterFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error
}

type VirtualRouterReconcileLoop interface {
	RunVirtualRouterReconciler(ctx context.Context, rec VirtualRouterReconciler, predicates ...predicate.Predicate) error
}

type virtualRouterReconcileLoop struct {
	loop reconcile.Loop
}

func NewVirtualRouterReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) VirtualRouterReconcileLoop {
	return &virtualRouterReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &appmesh_k8s_aws_v1beta2.VirtualRouter{}, options),
	}
}

func (c *virtualRouterReconcileLoop) RunVirtualRouterReconciler(ctx context.Context, reconciler VirtualRouterReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericVirtualRouterReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(VirtualRouterFinalizer); ok {
		reconcilerWrapper = genericVirtualRouterFinalizer{
			genericVirtualRouterReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericVirtualRouterHandler implements a generic reconcile.Reconciler
type genericVirtualRouterReconciler struct {
	reconciler VirtualRouterReconciler
}

func (r genericVirtualRouterReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualRouter handler received event for %T", object)
	}
	return r.reconciler.ReconcileVirtualRouter(obj)
}

func (r genericVirtualRouterReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(VirtualRouterDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualRouterDeletion(request)
	}
	return nil
}

// genericVirtualRouterFinalizer implements a generic reconcile.FinalizingReconciler
type genericVirtualRouterFinalizer struct {
	genericVirtualRouterReconciler
	finalizingReconciler VirtualRouterFinalizer
}

func (r genericVirtualRouterFinalizer) FinalizerName() string {
	return r.finalizingReconciler.VirtualRouterFinalizerName()
}

func (r genericVirtualRouterFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	if !ok {
		return errors.Errorf("internal error: VirtualRouter handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualRouter(obj)
}

// Reconcile Upsert events for the VirtualGateway Resource.
// implemented by the user
type VirtualGatewayReconciler interface {
	ReconcileVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualGateway Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualGatewayDeletionReconciler interface {
	ReconcileVirtualGatewayDeletion(req reconcile.Request) error
}

type VirtualGatewayReconcilerFuncs struct {
	OnReconcileVirtualGateway         func(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) (reconcile.Result, error)
	OnReconcileVirtualGatewayDeletion func(req reconcile.Request) error
}

func (f *VirtualGatewayReconcilerFuncs) ReconcileVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) (reconcile.Result, error) {
	if f.OnReconcileVirtualGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualGateway(obj)
}

func (f *VirtualGatewayReconcilerFuncs) ReconcileVirtualGatewayDeletion(req reconcile.Request) error {
	if f.OnReconcileVirtualGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualGatewayDeletion(req)
}

// Reconcile and finalize the VirtualGateway Resource
// implemented by the user
type VirtualGatewayFinalizer interface {
	VirtualGatewayReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	VirtualGatewayFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error
}

type VirtualGatewayReconcileLoop interface {
	RunVirtualGatewayReconciler(ctx context.Context, rec VirtualGatewayReconciler, predicates ...predicate.Predicate) error
}

type virtualGatewayReconcileLoop struct {
	loop reconcile.Loop
}

func NewVirtualGatewayReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) VirtualGatewayReconcileLoop {
	return &virtualGatewayReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &appmesh_k8s_aws_v1beta2.VirtualGateway{}, options),
	}
}

func (c *virtualGatewayReconcileLoop) RunVirtualGatewayReconciler(ctx context.Context, reconciler VirtualGatewayReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericVirtualGatewayReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(VirtualGatewayFinalizer); ok {
		reconcilerWrapper = genericVirtualGatewayFinalizer{
			genericVirtualGatewayReconciler: genericReconciler,
			finalizingReconciler:            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericVirtualGatewayHandler implements a generic reconcile.Reconciler
type genericVirtualGatewayReconciler struct {
	reconciler VirtualGatewayReconciler
}

func (r genericVirtualGatewayReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return r.reconciler.ReconcileVirtualGateway(obj)
}

func (r genericVirtualGatewayReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(VirtualGatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualGatewayDeletion(request)
	}
	return nil
}

// genericVirtualGatewayFinalizer implements a generic reconcile.FinalizingReconciler
type genericVirtualGatewayFinalizer struct {
	genericVirtualGatewayReconciler
	finalizingReconciler VirtualGatewayFinalizer
}

func (r genericVirtualGatewayFinalizer) FinalizerName() string {
	return r.finalizingReconciler.VirtualGatewayFinalizerName()
}

func (r genericVirtualGatewayFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualGateway(obj)
}

// Reconcile Upsert events for the GatewayRoute Resource.
// implemented by the user
type GatewayRouteReconciler interface {
	ReconcileGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) (reconcile.Result, error)
}

// Reconcile deletion events for the GatewayRoute Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type GatewayRouteDeletionReconciler interface {
	ReconcileGatewayRouteDeletion(req reconcile.Request) error
}

type GatewayRouteReconcilerFuncs struct {
	OnReconcileGatewayRoute         func(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) (reconcile.Result, error)
	OnReconcileGatewayRouteDeletion func(req reconcile.Request) error
}

func (f *GatewayRouteReconcilerFuncs) ReconcileGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) (reconcile.Result, error) {
	if f.OnReconcileGatewayRoute == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGatewayRoute(obj)
}

func (f *GatewayRouteReconcilerFuncs) ReconcileGatewayRouteDeletion(req reconcile.Request) error {
	if f.OnReconcileGatewayRouteDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayRouteDeletion(req)
}

// Reconcile and finalize the GatewayRoute Resource
// implemented by the user
type GatewayRouteFinalizer interface {
	GatewayRouteReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	GatewayRouteFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error
}

type GatewayRouteReconcileLoop interface {
	RunGatewayRouteReconciler(ctx context.Context, rec GatewayRouteReconciler, predicates ...predicate.Predicate) error
}

type gatewayRouteReconcileLoop struct {
	loop reconcile.Loop
}

func NewGatewayRouteReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) GatewayRouteReconcileLoop {
	return &gatewayRouteReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &appmesh_k8s_aws_v1beta2.GatewayRoute{}, options),
	}
}

func (c *gatewayRouteReconcileLoop) RunGatewayRouteReconciler(ctx context.Context, reconciler GatewayRouteReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericGatewayRouteReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(GatewayRouteFinalizer); ok {
		reconcilerWrapper = genericGatewayRouteFinalizer{
			genericGatewayRouteReconciler: genericReconciler,
			finalizingReconciler:          finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericGatewayRouteHandler implements a generic reconcile.Reconciler
type genericGatewayRouteReconciler struct {
	reconciler GatewayRouteReconciler
}

func (r genericGatewayRouteReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: GatewayRoute handler received event for %T", object)
	}
	return r.reconciler.ReconcileGatewayRoute(obj)
}

func (r genericGatewayRouteReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(GatewayRouteDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayRouteDeletion(request)
	}
	return nil
}

// genericGatewayRouteFinalizer implements a generic reconcile.FinalizingReconciler
type genericGatewayRouteFinalizer struct {
	genericGatewayRouteReconciler
	finalizingReconciler GatewayRouteFinalizer
}

func (r genericGatewayRouteFinalizer) FinalizerName() string {
	return r.finalizingReconciler.GatewayRouteFinalizerName()
}

func (r genericGatewayRouteFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	if !ok {
		return errors.Errorf("internal error: GatewayRoute handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeGatewayRoute(obj)
}