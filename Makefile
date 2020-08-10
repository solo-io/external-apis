
#----------------------------------------------------------------------------------
# Generated Code and Docs
#----------------------------------------------------------------------------------

DEPSGOBIN=$(shell pwd)/_output/.bin
PATH=$(DEPSGOBIN):$$PATH

.PHONY: mod-download
mod-download:
	go mod download

# Dependencies for code generation
.PHONY: codegen-deps
codegen-deps: mod-download
	mkdir -p $(DEPSGOBIN)
	GOBIN=$(DEPSGOBIN) go install github.com/golang/mock/mockgen
	GOBIN=$(DEPSGOBIN) go install golang.org/x/tools/cmd/goimports

.PHONY: generated-code
generated-code:
	go run generate.go
	go generate -v ./...
	goimports -w .
	go mod tidy

