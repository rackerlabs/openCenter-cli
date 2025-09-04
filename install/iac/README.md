Kubespray Provider Configuration Guide

We use OpenTofu to deploy virtual machines on OpenStack, using the outputs of the infra module, we generate the Kubespray YAML manifests, then run the Kubespray playbooks to deploy Kubernetes on the virtual machines.

**NOTE:** These steps will change once automation is built.

The terraform state file will be stored in an S3 bucket

# Steps to deploy

## Pre Requisites
- Requires-Python >=3.10
- python3.10-venv
- Terraform >=v1.11.1
- kubectl
```
# Ubuntu 24.04
apt install unzip -y
apt install make -y

```


**NOTE:** The brew installed Python works for Macs `brew install python@3.10`
### Create S3 Bucket
- Give it a unique name
- Enable `Block all public access`
- Enable Bucket Versioning
- Add Tags to know if this is important or not: `production` or `dev`
- Default Encryption: Server-side encryption with Amazon S3 managed keys

### Create access policy for Bucket

- In the Resource URN replace BUCKET_NAME with the name of the S3 bucket from the previous step
- This policy will allow a single account to access the OpenTofu state file of multiple clusters by allowing a directory structure:
```
├── BUCKET_NAME
│   ├── CLUSTER_NAME
│   │   ├── tfstate
│   │   |	└── terraform.tfstate
│   │   |	└── terraform.tfstate.tflock
```


IAM Policy:

``` json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "s3:ListBucket",
            "Resource": "arn:aws:s3:::BUCKET_NAME"
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:PutObject"
            ],
            "Resource": [
                "arn:aws:s3:::BUCKET_NAME/*/tfstate/terraform.tfstate"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:PutObject",
                "s3:DeleteObject"
            ],
            "Resource": [
                "arn:aws:s3:::BUCKET_NAME/*/tfstate/terraform.tfstate.tflock"
            ]
        }
    ]
}
```


### Create AWS User
- Give it a clear name like "customer name".
- Leave console access unchecked.
- Attach policies directly and pick the policy created above.
- Take note of the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY of the new user.

## Configure IaC files

The starting point is to copy the init directory into the new clusters directory

```

├── infrastructure
│   └── init
│       ├── main.tf
│       ├── provider.tf
│       └── variables.tf
│       └── Makefile
│   └── clusters
│       ├── demo-cluster
│       └── production
```


### Copy the base OpenTofu files to the new directory

```
# cd /etc/openCenter
# cp -r infrastructure/init infrastructure/clusters/demo-cluster
# cd infrastructure/clusters/demo-cluster
```

### Configure the OpenTofu files

**provider.tf**

Update the BUCKET_NAME wiht the S3 bucket and CLUSTER_NAME with the unique cluster name.

```
terraform {
  backend "s3" {
    bucket       = "BUCKET_NAME"
    key          = "CLUSTER_NAME/tfstate/terraform.tfstate"
    region       = "us-west-2"
    use_lockfile = true
    encrypt      = true
  }
}
```

**main.tf**

**NOTE:** The main.tf in base is configured to work on Openstack Flex in SJC3 it will use Kube-VIP with a floating IP associated to expose the Kubernetes API publicly. When deploying in another cloud more settings will need to be updated.

These are the minimum required changes;  further configuration options are documented later in this document

Replace the cluster and Tenant names accordignly.
Pick a CIDR that is available for the servers or VMs and reokace it in the subnet_nodes.

```
locals {

  cluster_name                            = "CLUSTER_NAME"
  openstack_tenant_name                   = "TENANT_NAME"
  #CIDR that the openstack VMs will use for K8s nodes
  subnet_nodes                            = "10.2.188.0/22"
  # ==================================== 
  #Kubespray Settings
  kubespray_version                       = "v2.28.1"
  kubernetes_version                      = "1.32.5"
}
```

### Export credentials

Export the openstack and S3 credentials
```

export TF_VAR_openstack_user_password='api-key'
export TF_VAR_openstack_user_name='fanatiguy@rackspace.com'
export AWS_ACCESS_KEY_ID=<KEY>
export AWS_SECRET_ACCESS_KEY=<KEY>
```

## Install Terraform binary
From within each cluster directory we want to install in a local .bin directory as there may be a case where newer clusters get deployed with a much newer and incompatible version.

`make terraform`

## Add terraform to PATH
```
export BIN=${PWD]/.bin
export PATH=${BIN}:${PATH}

```

## Deploy IAC

```
# terraform init
```
The terraform init needs to access modules in git which can be done with SSH keys or a Git Token.
If you want to use the SSH Key method each module source will use: `git@github.com:rackerlabs/openCenter.git`
For Token `github.com/rackerlabs/openCenter.git`

If the init succeeds you are good to apply

```
# terraform apply
```

## Use the cluster

### Install the kubectl binary
`make kubectl`

```
export KUBECONFIG=${PWD}/kubeconfig.yaml

kubectl get nodes
```

# Infrastructure

