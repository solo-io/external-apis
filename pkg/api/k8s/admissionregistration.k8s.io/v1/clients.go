// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	admissionregistration_k8s_io_v1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the admissionregistration.k8s.io/v1 APIs
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

// clienset for the admissionregistration.k8s.io/v1 APIs
type Clientset interface {
	// clienset for the admissionregistration.k8s.io/v1/v1 APIs
	ValidatingWebhookConfigurations() ValidatingWebhookConfigurationClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := admissionregistration_k8s_io_v1.SchemeBuilder.AddToScheme(scheme); err != nil {
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

// clienset for the admissionregistration.k8s.io/v1/v1 APIs
func (c *clientSet) ValidatingWebhookConfigurations() ValidatingWebhookConfigurationClient {
	return NewValidatingWebhookConfigurationClient(c.client)
}

// Reader knows how to read and list ValidatingWebhookConfigurations.
type ValidatingWebhookConfigurationReader interface {
	// Get retrieves a ValidatingWebhookConfiguration for the given object key
	GetValidatingWebhookConfiguration(ctx context.Context, key client.ObjectKey) (*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, error)

	// List retrieves list of ValidatingWebhookConfigurations for a given namespace and list options.
	ListValidatingWebhookConfiguration(ctx context.Context, opts ...client.ListOption) (*admissionregistration_k8s_io_v1.ValidatingWebhookConfigurationList, error)
}

// ValidatingWebhookConfigurationTransitionFunction instructs the ValidatingWebhookConfigurationWriter how to transition between an existing
// ValidatingWebhookConfiguration object and a desired on an Upsert
type ValidatingWebhookConfigurationTransitionFunction func(existing, desired *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) error

// Writer knows how to create, delete, and update ValidatingWebhookConfigurations.
type ValidatingWebhookConfigurationWriter interface {
	// Create saves the ValidatingWebhookConfiguration object.
	CreateValidatingWebhookConfiguration(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, opts ...client.CreateOption) error

	// Delete deletes the ValidatingWebhookConfiguration object.
	DeleteValidatingWebhookConfiguration(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ValidatingWebhookConfiguration object.
	UpdateValidatingWebhookConfiguration(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, opts ...client.UpdateOption) error

	// Patch patches the given ValidatingWebhookConfiguration object.
	PatchValidatingWebhookConfiguration(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ValidatingWebhookConfiguration objects matching the given options.
	DeleteAllOfValidatingWebhookConfiguration(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ValidatingWebhookConfiguration object.
	UpsertValidatingWebhookConfiguration(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, transitionFuncs ...ValidatingWebhookConfigurationTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ValidatingWebhookConfiguration object.
type ValidatingWebhookConfigurationStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ValidatingWebhookConfiguration object.
	UpdateValidatingWebhookConfigurationStatus(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, opts ...client.UpdateOption) error

	// Patch patches the given ValidatingWebhookConfiguration object's subresource.
	PatchValidatingWebhookConfigurationStatus(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ValidatingWebhookConfigurations.
type ValidatingWebhookConfigurationClient interface {
	ValidatingWebhookConfigurationReader
	ValidatingWebhookConfigurationWriter
	ValidatingWebhookConfigurationStatusWriter
}

type validatingWebhookConfigurationClient struct {
	client client.Client
}

func NewValidatingWebhookConfigurationClient(client client.Client) *validatingWebhookConfigurationClient {
	return &validatingWebhookConfigurationClient{client: client}
}

func (c *validatingWebhookConfigurationClient) GetValidatingWebhookConfiguration(ctx context.Context, key client.ObjectKey) (*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, error) {
	obj := &admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *validatingWebhookConfigurationClient) ListValidatingWebhookConfiguration(ctx context.Context, opts ...client.ListOption) (*admissionregistration_k8s_io_v1.ValidatingWebhookConfigurationList, error) {
	list := &admissionregistration_k8s_io_v1.ValidatingWebhookConfigurationList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *validatingWebhookConfigurationClient) CreateValidatingWebhookConfiguration(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) DeleteValidatingWebhookConfiguration(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) UpdateValidatingWebhookConfiguration(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) PatchValidatingWebhookConfiguration(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *validatingWebhookConfigurationClient) DeleteAllOfValidatingWebhookConfiguration(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) UpsertValidatingWebhookConfiguration(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, transitionFuncs ...ValidatingWebhookConfigurationTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration), desired.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *validatingWebhookConfigurationClient) UpdateValidatingWebhookConfigurationStatus(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *validatingWebhookConfigurationClient) PatchValidatingWebhookConfigurationStatus(ctx context.Context, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides ValidatingWebhookConfigurationClients for multiple clusters.
type MulticlusterValidatingWebhookConfigurationClient interface {
	// Cluster returns a ValidatingWebhookConfigurationClient for the given cluster
	Cluster(cluster string) (ValidatingWebhookConfigurationClient, error)
}

type multiclusterValidatingWebhookConfigurationClient struct {
	client multicluster.Client
}

func NewMulticlusterValidatingWebhookConfigurationClient(client multicluster.Client) MulticlusterValidatingWebhookConfigurationClient {
	return &multiclusterValidatingWebhookConfigurationClient{client: client}
}

func (m *multiclusterValidatingWebhookConfigurationClient) Cluster(cluster string) (ValidatingWebhookConfigurationClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewValidatingWebhookConfigurationClient(client), nil
}
