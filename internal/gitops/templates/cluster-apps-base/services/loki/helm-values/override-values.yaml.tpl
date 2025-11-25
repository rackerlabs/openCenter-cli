{{- $loki := index .OpenCenter.Services "loki" }}
loki:
  storage:
    bucketNames:
    chunks: {{ $loki.LokiBucketName }}
    ruler: {{ $loki.LokiBucketName }}
    admin: {{ $loki.LokiBucketName }}
    type: swift
    swift:
    auth_url: {{ $loki.SwiftAuthURL }}
    auth_version: 3
    internal: false
    application_credential_id: {{ $loki.SwiftUsername | quote }}
    application_credential_secret: {{ .Secrets.Loki.SwiftPassword | quote }}
    user_domain_name: {{ $loki.SwiftDomainName }}
    project_name: {{ $loki.SwiftProjectName }}
    project_domain_name: {{ $loki.SwiftDomainName }}
    region_name: {{ $loki.SwiftRegion }}
    container_name: {{ $loki.LokiBucketName }}
    max_retries: 5
    connect_timeout: 10s
    request_timeout: 30s
  # Local pathing used by the charted components
  storage_config:
    tsdb_shipper:
    active_index_directory: /var/loki/index
    cache_location: /var/loki/index-cache
  # Scraping (Prometheus)
  serviceMonitor:
    enabled: true
write:
  replicas: 3
  resources:
    requests:
    cpu: 100m
    memory: 500Mi
    limits:
    cpu: "1"
    memory: 1Gi
  persistence:
    enabled: true
    size: {{ $loki.LokiVolumeSize }}Gi
    storageClass: {{ $loki.LokiStorageClass }}
  podAntiAffinityPreset: soft
read:
  replicas: 3
  resources:
    requests:
    cpu: 100m
    memory: 500Mi
    limits:
    cpu: "1"
    memory: 1Gi
  persistence:
    enabled: true
    size: {{ $loki.LokiVolumeSize }}Gi
    storageClass: {{ $loki.LokiStorageClass }}
  podAntiAffinityPreset: soft
backend:
  replicas: 3
  resources:
    requests:
    cpu: 100m
    memory: 400Mi
    limits:
    cpu: "1"
    memory: 1Gi
  persistence:
    enabled: true
    size: {{ $loki.LokiVolumeSize }}Gi
    storageClass: {{ $loki.LokiStorageClass }}
  podAntiAffinityPreset: soft
gateway:
  replicas: 2
  resources:
    requests:
    cpu: 100m
    memory: 256Mi
    limits:
    cpu: 500m
    memory: 512Mi
  ingress:
    enabled: false
chunksCache:
  enabled: true
  memcached:
    replicaCount: 3
    resources:
    requests:
      cpu: 100m
      memory: 512Mi
    limits:
      cpu: "1"
      memory: 1Gi
resultsCache:
  enabled: true
  memcached:
    replicaCount: 3
    resources:
    requests:
      cpu: 100m
      memory: 512Mi
    limits:
      cpu: "1"
      memory: 1Gi
