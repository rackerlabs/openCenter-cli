---
config:
  oidc:
    enabled: {{ .OpenCenter.Cluster.Kubernetes.OIDC.Enabled }}
    externalSecret:
      enabled: false
    secret:
      create: true
    clientID: {{ .OpenCenter.Cluster.Kubernetes.OIDC.KubeOIDCClientID | default "kubernetes" }}
    clientSecret: ""
    issuerURL: {{ .OpenCenter.Cluster.Kubernetes.OIDC.KubeOIDCURL | default "https://auth.prosys.dev.dfw3.k8s.opencenter.cloud/realms/opencenter" }}
    scopes: "openid profile email groups"
    callbackURL: "https://headlamp.{{ .OpenCenter.Cluster.ClusterName }}.k8s.opencenter.cloud/oidc-callback"
