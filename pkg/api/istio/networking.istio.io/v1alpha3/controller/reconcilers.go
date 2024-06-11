// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_istio_io_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the DestinationRule Resource.
// implemented by the user
type DestinationRuleReconciler interface {
	ReconcileDestinationRule(obj *networking_istio_io_v1alpha3.DestinationRule) (reconcile.Result, error)
}

// Reconcile deletion events for the DestinationRule Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type DestinationRuleDeletionReconciler interface {
	ReconcileDestinationRuleDeletion(req reconcile.Request) error
}

type DestinationRuleReconcilerFuncs struct {
	OnReconcileDestinationRule         func(obj *networking_istio_io_v1alpha3.DestinationRule) (reconcile.Result, error)
	OnReconcileDestinationRuleDeletion func(req reconcile.Request) error
}

func (f *DestinationRuleReconcilerFuncs) ReconcileDestinationRule(obj *networking_istio_io_v1alpha3.DestinationRule) (reconcile.Result, error) {
	if f.OnReconcileDestinationRule == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileDestinationRule(obj)
}

func (f *DestinationRuleReconcilerFuncs) ReconcileDestinationRuleDeletion(req reconcile.Request) error {
	if f.OnReconcileDestinationRuleDeletion == nil {
		return nil
	}
	return f.OnReconcileDestinationRuleDeletion(req)
}

// Reconcile and finalize the DestinationRule Resource
// implemented by the user
type DestinationRuleFinalizer interface {
	DestinationRuleReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	DestinationRuleFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeDestinationRule(obj *networking_istio_io_v1alpha3.DestinationRule) error
}

type DestinationRuleReconcileLoop interface {
	RunDestinationRuleReconciler(ctx context.Context, rec DestinationRuleReconciler, predicates ...predicate.Predicate) error
}

type destinationRuleReconcileLoop struct {
	loop reconcile.Loop
}

func NewDestinationRuleReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) DestinationRuleReconcileLoop {
	return &destinationRuleReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_istio_io_v1alpha3.DestinationRule{}, options),
	}
}

func (c *destinationRuleReconcileLoop) RunDestinationRuleReconciler(ctx context.Context, reconciler DestinationRuleReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericDestinationRuleReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(DestinationRuleFinalizer); ok {
		reconcilerWrapper = genericDestinationRuleFinalizer{
			genericDestinationRuleReconciler: genericReconciler,
			finalizingReconciler:             finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericDestinationRuleHandler implements a generic reconcile.Reconciler
type genericDestinationRuleReconciler struct {
	reconciler DestinationRuleReconciler
}

func (r genericDestinationRuleReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.DestinationRule)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: DestinationRule handler received event for %T", object)
	}
	return r.reconciler.ReconcileDestinationRule(obj)
}

func (r genericDestinationRuleReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(DestinationRuleDeletionReconciler); ok {
		return deletionReconciler.ReconcileDestinationRuleDeletion(request)
	}
	return nil
}

// genericDestinationRuleFinalizer implements a generic reconcile.FinalizingReconciler
type genericDestinationRuleFinalizer struct {
	genericDestinationRuleReconciler
	finalizingReconciler DestinationRuleFinalizer
}

func (r genericDestinationRuleFinalizer) FinalizerName() string {
	return r.finalizingReconciler.DestinationRuleFinalizerName()
}

func (r genericDestinationRuleFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_istio_io_v1alpha3.DestinationRule)
	if !ok {
		return errors.Errorf("internal error: DestinationRule handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeDestinationRule(obj)
}

// Reconcile Upsert events for the Gateway Resource.
// implemented by the user
type GatewayReconciler interface {
	ReconcileGateway(obj *networking_istio_io_v1alpha3.Gateway) (reconcile.Result, error)
}

// Reconcile deletion events for the Gateway Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type GatewayDeletionReconciler interface {
	ReconcileGatewayDeletion(req reconcile.Request) error
}

type GatewayReconcilerFuncs struct {
	OnReconcileGateway         func(obj *networking_istio_io_v1alpha3.Gateway) (reconcile.Result, error)
	OnReconcileGatewayDeletion func(req reconcile.Request) error
}

func (f *GatewayReconcilerFuncs) ReconcileGateway(obj *networking_istio_io_v1alpha3.Gateway) (reconcile.Result, error) {
	if f.OnReconcileGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGateway(obj)
}

func (f *GatewayReconcilerFuncs) ReconcileGatewayDeletion(req reconcile.Request) error {
	if f.OnReconcileGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayDeletion(req)
}

// Reconcile and finalize the Gateway Resource
// implemented by the user
type GatewayFinalizer interface {
	GatewayReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	GatewayFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeGateway(obj *networking_istio_io_v1alpha3.Gateway) error
}

type GatewayReconcileLoop interface {
	RunGatewayReconciler(ctx context.Context, rec GatewayReconciler, predicates ...predicate.Predicate) error
}

type gatewayReconcileLoop struct {
	loop reconcile.Loop
}

func NewGatewayReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) GatewayReconcileLoop {
	return &gatewayReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_istio_io_v1alpha3.Gateway{}, options),
	}
}

