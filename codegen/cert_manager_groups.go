package codegen

import (
	v1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	"github.com/solo-io/skv2/codegen/model"
)

func init() {
	Groups = append(Groups, certManagerGroups()...)
}

const (
	certManagerApiRoot      = apiRoot + "/certmanager"
	certManagerModule       = "github.com/jetstack/cert-manager"
	certManagerTypesPackage = certManagerModule + "/pkg/apis/certmanager/v1"
)

func certManagerGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: v1.SchemeGroupVersion,
			Module:       certManagerModule,
			Resources: []model.Resource{
				{
					Kind: "CertificateRequest",
				},
			},
			CustomTypesImportPath: certManagerTypesPackage,
			ApiRoot:               certManagerApiRoot,
		},
	}
}
