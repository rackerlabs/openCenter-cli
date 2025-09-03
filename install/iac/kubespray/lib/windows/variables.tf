variable "deploy_cluster" {
  type    = bool
  default = false
}

variable "calico_version" {
  type    = string
  default = "3.29.4" # adjust as needed
}

variable "chart_name" {
  type    = string
  default = "calico"
}

variable "chart_namespace" {
  type    = string
  default = "tigera-operator"
}

variable "chart_repo" {
  type    = string
  default = "https://docs.tigera.io/calico/charts"
}

variable "cni_iface" {
  type    = string
  default = ""
}

variable "subnet_pods" {
  type    = string
  default = ""
}

variable "subnet_services" {
  type    = string
  default = ""
}

variable "windows_dataplane" {
  type    = string
  default = "HNS"
}