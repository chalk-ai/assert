.PHONY: all test cloc

SHELL := /bin/bash
TEMP_FILE := $(shell mktemp)

help:  ## Show this help.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

test:  ## Run tests.
	@which gotestsum > /dev/null || go get gotest.tools/gotestsum@latest
	@CGO_ENABLED=0 gotestsum --hide-summary=skipped --format-hide-empty-pkg -- -short -vet=all -shuffle=on ./...

build.test:  ## Build everything, including tests
	@CGO_ENABLED=0 go test -run=XXX_SHOULD_NEVER_MATCH_XXX ./... | grep -v "no test files" | grep -v "no tests to run" | grep -v "FAIL"

cloc:
	cloc . --vcs=git --exclude-lang JSON,SVG,.pyi --not-match-f generated.go

.DEFAULT:
	@echo Unknown command. Available commands below
	@echo
	@make help
