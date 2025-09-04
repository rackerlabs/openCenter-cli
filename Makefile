# ====================================================================================
# VARIABLES
# ====================================================================================

# Go variables
GO_PACKAGES := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name "*.go" -print)
BINARY_NAME := openCenter
BINARY_PATH := ./bin/$(BINARY_NAME)

# Tool versions
GOLANGCI_LINT_VERSION := v1.62.0
GOIMPORTS_VERSION := latest
MOCKERY_VERSION := v2.43.2

# Tools
TOOLS_DIR := bin
GOLANGCI_LINT := $(TOOLS_DIR)/golangci-lint
GOIMPORTS := $(TOOLS_DIR)/goimports
MOCKERY := $(TOOLS_DIR)/mockery

# Kubernetes (kind)
K8S_VERSION ?= 1.33.2
CNI ?= cilium
CNI_VERSION ?= 1.18.1

# ====================================================================================
# HELP
# ====================================================================================

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ====================================================================================
# DEVELOPMENT
# ====================================================================================

.PHONY: build
build: ## Build the application binary
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_PATH) main.go

.PHONY: tidy
tidy: ## Install and tidy dependencies
	@echo "Tidying go modules..."
	@go mod tidy

.PHONY: fmt
fmt: tools ## Format all go files using goimports
	@echo "Formatting go files..."
	@$(GOIMPORTS) -w -local $(shell go list -m) $(GO_FILES)

.PHONY: lint
lint: tools ## Lint the codebase with golangci-lint
	@echo "Linting..."
	@$(GOLANGCI_LINT) run ./...

# ====================================================================================
# TOOLS utilize for development
# ====================================================================================

.PHONY: tools
tools: $(GOLANGCI_LINT) $(GOIMPORTS) $(MOCKERY) ## Install development tools

$(GOLANGCI_LINT):
	@echo "Installing golangci-lint..."
	@GOBIN=$(shell pwd)/$(TOOLS_DIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

$(GOIMPORTS):
	@echo "Installing goimports..."
	@GOBIN=$(shell pwd)/$(TOOLS_DIR) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

$(MOCKERY):
	@echo "Installing mockery..."
	@GOBIN=$(shell pwd)/$(TOOLS_DIR) go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)


# ====================================================================================
# GITEA local development actions
# ====================================================================================
#
.PHONY: gitea-setup
gitea-setup: ## Create a local gitea setup for testing that listens on https port 3001
	@echo "Starting gitea..."
	@./hack/gitea-local/setup-gitea.sh

.PHONY: gitea-configure
gitea-configure: ## Generate tokens and save them for admin and newuser addiionally create a repo for testing
	@echo "Starting gitea..."
	@./hack/gitea-local/configure-gitea-user-tokens.sh

.PHONY: gitea-cleanup
gitea-cleanup: ## Remove gitea and all data that was used for testing
	@echo "Cleaning up local git..."
	@./hack/gitea-local/setup-gitea.sh destroy -y
	@rm -f .gitea_admin_token .gitea_newuser_token

# ====================================================================================
# PROJECT SPECIFIC
# ====================================================================================
.PHONY: clean
clean: ## Clean up build artifacts and tools
	@echo "Cleaning up..."
	@rm -f $(BINARY_PATH)
	@rm -rf $(TOOLS_DIR)

.DEFAULT_GOAL := help
