# Manual install




## Pre Requisites

### Openstack Flex
- An existing tenant account
- A username with API Key
- Ubuntu 24.04 Image
- Define a private /22 CIDR for Server network. Default 10.2.188.0/22
- Create S3 Bucket
- A IAM user with a policy that allows access to the S3 Bucket
- Access Keys for the IAM user
### Openstack Flex Outcome
Virtual machines in a private network with a bastion node for SSH access.


## Create a Kubernetes cluster 


### 1. Terraform boostrap 
   1. s3 backend
   2. ..


### 2. Deploy - CNI
  -  helm etc?

### 3 Finalize k8s harden 
   - CSR Approver

### 4 Install Flux


### 5 Install core services


