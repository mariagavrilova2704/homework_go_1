SHELL := /bin/bash

GO ?= go
GO_PACKAGES := ./...
UNIT_PACKAGES := ./internal/...
INTEGRATION_PACKAGES := ./test/integration/...
APP_PATH := ./cmd/demo
APP_BIN := bin/main
PACKAGE_FILE := bin/main-linux-amd64.tar.gz
ARGS ?=
LOCAL_BIN := $(CURDIR)/bin
GOLANGCI_LINT := $(LOCAL_BIN)/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.64.8
GOLANGCI_LINT_VERSION_NO_V := $(patsubst v%,%,$(GOLANGCI_LINT_VERSION))
GOLANGCI_LINT_DOWNLOAD_URL ?=
COVERAGE_FILE ?= coverage.out
COVERAGE_THRESHOLD ?= 90.0
GO_TEST_FLAGS ?= -race -covermode=atomic -coverprofile=$(COVERAGE_FILE)

.PHONY: \
	help \
	run \
	build \
	package \
	deps-check \
	mod-check \
	test \
	test-unit \
	test-integration \
	test-race \
	coverage \
	coverage-check \
	fmt \
	fmt-check \
	vet \
	tools \
	lint \
	lint-style \
	lint-static \
	lint-security \
	check \
	ci \
	clean

help:
	@echo "Available commands:"
	@echo "  make run              - run demo application with default input"
	@echo '  make run ARGS="Алексей 17 250 4 8 16 true false"'
	@echo "  make build            - build demo binary"
	@echo "  make package          - build tar.gz artifact"
	@echo "  make deps-check       - download and verify Go modules"
	@echo "  make mod-check        - check that go.mod/go.sum are tidy"
	@echo "  make test             - run unit and integration tests"
	@echo "  make test-unit        - run unit tests only"
	@echo "  make test-integration - build binary and verify CLI output"
	@echo "  make test-race        - run unit tests with race detector"
	@echo "  make coverage         - run unit tests and print coverage"
	@echo "  make coverage-check   - fail if unit coverage is below threshold"
	@echo "  make fmt              - format Go files"
	@echo "  make fmt-check        - check formatting without changing files"
	@echo "  make vet              - run go vet"
	@echo "  make lint-style       - run formatting and style linters"
	@echo "  make lint-static      - run static analysis linters"
	@echo "  make lint-security    - run security linters"
	@echo "  make lint             - run full golangci-lint config"
	@echo "  make check            - run fast local checks"
	@echo "  make ci               - run the same checks as CI"
	@echo "  make clean            - remove generated artifacts"

run:
	$(GO) run $(APP_PATH) $(ARGS)

build:
	@mkdir -p bin
	$(GO) build -o "$(APP_BIN)" $(APP_PATH)

package: build
	tar -czf "$(PACKAGE_FILE)" -C bin main

deps-check:
	$(GO) mod download
	$(GO) mod verify

mod-check:
	$(GO) mod tidy
	@if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then \
		git diff --exit-code -- go.mod go.sum; \
	else \
		echo "Skipping git diff because this directory is not a git repository"; \
	fi

test: test-unit test-integration

test-unit:
	$(GO) test $(UNIT_PACKAGES)

test-integration:
	$(GO) test $(INTEGRATION_PACKAGES)

test-race:
	$(GO) test -race $(UNIT_PACKAGES)

coverage:
	$(GO) test $(GO_TEST_FLAGS) $(UNIT_PACKAGES)
	$(GO) tool cover -func="$(COVERAGE_FILE)"

coverage-check: coverage
	@total="$$(go tool cover -func="$(COVERAGE_FILE)" | awk '/^total:/ { gsub(/%/, "", $$3); print $$3 }')"; \
	awk -v total="$$total" -v threshold="$(COVERAGE_THRESHOLD)" 'BEGIN { \
		if (total + 0 < threshold + 0) { \
			printf "Coverage %.1f%% is below required %.1f%%\n", total, threshold; \
			exit 1; \
		} \
		printf "Coverage %.1f%% is OK. Required %.1f%%\n", total, threshold; \
	}'

