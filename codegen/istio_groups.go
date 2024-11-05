package codegen

import (
	operator "github.com/jehawley/istio/operator/pkg/apis/istio/v1alpha1"
	"github.com/solo-io/skv2/codegen/model"
	extensionsv1alpha1 "istio.io/client-go/pkg/apis/extensions/v1alpha1"
	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	networkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	"istio.io/client-go/pkg/apis/security/v1beta1"
	istiotelemetryv1alpha1 "istio.io/client-go/pkg/apis/telemetry/v1alpha1"
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
			GroupVersion: networkingv1beta1.SchemeGroupVersion,
			Module:       istioModule,
			Resources: []model.Resource{
				{
					Kind: "DestinationRule",
				},
				{
					Kind: "Gateway",
				},
				{
					Kind: "ProxyConfig",
				},
				{
					Kind: "ServiceEntry",
				},
				{
					Kind: "WorkloadEntry",
				},
				{
					Kind: "WorkloadGroup",
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
			PointerSlices:         true,
		},
		{
			GroupVersion: networkingv1alpha3.SchemeGroupVersion,
			Module:       istioModule,
			Resources: []model.Resource{
				{
					Kind: "EnvoyFilter",
				},
			},
			CustomTypesImportPath: "istio.io/client-go/pkg/apis/networking/v1alpha3",
			ApiRoot:               istioApiRoot,
			PointerSlices:         true,
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
				{
					Kind: "RequestAuthentication",
				},
			},
			CustomTypesImportPath: "istio.io/client-go/pkg/apis/security/v1beta1",
			ApiRoot:               istioApiRoot,
			PointerSlices:         true,
		},
		{
			GroupVersion: operator.SchemeGroupVersion,
			Module:       "install.istio.io",
			Resources: []model.Resource{
				{
					Kind: "IstioOperator",
				},
			},
			CustomTypesImportPath: "github.com/jehawley/istio/operator/pkg/apis/istio/v1alpha1",
			ApiRoot:               istioApiRoot,
		},
		{
			GroupVersion: istiotelemetryv1alpha1.SchemeGroupVersion,
			Module:       istioModule,
			Resources: []model.Resource{
				{
					Kind: "Telemetry",
				},
			},
			CustomTypesImportPath: "istio.io/client-go/pkg/apis/telemetry/v1alpha1",
			ApiRoot:               istioApiRoot,
			PointerSlices:         true,
		},
		{
			GroupVersion: extensionsv1alpha1.SchemeGroupVersion,
			Module:       istioModule,
			Resources: []model.Resource{
				{
					Kind: "WasmPlugin",
				},
			},
			CustomTypesImportPath: "istio.io/client-go/pkg/apis/extensions/v1alpha1",
			ApiRoot:               istioApiRoot,
			PointerSlices:         true,
		},
	}
}
