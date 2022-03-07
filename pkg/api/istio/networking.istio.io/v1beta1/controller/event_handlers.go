// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_istio_io_v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the DestinationRule Resource
// DEPRECATED: Prefer reconciler pattern.
type DestinationRuleEventHandler interface {
	CreateDestinationRule(obj *networking_istio_io_v1beta1.DestinationRule) error
	UpdateDestinationRule(old, new *networking_istio_io_v1beta1.DestinationRule) error
	DeleteDestinationRule(obj *networking_istio_io_v1beta1.DestinationRule) error
	GenericDestinationRule(obj *networking_istio_io_v1beta1.DestinationRule) error
}

type DestinationRuleEventHandlerFuncs struct {
	OnCreate  func(obj *networking_istio_io_v1beta1.DestinationRule) error
	OnUpdate  func(old, new *networking_istio_io_v1beta1.DestinationRule) error
	OnDelete  func(obj *networking_istio_io_v1beta1.DestinationRule) error
	OnGeneric func(obj *networking_istio_io_v1beta1.DestinationRule) error
}

func (f *DestinationRuleEventHandlerFuncs) CreateDestinationRule(obj *networking_istio_io_v1beta1.DestinationRule) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *DestinationRuleEventHandlerFuncs) DeleteDestinationRule(obj *networking_istio_io_v1beta1.DestinationRule) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *DestinationRuleEventHandlerFuncs) UpdateDestinationRule(objOld, objNew *networking_istio_io_v1beta1.DestinationRule) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *DestinationRuleEventHandlerFuncs) GenericDestinationRule(obj *networking_istio_io_v1beta1.DestinationRule) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type DestinationRuleEventWatcher interface {
	AddEventHandler(ctx context.Context, h DestinationRuleEventHandler, predicates ...predicate.Predicate) error
}

type destinationRuleEventWatcher struct {
	watcher events.EventWatcher
}

func NewDestinationRuleEventWatcher(name string, mgr manager.Manager) DestinationRuleEventWatcher {
	return &destinationRuleEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_istio_io_v1beta1.DestinationRule{}),
	}
}

func (c *destinationRuleEventWatcher) AddEventHandler(ctx context.Context, h DestinationRuleEventHandler, predicates ...predicate.Predicate) error {
	handler := genericDestinationRuleHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericDestinationRuleHandler implements a generic events.EventHandler
type genericDestinationRuleHandler struct {
	handler DestinationRuleEventHandler
}

func (h genericDestinationRuleHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.DestinationRule)
	if !ok {
		return errors.Errorf("internal error: DestinationRule handler received event for %T", object)
	}
	return h.handler.CreateDestinationRule(obj)
}

func (h genericDestinationRuleHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.DestinationRule)
	if !ok {
		return errors.Errorf("internal error: DestinationRule handler received event for %T", object)
	}
	return h.handler.DeleteDestinationRule(obj)
}

func (h genericDestinationRuleHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_istio_io_v1beta1.DestinationRule)
	if !ok {
		return errors.Errorf("internal error: DestinationRule handler received event for %T", old)
	}
	objNew, ok := new.(*networking_istio_io_v1beta1.DestinationRule)
	if !ok {
		return errors.Errorf("internal error: DestinationRule handler received event for %T", new)
	}
	return h.handler.UpdateDestinationRule(objOld, objNew)
}

func (h genericDestinationRuleHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.DestinationRule)
	if !ok {
		return errors.Errorf("internal error: DestinationRule handler received event for %T", object)
	}
	return h.handler.GenericDestinationRule(obj)
}

// Handle events for the Gateway Resource
// DEPRECATED: Prefer reconciler pattern.
type GatewayEventHandler interface {
	CreateGateway(obj *networking_istio_io_v1beta1.Gateway) error
	UpdateGateway(old, new *networking_istio_io_v1beta1.Gateway) error
	DeleteGateway(obj *networking_istio_io_v1beta1.Gateway) error
	GenericGateway(obj *networking_istio_io_v1beta1.Gateway) error
}

type GatewayEventHandlerFuncs struct {
	OnCreate  func(obj *networking_istio_io_v1beta1.Gateway) error
	OnUpdate  func(old, new *networking_istio_io_v1beta1.Gateway) error
	OnDelete  func(obj *networking_istio_io_v1beta1.Gateway) error
	OnGeneric func(obj *networking_istio_io_v1beta1.Gateway) error
}

