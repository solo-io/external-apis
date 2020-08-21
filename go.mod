module github.com/solo-io/external-apis

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
)

require (
	github.com/coreos/etcd v3.3.15+incompatible // indirect
	github.com/golang/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	github.com/golang/mock v1.4.4
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.1.1
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/solo-io/skv2 v0.10.0
	gonum.org/v1/netlib v0.0.0-20191031114514-eccb95939662 // indirect
	google.golang.org/genproto v0.0.0-20200117163144-32f20d992d24 // indirect
	istio.io/client-go v0.0.0-20200610222318-1cfead1f1938
	k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/structured-merge-diff v1.0.1-0.20191108220359-b1b620dd3f06 // indirect
	sigs.k8s.io/testing_frameworks v0.1.2 // indirect
)
