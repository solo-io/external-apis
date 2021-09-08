module github.com/solo-io/external-apis

go 1.14

replace (
	// pinned to solo-io's fork of cue version 95a50cebaffb4bdba8c544601d8fb867990ad1ad
	cuelang.org/go => github.com/solo-io/cue v0.4.1-0.20210622213027-95a50cebaffb

	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

	github.com/golang/mock => github.com/golang/mock v1.4.4
)

require (
	github.com/aws/aws-app-mesh-controller-for-k8s v1.1.1
	github.com/golang/mock v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.1.1
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/solo-io/skv2 v0.19.8-0.20210908202014-78e62d569a05
	istio.io/client-go v1.9.5-0.20210607162355-6a6709ba5473
	istio.io/istio v0.0.0-20210903183209-9d4e257a8202
	k8s.io/api v0.20.1
	k8s.io/apiextensions-apiserver v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
	sigs.k8s.io/controller-runtime v0.7.0
)
