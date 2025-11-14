apiVersion: v1
data:
  admin-password: {{ .Secrets.Grafana.AdminPassword | b64enc }}
  admin-user: {{ "admin" | b64enc }}
kind: Secret
metadata:
  creationTimestamp: null
  name: grafana-admin-password
  namespace: observability
