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
	github.com/golang/mock v1.4.4-0.20200406172829-6d816de489c1
	github.com/linkerd/linkerd2 v0.5.1-0.20200402173539-fee70c064bc0
	github.com/pkg/errors v0.9.1
	github.com/servicemeshinterface/smi-sdk-go v0.3.0
	github.com/solo-io/skv2 v0.4.0
	istio.io/client-go v0.0.0-20200610222318-1cfead1f1938
	k8s.io/api v0.17.4
	k8s.io/apiextensions-apiserver v0.17.4
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v8.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.5.1
)
