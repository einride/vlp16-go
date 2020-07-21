# all: run a complete build
.PHONY: all
all: \
	go-generate \
	go-review \
	go-test \
	go-lint \
	go-mod-tidy \
	git-verify-nodiff

include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk
include tools/xtools/rules.mk

# go-mod-tidy: update go modules
.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

# go-test: run Go test suite
.PHONY: go-test
go-test:
	go test -race -cover ./...

# go-generate: generate Go code
.PHONY: go-generate
go-generate: returnmode_string.go productid_string.go

returnmode_string.go: returnmode.go $(stringer)
	$(stringer) -type ReturnMode -trimprefix ReturnMode -output $@ $<

productid_string.go: productid.go $(GOBIN)
	$(stringer) -type ProductID -trimprefix ProductID -output $@ $<
