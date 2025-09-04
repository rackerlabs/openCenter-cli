# openCenter


For instructions on how to do local development refer to the [local development setup](docs/local-development.md).


```
kubernetes-platform/
├── README.md
├── applications
│   ├── base
│   │   ├── genestack-sources
│   │   │   ├── genestack.yaml
│   │   │   ├── gitrepository-aggregator.yaml
│   │   │   ├── kustomization.yaml
│   │   │   └── openstack-helm.yaml
│   │   ├── managed-services
│   │   │   ├── cert-manager
│   │   │   │   └── placeholder.txt
│   │   │   ├── gateway-api
│   │   │   │   └── placeholder.txt
│   │   │   ├── ingress-nginx
│   │   │   │   └── placeholder.txt
│   │   │   ├── keycloak
│   │   │   │   └── placeholder.txt
│   │   │   ├── sealed-secrets
│   │   │   │   └── placeholder.txt
│   │   │   └── sources
│   │   │       ├── bitnami.yaml
│   │   │       ├── envoyproxy.yaml
│   │   │       ├── ingress-nginx.yaml
│   │   │       ├── jetstack.yaml
│   │   │       ├── kustomization.yaml
│   │   │       └── sealed-secrets.yaml
│   ├── overlays
│   │   ├── delta
│   │   │   ├── flux-system
│   │   │   │   ├── gotk-components.yaml
│   │   │   │   ├── gotk-sync.yaml
│   │   │   │   └── kustomization.yaml
│   │   │   ├── genestack
│   │   │   │   └── fluxcd
│   │   │   ├── kustomization.yaml
│   │   │   └── managed-services
│   │   │       ├── cert-manager
│   │   │       ├── fluxcd
│   │   │       ├── gateway
│   │   │       ├── gateway-api
│   │   │       ├── ingress-nginx
│   │   │       ├── keycloak
│   │   │       └── sealed-secrets
│   │   ├── dev
│   │   │   └── placeholder.txt
│   │   └── production
│   │       └── placeholder.txt
│   └── policies
│       ├── network-policies
│       │   └── placeholder.txt
│       ├── pod-security-policies
│       │   └── placeholder.txt
│       └── rbac
│           └── placeholder.txt
```




### Prerequisites

- [podman](https://podman.io/get-started)/[orbstack](https://orbstack.dev/) installed and running
- [Mise](https://mise.jdx.dev/) for tool version management

### Initial Setup

#### 1. Install Mise

```bash
# macOS
brew install mise

# Linux
curl https://mise.run | sh

# Add to your shell profile
# or for fish
echo '' >> ~/.zshrc
# or for bash
echo 'eval "$(mise activate bash)"' >> ~/.bashrc
# or for zsh
echo 'eval "$(mise activate zsh)"' >> ~/.zshrc
```
A quick-start guide for local development with Go CLI tools, FluxCD, and Kubernetes using Kind and Mise.

