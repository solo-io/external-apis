
#----------------------------------------------------------------------------------
# Generated Code and Docs
#----------------------------------------------------------------------------------

.PHONY: mod-download
mod-download:
	go mod download

# Dependencies for code generation
.PHONY: codegen-deps
codegen-deps: mod-download
	go get -v github.com/golang/mock/mockgen@v1.4.3
	go get -v golang.org/x/tools/cmd/goimports@v0.0.0-20200423205358-59e73619c742

.PHONY: generated-code
generated-code:
	go run generate.go
	go generate -v ./...
	goimports -w .

