apiVersion: v1
kind: Secret
metadata:
    name: cloud-credentials
    namespace: velero
type: Opaque
stringData:
    OS_AUTH_URL: {{ .OpenCenter.Infrastructure.Cloud.OpenStack.AuthURL }}
    OS_PROJECT_NAME: {{ .OpenCenter.Infrastructure.Cloud.OpenStack.TenantName }}
    OS_APPLICATION_CREDENTIAL_ID: {{ .OpenCenter.Infrastructure.Cloud.OpenStack.ApplicationCredentialID }}
    OS_APPLICATION_CREDENTIAL_SECRET: {{ .OpenCenter.Infrastructure.Cloud.OpenStack.ApplicationCredentialSecret }}
    OS_REGION_NAME: {{ .OpenCenter.Infrastructure.Cloud.OpenStack.Region }}
    OS_DOMAIN_NAME: {{ .OpenCenter.Infrastructure.Cloud.OpenStack.Domain }}
    OS_SWIFT_TEMP_URL_KEY: ""
    OS_SWIFT_TEMP_URL_DIGEST: sha256