func (c *gatewayReconcileLoop) RunGatewayReconciler(ctx context.Context, reconciler GatewayReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericGatewayReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(GatewayFinalizer); ok {
		reconcilerWrapper = genericGatewayFinalizer{
			genericGatewayReconciler: genericReconciler,
			finalizingReconciler:     finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericGatewayHandler implements a generic reconcile.Reconciler
type genericGatewayReconciler struct {
	reconciler GatewayReconciler
}

func (r genericGatewayReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.Gateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return r.reconciler.ReconcileGateway(obj)
}

func (r genericGatewayReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(GatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayDeletion(request)
	}
	return nil
}

// genericGatewayFinalizer implements a generic reconcile.FinalizingReconciler
type genericGatewayFinalizer struct {
	genericGatewayReconciler
	finalizingReconciler GatewayFinalizer
}

func (r genericGatewayFinalizer) FinalizerName() string {
	return r.finalizingReconciler.GatewayFinalizerName()
}

func (r genericGatewayFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_istio_io_v1alpha3.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeGateway(obj)
}

// Reconcile Upsert events for the ServiceEntry Resource.
// implemented by the user
type ServiceEntryReconciler interface {
	ReconcileServiceEntry(obj *networking_istio_io_v1alpha3.ServiceEntry) (reconcile.Result, error)
}

// Reconcile deletion events for the ServiceEntry Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ServiceEntryDeletionReconciler interface {
	ReconcileServiceEntryDeletion(req reconcile.Request) error
}

type ServiceEntryReconcilerFuncs struct {
	OnReconcileServiceEntry         func(obj *networking_istio_io_v1alpha3.ServiceEntry) (reconcile.Result, error)
	OnReconcileServiceEntryDeletion func(req reconcile.Request) error
}

func (f *ServiceEntryReconcilerFuncs) ReconcileServiceEntry(obj *networking_istio_io_v1alpha3.ServiceEntry) (reconcile.Result, error) {
	if f.OnReconcileServiceEntry == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileServiceEntry(obj)
}

func (f *ServiceEntryReconcilerFuncs) ReconcileServiceEntryDeletion(req reconcile.Request) error {
	if f.OnReconcileServiceEntryDeletion == nil {
		return nil
	}
	return f.OnReconcileServiceEntryDeletion(req)
}

// Reconcile and finalize the ServiceEntry Resource
// implemented by the user
type ServiceEntryFinalizer interface {
	ServiceEntryReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ServiceEntryFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeServiceEntry(obj *networking_istio_io_v1alpha3.ServiceEntry) error
}

type ServiceEntryReconcileLoop interface {
	RunServiceEntryReconciler(ctx context.Context, rec ServiceEntryReconciler, predicates ...predicate.Predicate) error
}

type serviceEntryReconcileLoop struct {
	loop reconcile.Loop
}

func NewServiceEntryReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) ServiceEntryReconcileLoop {
	return &serviceEntryReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_istio_io_v1alpha3.ServiceEntry{}, options),
	}
}

