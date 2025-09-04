variable "cni_iface" {
  type    = string
  default = ""
}

variable "deploy_cluster" {
  type    = bool
  default = false
}

variable "kubeconfig_path" {
  type    = string
  default = "./kubeconfig"
}

variable "master_nodes" {
  type = list(object({
    id           = string
    name         = string
    access_ip_v4 = string
  }))
}

variable "worker_nodes" {
  type = list(object({
    id           = string
    name         = string
    access_ip_v4 = string
  }))
}

variable "subnet_pods" {
  type    = string
  default = "10.236.0.0/14"
}

variable "subnet_services" {
  type    = string
  default = "10.233.0.0/18"
}

variable "subnet_join" {
  type    = string
  default = "100.64.0.0/16"
}

variable "kube_ovn_version" {
  type    = string
  default = "1.12.31" # adjust as needed

}

variable "kube_ovn_chart_name" {
  type    = string
  default = "kube-ovn"
}

variable "kube_ovn_chart_repo" {
  type    = string
  default = "https://kubeovn.github.io/kube-ovn"

}