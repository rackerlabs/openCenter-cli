#BUG wont install into correct namespace https://github.com/hashicorp/terraform-provider-helm/pull/1583
locals {
  ip_octets   = split(".", var.subnet_pods)
  pod_gateway = join(".", [element(local.ip_octets, 0), element(local.ip_octets, 1), element(local.ip_octets, 2)])
}

resource "helm_release" "kube_ovn" {
  count            = var.deploy_cluster ? 1 : 0
  name             = var.kube_ovn_chart_name
  namespace        = var.kube_ovn_chart_name
  create_namespace = true

  repository = var.kube_ovn_chart_repo
  chart      = var.kube_ovn_chart_name
  version    = var.kube_ovn_version # adjust as needed

  values = [
    yamlencode({
      ipv4 = {
        POD_CIDR    = var.subnet_pods          # ovn-default CIDR
        POD_GATEWAY = "${local.pod_gateway}.1" # ovn-default CIDR
        SVC_CIDR    = var.subnet_services      # Service CIDR
        JOIN_CIDR   = var.subnet_join          # join
      }
      networking = {
        IFACE = var.cni_iface
      }
      MASTER_NODES = join(",", [
        for k, v in var.master_nodes : v.access_ip_v4
      ])
      replicaCount = length(var.master_nodes)
      ovn-central = {
        requests = {
          cpu : "300m"
          memory : "200Mi"
        }
        limits = {
          cpu : "2"
          memory : "4Gi"
        }
      }
      # You can add more inline values here if needed
    })
  ]

  timeout = 120
  atomic  = true
}