func (c *serviceEntryReconcileLoop) RunServiceEntryReconciler(ctx context.Context, reconciler ServiceEntryReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericServiceEntryReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ServiceEntryFinalizer); ok {
		reconcilerWrapper = genericServiceEntryFinalizer{
			genericServiceEntryReconciler: genericReconciler,
			finalizingReconciler:          finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericServiceEntryHandler implements a generic reconcile.Reconciler
type genericServiceEntryReconciler struct {
	reconciler ServiceEntryReconciler
}

func (r genericServiceEntryReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.ServiceEntry)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ServiceEntry handler received event for %T", object)
	}
	return r.reconciler.ReconcileServiceEntry(obj)
}

func (r genericServiceEntryReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(ServiceEntryDeletionReconciler); ok {
		return deletionReconciler.ReconcileServiceEntryDeletion(request)
	}
	return nil
}

// genericServiceEntryFinalizer implements a generic reconcile.FinalizingReconciler
type genericServiceEntryFinalizer struct {
	genericServiceEntryReconciler
	finalizingReconciler ServiceEntryFinalizer
}

func (r genericServiceEntryFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ServiceEntryFinalizerName()
}

func (r genericServiceEntryFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_istio_io_v1alpha3.ServiceEntry)
	if !ok {
		return errors.Errorf("internal error: ServiceEntry handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeServiceEntry(obj)
}

// Reconcile Upsert events for the WorkloadEntry Resource.
// implemented by the user
type WorkloadEntryReconciler interface {
	ReconcileWorkloadEntry(obj *networking_istio_io_v1alpha3.WorkloadEntry) (reconcile.Result, error)
}

// Reconcile deletion events for the WorkloadEntry Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type WorkloadEntryDeletionReconciler interface {
	ReconcileWorkloadEntryDeletion(req reconcile.Request) error
}

type WorkloadEntryReconcilerFuncs struct {
	OnReconcileWorkloadEntry         func(obj *networking_istio_io_v1alpha3.WorkloadEntry) (reconcile.Result, error)
	OnReconcileWorkloadEntryDeletion func(req reconcile.Request) error
}

func (f *WorkloadEntryReconcilerFuncs) ReconcileWorkloadEntry(obj *networking_istio_io_v1alpha3.WorkloadEntry) (reconcile.Result, error) {
	if f.OnReconcileWorkloadEntry == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWorkloadEntry(obj)
}

func (f *WorkloadEntryReconcilerFuncs) ReconcileWorkloadEntryDeletion(req reconcile.Request) error {
	if f.OnReconcileWorkloadEntryDeletion == nil {
		return nil
	}
	return f.OnReconcileWorkloadEntryDeletion(req)
}

// Reconcile and finalize the WorkloadEntry Resource
// implemented by the user
type WorkloadEntryFinalizer interface {
	WorkloadEntryReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	WorkloadEntryFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeWorkloadEntry(obj *networking_istio_io_v1alpha3.WorkloadEntry) error
}

type WorkloadEntryReconcileLoop interface {
	RunWorkloadEntryReconciler(ctx context.Context, rec WorkloadEntryReconciler, predicates ...predicate.Predicate) error
}

type workloadEntryReconcileLoop struct {
	loop reconcile.Loop
}

func NewWorkloadEntryReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) WorkloadEntryReconcileLoop {
	return &workloadEntryReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_istio_io_v1alpha3.WorkloadEntry{}, options),
	}
}

