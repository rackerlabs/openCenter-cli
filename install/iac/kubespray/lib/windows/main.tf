resource "helm_release" "calico" {
  count            = var.deploy_cluster ? 1 : 0
  name             = var.chart_name
  namespace        = var.chart_namespace
  create_namespace = true

  repository = var.chart_repo
  chart      = var.chart_name
  version    = var.calico_version # adjust as needed

  values = [
    yamlencode({
      installation = {
        enabled : true
        kubernetesProvider : ""
        calicoNetwork = {
          windowsDataplane : var.windows_dataplane
          ipPools = [
            {
              cidr          = var.subnet_pods
              encapsulation = "VXLAN"
              natOutgoing   = "Enabled"
            }
          ]
        }
        serviceCIDRs = [var.subnet_services]
      }
    })
  ]

  timeout = 120
  atomic  = true
}