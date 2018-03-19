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
Run go script. The terraform infrastructure as code is defined in file [ghost_cluster.tf](https://github.com/mngaonkar/ghost-blog-cloud-native/blob/master/ghost-blog-kubernetes-infra/ghost_cluster.tf)
```
go run deploy_ghost_cluster.go 
```

## Install Kubernetes by running Ansible playbook
```
ansible-playbook playbook.yml
```

