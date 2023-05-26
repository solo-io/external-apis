// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	install_istio_io_v1alpha1 "istio.io/istio/operator/pkg/apis/istio/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the install.istio.io/v1alpha1 APIs
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

// clienset for the install.istio.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the install.istio.io/v1alpha1/v1alpha1 APIs
	IstioOperators() IstioOperatorClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := install_istio_io_v1alpha1.SchemeBuilder.AddToScheme(scheme); err != nil {
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

// clienset for the install.istio.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) IstioOperators() IstioOperatorClient {
	return NewIstioOperatorClient(c.client)
}

// Reader knows how to read and list IstioOperators.
type IstioOperatorReader interface {
	// Get retrieves a IstioOperator for the given object key
	GetIstioOperator(ctx context.Context, key client.ObjectKey) (*install_istio_io_v1alpha1.IstioOperator, error)

	// List retrieves list of IstioOperators for a given namespace and list options.
	ListIstioOperator(ctx context.Context, opts ...client.ListOption) (*install_istio_io_v1alpha1.IstioOperatorList, error)
}

// IstioOperatorTransitionFunction instructs the IstioOperatorWriter how to transition between an existing
// IstioOperator object and a desired on an Upsert
type IstioOperatorTransitionFunction func(existing, desired *install_istio_io_v1alpha1.IstioOperator) error

// Writer knows how to create, delete, and update IstioOperators.
type IstioOperatorWriter interface {
	// Create saves the IstioOperator object.
	CreateIstioOperator(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, opts ...client.CreateOption) error

	// Delete deletes the IstioOperator object.
	DeleteIstioOperator(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given IstioOperator object.
	UpdateIstioOperator(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, opts ...client.UpdateOption) error

	// Patch patches the given IstioOperator object.
	PatchIstioOperator(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all IstioOperator objects matching the given options.
	DeleteAllOfIstioOperator(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the IstioOperator object.
	UpsertIstioOperator(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, transitionFuncs ...IstioOperatorTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a IstioOperator object.
type IstioOperatorStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given IstioOperator object.
	UpdateIstioOperatorStatus(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given IstioOperator object's subresource.
	PatchIstioOperatorStatus(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on IstioOperators.
type IstioOperatorClient interface {
	IstioOperatorReader
	IstioOperatorWriter
	IstioOperatorStatusWriter
}

type istioOperatorClient struct {
	client client.Client
}

func NewIstioOperatorClient(client client.Client) *istioOperatorClient {
	return &istioOperatorClient{client: client}
}

func (c *istioOperatorClient) GetIstioOperator(ctx context.Context, key client.ObjectKey) (*install_istio_io_v1alpha1.IstioOperator, error) {
	obj := &install_istio_io_v1alpha1.IstioOperator{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *istioOperatorClient) ListIstioOperator(ctx context.Context, opts ...client.ListOption) (*install_istio_io_v1alpha1.IstioOperatorList, error) {
	list := &install_istio_io_v1alpha1.IstioOperatorList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *istioOperatorClient) CreateIstioOperator(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *istioOperatorClient) DeleteIstioOperator(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &install_istio_io_v1alpha1.IstioOperator{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *istioOperatorClient) UpdateIstioOperator(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *istioOperatorClient) PatchIstioOperator(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *istioOperatorClient) DeleteAllOfIstioOperator(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &install_istio_io_v1alpha1.IstioOperator{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *istioOperatorClient) UpsertIstioOperator(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, transitionFuncs ...IstioOperatorTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*install_istio_io_v1alpha1.IstioOperator), desired.(*install_istio_io_v1alpha1.IstioOperator)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *istioOperatorClient) UpdateIstioOperatorStatus(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *istioOperatorClient) PatchIstioOperatorStatus(ctx context.Context, obj *install_istio_io_v1alpha1.IstioOperator, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides IstioOperatorClients for multiple clusters.
type MulticlusterIstioOperatorClient interface {
	// Cluster returns a IstioOperatorClient for the given cluster
	Cluster(cluster string) (IstioOperatorClient, error)
}

type multiclusterIstioOperatorClient struct {
	client multicluster.Client
}

func NewMulticlusterIstioOperatorClient(client multicluster.Client) MulticlusterIstioOperatorClient {
	return &multiclusterIstioOperatorClient{client: client}
}

func (m *multiclusterIstioOperatorClient) Cluster(cluster string) (IstioOperatorClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewIstioOperatorClient(client), nil
}
