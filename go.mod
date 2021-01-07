module github.com/solo-io/external-apis

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
)

require (
	github.com/Azure/go-autorest/autorest v0.9.3 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.8.1 // indirect
	github.com/aws/aws-app-mesh-controller-for-k8s v1.1.1
	github.com/envoyproxy/protoc-gen-validate v0.4.0 // indirect
	github.com/golang/mock v1.4.4
	github.com/mattn/go-zglob v0.0.3 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.1.1
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/solo-io/skv2 v0.15.2
	github.com/spf13/afero v1.3.4 // indirect
	go.uber.org/multierr v1.4.0 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	istio.io/client-go v0.0.0-20200610222318-1cfead1f1938
	k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime v0.6.2
)
