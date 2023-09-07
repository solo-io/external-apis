// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	gateway_networking_k8s_io_v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the ReferenceGrant Resource
// DEPRECATED: Prefer reconciler pattern.
type ReferenceGrantEventHandler interface {
	CreateReferenceGrant(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error
	UpdateReferenceGrant(old, new *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error
	DeleteReferenceGrant(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error
	GenericReferenceGrant(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error
}

type ReferenceGrantEventHandlerFuncs struct {
	OnCreate  func(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error
	OnUpdate  func(old, new *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error
	OnDelete  func(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error
	OnGeneric func(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error
}

func (f *ReferenceGrantEventHandlerFuncs) CreateReferenceGrant(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ReferenceGrantEventHandlerFuncs) DeleteReferenceGrant(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ReferenceGrantEventHandlerFuncs) UpdateReferenceGrant(objOld, objNew *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ReferenceGrantEventHandlerFuncs) GenericReferenceGrant(obj *gateway_networking_k8s_io_v1alpha2.ReferenceGrant) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ReferenceGrantEventWatcher interface {
	AddEventHandler(ctx context.Context, h ReferenceGrantEventHandler, predicates ...predicate.Predicate) error
}

type referenceGrantEventWatcher struct {
	watcher events.EventWatcher
}

func NewReferenceGrantEventWatcher(name string, mgr manager.Manager) ReferenceGrantEventWatcher {
	return &referenceGrantEventWatcher{
		watcher: events.NewWatcher(name, mgr, &gateway_networking_k8s_io_v1alpha2.ReferenceGrant{}),
	}
}

func (c *referenceGrantEventWatcher) AddEventHandler(ctx context.Context, h ReferenceGrantEventHandler, predicates ...predicate.Predicate) error {
	handler := genericReferenceGrantHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericReferenceGrantHandler implements a generic events.EventHandler
type genericReferenceGrantHandler struct {
	handler ReferenceGrantEventHandler
}

func (h genericReferenceGrantHandler) Create(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.ReferenceGrant)
	if !ok {
		return errors.Errorf("internal error: ReferenceGrant handler received event for %T", object)
	}
	return h.handler.CreateReferenceGrant(obj)
}

func (h genericReferenceGrantHandler) Delete(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.ReferenceGrant)
	if !ok {
		return errors.Errorf("internal error: ReferenceGrant handler received event for %T", object)
	}
	return h.handler.DeleteReferenceGrant(obj)
}

func (h genericReferenceGrantHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*gateway_networking_k8s_io_v1alpha2.ReferenceGrant)
	if !ok {
		return errors.Errorf("internal error: ReferenceGrant handler received event for %T", old)
	}
	objNew, ok := new.(*gateway_networking_k8s_io_v1alpha2.ReferenceGrant)
	if !ok {
		return errors.Errorf("internal error: ReferenceGrant handler received event for %T", new)
	}
	return h.handler.UpdateReferenceGrant(objOld, objNew)
}

func (h genericReferenceGrantHandler) Generic(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.ReferenceGrant)
	if !ok {
		return errors.Errorf("internal error: ReferenceGrant handler received event for %T", object)
	}
	return h.handler.GenericReferenceGrant(obj)
}

// Handle events for the GRPCRoute Resource
// DEPRECATED: Prefer reconciler pattern.
type GRPCRouteEventHandler interface {
	CreateGRPCRoute(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error
	UpdateGRPCRoute(old, new *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error
	DeleteGRPCRoute(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error
	GenericGRPCRoute(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error
}

type GRPCRouteEventHandlerFuncs struct {
	OnCreate  func(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error
	OnUpdate  func(old, new *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error
	OnDelete  func(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error
	OnGeneric func(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error
}

func (f *GRPCRouteEventHandlerFuncs) CreateGRPCRoute(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *GRPCRouteEventHandlerFuncs) DeleteGRPCRoute(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *GRPCRouteEventHandlerFuncs) UpdateGRPCRoute(objOld, objNew *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *GRPCRouteEventHandlerFuncs) GenericGRPCRoute(obj *gateway_networking_k8s_io_v1alpha2.GRPCRoute) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type GRPCRouteEventWatcher interface {
	AddEventHandler(ctx context.Context, h GRPCRouteEventHandler, predicates ...predicate.Predicate) error
}

type gRPCRouteEventWatcher struct {
	watcher events.EventWatcher
}

func NewGRPCRouteEventWatcher(name string, mgr manager.Manager) GRPCRouteEventWatcher {
	return &gRPCRouteEventWatcher{
		watcher: events.NewWatcher(name, mgr, &gateway_networking_k8s_io_v1alpha2.GRPCRoute{}),
	}
}

func (c *gRPCRouteEventWatcher) AddEventHandler(ctx context.Context, h GRPCRouteEventHandler, predicates ...predicate.Predicate) error {
	handler := genericGRPCRouteHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericGRPCRouteHandler implements a generic events.EventHandler
type genericGRPCRouteHandler struct {
	handler GRPCRouteEventHandler
}

func (h genericGRPCRouteHandler) Create(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.GRPCRoute)
	if !ok {
		return errors.Errorf("internal error: GRPCRoute handler received event for %T", object)
	}
	return h.handler.CreateGRPCRoute(obj)
}

func (h genericGRPCRouteHandler) Delete(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.GRPCRoute)
	if !ok {
		return errors.Errorf("internal error: GRPCRoute handler received event for %T", object)
	}
	return h.handler.DeleteGRPCRoute(obj)
}

func (h genericGRPCRouteHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*gateway_networking_k8s_io_v1alpha2.GRPCRoute)
	if !ok {
		return errors.Errorf("internal error: GRPCRoute handler received event for %T", old)
	}
	objNew, ok := new.(*gateway_networking_k8s_io_v1alpha2.GRPCRoute)
	if !ok {
		return errors.Errorf("internal error: GRPCRoute handler received event for %T", new)
	}
	return h.handler.UpdateGRPCRoute(objOld, objNew)
}

func (h genericGRPCRouteHandler) Generic(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.GRPCRoute)
	if !ok {
		return errors.Errorf("internal error: GRPCRoute handler received event for %T", object)
	}
	return h.handler.GenericGRPCRoute(obj)
}

// Handle events for the TCPRoute Resource
// DEPRECATED: Prefer reconciler pattern.
type TCPRouteEventHandler interface {
	CreateTCPRoute(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error
	UpdateTCPRoute(old, new *gateway_networking_k8s_io_v1alpha2.TCPRoute) error
	DeleteTCPRoute(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error
	GenericTCPRoute(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error
}

type TCPRouteEventHandlerFuncs struct {
	OnCreate  func(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error
	OnUpdate  func(old, new *gateway_networking_k8s_io_v1alpha2.TCPRoute) error
	OnDelete  func(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error
	OnGeneric func(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error
}

func (f *TCPRouteEventHandlerFuncs) CreateTCPRoute(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *TCPRouteEventHandlerFuncs) DeleteTCPRoute(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *TCPRouteEventHandlerFuncs) UpdateTCPRoute(objOld, objNew *gateway_networking_k8s_io_v1alpha2.TCPRoute) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *TCPRouteEventHandlerFuncs) GenericTCPRoute(obj *gateway_networking_k8s_io_v1alpha2.TCPRoute) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type TCPRouteEventWatcher interface {
	AddEventHandler(ctx context.Context, h TCPRouteEventHandler, predicates ...predicate.Predicate) error
}

type tCPRouteEventWatcher struct {
	watcher events.EventWatcher
}

func NewTCPRouteEventWatcher(name string, mgr manager.Manager) TCPRouteEventWatcher {
	return &tCPRouteEventWatcher{
		watcher: events.NewWatcher(name, mgr, &gateway_networking_k8s_io_v1alpha2.TCPRoute{}),
	}
}

func (c *tCPRouteEventWatcher) AddEventHandler(ctx context.Context, h TCPRouteEventHandler, predicates ...predicate.Predicate) error {
	handler := genericTCPRouteHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericTCPRouteHandler implements a generic events.EventHandler
type genericTCPRouteHandler struct {
	handler TCPRouteEventHandler
}

func (h genericTCPRouteHandler) Create(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.TCPRoute)
	if !ok {
		return errors.Errorf("internal error: TCPRoute handler received event for %T", object)
	}
	return h.handler.CreateTCPRoute(obj)
}

func (h genericTCPRouteHandler) Delete(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.TCPRoute)
	if !ok {
		return errors.Errorf("internal error: TCPRoute handler received event for %T", object)
	}
	return h.handler.DeleteTCPRoute(obj)
}

func (h genericTCPRouteHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*gateway_networking_k8s_io_v1alpha2.TCPRoute)
	if !ok {
		return errors.Errorf("internal error: TCPRoute handler received event for %T", old)
	}
	objNew, ok := new.(*gateway_networking_k8s_io_v1alpha2.TCPRoute)
	if !ok {
		return errors.Errorf("internal error: TCPRoute handler received event for %T", new)
	}
	return h.handler.UpdateTCPRoute(objOld, objNew)
}

func (h genericTCPRouteHandler) Generic(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.TCPRoute)
	if !ok {
		return errors.Errorf("internal error: TCPRoute handler received event for %T", object)
	}
	return h.handler.GenericTCPRoute(obj)
}

// Handle events for the TLSRoute Resource
// DEPRECATED: Prefer reconciler pattern.
type TLSRouteEventHandler interface {
	CreateTLSRoute(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error
	UpdateTLSRoute(old, new *gateway_networking_k8s_io_v1alpha2.TLSRoute) error
	DeleteTLSRoute(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error
	GenericTLSRoute(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error
}

type TLSRouteEventHandlerFuncs struct {
	OnCreate  func(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error
	OnUpdate  func(old, new *gateway_networking_k8s_io_v1alpha2.TLSRoute) error
	OnDelete  func(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error
	OnGeneric func(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error
}

func (f *TLSRouteEventHandlerFuncs) CreateTLSRoute(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *TLSRouteEventHandlerFuncs) DeleteTLSRoute(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *TLSRouteEventHandlerFuncs) UpdateTLSRoute(objOld, objNew *gateway_networking_k8s_io_v1alpha2.TLSRoute) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *TLSRouteEventHandlerFuncs) GenericTLSRoute(obj *gateway_networking_k8s_io_v1alpha2.TLSRoute) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type TLSRouteEventWatcher interface {
	AddEventHandler(ctx context.Context, h TLSRouteEventHandler, predicates ...predicate.Predicate) error
}

type tLSRouteEventWatcher struct {
	watcher events.EventWatcher
}

func NewTLSRouteEventWatcher(name string, mgr manager.Manager) TLSRouteEventWatcher {
	return &tLSRouteEventWatcher{
		watcher: events.NewWatcher(name, mgr, &gateway_networking_k8s_io_v1alpha2.TLSRoute{}),
	}
}

func (c *tLSRouteEventWatcher) AddEventHandler(ctx context.Context, h TLSRouteEventHandler, predicates ...predicate.Predicate) error {
	handler := genericTLSRouteHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericTLSRouteHandler implements a generic events.EventHandler
type genericTLSRouteHandler struct {
	handler TLSRouteEventHandler
}

func (h genericTLSRouteHandler) Create(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.TLSRoute)
	if !ok {
		return errors.Errorf("internal error: TLSRoute handler received event for %T", object)
	}
	return h.handler.CreateTLSRoute(obj)
}

func (h genericTLSRouteHandler) Delete(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.TLSRoute)
	if !ok {
		return errors.Errorf("internal error: TLSRoute handler received event for %T", object)
	}
	return h.handler.DeleteTLSRoute(obj)
}

func (h genericTLSRouteHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*gateway_networking_k8s_io_v1alpha2.TLSRoute)
	if !ok {
		return errors.Errorf("internal error: TLSRoute handler received event for %T", old)
	}
	objNew, ok := new.(*gateway_networking_k8s_io_v1alpha2.TLSRoute)
	if !ok {
		return errors.Errorf("internal error: TLSRoute handler received event for %T", new)
	}
	return h.handler.UpdateTLSRoute(objOld, objNew)
}

func (h genericTLSRouteHandler) Generic(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.TLSRoute)
	if !ok {
		return errors.Errorf("internal error: TLSRoute handler received event for %T", object)
	}
	return h.handler.GenericTLSRoute(obj)
}

// Handle events for the UDPRoute Resource
// DEPRECATED: Prefer reconciler pattern.
type UDPRouteEventHandler interface {
	CreateUDPRoute(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error
	UpdateUDPRoute(old, new *gateway_networking_k8s_io_v1alpha2.UDPRoute) error
	DeleteUDPRoute(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error
	GenericUDPRoute(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error
}

type UDPRouteEventHandlerFuncs struct {
	OnCreate  func(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error
	OnUpdate  func(old, new *gateway_networking_k8s_io_v1alpha2.UDPRoute) error
	OnDelete  func(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error
	OnGeneric func(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error
}

func (f *UDPRouteEventHandlerFuncs) CreateUDPRoute(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *UDPRouteEventHandlerFuncs) DeleteUDPRoute(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *UDPRouteEventHandlerFuncs) UpdateUDPRoute(objOld, objNew *gateway_networking_k8s_io_v1alpha2.UDPRoute) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *UDPRouteEventHandlerFuncs) GenericUDPRoute(obj *gateway_networking_k8s_io_v1alpha2.UDPRoute) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type UDPRouteEventWatcher interface {
	AddEventHandler(ctx context.Context, h UDPRouteEventHandler, predicates ...predicate.Predicate) error
}

type uDPRouteEventWatcher struct {
	watcher events.EventWatcher
}

func NewUDPRouteEventWatcher(name string, mgr manager.Manager) UDPRouteEventWatcher {
	return &uDPRouteEventWatcher{
		watcher: events.NewWatcher(name, mgr, &gateway_networking_k8s_io_v1alpha2.UDPRoute{}),
	}
}

func (c *uDPRouteEventWatcher) AddEventHandler(ctx context.Context, h UDPRouteEventHandler, predicates ...predicate.Predicate) error {
	handler := genericUDPRouteHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericUDPRouteHandler implements a generic events.EventHandler
type genericUDPRouteHandler struct {
	handler UDPRouteEventHandler
}

func (h genericUDPRouteHandler) Create(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.UDPRoute)
	if !ok {
		return errors.Errorf("internal error: UDPRoute handler received event for %T", object)
	}
	return h.handler.CreateUDPRoute(obj)
}

func (h genericUDPRouteHandler) Delete(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.UDPRoute)
	if !ok {
		return errors.Errorf("internal error: UDPRoute handler received event for %T", object)
	}
	return h.handler.DeleteUDPRoute(obj)
}

func (h genericUDPRouteHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*gateway_networking_k8s_io_v1alpha2.UDPRoute)
	if !ok {
		return errors.Errorf("internal error: UDPRoute handler received event for %T", old)
	}
	objNew, ok := new.(*gateway_networking_k8s_io_v1alpha2.UDPRoute)
	if !ok {
		return errors.Errorf("internal error: UDPRoute handler received event for %T", new)
	}
	return h.handler.UpdateUDPRoute(objOld, objNew)
}

func (h genericUDPRouteHandler) Generic(object client.Object) error {
	obj, ok := object.(*gateway_networking_k8s_io_v1alpha2.UDPRoute)
	if !ok {
		return errors.Errorf("internal error: UDPRoute handler received event for %T", object)
	}
	return h.handler.GenericUDPRoute(obj)
}
