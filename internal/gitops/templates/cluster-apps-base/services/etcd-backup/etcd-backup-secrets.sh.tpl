{/*
This file was generated from overlay template comparison
Environment-specific values are templated with Go template syntax
Original source: dev environment overlay
*/}
kubectl --namespace etcd-backup \
create secret generic etcd-backup-secrets \
--type Opaque \
--from-literal=ACCESS_KEY="31d33dd98caa4d1d9a1d31406da95065" \
--from-literal=SECRET_KEY="5adc3d4ee8da4c67b9398ff07157a8fa" \
--from-literal=S3_HOST="https://swift.api.dfw3.rackspacecloud.com" \
--from-literal=S3_REGION="DFW3" \
--from-literal=ETCDCTL_API="3" \
--from-literal=ETCDCTL_ENDPOINTS="https://127.0.0.1:2379" \
--from-literal=ETCDCTL_CACERT="/etc/kubernetes/ssl/etcd/ca.crt" \
--from-literal=ETCDCTL_CERT="/etc/kubernetes/ssl/etcd/server.crt" \
--from-literal=ETCDCTL_KEY="/etc/kubernetes/ssl/etcd/server.key"
