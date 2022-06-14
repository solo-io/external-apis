// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	cilium_io_v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the CiliumNetworkPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type CiliumNetworkPolicyEventHandler interface {
	CreateCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) error
	UpdateCiliumNetworkPolicy(old, new *cilium_io_v2.CiliumNetworkPolicy) error
	DeleteCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) error
	GenericCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) error
}

type CiliumNetworkPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *cilium_io_v2.CiliumNetworkPolicy) error
	OnUpdate  func(old, new *cilium_io_v2.CiliumNetworkPolicy) error
	OnDelete  func(obj *cilium_io_v2.CiliumNetworkPolicy) error
	OnGeneric func(obj *cilium_io_v2.CiliumNetworkPolicy) error
}

func (f *CiliumNetworkPolicyEventHandlerFuncs) CreateCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *CiliumNetworkPolicyEventHandlerFuncs) DeleteCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *CiliumNetworkPolicyEventHandlerFuncs) UpdateCiliumNetworkPolicy(objOld, objNew *cilium_io_v2.CiliumNetworkPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *CiliumNetworkPolicyEventHandlerFuncs) GenericCiliumNetworkPolicy(obj *cilium_io_v2.CiliumNetworkPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type CiliumNetworkPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h CiliumNetworkPolicyEventHandler, predicates ...predicate.Predicate) error
}

type ciliumNetworkPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewCiliumNetworkPolicyEventWatcher(name string, mgr manager.Manager) CiliumNetworkPolicyEventWatcher {
	return &ciliumNetworkPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &cilium_io_v2.CiliumNetworkPolicy{}),
	}
}

func (c *ciliumNetworkPolicyEventWatcher) AddEventHandler(ctx context.Context, h CiliumNetworkPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericCiliumNetworkPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericCiliumNetworkPolicyHandler implements a generic events.EventHandler
type genericCiliumNetworkPolicyHandler struct {
	handler CiliumNetworkPolicyEventHandler
}

func (h genericCiliumNetworkPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*cilium_io_v2.CiliumNetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: CiliumNetworkPolicy handler received event for %T", object)
	}
	return h.handler.CreateCiliumNetworkPolicy(obj)
}

func (h genericCiliumNetworkPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*cilium_io_v2.CiliumNetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: CiliumNetworkPolicy handler received event for %T", object)
	}
	return h.handler.DeleteCiliumNetworkPolicy(obj)
}

func (h genericCiliumNetworkPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*cilium_io_v2.CiliumNetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: CiliumNetworkPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*cilium_io_v2.CiliumNetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: CiliumNetworkPolicy handler received event for %T", new)
	}
	return h.handler.UpdateCiliumNetworkPolicy(objOld, objNew)
}

func (h genericCiliumNetworkPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*cilium_io_v2.CiliumNetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: CiliumNetworkPolicy handler received event for %T", object)
	}
	return h.handler.GenericCiliumNetworkPolicy(obj)
}