func (c *workloadEntryReconcileLoop) RunWorkloadEntryReconciler(ctx context.Context, reconciler WorkloadEntryReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericWorkloadEntryReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(WorkloadEntryFinalizer); ok {
		reconcilerWrapper = genericWorkloadEntryFinalizer{
			genericWorkloadEntryReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericWorkloadEntryHandler implements a generic reconcile.Reconciler
type genericWorkloadEntryReconciler struct {
	reconciler WorkloadEntryReconciler
}

func (r genericWorkloadEntryReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.WorkloadEntry)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WorkloadEntry handler received event for %T", object)
	}
	return r.reconciler.ReconcileWorkloadEntry(obj)
}

func (r genericWorkloadEntryReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(WorkloadEntryDeletionReconciler); ok {
		return deletionReconciler.ReconcileWorkloadEntryDeletion(request)
	}
	return nil
}

// genericWorkloadEntryFinalizer implements a generic reconcile.FinalizingReconciler
type genericWorkloadEntryFinalizer struct {
	genericWorkloadEntryReconciler
	finalizingReconciler WorkloadEntryFinalizer
}

func (r genericWorkloadEntryFinalizer) FinalizerName() string {
	return r.finalizingReconciler.WorkloadEntryFinalizerName()
}

func (r genericWorkloadEntryFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_istio_io_v1alpha3.WorkloadEntry)
	if !ok {
		return errors.Errorf("internal error: WorkloadEntry handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeWorkloadEntry(obj)
}

// Reconcile Upsert events for the WorkloadGroup Resource.
// implemented by the user
type WorkloadGroupReconciler interface {
	ReconcileWorkloadGroup(obj *networking_istio_io_v1alpha3.WorkloadGroup) (reconcile.Result, error)
}

// Reconcile deletion events for the WorkloadGroup Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type WorkloadGroupDeletionReconciler interface {
	ReconcileWorkloadGroupDeletion(req reconcile.Request) error
}

type WorkloadGroupReconcilerFuncs struct {
	OnReconcileWorkloadGroup         func(obj *networking_istio_io_v1alpha3.WorkloadGroup) (reconcile.Result, error)
	OnReconcileWorkloadGroupDeletion func(req reconcile.Request) error
}

func (f *WorkloadGroupReconcilerFuncs) ReconcileWorkloadGroup(obj *networking_istio_io_v1alpha3.WorkloadGroup) (reconcile.Result, error) {
	if f.OnReconcileWorkloadGroup == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWorkloadGroup(obj)
}

func (f *WorkloadGroupReconcilerFuncs) ReconcileWorkloadGroupDeletion(req reconcile.Request) error {
	if f.OnReconcileWorkloadGroupDeletion == nil {
		return nil
	}
	return f.OnReconcileWorkloadGroupDeletion(req)
}

// Reconcile and finalize the WorkloadGroup Resource
// implemented by the user
type WorkloadGroupFinalizer interface {
	WorkloadGroupReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	WorkloadGroupFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeWorkloadGroup(obj *networking_istio_io_v1alpha3.WorkloadGroup) error
}

type WorkloadGroupReconcileLoop interface {
	RunWorkloadGroupReconciler(ctx context.Context, rec WorkloadGroupReconciler, predicates ...predicate.Predicate) error
}

type workloadGroupReconcileLoop struct {
	loop reconcile.Loop
}

func NewWorkloadGroupReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) WorkloadGroupReconcileLoop {
	return &workloadGroupReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_istio_io_v1alpha3.WorkloadGroup{}, options),
	}
}

