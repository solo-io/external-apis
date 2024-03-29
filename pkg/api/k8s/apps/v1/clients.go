// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	apps_v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the apps/v1 APIs
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

// clienset for the apps/v1 APIs
type Clientset interface {
	// clienset for the apps/v1/v1 APIs
	Deployments() DeploymentClient
	// clienset for the apps/v1/v1 APIs
	ReplicaSets() ReplicaSetClient
	// clienset for the apps/v1/v1 APIs
	DaemonSets() DaemonSetClient
	// clienset for the apps/v1/v1 APIs
	StatefulSets() StatefulSetClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := apps_v1.SchemeBuilder.AddToScheme(scheme); err != nil {
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

// clienset for the apps/v1/v1 APIs
func (c *clientSet) Deployments() DeploymentClient {
	return NewDeploymentClient(c.client)
}

// clienset for the apps/v1/v1 APIs
func (c *clientSet) ReplicaSets() ReplicaSetClient {
	return NewReplicaSetClient(c.client)
}

// clienset for the apps/v1/v1 APIs
func (c *clientSet) DaemonSets() DaemonSetClient {
	return NewDaemonSetClient(c.client)
}

// clienset for the apps/v1/v1 APIs
func (c *clientSet) StatefulSets() StatefulSetClient {
	return NewStatefulSetClient(c.client)
}

// Reader knows how to read and list Deployments.
type DeploymentReader interface {
	// Get retrieves a Deployment for the given object key
	GetDeployment(ctx context.Context, key client.ObjectKey) (*apps_v1.Deployment, error)

	// List retrieves list of Deployments for a given namespace and list options.
	ListDeployment(ctx context.Context, opts ...client.ListOption) (*apps_v1.DeploymentList, error)
}

// DeploymentTransitionFunction instructs the DeploymentWriter how to transition between an existing
// Deployment object and a desired on an Upsert
type DeploymentTransitionFunction func(existing, desired *apps_v1.Deployment) error

// Writer knows how to create, delete, and update Deployments.
type DeploymentWriter interface {
	// Create saves the Deployment object.
	CreateDeployment(ctx context.Context, obj *apps_v1.Deployment, opts ...client.CreateOption) error

	// Delete deletes the Deployment object.
	DeleteDeployment(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Deployment object.
	UpdateDeployment(ctx context.Context, obj *apps_v1.Deployment, opts ...client.UpdateOption) error

	// Patch patches the given Deployment object.
	PatchDeployment(ctx context.Context, obj *apps_v1.Deployment, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Deployment objects matching the given options.
	DeleteAllOfDeployment(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Deployment object.
	UpsertDeployment(ctx context.Context, obj *apps_v1.Deployment, transitionFuncs ...DeploymentTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Deployment object.
type DeploymentStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Deployment object.
	UpdateDeploymentStatus(ctx context.Context, obj *apps_v1.Deployment, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given Deployment object's subresource.
	PatchDeploymentStatus(ctx context.Context, obj *apps_v1.Deployment, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on Deployments.
type DeploymentClient interface {
	DeploymentReader
	DeploymentWriter
	DeploymentStatusWriter
}

type deploymentClient struct {
	client client.Client
}

func NewDeploymentClient(client client.Client) *deploymentClient {
	return &deploymentClient{client: client}
}

func (c *deploymentClient) GetDeployment(ctx context.Context, key client.ObjectKey) (*apps_v1.Deployment, error) {
	obj := &apps_v1.Deployment{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *deploymentClient) ListDeployment(ctx context.Context, opts ...client.ListOption) (*apps_v1.DeploymentList, error) {
	list := &apps_v1.DeploymentList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *deploymentClient) CreateDeployment(ctx context.Context, obj *apps_v1.Deployment, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *deploymentClient) DeleteDeployment(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &apps_v1.Deployment{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *deploymentClient) UpdateDeployment(ctx context.Context, obj *apps_v1.Deployment, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *deploymentClient) PatchDeployment(ctx context.Context, obj *apps_v1.Deployment, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *deploymentClient) DeleteAllOfDeployment(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &apps_v1.Deployment{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *deploymentClient) UpsertDeployment(ctx context.Context, obj *apps_v1.Deployment, transitionFuncs ...DeploymentTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*apps_v1.Deployment), desired.(*apps_v1.Deployment)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *deploymentClient) UpdateDeploymentStatus(ctx context.Context, obj *apps_v1.Deployment, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *deploymentClient) PatchDeploymentStatus(ctx context.Context, obj *apps_v1.Deployment, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides DeploymentClients for multiple clusters.
type MulticlusterDeploymentClient interface {
	// Cluster returns a DeploymentClient for the given cluster
	Cluster(cluster string) (DeploymentClient, error)
}

type multiclusterDeploymentClient struct {
	client multicluster.Client
}

func NewMulticlusterDeploymentClient(client multicluster.Client) MulticlusterDeploymentClient {
	return &multiclusterDeploymentClient{client: client}
}

func (m *multiclusterDeploymentClient) Cluster(cluster string) (DeploymentClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewDeploymentClient(client), nil
}

// Reader knows how to read and list ReplicaSets.
type ReplicaSetReader interface {
	// Get retrieves a ReplicaSet for the given object key
	GetReplicaSet(ctx context.Context, key client.ObjectKey) (*apps_v1.ReplicaSet, error)

	// List retrieves list of ReplicaSets for a given namespace and list options.
	ListReplicaSet(ctx context.Context, opts ...client.ListOption) (*apps_v1.ReplicaSetList, error)
}

// ReplicaSetTransitionFunction instructs the ReplicaSetWriter how to transition between an existing
// ReplicaSet object and a desired on an Upsert
type ReplicaSetTransitionFunction func(existing, desired *apps_v1.ReplicaSet) error

// Writer knows how to create, delete, and update ReplicaSets.
type ReplicaSetWriter interface {
	// Create saves the ReplicaSet object.
	CreateReplicaSet(ctx context.Context, obj *apps_v1.ReplicaSet, opts ...client.CreateOption) error

	// Delete deletes the ReplicaSet object.
	DeleteReplicaSet(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ReplicaSet object.
	UpdateReplicaSet(ctx context.Context, obj *apps_v1.ReplicaSet, opts ...client.UpdateOption) error

	// Patch patches the given ReplicaSet object.
	PatchReplicaSet(ctx context.Context, obj *apps_v1.ReplicaSet, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ReplicaSet objects matching the given options.
	DeleteAllOfReplicaSet(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ReplicaSet object.
	UpsertReplicaSet(ctx context.Context, obj *apps_v1.ReplicaSet, transitionFuncs ...ReplicaSetTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ReplicaSet object.
type ReplicaSetStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ReplicaSet object.
	UpdateReplicaSetStatus(ctx context.Context, obj *apps_v1.ReplicaSet, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given ReplicaSet object's subresource.
	PatchReplicaSetStatus(ctx context.Context, obj *apps_v1.ReplicaSet, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on ReplicaSets.
type ReplicaSetClient interface {
	ReplicaSetReader
	ReplicaSetWriter
	ReplicaSetStatusWriter
}

type replicaSetClient struct {
	client client.Client
}

func NewReplicaSetClient(client client.Client) *replicaSetClient {
	return &replicaSetClient{client: client}
}

func (c *replicaSetClient) GetReplicaSet(ctx context.Context, key client.ObjectKey) (*apps_v1.ReplicaSet, error) {
	obj := &apps_v1.ReplicaSet{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *replicaSetClient) ListReplicaSet(ctx context.Context, opts ...client.ListOption) (*apps_v1.ReplicaSetList, error) {
	list := &apps_v1.ReplicaSetList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *replicaSetClient) CreateReplicaSet(ctx context.Context, obj *apps_v1.ReplicaSet, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *replicaSetClient) DeleteReplicaSet(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &apps_v1.ReplicaSet{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *replicaSetClient) UpdateReplicaSet(ctx context.Context, obj *apps_v1.ReplicaSet, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *replicaSetClient) PatchReplicaSet(ctx context.Context, obj *apps_v1.ReplicaSet, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *replicaSetClient) DeleteAllOfReplicaSet(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &apps_v1.ReplicaSet{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *replicaSetClient) UpsertReplicaSet(ctx context.Context, obj *apps_v1.ReplicaSet, transitionFuncs ...ReplicaSetTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*apps_v1.ReplicaSet), desired.(*apps_v1.ReplicaSet)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *replicaSetClient) UpdateReplicaSetStatus(ctx context.Context, obj *apps_v1.ReplicaSet, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *replicaSetClient) PatchReplicaSetStatus(ctx context.Context, obj *apps_v1.ReplicaSet, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides ReplicaSetClients for multiple clusters.
type MulticlusterReplicaSetClient interface {
	// Cluster returns a ReplicaSetClient for the given cluster
	Cluster(cluster string) (ReplicaSetClient, error)
}

type multiclusterReplicaSetClient struct {
	client multicluster.Client
}

func NewMulticlusterReplicaSetClient(client multicluster.Client) MulticlusterReplicaSetClient {
	return &multiclusterReplicaSetClient{client: client}
}

func (m *multiclusterReplicaSetClient) Cluster(cluster string) (ReplicaSetClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewReplicaSetClient(client), nil
}

// Reader knows how to read and list DaemonSets.
type DaemonSetReader interface {
	// Get retrieves a DaemonSet for the given object key
	GetDaemonSet(ctx context.Context, key client.ObjectKey) (*apps_v1.DaemonSet, error)

	// List retrieves list of DaemonSets for a given namespace and list options.
	ListDaemonSet(ctx context.Context, opts ...client.ListOption) (*apps_v1.DaemonSetList, error)
}

// DaemonSetTransitionFunction instructs the DaemonSetWriter how to transition between an existing
// DaemonSet object and a desired on an Upsert
type DaemonSetTransitionFunction func(existing, desired *apps_v1.DaemonSet) error

// Writer knows how to create, delete, and update DaemonSets.
type DaemonSetWriter interface {
	// Create saves the DaemonSet object.
	CreateDaemonSet(ctx context.Context, obj *apps_v1.DaemonSet, opts ...client.CreateOption) error

	// Delete deletes the DaemonSet object.
	DeleteDaemonSet(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given DaemonSet object.
	UpdateDaemonSet(ctx context.Context, obj *apps_v1.DaemonSet, opts ...client.UpdateOption) error

	// Patch patches the given DaemonSet object.
	PatchDaemonSet(ctx context.Context, obj *apps_v1.DaemonSet, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all DaemonSet objects matching the given options.
	DeleteAllOfDaemonSet(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the DaemonSet object.
	UpsertDaemonSet(ctx context.Context, obj *apps_v1.DaemonSet, transitionFuncs ...DaemonSetTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a DaemonSet object.
type DaemonSetStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given DaemonSet object.
	UpdateDaemonSetStatus(ctx context.Context, obj *apps_v1.DaemonSet, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given DaemonSet object's subresource.
	PatchDaemonSetStatus(ctx context.Context, obj *apps_v1.DaemonSet, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on DaemonSets.
type DaemonSetClient interface {
	DaemonSetReader
	DaemonSetWriter
	DaemonSetStatusWriter
}

type daemonSetClient struct {
	client client.Client
}

func NewDaemonSetClient(client client.Client) *daemonSetClient {
	return &daemonSetClient{client: client}
}

func (c *daemonSetClient) GetDaemonSet(ctx context.Context, key client.ObjectKey) (*apps_v1.DaemonSet, error) {
	obj := &apps_v1.DaemonSet{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *daemonSetClient) ListDaemonSet(ctx context.Context, opts ...client.ListOption) (*apps_v1.DaemonSetList, error) {
	list := &apps_v1.DaemonSetList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *daemonSetClient) CreateDaemonSet(ctx context.Context, obj *apps_v1.DaemonSet, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *daemonSetClient) DeleteDaemonSet(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &apps_v1.DaemonSet{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *daemonSetClient) UpdateDaemonSet(ctx context.Context, obj *apps_v1.DaemonSet, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *daemonSetClient) PatchDaemonSet(ctx context.Context, obj *apps_v1.DaemonSet, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *daemonSetClient) DeleteAllOfDaemonSet(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &apps_v1.DaemonSet{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *daemonSetClient) UpsertDaemonSet(ctx context.Context, obj *apps_v1.DaemonSet, transitionFuncs ...DaemonSetTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*apps_v1.DaemonSet), desired.(*apps_v1.DaemonSet)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *daemonSetClient) UpdateDaemonSetStatus(ctx context.Context, obj *apps_v1.DaemonSet, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *daemonSetClient) PatchDaemonSetStatus(ctx context.Context, obj *apps_v1.DaemonSet, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides DaemonSetClients for multiple clusters.
type MulticlusterDaemonSetClient interface {
	// Cluster returns a DaemonSetClient for the given cluster
	Cluster(cluster string) (DaemonSetClient, error)
}

type multiclusterDaemonSetClient struct {
	client multicluster.Client
}

func NewMulticlusterDaemonSetClient(client multicluster.Client) MulticlusterDaemonSetClient {
	return &multiclusterDaemonSetClient{client: client}
}

func (m *multiclusterDaemonSetClient) Cluster(cluster string) (DaemonSetClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewDaemonSetClient(client), nil
}

// Reader knows how to read and list StatefulSets.
type StatefulSetReader interface {
	// Get retrieves a StatefulSet for the given object key
	GetStatefulSet(ctx context.Context, key client.ObjectKey) (*apps_v1.StatefulSet, error)

	// List retrieves list of StatefulSets for a given namespace and list options.
	ListStatefulSet(ctx context.Context, opts ...client.ListOption) (*apps_v1.StatefulSetList, error)
}

// StatefulSetTransitionFunction instructs the StatefulSetWriter how to transition between an existing
// StatefulSet object and a desired on an Upsert
type StatefulSetTransitionFunction func(existing, desired *apps_v1.StatefulSet) error

// Writer knows how to create, delete, and update StatefulSets.
type StatefulSetWriter interface {
	// Create saves the StatefulSet object.
	CreateStatefulSet(ctx context.Context, obj *apps_v1.StatefulSet, opts ...client.CreateOption) error

	// Delete deletes the StatefulSet object.
	DeleteStatefulSet(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given StatefulSet object.
	UpdateStatefulSet(ctx context.Context, obj *apps_v1.StatefulSet, opts ...client.UpdateOption) error

	// Patch patches the given StatefulSet object.
	PatchStatefulSet(ctx context.Context, obj *apps_v1.StatefulSet, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all StatefulSet objects matching the given options.
	DeleteAllOfStatefulSet(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the StatefulSet object.
	UpsertStatefulSet(ctx context.Context, obj *apps_v1.StatefulSet, transitionFuncs ...StatefulSetTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a StatefulSet object.
type StatefulSetStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given StatefulSet object.
	UpdateStatefulSetStatus(ctx context.Context, obj *apps_v1.StatefulSet, opts ...client.SubResourceUpdateOption) error

	// Patch patches the given StatefulSet object's subresource.
	PatchStatefulSetStatus(ctx context.Context, obj *apps_v1.StatefulSet, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on StatefulSets.
type StatefulSetClient interface {
	StatefulSetReader
	StatefulSetWriter
	StatefulSetStatusWriter
}

type statefulSetClient struct {
	client client.Client
}

func NewStatefulSetClient(client client.Client) *statefulSetClient {
	return &statefulSetClient{client: client}
}

func (c *statefulSetClient) GetStatefulSet(ctx context.Context, key client.ObjectKey) (*apps_v1.StatefulSet, error) {
	obj := &apps_v1.StatefulSet{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *statefulSetClient) ListStatefulSet(ctx context.Context, opts ...client.ListOption) (*apps_v1.StatefulSetList, error) {
	list := &apps_v1.StatefulSetList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *statefulSetClient) CreateStatefulSet(ctx context.Context, obj *apps_v1.StatefulSet, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *statefulSetClient) DeleteStatefulSet(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &apps_v1.StatefulSet{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *statefulSetClient) UpdateStatefulSet(ctx context.Context, obj *apps_v1.StatefulSet, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *statefulSetClient) PatchStatefulSet(ctx context.Context, obj *apps_v1.StatefulSet, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *statefulSetClient) DeleteAllOfStatefulSet(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &apps_v1.StatefulSet{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *statefulSetClient) UpsertStatefulSet(ctx context.Context, obj *apps_v1.StatefulSet, transitionFuncs ...StatefulSetTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*apps_v1.StatefulSet), desired.(*apps_v1.StatefulSet)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *statefulSetClient) UpdateStatefulSetStatus(ctx context.Context, obj *apps_v1.StatefulSet, opts ...client.SubResourceUpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *statefulSetClient) PatchStatefulSetStatus(ctx context.Context, obj *apps_v1.StatefulSet, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides StatefulSetClients for multiple clusters.
type MulticlusterStatefulSetClient interface {
	// Cluster returns a StatefulSetClient for the given cluster
	Cluster(cluster string) (StatefulSetClient, error)
}

type multiclusterStatefulSetClient struct {
	client multicluster.Client
}

func NewMulticlusterStatefulSetClient(client multicluster.Client) MulticlusterStatefulSetClient {
	return &multiclusterStatefulSetClient{client: client}
}

func (m *multiclusterStatefulSetClient) Cluster(cluster string) (StatefulSetClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewStatefulSetClient(client), nil
}
