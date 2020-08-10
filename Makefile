
#----------------------------------------------------------------------------------
# Generated Code and Docs
#----------------------------------------------------------------------------------

DEPSGOBIN=$(shell pwd)/_output/.bin

.PHONY: mod-download
mod-download:
	PATH=$(DEPSGOBIN):$$PATH go mod download

# Dependencies for code generation
.PHONY: codegen-deps
codegen-deps: mod-download
	mkdir -p $(DEPSGOBIN)
	GOBIN=$(DEPSGOBIN) go install github.com/golang/mock/mockgen
	GOBIN=$(DEPSGOBIN) go install golang.org/x/tools/cmd/goimports

.PHONY: generated-code
generated-code:
	PATH=$(DEPSGOBIN):$$PATH go run generate.go
	PATH=$(DEPSGOBIN):$$PATH go generate -v ./...
	PATH=$(DEPSGOBIN):$$PATH goimports -w .

