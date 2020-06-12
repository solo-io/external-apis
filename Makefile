
#----------------------------------------------------------------------------------
# Generated Code and Docs
#----------------------------------------------------------------------------------

.PHONY: mod-download
mod-download:
	go mod download

# Dependencies for code generation
.PHONY: codegen-deps
codegen-deps: mod-download
	GO111MODULE=off go get -v github.com/golang/mock/mockgen
	GO111MODULE=off go get -v golang.org/x/tools/cmd/goimports

.PHONY: generated-code
generated-code:
	go run generate.go
	go generate -v ./...
	goimports -w .

