# Code generated by Mage-tools. DO NOT EDIT.

.DEFAULT_GOAL := all

include .mage/tools.mk

.PHONY: all
all: $(mage)
	@$(mage) all

.PHONY: convco-check
convco-check: $(mage)
ifndef rev
	$(error missing argument rev="...")
endif
	@$(mage) convcoCheck $(rev)

.PHONY: format-markdown
format-markdown: $(mage)
	@$(mage) formatMarkdown

.PHONY: git-verify-no-diff
git-verify-no-diff: $(mage)
	@$(mage) gitVerifyNoDiff

.PHONY: go-mod-tidy
go-mod-tidy: $(mage)
	@$(mage) goModTidy

.PHONY: go-test
go-test: $(mage)
	@$(mage) goTest

.PHONY: golangci-lint
golangci-lint: $(mage)
	@$(mage) golangciLint
