# Host your Ghost blog Cloud Native way

There are lots of way to host your blog on internet. Broadly categorized in managed hosting and self hosting. This tutorial is all about self hosting, where you will deploy the blog hard way. The reason is obvious, great learning and full control.

## Selecting a Cloud Provider
The important criteria for choosing a particular cloud provider is cost. As a developer, one of the great recommendation is DigitalOcean. A virtual machine having 1 vCPU, 2GB memory and 40 GB of disk would be a good choice.

## Why Cloud Native?
The biggest driver for going cloud native is vendor agnostic. Most people want to have that freedom to move their blog to any cloud provider in future seamlessly. Going cloud native also means, moving away from the managed blog service provided by cloud provider. In this case, DigitalOcean does provide Ghost blog service. Using managed service helps to get the blog up and running in no time but provide no flexibility for workload movement. The idea here is to demonstrate the cloud native way so that it can be applied to any application not just a blog.

## Containerization
The first step toward cloud native is to get your application containerized. For Ghost is more straight forward to spin up a container with Ubuntu as base image. Refer https://github.com/mngaonkar/ghost-blog-cloud-native/tree/master/ghost-blog-containerization. The blog application is installed as user admin and uses SQLite as database for storing content.

## Building Kubernetes Infrastructure
The next step is to have a container orchestration platform to deploy the blog container. We will use kubernetes as container orchestrator. A typical kubernetes cluster is recommended to have three nodes, one for master and two for worker. In this example, we will build a single node cluster. This mean a single node will act as master and worker. Note, it is possible to add more workers at later time to the same cluster if needed.  

For infrastructure automation, we will use Terraform (https://www.terraform.io) to spin up a virtual machine in cloud. For host management, we will use Ansible (https://www.ansible.com) to install single node kubernetes cluster.

Refer https://github.com/mngaonkar/ghost-blog-cloud-native/tree/master/ghost-blog-kubernetes-infra
