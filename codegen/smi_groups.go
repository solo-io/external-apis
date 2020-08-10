package codegen

import (
	access_v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	specs_v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	split_v1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	split_v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha3"
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
			GroupVersion: split_v1alpha1.SchemeGroupVersion,
			Module:       "github.com/servicemeshinterface/smi-sdk-go",
			Resources: []model.Resource{
				{
					Kind: "TrafficSplit",
				},
			},
			CustomTypesImportPath: "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1",
			ApiRoot:               smiApiRoot,
		},
		{
			GroupVersion: split_v1alpha3.SchemeGroupVersion,
			Module:       "github.com/servicemeshinterface/smi-sdk-go",
			Resources: []model.Resource{
				{
					Kind: "TrafficSplit",
				},
			},
			CustomTypesImportPath: "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha3",
			ApiRoot:               smiApiRoot,
		},
		{
			GroupVersion: access_v1alpha2.SchemeGroupVersion,
			Module:       "github.com/servicemeshinterface/smi-sdk-go",
			Resources: []model.Resource{
				{
					Kind: "TrafficTarget",
				},
			},
			CustomTypesImportPath: "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2",
			ApiRoot:               smiApiRoot,
		},
		{
			GroupVersion: specs_v1alpha3.SchemeGroupVersion,
			Module:       "github.com/servicemeshinterface/smi-sdk-go",
			Resources: []model.Resource{
				{
					Kind: "HTTPRouteGroup",
				},
			},
			CustomTypesImportPath: "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3",
			ApiRoot:               smiApiRoot,
		},
	}
}