func (f *GatewayEventHandlerFuncs) CreateGateway(obj *networking_istio_io_v1beta1.Gateway) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GatewayEventHandlerFuncs) DeleteGateway(obj *networking_istio_io_v1beta1.Gateway) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GatewayEventHandlerFuncs) UpdateGateway(objOld, objNew *networking_istio_io_v1beta1.Gateway) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GatewayEventHandlerFuncs) GenericGateway(obj *networking_istio_io_v1beta1.Gateway) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GatewayEventWatcher interface {
	AddEventHandler(ctx context.Context, h GatewayEventHandler, predicates ...predicate.Predicate) error
}

type gatewayEventWatcher struct {
	watcher events.EventWatcher
}

func NewGatewayEventWatcher(name string, mgr manager.Manager) GatewayEventWatcher {
	return &gatewayEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_istio_io_v1beta1.Gateway{}),
	}
}

func (c *gatewayEventWatcher) AddEventHandler(ctx context.Context, h GatewayEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGatewayHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGatewayHandler implements a generic events.EventHandler
type genericGatewayHandler struct {
	handler GatewayEventHandler
}

func (h genericGatewayHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return h.handler.CreateGateway(obj)
}

func (h genericGatewayHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return h.handler.DeleteGateway(obj)
}

func (h genericGatewayHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_istio_io_v1beta1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", old)
	}
	objNew, ok := new.(*networking_istio_io_v1beta1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", new)
	}
	return h.handler.UpdateGateway(objOld, objNew)
}

func (h genericGatewayHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return h.handler.GenericGateway(obj)
}

// Handle events for the ServiceEntry Resource
// DEPRECATED: Prefer reconciler pattern.
type ServiceEntryEventHandler interface {
	CreateServiceEntry(obj *networking_istio_io_v1beta1.ServiceEntry) error
	UpdateServiceEntry(old, new *networking_istio_io_v1beta1.ServiceEntry) error
	DeleteServiceEntry(obj *networking_istio_io_v1beta1.ServiceEntry) error
	GenericServiceEntry(obj *networking_istio_io_v1beta1.ServiceEntry) error
}

type ServiceEntryEventHandlerFuncs struct {
	OnCreate  func(obj *networking_istio_io_v1beta1.ServiceEntry) error
	OnUpdate  func(old, new *networking_istio_io_v1beta1.ServiceEntry) error
	OnDelete  func(obj *networking_istio_io_v1beta1.ServiceEntry) error
	OnGeneric func(obj *networking_istio_io_v1beta1.ServiceEntry) error
}

func (f *ServiceEntryEventHandlerFuncs) CreateServiceEntry(obj *networking_istio_io_v1beta1.ServiceEntry) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ServiceEntryEventHandlerFuncs) DeleteServiceEntry(obj *networking_istio_io_v1beta1.ServiceEntry) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ServiceEntryEventHandlerFuncs) UpdateServiceEntry(objOld, objNew *networking_istio_io_v1beta1.ServiceEntry) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ServiceEntryEventHandlerFuncs) GenericServiceEntry(obj *networking_istio_io_v1beta1.ServiceEntry) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ServiceEntryEventWatcher interface {
	AddEventHandler(ctx context.Context, h ServiceEntryEventHandler, predicates ...predicate.Predicate) error
}

type serviceEntryEventWatcher struct {
	watcher events.EventWatcher
}

func NewServiceEntryEventWatcher(name string, mgr manager.Manager) ServiceEntryEventWatcher {
	return &serviceEntryEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_istio_io_v1beta1.ServiceEntry{}),
	}
}

func (c *serviceEntryEventWatcher) AddEventHandler(ctx context.Context, h ServiceEntryEventHandler, predicates ...predicate.Predicate) error {
	handler := genericServiceEntryHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericServiceEntryHandler implements a generic events.EventHandler
type genericServiceEntryHandler struct {
	handler ServiceEntryEventHandler
}

func (h genericServiceEntryHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.ServiceEntry)
	if !ok {
		return errors.Errorf("internal error: ServiceEntry handler received event for %T", object)
	}
	return h.handler.CreateServiceEntry(obj)
}

func (h genericServiceEntryHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.ServiceEntry)
	if !ok {
		return errors.Errorf("internal error: ServiceEntry handler received event for %T", object)
	}
	return h.handler.DeleteServiceEntry(obj)
}

func (h genericServiceEntryHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_istio_io_v1beta1.ServiceEntry)
	if !ok {
		return errors.Errorf("internal error: ServiceEntry handler received event for %T", old)
	}
	objNew, ok := new.(*networking_istio_io_v1beta1.ServiceEntry)
	if !ok {
		return errors.Errorf("internal error: ServiceEntry handler received event for %T", new)
	}
	return h.handler.UpdateServiceEntry(objOld, objNew)
}

func (h genericServiceEntryHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.ServiceEntry)
	if !ok {
		return errors.Errorf("internal error: ServiceEntry handler received event for %T", object)
	}
	return h.handler.GenericServiceEntry(obj)
}

