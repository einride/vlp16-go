.PHONY: all
all: \
	circleci-config-validate \
	mod-tidy \
	go-lint \
	go-review \
	go-test \
	git-verify-submodules \
	git-verify-nodiff


.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build:
	@git submodule update --init --recursive $@

include build/rules.mk
build/rules.mk: build
	@

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

# Dupl is disabled because of the testdata in spherical_point_cloud_test
.PHONY: go-lint
go-lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --enable-all --disable=dupl

.PHONY: go-test
go-test:
	go test -race -cover ./...

.PHONY: go-review
go-review: $(GOREVIEW)
	$(GOREVIEW) -c 1 ./...

.PHONY: circleci-config-validate
circleci-config-validate: $(CIRCLECI)
	$(CIRCLECI) config validate

.PHONY: doc
doc:
	godoc -http=:6060 &
ifeq ($(Uname),Linux)
	xdg-open http://localhost:6060/pkg/github.com/einride/vlp-16-go/pkg/vlp16
else ifeq ($(Uname),Darwin)
	open http://localhost:6060/pkg/github.com/einride/vlp-16-go/pkg/vlp16
endif
