package cillium

import (
	ciliumv2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/solo-io/external-apis/codegen"
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	codegen.Groups = append(codegen.Groups, ciliumGroups()...)
}

const (
	ciliumApiRoot = codegen.ApiRoot + "/cilium"
	ciliumModule  = "github.com/cilium/cilium"
)

func ciliumGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: ciliumv2.SchemeGroupVersion,
			Module:       ciliumModule,
			Resources: []model.Resource{
				{
					Kind: "CiliumNetworkPolicy",
				},
			},
			CustomTypesImportPath: "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2",
			ApiRoot:               ciliumApiRoot,
		},
	}
}