// Handle events for the WorkloadEntry Resource
// DEPRECATED: Prefer reconciler pattern.
type WorkloadEntryEventHandler interface {
	CreateWorkloadEntry(obj *networking_istio_io_v1beta1.WorkloadEntry) error
	UpdateWorkloadEntry(old, new *networking_istio_io_v1beta1.WorkloadEntry) error
	DeleteWorkloadEntry(obj *networking_istio_io_v1beta1.WorkloadEntry) error
	GenericWorkloadEntry(obj *networking_istio_io_v1beta1.WorkloadEntry) error
}

type WorkloadEntryEventHandlerFuncs struct {
	OnCreate  func(obj *networking_istio_io_v1beta1.WorkloadEntry) error
	OnUpdate  func(old, new *networking_istio_io_v1beta1.WorkloadEntry) error
	OnDelete  func(obj *networking_istio_io_v1beta1.WorkloadEntry) error
	OnGeneric func(obj *networking_istio_io_v1beta1.WorkloadEntry) error
}

func (f *WorkloadEntryEventHandlerFuncs) CreateWorkloadEntry(obj *networking_istio_io_v1beta1.WorkloadEntry) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *WorkloadEntryEventHandlerFuncs) DeleteWorkloadEntry(obj *networking_istio_io_v1beta1.WorkloadEntry) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *WorkloadEntryEventHandlerFuncs) UpdateWorkloadEntry(objOld, objNew *networking_istio_io_v1beta1.WorkloadEntry) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *WorkloadEntryEventHandlerFuncs) GenericWorkloadEntry(obj *networking_istio_io_v1beta1.WorkloadEntry) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type WorkloadEntryEventWatcher interface {
	AddEventHandler(ctx context.Context, h WorkloadEntryEventHandler, predicates ...predicate.Predicate) error
}

type workloadEntryEventWatcher struct {
	watcher events.EventWatcher
}

func NewWorkloadEntryEventWatcher(name string, mgr manager.Manager) WorkloadEntryEventWatcher {
	return &workloadEntryEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_istio_io_v1beta1.WorkloadEntry{}),
	}
}

func (c *workloadEntryEventWatcher) AddEventHandler(ctx context.Context, h WorkloadEntryEventHandler, predicates ...predicate.Predicate) error {
	handler := genericWorkloadEntryHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericWorkloadEntryHandler implements a generic events.EventHandler
type genericWorkloadEntryHandler struct {
	handler WorkloadEntryEventHandler
}

func (h genericWorkloadEntryHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.WorkloadEntry)
	if !ok {
		return errors.Errorf("internal error: WorkloadEntry handler received event for %T", object)
	}
	return h.handler.CreateWorkloadEntry(obj)
}

func (h genericWorkloadEntryHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.WorkloadEntry)
	if !ok {
		return errors.Errorf("internal error: WorkloadEntry handler received event for %T", object)
	}
	return h.handler.DeleteWorkloadEntry(obj)
}

func (h genericWorkloadEntryHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_istio_io_v1beta1.WorkloadEntry)
	if !ok {
		return errors.Errorf("internal error: WorkloadEntry handler received event for %T", old)
	}
	objNew, ok := new.(*networking_istio_io_v1beta1.WorkloadEntry)
	if !ok {
		return errors.Errorf("internal error: WorkloadEntry handler received event for %T", new)
	}
	return h.handler.UpdateWorkloadEntry(objOld, objNew)
}

func (h genericWorkloadEntryHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.WorkloadEntry)
	if !ok {
		return errors.Errorf("internal error: WorkloadEntry handler received event for %T", object)
	}
	return h.handler.GenericWorkloadEntry(obj)
}

// Handle events for the VirtualService Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualServiceEventHandler interface {
	CreateVirtualService(obj *networking_istio_io_v1beta1.VirtualService) error
	UpdateVirtualService(old, new *networking_istio_io_v1beta1.VirtualService) error
	DeleteVirtualService(obj *networking_istio_io_v1beta1.VirtualService) error
	GenericVirtualService(obj *networking_istio_io_v1beta1.VirtualService) error
}

type VirtualServiceEventHandlerFuncs struct {
	OnCreate  func(obj *networking_istio_io_v1beta1.VirtualService) error
	OnUpdate  func(old, new *networking_istio_io_v1beta1.VirtualService) error
	OnDelete  func(obj *networking_istio_io_v1beta1.VirtualService) error
	OnGeneric func(obj *networking_istio_io_v1beta1.VirtualService) error
}

