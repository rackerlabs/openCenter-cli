apiVersion: v1
data:
    admin-password: {{ .Secrets.Grafana.AdminPassword | quote }}
    admin-user: admin
kind: Secret
metadata:
    name: grafana-admin-password
    namespace: observability

