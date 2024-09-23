package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
	gwapiv1 "sigs.k8s.io/gateway-api/apis/v1"
)

const (
	k8sGwApiRoot = apiRoot + "/k8s-gw"
	k8sGwModule  = "sigs.k8s.io/gateway-api"
)

func init() {
	Groups = append(Groups, k8sGwGroups()...)
}

func k8sGwGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: gwapiv1.SchemeGroupVersion,
			Module:       k8sGwModule,
			Resources: []model.Resource{
				{
					Kind: "GatewayClass",
				},
				{
					Kind: "Gateway",
				},
				{
					Kind: "HTTPRoute",
				},
			},
			CustomTypesImportPath: k8sGwModule + "/apis/v1",
			ApiRoot:               k8sGwApiRoot + "/apis",
		},
	}
}
