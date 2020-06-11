package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/security/v1beta1"
)

func init() {
	groups = append(groups, istioGroups()...)
}

const (
	istioApiRoot = apiRoot + "/istio"
	istioModule  = "istio.io/client-go"
)

func istioGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: v1alpha3.SchemeGroupVersion,
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
					Kind: "VirtualService",
				},
			},
			CustomTypesImportPath: "istio.io/client-go/pkg/apis/networking/v1alpha3",
			ApiRoot:               istioApiRoot,
		},
		{
			GroupVersion: v1beta1.SchemeGroupVersion,
			Module:       istioModule,
			Resources: []model.Resource{
				{
					Kind: "AuthorizationPolicy",
				},
			},
			CustomTypesImportPath: "istio.io/client-go/pkg/apis/security/v1beta1",
			ApiRoot:               istioApiRoot,
		},
	}
}
