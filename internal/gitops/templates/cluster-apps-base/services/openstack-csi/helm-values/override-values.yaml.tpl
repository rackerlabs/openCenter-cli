---
secret:
    enabled:  {{ enabled }}
    hostMount: {{ HostName}}
    create: {{ create}}
    filename: {{ filename}}
    name: {{ clusterName}}
    data:
        cloud.conf: {{ cloud-conf }}
