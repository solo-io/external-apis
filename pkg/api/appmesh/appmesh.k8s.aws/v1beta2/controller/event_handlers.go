// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the Mesh Resource
// DEPRECATED: Prefer reconciler pattern.
type MeshEventHandler interface {
	CreateMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) error
	UpdateMesh(old, new *appmesh_k8s_aws_v1beta2.Mesh) error
	DeleteMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) error
	GenericMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) error
}

type MeshEventHandlerFuncs struct {
	OnCreate  func(obj *appmesh_k8s_aws_v1beta2.Mesh) error
	OnUpdate  func(old, new *appmesh_k8s_aws_v1beta2.Mesh) error
	OnDelete  func(obj *appmesh_k8s_aws_v1beta2.Mesh) error
	OnGeneric func(obj *appmesh_k8s_aws_v1beta2.Mesh) error
}

func (f *MeshEventHandlerFuncs) CreateMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *MeshEventHandlerFuncs) DeleteMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *MeshEventHandlerFuncs) UpdateMesh(objOld, objNew *appmesh_k8s_aws_v1beta2.Mesh) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *MeshEventHandlerFuncs) GenericMesh(obj *appmesh_k8s_aws_v1beta2.Mesh) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type MeshEventWatcher interface {
	AddEventHandler(ctx context.Context, h MeshEventHandler, predicates ...predicate.Predicate) error
}

type meshEventWatcher struct {
	watcher events.EventWatcher
}

func NewMeshEventWatcher(name string, mgr manager.Manager) MeshEventWatcher {
	return &meshEventWatcher{
		watcher: events.NewWatcher(name, mgr, &appmesh_k8s_aws_v1beta2.Mesh{}),
	}
}

func (c *meshEventWatcher) AddEventHandler(ctx context.Context, h MeshEventHandler, predicates ...predicate.Predicate) error {
	handler := genericMeshHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericMeshHandler implements a generic events.EventHandler
type genericMeshHandler struct {
	handler MeshEventHandler
}

func (h genericMeshHandler) Create(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.Mesh)
	if !ok {
		return errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return h.handler.CreateMesh(obj)
}

func (h genericMeshHandler) Delete(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.Mesh)
	if !ok {
		return errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return h.handler.DeleteMesh(obj)
}

func (h genericMeshHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*appmesh_k8s_aws_v1beta2.Mesh)
	if !ok {
		return errors.Errorf("internal error: Mesh handler received event for %T", old)
	}
	objNew, ok := new.(*appmesh_k8s_aws_v1beta2.Mesh)
	if !ok {
		return errors.Errorf("internal error: Mesh handler received event for %T", new)
	}
	return h.handler.UpdateMesh(objOld, objNew)
}

func (h genericMeshHandler) Generic(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.Mesh)
	if !ok {
		return errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return h.handler.GenericMesh(obj)
}

// Handle events for the VirtualService Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualServiceEventHandler interface {
	CreateVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) error
	UpdateVirtualService(old, new *appmesh_k8s_aws_v1beta2.VirtualService) error
	DeleteVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) error
	GenericVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) error
}

type VirtualServiceEventHandlerFuncs struct {
	OnCreate  func(obj *appmesh_k8s_aws_v1beta2.VirtualService) error
	OnUpdate  func(old, new *appmesh_k8s_aws_v1beta2.VirtualService) error
	OnDelete  func(obj *appmesh_k8s_aws_v1beta2.VirtualService) error
	OnGeneric func(obj *appmesh_k8s_aws_v1beta2.VirtualService) error
}

func (f *VirtualServiceEventHandlerFuncs) CreateVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualServiceEventHandlerFuncs) DeleteVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualServiceEventHandlerFuncs) UpdateVirtualService(objOld, objNew *appmesh_k8s_aws_v1beta2.VirtualService) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualServiceEventHandlerFuncs) GenericVirtualService(obj *appmesh_k8s_aws_v1beta2.VirtualService) error {
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
		watcher: events.NewWatcher(name, mgr, &appmesh_k8s_aws_v1beta2.VirtualService{}),
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
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.CreateVirtualService(obj)
}

