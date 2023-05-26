// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	security_istio_io_v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the AuthorizationPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type AuthorizationPolicyEventHandler interface {
	CreateAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) error
	UpdateAuthorizationPolicy(old, new *security_istio_io_v1beta1.AuthorizationPolicy) error
	DeleteAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) error
	GenericAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) error
}

type AuthorizationPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *security_istio_io_v1beta1.AuthorizationPolicy) error
	OnUpdate  func(old, new *security_istio_io_v1beta1.AuthorizationPolicy) error
	OnDelete  func(obj *security_istio_io_v1beta1.AuthorizationPolicy) error
	OnGeneric func(obj *security_istio_io_v1beta1.AuthorizationPolicy) error
}

func (f *AuthorizationPolicyEventHandlerFuncs) CreateAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *AuthorizationPolicyEventHandlerFuncs) DeleteAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *AuthorizationPolicyEventHandlerFuncs) UpdateAuthorizationPolicy(objOld, objNew *security_istio_io_v1beta1.AuthorizationPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *AuthorizationPolicyEventHandlerFuncs) GenericAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type AuthorizationPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h AuthorizationPolicyEventHandler, predicates ...predicate.Predicate) error
}

type authorizationPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewAuthorizationPolicyEventWatcher(name string, mgr manager.Manager) AuthorizationPolicyEventWatcher {
	return &authorizationPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_istio_io_v1beta1.AuthorizationPolicy{}),
	}
}

func (c *authorizationPolicyEventWatcher) AddEventHandler(ctx context.Context, h AuthorizationPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericAuthorizationPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericAuthorizationPolicyHandler implements a generic events.EventHandler
type genericAuthorizationPolicyHandler struct {
	handler AuthorizationPolicyEventHandler
}

func (h genericAuthorizationPolicyHandler) Create(object client.Object) error {
	obj, ok := object.(*security_istio_io_v1beta1.AuthorizationPolicy)
	if !ok {
		return errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", object)
	}
	return h.handler.CreateAuthorizationPolicy(obj)
}

func (h genericAuthorizationPolicyHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_istio_io_v1beta1.AuthorizationPolicy)
	if !ok {
		return errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", object)
	}
	return h.handler.DeleteAuthorizationPolicy(obj)
}

func (h genericAuthorizationPolicyHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_istio_io_v1beta1.AuthorizationPolicy)
	if !ok {
		return errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*security_istio_io_v1beta1.AuthorizationPolicy)
	if !ok {
		return errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", new)
	}
	return h.handler.UpdateAuthorizationPolicy(objOld, objNew)
}

func (h genericAuthorizationPolicyHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_istio_io_v1beta1.AuthorizationPolicy)
	if !ok {
		return errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", object)
	}
	return h.handler.GenericAuthorizationPolicy(obj)
}

// Handle events for the PeerAuthentication Resource
// DEPRECATED: Prefer reconciler pattern.
type PeerAuthenticationEventHandler interface {
	CreatePeerAuthentication(obj *security_istio_io_v1beta1.PeerAuthentication) error
	UpdatePeerAuthentication(old, new *security_istio_io_v1beta1.PeerAuthentication) error
	DeletePeerAuthentication(obj *security_istio_io_v1beta1.PeerAuthentication) error
	GenericPeerAuthentication(obj *security_istio_io_v1beta1.PeerAuthentication) error
}

type PeerAuthenticationEventHandlerFuncs struct {
	OnCreate  func(obj *security_istio_io_v1beta1.PeerAuthentication) error
	OnUpdate  func(old, new *security_istio_io_v1beta1.PeerAuthentication) error
	OnDelete  func(obj *security_istio_io_v1beta1.PeerAuthentication) error
	OnGeneric func(obj *security_istio_io_v1beta1.PeerAuthentication) error
}

func (f *PeerAuthenticationEventHandlerFuncs) CreatePeerAuthentication(obj *security_istio_io_v1beta1.PeerAuthentication) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *PeerAuthenticationEventHandlerFuncs) DeletePeerAuthentication(obj *security_istio_io_v1beta1.PeerAuthentication) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *PeerAuthenticationEventHandlerFuncs) UpdatePeerAuthentication(objOld, objNew *security_istio_io_v1beta1.PeerAuthentication) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *PeerAuthenticationEventHandlerFuncs) GenericPeerAuthentication(obj *security_istio_io_v1beta1.PeerAuthentication) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type PeerAuthenticationEventWatcher interface {
	AddEventHandler(ctx context.Context, h PeerAuthenticationEventHandler, predicates ...predicate.Predicate) error
}

type peerAuthenticationEventWatcher struct {
	watcher events.EventWatcher
}

func NewPeerAuthenticationEventWatcher(name string, mgr manager.Manager) PeerAuthenticationEventWatcher {
	return &peerAuthenticationEventWatcher{
		watcher: events.NewWatcher(name, mgr, &security_istio_io_v1beta1.PeerAuthentication{}),
	}
}

func (c *peerAuthenticationEventWatcher) AddEventHandler(ctx context.Context, h PeerAuthenticationEventHandler, predicates ...predicate.Predicate) error {
	handler := genericPeerAuthenticationHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericPeerAuthenticationHandler implements a generic events.EventHandler
type genericPeerAuthenticationHandler struct {
	handler PeerAuthenticationEventHandler
}

func (h genericPeerAuthenticationHandler) Create(object client.Object) error {
	obj, ok := object.(*security_istio_io_v1beta1.PeerAuthentication)
	if !ok {
		return errors.Errorf("internal error: PeerAuthentication handler received event for %T", object)
	}
	return h.handler.CreatePeerAuthentication(obj)
}

func (h genericPeerAuthenticationHandler) Delete(object client.Object) error {
	obj, ok := object.(*security_istio_io_v1beta1.PeerAuthentication)
	if !ok {
		return errors.Errorf("internal error: PeerAuthentication handler received event for %T", object)
	}
	return h.handler.DeletePeerAuthentication(obj)
}

func (h genericPeerAuthenticationHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*security_istio_io_v1beta1.PeerAuthentication)
	if !ok {
		return errors.Errorf("internal error: PeerAuthentication handler received event for %T", old)
	}
	objNew, ok := new.(*security_istio_io_v1beta1.PeerAuthentication)
	if !ok {
		return errors.Errorf("internal error: PeerAuthentication handler received event for %T", new)
	}
	return h.handler.UpdatePeerAuthentication(objOld, objNew)
}

func (h genericPeerAuthenticationHandler) Generic(object client.Object) error {
	obj, ok := object.(*security_istio_io_v1beta1.PeerAuthentication)
	if !ok {
		return errors.Errorf("internal error: PeerAuthentication handler received event for %T", object)
	}
	return h.handler.GenericPeerAuthentication(obj)
}