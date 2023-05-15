module github.com/solo-io/external-apis

go 1.16

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
)

require (
	github.com/Masterminds/sprig/v3 v3.2.0 // indirect
	github.com/golang/mock v1.6.0
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.1.1
	github.com/solo-io/skv2 v0.30.0
	k8s.io/api v0.26.4
	k8s.io/apiextensions-apiserver v0.26.4
	k8s.io/apimachinery v0.26.4
	k8s.io/client-go v0.26.4
	sigs.k8s.io/controller-runtime v0.14.6
)
