// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	extensions_istio_io_v1alpha1 "istio.io/client-go/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the extensions.istio.io/v1alpha1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the extensions.istio.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the extensions.istio.io/v1alpha1/v1alpha1 APIs
	WasmPlugins() WasmPluginClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := extensions_istio_io_v1alpha1.SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the extensions.istio.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) WasmPlugins() WasmPluginClient {
	return NewWasmPluginClient(c.client)
}

// Reader knows how to read and list WasmPlugins.
type WasmPluginReader interface {
	// Get retrieves a WasmPlugin for the given object key
	GetWasmPlugin(ctx context.Context, key client.ObjectKey) (*extensions_istio_io_v1alpha1.WasmPlugin, error)

	// List retrieves list of WasmPlugins for a given namespace and list options.
	ListWasmPlugin(ctx context.Context, opts ...client.ListOption) (*extensions_istio_io_v1alpha1.WasmPluginList, error)
}

// WasmPluginTransitionFunction instructs the WasmPluginWriter how to transition between an existing
// WasmPlugin object and a desired on an Upsert
type WasmPluginTransitionFunction func(existing, desired *extensions_istio_io_v1alpha1.WasmPlugin) error

// Writer knows how to create, delete, and update WasmPlugins.
type WasmPluginWriter interface {
	// Create saves the WasmPlugin object.
	CreateWasmPlugin(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, opts ...client.CreateOption) error

	// Delete deletes the WasmPlugin object.
	DeleteWasmPlugin(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given WasmPlugin object.
	UpdateWasmPlugin(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, opts ...client.UpdateOption) error

	// Patch patches the given WasmPlugin object.
	PatchWasmPlugin(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all WasmPlugin objects matching the given options.
	DeleteAllOfWasmPlugin(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the WasmPlugin object.
	UpsertWasmPlugin(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, transitionFuncs ...WasmPluginTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a WasmPlugin object.
type WasmPluginStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given WasmPlugin object.
	UpdateWasmPluginStatus(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given WasmPlugin object's subresource.
	PatchWasmPluginStatus(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on WasmPlugins.
type WasmPluginClient interface {
	WasmPluginReader
	WasmPluginWriter
	WasmPluginStatusWriter
}

type wasmPluginClient struct {
	client client.Client
}

func NewWasmPluginClient(client client.Client) *wasmPluginClient {
	return &wasmPluginClient{client: client}
}

func (c *wasmPluginClient) GetWasmPlugin(ctx context.Context, key client.ObjectKey) (*extensions_istio_io_v1alpha1.WasmPlugin, error) {
	obj := &extensions_istio_io_v1alpha1.WasmPlugin{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *wasmPluginClient) ListWasmPlugin(ctx context.Context, opts ...client.ListOption) (*extensions_istio_io_v1alpha1.WasmPluginList, error) {
	list := &extensions_istio_io_v1alpha1.WasmPluginList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *wasmPluginClient) CreateWasmPlugin(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *wasmPluginClient) DeleteWasmPlugin(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &extensions_istio_io_v1alpha1.WasmPlugin{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *wasmPluginClient) UpdateWasmPlugin(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *wasmPluginClient) PatchWasmPlugin(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *wasmPluginClient) DeleteAllOfWasmPlugin(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &extensions_istio_io_v1alpha1.WasmPlugin{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *wasmPluginClient) UpsertWasmPlugin(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, transitionFuncs ...WasmPluginTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*extensions_istio_io_v1alpha1.WasmPlugin), desired.(*extensions_istio_io_v1alpha1.WasmPlugin)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *wasmPluginClient) UpdateWasmPluginStatus(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *wasmPluginClient) PatchWasmPluginStatus(ctx context.Context, obj *extensions_istio_io_v1alpha1.WasmPlugin, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides WasmPluginClients for multiple clusters.
type MulticlusterWasmPluginClient interface {
	// Cluster returns a WasmPluginClient for the given cluster
	Cluster(cluster string) (WasmPluginClient, error)
}

type multiclusterWasmPluginClient struct {
	client multicluster.Client
}

func NewMulticlusterWasmPluginClient(client multicluster.Client) MulticlusterWasmPluginClient {
	return &multiclusterWasmPluginClient{client: client}
}

func (m *multiclusterWasmPluginClient) Cluster(cluster string) (WasmPluginClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewWasmPluginClient(client), nil
}
