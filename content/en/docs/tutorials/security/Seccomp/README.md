# Pre requisites
* Locally or cluster you fully control (=== configure control plane), where you can
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster tool
    * [minikube](https://minikube.sigs.k8s.io/docs/start/)
    * [kind](https://kind.sigs.k8s.io/)
  * curl
  * kubectl > v1.28
    * `kubectl get nodes -o=jsonpath=$'{range .items[*]}{@.metadata.name}: {@.status.nodeInfo.kubeletVersion}\n{end}'`
      * Check it in the current context's nodes
  * Linux environment
    * Reason: AppArmor is just valid in Linux

# Steps
* Go in this repo to "/content/en/examples/pods/security/seccomp/profiles"
* Follow "Run kind cluster with seccomp profiles" notes