func (h genericVirtualServiceHandler) Delete(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.DeleteVirtualService(obj)
}

func (h genericVirtualServiceHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*appmesh_k8s_aws_v1beta2.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", old)
	}
	objNew, ok := new.(*appmesh_k8s_aws_v1beta2.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", new)
	}
	return h.handler.UpdateVirtualService(objOld, objNew)
}

func (h genericVirtualServiceHandler) Generic(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualService)
	if !ok {
		return errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return h.handler.GenericVirtualService(obj)
}

// Handle events for the VirtualNode Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualNodeEventHandler interface {
	CreateVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error
	UpdateVirtualNode(old, new *appmesh_k8s_aws_v1beta2.VirtualNode) error
	DeleteVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error
	GenericVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error
}

type VirtualNodeEventHandlerFuncs struct {
	OnCreate  func(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error
	OnUpdate  func(old, new *appmesh_k8s_aws_v1beta2.VirtualNode) error
	OnDelete  func(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error
	OnGeneric func(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error
}

func (f *VirtualNodeEventHandlerFuncs) CreateVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualNodeEventHandlerFuncs) DeleteVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualNodeEventHandlerFuncs) UpdateVirtualNode(objOld, objNew *appmesh_k8s_aws_v1beta2.VirtualNode) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualNodeEventHandlerFuncs) GenericVirtualNode(obj *appmesh_k8s_aws_v1beta2.VirtualNode) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type VirtualNodeEventWatcher interface {
	AddEventHandler(ctx context.Context, h VirtualNodeEventHandler, predicates ...predicate.Predicate) error
}

type virtualNodeEventWatcher struct {
	watcher events.EventWatcher
}

func NewVirtualNodeEventWatcher(name string, mgr manager.Manager) VirtualNodeEventWatcher {
	return &virtualNodeEventWatcher{
		watcher: events.NewWatcher(name, mgr, &appmesh_k8s_aws_v1beta2.VirtualNode{}),
	}
}

func (c *virtualNodeEventWatcher) AddEventHandler(ctx context.Context, h VirtualNodeEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualNodeHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericVirtualNodeHandler implements a generic events.EventHandler
type genericVirtualNodeHandler struct {
	handler VirtualNodeEventHandler
}

func (h genericVirtualNodeHandler) Create(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	if !ok {
		return errors.Errorf("internal error: VirtualNode handler received event for %T", object)
	}
	return h.handler.CreateVirtualNode(obj)
}

func (h genericVirtualNodeHandler) Delete(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	if !ok {
		return errors.Errorf("internal error: VirtualNode handler received event for %T", object)
	}
	return h.handler.DeleteVirtualNode(obj)
}

func (h genericVirtualNodeHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	if !ok {
		return errors.Errorf("internal error: VirtualNode handler received event for %T", old)
	}
	objNew, ok := new.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	if !ok {
		return errors.Errorf("internal error: VirtualNode handler received event for %T", new)
	}
	return h.handler.UpdateVirtualNode(objOld, objNew)
}

func (h genericVirtualNodeHandler) Generic(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualNode)
	if !ok {
		return errors.Errorf("internal error: VirtualNode handler received event for %T", object)
	}
	return h.handler.GenericVirtualNode(obj)
}

// Handle events for the VirtualRouter Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualRouterEventHandler interface {
	CreateVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error
	UpdateVirtualRouter(old, new *appmesh_k8s_aws_v1beta2.VirtualRouter) error
	DeleteVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error
	GenericVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error
}

