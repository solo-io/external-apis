package codegen

import (
	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	Groups = append(Groups, smiGroups()...)
}

const (
	smiApiRoot = apiRoot + "/smi"
)

func smiGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: v1alpha1.SchemeGroupVersion,
			Module:       "github.com/servicemeshinterface/smi-sdk-go",
			Resources: []model.Resource{
				{
					Kind: "TrafficSplit",
				},
			},
			CustomTypesImportPath: "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1",
			ApiRoot:               smiApiRoot,
		},
	}
}