func (f *VirtualServiceEventHandlerFuncs) CreateVirtualService(obj *networking_istio_io_v1beta1.VirtualService) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualServiceEventHandlerFuncs) DeleteVirtualService(obj *networking_istio_io_v1beta1.VirtualService) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualServiceEventHandlerFuncs) UpdateVirtualService(objOld, objNew *networking_istio_io_v1beta1.VirtualService) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualServiceEventHandlerFuncs) GenericVirtualService(obj *networking_istio_io_v1beta1.VirtualService) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type VirtualServiceEventWatcher interface {
	AddEventHandler(ctx context.Context, h VirtualServiceEventHandler, predicates ...predicate.Predicate) error
}

type virtualServiceEventWatcher struct {
	watcher events.EventWatcher
}

func NewVirtualServiceEventWatcher(name string, mgr manager.Manager) VirtualServiceEventWatcher {
	return &virtualServiceEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_istio_io_v1beta1.VirtualService{}),
	}
}

func (c *virtualServiceEventWatcher) AddEventHandler(ctx context.Context, h VirtualServiceEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualServiceHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericVirtualServiceHandler implements a generic events.EventHandler
type genericVirtualServiceHandler struct {
	handler VirtualServiceEventHandler
}

func (h genericVirtualServiceHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.CreateVirtualService(obj)
}

func (h genericVirtualServiceHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.DeleteVirtualService(obj)
}

func (h genericVirtualServiceHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_istio_io_v1beta1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", old)
	}
	objNew, ok := new.(*networking_istio_io_v1beta1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", new)
	}
	return h.handler.UpdateVirtualService(objOld, objNew)
}

func (h genericVirtualServiceHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.GenericVirtualService(obj)
}

// Handle events for the Sidecar Resource
// DEPRECATED: Prefer reconciler pattern.
type SidecarEventHandler interface {
	CreateSidecar(obj *networking_istio_io_v1beta1.Sidecar) error
	UpdateSidecar(old, new *networking_istio_io_v1beta1.Sidecar) error
	DeleteSidecar(obj *networking_istio_io_v1beta1.Sidecar) error
	GenericSidecar(obj *networking_istio_io_v1beta1.Sidecar) error
}

type SidecarEventHandlerFuncs struct {
	OnCreate  func(obj *networking_istio_io_v1beta1.Sidecar) error
	OnUpdate  func(old, new *networking_istio_io_v1beta1.Sidecar) error
	OnDelete  func(obj *networking_istio_io_v1beta1.Sidecar) error
	OnGeneric func(obj *networking_istio_io_v1beta1.Sidecar) error
}

func (f *SidecarEventHandlerFuncs) CreateSidecar(obj *networking_istio_io_v1beta1.Sidecar) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *SidecarEventHandlerFuncs) DeleteSidecar(obj *networking_istio_io_v1beta1.Sidecar) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *SidecarEventHandlerFuncs) UpdateSidecar(objOld, objNew *networking_istio_io_v1beta1.Sidecar) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *SidecarEventHandlerFuncs) GenericSidecar(obj *networking_istio_io_v1beta1.Sidecar) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type SidecarEventWatcher interface {
	AddEventHandler(ctx context.Context, h SidecarEventHandler, predicates ...predicate.Predicate) error
}

type sidecarEventWatcher struct {
	watcher events.EventWatcher
}

func NewSidecarEventWatcher(name string, mgr manager.Manager) SidecarEventWatcher {
	return &sidecarEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_istio_io_v1beta1.Sidecar{}),
	}
}

func (c *sidecarEventWatcher) AddEventHandler(ctx context.Context, h SidecarEventHandler, predicates ...predicate.Predicate) error {
	handler := genericSidecarHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericSidecarHandler implements a generic events.EventHandler
type genericSidecarHandler struct {
	handler SidecarEventHandler
}

func (h genericSidecarHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.Sidecar)
	if !ok {
		return errors.Errorf("internal error: Sidecar handler received event for %T", object)
	}
	return h.handler.CreateSidecar(obj)
}

func (h genericSidecarHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.Sidecar)
	if !ok {
		return errors.Errorf("internal error: Sidecar handler received event for %T", object)
	}
	return h.handler.DeleteSidecar(obj)
}

func (h genericSidecarHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_istio_io_v1beta1.Sidecar)
	if !ok {
		return errors.Errorf("internal error: Sidecar handler received event for %T", old)
	}
	objNew, ok := new.(*networking_istio_io_v1beta1.Sidecar)
	if !ok {
		return errors.Errorf("internal error: Sidecar handler received event for %T", new)
	}
	return h.handler.UpdateSidecar(objOld, objNew)
}

func (h genericSidecarHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_istio_io_v1beta1.Sidecar)
	if !ok {
		return errors.Errorf("internal error: Sidecar handler received event for %T", object)
	}
	return h.handler.GenericSidecar(obj)
}
