// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	rbac_authorization_k8s_io_v1 "k8s.io/api/rbac/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the Role Resource
// DEPRECATED: Prefer reconciler pattern.
type RoleEventHandler interface {
	CreateRole(obj *rbac_authorization_k8s_io_v1.Role) error
	UpdateRole(old, new *rbac_authorization_k8s_io_v1.Role) error
	DeleteRole(obj *rbac_authorization_k8s_io_v1.Role) error
	GenericRole(obj *rbac_authorization_k8s_io_v1.Role) error
}

type RoleEventHandlerFuncs struct {
	OnCreate  func(obj *rbac_authorization_k8s_io_v1.Role) error
	OnUpdate  func(old, new *rbac_authorization_k8s_io_v1.Role) error
	OnDelete  func(obj *rbac_authorization_k8s_io_v1.Role) error
	OnGeneric func(obj *rbac_authorization_k8s_io_v1.Role) error
}

func (f *RoleEventHandlerFuncs) CreateRole(obj *rbac_authorization_k8s_io_v1.Role) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RoleEventHandlerFuncs) DeleteRole(obj *rbac_authorization_k8s_io_v1.Role) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RoleEventHandlerFuncs) UpdateRole(objOld, objNew *rbac_authorization_k8s_io_v1.Role) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RoleEventHandlerFuncs) GenericRole(obj *rbac_authorization_k8s_io_v1.Role) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RoleEventWatcher interface {
	AddEventHandler(ctx context.Context, h RoleEventHandler, predicates ...predicate.Predicate) error
}

type roleEventWatcher struct {
	watcher events.EventWatcher
}

func NewRoleEventWatcher(name string, mgr manager.Manager) RoleEventWatcher {
	return &roleEventWatcher{
		watcher: events.NewWatcher(name, mgr, &rbac_authorization_k8s_io_v1.Role{}),
	}
}

func (c *roleEventWatcher) AddEventHandler(ctx context.Context, h RoleEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRoleHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRoleHandler implements a generic events.EventHandler
type genericRoleHandler struct {
	handler RoleEventHandler
}

func (h genericRoleHandler) Create(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.Role)
	if !ok {
		return errors.Errorf("internal error: Role handler received event for %T", object)
	}
	return h.handler.CreateRole(obj)
}

func (h genericRoleHandler) Delete(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.Role)
	if !ok {
		return errors.Errorf("internal error: Role handler received event for %T", object)
	}
	return h.handler.DeleteRole(obj)
}

func (h genericRoleHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*rbac_authorization_k8s_io_v1.Role)
	if !ok {
		return errors.Errorf("internal error: Role handler received event for %T", old)
	}
	objNew, ok := new.(*rbac_authorization_k8s_io_v1.Role)
	if !ok {
		return errors.Errorf("internal error: Role handler received event for %T", new)
	}
	return h.handler.UpdateRole(objOld, objNew)
}

func (h genericRoleHandler) Generic(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.Role)
	if !ok {
		return errors.Errorf("internal error: Role handler received event for %T", object)
	}
	return h.handler.GenericRole(obj)
}

// Handle events for the RoleBinding Resource
// DEPRECATED: Prefer reconciler pattern.
type RoleBindingEventHandler interface {
	CreateRoleBinding(obj *rbac_authorization_k8s_io_v1.RoleBinding) error
	UpdateRoleBinding(old, new *rbac_authorization_k8s_io_v1.RoleBinding) error
	DeleteRoleBinding(obj *rbac_authorization_k8s_io_v1.RoleBinding) error
	GenericRoleBinding(obj *rbac_authorization_k8s_io_v1.RoleBinding) error
}

type RoleBindingEventHandlerFuncs struct {
	OnCreate  func(obj *rbac_authorization_k8s_io_v1.RoleBinding) error
	OnUpdate  func(old, new *rbac_authorization_k8s_io_v1.RoleBinding) error
	OnDelete  func(obj *rbac_authorization_k8s_io_v1.RoleBinding) error
	OnGeneric func(obj *rbac_authorization_k8s_io_v1.RoleBinding) error
}

