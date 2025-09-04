resource "local_file" "calico_values" {
  content = templatefile("${path.module}/calico-values.tpl",
    {
      cni_iface                        = var.cni_iface
      subnet_pods                      = var.subnet_pods
      subnet_services                  = var.subnet_services
      windows_dataplane                = var.windows_dataplane
      calico_nat_outgoing              = var.calico_nat_outgoing == true ? "Enabled" : "Disabled"
      calico_encapsulation_type        = var.calico_encapsulation_type
      calico_interface_autodetect      = var.calico_interface_autodetect
      calico_interface_autodetect_cidr = var.calico_interface_autodetect_cidr
      calico_version                   = var.calico_version
      k8s_internal_vip                 = var.k8s_internal_vip
      k8s_api_port                     = var.k8s_api_port
  })

  filename = "${path.cwd}/cni-values.yml"

}


# module "calico" {
#   source = "./lib/calico"
#   count  = var.network_plugin == "calico" ? 1 : 0
#   # depends_on = [null_resource.copy_and_update_kubeconfig]
#   calico_interface_autodetect      = var.calico_interface_autodetect
#   calico_encapsulation_type        = var.calico_encapsulation_type
#   calico_nat_outgoing              = var.calico_nat_outgoing
#   calico_interface_autodetect_cidr = var.calico_interface_autodetect_cidr == "" ? var.subnet_nodes : var.calico_interface_autodetect_cidr
#   cni_iface                        = var.cni_iface
#   deploy_cluster                   = var.deploy_cluster
#   k8s_internal_vip                 = local.k8s_internal_vip
#   k8s_api_port                     = var.k8s_api_port
#   subnet_nodes                     = var.subnet_nodes
#   subnet_pods                      = var.subnet_pods
#   subnet_services                  = var.subnet_services
#   windows_dataplane                = length(var.windows_nodes) > 0 ? "HSN" : ""
# }