func (c *workloadGroupReconcileLoop) RunWorkloadGroupReconciler(ctx context.Context, reconciler WorkloadGroupReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericWorkloadGroupReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(WorkloadGroupFinalizer); ok {
		reconcilerWrapper = genericWorkloadGroupFinalizer{
			genericWorkloadGroupReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericWorkloadGroupHandler implements a generic reconcile.Reconciler
type genericWorkloadGroupReconciler struct {
	reconciler WorkloadGroupReconciler
}

func (r genericWorkloadGroupReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.WorkloadGroup)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WorkloadGroup handler received event for %T", object)
	}
	return r.reconciler.ReconcileWorkloadGroup(obj)
}

func (r genericWorkloadGroupReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(WorkloadGroupDeletionReconciler); ok {
		return deletionReconciler.ReconcileWorkloadGroupDeletion(request)
	}
	return nil
}

// genericWorkloadGroupFinalizer implements a generic reconcile.FinalizingReconciler
type genericWorkloadGroupFinalizer struct {
	genericWorkloadGroupReconciler
	finalizingReconciler WorkloadGroupFinalizer
}

func (r genericWorkloadGroupFinalizer) FinalizerName() string {
	return r.finalizingReconciler.WorkloadGroupFinalizerName()
}

func (r genericWorkloadGroupFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_istio_io_v1alpha3.WorkloadGroup)
	if !ok {
		return errors.Errorf("internal error: WorkloadGroup handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeWorkloadGroup(obj)
}

// Reconcile Upsert events for the VirtualService Resource.
// implemented by the user
type VirtualServiceReconciler interface {
	ReconcileVirtualService(obj *networking_istio_io_v1alpha3.VirtualService) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualService Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualServiceDeletionReconciler interface {
	ReconcileVirtualServiceDeletion(req reconcile.Request) error
}

type VirtualServiceReconcilerFuncs struct {
	OnReconcileVirtualService         func(obj *networking_istio_io_v1alpha3.VirtualService) (reconcile.Result, error)
	OnReconcileVirtualServiceDeletion func(req reconcile.Request) error
}

func (f *VirtualServiceReconcilerFuncs) ReconcileVirtualService(obj *networking_istio_io_v1alpha3.VirtualService) (reconcile.Result, error) {
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
	FinalizeVirtualService(obj *networking_istio_io_v1alpha3.VirtualService) error
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
		loop: reconcile.NewLoop(name, "", mgr, &networking_istio_io_v1alpha3.VirtualService{}, options),
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
	obj, ok := object.(*networking_istio_io_v1alpha3.VirtualService)
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
	obj, ok := object.(*networking_istio_io_v1alpha3.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualService(obj)
}

// Reconcile Upsert events for the Sidecar Resource.
// implemented by the user
type SidecarReconciler interface {
	ReconcileSidecar(obj *networking_istio_io_v1alpha3.Sidecar) (reconcile.Result, error)
}

// Reconcile deletion events for the Sidecar Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type SidecarDeletionReconciler interface {
	ReconcileSidecarDeletion(req reconcile.Request) error
}

type SidecarReconcilerFuncs struct {
	OnReconcileSidecar         func(obj *networking_istio_io_v1alpha3.Sidecar) (reconcile.Result, error)
	OnReconcileSidecarDeletion func(req reconcile.Request) error
}

func (f *SidecarReconcilerFuncs) ReconcileSidecar(obj *networking_istio_io_v1alpha3.Sidecar) (reconcile.Result, error) {
	if f.OnReconcileSidecar == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSidecar(obj)
}

func (f *SidecarReconcilerFuncs) ReconcileSidecarDeletion(req reconcile.Request) error {
	if f.OnReconcileSidecarDeletion == nil {
		return nil
	}
	return f.OnReconcileSidecarDeletion(req)
}

// Reconcile and finalize the Sidecar Resource
// implemented by the user
type SidecarFinalizer interface {
	SidecarReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	SidecarFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeSidecar(obj *networking_istio_io_v1alpha3.Sidecar) error
}

type SidecarReconcileLoop interface {
	RunSidecarReconciler(ctx context.Context, rec SidecarReconciler, predicates ...predicate.Predicate) error
}

type sidecarReconcileLoop struct {
	loop reconcile.Loop
}

func NewSidecarReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) SidecarReconcileLoop {
	return &sidecarReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_istio_io_v1alpha3.Sidecar{}, options),
	}
}

