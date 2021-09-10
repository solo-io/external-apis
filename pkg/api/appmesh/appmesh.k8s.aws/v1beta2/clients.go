// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1beta2

import (
	"context"

	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the appmesh.k8s.aws/v1beta2 APIs
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

// clienset for the appmesh.k8s.aws/v1beta2 APIs
type Clientset interface {
	// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
	Meshes() MeshClient
	// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
	VirtualServices() VirtualServiceClient
	// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
	VirtualNodes() VirtualNodeClient
	// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
	VirtualRouters() VirtualRouterClient
	// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
	VirtualGateways() VirtualGatewayClient
	// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
	GatewayRoutes() GatewayRouteClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := appmesh_k8s_aws_v1beta2.SchemeBuilder.AddToScheme(scheme); err != nil {
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

// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
func (c *clientSet) Meshes() MeshClient {
	return NewMeshClient(c.client)
}

// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
func (c *clientSet) VirtualServices() VirtualServiceClient {
	return NewVirtualServiceClient(c.client)
}

// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
func (c *clientSet) VirtualNodes() VirtualNodeClient {
	return NewVirtualNodeClient(c.client)
}

// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
func (c *clientSet) VirtualRouters() VirtualRouterClient {
	return NewVirtualRouterClient(c.client)
}

// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
func (c *clientSet) VirtualGateways() VirtualGatewayClient {
	return NewVirtualGatewayClient(c.client)
}

// clienset for the appmesh.k8s.aws/v1beta2/v1beta2 APIs
func (c *clientSet) GatewayRoutes() GatewayRouteClient {
	return NewGatewayRouteClient(c.client)
}

// Reader knows how to read and list Meshs.
type MeshReader interface {
	// Get retrieves a Mesh for the given object key
	GetMesh(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.Mesh, error)

	// List retrieves list of Meshs for a given namespace and list options.
	ListMesh(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.MeshList, error)
}

// MeshTransitionFunction instructs the MeshWriter how to transition between an existing
// Mesh object and a desired on an Upsert
type MeshTransitionFunction func(existing, desired *appmesh_k8s_aws_v1beta2.Mesh) error

// Writer knows how to create, delete, and update Meshs.
type MeshWriter interface {
	// Create saves the Mesh object.
	CreateMesh(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, opts ...client.CreateOption) error

	// Delete deletes the Mesh object.
	DeleteMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Mesh object.
	UpdateMesh(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, opts ...client.UpdateOption) error

	// Patch patches the given Mesh object.
	PatchMesh(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Mesh objects matching the given options.
	DeleteAllOfMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Mesh object.
	UpsertMesh(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, transitionFuncs ...MeshTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Mesh object.
type MeshStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Mesh object.
	UpdateMeshStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, opts ...client.UpdateOption) error

	// Patch patches the given Mesh object's subresource.
	PatchMeshStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Meshs.
type MeshClient interface {
	MeshReader
	MeshWriter
	MeshStatusWriter
}

type meshClient struct {
	client client.Client
}

func NewMeshClient(client client.Client) *meshClient {
	return &meshClient{client: client}
}

func (c *meshClient) GetMesh(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.Mesh, error) {
	obj := &appmesh_k8s_aws_v1beta2.Mesh{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *meshClient) ListMesh(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.MeshList, error) {
	list := &appmesh_k8s_aws_v1beta2.MeshList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *meshClient) CreateMesh(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *meshClient) DeleteMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &appmesh_k8s_aws_v1beta2.Mesh{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *meshClient) UpdateMesh(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *meshClient) PatchMesh(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *meshClient) DeleteAllOfMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &appmesh_k8s_aws_v1beta2.Mesh{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *meshClient) UpsertMesh(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, transitionFuncs ...MeshTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*appmesh_k8s_aws_v1beta2.Mesh), desired.(*appmesh_k8s_aws_v1beta2.Mesh)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *meshClient) UpdateMeshStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *meshClient) PatchMeshStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.Mesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides MeshClients for multiple clusters.
type MulticlusterMeshClient interface {
	// Cluster returns a MeshClient for the given cluster
	Cluster(cluster string) (MeshClient, error)
}

type multiclusterMeshClient struct {
	client multicluster.Client
}

func NewMulticlusterMeshClient(client multicluster.Client) MulticlusterMeshClient {
	return &multiclusterMeshClient{client: client}
}

func (m *multiclusterMeshClient) Cluster(cluster string) (MeshClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewMeshClient(client), nil
}

// Reader knows how to read and list VirtualServices.
type VirtualServiceReader interface {
	// Get retrieves a VirtualService for the given object key
	GetVirtualService(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.VirtualService, error)

	// List retrieves list of VirtualServices for a given namespace and list options.
	ListVirtualService(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.VirtualServiceList, error)
}

// VirtualServiceTransitionFunction instructs the VirtualServiceWriter how to transition between an existing
// VirtualService object and a desired on an Upsert
type VirtualServiceTransitionFunction func(existing, desired *appmesh_k8s_aws_v1beta2.VirtualService) error

// Writer knows how to create, delete, and update VirtualServices.
type VirtualServiceWriter interface {
	// Create saves the VirtualService object.
	CreateVirtualService(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, opts ...client.CreateOption) error

	// Delete deletes the VirtualService object.
	DeleteVirtualService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualService object.
	UpdateVirtualService(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, opts ...client.UpdateOption) error

	// Patch patches the given VirtualService object.
	PatchVirtualService(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualService objects matching the given options.
	DeleteAllOfVirtualService(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualService object.
	UpsertVirtualService(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, transitionFuncs ...VirtualServiceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualService object.
type VirtualServiceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualService object.
	UpdateVirtualServiceStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, opts ...client.UpdateOption) error

	// Patch patches the given VirtualService object's subresource.
	PatchVirtualServiceStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on VirtualServices.
type VirtualServiceClient interface {
	VirtualServiceReader
	VirtualServiceWriter
	VirtualServiceStatusWriter
}

type virtualServiceClient struct {
	client client.Client
}

func NewVirtualServiceClient(client client.Client) *virtualServiceClient {
	return &virtualServiceClient{client: client}
}

func (c *virtualServiceClient) GetVirtualService(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.VirtualService, error) {
	obj := &appmesh_k8s_aws_v1beta2.VirtualService{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualServiceClient) ListVirtualService(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.VirtualServiceList, error) {
	list := &appmesh_k8s_aws_v1beta2.VirtualServiceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualServiceClient) CreateVirtualService(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualServiceClient) DeleteVirtualService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &appmesh_k8s_aws_v1beta2.VirtualService{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualServiceClient) UpdateVirtualService(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualServiceClient) PatchVirtualService(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualServiceClient) DeleteAllOfVirtualService(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &appmesh_k8s_aws_v1beta2.VirtualService{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualServiceClient) UpsertVirtualService(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, transitionFuncs ...VirtualServiceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*appmesh_k8s_aws_v1beta2.VirtualService), desired.(*appmesh_k8s_aws_v1beta2.VirtualService)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualServiceClient) UpdateVirtualServiceStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualServiceClient) PatchVirtualServiceStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides VirtualServiceClients for multiple clusters.
type MulticlusterVirtualServiceClient interface {
	// Cluster returns a VirtualServiceClient for the given cluster
	Cluster(cluster string) (VirtualServiceClient, error)
}

type multiclusterVirtualServiceClient struct {
	client multicluster.Client
}

func NewMulticlusterVirtualServiceClient(client multicluster.Client) MulticlusterVirtualServiceClient {
	return &multiclusterVirtualServiceClient{client: client}
}

func (m *multiclusterVirtualServiceClient) Cluster(cluster string) (VirtualServiceClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewVirtualServiceClient(client), nil
}

// Reader knows how to read and list VirtualNodes.
type VirtualNodeReader interface {
	// Get retrieves a VirtualNode for the given object key
	GetVirtualNode(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.VirtualNode, error)

	// List retrieves list of VirtualNodes for a given namespace and list options.
	ListVirtualNode(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.VirtualNodeList, error)
}

// VirtualNodeTransitionFunction instructs the VirtualNodeWriter how to transition between an existing
// VirtualNode object and a desired on an Upsert
type VirtualNodeTransitionFunction func(existing, desired *appmesh_k8s_aws_v1beta2.VirtualNode) error

// Writer knows how to create, delete, and update VirtualNodes.
type VirtualNodeWriter interface {
	// Create saves the VirtualNode object.
	CreateVirtualNode(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, opts ...client.CreateOption) error

	// Delete deletes the VirtualNode object.
	DeleteVirtualNode(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualNode object.
	UpdateVirtualNode(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, opts ...client.UpdateOption) error

	// Patch patches the given VirtualNode object.
	PatchVirtualNode(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualNode objects matching the given options.
	DeleteAllOfVirtualNode(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualNode object.
	UpsertVirtualNode(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, transitionFuncs ...VirtualNodeTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualNode object.
type VirtualNodeStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualNode object.
	UpdateVirtualNodeStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, opts ...client.UpdateOption) error

	// Patch patches the given VirtualNode object's subresource.
	PatchVirtualNodeStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on VirtualNodes.
type VirtualNodeClient interface {
	VirtualNodeReader
	VirtualNodeWriter
	VirtualNodeStatusWriter
}

type virtualNodeClient struct {
	client client.Client
}

func NewVirtualNodeClient(client client.Client) *virtualNodeClient {
	return &virtualNodeClient{client: client}
}

func (c *virtualNodeClient) GetVirtualNode(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.VirtualNode, error) {
	obj := &appmesh_k8s_aws_v1beta2.VirtualNode{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualNodeClient) ListVirtualNode(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.VirtualNodeList, error) {
	list := &appmesh_k8s_aws_v1beta2.VirtualNodeList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualNodeClient) CreateVirtualNode(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualNodeClient) DeleteVirtualNode(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &appmesh_k8s_aws_v1beta2.VirtualNode{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualNodeClient) UpdateVirtualNode(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualNodeClient) PatchVirtualNode(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualNodeClient) DeleteAllOfVirtualNode(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &appmesh_k8s_aws_v1beta2.VirtualNode{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualNodeClient) UpsertVirtualNode(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, transitionFuncs ...VirtualNodeTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*appmesh_k8s_aws_v1beta2.VirtualNode), desired.(*appmesh_k8s_aws_v1beta2.VirtualNode)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualNodeClient) UpdateVirtualNodeStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualNodeClient) PatchVirtualNodeStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualNode, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides VirtualNodeClients for multiple clusters.
type MulticlusterVirtualNodeClient interface {
	// Cluster returns a VirtualNodeClient for the given cluster
	Cluster(cluster string) (VirtualNodeClient, error)
}

type multiclusterVirtualNodeClient struct {
	client multicluster.Client
}

func NewMulticlusterVirtualNodeClient(client multicluster.Client) MulticlusterVirtualNodeClient {
	return &multiclusterVirtualNodeClient{client: client}
}

func (m *multiclusterVirtualNodeClient) Cluster(cluster string) (VirtualNodeClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewVirtualNodeClient(client), nil
}

// Reader knows how to read and list VirtualRouters.
type VirtualRouterReader interface {
	// Get retrieves a VirtualRouter for the given object key
	GetVirtualRouter(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.VirtualRouter, error)

	// List retrieves list of VirtualRouters for a given namespace and list options.
	ListVirtualRouter(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.VirtualRouterList, error)
}

// VirtualRouterTransitionFunction instructs the VirtualRouterWriter how to transition between an existing
// VirtualRouter object and a desired on an Upsert
type VirtualRouterTransitionFunction func(existing, desired *appmesh_k8s_aws_v1beta2.VirtualRouter) error

// Writer knows how to create, delete, and update VirtualRouters.
type VirtualRouterWriter interface {
	// Create saves the VirtualRouter object.
	CreateVirtualRouter(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, opts ...client.CreateOption) error

	// Delete deletes the VirtualRouter object.
	DeleteVirtualRouter(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualRouter object.
	UpdateVirtualRouter(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, opts ...client.UpdateOption) error

	// Patch patches the given VirtualRouter object.
	PatchVirtualRouter(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualRouter objects matching the given options.
	DeleteAllOfVirtualRouter(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualRouter object.
	UpsertVirtualRouter(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, transitionFuncs ...VirtualRouterTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualRouter object.
type VirtualRouterStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualRouter object.
	UpdateVirtualRouterStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, opts ...client.UpdateOption) error

	// Patch patches the given VirtualRouter object's subresource.
	PatchVirtualRouterStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on VirtualRouters.
type VirtualRouterClient interface {
	VirtualRouterReader
	VirtualRouterWriter
	VirtualRouterStatusWriter
}

type virtualRouterClient struct {
	client client.Client
}

func NewVirtualRouterClient(client client.Client) *virtualRouterClient {
	return &virtualRouterClient{client: client}
}

func (c *virtualRouterClient) GetVirtualRouter(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.VirtualRouter, error) {
	obj := &appmesh_k8s_aws_v1beta2.VirtualRouter{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualRouterClient) ListVirtualRouter(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.VirtualRouterList, error) {
	list := &appmesh_k8s_aws_v1beta2.VirtualRouterList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualRouterClient) CreateVirtualRouter(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualRouterClient) DeleteVirtualRouter(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &appmesh_k8s_aws_v1beta2.VirtualRouter{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualRouterClient) UpdateVirtualRouter(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualRouterClient) PatchVirtualRouter(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualRouterClient) DeleteAllOfVirtualRouter(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &appmesh_k8s_aws_v1beta2.VirtualRouter{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualRouterClient) UpsertVirtualRouter(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, transitionFuncs ...VirtualRouterTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*appmesh_k8s_aws_v1beta2.VirtualRouter), desired.(*appmesh_k8s_aws_v1beta2.VirtualRouter)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualRouterClient) UpdateVirtualRouterStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualRouterClient) PatchVirtualRouterStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualRouter, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides VirtualRouterClients for multiple clusters.
type MulticlusterVirtualRouterClient interface {
	// Cluster returns a VirtualRouterClient for the given cluster
	Cluster(cluster string) (VirtualRouterClient, error)
}

type multiclusterVirtualRouterClient struct {
	client multicluster.Client
}

func NewMulticlusterVirtualRouterClient(client multicluster.Client) MulticlusterVirtualRouterClient {
	return &multiclusterVirtualRouterClient{client: client}
}

func (m *multiclusterVirtualRouterClient) Cluster(cluster string) (VirtualRouterClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewVirtualRouterClient(client), nil
}

// Reader knows how to read and list VirtualGateways.
type VirtualGatewayReader interface {
	// Get retrieves a VirtualGateway for the given object key
	GetVirtualGateway(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.VirtualGateway, error)

	// List retrieves list of VirtualGateways for a given namespace and list options.
	ListVirtualGateway(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.VirtualGatewayList, error)
}

// VirtualGatewayTransitionFunction instructs the VirtualGatewayWriter how to transition between an existing
// VirtualGateway object and a desired on an Upsert
type VirtualGatewayTransitionFunction func(existing, desired *appmesh_k8s_aws_v1beta2.VirtualGateway) error

// Writer knows how to create, delete, and update VirtualGateways.
type VirtualGatewayWriter interface {
	// Create saves the VirtualGateway object.
	CreateVirtualGateway(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, opts ...client.CreateOption) error

	// Delete deletes the VirtualGateway object.
	DeleteVirtualGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualGateway object.
	UpdateVirtualGateway(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, opts ...client.UpdateOption) error

	// Patch patches the given VirtualGateway object.
	PatchVirtualGateway(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualGateway objects matching the given options.
	DeleteAllOfVirtualGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualGateway object.
	UpsertVirtualGateway(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, transitionFuncs ...VirtualGatewayTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualGateway object.
type VirtualGatewayStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualGateway object.
	UpdateVirtualGatewayStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, opts ...client.UpdateOption) error

	// Patch patches the given VirtualGateway object's subresource.
	PatchVirtualGatewayStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on VirtualGateways.
type VirtualGatewayClient interface {
	VirtualGatewayReader
	VirtualGatewayWriter
	VirtualGatewayStatusWriter
}

type virtualGatewayClient struct {
	client client.Client
}

func NewVirtualGatewayClient(client client.Client) *virtualGatewayClient {
	return &virtualGatewayClient{client: client}
}

func (c *virtualGatewayClient) GetVirtualGateway(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.VirtualGateway, error) {
	obj := &appmesh_k8s_aws_v1beta2.VirtualGateway{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualGatewayClient) ListVirtualGateway(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.VirtualGatewayList, error) {
	list := &appmesh_k8s_aws_v1beta2.VirtualGatewayList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualGatewayClient) CreateVirtualGateway(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualGatewayClient) DeleteVirtualGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &appmesh_k8s_aws_v1beta2.VirtualGateway{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualGatewayClient) UpdateVirtualGateway(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualGatewayClient) PatchVirtualGateway(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualGatewayClient) DeleteAllOfVirtualGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &appmesh_k8s_aws_v1beta2.VirtualGateway{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualGatewayClient) UpsertVirtualGateway(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, transitionFuncs ...VirtualGatewayTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*appmesh_k8s_aws_v1beta2.VirtualGateway), desired.(*appmesh_k8s_aws_v1beta2.VirtualGateway)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualGatewayClient) UpdateVirtualGatewayStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualGatewayClient) PatchVirtualGatewayStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.VirtualGateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides VirtualGatewayClients for multiple clusters.
type MulticlusterVirtualGatewayClient interface {
	// Cluster returns a VirtualGatewayClient for the given cluster
	Cluster(cluster string) (VirtualGatewayClient, error)
}

type multiclusterVirtualGatewayClient struct {
	client multicluster.Client
}

func NewMulticlusterVirtualGatewayClient(client multicluster.Client) MulticlusterVirtualGatewayClient {
	return &multiclusterVirtualGatewayClient{client: client}
}

func (m *multiclusterVirtualGatewayClient) Cluster(cluster string) (VirtualGatewayClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewVirtualGatewayClient(client), nil
}

// Reader knows how to read and list GatewayRoutes.
type GatewayRouteReader interface {
	// Get retrieves a GatewayRoute for the given object key
	GetGatewayRoute(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.GatewayRoute, error)

	// List retrieves list of GatewayRoutes for a given namespace and list options.
	ListGatewayRoute(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.GatewayRouteList, error)
}

// GatewayRouteTransitionFunction instructs the GatewayRouteWriter how to transition between an existing
// GatewayRoute object and a desired on an Upsert
type GatewayRouteTransitionFunction func(existing, desired *appmesh_k8s_aws_v1beta2.GatewayRoute) error

// Writer knows how to create, delete, and update GatewayRoutes.
type GatewayRouteWriter interface {
	// Create saves the GatewayRoute object.
	CreateGatewayRoute(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, opts ...client.CreateOption) error

	// Delete deletes the GatewayRoute object.
	DeleteGatewayRoute(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given GatewayRoute object.
	UpdateGatewayRoute(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, opts ...client.UpdateOption) error

	// Patch patches the given GatewayRoute object.
	PatchGatewayRoute(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all GatewayRoute objects matching the given options.
	DeleteAllOfGatewayRoute(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the GatewayRoute object.
	UpsertGatewayRoute(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, transitionFuncs ...GatewayRouteTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a GatewayRoute object.
type GatewayRouteStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given GatewayRoute object.
	UpdateGatewayRouteStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, opts ...client.UpdateOption) error

	// Patch patches the given GatewayRoute object's subresource.
	PatchGatewayRouteStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on GatewayRoutes.
type GatewayRouteClient interface {
	GatewayRouteReader
	GatewayRouteWriter
	GatewayRouteStatusWriter
}

type gatewayRouteClient struct {
	client client.Client
}

func NewGatewayRouteClient(client client.Client) *gatewayRouteClient {
	return &gatewayRouteClient{client: client}
}

func (c *gatewayRouteClient) GetGatewayRoute(ctx context.Context, key client.ObjectKey) (*appmesh_k8s_aws_v1beta2.GatewayRoute, error) {
	obj := &appmesh_k8s_aws_v1beta2.GatewayRoute{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *gatewayRouteClient) ListGatewayRoute(ctx context.Context, opts ...client.ListOption) (*appmesh_k8s_aws_v1beta2.GatewayRouteList, error) {
	list := &appmesh_k8s_aws_v1beta2.GatewayRouteList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *gatewayRouteClient) CreateGatewayRoute(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *gatewayRouteClient) DeleteGatewayRoute(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &appmesh_k8s_aws_v1beta2.GatewayRoute{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *gatewayRouteClient) UpdateGatewayRoute(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *gatewayRouteClient) PatchGatewayRoute(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *gatewayRouteClient) DeleteAllOfGatewayRoute(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &appmesh_k8s_aws_v1beta2.GatewayRoute{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *gatewayRouteClient) UpsertGatewayRoute(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, transitionFuncs ...GatewayRouteTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*appmesh_k8s_aws_v1beta2.GatewayRoute), desired.(*appmesh_k8s_aws_v1beta2.GatewayRoute)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *gatewayRouteClient) UpdateGatewayRouteStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *gatewayRouteClient) PatchGatewayRouteStatus(ctx context.Context, obj *appmesh_k8s_aws_v1beta2.GatewayRoute, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides GatewayRouteClients for multiple clusters.
type MulticlusterGatewayRouteClient interface {
	// Cluster returns a GatewayRouteClient for the given cluster
	Cluster(cluster string) (GatewayRouteClient, error)
}

type multiclusterGatewayRouteClient struct {
	client multicluster.Client
}

func NewMulticlusterGatewayRouteClient(client multicluster.Client) MulticlusterGatewayRouteClient {
	return &multiclusterGatewayRouteClient{client: client}
}

func (m *multiclusterGatewayRouteClient) Cluster(cluster string) (GatewayRouteClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewGatewayRouteClient(client), nil
}