| Key | Type | Default | Description |
| --- | --- | --- | --- |
| cluster_name | string | ""  | Sets the name of the Cluster, Openstack Project and User. |
| naming_prefix | string | ""  | Prefix to add to Resource Names |
| openstack_auth_url | string | ""  | Openstack Keystone Endpoint |
| openstack_insecure | bool | false | Trust self-signed SSL certificates. |
| openstack_region | string | "RegionOne" | The region of the OpenStack cloud to use. |
| openstack_user_name | string | ""  | The Openstack Username to create. |
| openstack_user_password | string | ""  | The password to set on the Openstack Username. |
| openstack_admin_password | string | ""  | The password of the Openstack Administrator account. |
| openstack_project_domain_name | string | ""  | The openstack project domain name. |
| openstack_user_domain_name | string | ""  | The openstack user domain name. |
| openstack_tenant_name | string | ""  | The openstack tenant name if it already exists. |
| availability_zone | string | ""  | OpenStack availability zone for resource placement |
| floatingip_pool | string | ""  | The name of the floating IP pool to use for external access. |
| router_external_network_id | string | ""  | The UUID of the openstack network to attach to the router for external access. |
| disable_bastion | bool | false | To disable the bastion set it true. Will open port 22 on nodes |
| vlan_id | string | ""  | If set, it will create a VLAN network for the node network. |
| vlan_mtu | string | "1440" | MTU for the VLAN. If VxLAN it will use the environment's default. |
| network_provider | string | ""  | Network provider for the VLAN network interface. |
| subnet_nodes | string | "10.0.0.0/16" | CIDR for Openstack Network for nodes. |
| allocation_pool_start | string | ""  | Start IP of the DHCP allocation IPs of the subnet_nodes network. |
| allocation_pool_end | string | ""  | End IP of the DHCP allocation IPs of the subnet_nodes network. |
| vrrp_ip | string | ""  | Must be an IP from subnet_nodes and will be used as the internal Kubernetes API VIP. |
| subnet_services | string | "10.43.0.0/16" | CIDR to use for Kubernetes services. |
| subnet_pods | string | "10.42.0.0/16" | CIDR to use for Kubernetes pods. |
| use_octavia | bool | true | Use Octavia Load Balancer for Kubernetes API. If False, the vrrp_ip will se used with keepalived. |
| loadbalancer_provider | string | "amphora" | Openstack Octavia loadbalancer provider. |
| vrrp_enabled | bool | ""  | Will use vrrp_ip as the vip to be used with kube-vip. cannot be set to true if use_octavia is true |
| use_designate | bool | true | Creates a DNS record using the LB floating IP and dns_zone_name |
| dns_zone_name | string | ""  | dns_zone_name is the dns zone to create if use_designate is true. The k8s.dns_zone_name record will be added to the Kubernetes api SSL cert. |
| dns_nameservers | list(string) | \["8.8.8.8", "8.8.4.4"\] | DNS servers to configure on the nodes |
| image_id | string | ""  | Glance ImageID for the node Operating System |
| image_id_windows | string | ""  | Glance ImageID for the node Windows Operating System |
| worker_count | number | 0   | Number of worker node VMs to build |
| worker_count_windows | number | 0   | Number of Windows worker node VMs to build |
| master_count | number | 0   | Number of master node VMs to build |
| node_master | string | "master" | Define the role to customize the hostname for eg. the default role is master |
| node_worker | string | "worker" | Define the role to customize the hostname for eg. the default role is worker |
| node_windows | string | "win_wn" | Define the role to customize the hostname for eg. the default role is windows |
| master_node_bfv_size | number | 100 | boot from volume size for the master nodes |
| master_node_bfv_type | string | local | boot from volume type for the master nodes |
| worker_node_bfv_size | number | 100 | boot from volume size for the worker nodes |
| worker_node_bfv_type | string | local | boot from volume type for the worker nodes |
| ssh_user | string | "ubuntu" | Username with SSH access |
| openstack_ca | string | ""  | Signing CA certificate for TLS Openstack Endpoints |
| ca_certificates | string | ""  | Certificates to add to the node OS trusts |
| k8s_api_port | number | 443 | Port number for Kubernetes API server |
| flavor_bastion | string | ""  | OpenStack flavor for bastion host instances |
| flavor_master | string | ""  | OpenStack flavor for master node instances |
| flavor_worker | string | ""  | OpenStack flavor for worker node instances |
| flavor_worker_windows | string | ""  | OpenStack flavor for Windows worker node instances |
| ssh_authorized_keys | list(string) | \[\] | List of SSH public keys for cluster access |
| ub_version | string | "20" | Ubuntu version to use for nodes |
| windows_user | string | "Administrator" | Username for Windows nodes |
| windows_admin_password | string | ""  | Administrator password for Windows nodes |
| worker_node_bfv_size_windows | number | 0   | Boot from volume size for Windows worker nodes |
| worker_node_bfv_type_windows | string | "local" | Boot from volume type for Windows worker nodes |
