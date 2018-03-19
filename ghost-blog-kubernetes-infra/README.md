## Prerequisite
Set DigitalOcean token environment variable
```
export DO_TOKEN=<your token here>
```

Make sure public and private keys are generated for local machine. The kyes are located at ~/.ssh/id_rsa as private and ~/.ssh/id_rsa.pub as public keys. If the keys are not found, generate the key pair.
```
ssh-keygen -t rsa
```

## Create a virtual machine in cloud
Run go script. The terraform infrastructure as code is defined in file [ghost_cluster.tf](https://github.com/mngaonkar/ghost-blog-cloud-native/blob/master/ghost-blog-kubernetes-infra/ghost_cluster.tf). This command will create a virtual machine in cloud. 
```
go run deploy_ghost_cluster.go 
```
Once the virtual machine is created, update [hosts](https://github.com/mngaonkar/ghost-blog-cloud-native/blob/master/ghost-blog-kubernetes-infra/hosts) file with public IP address of virtual machine.

## Create Kubernetes cluster
Install kubernetes by running Ansible playbook
```
ansible-playbook playbook.yml
```
Once the cluster is up and running. Copy /etc/kubernetes/admin.conf as ~/.kube/config in virutal machine. Check the cluster state by running kubectl command in virtual machine.
```
kubectl get nodes
```
Copy kubernetes config from virtual machine to local machine. Set environment variable KUBECONFIG to point to config file. Now you can run kubectl command from local machine to manage kubernetes cluster.