type VirtualRouterEventHandlerFuncs struct {
	OnCreate  func(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error
	OnUpdate  func(old, new *appmesh_k8s_aws_v1beta2.VirtualRouter) error
	OnDelete  func(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error
	OnGeneric func(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error
}

func (f *VirtualRouterEventHandlerFuncs) CreateVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualRouterEventHandlerFuncs) DeleteVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualRouterEventHandlerFuncs) UpdateVirtualRouter(objOld, objNew *appmesh_k8s_aws_v1beta2.VirtualRouter) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualRouterEventHandlerFuncs) GenericVirtualRouter(obj *appmesh_k8s_aws_v1beta2.VirtualRouter) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type VirtualRouterEventWatcher interface {
	AddEventHandler(ctx context.Context, h VirtualRouterEventHandler, predicates ...predicate.Predicate) error
}

type virtualRouterEventWatcher struct {
	watcher events.EventWatcher
}

func NewVirtualRouterEventWatcher(name string, mgr manager.Manager) VirtualRouterEventWatcher {
	return &virtualRouterEventWatcher{
		watcher: events.NewWatcher(name, mgr, &appmesh_k8s_aws_v1beta2.VirtualRouter{}),
	}
}

func (c *virtualRouterEventWatcher) AddEventHandler(ctx context.Context, h VirtualRouterEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualRouterHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericVirtualRouterHandler implements a generic events.EventHandler
type genericVirtualRouterHandler struct {
	handler VirtualRouterEventHandler
}

func (h genericVirtualRouterHandler) Create(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	if !ok {
		return errors.Errorf("internal error: VirtualRouter handler received event for %T", object)
	}
	return h.handler.CreateVirtualRouter(obj)
}

func (h genericVirtualRouterHandler) Delete(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	if !ok {
		return errors.Errorf("internal error: VirtualRouter handler received event for %T", object)
	}
	return h.handler.DeleteVirtualRouter(obj)
}

func (h genericVirtualRouterHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	if !ok {
		return errors.Errorf("internal error: VirtualRouter handler received event for %T", old)
	}
	objNew, ok := new.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	if !ok {
		return errors.Errorf("internal error: VirtualRouter handler received event for %T", new)
	}
	return h.handler.UpdateVirtualRouter(objOld, objNew)
}

func (h genericVirtualRouterHandler) Generic(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualRouter)
	if !ok {
		return errors.Errorf("internal error: VirtualRouter handler received event for %T", object)
	}
	return h.handler.GenericVirtualRouter(obj)
}

// Handle events for the VirtualGateway Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualGatewayEventHandler interface {
	CreateVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error
	UpdateVirtualGateway(old, new *appmesh_k8s_aws_v1beta2.VirtualGateway) error
	DeleteVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error
	GenericVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error
}

type VirtualGatewayEventHandlerFuncs struct {
	OnCreate  func(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error
	OnUpdate  func(old, new *appmesh_k8s_aws_v1beta2.VirtualGateway) error
	OnDelete  func(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error
	OnGeneric func(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error
}

func (f *VirtualGatewayEventHandlerFuncs) CreateVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualGatewayEventHandlerFuncs) DeleteVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualGatewayEventHandlerFuncs) UpdateVirtualGateway(objOld, objNew *appmesh_k8s_aws_v1beta2.VirtualGateway) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualGatewayEventHandlerFuncs) GenericVirtualGateway(obj *appmesh_k8s_aws_v1beta2.VirtualGateway) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type VirtualGatewayEventWatcher interface {
	AddEventHandler(ctx context.Context, h VirtualGatewayEventHandler, predicates ...predicate.Predicate) error
}

type virtualGatewayEventWatcher struct {
	watcher events.EventWatcher
}

func NewVirtualGatewayEventWatcher(name string, mgr manager.Manager) VirtualGatewayEventWatcher {
	return &virtualGatewayEventWatcher{
		watcher: events.NewWatcher(name, mgr, &appmesh_k8s_aws_v1beta2.VirtualGateway{}),
	}
}

