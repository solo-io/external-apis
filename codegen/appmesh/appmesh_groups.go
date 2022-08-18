package appmesh

import (
	"github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	"github.com/solo-io/external-apis/codegen"
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	codegen.Groups = append(codegen.Groups, appmeshGroups()...)
}

const (
	appmeshApiRoot      = codegen.ApiRoot + "/appmesh"
	appmeshModule       = "github.com/aws/aws-app-mesh-controller-for-k8s"
	appmeshTypesPackage = appmeshModule + "/apis/appmesh/v1beta2"
)

func appmeshGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: v1beta2.GroupVersion,
			Module:       appmeshModule,
			Resources: []model.Resource{
				{
					Kind: "Mesh",
				},
				{
					Kind: "VirtualService",
				},
				{
					Kind: "VirtualNode",
				},
				{
					Kind: "VirtualRouter",
				},
				{
					Kind: "VirtualGateway",
				},
				{
					Kind: "GatewayRoute",
				},
			},
			CustomTypesImportPath: appmeshTypesPackage,
			ApiRoot:               appmeshApiRoot,
		},
	}
}