func (c *sidecarReconcileLoop) RunSidecarReconciler(ctx context.Context, reconciler SidecarReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericSidecarReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(SidecarFinalizer); ok {
		reconcilerWrapper = genericSidecarFinalizer{
			genericSidecarReconciler: genericReconciler,
			finalizingReconciler:     finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericSidecarHandler implements a generic reconcile.Reconciler
type genericSidecarReconciler struct {
	reconciler SidecarReconciler
}

func (r genericSidecarReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.Sidecar)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Sidecar handler received event for %T", object)
	}
	return r.reconciler.ReconcileSidecar(obj)
}

func (r genericSidecarReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(SidecarDeletionReconciler); ok {
		return deletionReconciler.ReconcileSidecarDeletion(request)
	}
	return nil
}

// genericSidecarFinalizer implements a generic reconcile.FinalizingReconciler
type genericSidecarFinalizer struct {
	genericSidecarReconciler
	finalizingReconciler SidecarFinalizer
}

func (r genericSidecarFinalizer) FinalizerName() string {
	return r.finalizingReconciler.SidecarFinalizerName()
}

func (r genericSidecarFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_istio_io_v1alpha3.Sidecar)
	if !ok {
		return errors.Errorf("internal error: Sidecar handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeSidecar(obj)
}

// Reconcile Upsert events for the EnvoyFilter Resource.
// implemented by the user
type EnvoyFilterReconciler interface {
	ReconcileEnvoyFilter(obj *networking_istio_io_v1alpha3.EnvoyFilter) (reconcile.Result, error)
}

// Reconcile deletion events for the EnvoyFilter Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type EnvoyFilterDeletionReconciler interface {
	ReconcileEnvoyFilterDeletion(req reconcile.Request) error
}

type EnvoyFilterReconcilerFuncs struct {
	OnReconcileEnvoyFilter         func(obj *networking_istio_io_v1alpha3.EnvoyFilter) (reconcile.Result, error)
	OnReconcileEnvoyFilterDeletion func(req reconcile.Request) error
}

func (f *EnvoyFilterReconcilerFuncs) ReconcileEnvoyFilter(obj *networking_istio_io_v1alpha3.EnvoyFilter) (reconcile.Result, error) {
	if f.OnReconcileEnvoyFilter == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileEnvoyFilter(obj)
}

func (f *EnvoyFilterReconcilerFuncs) ReconcileEnvoyFilterDeletion(req reconcile.Request) error {
	if f.OnReconcileEnvoyFilterDeletion == nil {
		return nil
	}
	return f.OnReconcileEnvoyFilterDeletion(req)
}

// Reconcile and finalize the EnvoyFilter Resource
// implemented by the user
type EnvoyFilterFinalizer interface {
	EnvoyFilterReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	EnvoyFilterFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeEnvoyFilter(obj *networking_istio_io_v1alpha3.EnvoyFilter) error
}

type EnvoyFilterReconcileLoop interface {
	RunEnvoyFilterReconciler(ctx context.Context, rec EnvoyFilterReconciler, predicates ...predicate.Predicate) error
}

type envoyFilterReconcileLoop struct {
	loop reconcile.Loop
}

func NewEnvoyFilterReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) EnvoyFilterReconcileLoop {
	return &envoyFilterReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &networking_istio_io_v1alpha3.EnvoyFilter{}, options),
	}
}

func (c *envoyFilterReconcileLoop) RunEnvoyFilterReconciler(ctx context.Context, reconciler EnvoyFilterReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericEnvoyFilterReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(EnvoyFilterFinalizer); ok {
		reconcilerWrapper = genericEnvoyFilterFinalizer{
			genericEnvoyFilterReconciler: genericReconciler,
			finalizingReconciler:         finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericEnvoyFilterHandler implements a generic reconcile.Reconciler
type genericEnvoyFilterReconciler struct {
	reconciler EnvoyFilterReconciler
}

func (r genericEnvoyFilterReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.EnvoyFilter)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: EnvoyFilter handler received event for %T", object)
	}
	return r.reconciler.ReconcileEnvoyFilter(obj)
}

func (r genericEnvoyFilterReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(EnvoyFilterDeletionReconciler); ok {
		return deletionReconciler.ReconcileEnvoyFilterDeletion(request)
	}
	return nil
}

// genericEnvoyFilterFinalizer implements a generic reconcile.FinalizingReconciler
type genericEnvoyFilterFinalizer struct {
	genericEnvoyFilterReconciler
	finalizingReconciler EnvoyFilterFinalizer
}

func (r genericEnvoyFilterFinalizer) FinalizerName() string {
	return r.finalizingReconciler.EnvoyFilterFinalizerName()
}

func (r genericEnvoyFilterFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_istio_io_v1alpha3.EnvoyFilter)
	if !ok {
		return errors.Errorf("internal error: EnvoyFilter handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeEnvoyFilter(obj)
}
