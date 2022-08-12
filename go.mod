module github.com/solo-io/external-apis

go 1.16

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b

	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d

	k8s.io/api => k8s.io/api v0.22.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.22.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.2
	k8s.io/apiserver => k8s.io/apiserver v0.22.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.22.2
	k8s.io/client-go => k8s.io/client-go v0.22.2
	k8s.io/kubectl => k8s.io/kubectl v0.22.2
)

require (
	github.com/aws/aws-app-mesh-controller-for-k8s v1.1.1
	github.com/cilium/cilium v1.11.6
	github.com/golang/mock v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.1.1
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/solo-io/skv2 v0.22.8
	istio.io/client-go v1.13.7
	istio.io/istio v0.0.0-20220812032556-81e2be94452f
	k8s.io/api v0.24.2
	k8s.io/apiextensions-apiserver v0.24.2
	k8s.io/apimachinery v0.24.2
	k8s.io/client-go v0.24.2
	sigs.k8s.io/controller-runtime v0.12.2
)
