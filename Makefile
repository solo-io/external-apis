
#----------------------------------------------------------------------------------
# Generated Code and Docs
#----------------------------------------------------------------------------------

DEPSGOBIN=$(shell pwd)/.bin
export PATH:=$(DEPSGOBIN):$(PATH)
export GOBIN:=$(DEPSGOBIN)

.PHONY: mod-download
mod-download:
	go mod download

# Dependencies for code generation
.PHONY: codegen-deps
codegen-deps: mod-download
	mkdir -p $(DEPSGOBIN)
	go install github.com/golang/mock/mockgen
	go install golang.org/x/tools/cmd/goimports

.PHONY: generated-code
generated-code:
	go run generate.go
	go generate -v ./...
	goimports -w .
	go mod tidy 

