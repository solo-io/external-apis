package codegen

import (
	ciliumv2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	Groups = append(Groups, ciliumGroups()...)
}

const (
	ciliumApiRoot = apiRoot + "/cilium"
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
