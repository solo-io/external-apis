module github.com/solo-io/external-apis

go 1.16

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b

	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d

	istio.io/client-go => istio.io/client-go v1.12.2

	// Using private fork of controller-tools. See commit msg for more context
	// as to why we are using a private fork.
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
//github.com/envoyproxy/go-control-plane => github.com/envoyproxy/go-control-plane v0.9.0
//github.com/envoyproxy/go-control-plane => 	github.com/envoyproxy/go-control-plane v0.10.2-0.20220420171917-689c2bccf0ec
//github.com/envoyproxy/go-control-plane => 	github.com/envoyproxy/go-control-plane v0.10.2-0.20220325020618-49ff273808a1
// Using private fork of controller-tools. See commit msg for more context
// as to why we are using a private fork.
//sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
)

require (
	github.com/aws/aws-app-mesh-controller-for-k8s v1.1.1
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/golang/mock v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_model v0.2.1-0.20210607210712-147c58e9608a // indirect
	github.com/rotisserie/eris v0.1.1
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/solo-io/skv2 v0.22.8
	istio.io/client-go v1.12.2
	istio.io/istio v0.0.0-20220606165002-e8b889bc440d
	k8s.io/api v0.24.1
	k8s.io/apiextensions-apiserver v0.24.0
	k8s.io/apimachinery v0.24.1
	k8s.io/client-go v0.24.1
	sigs.k8s.io/controller-runtime v0.12.1
)
