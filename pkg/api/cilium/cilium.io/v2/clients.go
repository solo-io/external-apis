// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v2

import (
	"context"

	cilium_io_v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the cilium.io/v2 APIs
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

// clienset for the cilium.io/v2 APIs
type Clientset interface {
	// clienset for the cilium.io/v2/v2 APIs
	CiliumNetworkPolicies() CiliumNetworkPolicyClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := cilium_io_v2.SchemeBuilder.AddToScheme(scheme); err != nil {
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

// clienset for the cilium.io/v2/v2 APIs
func (c *clientSet) CiliumNetworkPolicies() CiliumNetworkPolicyClient {
	return NewCiliumNetworkPolicyClient(c.client)
}

// Reader knows how to read and list CiliumNetworkPolicys.
type CiliumNetworkPolicyReader interface {
	// Get retrieves a CiliumNetworkPolicy for the given object key
	GetCiliumNetworkPolicy(ctx context.Context, key client.ObjectKey) (*cilium_io_v2.CiliumNetworkPolicy, error)

	// List retrieves list of CiliumNetworkPolicys for a given namespace and list options.
	ListCiliumNetworkPolicy(ctx context.Context, opts ...client.ListOption) (*cilium_io_v2.CiliumNetworkPolicyList, error)
}

// CiliumNetworkPolicyTransitionFunction instructs the CiliumNetworkPolicyWriter how to transition between an existing
// CiliumNetworkPolicy object and a desired on an Upsert
type CiliumNetworkPolicyTransitionFunction func(existing, desired *cilium_io_v2.CiliumNetworkPolicy) error

// Writer knows how to create, delete, and update CiliumNetworkPolicys.
type CiliumNetworkPolicyWriter interface {
	// Create saves the CiliumNetworkPolicy object.
	CreateCiliumNetworkPolicy(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, opts ...client.CreateOption) error

	// Delete deletes the CiliumNetworkPolicy object.
	DeleteCiliumNetworkPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given CiliumNetworkPolicy object.
	UpdateCiliumNetworkPolicy(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, opts ...client.UpdateOption) error

	// Patch patches the given CiliumNetworkPolicy object.
	PatchCiliumNetworkPolicy(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all CiliumNetworkPolicy objects matching the given options.
	DeleteAllOfCiliumNetworkPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the CiliumNetworkPolicy object.
	UpsertCiliumNetworkPolicy(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, transitionFuncs ...CiliumNetworkPolicyTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a CiliumNetworkPolicy object.
type CiliumNetworkPolicyStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given CiliumNetworkPolicy object.
	UpdateCiliumNetworkPolicyStatus(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given CiliumNetworkPolicy object's subresource.
	PatchCiliumNetworkPolicyStatus(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on CiliumNetworkPolicys.
type CiliumNetworkPolicyClient interface {
	CiliumNetworkPolicyReader
	CiliumNetworkPolicyWriter
	CiliumNetworkPolicyStatusWriter
}

type ciliumNetworkPolicyClient struct {
	client client.Client
}

func NewCiliumNetworkPolicyClient(client client.Client) *ciliumNetworkPolicyClient {
	return &ciliumNetworkPolicyClient{client: client}
}

func (c *ciliumNetworkPolicyClient) GetCiliumNetworkPolicy(ctx context.Context, key client.ObjectKey) (*cilium_io_v2.CiliumNetworkPolicy, error) {
	obj := &cilium_io_v2.CiliumNetworkPolicy{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *ciliumNetworkPolicyClient) ListCiliumNetworkPolicy(ctx context.Context, opts ...client.ListOption) (*cilium_io_v2.CiliumNetworkPolicyList, error) {
	list := &cilium_io_v2.CiliumNetworkPolicyList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *ciliumNetworkPolicyClient) CreateCiliumNetworkPolicy(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *ciliumNetworkPolicyClient) DeleteCiliumNetworkPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &cilium_io_v2.CiliumNetworkPolicy{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *ciliumNetworkPolicyClient) UpdateCiliumNetworkPolicy(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *ciliumNetworkPolicyClient) PatchCiliumNetworkPolicy(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *ciliumNetworkPolicyClient) DeleteAllOfCiliumNetworkPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &cilium_io_v2.CiliumNetworkPolicy{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *ciliumNetworkPolicyClient) UpsertCiliumNetworkPolicy(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, transitionFuncs ...CiliumNetworkPolicyTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*cilium_io_v2.CiliumNetworkPolicy), desired.(*cilium_io_v2.CiliumNetworkPolicy)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *ciliumNetworkPolicyClient) UpdateCiliumNetworkPolicyStatus(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *ciliumNetworkPolicyClient) PatchCiliumNetworkPolicyStatus(ctx context.Context, obj *cilium_io_v2.CiliumNetworkPolicy, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides CiliumNetworkPolicyClients for multiple clusters.
type MulticlusterCiliumNetworkPolicyClient interface {
	// Cluster returns a CiliumNetworkPolicyClient for the given cluster
	Cluster(cluster string) (CiliumNetworkPolicyClient, error)
}

type multiclusterCiliumNetworkPolicyClient struct {
	client multicluster.Client
}

func NewMulticlusterCiliumNetworkPolicyClient(client multicluster.Client) MulticlusterCiliumNetworkPolicyClient {
	return &multiclusterCiliumNetworkPolicyClient{client: client}
}

func (m *multiclusterCiliumNetworkPolicyClient) Cluster(cluster string) (CiliumNetworkPolicyClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewCiliumNetworkPolicyClient(client), nil
}
