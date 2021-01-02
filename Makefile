SHELL := /bin/bash

.PHONY: all
all: \
	commitlint \
	go-generate \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy \
	git-verify-nodiff

include tools/commitlint/rules.mk
include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk
include tools/semantic-release/rules.mk
include tools/stringer/rules.mk

.PHONY: clean
clean:
	$(info [$@] cleaning generated files...)
	@find -name '*_string.go' -exec rm {} \+
	@rm -rf build

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@mkdir -p build/coverage
	@go test -short -race -coverprofile=build/coverage/$@.txt -covermode=atomic ./...

.PHONY: go-generate
go-generate: \
	productid_string.go \
	returnmode_string.go

%_string.go: %.go $(stringer)
	$(info generating $@...)
	@go generate ./$<
