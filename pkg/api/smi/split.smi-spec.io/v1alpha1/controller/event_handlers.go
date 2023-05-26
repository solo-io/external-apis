// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	split_smi_spec_io_v1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the TrafficSplit Resource
// DEPRECATED: Prefer reconciler pattern.
type TrafficSplitEventHandler interface {
	CreateTrafficSplit(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error
	UpdateTrafficSplit(old, new *split_smi_spec_io_v1alpha1.TrafficSplit) error
	DeleteTrafficSplit(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error
	GenericTrafficSplit(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error
}

type TrafficSplitEventHandlerFuncs struct {
	OnCreate  func(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error
	OnUpdate  func(old, new *split_smi_spec_io_v1alpha1.TrafficSplit) error
	OnDelete  func(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error
	OnGeneric func(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error
}

func (f *TrafficSplitEventHandlerFuncs) CreateTrafficSplit(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *TrafficSplitEventHandlerFuncs) DeleteTrafficSplit(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *TrafficSplitEventHandlerFuncs) UpdateTrafficSplit(objOld, objNew *split_smi_spec_io_v1alpha1.TrafficSplit) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *TrafficSplitEventHandlerFuncs) GenericTrafficSplit(obj *split_smi_spec_io_v1alpha1.TrafficSplit) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type TrafficSplitEventWatcher interface {
	AddEventHandler(ctx context.Context, h TrafficSplitEventHandler, predicates ...predicate.Predicate) error
}

type trafficSplitEventWatcher struct {
	watcher events.EventWatcher
}

func NewTrafficSplitEventWatcher(name string, mgr manager.Manager) TrafficSplitEventWatcher {
	return &trafficSplitEventWatcher{
		watcher: events.NewWatcher(name, mgr, &split_smi_spec_io_v1alpha1.TrafficSplit{}),
	}
}

func (c *trafficSplitEventWatcher) AddEventHandler(ctx context.Context, h TrafficSplitEventHandler, predicates ...predicate.Predicate) error {
	handler := genericTrafficSplitHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericTrafficSplitHandler implements a generic events.EventHandler
type genericTrafficSplitHandler struct {
	handler TrafficSplitEventHandler
}

func (h genericTrafficSplitHandler) Create(object client.Object) error {
	obj, ok := object.(*split_smi_spec_io_v1alpha1.TrafficSplit)
	if !ok {
		return errors.Errorf("internal error: TrafficSplit handler received event for %T", object)
	}
	return h.handler.CreateTrafficSplit(obj)
}

func (h genericTrafficSplitHandler) Delete(object client.Object) error {
	obj, ok := object.(*split_smi_spec_io_v1alpha1.TrafficSplit)
	if !ok {
		return errors.Errorf("internal error: TrafficSplit handler received event for %T", object)
	}
	return h.handler.DeleteTrafficSplit(obj)
}

func (h genericTrafficSplitHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*split_smi_spec_io_v1alpha1.TrafficSplit)
	if !ok {
		return errors.Errorf("internal error: TrafficSplit handler received event for %T", old)
	}
	objNew, ok := new.(*split_smi_spec_io_v1alpha1.TrafficSplit)
	if !ok {
		return errors.Errorf("internal error: TrafficSplit handler received event for %T", new)
	}
	return h.handler.UpdateTrafficSplit(objOld, objNew)
}

func (h genericTrafficSplitHandler) Generic(object client.Object) error {
	obj, ok := object.(*split_smi_spec_io_v1alpha1.TrafficSplit)
	if !ok {
		return errors.Errorf("internal error: TrafficSplit handler received event for %T", object)
	}
	return h.handler.GenericTrafficSplit(obj)
}