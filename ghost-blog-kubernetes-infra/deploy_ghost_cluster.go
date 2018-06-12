/*
Create a virtual machine in cloud and setup single node Kubernetes cluster
*/
package main

import (
  "os"
  "log"
  "os/exec"
)

func main() {
  log.Printf("Deploying, wait for few min...")

  //parameters = append(parameters, "-var \"do_token=$DO_TOKEN\"")
  //parameters = append(parameters, "-var \"pub_key=$HOME/.ssh/id_rsa.pub\"")
  //parameters = append(parameters, "-var \"pvt_key=$HOME/.ssh/id_rsa\"")
  //parameters = append(parameters, "-var \"ssh_fingerprint=$SSH_FINGERPRINT\"")

  os.Setenv("TF_VAR_do_token", os.Getenv("DO_TOKEN"))
  os.Setenv("TF_VAR_pub_key", os.Getenv("HOME") + "/.ssh/id_rsa.pub")
  os.Setenv("TF_VAR_pvt_key", os.Getenv("HOME") + "/.ssh/id_rsa")

  //ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub
  os.Setenv("TF_VAR_ssh_fingerprint", os.Getenv("SSH_FINGERPRINT"))

  get_version()
  create_virtual_machine()
  //setup_kubernetes_cluster()

  //delete_virtual_machine()
  //log.Printf(parameters)

}

// get terraform version
func get_version(){
  execute_command("terraform", []string{"version"})
}

// create virual machine, change parameters in tf file
func create_virtual_machine(){
  parameters := make([]string, 0)
  parameters = append(parameters, "apply")
  parameters = append(parameters, "--auto-approve")

  execute_command("terraform", parameters)
}

// create Kubernetes cluster
func setup_kubernetes_cluster() {
  parameters := make([]string, 0)
  parameters = append(parameters, "playbook.yml")
  execute_command("ansible-playbook", parameters)
}

// delete virtual machine
func delete_virtual_machine(){
  parameters := make([]string, 0)
  parameters = append(parameters, "destroy")
  parameters = append(parameters, "--force")

  execute_command("terraform", parameters)
}

func execute_command(cmd string, param []string){
  //log.Printf("Executing command = " + cmd + " " + subcmd + " " + param)
  log.Println(param)
  var output, err = exec.Command(cmd, param...).CombinedOutput()
  if err != nil {
    log.Printf("Command finished with error " + string(output))
    log.Fatal(err)
  } else {
    log.Printf("Command finished with success " + string(output))
  }
}
