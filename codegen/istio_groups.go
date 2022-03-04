package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
	networking "istio.io/client-go/pkg/apis/networking/v1beta1"
	"istio.io/client-go/pkg/apis/security/v1beta1"
	operator "istio.io/istio/operator/pkg/apis/istio/v1alpha1"
)

func init() {
	Groups = append(Groups, istioGroups()...)
}

const (
	istioApiRoot = apiRoot + "/istio"
	istioModule  = "istio.io/client-go"
)

func istioGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: networking.SchemeGroupVersion,
			Module:       istioModule,
			Resources: []model.Resource{
				{
					Kind: "DestinationRule",
				},
				{
					Kind: "EnvoyFilter",
				},
				{
					Kind: "Gateway",
				},
				{
					Kind: "ServiceEntry",
				},
				{
					Kind: "WorkloadEntry",
				},
				{
					Kind: "VirtualService",
				},
				{
					Kind: "Sidecar",
				},
			},
			CustomTypesImportPath: "istio.io/client-go/pkg/apis/networking/v1beta1",
			ApiRoot:               istioApiRoot,
		},
		{
			GroupVersion: v1beta1.SchemeGroupVersion,
			Module:       istioModule,
			Resources: []model.Resource{
				{
					Kind: "AuthorizationPolicy",
				},
				{
					Kind: "PeerAuthentication",
				},
			},
			CustomTypesImportPath: "istio.io/client-go/pkg/apis/security/v1beta1",
			ApiRoot:               istioApiRoot,
		},
		{
			GroupVersion: operator.SchemeGroupVersion,
			Module:       "install.istio.io",
			Resources: []model.Resource{
				{
					Kind: "IstioOperator",
				},
			},
			CustomTypesImportPath: "istio.io/istio/operator/pkg/apis/istio/v1alpha1",
			ApiRoot:               istioApiRoot,
		},
	}
}
