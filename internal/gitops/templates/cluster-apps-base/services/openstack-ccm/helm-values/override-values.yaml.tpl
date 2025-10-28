---
cluster:
    name: {{ ClusterName}}
cloudConfig:
    global:
        auth-url: {{ auth-url}}
        application-credential-id: {{ application-credential-id }}
        application-credential-secret: {{ application-credential-secret }}
        domain-name: {{ domain-name }}
        region: {{ Region }}
        tenant-name:  {{ tenant-name }}
        tls-insecure: {{ tls-insecure | default ("false") }}
    loadBalancer:
        floating-network-id: {{ floating-network-id }}
        subnet-id:  {{ subnet-id }}
