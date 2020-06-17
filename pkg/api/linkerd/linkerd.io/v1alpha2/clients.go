// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha2

import (
	"context"

	linkerd_io_v1alpha2 "github.com/linkerd/linkerd2/controller/gen/apis/serviceprofile/v1alpha2"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the linkerd.io/v1alpha2 APIs
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

// clienset for the linkerd.io/v1alpha2 APIs
type Clientset interface {
	// clienset for the linkerd.io/v1alpha2/v1alpha2 APIs
	ServiceProfiles() ServiceProfileClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := linkerd_io_v1alpha2.AddToScheme(scheme); err != nil {
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

// clienset for the linkerd.io/v1alpha2/v1alpha2 APIs
func (c *clientSet) ServiceProfiles() ServiceProfileClient {
	return NewServiceProfileClient(c.client)
}

// Reader knows how to read and list ServiceProfiles.
type ServiceProfileReader interface {
	// Get retrieves a ServiceProfile for the given object key
	GetServiceProfile(ctx context.Context, key client.ObjectKey) (*linkerd_io_v1alpha2.ServiceProfile, error)

	// List retrieves list of ServiceProfiles for a given namespace and list options.
	ListServiceProfile(ctx context.Context, opts ...client.ListOption) (*linkerd_io_v1alpha2.ServiceProfileList, error)
}

// ServiceProfileTransitionFunction instructs the ServiceProfileWriter how to transition between an existing
// ServiceProfile object and a desired on an Upsert
type ServiceProfileTransitionFunction func(existing, desired *linkerd_io_v1alpha2.ServiceProfile) error

// Writer knows how to create, delete, and update ServiceProfiles.
type ServiceProfileWriter interface {
	// Create saves the ServiceProfile object.
	CreateServiceProfile(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, opts ...client.CreateOption) error

	// Delete deletes the ServiceProfile object.
	DeleteServiceProfile(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ServiceProfile object.
	UpdateServiceProfile(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, opts ...client.UpdateOption) error

	// Patch patches the given ServiceProfile object.
	PatchServiceProfile(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ServiceProfile objects matching the given options.
	DeleteAllOfServiceProfile(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ServiceProfile object.
	UpsertServiceProfile(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, transitionFuncs ...ServiceProfileTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ServiceProfile object.
type ServiceProfileStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ServiceProfile object.
	UpdateServiceProfileStatus(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, opts ...client.UpdateOption) error

	// Patch patches the given ServiceProfile object's subresource.
	PatchServiceProfileStatus(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ServiceProfiles.
type ServiceProfileClient interface {
	ServiceProfileReader
	ServiceProfileWriter
	ServiceProfileStatusWriter
}

type serviceProfileClient struct {
	client client.Client
}

func NewServiceProfileClient(client client.Client) *serviceProfileClient {
	return &serviceProfileClient{client: client}
}

func (c *serviceProfileClient) GetServiceProfile(ctx context.Context, key client.ObjectKey) (*linkerd_io_v1alpha2.ServiceProfile, error) {
	obj := &linkerd_io_v1alpha2.ServiceProfile{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *serviceProfileClient) ListServiceProfile(ctx context.Context, opts ...client.ListOption) (*linkerd_io_v1alpha2.ServiceProfileList, error) {
	list := &linkerd_io_v1alpha2.ServiceProfileList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *serviceProfileClient) CreateServiceProfile(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *serviceProfileClient) DeleteServiceProfile(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &linkerd_io_v1alpha2.ServiceProfile{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *serviceProfileClient) UpdateServiceProfile(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *serviceProfileClient) PatchServiceProfile(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *serviceProfileClient) DeleteAllOfServiceProfile(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &linkerd_io_v1alpha2.ServiceProfile{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *serviceProfileClient) UpsertServiceProfile(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, transitionFuncs ...ServiceProfileTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*linkerd_io_v1alpha2.ServiceProfile), desired.(*linkerd_io_v1alpha2.ServiceProfile)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *serviceProfileClient) UpdateServiceProfileStatus(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *serviceProfileClient) PatchServiceProfileStatus(ctx context.Context, obj *linkerd_io_v1alpha2.ServiceProfile, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides ServiceProfileClients for multiple clusters.
type MulticlusterServiceProfileClient interface {
	// Cluster returns a ServiceProfileClient for the given cluster
	Cluster(cluster string) (ServiceProfileClient, error)
}

type multiclusterServiceProfileClient struct {
	client multicluster.Client
}

func NewMulticlusterServiceProfileClient(client multicluster.Client) MulticlusterServiceProfileClient {
	return &multiclusterServiceProfileClient{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewServiceProfileClient(client), nil
}