fmt:
	find . -type f -name '*.go' -not -path './.git/*' -not -path './bin/*' -print0 \
		| xargs -0 gofmt -w

fmt-check:
	@files="$$(find . -type f -name '*.go' -not -path './.git/*' -not -path './bin/*' -print0 \
		| xargs -0 gofmt -l)"; \
	if [[ -n "$$files" ]]; then \
		echo "Go files are not formatted. Run: make fmt"; \
		echo "$$files"; \
		exit 1; \
	fi

vet:
	$(GO) vet $(GO_PACKAGES)

tools:
	@mkdir -p "$(LOCAL_BIN)"
	@if [[ ! -x "$(GOLANGCI_LINT)" ]]; then \
		tmp_dir="$$(mktemp -d)"; \
		trap 'rm -rf "$$tmp_dir"' EXIT; \
		os="$$(uname -s | tr '[:upper:]' '[:lower:]')"; \
		arch="$$(uname -m)"; \
		case "$$os" in \
			darwin|linux) ;; \
			*) echo "Unsupported OS for golangci-lint archive: $$os"; exit 1 ;; \
		esac; \
		case "$$arch" in \
			x86_64|amd64) arch="amd64" ;; \
			aarch64|arm64) arch="arm64" ;; \
			*) echo "Unsupported architecture for golangci-lint archive: $$arch"; exit 1 ;; \
		esac; \
		archive="golangci-lint-$(GOLANGCI_LINT_VERSION_NO_V)-$$os-$$arch.tar.gz"; \
		url="$(GOLANGCI_LINT_DOWNLOAD_URL)"; \
		if [[ -z "$$url" ]]; then \
			url="https://github.com/golangci/golangci-lint/releases/download/$(GOLANGCI_LINT_VERSION)/$$archive"; \
		fi; \
		echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION) from $$url"; \
		curl -fsSL "$$url" -o "$$tmp_dir/$$archive"; \
		tar -xzf "$$tmp_dir/$$archive" -C "$$tmp_dir"; \
		binary="$$(find "$$tmp_dir" -type f -name golangci-lint | head -n 1)"; \
		if [[ -z "$$binary" ]]; then \
			echo "golangci-lint binary was not found inside $$archive"; \
			exit 1; \
		fi; \
		cp "$$binary" "$(GOLANGCI_LINT)"; \
		chmod +x "$(GOLANGCI_LINT)"; \
	fi

lint: tools
	"$(GOLANGCI_LINT)" run $(GO_PACKAGES)

lint-style: tools
	"$(GOLANGCI_LINT)" run \
		--disable-all \
		--enable=gofmt \
		--enable=gofumpt \
		--enable=lll \
		--enable=funlen \
		--enable=whitespace \
		$(GO_PACKAGES)

lint-static: tools
	"$(GOLANGCI_LINT)" run \
		--disable-all \
		--enable=asciicheck \
		--enable=bidichk \
		--enable=bodyclose \
		--enable=cyclop \
		--enable=errcheck \
		--enable=errorlint \
		--enable=exhaustive \
		--enable=gocognit \
		--enable=goconst \
		--enable=gocritic \
		--enable=gosimple \
		--enable=govet \
		--enable=ineffassign \
		--enable=misspell \
		--enable=nakedret \
		--enable=nestif \
		--enable=nilerr \
		--enable=noctx \
		--enable=prealloc \
		--enable=revive \
		--enable=staticcheck \
		--enable=unconvert \
		--enable=unused \
		$(GO_PACKAGES)

lint-security: tools
	"$(GOLANGCI_LINT)" run \
		--disable-all \
		--enable=gosec \
		$(GO_PACKAGES)

check: fmt-check vet lint-style lint-static lint-security test-unit test-integration

ci: deps-check mod-check fmt-check vet lint test-unit test-integration test-race coverage-check build

clean:
	rm -rf bin "$(COVERAGE_FILE)"