func (f *RoleBindingEventHandlerFuncs) CreateRoleBinding(obj *rbac_authorization_k8s_io_v1.RoleBinding) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *RoleBindingEventHandlerFuncs) DeleteRoleBinding(obj *rbac_authorization_k8s_io_v1.RoleBinding) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *RoleBindingEventHandlerFuncs) UpdateRoleBinding(objOld, objNew *rbac_authorization_k8s_io_v1.RoleBinding) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *RoleBindingEventHandlerFuncs) GenericRoleBinding(obj *rbac_authorization_k8s_io_v1.RoleBinding) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type RoleBindingEventWatcher interface {
	AddEventHandler(ctx context.Context, h RoleBindingEventHandler, predicates ...predicate.Predicate) error
}

type roleBindingEventWatcher struct {
	watcher events.EventWatcher
}

func NewRoleBindingEventWatcher(name string, mgr manager.Manager) RoleBindingEventWatcher {
	return &roleBindingEventWatcher{
		watcher: events.NewWatcher(name, mgr, &rbac_authorization_k8s_io_v1.RoleBinding{}),
	}
}

func (c *roleBindingEventWatcher) AddEventHandler(ctx context.Context, h RoleBindingEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRoleBindingHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericRoleBindingHandler implements a generic events.EventHandler
type genericRoleBindingHandler struct {
	handler RoleBindingEventHandler
}

func (h genericRoleBindingHandler) Create(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.RoleBinding)
	if !ok {
		return errors.Errorf("internal error: RoleBinding handler received event for %T", object)
	}
	return h.handler.CreateRoleBinding(obj)
}

func (h genericRoleBindingHandler) Delete(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.RoleBinding)
	if !ok {
		return errors.Errorf("internal error: RoleBinding handler received event for %T", object)
	}
	return h.handler.DeleteRoleBinding(obj)
}

func (h genericRoleBindingHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*rbac_authorization_k8s_io_v1.RoleBinding)
	if !ok {
		return errors.Errorf("internal error: RoleBinding handler received event for %T", old)
	}
	objNew, ok := new.(*rbac_authorization_k8s_io_v1.RoleBinding)
	if !ok {
		return errors.Errorf("internal error: RoleBinding handler received event for %T", new)
	}
	return h.handler.UpdateRoleBinding(objOld, objNew)
}

func (h genericRoleBindingHandler) Generic(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.RoleBinding)
	if !ok {
		return errors.Errorf("internal error: RoleBinding handler received event for %T", object)
	}
	return h.handler.GenericRoleBinding(obj)
}

// Handle events for the ClusterRole Resource
// DEPRECATED: Prefer reconciler pattern.
type ClusterRoleEventHandler interface {
	CreateClusterRole(obj *rbac_authorization_k8s_io_v1.ClusterRole) error
	UpdateClusterRole(old, new *rbac_authorization_k8s_io_v1.ClusterRole) error
	DeleteClusterRole(obj *rbac_authorization_k8s_io_v1.ClusterRole) error
	GenericClusterRole(obj *rbac_authorization_k8s_io_v1.ClusterRole) error
}

type ClusterRoleEventHandlerFuncs struct {
	OnCreate  func(obj *rbac_authorization_k8s_io_v1.ClusterRole) error
	OnUpdate  func(old, new *rbac_authorization_k8s_io_v1.ClusterRole) error
	OnDelete  func(obj *rbac_authorization_k8s_io_v1.ClusterRole) error
	OnGeneric func(obj *rbac_authorization_k8s_io_v1.ClusterRole) error
}

func (f *ClusterRoleEventHandlerFuncs) CreateClusterRole(obj *rbac_authorization_k8s_io_v1.ClusterRole) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ClusterRoleEventHandlerFuncs) DeleteClusterRole(obj *rbac_authorization_k8s_io_v1.ClusterRole) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ClusterRoleEventHandlerFuncs) UpdateClusterRole(objOld, objNew *rbac_authorization_k8s_io_v1.ClusterRole) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ClusterRoleEventHandlerFuncs) GenericClusterRole(obj *rbac_authorization_k8s_io_v1.ClusterRole) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ClusterRoleEventWatcher interface {
	AddEventHandler(ctx context.Context, h ClusterRoleEventHandler, predicates ...predicate.Predicate) error
}

type clusterRoleEventWatcher struct {
	watcher events.EventWatcher
}

func NewClusterRoleEventWatcher(name string, mgr manager.Manager) ClusterRoleEventWatcher {
	return &clusterRoleEventWatcher{
		watcher: events.NewWatcher(name, mgr, &rbac_authorization_k8s_io_v1.ClusterRole{}),
	}
}

