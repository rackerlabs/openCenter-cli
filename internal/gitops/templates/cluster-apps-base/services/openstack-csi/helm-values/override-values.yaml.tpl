secret:
  enabled: true
  hostMount: true
  create: true
  filename: cloud.conf
  name: cinder-csi-cloud-config
  data:
    cloud.conf: |-
      [Global]
      auth-url = {{ .OpenCenter.Infrastructure.Cloud.OpenStack.AuthURL }}
      application-credential-id = {{ .OpenCenter.Infrastructure.Cloud.OpenStack.ApplicationCredentialID }}
      application-credential-secret = {{ .OpenCenter.Infrastructure.Cloud.OpenStack.ApplicationCredentialSecret }}
      domain-name = {{ .OpenCenter.Infrastructure.Cloud.OpenStack.Domain }}
      region = {{ .OpenCenter.Infrastructure.Cloud.OpenStack.Region }}
      tenant-name = {{ .OpenCenter.Infrastructure.Cloud.OpenStack.TenantName }}
      tls-insecure = {{ .OpenCenter.Infrastructure.Cloud.OpenStack.Insecure | default false }}