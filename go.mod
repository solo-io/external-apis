module github.com/solo-io/external-apis

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.1
	k8s.io/client-go => k8s.io/client-go v0.17.1
)

require (
	github.com/aws/aws-sdk-go v1.30.15 // indirect
	github.com/gertd/go-pluralize v0.1.1 // indirect
	github.com/gobuffalo/envy v1.8.1 // indirect
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/golang/mock v1.4.4
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/gophercloud/gophercloud v0.2.0 // indirect
	github.com/linkerd/linkerd2 v0.5.1-0.20200402173539-fee70c064bc0
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rogpeppe/go-internal v1.5.2 // indirect
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/solo-io/go-utils v0.15.2 // indirect
	github.com/solo-io/protoc-gen-ext v0.0.9 // indirect
	github.com/solo-io/skv2 v0.8.1
	github.com/solo-io/solo-kit v0.13.8 // indirect
	golang.org/x/mod v0.3.0 // indirect
	gonum.org/v1/netlib v0.0.0-20191031114514-eccb95939662 // indirect
	istio.io/client-go v0.0.0-20200807223845-61c70ad04ec9
	k8s.io/api v0.18.2
	k8s.io/apiextensions-apiserver v0.18.2
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
	rsc.io/quote/v3 v3.1.0 // indirect
	sigs.k8s.io/aws-iam-authenticator v0.5.0 // indirect
	sigs.k8s.io/controller-runtime v0.5.6
)
