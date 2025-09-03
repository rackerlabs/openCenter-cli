
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
