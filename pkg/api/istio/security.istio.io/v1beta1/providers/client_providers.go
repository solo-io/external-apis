// Code generated by skv2. DO NOT EDIT.

package v1beta1

import (
	security_istio_io_v1beta1 "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for AuthorizationPolicyClient from Clientset
func AuthorizationPolicyClientFromClientsetProvider(clients security_istio_io_v1beta1.Clientset) security_istio_io_v1beta1.AuthorizationPolicyClient {
	return clients.AuthorizationPolicies()
}

// Provider for AuthorizationPolicy Client from Client
func AuthorizationPolicyClientProvider(client client.Client) security_istio_io_v1beta1.AuthorizationPolicyClient {
	return security_istio_io_v1beta1.NewAuthorizationPolicyClient(client)
}

type AuthorizationPolicyClientFactory func(client client.Client) security_istio_io_v1beta1.AuthorizationPolicyClient

func AuthorizationPolicyClientFactoryProvider() AuthorizationPolicyClientFactory {
	return AuthorizationPolicyClientProvider
}

type AuthorizationPolicyClientFromConfigFactory func(cfg *rest.Config) (security_istio_io_v1beta1.AuthorizationPolicyClient, error)

func AuthorizationPolicyClientFromConfigFactoryProvider() AuthorizationPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (security_istio_io_v1beta1.AuthorizationPolicyClient, error) {
		clients, err := security_istio_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.AuthorizationPolicies(), nil
	}
}

// Provider for PeerAuthenticationClient from Clientset
func PeerAuthenticationClientFromClientsetProvider(clients security_istio_io_v1beta1.Clientset) security_istio_io_v1beta1.PeerAuthenticationClient {
	return clients.PeerAuthentications()
}

// Provider for PeerAuthentication Client from Client
func PeerAuthenticationClientProvider(client client.Client) security_istio_io_v1beta1.PeerAuthenticationClient {
	return security_istio_io_v1beta1.NewPeerAuthenticationClient(client)
}

type PeerAuthenticationClientFactory func(client client.Client) security_istio_io_v1beta1.PeerAuthenticationClient

func PeerAuthenticationClientFactoryProvider() PeerAuthenticationClientFactory {
	return PeerAuthenticationClientProvider
}

type PeerAuthenticationClientFromConfigFactory func(cfg *rest.Config) (security_istio_io_v1beta1.PeerAuthenticationClient, error)

func PeerAuthenticationClientFromConfigFactoryProvider() PeerAuthenticationClientFromConfigFactory {
	return func(cfg *rest.Config) (security_istio_io_v1beta1.PeerAuthenticationClient, error) {
		clients, err := security_istio_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.PeerAuthentications(), nil
	}
}