func (c *virtualGatewayEventWatcher) AddEventHandler(ctx context.Context, h VirtualGatewayEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualGatewayHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericVirtualGatewayHandler implements a generic events.EventHandler
type genericVirtualGatewayHandler struct {
	handler VirtualGatewayEventHandler
}

func (h genericVirtualGatewayHandler) Create(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return h.handler.CreateVirtualGateway(obj)
}

func (h genericVirtualGatewayHandler) Delete(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return h.handler.DeleteVirtualGateway(obj)
}

func (h genericVirtualGatewayHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", old)
	}
	objNew, ok := new.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", new)
	}
	return h.handler.UpdateVirtualGateway(objOld, objNew)
}

func (h genericVirtualGatewayHandler) Generic(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.VirtualGateway)
	if !ok {
		return errors.Errorf("internal error: VirtualGateway handler received event for %T", object)
	}
	return h.handler.GenericVirtualGateway(obj)
}

// Handle events for the GatewayRoute Resource
// DEPRECATED: Prefer reconciler pattern.
type GatewayRouteEventHandler interface {
	CreateGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error
	UpdateGatewayRoute(old, new *appmesh_k8s_aws_v1beta2.GatewayRoute) error
	DeleteGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error
	GenericGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error
}

type GatewayRouteEventHandlerFuncs struct {
	OnCreate  func(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error
	OnUpdate  func(old, new *appmesh_k8s_aws_v1beta2.GatewayRoute) error
	OnDelete  func(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error
	OnGeneric func(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error
}

func (f *GatewayRouteEventHandlerFuncs) CreateGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GatewayRouteEventHandlerFuncs) DeleteGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GatewayRouteEventHandlerFuncs) UpdateGatewayRoute(objOld, objNew *appmesh_k8s_aws_v1beta2.GatewayRoute) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GatewayRouteEventHandlerFuncs) GenericGatewayRoute(obj *appmesh_k8s_aws_v1beta2.GatewayRoute) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GatewayRouteEventWatcher interface {
	AddEventHandler(ctx context.Context, h GatewayRouteEventHandler, predicates ...predicate.Predicate) error
}

type gatewayRouteEventWatcher struct {
	watcher events.EventWatcher
}

func NewGatewayRouteEventWatcher(name string, mgr manager.Manager) GatewayRouteEventWatcher {
	return &gatewayRouteEventWatcher{
		watcher: events.NewWatcher(name, mgr, &appmesh_k8s_aws_v1beta2.GatewayRoute{}),
	}
}

func (c *gatewayRouteEventWatcher) AddEventHandler(ctx context.Context, h GatewayRouteEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGatewayRouteHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGatewayRouteHandler implements a generic events.EventHandler
type genericGatewayRouteHandler struct {
	handler GatewayRouteEventHandler
}

func (h genericGatewayRouteHandler) Create(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	if !ok {
		return errors.Errorf("internal error: GatewayRoute handler received event for %T", object)
	}
	return h.handler.CreateGatewayRoute(obj)
}

func (h genericGatewayRouteHandler) Delete(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	if !ok {
		return errors.Errorf("internal error: GatewayRoute handler received event for %T", object)
	}
	return h.handler.DeleteGatewayRoute(obj)
}

func (h genericGatewayRouteHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	if !ok {
		return errors.Errorf("internal error: GatewayRoute handler received event for %T", old)
	}
	objNew, ok := new.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	if !ok {
		return errors.Errorf("internal error: GatewayRoute handler received event for %T", new)
	}
	return h.handler.UpdateGatewayRoute(objOld, objNew)
}

func (h genericGatewayRouteHandler) Generic(object client.Object) error {
	obj, ok := object.(*appmesh_k8s_aws_v1beta2.GatewayRoute)
	if !ok {
		return errors.Errorf("internal error: GatewayRoute handler received event for %T", object)
	}
	return h.handler.GenericGatewayRoute(obj)
}
