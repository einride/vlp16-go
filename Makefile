.PHONY: all
all: go-test test

include build/rules.mk
build/rules.mk:
	git submodule update --init

.PHONY: dep-ensure
dep-ensure: $(DEP)
	$(DEP) ensure -v

.PHONY: dep-check
dep-check: $(DEP)
	$(DEP) check

.PHONY: go-test
go-test: $(GOLANGCI_LINT) dep-ensure
	$(GOLANGCI_LINT) run --enable-all --disable=dupl

.PHONY: test
test:
	go test -race ./...

.PHONY: doc
doc:
	godoc -http=:6060 &
ifeq ($(Uname),Linux)
	xdg-open http://localhost:6060/pkg/github.com/einride/vlp-16-go/pkg/vlp16
else ifeq ($(Uname),Darwin)
	open http://localhost:6060/pkg/github.com/einride/vlp-16-go/pkg/vlp16
endif
