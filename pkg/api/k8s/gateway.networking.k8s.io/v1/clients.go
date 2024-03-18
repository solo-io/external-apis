// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gateway_networking_k8s_io_v1 "sigs.k8s.io/gateway-api/apis/v1"
)

// MulticlusterClientset for the gateway.networking.k8s.io/v1 APIs
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

// clienset for the gateway.networking.k8s.io/v1 APIs
type Clientset interface {
	// clienset for the gateway.networking.k8s.io/v1/v1 APIs
	Gateways() GatewayClient
	// clienset for the gateway.networking.k8s.io/v1/v1 APIs
	GatewayClasses() GatewayClassClient
	// clienset for the gateway.networking.k8s.io/v1/v1 APIs
	HTTPRoutes() HTTPRouteClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := gateway_networking_k8s_io_v1.SchemeBuilder.AddToScheme(scheme); err != nil {
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

// clienset for the gateway.networking.k8s.io/v1/v1 APIs
func (c *clientSet) Gateways() GatewayClient {
	return NewGatewayClient(c.client)
}

// clienset for the gateway.networking.k8s.io/v1/v1 APIs
func (c *clientSet) GatewayClasses() GatewayClassClient {
	return NewGatewayClassClient(c.client)
}

// clienset for the gateway.networking.k8s.io/v1/v1 APIs
func (c *clientSet) HTTPRoutes() HTTPRouteClient {
	return NewHTTPRouteClient(c.client)
}

// Reader knows how to read and list Gateways.
type GatewayReader interface {
	// Get retrieves a Gateway for the given object key
	GetGateway(ctx context.Context, key client.ObjectKey) (*gateway_networking_k8s_io_v1.Gateway, error)

	// List retrieves list of Gateways for a given namespace and list options.
	ListGateway(ctx context.Context, opts ...client.ListOption) (*gateway_networking_k8s_io_v1.GatewayList, error)
}

// GatewayTransitionFunction instructs the GatewayWriter how to transition between an existing
// Gateway object and a desired on an Upsert
type GatewayTransitionFunction func(existing, desired *gateway_networking_k8s_io_v1.Gateway) error

// Writer knows how to create, delete, and update Gateways.
type GatewayWriter interface {
	// Create saves the Gateway object.
	CreateGateway(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, opts ...client.CreateOption) error

	// Delete deletes the Gateway object.
	DeleteGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Gateway object.
	UpdateGateway(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, opts ...client.UpdateOption) error

	// Patch patches the given Gateway object.
	PatchGateway(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Gateway objects matching the given options.
	DeleteAllOfGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Gateway object.
	UpsertGateway(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, transitionFuncs ...GatewayTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Gateway object.
type GatewayStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Gateway object.
	UpdateGatewayStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given Gateway object's subresource.
	PatchGatewayStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on Gateways.
type GatewayClient interface {
	GatewayReader
	GatewayWriter
	GatewayStatusWriter
}

type gatewayClient struct {
	client client.Client
}

func NewGatewayClient(client client.Client) *gatewayClient {
	return &gatewayClient{client: client}
}

func (c *gatewayClient) GetGateway(ctx context.Context, key client.ObjectKey) (*gateway_networking_k8s_io_v1.Gateway, error) {
	obj := &gateway_networking_k8s_io_v1.Gateway{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *gatewayClient) ListGateway(ctx context.Context, opts ...client.ListOption) (*gateway_networking_k8s_io_v1.GatewayList, error) {
	list := &gateway_networking_k8s_io_v1.GatewayList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *gatewayClient) CreateGateway(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *gatewayClient) DeleteGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &gateway_networking_k8s_io_v1.Gateway{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *gatewayClient) UpdateGateway(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *gatewayClient) PatchGateway(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *gatewayClient) DeleteAllOfGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &gateway_networking_k8s_io_v1.Gateway{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *gatewayClient) UpsertGateway(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, transitionFuncs ...GatewayTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*gateway_networking_k8s_io_v1.Gateway), desired.(*gateway_networking_k8s_io_v1.Gateway)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *gatewayClient) UpdateGatewayStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *gatewayClient) PatchGatewayStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.Gateway, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides GatewayClients for multiple clusters.
type MulticlusterGatewayClient interface {
	// Cluster returns a GatewayClient for the given cluster
	Cluster(cluster string) (GatewayClient, error)
}

type multiclusterGatewayClient struct {
	client multicluster.Client
}

func NewMulticlusterGatewayClient(client multicluster.Client) MulticlusterGatewayClient {
	return &multiclusterGatewayClient{client: client}
}

func (m *multiclusterGatewayClient) Cluster(cluster string) (GatewayClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewGatewayClient(client), nil
}

// Reader knows how to read and list GatewayClasss.
type GatewayClassReader interface {
	// Get retrieves a GatewayClass for the given object key
	GetGatewayClass(ctx context.Context, name string) (*gateway_networking_k8s_io_v1.GatewayClass, error)

	// List retrieves list of GatewayClasss for a given namespace and list options.
	ListGatewayClass(ctx context.Context, opts ...client.ListOption) (*gateway_networking_k8s_io_v1.GatewayClassList, error)
}

// GatewayClassTransitionFunction instructs the GatewayClassWriter how to transition between an existing
// GatewayClass object and a desired on an Upsert
type GatewayClassTransitionFunction func(existing, desired *gateway_networking_k8s_io_v1.GatewayClass) error

// Writer knows how to create, delete, and update GatewayClasss.
type GatewayClassWriter interface {
	// Create saves the GatewayClass object.
	CreateGatewayClass(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, opts ...client.CreateOption) error

	// Delete deletes the GatewayClass object.
	DeleteGatewayClass(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given GatewayClass object.
	UpdateGatewayClass(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, opts ...client.UpdateOption) error

	// Patch patches the given GatewayClass object.
	PatchGatewayClass(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all GatewayClass objects matching the given options.
	DeleteAllOfGatewayClass(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the GatewayClass object.
	UpsertGatewayClass(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, transitionFuncs ...GatewayClassTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a GatewayClass object.
type GatewayClassStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given GatewayClass object.
	UpdateGatewayClassStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given GatewayClass object's subresource.
	PatchGatewayClassStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on GatewayClasss.
type GatewayClassClient interface {
	GatewayClassReader
	GatewayClassWriter
	GatewayClassStatusWriter
}

type gatewayClassClient struct {
	client client.Client
}

func NewGatewayClassClient(client client.Client) *gatewayClassClient {
	return &gatewayClassClient{client: client}
}

func (c *gatewayClassClient) GetGatewayClass(ctx context.Context, name string) (*gateway_networking_k8s_io_v1.GatewayClass, error) {
	obj := &gateway_networking_k8s_io_v1.GatewayClass{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *gatewayClassClient) ListGatewayClass(ctx context.Context, opts ...client.ListOption) (*gateway_networking_k8s_io_v1.GatewayClassList, error) {
	list := &gateway_networking_k8s_io_v1.GatewayClassList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *gatewayClassClient) CreateGatewayClass(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *gatewayClassClient) DeleteGatewayClass(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &gateway_networking_k8s_io_v1.GatewayClass{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *gatewayClassClient) UpdateGatewayClass(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *gatewayClassClient) PatchGatewayClass(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *gatewayClassClient) DeleteAllOfGatewayClass(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &gateway_networking_k8s_io_v1.GatewayClass{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *gatewayClassClient) UpsertGatewayClass(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, transitionFuncs ...GatewayClassTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*gateway_networking_k8s_io_v1.GatewayClass), desired.(*gateway_networking_k8s_io_v1.GatewayClass)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *gatewayClassClient) UpdateGatewayClassStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *gatewayClassClient) PatchGatewayClassStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.GatewayClass, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides GatewayClassClients for multiple clusters.
type MulticlusterGatewayClassClient interface {
	// Cluster returns a GatewayClassClient for the given cluster
	Cluster(cluster string) (GatewayClassClient, error)
}

type multiclusterGatewayClassClient struct {
	client multicluster.Client
}

func NewMulticlusterGatewayClassClient(client multicluster.Client) MulticlusterGatewayClassClient {
	return &multiclusterGatewayClassClient{client: client}
}

func (m *multiclusterGatewayClassClient) Cluster(cluster string) (GatewayClassClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewGatewayClassClient(client), nil
}

// Reader knows how to read and list HTTPRoutes.
type HTTPRouteReader interface {
	// Get retrieves a HTTPRoute for the given object key
	GetHTTPRoute(ctx context.Context, key client.ObjectKey) (*gateway_networking_k8s_io_v1.HTTPRoute, error)

	// List retrieves list of HTTPRoutes for a given namespace and list options.
	ListHTTPRoute(ctx context.Context, opts ...client.ListOption) (*gateway_networking_k8s_io_v1.HTTPRouteList, error)
}

// HTTPRouteTransitionFunction instructs the HTTPRouteWriter how to transition between an existing
// HTTPRoute object and a desired on an Upsert
type HTTPRouteTransitionFunction func(existing, desired *gateway_networking_k8s_io_v1.HTTPRoute) error

// Writer knows how to create, delete, and update HTTPRoutes.
type HTTPRouteWriter interface {
	// Create saves the HTTPRoute object.
	CreateHTTPRoute(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, opts ...client.CreateOption) error

	// Delete deletes the HTTPRoute object.
	DeleteHTTPRoute(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given HTTPRoute object.
	UpdateHTTPRoute(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, opts ...client.UpdateOption) error

	// Patch patches the given HTTPRoute object.
	PatchHTTPRoute(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all HTTPRoute objects matching the given options.
	DeleteAllOfHTTPRoute(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the HTTPRoute object.
	UpsertHTTPRoute(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, transitionFuncs ...HTTPRouteTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a HTTPRoute object.
type HTTPRouteStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given HTTPRoute object.
	UpdateHTTPRouteStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given HTTPRoute object's subresource.
	PatchHTTPRouteStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on HTTPRoutes.
type HTTPRouteClient interface {
	HTTPRouteReader
	HTTPRouteWriter
	HTTPRouteStatusWriter
}

type hTTPRouteClient struct {
	client client.Client
}

func NewHTTPRouteClient(client client.Client) *hTTPRouteClient {
	return &hTTPRouteClient{client: client}
}

func (c *hTTPRouteClient) GetHTTPRoute(ctx context.Context, key client.ObjectKey) (*gateway_networking_k8s_io_v1.HTTPRoute, error) {
	obj := &gateway_networking_k8s_io_v1.HTTPRoute{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *hTTPRouteClient) ListHTTPRoute(ctx context.Context, opts ...client.ListOption) (*gateway_networking_k8s_io_v1.HTTPRouteList, error) {
	list := &gateway_networking_k8s_io_v1.HTTPRouteList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *hTTPRouteClient) CreateHTTPRoute(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *hTTPRouteClient) DeleteHTTPRoute(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &gateway_networking_k8s_io_v1.HTTPRoute{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *hTTPRouteClient) UpdateHTTPRoute(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *hTTPRouteClient) PatchHTTPRoute(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *hTTPRouteClient) DeleteAllOfHTTPRoute(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &gateway_networking_k8s_io_v1.HTTPRoute{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *hTTPRouteClient) UpsertHTTPRoute(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, transitionFuncs ...HTTPRouteTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*gateway_networking_k8s_io_v1.HTTPRoute), desired.(*gateway_networking_k8s_io_v1.HTTPRoute)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *hTTPRouteClient) UpdateHTTPRouteStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *hTTPRouteClient) PatchHTTPRouteStatus(ctx context.Context, obj *gateway_networking_k8s_io_v1.HTTPRoute, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides HTTPRouteClients for multiple clusters.
type MulticlusterHTTPRouteClient interface {
	// Cluster returns a HTTPRouteClient for the given cluster
	Cluster(cluster string) (HTTPRouteClient, error)
}

type multiclusterHTTPRouteClient struct {
	client multicluster.Client
}

func NewMulticlusterHTTPRouteClient(client multicluster.Client) MulticlusterHTTPRouteClient {
	return &multiclusterHTTPRouteClient{client: client}
}

func (m *multiclusterHTTPRouteClient) Cluster(cluster string) (HTTPRouteClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewHTTPRouteClient(client), nil
}