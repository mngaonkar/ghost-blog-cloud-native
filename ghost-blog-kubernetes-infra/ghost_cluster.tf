variable "do_token" {}
variable "pub_key" {}
variable "pvt_key" {}
variable "ssh_fingerprint" {}

provider "digitalocean" {
  token = "${var.do_token}"
}

resource "digitalocean_droplet" "ghost_master" {
  image = "ubuntu-16-04-x64"
  name = "ghost-master"
  region = "sfo2"
  size = "2gb"
  private_networking = true
  ssh_keys = [
    "${var.ssh_fingerprint}"
  ]
  connection {
       user = "root"
       type = "ssh"
       private_key = "${file(var.pvt_key)}"
       timeout = "2m"
   }
}
