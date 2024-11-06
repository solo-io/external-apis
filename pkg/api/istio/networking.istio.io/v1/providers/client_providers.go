// Code generated by skv2. DO NOT EDIT.

package v1

import (
	networking_istio_io_v1 "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for DestinationRuleClient from Clientset
func DestinationRuleClientFromClientsetProvider(clients networking_istio_io_v1.Clientset) networking_istio_io_v1.DestinationRuleClient {
	return clients.DestinationRules()
}

// Provider for DestinationRule Client from Client
func DestinationRuleClientProvider(client client.Client) networking_istio_io_v1.DestinationRuleClient {
	return networking_istio_io_v1.NewDestinationRuleClient(client)
}

type DestinationRuleClientFactory func(client client.Client) networking_istio_io_v1.DestinationRuleClient

func DestinationRuleClientFactoryProvider() DestinationRuleClientFactory {
	return DestinationRuleClientProvider
}

type DestinationRuleClientFromConfigFactory func(cfg *rest.Config) (networking_istio_io_v1.DestinationRuleClient, error)

func DestinationRuleClientFromConfigFactoryProvider() DestinationRuleClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_istio_io_v1.DestinationRuleClient, error) {
		clients, err := networking_istio_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.DestinationRules(), nil
	}
}

// Provider for GatewayClient from Clientset
func GatewayClientFromClientsetProvider(clients networking_istio_io_v1.Clientset) networking_istio_io_v1.GatewayClient {
	return clients.Gateways()
}

// Provider for Gateway Client from Client
func GatewayClientProvider(client client.Client) networking_istio_io_v1.GatewayClient {
	return networking_istio_io_v1.NewGatewayClient(client)
}

type GatewayClientFactory func(client client.Client) networking_istio_io_v1.GatewayClient

func GatewayClientFactoryProvider() GatewayClientFactory {
	return GatewayClientProvider
}

type GatewayClientFromConfigFactory func(cfg *rest.Config) (networking_istio_io_v1.GatewayClient, error)

func GatewayClientFromConfigFactoryProvider() GatewayClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_istio_io_v1.GatewayClient, error) {
		clients, err := networking_istio_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Gateways(), nil
	}
}

// Provider for ServiceEntryClient from Clientset
func ServiceEntryClientFromClientsetProvider(clients networking_istio_io_v1.Clientset) networking_istio_io_v1.ServiceEntryClient {
	return clients.ServiceEntries()
}

// Provider for ServiceEntry Client from Client
func ServiceEntryClientProvider(client client.Client) networking_istio_io_v1.ServiceEntryClient {
	return networking_istio_io_v1.NewServiceEntryClient(client)
}

type ServiceEntryClientFactory func(client client.Client) networking_istio_io_v1.ServiceEntryClient

func ServiceEntryClientFactoryProvider() ServiceEntryClientFactory {
	return ServiceEntryClientProvider
}

type ServiceEntryClientFromConfigFactory func(cfg *rest.Config) (networking_istio_io_v1.ServiceEntryClient, error)

func ServiceEntryClientFromConfigFactoryProvider() ServiceEntryClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_istio_io_v1.ServiceEntryClient, error) {
		clients, err := networking_istio_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ServiceEntries(), nil
	}
}

// Provider for WorkloadEntryClient from Clientset
func WorkloadEntryClientFromClientsetProvider(clients networking_istio_io_v1.Clientset) networking_istio_io_v1.WorkloadEntryClient {
	return clients.WorkloadEntries()
}

// Provider for WorkloadEntry Client from Client
func WorkloadEntryClientProvider(client client.Client) networking_istio_io_v1.WorkloadEntryClient {
	return networking_istio_io_v1.NewWorkloadEntryClient(client)
}

type WorkloadEntryClientFactory func(client client.Client) networking_istio_io_v1.WorkloadEntryClient

func WorkloadEntryClientFactoryProvider() WorkloadEntryClientFactory {
	return WorkloadEntryClientProvider
}

type WorkloadEntryClientFromConfigFactory func(cfg *rest.Config) (networking_istio_io_v1.WorkloadEntryClient, error)

func WorkloadEntryClientFromConfigFactoryProvider() WorkloadEntryClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_istio_io_v1.WorkloadEntryClient, error) {
		clients, err := networking_istio_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.WorkloadEntries(), nil
	}
}

// Provider for WorkloadGroupClient from Clientset
func WorkloadGroupClientFromClientsetProvider(clients networking_istio_io_v1.Clientset) networking_istio_io_v1.WorkloadGroupClient {
	return clients.WorkloadGroups()
}

// Provider for WorkloadGroup Client from Client
func WorkloadGroupClientProvider(client client.Client) networking_istio_io_v1.WorkloadGroupClient {
	return networking_istio_io_v1.NewWorkloadGroupClient(client)
}

type WorkloadGroupClientFactory func(client client.Client) networking_istio_io_v1.WorkloadGroupClient

func WorkloadGroupClientFactoryProvider() WorkloadGroupClientFactory {
	return WorkloadGroupClientProvider
}

type WorkloadGroupClientFromConfigFactory func(cfg *rest.Config) (networking_istio_io_v1.WorkloadGroupClient, error)

func WorkloadGroupClientFromConfigFactoryProvider() WorkloadGroupClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_istio_io_v1.WorkloadGroupClient, error) {
		clients, err := networking_istio_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.WorkloadGroups(), nil
	}
}

// Provider for VirtualServiceClient from Clientset
func VirtualServiceClientFromClientsetProvider(clients networking_istio_io_v1.Clientset) networking_istio_io_v1.VirtualServiceClient {
	return clients.VirtualServices()
}

// Provider for VirtualService Client from Client
func VirtualServiceClientProvider(client client.Client) networking_istio_io_v1.VirtualServiceClient {
	return networking_istio_io_v1.NewVirtualServiceClient(client)
}

type VirtualServiceClientFactory func(client client.Client) networking_istio_io_v1.VirtualServiceClient

func VirtualServiceClientFactoryProvider() VirtualServiceClientFactory {
	return VirtualServiceClientProvider
}

type VirtualServiceClientFromConfigFactory func(cfg *rest.Config) (networking_istio_io_v1.VirtualServiceClient, error)

func VirtualServiceClientFromConfigFactoryProvider() VirtualServiceClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_istio_io_v1.VirtualServiceClient, error) {
		clients, err := networking_istio_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualServices(), nil
	}
}

// Provider for SidecarClient from Clientset
func SidecarClientFromClientsetProvider(clients networking_istio_io_v1.Clientset) networking_istio_io_v1.SidecarClient {
	return clients.Sidecars()
}

// Provider for Sidecar Client from Client
func SidecarClientProvider(client client.Client) networking_istio_io_v1.SidecarClient {
	return networking_istio_io_v1.NewSidecarClient(client)
}

type SidecarClientFactory func(client client.Client) networking_istio_io_v1.SidecarClient

func SidecarClientFactoryProvider() SidecarClientFactory {
	return SidecarClientProvider
}

type SidecarClientFromConfigFactory func(cfg *rest.Config) (networking_istio_io_v1.SidecarClient, error)

func SidecarClientFromConfigFactoryProvider() SidecarClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_istio_io_v1.SidecarClient, error) {
		clients, err := networking_istio_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Sidecars(), nil
	}
}
