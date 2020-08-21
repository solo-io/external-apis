package codegen

// TODO: Uncomment to reenable generating linkerd types once they bump to kube 1.18.x
// import (
// 	"github.com/linkerd/linkerd2/controller/gen/apis/serviceprofile/v1alpha2"
// 	"github.com/solo-io/skv2/codegen/model"
// )

// func init() {
// 	Groups = append(Groups, linkerdGroups()...)
// }

// const (
// 	linkerdApiRoot = apiRoot + "/linkerd"
// )

// func linkerdGroups() []model.Group {
// 	return []model.Group{
// 		{
// 			GroupVersion: v1alpha2.SchemeGroupVersion,
// 			Module:       "github.com/linkerd/linkerd2",
// 			Resources: []model.Resource{
// 				{
// 					Kind: "ServiceProfile",
// 				},
// 			},
// 			CustomTypesImportPath: "github.com/linkerd/linkerd2/controller/gen/apis/serviceprofile/v1alpha2",
// 			ApiRoot:               linkerdApiRoot,
// 		},
// 	}
// }
