// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha3

import (
	"context"

	specs_smi_spec_io_v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the specs.smi-spec.io/v1alpha3 APIs
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

// clienset for the specs.smi-spec.io/v1alpha3 APIs
type Clientset interface {
	// clienset for the specs.smi-spec.io/v1alpha3/v1alpha3 APIs
	HTTPRouteGroups() HTTPRouteGroupClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := specs_smi_spec_io_v1alpha3.AddToScheme(scheme); err != nil {
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

// clienset for the specs.smi-spec.io/v1alpha3/v1alpha3 APIs
func (c *clientSet) HTTPRouteGroups() HTTPRouteGroupClient {
	return NewHTTPRouteGroupClient(c.client)
}

// Reader knows how to read and list HTTPRouteGroups.
type HTTPRouteGroupReader interface {
	// Get retrieves a HTTPRouteGroup for the given object key
	GetHTTPRouteGroup(ctx context.Context, key client.ObjectKey) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, error)

	// List retrieves list of HTTPRouteGroups for a given namespace and list options.
	ListHTTPRouteGroup(ctx context.Context, opts ...client.ListOption) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroupList, error)
}

// HTTPRouteGroupTransitionFunction instructs the HTTPRouteGroupWriter how to transition between an existing
// HTTPRouteGroup object and a desired on an Upsert
type HTTPRouteGroupTransitionFunction func(existing, desired *specs_smi_spec_io_v1alpha3.HTTPRouteGroup) error

// Writer knows how to create, delete, and update HTTPRouteGroups.
type HTTPRouteGroupWriter interface {
	// Create saves the HTTPRouteGroup object.
	CreateHTTPRouteGroup(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, opts ...client.CreateOption) error

	// Delete deletes the HTTPRouteGroup object.
	DeleteHTTPRouteGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given HTTPRouteGroup object.
	UpdateHTTPRouteGroup(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, opts ...client.UpdateOption) error

	// Patch patches the given HTTPRouteGroup object.
	PatchHTTPRouteGroup(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all HTTPRouteGroup objects matching the given options.
	DeleteAllOfHTTPRouteGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the HTTPRouteGroup object.
	UpsertHTTPRouteGroup(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, transitionFuncs ...HTTPRouteGroupTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a HTTPRouteGroup object.
type HTTPRouteGroupStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given HTTPRouteGroup object.
	UpdateHTTPRouteGroupStatus(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, opts ...client.UpdateOption) error

	// Patch patches the given HTTPRouteGroup object's subresource.
	PatchHTTPRouteGroupStatus(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on HTTPRouteGroups.
type HTTPRouteGroupClient interface {
	HTTPRouteGroupReader
	HTTPRouteGroupWriter
	HTTPRouteGroupStatusWriter
}

type hTTPRouteGroupClient struct {
	client client.Client
}

func NewHTTPRouteGroupClient(client client.Client) *hTTPRouteGroupClient {
	return &hTTPRouteGroupClient{client: client}
}

func (c *hTTPRouteGroupClient) GetHTTPRouteGroup(ctx context.Context, key client.ObjectKey) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, error) {
	obj := &specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *hTTPRouteGroupClient) ListHTTPRouteGroup(ctx context.Context, opts ...client.ListOption) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroupList, error) {
	list := &specs_smi_spec_io_v1alpha3.HTTPRouteGroupList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *hTTPRouteGroupClient) CreateHTTPRouteGroup(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *hTTPRouteGroupClient) DeleteHTTPRouteGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *hTTPRouteGroupClient) UpdateHTTPRouteGroup(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *hTTPRouteGroupClient) PatchHTTPRouteGroup(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *hTTPRouteGroupClient) DeleteAllOfHTTPRouteGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *hTTPRouteGroupClient) UpsertHTTPRouteGroup(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, transitionFuncs ...HTTPRouteGroupTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup), desired.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *hTTPRouteGroupClient) UpdateHTTPRouteGroupStatus(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *hTTPRouteGroupClient) PatchHTTPRouteGroupStatus(ctx context.Context, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides HTTPRouteGroupClients for multiple clusters.
type MulticlusterHTTPRouteGroupClient interface {
	// Cluster returns a HTTPRouteGroupClient for the given cluster
	Cluster(cluster string) (HTTPRouteGroupClient, error)
}

type multiclusterHTTPRouteGroupClient struct {
	client multicluster.Client
}

func NewMulticlusterHTTPRouteGroupClient(client multicluster.Client) MulticlusterHTTPRouteGroupClient {
	return &multiclusterHTTPRouteGroupClient{client: client}
}

func (m *multiclusterHTTPRouteGroupClient) Cluster(cluster string) (HTTPRouteGroupClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewHTTPRouteGroupClient(client), nil
}
