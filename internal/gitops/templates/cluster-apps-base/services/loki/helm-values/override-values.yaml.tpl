loki:
  auth_enabled: true
  storage:
    bucketNames:
      chunks: {{ .OpenCenter.Services.loki.LokiBucketName | default (printf "%s-loki" .OpenCenter.Cluster.ClusterName) }}
      ruler: {{ .OpenCenter.Services.loki.LokiBucketName | default (printf "%s-loki" .OpenCenter.Cluster.ClusterName) }}
      admin: {{ .OpenCenter.Services.loki.LokiBucketName | default (printf "%s-loki" .OpenCenter.Cluster.ClusterName) }}
    type: swift
    swift:
      auth_url: {{ .OpenCenter.Services.loki.SwiftAuthURL }}
      auth_version: 3
      internal: false
      username: {{ .OpenCenter.Services.loki.SwiftUsername }}
      password: {{ .Secrets.Loki.SwiftPassword }}
      user_domain_name: {{ .OpenCenter.Services.loki.SwiftDomainName }}
      project_name: {{ .OpenCenter.Services.loki.SwiftProjectName }}
      project_domain_name: {{ .OpenCenter.Services.loki.SwiftDomainName }}
      region_name: {{ .OpenCenter.Services.loki.SwiftRegion }}
      container_name: {{ .OpenCenter.Services.loki.LokiBucketName | default (printf "%s-loki" .OpenCenter.Cluster.ClusterName) }}
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
    size: {{ .OpenCenter.Services.loki.LokiVolumeSize | default 20 }}Gi
    storageClass: {{ .OpenCenter.Services.loki.LokiStorageClass | default .OpenCenter.Storage.DefaultStorageClass | default "csi-cinder-sc-delete" }}
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
    size: {{ .OpenCenter.Services.loki.LokiVolumeSize | default 20 }}Gi
    storageClass: {{ .OpenCenter.Services.loki.LokiStorageClass | default .OpenCenter.Storage.DefaultStorageClass | default "csi-cinder-sc-delete" }}
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
    size: {{ .OpenCenter.Services.loki.LokiVolumeSize | default 20 }}Gi
    storageClass: {{ .OpenCenter.Services.loki.LokiStorageClass | default .OpenCenter.Storage.DefaultStorageClass | default "csi-cinder-sc-delete" }}
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
