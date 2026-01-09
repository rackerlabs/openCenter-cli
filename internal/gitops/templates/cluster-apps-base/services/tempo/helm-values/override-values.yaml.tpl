global:
    storageClass: {{ .OpenCenter.Services.CSI.StorageClass }}
storage:
    trace:
        backend: s3
        s3:
            bucket: {{ .OpenCenter.Cluster.ClusterName }}-tempo
            endpoint: swift.api.{{ .OpenCenter.Meta.Region }}.rackspacecloud.com
            access_key: {{ .GetTempoS3AccessKey }}
            secret_key: {{ .GetTempoS3SecretKey }}
            region: {{ (index .OpenCenter.Services "tempo").Region }}
            insecure: false
multitenancyEnabled: true
ingester:
    persistence:
        size: 50Gi
