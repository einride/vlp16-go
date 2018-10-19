# -------------
# Makefile
# -------------
#
# This Makefile is designed to run reproducible builds locally (on macOS and
# Linux) and on a CI server. Windows builds are not supported.
.PHONY: all
all: lint-go test
Uname := $(shell uname -s)
uname := $(shell uname -s | tr '[:upper:]' '[:lower:]')

ifeq ($(Uname),Linux)
else ifeq ($(Uname),Darwin)
else
$(error This Makefile only supports Linux and OSX build agents)
endif

ifneq ($(shell uname -m),x86_64)
$(error This Makefile only supports x86_64 build agents)
endif

# -------------------------------------
# Download vendor dependencies with dep
# -------------------------------------

dep_version := 0.5.0
dep_dir := build/dep/$(dep_version)
dep := $(dep_dir)/dep

$(dep):
	mkdir -p $(dep_dir)
	curl -s -L https://github.com/golang/dep/releases/download/v$(dep_version)/dep-$(uname)-amd64 -o $(dep_dir)/dep
	chmod +x $(dep_dir)/dep

vendor: $(dep) Gopkg.toml Gopkg.lock
	$(dep) ensure

# -------------------------------
# Lint Go code with GolangCI-Lint
# -------------------------------
golangci_lint_version := 1.10.2
golangci_lint_dir := build/golangci-lint/$(golangci_lint_version)
golangci_lint := $(golangci_lint_dir)/golangci-lint

$(golangci_lint):
	mkdir -p $(golangci_lint_dir)
	curl -s -L https://github.com/golangci/golangci-lint/releases/download/v$(golangci_lint_version)/golangci-lint-$(golangci_lint_version)-$(uname)-amd64.tar.gz -o $(golangci_lint_dir)/archive.tar.gz
	tar xzf $(golangci_lint_dir)/archive.tar.gz -C $(golangci_lint_dir) --strip 1

.PHONY: default
	echo "Enter run-vlp16"

# Check linting
# Maligned disabled in golangci-lint since structs are so large anyway.
# Might want to look at optimising structs in the future.
.PHONY: lint-go
lint-go: vendor $(golangci_lint)
	$(golangci_lint) run ./... --enable-all --disable=gas

# -------------------------------------
# Test
# -------------------------------------

.PHONY: test
test: lint-go
	go test -race ./...
# -------------------------------------
# Build
# -------------------------------------

.PHONY: build-vlp16
build-vlp16:
	go build ./cmd/vlp16

# -------------------------------------
# Run
# -------------------------------------

.PHONY: run-vlp16
run-vlp16: build-vlp16