func (c *clusterRoleEventWatcher) AddEventHandler(ctx context.Context, h ClusterRoleEventHandler, predicates ...predicate.Predicate) error {
	handler := genericClusterRoleHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericClusterRoleHandler implements a generic events.EventHandler
type genericClusterRoleHandler struct {
	handler ClusterRoleEventHandler
}

func (h genericClusterRoleHandler) Create(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.ClusterRole)
	if !ok {
		return errors.Errorf("internal error: ClusterRole handler received event for %T", object)
	}
	return h.handler.CreateClusterRole(obj)
}

func (h genericClusterRoleHandler) Delete(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.ClusterRole)
	if !ok {
		return errors.Errorf("internal error: ClusterRole handler received event for %T", object)
	}
	return h.handler.DeleteClusterRole(obj)
}

func (h genericClusterRoleHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*rbac_authorization_k8s_io_v1.ClusterRole)
	if !ok {
		return errors.Errorf("internal error: ClusterRole handler received event for %T", old)
	}
	objNew, ok := new.(*rbac_authorization_k8s_io_v1.ClusterRole)
	if !ok {
		return errors.Errorf("internal error: ClusterRole handler received event for %T", new)
	}
	return h.handler.UpdateClusterRole(objOld, objNew)
}

func (h genericClusterRoleHandler) Generic(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.ClusterRole)
	if !ok {
		return errors.Errorf("internal error: ClusterRole handler received event for %T", object)
	}
	return h.handler.GenericClusterRole(obj)
}

// Handle events for the ClusterRoleBinding Resource
// DEPRECATED: Prefer reconciler pattern.
type ClusterRoleBindingEventHandler interface {
	CreateClusterRoleBinding(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error
	UpdateClusterRoleBinding(old, new *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error
	DeleteClusterRoleBinding(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error
	GenericClusterRoleBinding(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error
}

type ClusterRoleBindingEventHandlerFuncs struct {
	OnCreate  func(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error
	OnUpdate  func(old, new *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error
	OnDelete  func(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error
	OnGeneric func(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error
}

func (f *ClusterRoleBindingEventHandlerFuncs) CreateClusterRoleBinding(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ClusterRoleBindingEventHandlerFuncs) DeleteClusterRoleBinding(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ClusterRoleBindingEventHandlerFuncs) UpdateClusterRoleBinding(objOld, objNew *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ClusterRoleBindingEventHandlerFuncs) GenericClusterRoleBinding(obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ClusterRoleBindingEventWatcher interface {
	AddEventHandler(ctx context.Context, h ClusterRoleBindingEventHandler, predicates ...predicate.Predicate) error
}

type clusterRoleBindingEventWatcher struct {
	watcher events.EventWatcher
}

func NewClusterRoleBindingEventWatcher(name string, mgr manager.Manager) ClusterRoleBindingEventWatcher {
	return &clusterRoleBindingEventWatcher{
		watcher: events.NewWatcher(name, mgr, &rbac_authorization_k8s_io_v1.ClusterRoleBinding{}),
	}
}

func (c *clusterRoleBindingEventWatcher) AddEventHandler(ctx context.Context, h ClusterRoleBindingEventHandler, predicates ...predicate.Predicate) error {
	handler := genericClusterRoleBindingHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericClusterRoleBindingHandler implements a generic events.EventHandler
type genericClusterRoleBindingHandler struct {
	handler ClusterRoleBindingEventHandler
}

func (h genericClusterRoleBindingHandler) Create(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding)
	if !ok {
		return errors.Errorf("internal error: ClusterRoleBinding handler received event for %T", object)
	}
	return h.handler.CreateClusterRoleBinding(obj)
}

func (h genericClusterRoleBindingHandler) Delete(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding)
	if !ok {
		return errors.Errorf("internal error: ClusterRoleBinding handler received event for %T", object)
	}
	return h.handler.DeleteClusterRoleBinding(obj)
}

func (h genericClusterRoleBindingHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding)
	if !ok {
		return errors.Errorf("internal error: ClusterRoleBinding handler received event for %T", old)
	}
	objNew, ok := new.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding)
	if !ok {
		return errors.Errorf("internal error: ClusterRoleBinding handler received event for %T", new)
	}
	return h.handler.UpdateClusterRoleBinding(objOld, objNew)
}

func (h genericClusterRoleBindingHandler) Generic(object client.Object) error {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding)
	if !ok {
		return errors.Errorf("internal error: ClusterRoleBinding handler received event for %T", object)
	}
	return h.handler.GenericClusterRoleBinding(obj)
}
