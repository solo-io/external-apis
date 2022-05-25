// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_k8s_io_v1 "k8s.io/api/networking/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the NetworkPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type NetworkPolicyEventHandler interface {
	CreateNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) error
	UpdateNetworkPolicy(old, new *networking_k8s_io_v1.NetworkPolicy) error
	DeleteNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) error
	GenericNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) error
}

type NetworkPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *networking_k8s_io_v1.NetworkPolicy) error
	OnUpdate  func(old, new *networking_k8s_io_v1.NetworkPolicy) error
	OnDelete  func(obj *networking_k8s_io_v1.NetworkPolicy) error
	OnGeneric func(obj *networking_k8s_io_v1.NetworkPolicy) error
}

func (f *NetworkPolicyEventHandlerFuncs) CreateNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *NetworkPolicyEventHandlerFuncs) DeleteNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *NetworkPolicyEventHandlerFuncs) UpdateNetworkPolicy(objOld, objNew *networking_k8s_io_v1.NetworkPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *NetworkPolicyEventHandlerFuncs) GenericNetworkPolicy(obj *networking_k8s_io_v1.NetworkPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type NetworkPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h NetworkPolicyEventHandler, predicates ...predicate.Predicate) error
}

type networkPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewNetworkPolicyEventWatcher(name string, mgr manager.Manager) NetworkPolicyEventWatcher {
	return &networkPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_k8s_io_v1.NetworkPolicy{}),
	}
}

func (c *networkPolicyEventWatcher) AddEventHandler(ctx context.Context, h NetworkPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericNetworkPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericNetworkPolicyHandler implements a generic events.EventHandler
type genericNetworkPolicyHandler struct {
	handler NetworkPolicyEventHandler
}

func (h genericNetworkPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*networking_k8s_io_v1.NetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: NetworkPolicy handler received event for %T", object)
	}
	return h.handler.CreateNetworkPolicy(obj)
}

func (h genericNetworkPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*networking_k8s_io_v1.NetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: NetworkPolicy handler received event for %T", object)
	}
	return h.handler.DeleteNetworkPolicy(obj)
}

func (h genericNetworkPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*networking_k8s_io_v1.NetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: NetworkPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*networking_k8s_io_v1.NetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: NetworkPolicy handler received event for %T", new)
	}
	return h.handler.UpdateNetworkPolicy(objOld, objNew)
}

func (h genericNetworkPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*networking_k8s_io_v1.NetworkPolicy)
	if !ok {
		return errors.Errorf("internal error: NetworkPolicy handler received event for %T", object)
	}
	return h.handler.GenericNetworkPolicy(obj)
}
