# all: run a complete build
.PHONY: all
all: \
	circleci-config-validate \
	go-generate \
	go-mocks \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy \
	git-verify-submodules \
	git-verify-nodiff

export GO111MODULE := on

# clean: remove generated build files
.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build:
	@git submodule update --init --recursive $@

include build/rules.mk
build/rules.mk: build
	@# included in submodule: build

# go-mod-tidy: update go modules
.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

# go-lint: lint Go files
.PHONY: go-lint
go-lint: $(GOLANGCI_LINT)
	# dupl: disabled due to the testdata in spherical_point_cloud_test
	$(GOLANGCI_LINT) run --enable-all --disable=dupl

# go-test: run Go test suite
.PHONY: go-test
go-test:
	go test -race -cover ./...

# go-review: run goreview linter
.PHONY: go-review
go-review: $(GOREVIEW)
	$(GOREVIEW) -c 1 ./...

# circleci-config-validate: validate CircleCI config
.PHONY: circleci-config-validate
circleci-config-validate: $(CIRCLECI)
	$(CIRCLECI) config validate

# go-generate: generate Go code
.PHONY: go-generate
go-generate: returnmode_string.go productid_string.go

returnmode_string.go: returnmode.go $(GOBIN)
	$(GOBIN) -m -run golang.org/x/tools/cmd/stringer \
		-type ReturnMode -trimprefix ReturnMode -output $@ $<

productid_string.go: productid.go $(GOBIN)
	$(GOBIN) -m -run golang.org/x/tools/cmd/stringer \
		-type ProductID -trimprefix ProductID -output $@ $<

# go-mocks: generate Go mocks
.PHONY: go-mocks
go-mocks: test/mocks/vlp16/mocks.go

test/mocks/vlp16/mocks.go: client.go $(GOBIN)
	$(GOBIN) -m -run github.com/golang/mock/mockgen -destination $@ -package mockvlp16 \
		github.com/einride/vlp-16-go UDPConn
