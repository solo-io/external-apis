module github.com/solo-io/external-apis

go 1.16

replace (
	// pinned to solo-io's fork of cue version 95a50cebaffb4bdba8c544601d8fb867990ad1ad
	cuelang.org/go => github.com/solo-io/cue v0.4.1-0.20210622213027-95a50cebaffb

	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
)

require (
	github.com/aws/aws-app-mesh-controller-for-k8s v1.1.1
	github.com/golang/mock v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.1.1
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/solo-io/skv2 v0.19.8
	istio.io/client-go v1.12.0-alpha.5.0.20211124144007-2c98ed70c967
	istio.io/istio v0.0.0-20211202012955-9928258d2089
	k8s.io/api v0.22.4
	k8s.io/apiextensions-apiserver v0.22.4
	k8s.io/apimachinery v0.22.4
	k8s.io/client-go v0.22.4
	sigs.k8s.io/controller-runtime v0.10.3
)
