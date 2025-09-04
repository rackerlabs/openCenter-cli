# ====================================================================================
# VARIABLES
# ====================================================================================

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

.DEFAULT_GOAL := help
