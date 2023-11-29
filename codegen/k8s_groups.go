package codegen

import (
	"github.com/solo-io/skv2/codegen/model"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	gatewayapiv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

const (
	k8sApiRoot = apiRoot + "/k8s"
	k8sModule  = "k8s.io/api"
)

func init() {
	Groups = append(Groups, k8sGroups()...)
}

func k8sGroups() []model.Group {
	return []model.Group{
		{
			GroupVersion: corev1.SchemeGroupVersion,
			Module:       k8sModule,
			Resources: []model.Resource{
				{
					Kind: "Secret",
				},
				{
					Kind: "ServiceAccount",
				},
				{
					Kind: "ConfigMap",
				},
				{
					Kind: "Service",
				},
				{
					Kind: "Pod",
				},
				{
					Kind: "Endpoints",
				},
				{
					Kind:          "Namespace",
					ClusterScoped: true,
				},
				{
					Kind:          "Node",
					ClusterScoped: true,
				},
			},
			CustomTypesImportPath: "k8s.io/api/core/v1",
			ApiRoot:               k8sApiRoot + "/core",
		},
		{
			GroupVersion: appsv1.SchemeGroupVersion,
			Module:       k8sModule,
			Resources: []model.Resource{
				{
					Kind: "Deployment",
				},
				{
					Kind: "ReplicaSet",
				},
				{
					Kind: "DaemonSet",
				},
				{
					Kind: "StatefulSet",
				},
			},
			CustomTypesImportPath: "k8s.io/api/apps/v1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: batchv1.SchemeGroupVersion,
			Module:       k8sModule,
			Resources: []model.Resource{
				{
					Kind: "Job",
				},
			},
			CustomTypesImportPath: "k8s.io/api/batch/v1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: admissionregistrationv1.SchemeGroupVersion,
			Module:       "k8s.io/apiextensions-apiserver",
			Resources: []model.Resource{
				{
					Kind: "ValidatingWebhookConfiguration",
				},
				{
					Kind: "MutatingWebhookConfiguration",
				},
			},
			CustomTypesImportPath: "k8s.io/api/admissionregistration/v1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: rbacv1.SchemeGroupVersion,
			Resources: []model.Resource{
				{
					Kind: "Role",
				},
				{
					Kind: "RoleBinding",
				},
				{
					Kind:          "ClusterRole",
					ClusterScoped: true,
				},
				{
					Kind:          "ClusterRoleBinding",
					ClusterScoped: true,
				},
			},
			CustomTypesImportPath: "k8s.io/api/rbac/v1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: networkingv1.SchemeGroupVersion,
			Resources: []model.Resource{
				{
					Kind: "NetworkPolicy",
				},
			},
			CustomTypesImportPath: "k8s.io/api/networking/v1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: certificatesv1beta1.SchemeGroupVersion,
			Resources: []model.Resource{
				{
					Kind: "CertificateSigningRequest",
				},
			},
			CustomTypesImportPath: "k8s.io/api/certificates/v1beta1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: apiextensionsv1beta1.SchemeGroupVersion,
			Module:       "k8s.io/apiextensions-apiserver",
			Resources: []model.Resource{
				{
					Kind:          "CustomResourceDefinition",
					ClusterScoped: true,
				},
			},
			CustomTypesImportPath: "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: apiextensionsv1.SchemeGroupVersion,
			Module:       "k8s.io/apiextensions-apiserver",
			Resources: []model.Resource{
				{
					Kind:          "CustomResourceDefinition",
					ClusterScoped: true,
				},
			},
			CustomTypesImportPath: "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: schema.GroupVersion{
				Group:   gatewayapiv1beta1.GroupName,
				Version: gatewayapiv1beta1.GroupVersion.Version,
			},
			Module: "sigs.k8s.io/gateway-api",
			Resources: []model.Resource{
				{
					Kind: "Gateway",
				},
				{
					Kind:          "GatewayClass",
					ClusterScoped: true,
				},
				{
					Kind: "HTTPRoute",
				},
			},
			CustomTypesImportPath: "sigs.k8s.io/gateway-api/apis/v1",
			ApiRoot:               k8sApiRoot,
		},
	}
}
