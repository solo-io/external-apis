package smi

import (
	accessv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	specsv1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	splitv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"
	splitv1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha3"
	"github.com/solo-io/external-apis/codegen"
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	codegen.Groups = append(codegen.Groups, smiGroups()...)
}

const (
	smiApiRoot = codegen.ApiRoot + "/smi"
)

func smiGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: splitv1alpha1.SchemeGroupVersion,
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
			GroupVersion: splitv1alpha2.SchemeGroupVersion,
			Module:       "github.com/servicemeshinterface/smi-sdk-go",
			Resources: []model.Resource{
				{
					Kind: "TrafficSplit",
				},
			},
			CustomTypesImportPath: "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2",
			ApiRoot:               smiApiRoot,
		},
		{
			GroupVersion: splitv1alpha3.SchemeGroupVersion,
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
			GroupVersion: accessv1alpha2.SchemeGroupVersion,
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
			GroupVersion: specsv1alpha3.SchemeGroupVersion,
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
